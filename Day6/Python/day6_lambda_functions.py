"""
Day 6 Topic: Lambda Functions (Python) — In-Depth Theory + Hands-on Lab
=====================================================================

Template followed:
What & Why → Mind Map → Engineering → Java↔Python → Mistakes → Lab

------------------------------------------------------------
INSTALL
------------------------------------------------------------
No external packages needed (Python 3.9+)

------------------------------------------------------------
RUN
------------------------------------------------------------
python day6_lambda_functions.py
"""

# ============================================================
# WHAT & WHY
# ============================================================
# WHAT:
# - lambda = small anonymous function: lambda <params>: <expression>
# - It returns a callable (function object).
#
# WHY:
# - Avoid writing tiny helper functions that are used only once
# - Useful in pipelines: transform/filter/sort in a compact way
#
# WHEN:
# - Sorting keys, simple predicates, simple transformations
#
# WHEN NOT:
# - If logic needs comments, multiple steps, error handling → use def

# ============================================================
# MIND MAP (Normal-life use cases)
# ============================================================
# Office HR: "give me active employees sorted by salary"
# E-commerce: "sort products by discount%"
# Data cleaning: "trim + lower while filtering invalid rows"

# ============================================================
# ENGINEERING CONCEPT
# ============================================================
# - Functions are first-class (can be passed as values)
# - lambda is expression-only (no statements like for/try/assign)
# - Readability rule: prefer list comprehensions when clearer

# ============================================================
# JAVA ↔ PYTHON (Concept + Code)
# ============================================================
# Java (Streams):
#   nums.stream().map(n -> n*2).collect(...)
# Python:
#   list(map(lambda n: n*2, nums))
# Often clearer in Python:
#   [n*2 for n in nums]

# ============================================================
# CODE EXAMPLES
# ============================================================

def demo_lambda_basics() -> None:
    square = lambda x: x * x
    print("square(5) =", square(5))

    nums = [1, 2, 3, 4]
    doubled = list(map(lambda n: n * 2, nums))
    print("doubled =", doubled)

    # Sorting real objects/dicts using lambda as key
    employees = [
        {"name": "Asha", "netSalary": 95000, "active": True},
        {"name": "Ravi", "netSalary": 120000, "active": True},
        {"name": "Nina", "netSalary": 80000, "active": False},
        {"name": "Irfan", "netSalary": 110000, "active": True},
    ]
    by_salary_desc = sorted(employees, key=lambda e: e["netSalary"], reverse=True)
    print("Top by salary:", [e["name"] for e in by_salary_desc])


# ============================================================
# COMMON MISTAKES
# ============================================================
# ❌ Complex lambda: lambda x: (x+1 if ... else ... if ... else ...)
# ✅ Use def for clarity
#
# ❌ Overusing map/filter when list comprehension is clearer
# ✅ Prefer: [x*2 for x in nums if x>2]

# ============================================================
# HANDS-ON LAB (Real-life Scenario)
# ============================================================
# Scenario:
# - HR wants a quick report: Active employees sorted by netSalary (desc),
#   then show Top 3.
#
# Requirements:
# 1) Filter active employees
# 2) Sort by netSalary descending
# 3) Print top 3 with name + salary
# 4) Show a list-comprehension alternative

def lab_hr_filter_sort(employees: list[dict]) -> list[dict]:
    # STEP 1: filter active (predicate)
    active = list(filter(lambda e: e.get("active") is True, employees))
    print("Active",active)

    # STEP 2: sort by netSalary desc (key function)
    sorted_active = sorted(active, key=lambda e: e.get("netSalary", 0), reverse=True)
    print("Sorted Active",sorted_active)

    # STEP 3: return top 3
    return sorted_active[:3]


def lab_hr_filter_sort_comprehension(employees: list[dict]) -> list[dict]:
    # Alternative: list comprehension for filtering + sorted for ordering
    active = [e for e in employees if e.get("active") is True]
    return sorted(active, key=lambda e: e.get("netSalary", 0), reverse=True)[:3]


def main() -> None:
    demo_lambda_basics()

    employees = [
        {"name": "Asha", "netSalary": 95000, "active": True},
        {"name": "Ravi", "netSalary": 120000, "active": True},
        {"name": "Nina", "netSalary": 80000, "active": False},
        {"name": "Irfan", "netSalary": 110000, "active": True},
        {"name": "Meera", "netSalary": 105000, "active": True},
    ]

    top3 = lab_hr_filter_sort(employees)
    print("\nLAB RESULT (lambda filter/sort):")
    for e in top3:
        print(f"- {e['name']}: {e['netSalary']}")

    # top3b = lab_hr_filter_sort_comprehension(employees)
    # print("\nLAB RESULT (list comprehension + sorted):")
    # for e in top3b:
    #     print(f"- {e['name']}: {e['netSalary']}")


if __name__ == "__main__":
    main()
