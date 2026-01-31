// File: JavaThreadDemo.java
// Topic: Go Concurrency vs Java Threads (Java side)
// Compile: javac JavaThreadDemo.java
// Run: java JavaThreadDemo

public class JavaThreadDemo {

    public static void main(String[] args) throws InterruptedException {
        System.out.println("=== Java Threads Demo ===");

        Thread[] threads = new Thread[5];

        for (int i = 0; i < 5; i++) {
            int id = i + 1;
            threads[i] = new Thread(() -> {
                try {
                    Thread.sleep(200);
                    System.out.println("Worker " + id + " done");
                } catch (InterruptedException e) {
                    Thread.currentThread().interrupt();
                }
            });
            threads[i].start();
        }

        for (Thread t : threads) {
            t.join();
        }

        System.out.println("All threads completed");
    }
}
