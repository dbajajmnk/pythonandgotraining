from typing import List
from fastapi import APIRouter, status
from app.schemas.user import UserCreate, UserUpdate, UserOut
from app.services.user_service import create_user, list_users, get_user, update_user, delete_user

router = APIRouter()

@router.post("/", response_model=UserOut, status_code=status.HTTP_201_CREATED)
def create(payload: UserCreate):
    return create_user(payload)

@router.get("/", response_model=List[UserOut])
def read_all():
    return list_users()

@router.get("/{user_id}", response_model=UserOut)
def read_one(user_id: int):
    return get_user(user_id)

@router.patch("/{user_id}", response_model=UserOut)
def patch_one(user_id: int, payload: UserUpdate):
    return update_user(user_id, payload)

@router.delete("/{user_id}", status_code=status.HTTP_204_NO_CONTENT)
def remove_one(user_id: int):
    delete_user(user_id)
    return None