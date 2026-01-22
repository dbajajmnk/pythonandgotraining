"""
money.py
---------
Salary and tax utilities for the office toolkit.
"""

def net_salary(gross: float, tax_rate: float, deductions: float) -> float:
    """
    Calculate net salary after tax and deductions.

    Example:
    gross=100000
    tax_rate=0.10
    deductions=2500
    """
    tax_amount = gross * tax_rate
    net = gross - tax_amount - deductions
    return round(net, 2)
