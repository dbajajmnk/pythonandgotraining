"""
office_utils package
--------------------
Public API for office utilities.
"""

from .money import net_salary
from .textutils import normalize_employee_code

__all__ = [
    "net_salary",
    "normalize_employee_code",
]
