# Pydantic Request/Response Models, Validation, Custom Validators

## 1) What this file set gives you
This package contains two files:

- `pydantic_models_demo.py` → one single Python source file
- `USERMANUAL_Pydantic_Request_Response_Validation.md` → step-by-step learner manual

The Python file demonstrates:

- request models
- response models
- built-in validation
- custom field validators
- custom model validators
- nested models
- partial update models
- response filtering with FastAPI

---

## 2) Why this topic matters
In real APIs, users and frontend apps send JSON data.
That data may be:

- missing required values
- in the wrong type
- invalid by business rules
- unsafe to store directly

Pydantic solves this by validating data before your business logic runs.
FastAPI uses Pydantic models for both request input and response output.

---

## 3) Real-life mind mapping
Think about an **online training institute admission form**.

A student submits:

- full name
- email
- age
- selected course
- fees paid
- discount
- skills
- address
- password

### Real-life checks
Before accepting the form, the institute must verify:

- name is complete
- email format is valid
- age is allowed
- discount is not greater than fees paid
- password is strong
- confirm password matches password
- phone number is valid

That is exactly what Pydantic validation does for your backend.

---

## 4) What are request models?
A **request model** defines the shape of incoming data.

Example from the file:

- `StudentCreateRequest`
- `StudentUpdateRequest`

### Why request models are useful
They help you:

- define what the client is allowed to send
- reject bad input automatically
- reduce manual validation code
- generate API docs automatically

---

## 5) What are response models?
A **response model** defines the shape of outgoing API data.

Example from the file:

- `StudentResponse`
- `StudentListResponse`
- `ErrorResponse`

### Why response models are useful
They help you:

- return clean structured output
- hide internal fields like passwords
- keep API responses consistent
- document your API clearly

---

## 6) Built-in validation used in the source file
The code uses built-in validation features such as:

- `EmailStr` for email format
- `Field(min_length=..., max_length=...)`
- `Field(ge=..., le=..., gt=...)`
- `Literal[...]` for allowed course values
- nested child model validation using `Address`
- `extra="forbid"` to reject unexpected fields

---

## 7) Custom validators used in the source file
The code shows two kinds of custom validators.

### A. Field validators
Used when one field needs its own rule.

Examples:

- normalize city and state
- postal code must be alphanumeric
- full name must contain at least 2 words
- phone must have exactly 10 digits
- password must include uppercase, lowercase, digit, and special character
- skills list must be cleaned and deduplicated

### B. Model validators
Used when multiple fields must be checked together.

Examples:

- `discount < fees_paid`
- `password == confirm_password`

---

## 8) When should you use custom validators?
Use custom validators when built-in validation is not enough.

Common cases:

- password policy
- age rules based on business logic
- cross-field comparison
- discount calculations
- date relationships
- normalized formatting rules

---

## 9) How the source file is structured
The file is intentionally written in one source file so students can understand the complete flow in one place.

### Section breakdown
1. imports and FastAPI app creation
2. fake in-memory database
3. nested shared model `Address`
4. shared base model `StudentBase`
5. request models
6. response models
7. helper function to shape final output
8. CRUD endpoints
9. ready-made valid and invalid test payloads

---

## 10) API endpoints in this example

### `GET /`
Health message and docs URL.

### `POST /students`
Creates a student.
Uses request model:
- `StudentCreateRequest`

Returns response model:
- `StudentResponse`

### `GET /students`
Lists all students.
Returns:
- `StudentListResponse`

### `GET /students/{student_id}`
Fetches one student.
Returns:
- `StudentResponse`

### `PATCH /students/{student_id}`
Updates selected fields only.
Uses:
- `StudentUpdateRequest`

### `DELETE /students/{student_id}`
Deletes a student.

---

## 11) How to run the project

### Step 1: create a virtual environment
```bash
python -m venv .venv
```

### Step 2: activate it
#### Windows
```bash
.venv\Scripts\activate
```

#### macOS/Linux
```bash
source .venv/bin/activate
```

### Step 3: install dependencies
```bash
pip install fastapi uvicorn pydantic[email]
```

### Step 4: run the app
```bash
uvicorn pydantic_models_demo:app --reload
```

### Step 5: open Swagger UI
Open this in your browser:
```text
http://127.0.0.1:8000/docs
```

---

## 12) Valid payload for testing
Use this payload in `POST /students`:

