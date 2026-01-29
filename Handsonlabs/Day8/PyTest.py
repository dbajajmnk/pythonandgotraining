import pytest

# Import the Flask app object from the lab file.
from Flaskdemo import app as flask_app
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
