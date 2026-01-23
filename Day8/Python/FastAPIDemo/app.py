"""Flask Serialization & Deserialization (single-file demo).

This file demonstrates:
1) Serialization: Python objects -> JSON responses
   - Flask's `jsonify` (dict/list -> JSON)
   - Marshmallow `dump` for structured output

2) Deserialization: JSON requests -> Python objects
   - Flask's `request.get_json()` (JSON body -> dict)
   - Marshmallow `load` for validation + object creation

Install:
  python -m pip install Flask marshmallow

Run:
  python app.py

Try:
  # 1) Serialization with jsonify
  curl http://127.0.0.1:5000/data

  # 2) Serialization of a dataclass via Marshmallow
  curl http://127.0.0.1:5000/users

  # 3) Deserialization with request.get_json()
  curl -X POST http://127.0.0.1:5000/submit \
    -H 'Content-Type: application/json' \
    -d '{"name":"Alice","age":25,"city":"New York"}'

  # 4) Deserialization + validation with Marshmallow
  curl -X POST http://127.0.0.1:5000/validate \
    -H 'Content-Type: application/json' \
    -d '{"username":"alice","email":"alice@example.com"}'

  # 5) Create a user (load -> object), then list users (dump)
  curl -X POST http://127.0.0.1:5000/users \
    -H 'Content-Type: application/json' \
    -d '{"username":"bob","email":"bob@example.com","age":30}'
"""

from __future__ import annotations

from dataclasses import dataclass
from typing import List

from flask import Flask, jsonify, request
from marshmallow import Schema, ValidationError, fields, post_load


app = Flask(__name__)


# ---------------------------
# Example 1: Flask jsonify (serialization)
# ---------------------------


@app.get("/data")
def get_data():
    """Serialize a plain Python dict into JSON."""
    data = {"name": "Alice", "age": 25, "city": "New York"}
    return jsonify(data)


# ---------------------------
# Example 2: request.get_json (deserialization)
# ---------------------------


@app.post("/submit")
def submit():
    """Deserialize request JSON into a Python dict."""
    if not request.is_json:
        return jsonify({"error": "Expected application/json"}), 415

    payload = request.get_json(silent=True)
    if payload is None:
        return jsonify({"error": "Invalid JSON"}), 400

    # payload is now a Python dict/list
    return jsonify({"message": "Data received", "data": payload})


# ---------------------------
# Example 3: Marshmallow schema for validation + (de)serialization
# ---------------------------


@dataclass
class User:
    username: str
    email: str
    age: int | None = None


class UserSchema(Schema):
    username = fields.String(required=True)
    email = fields.Email(required=True)
    age = fields.Integer(required=False, allow_none=True, strict=True)

    @post_load
    def make_user(self, data, **kwargs):  # noqa: ANN001
        # Convert validated dict -> User object
        return User(**data)


user_schema = UserSchema()
users_schema = UserSchema(many=True)


# In-memory "database"
USERS: List[User] = [
    User(username="alice", email="alice@example.com", age=25),
    User(username="charlie", email="charlie@example.com"),
]


@app.post("/validate")
def validate_user():
    """Validate incoming JSON against a schema (no persistence)."""
    if not request.is_json:
        return jsonify({"error": "Expected application/json"}), 415

    payload = request.get_json(silent=True)
    if payload is None:
        return jsonify({"error": "Invalid JSON"}), 400

    errors = user_schema.validate(payload)
    if errors:
        return jsonify(errors), 400

    return jsonify({"message": "Valid data", "data": payload})


@app.get("/users")
def list_users():
    """Serialize dataclass objects via schema.dump."""
    return jsonify({"count": len(USERS), "users": users_schema.dump(USERS)})


@app.post("/users")
def create_user():
    """Deserialize + validate JSON into a User object, store it, return serialized."""
    if not request.is_json:
        return jsonify({"error": "Expected application/json"}), 415

    try:
        user: User = user_schema.load(request.get_json())
    except ValidationError as err:
        return jsonify({"errors": err.messages}), 400

    USERS.append(user)
    return jsonify({"message": "Created", "user": user_schema.dump(user)}), 201


# ---------------------------
# Small quality-of-life: consistent error JSON
# ---------------------------


@app.errorhandler(404)
def not_found(_e):  # noqa: ANN001
    return jsonify({"error": "Not found"}), 404


@app.errorhandler(405)
def method_not_allowed(_e):  # noqa: ANN001
    return jsonify({"error": "Method not allowed"}), 405


if __name__ == "__main__":
    # Debug is convenient for learning; turn off in production.
    app.run(debug=True)
