"""
textutils.py
-------------
Text cleaning utilities.
"""

def normalizeMyStuff(code: str) -> str:
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
