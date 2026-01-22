"""02_comprehensions_demo.py

Collection comprehensions in Python:
  - list / set / dict comprehensions
  - generator expressions (lazy)

Hands-on lab: clean and summarize sales transactions.

Run:
  python 02_comprehensions_demo.py
"""

from __future__ import annotations

from collections import defaultdict


def basics_demo() -> None:
    print("\n=== BASICS: list / set / dict comprehensions ===")

    nums = [1, 2, 3, 4, 5, 6]

    squares = [n * n for n in nums]
    evens = [n for n in nums if n % 2 == 0]
    unique_mod_3 = {n % 3 for n in nums}
    as_text = {n: f"#{n}" for n in nums}

    print("Squares:", squares)
    print("Evens:", evens)
    print("Unique (n % 3):", unique_mod_3)
    print("Dict mapping:", as_text)

    # Generator expression: lazy (doesn't build a list)
    total = sum(n * n for n in nums)
    print("Sum of squares (generator):", total)


# -------------------------
# Hands-on Lab
# -------------------------

Transaction = dict[str, object]


def normalize(transactions: list[Transaction]) -> list[Transaction]:
    """Clean input rows.

    Rules:
      - keep rows with valid customer + positive amount
      - strip whitespace, lower-case customer ids
      - coerce amount to int if possible
    """
    cleaned: list[Transaction] = []
    for t in transactions:
        cust_raw = str(t.get("customer", "")).strip().lower()
        if not cust_raw:
            continue

        try:
            amt = int(t.get("amount", 0))
        except (TypeError, ValueError):
            continue

        if amt <= 0:
            continue

        cleaned.append({"customer": cust_raw, "amount": amt, "item": t.get("item", "")})
    return cleaned


def lab_sales_summary() -> None:
    print("\n=== LAB: Sales cleanup + summary ===")

    raw: list[Transaction] = [
        {"customer": " Alice ", "amount": "120", "item": "Book"},
        {"customer": "bob", "amount": 50, "item": "Pen"},
        {"customer": "", "amount": 10, "item": "Sticker"},  # bad customer
        {"customer": "alice", "amount": -5, "item": "Refund"},  # bad amount
        {"customer": "carol", "amount": "oops", "item": "Bag"},  # bad amount
        {"customer": "bob ", "amount": 70, "item": "Notebook"},
    ]

    cleaned = normalize(raw)
    print("Cleaned rows:")
    for row in cleaned:
        print("  ", row)

    # 1) Build a list of amounts (list comprehension)
    amounts = [row["amount"] for row in cleaned]
    print("\nAmounts:", amounts)

    # 2) Unique customers (set comprehension)
    customers = {row["customer"] for row in cleaned}
    print("Unique customers:", customers)

    # 3) Totals per customer (dict comprehension + defaultdict)
    totals: defaultdict[str, int] = defaultdict(int)
    for row in cleaned:
        totals[str(row["customer"])] += int(row["amount"])

    totals_dict = {cust: totals[cust] for cust in sorted(totals)}
    print("Totals per customer:", totals_dict)

    # 4) High-value customers (comprehension with if)
    high_value = [cust for cust, total in totals_dict.items() if total >= 100]
    print("High-value customers (>=100):", high_value)


def main() -> None:
    basics_demo()
    lab_sales_summary()


if __name__ == "__main__":
    main()
