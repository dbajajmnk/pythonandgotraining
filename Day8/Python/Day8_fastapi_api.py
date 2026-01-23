"""
Day 8 — pytest for FastAPI API (tests)
=====================================

------------------------------------------------------------
INSTALL
------------------------------------------------------------
pip install pytest fastapi uvicorn

------------------------------------------------------------
RUN
------------------------------------------------------------
pytest -q

NOTE:
- Targets the in-memory FastAPI API in day8_fastapi_rest_api.py
"""

from fastapi.testclient import TestClient
from day8_fastapi_rest_api import app

client = TestClient(app)


def test_create_and_get_task():
    r = client.post("/api/v1/tasks", json={"title": "Write tests"})
    assert r.status_code == 201
    task = r.json()
    task_id = task["id"]

    r2 = client.get(f"/api/v1/tasks/{task_id}")
    assert r2.status_code == 200
    assert r2.json()["title"] == "Write tests"


def test_validation_missing_title():
    r = client.post("/api/v1/tasks", json={})
    # FastAPI returns 422 for validation errors by default
    assert r.status_code == 422


def test_patch_invalid_status():
    r = client.post("/api/v1/tasks", json={"title": "X"})
    task_id = r.json()["id"]
    r2 = client.patch(f"/api/v1/tasks/{task_id}", json={"status": "INVALID"})
    assert r2.status_code == 400


def test_pagination():
    for i in range(5):
        client.post("/api/v1/tasks", json={"title": f"T{i}"})
    r = client.get("/api/v1/tasks?limit=2&offset=0")
    assert r.status_code == 200
    data = r.json()
    assert data["limit"] == 2
    assert len(data["items"]) <= 2
