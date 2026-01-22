"""04_custom_exceptions_demo.py

Custom exceptions in Python:
  - define your own exception classes
  - use them to represent domain failures

Hands-on lab: e-commerce checkout with domain-specific failures.

Run:
  python 04_custom_exceptions_demo.py
"""

from __future__ import annotations


class CheckoutError(Exception):
    """Base class for all checkout-related errors."""


class InvalidCouponError(CheckoutError):
    pass


class OutOfStockError(CheckoutError):
    def __init__(self, sku: str, requested: int, available: int):
        super().__init__(f"Out of stock: {sku} (requested {requested}, available {available})")
        self.sku = sku
        self.requested = requested
        self.available = available


class PaymentDeclinedError(CheckoutError):
    pass


def apply_coupon(subtotal: int, coupon: str | None) -> int:
    """Apply a simple coupon.

    - "SAVE10" gives 10% off
    - unknown coupon raises InvalidCouponError
    """
    if not coupon:
        return subtotal
    coupon = coupon.strip().upper()
    if coupon == "SAVE10":
        return int(subtotal * 0.90)
    raise InvalidCouponError(f"Unknown coupon: {coupon}")


def checkout(
    cart: dict[str, int],
    inventory: dict[str, int],
    prices: dict[str, int],
    coupon: str | None,
    wallet_cents: int,
) -> int:
    """Charge the customer and reduce inventory.

    Returns the charged amount (cents).
    """
    # 1) Validate inventory
    for sku, qty in cart.items():
        available = inventory.get(sku, 0)
        if qty > available:
            raise OutOfStockError(sku, qty, available)

    # 2) Compute subtotal
    subtotal = sum(prices[sku] * qty for sku, qty in cart.items())

    # 3) Apply coupon
    total = apply_coupon(subtotal, coupon)

    # 4) “Charge” the wallet
    if wallet_cents < total:
        raise PaymentDeclinedError(f"Insufficient funds: need {total}, have {wallet_cents}")

    # 5) Commit inventory changes
    for sku, qty in cart.items():
        inventory[sku] -= qty

    return total


def lab_checkout() -> None:
    print("\n=== LAB: Checkout with custom exceptions ===")

    inventory = {"SKU1": 3, "SKU2": 10}
    prices = {"SKU1": 500, "SKU2": 200}  # cents

    # Scenario 1: out of stock
    cart1 = {"SKU1": 5}
    try:
        checkout(cart1, inventory, prices, coupon=None, wallet_cents=10_000)
    except CheckoutError as e:
        print("Scenario 1 failed:", e)

    # Scenario 2: invalid coupon
    cart2 = {"SKU1": 1, "SKU2": 2}
    try:
        checkout(cart2, inventory, prices, coupon="nope", wallet_cents=10_000)
    except InvalidCouponError as e:
        print("Scenario 2 failed:", e)

    # Scenario 3: successful checkout
    cart3 = {"SKU1": 1, "SKU2": 2}
    try:
        charged = checkout(cart3, inventory, prices, coupon="SAVE10", wallet_cents=1_000)
        print(f"Scenario 3 success: charged {charged} cents")
        print("Remaining inventory:", inventory)
    except CheckoutError as e:
        print("Scenario 3 failed:", e)


def main() -> None:
    lab_checkout()


if __name__ == "__main__":
    main()
