/*
Day 5 Topic 3: File I/O (text, CSV, JSON) — Java version of same scenario

Scenario:
- Read employees.csv
- Read tax_rates.json
- Compute net salary
- Write salary_report.txt

NOTE: JSON in Java typically uses a library (Jackson/Gson).
To keep this file dependency-free, we'll parse a very simple JSON object manually.
In real projects, use Jackson (com.fasterxml.jackson) or Gson.
*/
package day5;

import java.io.*;
import java.nio.charset.StandardCharsets;
import java.nio.file.*;
import java.util.*;

public class FileIODemo {
    static final Path DATA_DIR = Paths.get("data_day5_java");
    static final Path EMP_CSV = DATA_DIR.resolve("employees.csv");
    static final Path TAX_JSON = DATA_DIR.resolve("tax_rates.json");
    static final Path REPORT_TXT = DATA_DIR.resolve("salary_report.txt");

    public static void main(String[] args) throws Exception {
        ensureSampleFiles();

        List<Employee> employees = readEmployees();
        Map<String, Double> taxRates = readSimpleJsonTaxRates();

        List<String> lines = new ArrayList<>();
        lines.add("Salary Report");
        lines.add("============");
        lines.add("");

        for (Employee e : employees) {
            double rate = taxRates.getOrDefault(e.id, 0.10);
            int net = netSalary(e.gross, rate, 2500);
            lines.add(e.id + " | " + e.name + " | gross=" + e.gross + " | tax=" + (int)(rate*100) + "% | net=" + net);
        }

        Files.write(REPORT_TXT, String.join("\n", lines).concat("\n").getBytes(StandardCharsets.UTF_8));
        System.out.println("Report created: " + REPORT_TXT.toAbsolutePath());
        System.out.println(new String(Files.readAllBytes(REPORT_TXT), StandardCharsets.UTF_8));
    }

    static void ensureSampleFiles() throws IOException {
        if (!Files.exists(DATA_DIR)) Files.createDirectories(DATA_DIR);

        if (!Files.exists(EMP_CSV)) {
            List<String> csv = Arrays.asList(
                "id,name,gross",
                "E001,Asha,100000",
                "E002,Ravi,85000",
                "E003,Neha,120000"
            );
            Files.write(EMP_CSV, String.join("\n", csv).concat("\n").getBytes(StandardCharsets.UTF_8));
        }

        if (!Files.exists(TAX_JSON)) {
            String json = "{\n  \"E001\": 0.10,\n  \"E002\": 0.08,\n  \"E003\": 0.12\n}\n";
            Files.write(TAX_JSON, json.getBytes(StandardCharsets.UTF_8));
        }
    }

    static List<Employee> readEmployees() throws IOException {
        List<Employee> out = new ArrayList<>();
        List<String> lines = Files.readAllLines(EMP_CSV, StandardCharsets.UTF_8);
        for (int i = 1; i < lines.size(); i++) { // skip header
            String[] parts = lines.get(i).split(",");
            out.add(new Employee(parts[0], parts[1], Integer.parseInt(parts[2])));
        }
        return out;
    }

    // VERY simple JSON parser: expects {"E001":0.10, ...}
    static Map<String, Double> readSimpleJsonTaxRates() throws IOException {
        String s = new String(Files.readAllBytes(TAX_JSON), StandardCharsets.UTF_8);
        s = s.replaceAll("[\\{\\}\\s\"]", ""); // remove { } spaces and quotes
        Map<String, Double> map = new HashMap<>();
        if (s.isEmpty()) return map;

        for (String pair : s.split(",")) {
            String[] kv = pair.split(":");
            map.put(kv[0], Double.parseDouble(kv[1]));
        }
        return map;
    }

    static int netSalary(int gross, double taxRate, int deductions) {
        int tax = (int) (gross * taxRate);
        return gross - tax - deductions;
    }

    static class Employee {
        final String id;
        final String name;
        final int gross;
        Employee(String id, String name, int gross) {
            this.id = id; this.name = name; this.gross = gross;
        }
    }
}
