/*
Day 5 Topic 5: Virtual environments — Java mapping

Python venv isolates packages per project.
Java equivalent ideas:
- Separate project builds with Maven/Gradle
- Dependencies are per project (pom.xml / build.gradle)
- Use different JDK versions via SDKMAN/Toolchains
- Use containers for full isolation

This file prints guidance only.
*/
package day5;

public class VirtualEnvironmentsDemo {
    public static void main(String[] args) {
        System.out.println("Python venv ≈ per-project dependency isolation.");
        System.out.println("Java equivalent: Maven/Gradle manage dependencies per project.");
        System.out.println();
        System.out.println("Python steps:");
        System.out.println("  python -m venv .venv");
        System.out.println("  activate");
        System.out.println("  pip install -r requirements.txt");
        System.out.println();
        System.out.println("Java steps (conceptual):");
        System.out.println("  create new Maven/Gradle project");
        System.out.println("  add dependencies in pom.xml/build.gradle");
        System.out.println("  build produces an artifact (JAR) with managed deps");
    }
}
