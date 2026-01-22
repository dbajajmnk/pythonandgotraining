/*
CustomExceptionsDemo.java

Custom exceptions in Java:
  - define your own exception classes
  - throw them for domain failures

This demo uses CHECKED exceptions (extends Exception) so callers must handle/declare.
In real systems, you might choose unchecked (extends RuntimeException) depending on policy.

Hands-on lab: e-commerce checkout (same scenario as 04_custom_exceptions_demo.py)

Compile & run:
  javac CustomExceptionsDemo.java
  java CustomExceptionsDemo
*/

import java.util.*;

public class CustomExceptionsDemo {

  // Base domain exception
  static class CheckoutException extends Exception {
    CheckoutException(String message) { super(message); }
    CheckoutException(String message, Throwable cause) { super(message, cause); }
  }

  static class InvalidCouponException extends CheckoutException {
    InvalidCouponException(String message) { super(message); }
  }

  static class OutOfStockException extends CheckoutException {
    final String sku;
    final int requested;
    final int available;
    OutOfStockException(String sku, int requested, int available) {
      super("Out of stock: " + sku + " (requested " + requested + ", available " + available + ")");
      this.sku = sku;
      this.requested = requested;
      this.available = available;
    }
  }

  static class PaymentDeclinedException extends CheckoutException {
    PaymentDeclinedException(String message) { super(message); }
  }

  static int applyCoupon(int subtotalCents, String coupon) throws InvalidCouponException {
    if (coupon == null || coupon.trim().isEmpty()) return subtotalCents;
    String c = coupon.trim().toUpperCase(Locale.ROOT);
    if (c.equals("SAVE10")) {
      return (int) Math.floor(subtotalCents * 0.90);
    }
    throw new InvalidCouponException("Unknown coupon: " + c);
  }

  static int checkout(
      Map<String, Integer> cart,
      Map<String, Integer> inventory,
      Map<String, Integer> prices,
      String coupon,
      int walletCents
  ) throws CheckoutException {

    // 1) Validate inventory
    for (Map.Entry<String, Integer> e : cart.entrySet()) {
      String sku = e.getKey();
      int qty = e.getValue();
      int available = inventory.getOrDefault(sku, 0);
      if (qty > available) {
        throw new OutOfStockException(sku, qty, available);
      }
    }

    // 2) Compute subtotal
    int subtotal = 0;
    for (Map.Entry<String, Integer> e : cart.entrySet()) {
      String sku = e.getKey();
      int qty = e.getValue();
      subtotal += prices.get(sku) * qty;
    }

    // 3) Apply coupon
    int total = applyCoupon(subtotal, coupon);

    // 4) “Charge”
    if (walletCents < total) {
      throw new PaymentDeclinedException("Insufficient funds: need " + total + ", have " + walletCents);
    }

    // 5) Commit inventory changes
    for (Map.Entry<String, Integer> e : cart.entrySet()) {
      String sku = e.getKey();
      int qty = e.getValue();
      inventory.put(sku, inventory.get(sku) - qty);
    }

    return total;
  }

  static void labCheckout() {
    System.out.println("\n=== LAB: Checkout with custom exceptions ===");

    Map<String, Integer> inventory = new HashMap<>();
    inventory.put("SKU1", 3);
    inventory.put("SKU2", 10);

    Map<String, Integer> prices = new HashMap<>();
    prices.put("SKU1", 500); // cents
    prices.put("SKU2", 200);

    // Scenario 1: out of stock
    Map<String, Integer> cart1 = new HashMap<>();
    cart1.put("SKU1", 5);
    try {
      checkout(cart1, inventory, prices, null, 10_000);
    } catch (CheckoutException e) {
      System.out.println("Scenario 1 failed: " + e.getMessage());
    }

    // Scenario 2: invalid coupon
    Map<String, Integer> cart2 = new HashMap<>();
    cart2.put("SKU1", 1);
    cart2.put("SKU2", 2);
    try {
      checkout(cart2, inventory, prices, "nope", 10_000);
    } catch (InvalidCouponException e) {
      System.out.println("Scenario 2 failed: " + e.getMessage());
    } catch (CheckoutException e) {
      System.out.println("Scenario 2 failed (other): " + e.getMessage());
    }

    // Scenario 3: success
    Map<String, Integer> cart3 = new HashMap<>();
    cart3.put("SKU1", 1);
    cart3.put("SKU2", 2);
    try {
      int charged = checkout(cart3, inventory, prices, "SAVE10", 1_000);
      System.out.println("Scenario 3 success: charged " + charged + " cents");
      System.out.println("Remaining inventory: " + inventory);
    } catch (CheckoutException e) {
      System.out.println("Scenario 3 failed: " + e.getMessage());
    }
  }

  public static void main(String[] args) {
    labCheckout();
  }
}
