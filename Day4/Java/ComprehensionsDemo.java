/*
ComprehensionsDemo.java

Java equivalents for Python comprehensions:
  - Python: [expr for x in xs if cond]
  - Java:   streams (map/filter/collect) or classic loops

Hands-on lab: clean and summarize sales transactions (same as 02_comprehensions_demo.py)

Compile & run:
  javac ComprehensionsDemo.java
  java ComprehensionsDemo
*/

import java.util.*;
import java.util.stream.*;

public class ComprehensionsDemo {

  static final class Tx {
    final String customer;
    final int amount;
    final String item;
    Tx(String customer, int amount, String item) {
      this.customer = customer;
      this.amount = amount;
      this.item = item;
    }
    @Override public String toString() {
      return "{customer=" + customer + ", amount=" + amount + ", item=" + item + "}";
    }
  }

  static void basicsDemo() {
    System.out.println("\n=== BASICS: map/filter/collect ===");

    List<Integer> nums = Arrays.asList(1, 2, 3, 4, 5, 6);

    List<Integer> squares = nums.stream().map(n -> n * n).collect(Collectors.toList());
    List<Integer> evens = nums.stream().filter(n -> n % 2 == 0).collect(Collectors.toList());
    Set<Integer> uniqueMod3 = nums.stream().map(n -> n % 3).collect(Collectors.toSet());
    Map<Integer, String> asText = nums.stream().collect(Collectors.toMap(
        n -> n,
        n -> "#" + n
    ));
    int sumSquares = nums.stream().mapToInt(n -> n * n).sum();

    System.out.println("Squares: " + squares);
    System.out.println("Evens: " + evens);
    System.out.println("Unique (n % 3): " + uniqueMod3);
    System.out.println("Dict mapping: " + asText);
    System.out.println("Sum of squares: " + sumSquares);
  }

  // -------------------------
  // Hands-on Lab
  // -------------------------

  static List<Tx> normalize(List<Map<String, Object>> raw) {
    List<Tx> cleaned = new ArrayList<>();
    for (Map<String, Object> row : raw) {
      String cust = String.valueOf(row.getOrDefault("customer", "")).trim().toLowerCase(Locale.ROOT);
      if (cust.isEmpty()) continue;

      int amt;
      try {
        Object a = row.get("amount");
        amt = (a instanceof Number) ? ((Number) a).intValue() : Integer.parseInt(String.valueOf(a));
      } catch (Exception e) {
        continue;
      }
      if (amt <= 0) continue;

      String item = String.valueOf(row.getOrDefault("item", ""));
      cleaned.add(new Tx(cust, amt, item));
    }
    return cleaned;
  }

  static void labSalesSummary() {
    System.out.println("\n=== LAB: Sales cleanup + summary ===");

    List<Map<String, Object>> raw = Arrays.asList(
        mapOf("customer", " Alice ", "amount", "120", "item", "Book"),
        mapOf("customer", "bob", "amount", 50, "item", "Pen"),
        mapOf("customer", "", "amount", 10, "item", "Sticker"),
        mapOf("customer", "alice", "amount", -5, "item", "Refund"),
        mapOf("customer", "carol", "amount", "oops", "item", "Bag"),
        mapOf("customer", "bob ", "amount", 70, "item", "Notebook")
    );

    List<Tx> cleaned = normalize(raw);
    System.out.println("Cleaned rows:");
    cleaned.forEach(tx -> System.out.println("  " + tx));

    // 1) List of amounts
    List<Integer> amounts = cleaned.stream().map(tx -> tx.amount).collect(Collectors.toList());
    System.out.println("\nAmounts: " + amounts);

    // 2) Unique customers
    Set<String> customers = cleaned.stream().map(tx -> tx.customer).collect(Collectors.toSet());
    System.out.println("Unique customers: " + customers);

    // 3) Totals per customer
    Map<String, Integer> totals = cleaned.stream().collect(Collectors.groupingBy(
        tx -> tx.customer,
        TreeMap::new, // sorted keys for deterministic output
        Collectors.summingInt(tx -> tx.amount)
    ));
    System.out.println("Totals per customer: " + totals);

    // 4) High-value customers (>=100)
    List<String> highValue = totals.entrySet().stream()
        .filter(e -> e.getValue() >= 100)
        .map(Map.Entry::getKey)
        .collect(Collectors.toList());
    System.out.println("High-value customers (>=100): " + highValue);
  }

  // Small helper to build maps inline (Java 8 friendly)
  static Map<String, Object> mapOf(Object... kv) {
    Map<String, Object> m = new HashMap<>();
    for (int i = 0; i < kv.length; i += 2) {
      m.put(String.valueOf(kv[i]), kv[i + 1]);
    }
    return m;
  }

  public static void main(String[] args) {
    basicsDemo();
    labSalesSummary();
  }
}
