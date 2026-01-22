/*
Day 5 Topic 4: Logging — Java mapping

Python: logging module
Java: java.util.logging (JUL) or SLF4J + Logback in real projects

This demo uses java.util.logging for zero dependencies.
*/
package day5;

import java.io.IOException;
import java.nio.file.*;
import java.util.logging.*;

public class LoggingDemo {
    static final Logger LOG = Logger.getLogger("salary");

    public static void main(String[] args) throws Exception {
        configureLogging();

        LOG.info("App started");
        try {
            int net1 = processSalary("E001", 100000, 0.10);
            int net2 = processSalary("E002", 85000, 5.0); // warning
            LOG.info("Results: E001=" + net1 + ", E002=" + net2);
        } catch (Exception e) {
            LOG.log(Level.SEVERE, "Unexpected failure", e);
        }
        LOG.info("App finished");
    }

    static void configureLogging() throws IOException {
        Path logDir = Paths.get("logs_day5_java");
        Files.createDirectories(logDir);
        Path logFile = logDir.resolve("app.log");

        LOG.setUseParentHandlers(false);
        LOG.setLevel(Level.INFO);

        // Console
        ConsoleHandler ch = new ConsoleHandler();
        ch.setLevel(Level.INFO);

        // File
        FileHandler fh = new FileHandler(logFile.toString(), true);
        fh.setLevel(Level.INFO);

        SimpleFormatter fmt = new SimpleFormatter();
        ch.setFormatter(fmt);
        fh.setFormatter(fmt);

        LOG.addHandler(ch);
        LOG.addHandler(fh);

        LOG.info("Log file: " + logFile.toAbsolutePath());
    }

    static int processSalary(String employeeId, int gross, double taxRate) {
        LOG.info("Processing salary for " + employeeId);

        if (gross < 0) {
            LOG.severe("Invalid gross salary: " + gross);
            throw new IllegalArgumentException("gross must be >= 0");
        }
        if (taxRate < 0.0 || taxRate > 1.0) {
            LOG.warning("Suspicious taxRate=" + taxRate + "; defaulting to 0.10");
            taxRate = 0.10;
        }
        int tax = (int) (gross * taxRate);
        int net = gross - tax;
        LOG.info("Computed net=" + net + " for " + employeeId);
        return net;
    }
}
