"""
monthly.py
-----------
Monthly report builders.
"""

def build_monthly_report(employee_code: str, net_salary: float) -> str:
    """
    Build a simple monthly salary report.
    """
    return f"""
Employee Code : {employee_code}
Net Salary    : {net_salary}
Status        : PROCESSED
"""
