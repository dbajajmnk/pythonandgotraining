from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from typing import List, Optional

app = FastAPI()

# 1. The Schema (Data Validation)
class Task(BaseModel):
    id: int
    title: str
    done: bool = False

# In-memory database
tasks_db = []

# 2. GET - Retrieve all
@app.get("/tasks", response_model=List[Task])
async def get_tasks():
    return tasks_db

# 3. POST - Create with automatic validation
@app.post("/tasks", status_code=201)
async def create_task(task: Task):
    tasks_db.append(task)
    return task

# 4. DELETE - Specific resource
@app.delete("/tasks/{task_id}")
async def delete_task(task_id: int):
    global tasks_db
    tasks_db = [t for t in tasks_db if t.id != task_id]
    return {"message": "Task deleted successfully"}