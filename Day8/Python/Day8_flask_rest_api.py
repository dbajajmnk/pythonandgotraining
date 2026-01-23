"""
Day 8 — Session 1A: REST API Basics with Flask (In-Depth + Lab)
==============================================================

This file is COMMENT-HEAVY for teaching (What/Why/When + mental mapping).

------------------------------------------------------------
INSTALL
------------------------------------------------------------
pip install flask

------------------------------------------------------------
RUN
------------------------------------------------------------
python day8_flask_rest_api.py

Then open:
- http://127.0.0.1:5000/api/v1/tasks

Try with curl:
- curl -s http://127.0.0.1:5000/api/v1/tasks | python -m json.tool
- curl -s -X POST http://127.0.0.1:5000/api/v1/tasks -H "Content-Type: application/json" -d '{"title":"Buy milk"}'
"""

from __future__ import annotations

from flask import Flask, request, jsonify
from datetime import datetime
from uuid import uuid4

# ============================================================
# WHAT: REST API mental model
# ============================================================
# Resource = noun (tasks). HTTP method = verb (GET/POST/PATCH/DELETE).
# - GET    /tasks        -> list
# - GET    /tasks/{id}   -> read one
# - POST   /tasks        -> create
# - PATCH  /tasks/{id}   -> partial update
# - DELETE /tasks/{id}   -> delete
#
# WHY: Predictable, testable, easy for clients, standard tooling.
# WHEN: Internal services, mobile apps, partner integrations.

app = Flask(__name__)

# In-memory storage for training.
# NOTE: In production, you would use a database. In-memory resets on restart,
# and is NOT safe across multiple worker processes.
TASKS: dict[str, dict] = {}

# -----------------------------
# Helpers: consistent errors
# -----------------------------
def error_response(status: int, code: str, message: str, details: dict | None = None):
    """
    WHY: Clients need stable error formats. Status code communicates the class of error.
    """
    payload = {"error": {"code": code, "message": message, "details": details or {}}}
    return jsonify(payload), status

def serialize_task(task: dict) -> dict:
    """
    WHY: JSON can’t serialize datetime automatically.
    Convert datetime to ISO string to be safe/stable.
    """
    return {
        "id": task["id"],
        "title": task["title"],
        "status": task["status"],
        "createdAt": task["createdAt"].isoformat(),
        "updatedAt": task["updatedAt"].isoformat(),
    }

def validate_task_create(data: dict | None):
    """
    WHEN: Always validate user input at the edge (API boundary).
    WHY: Prevent bad data from entering your system; reduce downstream bugs.
    """
    if not isinstance(data, dict):
        return "Body must be JSON object"
    title = data.get("title")
    if not isinstance(title, str) or not title.strip():
        return "Field 'title' is required and must be a non-empty string"
    return None

def validate_task_patch(data: dict | None):
    if not isinstance(data, dict):
        return "Body must be JSON object"
    if "title" in data and (not isinstance(data["title"], str) or not data["title"].strip()):
        return "Field 'title' must be a non-empty string"
    if "status" in data and data["status"] not in {"todo", "doing", "done"}:
        return "Field 'status' must be one of: todo, doing, done"
    return None

# -----------------------------
# Routes
# -----------------------------
@app.get("/api/v1/tasks")
def list_tasks():
    """
    Supports:
    - filtering: ?status=done
    - pagination: ?limit=10&offset=0
    """
    status = request.args.get("status")
    limit = request.args.get("limit", type=int) or 50
    offset = request.args.get("offset", type=int) or 0

    tasks = list(TASKS.values())
    if status:
        tasks = [t for t in tasks if t["status"] == status]

    # Pagination (simple offset/limit)
    paged = tasks[offset: offset + limit]

    return jsonify({
        "items": [serialize_task(t) for t in paged],
        "total": len(tasks),
        "limit": limit,
        "offset": offset
    }), 200

@app.get("/api/v1/tasks/<task_id>")
def get_task(task_id: str):
    task = TASKS.get(task_id)
    if not task:
        return error_response(404, "TASK_NOT_FOUND", f"Task '{task_id}' not found")
    return jsonify(serialize_task(task)), 200

@app.post("/api/v1/tasks")
def create_task():
    data = request.get_json(silent=True)
    err = validate_task_create(data)
    if err:
        return error_response(400, "VALIDATION_ERROR", err)

    now = datetime.utcnow()
    task_id = str(uuid4())
    task = {
        "id": task_id,
        "title": data["title"].strip(),
        "status": "todo",
        "createdAt": now,
        "updatedAt": now,
    }
    TASKS[task_id] = task

    # 201 Created is the REST-friendly status for creation
    return jsonify(serialize_task(task)), 201

@app.patch("/api/v1/tasks/<task_id>")
def patch_task(task_id: str):
    task = TASKS.get(task_id)
    if not task:
        return error_response(404, "TASK_NOT_FOUND", f"Task '{task_id}' not found")

    data = request.get_json(silent=True)
    err = validate_task_patch(data)
    if err:
        return error_response(400, "VALIDATION_ERROR", err)

    if "title" in data:
        task["title"] = data["title"].strip()
    if "status" in data:
        task["status"] = data["status"]

    task["updatedAt"] = datetime.utcnow()
    return jsonify(serialize_task(task)), 200

@app.delete("/api/v1/tasks/<task_id>")
def delete_task(task_id: str):
    if task_id not in TASKS:
        return error_response(404, "TASK_NOT_FOUND", f"Task '{task_id}' not found")
    del TASKS[task_id]
    # 204: no body returned
    return "", 204


if __name__ == "__main__":
    # Debug only. Production uses a WSGI server like gunicorn/uwsgi.
    app.run(host="127.0.0.1", port=5000, debug=True)
