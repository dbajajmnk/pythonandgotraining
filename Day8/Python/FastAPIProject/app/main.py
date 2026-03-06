from fastapi import FastAPI
from app.core.middleware import setup_middleware
from app.core.exceptions import register_exception_handlers
from app.api.routes import router

app = FastAPI(title="Production FastAPI Project")

setup_middleware(app)
register_exception_handlers(app)

app.include_router(router)
