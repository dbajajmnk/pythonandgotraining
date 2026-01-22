/*
 Duck Typing vs Liskov Substitution Principle (LSP) — Java + Python mapping.

 How to run:
   javac DuckTypingVsLSPDemo.java && java DuckTypingVsLSPDemo

 In Java:
   - "Duck typing" is not built-in like Python.
     The closest safe equivalent is programming to an interface (capabilities).
   - LSP is mainly about inheritance contracts:
     If a method expects a Base, any Subclass must behave in a way that
     doesn't break the method's assumptions.

 This file demonstrates:
   1) Capability-based design (interface Notifier) = "Java-style duck typing".
   2) A classic LSP violation: Square extends Rectangle (independent setters contract breaks).
   3) A safe alternative: Shape interface + separate Rectangle/Square implementations.
*/

import java.util.*;

public class DuckTypingVsLSPDemo {

  // ==========================================================
  // 1) "JAVA-STYLE DUCK TYPING" (CAPABILITY VIA INTERFACE)
  // ==========================================================

  interface Notifier {
    void send(String message);
  }

  static class EmailNotifier implements Notifier {
    private final String email;
    EmailNotifier(String email) { this.email = email; }
    @Override public void send(String message) {
      System.out.println("[EMAIL to " + email + "] " + message);
    }
  }

  static class SmsNotifier implements Notifier {
    private final String phone;
    SmsNotifier(String phone) { this.phone = phone; }
    @Override public void send(String message) {
      System.out.println("[SMS to " + phone + "] " + message);
    }
  }

  static class SlackNotifier implements Notifier {
    private final String channel;
    SlackNotifier(String channel) { this.channel = channel; }
    @Override public void send(String message) {
      System.out.println("[SLACK #" + channel + "] " + message);
    }
  }

  static void alertUser(Notifier notifier, String message) {
    // We don't care about the class name — only that it can send.
    notifier.send(message);
  }

  static void demoCapabilityDesign() {
    System.out.println("\n=== Demo 1: Capability-based design (interface) ===");
    List<Notifier> notifiers = Arrays.asList(
        new EmailNotifier("deepak@example.com"),
        new SmsNotifier("+91-99999-11111"),
        new SlackNotifier("support")
    );
    for (Notifier n : notifiers) {
      alertUser(n, "Your order is shipped.");
    }
  }

  // // ==========================================================
  // // 2) LSP VIOLATION DEMO (SQUARE EXTENDS RECTANGLE)
  // // ==========================================================

  // static class Rectangle {
  //   protected int width;
  //   protected int height;

  //   Rectangle(int width, int height) {
  //     this.width = width;
  //     this.height = height;
  //   }

  //   public void setWidth(int w) { this.width = w; }
  //   public void setHeight(int h) { this.height = h; }
  //   public int area() { return width * height; }
  // }

  // static class Square extends Rectangle {
  //   Square(int side) { super(side, side); }

  //   @Override public void setWidth(int w) {
  //     this.width = w;
  //     this.height = w;
  //   }

  //   @Override public void setHeight(int h) {
  //     this.height = h;
  //     this.width = h;
  //   }
  // }

  // static int resizeAndMeasure(Rectangle r) {
  //   // This method assumes that width and height are independent.
  //   r.setWidth(5);
  //   r.setHeight(2);
  //   return r.area(); // expected 10
  // }

  // static void demoLspViolation() {
  //   System.out.println("\n=== Demo 2: LSP violation (Square extends Rectangle) ===");

  //   Rectangle rect = new Rectangle(1, 1);
  //   Rectangle square = new Square(1);

  //   int rectArea = resizeAndMeasure(rect);
  //   int squareArea = resizeAndMeasure(square);

  //   System.out.println("Rectangle area expected=10, got=" + rectArea);
  //   System.out.println("Square area expected=10, got=" + squareArea + "  <-- violates expectation");

  //   System.out.println("\nWhy this violates LSP:");
  //   System.out.println("- resizeAndMeasure assumes independent setters");
  //   System.out.println("- Square changes the meaning of setWidth/setHeight");
  //   System.out.println("- so Square can't safely substitute Rectangle in this API");
  // }

  // // ==========================================================
  // // 3) LSP-SAFE DESIGN (SHARED INTERFACE + SEPARATE TYPES)
  // // ==========================================================

  // interface Shape {
  //   int area();
  // }

  // static class SafeRectangle implements Shape {
  //   private final int width;
  //   private final int height;
  //   SafeRectangle(int width, int height) { this.width = width; this.height = height; }
  //   @Override public int area() { return width * height; }
  // }

  // static class SafeSquare implements Shape {
  //   private final int side;
  //   SafeSquare(int side) { this.side = side; }
  //   @Override public int area() { return side * side; }
  // }

  // static int totalArea(List<Shape> shapes) {
  //   int sum = 0;
  //   for (Shape s : shapes) sum += s.area();
  //   return sum;
  // }

  // static void demoLspSafeDesign() {
  //   System.out.println("\n=== Demo 3: LSP-safe design (Shape + separate types) ===");
  //   List<Shape> shapes = Arrays.asList(new SafeRectangle(5, 2), new SafeSquare(3));
  //   System.out.println("Total area: " + totalArea(shapes));
  //   System.out.println("\nWhy this is safe:");
  //   System.out.println("- Rectangle and Square share only true common behavior: area()");
  //   System.out.println("- No misleading setters that break assumptions");
  // }

  // // ==========================================================
  // // Hands-on Lab (in main output)
  // // ==========================================================

  // static void labPrompt() {
  //   System.out.println("\n=== Hands-on Lab (Scenario) ===");
  //   System.out.println("Scenario: Notification system for an e-commerce app");
  //   System.out.println("Requirement:");
  //   System.out.println("  - You can send updates via Email/SMS/Slack today.");
  //   System.out.println("  - Tomorrow, you might add WhatsApp / PushNotification.");
  //   System.out.println("Task:");
  //   System.out.println("  1) Add a new notifier WITHOUT changing alertUser().");
  //   System.out.println("  2) Ensure any subtype you pass to business methods does not break assumptions (LSP).");
  //   System.out.println("Hint:");
  //   System.out.println("  - Use interface Notifier for capability-based design.");
  //   System.out.println("  - Avoid inheritance where the child changes parent meaning.");
  // }

  public static void main(String[] args) {
    demoCapabilityDesign();
    // demoLspViolation();
    // demoLspSafeDesign();
    // labPrompt();
  }
}
