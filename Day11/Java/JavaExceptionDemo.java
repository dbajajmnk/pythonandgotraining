// File: JavaExceptionDemo.java
// Topic: Comparison with Go Error Handling
// Compile: javac JavaExceptionDemo.java
// Run: java JavaExceptionDemo

public class JavaExceptionDemo {

    public static void main(String[] args) {
        System.out.println("=== Java Exceptions Demo ===");

        // Lab 1: Basic try-catch
        try {
            int result = divide(10, 0);
            System.out.println("Result: " + result);
        } catch (ArithmeticException e) {
            System.out.println("Caught exception: " + e.getMessage());
        }

        // Lab 2: Propagation with throws
        try {
            serviceLayer();
        } catch (Exception e) {
            System.out.println("Service error: " + e.getMessage());
        }

        // Lab 3: Runtime exception (unchecked)
        try {
            triggerBug();
        } catch (RuntimeException e) {
            System.out.println("Runtime exception caught: " + e.getMessage());
        }

        System.out.println("=== Done ===");
    }

    // Checked-style handling via exception
    static int divide(int a, int b) {
        if (b == 0) {
            throw new ArithmeticException("division by zero");
        }
        return a / b;
    }

    // Propagation using throws
    static void serviceLayer() throws Exception {
        repositoryLayer();
    }

    static void repositoryLayer() throws Exception {
        throw new Exception("database unavailable");
    }

    // Unchecked exception example
    static void triggerBug() {
        throw new RuntimeException("programmer bug");
    }
}
