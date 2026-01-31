// File: UserServiceRefactor.java
// Topic: Structs vs Java Classes — Side-by-side Refactor Example (Java Version)
//
// Scenario: Register a user, prevent duplicates, and send a welcome notification.
//
// Key Java traits:
// - Class bundles state + behavior
// - Explicit encapsulation via private/public
// - Exceptions often used (here we keep checked/unchecked simple)
//
// NOTE: This is a simplified, framework-free example.

import java.util.HashMap;
import java.util.Map;

public class UserServiceRefactor {

    // --- Domain model (class = data + behavior + lifecycle) ---
    static class User {
        private final String name;
        private String email;

        public User(String name, String email) {
            if (name == null || name.trim().isEmpty()) throw new IllegalArgumentException("name empty");
            if (email == null || !email.contains("@")) throw new IllegalArgumentException("email invalid");
            this.name = name.trim();
            this.email = email.trim();
        }

        public String getName() { return name; }
        public String getEmail() { return email; }

        public void setEmail(String email) {
            if (email == null || !email.contains("@")) throw new IllegalArgumentException("email invalid");
            this.email = email.trim();
        }
    }

    // --- Interface + explicit implements ---
    interface Notifier {
        void notify(User to, String message);
    }

    static class EmailNotifier implements Notifier {
        @Override public void notify(User to, String message) {
            System.out.println("Email sent to " + to.getEmail() + ": \"" + message + "\"");
        }
    }

    // --- Repository ---
    interface UserRepo {
        void save(User u);
        User findByEmail(String email);
    }

    static class InMemoryUserRepo implements UserRepo {
        private final Map<String, User> store = new HashMap<>();

        @Override public void save(User u) {
            String key = u.getEmail();
            if (store.containsKey(key)) throw new IllegalStateException("duplicate user: " + key);
            store.put(key, u);
        }

        @Override public User findByEmail(String email) {
            return store.get(email);
        }
    }

    // --- Service class (logic inside class) ---
    static class UserService {
        private final UserRepo repo;
        private final Notifier notifier;

        public UserService(UserRepo repo, Notifier notifier) {
            this.repo = repo;
            this.notifier = notifier;
        }

        public void register(String name, String email) {
            User u = new User(name, email);
            repo.save(u);
            notifier.notify(u, "Welcome!");
        }

        public User getByEmail(String email) {
            User u = repo.findByEmail(email);
            if (u == null) throw new IllegalStateException("user not found: " + email);
            return u;
        }
    }

    // --- Demo ---
    public static void main(String[] args) {
        UserRepo repo = new InMemoryUserRepo();
        Notifier notifier = new EmailNotifier();
        UserService svc = new UserService(repo, notifier);

        svc.register("Deepak", "deepak@example.com");

        try {
            svc.register("Deepak", "deepak@example.com"); // duplicate
        } catch (Exception e) {
            System.out.println("Expected error: " + e.getMessage());
        }

        try {
            svc.getByEmail("missing@example.com");
        } catch (Exception e) {
            System.out.println("Expected error: " + e.getMessage());
        }
    }
}
