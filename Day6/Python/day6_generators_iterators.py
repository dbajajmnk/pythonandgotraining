"""
Day 6 Topic: Generators & Iterators (Python) — In-Depth Theory + Hands-on Lab
============================================================================

Template followed:
What & Why → Mind Map → Engineering → Java↔Python → Mistakes → Lab

------------------------------------------------------------
INSTALL
------------------------------------------------------------
No external packages needed (Python 3.9+)

------------------------------------------------------------
RUN
------------------------------------------------------------
python day6_generators_iterators.py
"""

from __future__ import annotations

# ============================================================
# WHAT & WHY
# ============================================================
# WHAT:
# - Iterator: object you can loop over; has __iter__() and __next__()
# - Generator: easier iterator using 'yield' (lazy)
#
# WHY:
# - Saves memory (stream values one by one instead of storing all)
# - Makes pipelines for large data possible
#
# Java mapping:
# - Similar to lazy Streams (Files.lines, stream().filter())

# ============================================================
# ENGINEERING CONCEPT
# ============================================================
# - yield pauses function state and resumes later
# - generator is single-use; once exhausted, recreate it

def critical_lines(path: str):
    # Generator: yields only CRITICAL lines
    with open(path, "r", encoding="utf-8") as f:
        for line in f:
            if "CRITICAL" in line:
                yield line.rstrip("\n")

def make_sample_log(path: str) -> None:
    lines = [
        "INFO app started",
        "WARN cache low",
        "CRITICAL database down",
        "INFO retrying",
        "CRITICAL payment gateway timeout",
        "ERROR minor issue",
        "CRITICAL disk full",
    ]
    with open(path, "w", encoding="utf-8") as f:
        for l in lines:
            f.write(l + "\n")

# ============================================================
# HANDS-ON LAB: Large log scanning (streaming)
# ============================================================
# Scenario:
# - IT scans huge log files and wants only CRITICAL events.
#
# Requirements:
# 1) Create sample log
# 2) Use generator to stream CRITICAL lines
# 3) Count CRITICAL lines and print first N

def main() -> None:
    path = "sample.log"
    make_sample_log(path)

    gen = critical_lines(path)

    # Count lazily
    count = 0
    first_two = []
    for line in gen:
        count += 1
        if len(first_two) < 2:
            first_two.append(line)

    print("Total CRITICAL lines:", count)
    print("First 2 CRITICAL lines:", first_two)

    # COMMON MISTAKE demo: generator exhausted
    # If you try to loop again over 'gen', you'll get nothing.
    # Fix: recreate it:
    gen2 = critical_lines(path)
    print("Print all CRITICAL lines (fresh generator):")
    for line in gen2:
        print("-", line)

if __name__ == "__main__":
    main()
