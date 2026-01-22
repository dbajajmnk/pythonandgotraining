/*
01_CollectionsDemo.java

Java equivalents for Python list/tuple/set/dict:
  - List  (ArrayList)        ~ Python list
  - “Tuple” (small class)    ~ Python tuple (fixed bundle)
  - Set   (HashSet)          ~ Python set
  - Map   (HashMap)          ~ Python dict

Hands-on lab: conference registration (same scenario as 01_collections_demo.py)

Compile & run:
  javac CollectionsDemo.java
  java CollectionsDemo
*/

import java.util.*;

public class CollectionsDemo {

  // A tiny immutable "tuple" replacement (row, seatNumber)
  static final class Seat {
    final String row;
    final int number;
    Seat(String row, int number) {
      this.row = row;
      this.number = number;
    }
    @Override public String toString() {
      return "(" + row + ", " + number + ")";
    }
  }

  static void basicsDemo() {
    System.out.println("\n=== BASICS: List / Seat(tuple) / Set / Map ===");

    // LIST: ordered, mutable, allows duplicates
    List<String> shoppingList = new ArrayList<>(Arrays.asList(
        "milk", "bread", "eggs", "bread"
    ));
    shoppingList.add("coffee");
    System.out.println("List (ordered, duplicates ok): " + shoppingList);

    // "TUPLE": fixed bundle (row, seatNumber)
    Seat seat = new Seat("Row A", 12);
    System.out.println("Seat (fixed bundle): " + seat);

    // SET: unique items
    Set<String> uniqueItems = new HashSet<>(shoppingList);
    System.out.println("Set (unique): " + uniqueItems);
    System.out.println("Is 'milk' in set? " + uniqueItems.contains("milk"));

    // MAP: key -> value
    Map<String, Double> prices = new HashMap<>();
    prices.put("milk", 2.50);
    prices.put("bread", 1.25);
    prices.put("eggs", 3.10);
    prices.put("coffee", 8.99);
    System.out.println("Map (key -> value): " + prices);
    System.out.println("Price of eggs: " + prices.get("eggs"));
  }

  // -------------------------
  // Hands-on Lab
  // -------------------------

  static Map<String, Seat> assignSeats(List<String> emailsInOrder, int capacity, List<String> waitlistOut) {
    // LinkedHashMap keeps insertion order (nice for deterministic printing)
    Map<String, Seat> seatMap = new LinkedHashMap<>();
    Set<String> seen = new HashSet<>();

    for (String raw : emailsInOrder) {
      if (raw == null) continue;
      String email = raw.trim().toLowerCase(Locale.ROOT);
      if (email.isEmpty() || !email.contains("@")) continue;
      if (seen.contains(email)) continue;
      seen.add(email);

      if (seatMap.size() < capacity) {
        int i = seatMap.size();
        char rowLetter = (char) ('A' + (i / 10));
        int seatNumber = (i % 10) + 1;
        seatMap.put(email, new Seat("Row " + rowLetter, seatNumber));
      } else {
        waitlistOut.add(email);
      }
    }

    return seatMap;
  }

  static void labConferenceRegistration() {
    System.out.println("\n=== LAB: Conference registration ===");
    List<String> signups = Arrays.asList(
        "Alice@Example.com",
        "bob@example.com",
        "ALICE@example.com", // duplicate
        "carol@example.com",
        "dave@example.com",
        "",                 // bad
        "not-an-email",     // bad
        "erin@example.com"
    );
    int capacity = 4;

    List<String> waitlist = new ArrayList<>();
    Map<String, Seat> seatMap = assignSeats(signups, capacity, waitlist);

    System.out.println("Capacity: " + capacity);
    System.out.println("\nSeat assignments (Map: email -> Seat)");
    for (Map.Entry<String, Seat> e : seatMap.entrySet()) {
      System.out.printf("  %-18s -> %s%n", e.getKey(), e.getValue());
    }
    System.out.println("\nWaitlist (List): " + waitlist);

    Set<String> unique = new HashSet<>();
    unique.addAll(seatMap.keySet());
    unique.addAll(waitlist);
    System.out.println("\nUnique registrants (Set size): " + unique.size());
  }

  public static void main(String[] args) {
    basicsDemo();
    labConferenceRegistration();
  }
}
