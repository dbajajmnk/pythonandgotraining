"""
Day 5 Topic 4: Logging in Python — practical, production-friendly

Scenario:
- We process salary data.
- Instead of print(), we log:
  - INFO for normal flow
  - WARNING for unusual but recoverable
  - ERROR for failures
- Demonstrates: logger, formatter, file handler, and rotating logs.
"""

from __future__ import annotations

import logging
from logging.handlers import RotatingFileHandler
from pathlib import Path


LOG_DIR = Path("logs_day5")
LOG_FILE = LOG_DIR / "app.log"


def configure_logging() -> None:
    LOG_DIR.mkdir(exist_ok=True)

    logger = logging.getLogger()  # root logger
    logger.setLevel(logging.INFO)

    # Console handler
    ch = logging.StreamHandler()
    ch.setLevel(logging.INFO)

    # File handler with rotation
    fh = RotatingFileHandler(LOG_FILE, maxBytes=200_000, backupCount=3, encoding="utf-8")
    fh.setLevel(logging.INFO)

    fmt = logging.Formatter(
        fmt="%(asctime)s | %(levelname)s | %(name)s | %(message)s",
        datefmt="%Y-%m-%d %H:%M:%S",
    )
    ch.setFormatter(fmt)
    fh.setFormatter(fmt)

    # Avoid duplicate handlers when running multiple times (like in notebooks)
    if not logger.handlers:
        logger.addHandler(ch)
        logger.addHandler(fh)
    else:
        # Replace handlers to keep output clean
        logger.handlers = [ch, fh]


def process_salary(employee_id: str, gross: int, tax_rate: float) -> int:
    log = logging.getLogger("salary")
    log.info("Processing salary for %s", employee_id)

    if gross < 0:
        log.error("Invalid gross salary: %s", gross)
        raise ValueError("gross must be >= 0")

    if not (0.0 <= tax_rate <= 1.0):
        log.warning("Suspicious tax_rate=%s; defaulting to 0.10", tax_rate)
        tax_rate = 0.10

    tax = int(gross * tax_rate)
    net = gross - tax
    log.info("Computed net=%s for %s", net, employee_id)
    return net


def main() -> None:
    configure_logging()

    log = logging.getLogger("main")
    log.info("App started")

    try:
        net1 = process_salary("E001", 100000, 0.10)
        net2 = process_salary("E002", 85000, 5.0)  # triggers warning
        log.info("Results: E001=%s, E002=%s", net1, net2)
    except Exception:
        log.exception("Unexpected failure")  # includes stack trace

    log.info("App finished. Log file: %s", LOG_FILE.resolve())


if __name__ == "__main__":
    main()
