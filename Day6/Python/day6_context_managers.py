"""
Day 6 Topic: Context Managers (with) (Python) — In-Depth Theory + Hands-on Lab
=============================================================================

Template followed:
What & Why → Mind Map → Engineering → Java↔Python → Mistakes → Lab

------------------------------------------------------------
INSTALL
------------------------------------------------------------
No external packages needed (Python 3.9+)

------------------------------------------------------------
RUN
------------------------------------------------------------
python day6_context_managers.py
"""

from __future__ import annotations
from contextlib import contextmanager
import threading

# ============================================================
# WHAT & WHY
# ============================================================
# WHAT:
# - Context manager controls setup/cleanup using __enter__/__exit__
# - with-statement guarantees cleanup even if exceptions happen
#
# WHY:
# - Avoid resource leaks (files, DB connections, locks)
# - Replace try/finally boilerplate
#
# Java mapping:
# - try-with-resources

# ============================================================
# HANDS-ON LAB: Safe file import + lock handling
# ============================================================
# Scenario:
# - Employee import tool uses a lock (shared resource) and a file.
# - We must always release lock and close file even if an error happens.

LOCK = threading.Lock()

@contextmanager
def file_transaction(path: str, mode: str):
    print(f"[TX] OPEN {path}")
    f = open(path, mode, encoding="utf-8")
    try:
        yield f
        print(f"[TX] COMMIT {path}")
    except Exception as e:
        print(f"[TX] ROLLBACK {path} because {e!r}")
        raise
    finally:
        f.close()
        print(f"[TX] CLOSE {path}")

def main() -> None:
    # Create a demo file
    with open("employees.txt", "w", encoding="utf-8") as f:
        f.write("Asha\nRavi\nNina\n")

    try:
        with LOCK:  # lock automatically released
            with file_transaction("employees.txt", "r") as f:
                for line in f:
                    name = line.strip()
                    print("Importing:", name)
                    if name == "Ravi":
                        # Simulate an exception in the middle
                        raise RuntimeError("Simulated failure during import")
    except RuntimeError:
        print("Handled error; notice cleanup still happened.")

if __name__ == "__main__":
    main()
