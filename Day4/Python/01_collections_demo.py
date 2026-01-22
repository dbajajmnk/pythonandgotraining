"""01_collections_demo.py

Python collections: list, tuple, set, dict — with a hands-on lab scenario.

Run:
  python 01_collections_demo.py
"""

from __future__ import annotations


def basics_demo() -> None:
    print("\n=== BASICS: list / tuple / set / dict ===")

    # LIST: ordered, mutable, allows duplicates
    shopping_list = ["milk", "bread", "eggs", "bread"]
    shopping_list.append("coffee")
    print("List (ordered, duplicates ok):", shopping_list)

    # TUPLE: ordered, immutable (great for fixed "bundles" like coordinates)
    seat = ("Row A", 12)  # (row, seat_number)
    print("Tuple (fixed bundle):", seat)

    # SET: unique items only (fast membership checks)
    unique_items = set(shopping_list)
    print("Set (unique):", unique_items)
    print("Is 'milk' in set?", "milk" in unique_items)

    # DICT: key -> value (fast lookup by key)
    prices = {"milk": 2.50, "bread": 1.25, "eggs": 3.10}
    prices["coffee"] = 8.99
    print("Dict (key -> value):", prices)
    print("Price of eggs:", prices["eggs"])


# -------------------------
# Hands-on Lab
# -------------------------

def assign_seats(emails_in_order: list[str], capacity: int) -> tuple[dict[str, tuple[str, int]], list[str]]:
    """Assign seats to unique registrants.

    Uses:
      - list: input order
      - set: de-duplicate
      - dict: email -> seat tuple
      - tuple: seat is a fixed bundle (row, number)

    Returns:
      (seat_map, waitlist)
    """
    seen: set[str] = set()
    seat_map: dict[str, tuple[str, int]] = {}
    waitlist: list[str] = []

    # Seat layout: 10 seats per row => A1..A10, B1..B10, ...
    def seat_for_index(i: int) -> tuple[str, int]:
        row_letter = chr(ord("A") + (i // 10))
        seat_number = (i % 10) + 1
        return (f"Row {row_letter}", seat_number)

    for email in emails_in_order:
        email_norm = email.strip().lower()
        if not email_norm or "@" not in email_norm:
            # skip obviously bad entries
            continue
        if email_norm in seen:
            continue

        seen.add(email_norm)
        if len(seat_map) < capacity:
            seat_map[email_norm] = seat_for_index(len(seat_map))
        else:
            waitlist.append(email_norm)

    return seat_map, waitlist


def lab_conference_registration() -> None:
    print("\n=== LAB: Conference registration ===")

    signups = [
        "Alice@Example.com",
        "bob@example.com",
        "ALICE@example.com",  # duplicate
        "carol@example.com",
        "dave@example.com",
        "",  # bad
        "not-an-email",  # bad
        "erin@example.com",
    ]

    capacity = 4
    seat_map, waitlist = assign_seats(signups, capacity=capacity)

    print(f"Capacity: {capacity}")
    print("\nSeat assignments (dict: email -> tuple(row, seat))")
    for email, seat in seat_map.items():
        print(f"  {email:18s} -> {seat}")

    print("\nWaitlist (list):", waitlist)
    print("\nUnique registrants (set size):", len(set(seat_map.keys()) | set(waitlist)))


def main() -> None:
    basics_demo()
    lab_conference_registration()


if __name__ == "__main__":
    main()
