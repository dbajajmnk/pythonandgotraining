from typing import Optional
from pydantic import BaseModel, EmailStr, Field

class UserCreate(BaseModel):
    email: EmailStr
    password: str = Field(min_length=8)

class UserUpdate(BaseModel):
    email: Optional[EmailStr] = None
    # password update optional (demo). In real app: separate flow.
    password: Optional[str] = Field(default=None, min_length=8)

class UserOut(BaseModel):
    id: int
    email: EmailStr

    class Config:
        from_attributes = True
