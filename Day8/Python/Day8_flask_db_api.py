"""
Day 8 — Section 3A: Database Access for APIs with Flask (SQLite/Postgres)
=======================================================================

This file shows an API implementation perspective:
- Connection handling
- Transactions
- Parameterized queries
- Switching SQLite <-> PostgreSQL using DATABASE_URL

------------------------------------------------------------
INSTALL
------------------------------------------------------------
pip install flask sqlalchemy

# Optional PostgreSQL driver:
pip install psycopg2-binary

------------------------------------------------------------
RUN (SQLite default)
------------------------------------------------------------
python day8_flask_db_api.py

------------------------------------------------------------
RUN (PostgreSQL)
------------------------------------------------------------
# Example (edit credentials):
# Linux/macOS:
export DATABASE_URL="postgresql+psycopg2://user:pass@localhost:5432/mydb"
python day8_flask_db_api.py

# Windows PowerShell:
# $env:DATABASE_URL="postgresql+psycopg2://user:pass@localhost:5432/mydb"
# python day8_flask_db_api.py
"""

from __future__ import annotations

import os
from datetime import datetime
from uuid import uuid4

from flask import Flask, request, jsonify
from sqlalchemy import create_engine, text

# ============================================================
# DB CONFIG
# ============================================================
# WHY DATABASE_URL:
# - Same code in dev/prod; only config changes (12-factor app principle)
DATABASE_URL = os.getenv("DATABASE_URL", "sqlite+pysqlite:///./tasks.db")

# create_engine creates a connection pool internally.
# future=True uses SQLAlchemy 2.0 style.
engine = create_engine(DATABASE_URL, future=True)

app = Flask(__name__)

# ============================================================
# DB SCHEMA (simple tasks table)
# ============================================================
# NOTE: Real projects use migrations (Alembic).
SCHEMA_SQL = """
CREATE TABLE IF NOT EXISTS tasks (
  id TEXT PRIMARY KEY,
  title TEXT NOT NULL,
  status TEXT NOT NULL,
  created_at TEXT NOT NULL,
  updated_at TEXT NOT NULL
);
"""

def init_db():
    with engine.begin() as conn:
        conn.execute(text(SCHEMA_SQL))

def error_response(status: int, code: str, message: str, details: dict | None = None):
    return jsonify({"error": {"code": code, "message": message, "details": details or {}}}), status

def row_to_task(r: dict) -> dict:
    return {
        "id": r["id"],
        "title": r["title"],
        "status": r["status"],
        "createdAt": r["created_at"],
        "updatedAt": r["updated_at"],
    }

# ============================================================
# ROUTES
# ============================================================
@app.get("/api/v1/tasks")
def list_tasks():
    status = request.args.get("status")
    limit = request.args.get("limit", type=int) or 50
    offset = request.args.get("offset", type=int) or 0

    sql = "SELECT * FROM tasks"
    params = {}
    if status:
        sql += " WHERE status = :status"
        params["status"] = status
    sql += " ORDER BY created_at DESC LIMIT :limit OFFSET :offset"
    params["limit"] = limit
    params["offset"] = offset

    with engine.connect() as conn:
        rows = conn.execute(text(sql), params).mappings().all()

        # total count (for pagination UI)
        count_sql = "SELECT COUNT(*) AS c FROM tasks"
        if status:
            count_sql += " WHERE status = :status"
        total = conn.execute(text(count_sql), {"status": status} if status else {}).mappings().first()["c"]

    return jsonify({"items": [row_to_task(dict(r)) for r in rows], "total": total, "limit": limit, "offset": offset})

@app.get("/api/v1/tasks/<task_id>")
def get_task(task_id: str):
    with engine.connect() as conn:
        row = conn.execute(text("SELECT * FROM tasks WHERE id=:id"), {"id": task_id}).mappings().first()
    if not row:
        return error_response(404, "TASK_NOT_FOUND", f"Task '{task_id}' not found")
    return jsonify(row_to_task(dict(row)))

@app.post("/api/v1/tasks")
def create_task():
    data = request.get_json(silent=True) or {}
    title = data.get("title")
    if not isinstance(title, str) or not title.strip():
        return error_response(400, "VALIDATION_ERROR", "Field 'title' is required and must be non-empty")

    now = datetime.utcnow().isoformat()
    task_id = str(uuid4())
    task = {"id": task_id, "title": title.strip(), "status": "todo", "created_at": now, "updated_at": now}

    # engine.begin() = transaction (commit on success, rollback on exception)
    with engine.begin() as conn:
        conn.execute(
            text("INSERT INTO tasks (id,title,status,created_at,updated_at) VALUES (:id,:t,:s,:c,:u)"),
            {"id": task["id"], "t": task["title"], "s": task["status"], "c": task["created_at"], "u": task["updated_at"]},
        )

    return jsonify(row_to_task(task)), 201

@app.patch("/api/v1/tasks/<task_id>")
def patch_task(task_id: str):
    data = request.get_json(silent=True) or {}
    if "status" in data and data["status"] not in {"todo", "doing", "done"}:
        return error_response(400, "VALIDATION_ERROR", "status must be one of todo/doing/done")
    if "title" in data and (not isinstance(data["title"], str) or not data["title"].strip()):
        return error_response(400, "VALIDATION_ERROR", "title must be non-empty string")

    # read current
    with engine.connect() as conn:
        current = conn.execute(text("SELECT * FROM tasks WHERE id=:id"), {"id": task_id}).mappings().first()
    if not current:
        return error_response(404, "TASK_NOT_FOUND", f"Task '{task_id}' not found")

    new_title = data.get("title", current["title"])
    new_status = data.get("status", current["status"])
    now = datetime.utcnow().isoformat()

    with engine.begin() as conn:
        conn.execute(
            text("UPDATE tasks SET title=:t, status=:s, updated_at=:u WHERE id=:id"),
            {"t": new_title.strip(), "s": new_status, "u": now, "id": task_id},
        )

    return jsonify({"id": task_id, "title": new_title.strip(), "status": new_status, "createdAt": current["created_at"], "updatedAt": now})

@app.delete("/api/v1/tasks/<task_id>")
def delete_task(task_id: str):
    with engine.begin() as conn:
        res = conn.execute(text("DELETE FROM tasks WHERE id=:id"), {"id": task_id})
    if res.rowcount == 0:
        return error_response(404, "TASK_NOT_FOUND", f"Task '{task_id}' not found")
    return "", 204

@app.get("/health/db")
def health_db():
    try:
        with engine.connect() as conn:
            conn.execute(text("SELECT 1"))
        return jsonify({"status": "ok"})
    except Exception as e:
        return error_response(500, "DB_ERROR", "Database not reachable", {"reason": str(e)})


if __name__ == "__main__":
    init_db()
    app.run(host="127.0.0.1", port=5001, debug=True)
