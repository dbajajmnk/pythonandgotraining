"""
Day 5 Topic 2: Python Import System — in depth

Key ideas:
- sys.path determines where Python searches modules
- Absolute import vs relative import
- Import caching (sys.modules)
- import side effects
- Circular import avoidance
"""

from __future__ import annotations

import sys
import importlib
from types import ModuleType


def show_sys_path() -> None:
    print("---- sys.path (module search locations) ----")
    for i, p in enumerate(sys.path[:10], start=1):
        print(f"{i:02d}. {p}")
    if len(sys.path) > 10:
        print(f"... ({len(sys.path) - 10} more)")


def demonstrate_import_cache() -> None:
    print("\n---- import caching demo ----")
    import math  # noqa: F401

    # Once imported, module exists in sys.modules
    print("math in sys.modules:", "math" in sys.modules)

    # Importing again doesn't re-execute math; it returns cached module object
    import math as math2
    print("same module object:", sys.modules["math"] is math2)


def reload_module_demo() -> None:
    """
    Reload is useful for interactive sessions, not typical in production apps.
    """
    print("\n---- reload demo (safe example) ----")
    import random as rnd
    m: ModuleType = rnd
    print("random module id before:", id(m))
    m2 = importlib.reload(rnd)
    print("random module id after :", id(m2), "(same object but reloaded code)")


def common_import_problems() -> None:
    print("\n---- common import problems ----")
    print("1) Shadowing stdlib: do NOT create files like json.py, logging.py, sys.py")
    print("2) Circular imports: A imports B and B imports A")
    print("   Fix by moving imports inside functions, or reorganizing responsibilities.")
    print("3) Running a package module incorrectly:")
    print("   If you do: python office_utils/reports/monthly.py")
    print("   relative imports may break. Prefer: python -m office_utils.reports.monthly")


def main() -> None:
    show_sys_path()
    demonstrate_import_cache()
    reload_module_demo()
    common_import_problems()


if __name__ == "__main__":
    main()
