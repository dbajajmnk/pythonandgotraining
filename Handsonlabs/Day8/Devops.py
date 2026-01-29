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

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s %(levelname)s %(message)s",
)
log = logging.getLogger("devops_toolkit")

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
def build_parser() -> argparse.ArgumentParser:
    p = argparse.ArgumentParser(prog="devops_toolkit", description="DevOps automation toolkit demo")
    sub = p.add_subparsers(dest="command", required=True)


    # backup
    p_b = sub.add_parser("backup", help="Zip a folder and keep last N backups")
    p_b.add_argument("--src", required=True)
    p_b.add_argument("--dest", default="backups")
    p_b.add_argument("--keep", type=int, default=5)
    

    

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
