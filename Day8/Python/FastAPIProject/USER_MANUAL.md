# FastAPI Production Project Manual

## What
Production-ready FastAPI structure covering:
- Pydantic validation
- CRUD
- JWT auth
- Middleware
- Exception handling
- Testing
- Serving concepts

## Why
This structure separates concerns for scalability and maintainability.

## How to Run
1. Install requirements:
   pip install -r requirements.txt

2. Run server:
   uvicorn app.main:app --reload

3. Run tests:
   pytest

## Serving
- Development: uvicorn
- Production: gunicorn -k uvicorn.workers.UvicornWorker -w 4 app.main:app
