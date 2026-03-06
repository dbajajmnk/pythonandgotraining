"""
office_utils package
--------------------
Public API for office utilities.
"""

from .money import net_salary
from .textutils import normalize_employee_code
from .myCode import normalizeMyStuff
from .reports.monthly import build_monthly_report
__all__ = [
    "net_salary",
    "normalize_employee_code",
    "normalizeMyStuff",
    "build_monthly_report"
]
