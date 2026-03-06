from typing import Dict, List
from fastapi import HTTPException
from app.schemas.user import UserCreate, UserUpdate

# Simple in-memory store (for demo)
# NOTE: In production, replace with DB (SQLAlchemy) repo layer.
_USERS: Dict[int, dict] = {}
_NEXT_ID = 1

def create_user(payload: UserCreate) -> dict:
    global _NEXT_ID

    # Duplicate email check
    for u in _USERS.values():
        if u["email"] == str(payload.email):
            raise HTTPException(status_code=409, detail="Email already exists")

    user = {"id": _NEXT_ID, "email": str(payload.email)}
    _USERS[_NEXT_ID] = user
    _NEXT_ID += 1
    return user

def list_users() -> List[dict]:
    return list(_USERS.values())

def get_user(user_id: int) -> dict:
    user = _USERS.get(user_id)
    if not user:
        raise HTTPException(status_code=404, detail="User not found")
    return user

def update_user(user_id: int, payload: UserUpdate) -> dict:
    user = get_user(user_id)

    if payload.email is not None:
        # Duplicate email check (excluding self)
        for uid, u in _USERS.items():
            if uid != user_id and u["email"] == str(payload.email):
                raise HTTPException(status_code=409, detail="Email already exists")
        user["email"] = str(payload.email)

    _USERS[user_id] = user
    return user

def delete_user(user_id: int) -> None:
    _ = get_user(user_id)
    del _USERS[user_id]