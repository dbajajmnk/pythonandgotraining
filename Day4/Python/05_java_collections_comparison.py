"""05_java_collections_comparison.py

Python collections vs Java Collections Framework (quick, practical mapping).

This script focuses on *how you think* when switching languages:
  - list  ↔ List (ArrayList)
  - tuple ↔ immutable bundle (unmodifiable List / small class)
  - set   ↔ Set (HashSet)
  - dict  ↔ Map (HashMap)

Hands-on lab: group student roster by class, dedup, and sort.

Run:
  python 05_java_collections_comparison.py
"""

from __future__ import annotations


def quick_mapping() -> None:
    print("\n=== Quick mapping ===")
    print("Python list  -> Java List (e.g., ArrayList)")
    print("Python tuple -> Java immutable bundle (e.g., List.of / small class)")
    print("Python set   -> Java Set (e.g., HashSet)")
    print("Python dict  -> Java Map (e.g., HashMap)")


def lab_student_roster() -> None:
    print("\n=== LAB: Student roster management ===")

    # Input: (student, className) pairs; duplicates possible
    enrollments = [
        ("Alice", "Math"),
        ("Bob", "Math"),
        ("Alice", "Math"),  # duplicate
        ("Carol", "Science"),
        ("Bob", "Science"),
        ("Dave", "Math"),
        ("Alice", "Science"),
    ]

    # Core structure: Map<className, Set<student>>
    roster: dict[str, set[str]] = {}
    for student, cls in enrollments:
        roster.setdefault(cls, set()).add(student)

    # Deterministic output: sort classes and students
    all_students: set[str] = set()
    for cls in sorted(roster):
        students_sorted = sorted(roster[cls])
        all_students.update(roster[cls])
        print(f"{cls}: {students_sorted}")

    print("Total unique students:", len(all_students))


def main() -> None:
    quick_mapping()
    lab_student_roster()


if __name__ == "__main__":
    main()
