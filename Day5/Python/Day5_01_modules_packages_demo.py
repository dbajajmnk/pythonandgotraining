"""
Day 5 Topic 1: Modules & Packages (Python) — with real-life scenario

Scenario: "Office Utilities Toolkit"
- You are building a small toolkit for an office team:
  - money.py: salary/tax utilities
  - textutils.py: clean names, normalize IDs
  - reports/ : report builders
Goal:
- Show how to structure code as modules and packages
- Show imports, __init__.py, and best practices
"""

from __future__ import annotations

# --- Module-level constants (safe: no heavy side effects) ---
APP_NAME = "Office Utilities Toolkit"


def main() -> None:
    """
    Hands-on Lab:
    1) Create a package folder structure (see comment below)
    2) Import and run functions from different modules
    3) Learn how package __init__.py exposes a clean API
    """
    print(f"Running: {APP_NAME}")

    # Import from modules (local package)
    # NOTE: This will work after you create the folder structure and run as a package.
    # See "How to run" instructions at bottom.
    from office_utils.money import net_salary
    from office_utils.textutils import normalize_employee_code
    from office_utils.reports.monthly import build_monthly_report

    emp_code = "  emp-  009  "
    clean_code = normalize_employee_code(emp_code)
    take_home = net_salary(gross=100000, tax_rate=0.10, deductions=2500)

    print("Clean employee code:", clean_code)
    print("Net salary:", take_home)

    report = build_monthly_report(employee_code=clean_code, net_salary=take_home)
    print("\n--- Monthly Report ---")
    print(report)


"""
✅ Folder structure to create (same folder where you run python):

project_root/
  day5_modules_packages_demo.py   (this file)
  office_utils/
    __init__.py
    money.py
    textutils.py
    reports/
      __init__.py
      monthly.py

✅ office_utils/__init__.py (example)
-----------------------------------
from .money import net_salary
from .textutils import normalize_employee_code
__all__ = ["net_salary", "normalize_employee_code"]

✅ How to run (recommended):
---------------------------
From project_root:
  python -m day5_modules_packages_demo  (only if file is a module without .py? not typical)
Better:
  python day5_modules_packages_demo.py

If you run into import errors:
- Ensure you are running from project_root
- Or install your package in editable mode (see virtual env topic)

Common mistakes:
- Name your file json.py / logging.py => shadows stdlib
- Circular imports (A imports B imports A)
- Putting heavy code at import time (DB calls, network calls)
"""

if __name__ == "__main__":
    main()
