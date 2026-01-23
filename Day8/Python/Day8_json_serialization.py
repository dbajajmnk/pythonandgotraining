"""
Day 8 — JSON Serialization (Backend Perspective) — In-Depth + Lab
================================================================

------------------------------------------------------------
INSTALL
------------------------------------------------------------
# Standard library json is enough for basics.
# Optional (recommended for APIs):
pip install pydantic

------------------------------------------------------------
RUN
------------------------------------------------------------
python day8_json_serialization.py
"""

from __future__ import annotations

import json
from dataclasses import dataclass, asdict
from datetime import datetime
from decimal import Decimal
from uuid import UUID, uuid4

# ============================================================
# WHAT:
# - Serialization: object -> JSON string (for HTTP, logs, storage)
# - Deserialization: JSON string -> Python objects
#
# WHY:
# - APIs talk JSON; contracts must be stable, safe, validated.
# ============================================================


# ============================================================
# ENGINEERING: Custom JSON encoder
# ============================================================
class SafeJSONEncoder(json.JSONEncoder):
    """
    WHY: json.dumps can't serialize datetime/UUID/Decimal by default.
    Approach: convert to stable string representations.
    """
    def default(self, o):  # noqa: D401
        if isinstance(o, datetime):
            return o.isoformat()
        if isinstance(o, UUID):
            return str(o)
        if isinstance(o, Decimal):
            # Use string to avoid floating-point rounding surprises
            return str(o)
        return super().default(o)


@dataclass
class Invoice:
    id: UUID
    customer: str
    amount: Decimal
    created_at: datetime


def lab_serialize_deserialize():
    inv = Invoice(
        id=uuid4(),
        customer="ACME Corp",
        amount=Decimal("1234.50"),
        created_at=datetime.utcnow(),
    )

    # Convert dataclass -> dict
    inv_dict = asdict(inv)

    # Serialize
    json_str = json.dumps(inv_dict, cls=SafeJSONEncoder, indent=2)
    print("Serialized JSON:\n", json_str)

    # Write to file (common DevOps/logging use case)
    with open("invoice.json", "w", encoding="utf-8") as f:
        f.write(json_str)

    # Read back and deserialize
    with open("invoice.json", "r", encoding="utf-8") as f:
        loaded = json.loads(f.read())

    print("\nDeserialized dict:\n", loaded)

    # NOTE:
    # Types are now strings in JSON.
    # In real APIs, you'd validate and convert using a schema library (Pydantic).
    try:
        from pydantic import BaseModel

        class InvoiceSchema(BaseModel):
            id: UUID
            customer: str
            amount: Decimal
            created_at: datetime

        obj = InvoiceSchema.model_validate(loaded)
        print("\nValidated + typed object (Pydantic):\n", obj)
    except Exception as e:
        print("\nPydantic not installed or validation failed:", e)


if __name__ == "__main__":
    lab_serialize_deserialize()
