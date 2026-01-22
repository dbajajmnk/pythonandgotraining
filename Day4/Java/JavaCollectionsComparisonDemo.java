/*
JavaCollectionsComparisonDemo.java

Python collections vs Java Collections Framework (practical mapping).

Hands-on lab: student roster management.
  - Input: (student, className) pairs; duplicates possible
  - Build: Map<className, Set<student>>
  - Print classes and students in sorted order

Compile & run:
  javac JavaCollectionsComparisonDemo.java
  java JavaCollectionsComparisonDemo
*/

import java.util.*;

public class JavaCollectionsComparisonDemo {

  static void quickMapping() {
    System.out.println("\n=== Quick mapping ===");
    System.out.println("Python list  -> Java List (e.g., ArrayList)");
    System.out.println("Python tuple -> Java immutable bundle (List.of / small class)");
    System.out.println("Python set   -> Java Set (e.g., HashSet)");
    System.out.println("Python dict  -> Java Map (e.g., HashMap)");
  }

  static void labStudentRoster() {
    System.out.println("\n=== LAB: Student roster management ===");

    List<String[]> enrollments = Arrays.asList(
        new String[]{"Alice", "Math"},
        new String[]{"Bob", "Math"},
        new String[]{"Alice", "Math"}, // duplicate
        new String[]{"Carol", "Science"},
        new String[]{"Bob", "Science"},
        new String[]{"Dave", "Math"},
        new String[]{"Alice", "Science"}
    );

    // Core structure: Map<className, Set<student>>
    Map<String, Set<String>> roster = new HashMap<>();
    for (String[] pair : enrollments) {
      String student = pair[0];
      String cls = pair[1];
      roster.computeIfAbsent(cls, k -> new HashSet<>()).add(student);
    }

    // Deterministic output: sort keys and values
    Set<String> allStudents = new HashSet<>();
    List<String> classes = new ArrayList<>(roster.keySet());
    Collections.sort(classes);

    for (String cls : classes) {
      List<String> students = new ArrayList<>(roster.get(cls));
      Collections.sort(students);
      allStudents.addAll(roster.get(cls));
      System.out.println(cls + ": " + students);
    }
    System.out.println("Total unique students: " + allStudents.size());
  }

  public static void main(String[] args) {
    quickMapping();
    labStudentRoster();
  }
}
