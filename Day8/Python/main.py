"""
Day 8 — Session 1B: REST API Basics with FastAPI (In-Depth + Lab)
================================================================

------------------------------------------------------------
INSTALL
------------------------------------------------------------
pip install fastapi uvicorn

------------------------------------------------------------
RUN
------------------------------------------------------------
uvicorn day8_fastapi_rest_api:app --reload

Open docs:
- http://127.0.0.1:8000/docs
"""

from __future__ import annotations

from fastapi import FastAPI, HTTPException, Query
from pydantic import BaseModel, Field
from datetime import datetime
from uuid import uuid4, UUID

# ============================================================
# WHAT/WHY FastAPI:
# - Type hints + Pydantic give automatic validation + OpenAPI docs.
# - Great for contract-driven APIs and teams.
# ============================================================

app = FastAPI(title="Task Tracker API", version="1.0.0")

# In-memory storage for training (same warning as Flask).
TASKS: dict[UUID, dict] = {}

# -----------------------------
# Request/Response Models
# -----------------------------
class TaskIn(BaseModel):
    title: str = Field(min_length=1, description="Task title")

class TaskPatch(BaseModel):
    title: str | None = Field(default=None, min_length=1)
    status: str | None = Field(default=None, description="todo|doing|done")

class TaskOut(BaseModel):
    id: UUID
    title: str
    status: str
    createdAt: datetime
    updatedAt: datetime

class ListEnvelope(BaseModel):
    items: list[TaskOut]
    total: int
    limit: int
    offset: int

def to_out(task: dict) -> TaskOut:
    return TaskOut(**task)

# -----------------------------
# Routes
# -----------------------------
@app.get("/api/v1/tasks", response_model=ListEnvelope)
def list_tasks(
    status: str | None = None,
    limit: int = Query(default=50, ge=1, le=200),
    offset: int = Query(default=0, ge=0),
):
    tasks = list(TASKS.values())
    if status:
        tasks = [t for t in tasks if t["status"] == status]

    paged = tasks[offset: offset + limit]
    return ListEnvelope(
        items=[to_out(t) for t in paged],
        total=len(tasks),
        limit=limit,
        offset=offset,
    )

@app.get("/api/v1/tasks/{task_id}", response_model=TaskOut)
def get_task(task_id: UUID):
    task = TASKS.get(task_id)
    if not task:
        raise HTTPException(status_code=404, detail="Task not found")
    return to_out(task)

@app.post("/api/v1/tasks", response_model=TaskOut, status_code=201)
def create_task(payload: TaskIn):
    now = datetime.utcnow()
    task_id = uuid4()
    task = {
        "id": task_id,
        "title": payload.title.strip(),
        "status": "todo",
        "createdAt": now,
        "updatedAt": now,
    }
    TASKS[task_id] = task
    return to_out(task)

@app.patch("/api/v1/tasks/{task_id}", response_model=TaskOut)
def patch_task(task_id: UUID, payload: TaskPatch):
    task = TASKS.get(task_id)
    if not task:
        raise HTTPException(status_code=404, detail="Task not found")

    if payload.title is not None:
        task["title"] = payload.title.strip()
    if payload.status is not None:
        if payload.status not in {"todo", "doing", "done"}:
            raise HTTPException(status_code=400, detail="Invalid status")
        task["status"] = payload.status

    task["updatedAt"] = datetime.utcnow()
    return to_out(task)

@app.delete("/api/v1/tasks/{task_id}", status_code=204)
def delete_task(task_id: UUID):
    if task_id not in TASKS:
        raise HTTPException(status_code=404, detail="Task not found")
    del TASKS[task_id]
    return None
