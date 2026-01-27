"""
Day 6 Topic: Decorators (Python) — In-Depth Theory + Hands-on Lab
================================================================

Template followed:
What & Why → Mind Map → Engineering → Java↔Python → Mistakes → Lab

------------------------------------------------------------
INSTALL
------------------------------------------------------------
No external packages needed (Python 3.9+)

------------------------------------------------------------
RUN
------------------------------------------------------------
python day6_decorators.py
"""

from __future__ import annotations
import time
import functools

# ============================================================
# WHAT & WHY
# ============================================================
# WHAT:
# - Decorator = function that takes a function and returns a new function.
# - It "wraps" behavior around original function.
#
# WHY:
# - Add cross-cutting concerns (logging, timing, auth) without duplicating code.
#
# WHEN:
# - Same pre/post behavior required on many functions.
#
# Java mapping:
# - Similar to annotations + AOP (e.g., Spring @Transactional, @Timed)

# ============================================================
# ENGINEERING CONCEPT
# ============================================================
# - Higher-order function: accepts function, returns function
# - Closure: wrapper remembers 'fn'
# - Use functools.wraps to preserve name/docstring

def logged(fn):
    @functools.wraps(fn)
    def wrapper(*args, **kwargs):
        print(f"[LOG] START {fn.__name__} args={args} kwargs={kwargs}")
        try:
            return fn(*args, **kwargs)
        finally:
            print(f"[LOG] END   {fn.__name__}")
    return wrapper

def timed(fn):
    @functools.wraps(fn)
    def wrapper(*args, **kwargs):
        start = time.time()
        try:
            return fn(*args, **kwargs)
        finally:
            ms = (time.time() - start) * 1000
            print(f"[TIME] {fn.__name__} took {ms:.2f} ms")
    return wrapper

# ============================================================
# COMMON MISTAKES
# ============================================================
# ❌ wrapper without *args/**kwargs breaks for different signatures
# ❌ not using wraps: fn.__name__ becomes 'wrapper'
# ❌ forgetting to return the original function result

# ============================================================
# HANDS-ON LAB: Office Report Service
# ============================================================
# Scenario:
# - Office "Report Service" wants consistent logging + timing across methods.
#
# Requirements:
# 1) Build @logged and @timed decorators
# 2) Apply to 3 business functions
# 3) Ensure logs print even if exception occurs

@logged
@timed
def compute_tax(gross_salary: float) -> float:
    # Dummy computation
    time.sleep(0.05)  # simulate work
    return gross_salary * 0.20

@logged
@timed
def build_monthly_report(employee_name: str, gross_salary: float) -> str:
    time.sleep(0.03)  # simulate work
    tax = compute_tax(gross_salary)
    net = gross_salary - tax
    return f"Monthly Report: {employee_name} | Gross={gross_salary} | Tax={tax} | Net={net}"

@logged
@timed
def export_pdf(report_text: str) -> None:
    time.sleep(0.02)  # simulate work
    # In real life you'd write to a PDF file.
    print("[EXPORT] (pretend) PDF exported")

def main() -> None:
    report = build_monthly_report(employee_name="Asha",gross_salary=100000)
    print(report)
    export_pdf(report)

if __name__ == "__main__":
    main()
