"""
textutils.py
-------------
Text cleaning utilities.
"""

def normalize_employee_code(code: str) -> str:
    """
    Normalize employee code:
    - Remove spaces
    - Uppercase
    - Standardize format

    Example:
    "  emp-  009  " → "EMP-009"
    """
    cleaned = code.strip().upper()
    cleaned = cleaned.replace(" ", "")
    return cleaned
