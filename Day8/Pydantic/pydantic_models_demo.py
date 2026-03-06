"""
Pydantic + FastAPI single-file teaching example
Topic: request models, response models, validation, custom validators

Run:
    pip install fastapi uvicorn pydantic[email]
    uvicorn pydantic_models_demo:app --reload

Open:
    http://127.0.0.1:8000/docs
"""

from __future__ import annotations

from datetime import datetime
from decimal import Decimal
from typing import Dict, List, Literal, Optional
from uuid import uuid4

from fastapi import FastAPI, HTTPException, Path, Query
from pydantic import BaseModel, ConfigDict, EmailStr, Field, field_validator, model_validator

app = FastAPI(
    title="Pydantic Request/Response Models Demo",
    version="1.0.0",
    description=(
        "Single-source example that demonstrates request models, response models, "
        "built-in validation, and custom validators using FastAPI + Pydantic v2."
    ),
)


# -----------------------------------------------------------------------------
# In-memory store
# -----------------------------------------------------------------------------
FAKE_DB: Dict[str, dict] = {}


# -----------------------------------------------------------------------------
# Shared base models
# -----------------------------------------------------------------------------
class Address(BaseModel):
    """Nested model to show validation for child objects."""

    city: str = Field(min_length=2, max_length=60, description="City name")
    state: str = Field(min_length=2, max_length=60, description="State name")
    postal_code: str = Field(min_length=4, max_length=10, description="Postal code")

    @field_validator("city", "state")
    @classmethod
    def normalize_title_case(cls, value: str) -> str:
        """Trim spaces and make values easier to store consistently."""
        return value.strip().title()

    @field_validator("postal_code")
    @classmethod
    def postal_code_should_be_alnum(cls, value: str) -> str:
        cleaned = value.strip().replace(" ", "")
        if not cleaned.isalnum():
            raise ValueError("postal_code must contain only letters and numbers")
        return cleaned.upper()


class StudentBase(BaseModel):
    """Base shape shared by request and response models."""

    model_config = ConfigDict(extra="forbid", str_strip_whitespace=True)

    full_name: str = Field(
        min_length=3,
        max_length=80,
        description="Student full name",
        examples=["Deepak Bajaj"],
    )
    email: EmailStr
    age: int = Field(ge=18, le=60, description="Allowed age range: 18 to 60")
    course: Literal["fastapi", "python", "backend", "ai"]
    fees_paid: Decimal = Field(gt=0, description="Course fee must be greater than 0")
    discount: Decimal = Field(ge=0, description="Discount cannot be negative")
    skills: List[str] = Field(min_length=1, description="At least one skill is required")
    address: Address
    phone: Optional[str] = Field(default=None, description="10-digit contact number")

    @field_validator("full_name")
    @classmethod
    def full_name_should_have_two_words(cls, value: str) -> str:
        parts = [part for part in value.split(" ") if part]
        if len(parts) < 2:
            raise ValueError("full_name must contain at least first name and last name")
        return " ".join(word.capitalize() for word in parts)

    @field_validator("skills")
    @classmethod
    def skills_should_be_clean_and_unique(cls, value: List[str]) -> List[str]:
        cleaned: List[str] = []
        seen = set()
        for skill in value:
            normalized = skill.strip().lower()
            if not normalized:
                continue
            if normalized not in seen:
                seen.add(normalized)
                cleaned.append(normalized)
        if not cleaned:
            raise ValueError("skills must contain at least one non-empty item")
        return cleaned

    @field_validator("phone")
    @classmethod
    def phone_must_be_ten_digits(cls, value: Optional[str]) -> Optional[str]:
        if value is None:
            return value
        digits = "".join(ch for ch in value if ch.isdigit())
        if len(digits) != 10:
            raise ValueError("phone must contain exactly 10 digits")
        return digits

    @model_validator(mode="after")
    def discount_cannot_exceed_fees(self) -> "StudentBase":
        if self.discount >= self.fees_paid:
            raise ValueError("discount must be smaller than fees_paid")
        return self


# -----------------------------------------------------------------------------
# Request models
# -----------------------------------------------------------------------------
class StudentCreateRequest(StudentBase):
    password: str = Field(min_length=8, description="Login password")
    confirm_password: str = Field(min_length=8, description="Must match password")

    @field_validator("password")
    @classmethod
    def password_strength_rules(cls, value: str) -> str:
        has_upper = any(ch.isupper() for ch in value)
        has_lower = any(ch.islower() for ch in value)
        has_digit = any(ch.isdigit() for ch in value)
        has_special = any(not ch.isalnum() for ch in value)

        if not (has_upper and has_lower and has_digit and has_special):
            raise ValueError(
                "password must include uppercase, lowercase, number, and special character"
            )
        return value

    @model_validator(mode="after")
    def passwords_must_match(self) -> "StudentCreateRequest":
        if self.password != self.confirm_password:
            raise ValueError("password and confirm_password must match")
        return self


