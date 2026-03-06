from fastapi import APIRouter, Depends
from app.schemas.auth import Token
from app.services.auth_service import login

router = APIRouter()

@router.post("/login", response_model=Token)
def login_user(username: str, password: str):
    return login(username, password)
