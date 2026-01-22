"""
Day 5 Topic 3: File I/O (text, CSV, JSON) — with office scenario

Scenario:
- Read employees.csv (id,name,gross)
- Read tax_rates.json (id -> tax_rate)
- Compute net salary and write report.txt

This file is self-contained: it creates sample CSV/JSON files if missing.
"""

from __future__ import annotations

import csv
import json
from pathlib import Path
from typing import Dict, List, Any


DATA_DIR = Path("data_day5")
EMP_CSV = DATA_DIR / "employees.csv"
TAX_JSON = DATA_DIR / "tax_rates.json"
REPORT_TXT = DATA_DIR / "salary_report.txt"


def ensure_sample_files() -> None:
    DATA_DIR.mkdir(exist_ok=True)

    if not EMP_CSV.exists():
        with EMP_CSV.open("w", newline="", encoding="utf-8") as f:
            writer = csv.writer(f)
            writer.writerow(["id", "name", "gross"])
            writer.writerow(["E001", "Asha", "100000"])
            writer.writerow(["E002", "Ravi", "85000"])
            writer.writerow(["E003", "Neha", "120000"])

    if not TAX_JSON.exists():
        tax_rates = {"E001": 0.10, "E002": 0.08, "E003": 0.12}
        TAX_JSON.write_text(json.dumps(tax_rates, indent=2), encoding="utf-8")


def read_employees() -> List[Dict[str, Any]]:
    with EMP_CSV.open("r", newline="", encoding="utf-8") as f:
        reader = csv.DictReader(f)
        employees: List[Dict[str, Any]] = []
        for row in reader:
            employees.append({"id": row["id"], "name": row["name"], "gross": int(row["gross"])})
        return employees


def read_tax_rates() -> Dict[str, float]:
    return json.loads(TAX_JSON.read_text(encoding="utf-8"))


def net_salary(gross: int, tax_rate: float, deductions: int = 0) -> int:
    tax = int(gross * tax_rate)
    return gross - tax - deductions


def write_report(lines: List[str]) -> None:
    REPORT_TXT.write_text("\n".join(lines) + "\n", encoding="utf-8")


def main() -> None:
    ensure_sample_files()

    employees = read_employees()
    tax_rates = read_tax_rates()

    lines = ["Salary Report", "============", ""]
    for emp in employees:
        rate = float(tax_rates.get(emp["id"], 0.10))  # default tax rate
        take_home = net_salary(emp["gross"], rate, deductions=2500)
        lines.append(f'{emp["id"]} | {emp["name"]} | gross={emp["gross"]} | tax={rate:.0%} | net={take_home}')

    write_report(lines)
    print(f"Report created: {REPORT_TXT.resolve()}")
    print(REPORT_TXT.read_text(encoding="utf-8"))

    print("\nCommon mistakes:")
    print("- Not using encoding='utf-8' (especially on Windows)")
    print("- Forgetting newline='' with csv module -> blank lines on Windows")
    print("- Using open() without 'with' -> file handle leaks")
    print("- Assuming JSON keys are ints; JSON keys are strings")


if __name__ == "__main__":
    main()