class StudentUpdateRequest(BaseModel):
    """Partial update model to show optional fields in request bodies."""

    model_config = ConfigDict(extra="forbid", str_strip_whitespace=True)

    full_name: Optional[str] = Field(default=None, min_length=3, max_length=80)
    age: Optional[int] = Field(default=None, ge=18, le=60)
    phone: Optional[str] = None
    skills: Optional[List[str]] = None

    @field_validator("full_name")
    @classmethod
    def update_name_should_have_two_words(cls, value: Optional[str]) -> Optional[str]:
        if value is None:
            return value
        parts = [part for part in value.split(" ") if part]
        if len(parts) < 2:
            raise ValueError("full_name must contain at least first name and last name")
        return " ".join(word.capitalize() for word in parts)

    @field_validator("phone")
    @classmethod
    def update_phone_must_be_ten_digits(cls, value: Optional[str]) -> Optional[str]:
        if value is None:
            return value
        digits = "".join(ch for ch in value if ch.isdigit())
        if len(digits) != 10:
            raise ValueError("phone must contain exactly 10 digits")
        return digits

    @field_validator("skills")
    @classmethod
    def update_skills_should_be_clean_and_unique(
        cls, value: Optional[List[str]]
    ) -> Optional[List[str]]:
        if value is None:
            return value
        cleaned: List[str] = []
        seen = set()
        for skill in value:
            normalized = skill.strip().lower()
            if normalized and normalized not in seen:
                seen.add(normalized)
                cleaned.append(normalized)
        if not cleaned:
            raise ValueError("skills must contain at least one non-empty item")
        return cleaned


# -----------------------------------------------------------------------------
# Response models
# -----------------------------------------------------------------------------
class StudentResponse(BaseModel):
    id: str
    full_name: str
    email: EmailStr
    age: int
    course: str
    fees_paid: Decimal
    discount: Decimal
    net_fees: Decimal
    skills: List[str]
    address: Address
    phone: Optional[str]
    is_adult: bool
    created_at: datetime


class StudentListResponse(BaseModel):
    total: int
    items: List[StudentResponse]


class ErrorResponse(BaseModel):
    detail: str


# -----------------------------------------------------------------------------
# Helper function
# -----------------------------------------------------------------------------
def make_student_response(record: dict) -> StudentResponse:
    net_fees = Decimal(record["fees_paid"]) - Decimal(record["discount"])
    return StudentResponse(
        id=record["id"],
        full_name=record["full_name"],
        email=record["email"],
        age=record["age"],
        course=record["course"],
        fees_paid=record["fees_paid"],
        discount=record["discount"],
        net_fees=net_fees,
        skills=record["skills"],
        address=record["address"],
        phone=record.get("phone"),
        is_adult=record["age"] >= 18,
        created_at=record["created_at"],
    )


# -----------------------------------------------------------------------------
# Endpoints
# -----------------------------------------------------------------------------
@app.get("/", tags=["Health"])
def home() -> dict:
    return {
        "message": "Pydantic request/response models demo is running",
        "docs": "/docs",
    }


@app.post(
    "/students",
    response_model=StudentResponse,
    status_code=201,
    tags=["Students"],
    responses={400: {"model": ErrorResponse}},
)
def create_student(payload: StudentCreateRequest) -> StudentResponse:
    # Business rule beyond schema validation: email uniqueness in our fake DB.
    if any(row["email"] == payload.email for row in FAKE_DB.values()):
        raise HTTPException(status_code=400, detail="email already exists")

    student_id = str(uuid4())
    row = payload.model_dump(exclude={"password", "confirm_password"})
    row["id"] = student_id
    row["created_at"] = datetime.utcnow()
    FAKE_DB[student_id] = row
    return make_student_response(row)


@app.get(
    "/students",
    response_model=StudentListResponse,
    tags=["Students"],
)
def list_students(course: Optional[str] = Query(default=None)) -> StudentListResponse:
    rows = list(FAKE_DB.values())
    if course:
        rows = [row for row in rows if row["course"] == course]
    items = [make_student_response(row) for row in rows]
    return StudentListResponse(total=len(items), items=items)


@app.get(
    "/students/{student_id}",
    response_model=StudentResponse,
    tags=["Students"],
    responses={404: {"model": ErrorResponse}},
)
def get_student(student_id: str = Path(..., min_length=10)) -> StudentResponse:
    row = FAKE_DB.get(student_id)
    if not row:
        raise HTTPException(status_code=404, detail="student not found")
    return make_student_response(row)


@app.patch(
    "/students/{student_id}",
    response_model=StudentResponse,
    tags=["Students"],
    responses={404: {"model": ErrorResponse}},
)
def update_student(student_id: str, payload: StudentUpdateRequest) -> StudentResponse:
    row = FAKE_DB.get(student_id)
    if not row:
        raise HTTPException(status_code=404, detail="student not found")

    updates = payload.model_dump(exclude_none=True)
    row.update(updates)
    FAKE_DB[student_id] = row
    return make_student_response(row)


@app.delete(
    "/students/{student_id}",
    tags=["Students"],
    responses={404: {"model": ErrorResponse}},
)
def delete_student(student_id: str) -> dict:
    if student_id not in FAKE_DB:
        raise HTTPException(status_code=404, detail="student not found")
    del FAKE_DB[student_id]
    return {"message": "student deleted successfully"}


# -----------------------------------------------------------------------------
# Sample payloads for quick testing
# -----------------------------------------------------------------------------
VALID_CREATE_PAYLOAD = {
    "full_name": "Deepak Bajaj",
    "email": "deepak@example.com",
    "age": 28,
    "course": "fastapi",
    "fees_paid": 25000,
    "discount": 3000,
    "skills": ["Python", "FastAPI", "Python"],
    "address": {"city": "daulatabad", "state": "maharashtra", "postal_code": "431002"},
    "phone": "98765-43210",
    "password": "Strong@123",
    "confirm_password": "Strong@123",
}

INVALID_CREATE_PAYLOAD = {
    "full_name": "Deepak",
    "email": "not-an-email",
    "age": 16,
    "course": "unknown",
    "fees_paid": 1000,
    "discount": 1500,
    "skills": ["   ", ""],
    "address": {"city": "x", "state": "1", "postal_code": "43#1002"},
    "phone": "12345",
    "password": "weak",
    "confirm_password": "different",
}


if __name__ == "__main__":
    import uvicorn

    uvicorn.run("pydantic_models_demo:app", host="127.0.0.1", port=8000, reload=True)
