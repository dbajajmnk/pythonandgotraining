"""03_exceptions_demo.py

Exception handling in Python:
  - try / except / else / finally
  - raising and chaining exceptions

Hands-on lab: robust import of "id,amount" lines.

Run:
  python 03_exceptions_demo.py
"""

from __future__ import annotations


def parse_amount(text: str) -> int:
    """Parse an integer amount with clear errors.

    - ValueError is raised for invalid input.
    - We keep the original cause using exception chaining (raise ... from e).
    """
    try:
        amt = int(text)
    except ValueError as e:
        raise ValueError(f"Bad amount: {text!r}") from e
    else:
        return amt
    finally:
        # finally always runs (useful for metrics/counters/log flushing)
        pass


def basics_demo() -> None:
    print("\n=== BASICS: try/except/else/finally ===")
    for s in ["42", "oops"]:
        try:
            print("Parsing", s, "->", parse_amount(s))
        except ValueError as e:
            print("  ERROR:", e)


# -------------------------
# Hands-on Lab
# -------------------------


def import_lines(lines: list[str]) -> dict[int, int]:
    """Import lines like 'id,amount'.

    Requirements:
      - Skip malformed lines
      - Count bad lines
      - Sum amounts per id
      - Always print a final report (finally)

    Returns:
      dict[id] = total_amount
    """
    totals: dict[int, int] = {}
    bad = 0
    processed = 0

    try:
        for raw in lines:
            processed += 1
            raw = raw.strip()
            if not raw:
                bad += 1
                continue

            try:
                left, right = raw.split(",")
                rec_id = int(left)
                amt = parse_amount(right)
            except (ValueError, IndexError) as e:
                bad += 1
                # In real apps: log e + the raw line
                continue

            totals[rec_id] = totals.get(rec_id, 0) + amt

        return totals
    finally:
        print("\n[FINAL REPORT]")
        print("  processed:", processed)
        print("  bad lines:", bad)
        print("  valid ids :", len(totals))


def lab_import() -> None:
    print("\n=== LAB: Robust import of 'id,amount' ===")
    lines = [
        "101,250",
        "102,90",
        "bad-line",
        "103,not-a-number",
        "101,10",
        "",
    ]
    totals = import_lines(lines)
    print("\nTotals:")
    for k in sorted(totals):
        print(f"  {k}: {totals[k]}")


def main() -> None:
    basics_demo()
    lab_import()


if __name__ == "__main__":
    main()
