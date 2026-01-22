/*
Day 5 Topic 1: Modules & Packages (Java) — mapping to Python packages

Java concept:
- package = namespace + folder structure
- classes grouped under packages
- build output packaged as JAR
Python concept:
- package = folder with __init__.py (or namespace package)
- module = .py file

Scenario: Office Utilities Toolkit
*/
package day5;

import java.util.Locale;

public class ModulesPackagesDemo {
    public static void main(String[] args) {
        System.out.println("Running: Office Utilities Toolkit (Java mapping)");

        String empCode = "  emp-  009  ";
        String clean = normalizeEmployeeCode(empCode);
        int takeHome = netSalary(100000, 0.10, 2500);

        System.out.println("Clean employee code: " + clean);
        System.out.println("Net salary: " + takeHome);

        String report = buildMonthlyReport(clean, takeHome);
        System.out.println("\n--- Monthly Report ---");
        System.out.println(report);
    }

    // Equivalent of office_utils/textutils.py
    static String normalizeEmployeeCode(String s) {
        String cleaned = s.trim().toUpperCase(Locale.ROOT).replaceAll("\\s+", "");
        cleaned = cleaned.replace("-", "");
        return cleaned; // "EMP009"
    }

    // Equivalent of office_utils/money.py
    static int netSalary(int gross, double taxRate, int deductions) {
        int tax = (int) (gross * taxRate);
        return gross - tax - deductions;
    }

    // Equivalent of office_utils/reports/monthly.py
    static String buildMonthlyReport(String employeeCode, int netSalary) {
        return "Employee: " + employeeCode + "\nNet Salary: " + netSalary + "\nStatus: PROCESSED";
    }
}
