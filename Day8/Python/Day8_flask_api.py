"""
Day 8 — pytest for Flask API (tests)
===================================

------------------------------------------------------------
INSTALL
------------------------------------------------------------
pip install pytest flask

------------------------------------------------------------
RUN
------------------------------------------------------------
pytest -q

NOTE:
- These tests target the in-memory Flask API in day8_flask_rest_api.py
- If you switch to DB API, you can adapt tests to point at that app instead.
"""

import pytest

# Import the Flask app object from the lab file.
from day8_flask_rest_api import app as flask_app


@pytest.fixture()
def client():
    flask_app.config.update({"TESTING": True})
    with flask_app.test_client() as c:
        yield c


def test_create_and_get_task(client):
    r = client.post("/api/v1/tasks", json={"title": "Write tests"})
    assert r.status_code == 201
    data = r.get_json()
    assert "id" in data
    task_id = data["id"]

    r2 = client.get(f"/api/v1/tasks/{task_id}")
    assert r2.status_code == 200
    data2 = r2.get_json()
    assert data2["title"] == "Write tests"
    assert data2["status"] == "todo"


def test_validation_missing_title(client):
    r = client.post("/api/v1/tasks", json={})
    assert r.status_code == 400
    data = r.get_json()
    assert data["error"]["code"] == "VALIDATION_ERROR"


def test_patch_invalid_status(client):
    r = client.post("/api/v1/tasks", json={"title": "X"})
    task_id = r.get_json()["id"]

    r2 = client.patch(f"/api/v1/tasks/{task_id}", json={"status": "INVALID"})
    assert r2.status_code == 400


def test_pagination(client):
    # create a few tasks
    for i in range(5):
        client.post("/api/v1/tasks", json={"title": f"T{i}"})
    r = client.get("/api/v1/tasks?limit=2&offset=0")
    assert r.status_code == 200
    data = r.get_json()
    assert data["limit"] == 2
    assert len(data["items"]) <= 2
