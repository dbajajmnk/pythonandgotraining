from fastapi.testclient import TestClient
from app.main import app

client = TestClient(app)

def test_create_user():
    response = client.post("/users/", json={"email": "test@example.com", "password": "Strong@123"})
    assert response.status_code == 200
