/*
ExceptionsDemo.java

Exception handling in Java:
  - try / catch / finally
  - exception chaining (new Exception(msg, cause))

Hands-on lab: robust import of "id,amount" lines (same as 03_exceptions_demo.py)

Compile & run:
  javac ExceptionsDemo.java
  java ExceptionsDemo
*/

import java.util.*;

public class ExceptionsDemo {

  static int parseAmount(String text) {
    try {
      return Integer.parseInt(text);
    } catch (NumberFormatException e) {
      // Keep the cause (chaining)
      throw new IllegalArgumentException("Bad amount: " + text, e);
    } finally {
      // finally always runs
    }
  }

  static void basicsDemo() {
    System.out.println("\n=== BASICS: try/catch/finally ===");
    for (String s : Arrays.asList("42", "oops")) {
      try {
        System.out.println("Parsing " + s + " -> " + parseAmount(s));
      } catch (IllegalArgumentException e) {
        System.out.println("  ERROR: " + e.getMessage());
      }
    }
  }

  // -------------------------
  // Hands-on Lab
  // -------------------------

  static Map<Integer, Integer> importLines(List<String> lines) {
    Map<Integer, Integer> totals = new HashMap<>();
    int bad = 0;
    int processed = 0;

    try {
      for (String raw : lines) {
        processed++;
        if (raw == null) {
          bad++;
          continue;
        }
        raw = raw.trim();
        if (raw.isEmpty()) {
          bad++;
          continue;
        }

        try {
          String[] parts = raw.split(",");
          if (parts.length != 2) throw new IllegalArgumentException("wrong parts");
          int id = Integer.parseInt(parts[0]);
          int amt = parseAmount(parts[1]);
          totals.put(id, totals.getOrDefault(id, 0) + amt);
        } catch (Exception e) {
          bad++;
          // In real apps: log e + the raw line
        }
      }
      return totals;
    } finally {
      System.out.println("\n[FINAL REPORT]");
      System.out.println("  processed: " + processed);
      System.out.println("  bad lines: " + bad);
      System.out.println("  valid ids : " + totals.size());
    }
  }

  static void labImport() {
    System.out.println("\n=== LAB: Robust import of 'id,amount' ===");
    List<String> lines = Arrays.asList(
        "101,250",
        "102,90",
        "bad-line",
        "103,not-a-number",
        "101,10",
        ""
    );
    Map<Integer, Integer> totals = importLines(lines);

    System.out.println("\nTotals:");
    totals.keySet().stream().sorted().forEach(k ->
        System.out.println("  " + k + ": " + totals.get(k))
    );
  }

  public static void main(String[] args) {
    basicsDemo();
    labImport();
  }
}
