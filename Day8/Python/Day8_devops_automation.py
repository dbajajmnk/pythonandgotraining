"""
Day 8 — Section 4: Python in DevOps & Automation (Detailed) — Lab Toolkit
========================================================================

This file builds a mini DevOps CLI toolkit (real-world style).

------------------------------------------------------------
INSTALL
------------------------------------------------------------
No external packages needed (Python 3.9+).

------------------------------------------------------------
RUN
------------------------------------------------------------
python day8_devops_automation.py --help

Examples:
1) Healthcheck a URL (exit code non-zero on failure)
   python Day8_devops_automation.py healthcheck --url http://127.0.0.1:5000/health/db

2) Backup a folder to zip and keep last 5
   python day8_devops_automation.py backup --src . --dest backups --keep 5

3) Summarize a log file
   python day8_devops_automation.py logsummary --file sample.log

4) Run a command safely
   python day8_devops_automation.py run -- cmd /c dir       (Windows)
   python day8_devops_automation.py run -- ls -la           (Linux/macOS)

WHY THIS MATTERS:
- In CI/CD and operations, scripts must be: configurable, logged, and reliable.
"""

from __future__ import annotations

import argparse
import datetime as dt
import json
import logging
import os
import shutil
import subprocess
import sys
import urllib.request
import zipfile
from pathlib import Path


# ============================================================
# Logging setup (production scripts use logging, not print)
# ============================================================
logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s %(levelname)s %(message)s",
)
log = logging.getLogger("devops_toolkit")


# ============================================================
# Subcommand: healthcheck
# ============================================================
def cmd_healthcheck(url: str, timeout: float = 3.0) -> int:
    """
    WHAT: Calls an HTTP endpoint and checks status code.
    WHY: Used in readiness/liveness checks, smoke tests, deployment gates.
    WHEN: After deploy, before routing traffic.
    """
    try:
        req = urllib.request.Request(url, headers={"User-Agent": "python-healthcheck"})
        with urllib.request.urlopen(req, timeout=timeout) as resp:
            status = resp.status
            body = resp.read(200).decode("utf-8", errors="ignore")
            log.info("Healthcheck status=%s body_snippet=%r", status, body)
            return 0 if status == 200 else 2
    except Exception as e:
        log.error("Healthcheck failed: %s", e)
        return 2


# ============================================================
# Subcommand: backup (zip folder, keep last N)
# ============================================================
def cmd_backup(src: Path, dest: Path, keep: int = 5) -> int:
    """
    WHAT: Zip a folder as a timestamped backup.
    WHY: common for logs/config artifacts before changes.
    WHEN: pre-deploy backups, nightly backups (cron).
    """
    src = src.resolve()
    dest.mkdir(parents=True, exist_ok=True)

    ts = dt.datetime.utcnow().strftime("%Y%m%d_%H%M%S")
    zip_name = dest / f"backup_{ts}.zip"

    log.info("Creating backup: %s -> %s", src, zip_name)
    with zipfile.ZipFile(zip_name, "w", zipfile.ZIP_DEFLATED) as z:
        for p in src.rglob("*"):
            if p.is_file():
                z.write(p, arcname=p.relative_to(src))

    # Rotate old backups
    backups = sorted(dest.glob("backup_*.zip"))
    if len(backups) > keep:
        for old in backups[: len(backups) - keep]:
            log.info("Removing old backup: %s", old)
            old.unlink()

    log.info("Backup complete. Total backups kept: %s", min(len(backups), keep))
    return 0


# ============================================================
# Subcommand: logsummary
# ============================================================
def cmd_logsummary(file: Path) -> int:
    """
    WHAT: Parse a log file and count ERROR/CRITICAL lines.
    WHY: daily ops summary, alerting signals, postmortems.
    """
    if not file.exists():
        log.error("Log file not found: %s", file)
        return 2

    counts = {"ERROR": 0, "CRITICAL": 0, "WARN": 0, "INFO": 0}
    with file.open("r", encoding="utf-8", errors="ignore") as f:
        for line in f:
            for k in list(counts.keys()):
                if k in line:
                    counts[k] += 1

    print(json.dumps(counts, indent=2))
    return 0


# ============================================================
# Subcommand: run (execute OS command)
# ============================================================
def cmd_run(cmd: list[str]) -> int:
    """
    WHAT: Safely run a command and capture output.
    WHY: automation often wraps OS tools (docker, kubectl, terraform).
    """
    log.info("Running command: %s", cmd)
    r = subprocess.run(cmd, capture_output=True, text=True)
    print(r.stdout)
    if r.stderr:
        print(r.stderr, file=sys.stderr)
    return r.returncode


def build_parser() -> argparse.ArgumentParser:
    p = argparse.ArgumentParser(prog="devops_toolkit", description="DevOps automation toolkit demo")
    sub = p.add_subparsers(dest="command", required=True)

    # healthcheck
    p_h = sub.add_parser("healthcheck", help="Hit a URL and return non-zero if not 200")
    p_h.add_argument("--url", required=True)
    p_h.add_argument("--timeout", type=float, default=3.0)

    # backup
    p_b = sub.add_parser("backup", help="Zip a folder and keep last N backups")
    p_b.add_argument("--src", required=True)
    p_b.add_argument("--dest", default="backups")
    p_b.add_argument("--keep", type=int, default=5)

    # logsummary
    p_l = sub.add_parser("logsummary", help="Summarize a log file")
    p_l.add_argument("--file", required=True)

    # run
    p_r = sub.add_parser("run", help="Run an OS command safely")
    p_r.add_argument("cmd", nargs=argparse.REMAINDER, help="Command after --")

    return p


def main() -> int:
    parser = build_parser()
    args = parser.parse_args()

    if args.command == "healthcheck":
        return cmd_healthcheck(args.url, args.timeout)

    if args.command == "backup":
        return cmd_backup(Path(args.src), Path(args.dest), args.keep)

    if args.command == "logsummary":
        return cmd_logsummary(Path(args.file))

    if args.command == "run":
        # Expect usage: run -- <command...>
        cmd = args.cmd
        if cmd and cmd[0] == "--":
            cmd = cmd[1:]
        if not cmd:
            log.error("No command provided. Example: run -- ls -la")
            return 2
        return cmd_run(cmd)

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
