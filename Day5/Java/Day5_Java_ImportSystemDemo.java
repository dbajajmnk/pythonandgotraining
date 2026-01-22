/*
Day 5 Topic 2: Import system mapping (Java vs Python)

Java:
- import is compile-time name resolution for classes
- classpath decides where classes are found
Python:
- import executes module code and caches it
- sys.path decides where modules are found
*/
package day5;

import java.util.Arrays;

public class ImportSystemDemo {
    public static void main(String[] args) {
        System.out.println("Java import is not dynamic execution like Python.");
        System.out.println("JVM classpath plays a role similar to Python sys.path.");

        // Example: using a class from java.util
        System.out.println("Arrays.asList demo: " + Arrays.asList("A", "B", "C"));
    }
}
