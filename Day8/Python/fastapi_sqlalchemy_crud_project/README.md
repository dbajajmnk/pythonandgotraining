
# FastAPI + SQLAlchemy CRUD Project

## Setup
pip install -r requirements.txt

## Run Server
uvicorn main:app --reload

## API Docs
http://127.0.0.1:8000/docs

## Test with Curl

Create Book
curl -X POST http://127.0.0.1:8000/books -H "Content-Type: application/json" -d '{"title":"AI Book","author":"Deepak","price":20}'

Get Books
curl http://127.0.0.1:8000/books