```json
{
  "full_name": "Deepak Bajaj",
  "email": "deepak@example.com",
  "age": 28,
  "course": "fastapi",
  "fees_paid": 25000,
  "discount": 3000,
  "skills": ["Python", "FastAPI", "Python"],
  "address": {
    "city": "daulatabad",
    "state": "maharashtra",
    "postal_code": "431002"
  },
  "phone": "98765-43210",
  "password": "Strong@123",
  "confirm_password": "Strong@123"
}
```

### What you should observe
- email is validated
- duplicate skills become unique
- city/state are normalized
- password fields are not returned in the response
- net fees are calculated in the response helper

---

## 13) Invalid payload for testing
Use this to observe validation errors:

```json
{
  "full_name": "Deepak",
  "email": "not-an-email",
  "age": 16,
  "course": "unknown",
  "fees_paid": 1000,
  "discount": 1500,
  "skills": ["   ", ""],
  "address": {
    "city": "x",
    "state": "1",
    "postal_code": "43#1002"
  },
  "phone": "12345",
  "password": "weak",
  "confirm_password": "different"
}
```

### What you should observe
FastAPI + Pydantic return structured validation errors automatically.

---

## 14) Engineering concept behind this example

### Input side
Incoming JSON is parsed into a request model.
If validation fails, the request is rejected before business logic runs.

### Processing side
Business logic checks additional runtime conditions.
Example in code:
- duplicate email check against in-memory store

### Output side
Response models shape the final output.
This prevents accidental leakage of internal fields.

---

## 15) Why this design is good for real projects
This pattern improves:

- maintainability
- readability
- API contract clarity
- testing quality
- frontend/backend coordination
- production safety

---

## 16) Common mistakes students make

### Mistake 1: using one model for everything
Better approach:
- separate create model
- separate update model
- separate response model

### Mistake 2: returning password fields in API response
Never expose sensitive request-only data.

### Mistake 3: only checking type, not business rule
Example:
- `discount` may be numeric but still invalid if it exceeds `fees_paid`

### Mistake 4: not forbidding extra fields
Unexpected fields can create messy or unsafe payloads.

### Mistake 5: mixing validation and database logic badly
Schema validation and business validation should stay organized.

---

## 17) Hands-on labs

### Lab 1 — Student registration
Add a new field:
- `experience_years`

Rules:
- must be `>= 0`
- must be `<= 25`

### Lab 2 — Course-specific rule
Add a custom validation rule:
- if course is `ai`, student must have `python` in skills

### Lab 3 — Date validation
Add fields:
- `start_date`
- `end_date`

Rule:
- `end_date` must be after `start_date`

### Lab 4 — Response filtering
Add an admin-only response model that shows more fields.
Keep the public response model smaller.

### Lab 5 — Reusable nested models
Create a separate `GuardianContact` nested model and add it to student registration.

---

## 18) Knowledge check
Answer these after studying the code.

1. What is the difference between a request model and a response model?
2. Why do we use `extra="forbid"`?
3. Why is `EmailStr` better than plain `str` for emails?
4. When should you use `field_validator`?
5. When should you use `model_validator`?
6. Why should password not appear in the response model?
7. Why is `StudentUpdateRequest` separate from `StudentCreateRequest`?
8. What is the benefit of nested models like `Address`?
9. What happens before FastAPI reaches your endpoint logic when data is invalid?
10. Why is `discount < fees_paid` a model-level rule?
11. Why are skills normalized and deduplicated?
12. What is the benefit of response models in team development?

---

## 19) Mini interview questions

1. Explain Pydantic in plain English.
2. What is the role of request models in FastAPI?
3. What is the role of response models in FastAPI?
4. What is the difference between built-in validation and custom validation?
5. Why would you create separate models for create, update, and response?
6. Give a real example of a field validator.
7. Give a real example of a model validator.
8. How does validation improve API security and reliability?

---

## 20) Practice path for students
Recommended order:

1. Run the file
2. Open `/docs`
3. Test valid payload
4. Test invalid payload
5. Read each validation error carefully
6. Add one new field validation rule
7. Add one cross-field validator
8. Add one new endpoint and response model

---

## 21) Final takeaway
Remember this simple flow:

**Client JSON → Request Model → Validation → Business Logic → Response Model → Safe API Output**

That is the core idea of Pydantic in FastAPI.
