"""Duck Typing vs Liskov Substitution Principle (LSP) — Practical demos.

How to run:
    python duck_typing_vs_lsp.py

What you'll learn:
  1) Duck typing: accept any object that supports the needed behavior.
  2) LSP: if B is a subtype of A, then B must be usable anywhere A is expected
     without breaking the program's correctness.
  3) A classic LSP violation: Square as a subclass of Rectangle.
  4) A safe design: composition / separate interfaces (protocols).

This file is intentionally "teaching-first": lots of comments and printed output.
"""

from __future__ import annotations

from dataclasses import dataclass
from typing import Protocol, runtime_checkable


# ==========================================================
# 1) DUCK TYPING (CAPABILITY-BASED)
# ==========================================================

@runtime_checkable
class Notifier(Protocol):
    """A *behavior contract* (not inheritance) — anything with send(message) works."""

    def send(self, message: str) -> None: ...


class EmailNotifier:
    def __init__(self, email: str) -> None:
        self.email = email

    def send(self, message: str) -> None:
        print(f"[EMAIL to {self.email}] {message}")


class SmsNotifier:
    def __init__(self, phone: str) -> None:
        self.phone = phone

    def send(self, message: str) -> None:
        print(f"[SMS to {self.phone}] {message}")


class SlackNotifier:
    def __init__(self, channel: str) -> None:
        self.channel = channel

    def send(self, message: str) -> None:
        print(f"[SLACK #{self.channel}] {message}")


def alert_user(notifier: Notifier, message: str) -> None:
    """Duck typing in action.

    We do NOT care if notifier is EmailNotifier, SmsNotifier, SlackNotifier,
    or something else.

    We only care that it supports: notifier.send(str)
    """

    notifier.send(message)


def demo_duck_typing() -> None:
    print("\n=== Demo 1: Duck Typing ===")

    notifiers: list[Notifier] = [
        EmailNotifier("deepak@example.com"),
        SmsNotifier("+91-99999-11111"),
        SlackNotifier("support"),
    ]

    for n in notifiers:
        alert_user(n, "Your order is shipped.")

    # IMPORTANT: Duck typing fails at runtime if capability is missing.
    class BadNotifier:
        def push(self, message: str) -> None:
            print("pushed", message)

    print("\nDuck typing runtime failure example:")
    try:
        alert_user(BadNotifier(), "Hello")  # type: ignore[arg-type]
    except AttributeError as e:
        print("Expected error:", e)


# # ==========================================================
# # 2) LSP (INHERITANCE CONTRACT)
# # ==========================================================

# @dataclass
# class Rectangle:
#     """A rectangle is defined by independent width and height."""

#     width: int
#     height: int

#     def set_width(self, w: int) -> None:
#         self.width = w

#     def set_height(self, h: int) -> None:
#         self.height = h

#     def area(self) -> int:
#         return self.width * self.height


# class Square(Rectangle):
#     """Classic *bad* inheritance example.

#     A square is *not* a rectangle in behavior terms if Rectangle promises
#     independent set_width / set_height.
#     """

#     def set_width(self, w: int) -> None:
#         self.width = w
#         self.height = w

#     def set_height(self, h: int) -> None:
#         self.height = h
#         self.width = h


# def resize_and_measure(rect: Rectangle) -> int:
#     """A function that assumes Rectangle's contract.

#     If LSP holds, ANY subtype of Rectangle should work here.
#     """

#     rect.set_width(5)
#     rect.set_height(2)
#     return rect.area()  # expected 10


# def demo_lsp_violation() -> None:
#     print("\n=== Demo 2: LSP Violation (Square extends Rectangle) ===")

#     r = Rectangle(1, 1)
#     s = Square(1, 1)

#     r_area = resize_and_measure(r)
#     s_area = resize_and_measure(s)

#     print(f"Rectangle area expected=10, got={r_area}")
#     print(f"Square area expected=10, got={s_area}  <-- violates expectation")

#     print("\nWhy this is an LSP violation:")
#     print("- resize_and_measure expects independent width and height setters")
#     print("- Square changes meaning of set_width/set_height")
#     print("- So Square cannot safely substitute Rectangle")


# # ==========================================================
# # 3) LSP-SAFE DESIGN: COMPOSITION / SEPARATE SHAPES
# # ==========================================================

# @runtime_checkable
# class Shape(Protocol):
#     def area(self) -> int: ...


# @dataclass(frozen=True)
# class SafeRectangle:
#     width: int
#     height: int

#     def area(self) -> int:
#         return self.width * self.height


# @dataclass(frozen=True)
# class SafeSquare:
#     side: int

#     def area(self) -> int:
#         return self.side * self.side


# def total_area(shapes: list[Shape]) -> int:
#     return sum(s.area() for s in shapes)


# def demo_lsp_safe_design() -> None:
#     print("\n=== Demo 3: LSP-safe design (separate types + shared protocol) ===")

#     shapes: list[Shape] = [SafeRectangle(5, 2), SafeSquare(3)]
#     print("Total area:", total_area(shapes))

#     print("\nWhy this is safe:")
#     print("- Rectangle and Square share only what is truly common: area()")
#     print("- No misleading setters that break assumptions")


# # ==========================================================
# # 4) HANDS-ON LAB
# # ==========================================================

# def lab_prompt() -> None:
#     print("\n=== Hands-on Lab (with solution below) ===")
#     print("Scenario: Notification system for a small business")
#     print("1) Implement EmailNotifier and WhatsAppNotifier")
#     print("2) Write send_bulk(notifiers, message) using duck typing")
#     print("3) Add a failing notifier (missing send) and handle it safely")


# class WhatsAppNotifier:
#     def __init__(self, wa_number: str) -> None:
#         self.wa_number = wa_number

#     def send(self, message: str) -> None:
#         print(f"[WHATSAPP to {self.wa_number}] {message}")


# def send_bulk(notifiers: list[object], message: str) -> None:
#     """Duck typing with safety.

#     We accept any objects, but we guard at runtime:
#       - has attribute 'send'
#       - callable
#     """

#     for n in notifiers:
#         send_attr = getattr(n, "send", None)
#         if callable(send_attr):
#             send_attr(message)
#         else:
#             print(f"[SKIP] {type(n).__name__} has no usable send()")


# def lab_solution() -> None:
#     class BrokenNotifier:
#         def send_message(self, msg: str) -> None:
#             print("broken", msg)

#     notifiers: list[object] = [
#         EmailNotifier("billing@example.com"),
#         WhatsAppNotifier("+91-88888-22222"),
#         BrokenNotifier(),
#     ]
#     send_bulk(notifiers, "Invoice generated.")


def main() -> None:
    demo_duck_typing()
    # demo_lsp_violation()
    # demo_lsp_safe_design()
    # lab_prompt()
    # lab_solution()


if __name__ == "__main__":
    main()
