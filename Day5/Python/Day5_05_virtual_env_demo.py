"""
Day 5 Topic 5: Virtual Environments (venv) — practical guide + checks

This script doesn't create a venv (you do that in terminal),
but it helps you VERIFY you are in the correct environment.

Scenario:
- Two projects require different library versions.
- venv keeps dependencies isolated so "Project A" doesn't break "Project B".
"""

from __future__ import annotations

import os
import sys
import site
from pathlib import Path


def show_environment_info() -> None:
    print("Python executable:", sys.executable)
    print("Python version   :", sys.version.split()[0])
    print("VIRTUAL_ENV      :", os.environ.get("VIRTUAL_ENV"))
    print("sys.prefix       :", sys.prefix)
    print("base_prefix      :", getattr(sys, "base_prefix", None))
    print("site-packages    :", site.getsitepackages())


def in_venv() -> bool:
    base = getattr(sys, "base_prefix", sys.prefix)
    return sys.prefix != base


def print_steps() -> None:
    print("\n--- How to create venv (Windows) ---")
    print("python -m venv .venv")
    print(r".\.venv\Scripts\activate")
    print("python -m pip install --upgrade pip")
    print("pip install requests")
    print("pip freeze > requirements.txt")

    print("\n--- How to create venv (macOS/Linux) ---")
    print("python3 -m venv .venv")
    print("source .venv/bin/activate")
    print("python -m pip install --upgrade pip")
    print("pip install requests")
    print("pip freeze > requirements.txt")

    print("\n--- Recreate later ---")
    print("python -m venv .venv")
    print("activate venv")
    print("pip install -r requirements.txt")

    print("\nCommon mistakes:")
    print("- Installing packages globally instead of inside venv")
    print("- Forgetting to activate venv before running code")
    print("- Committing the .venv folder to git (usually avoid)")
    print("- Not pinning versions (requirements.txt)")


def main() -> None:
    show_environment_info()
    print("\nIs this a virtual environment?", in_venv())
    print_steps()

    # Hands-on mini lab:
    # 1) Create venv
    # 2) Install a package (requests)
    # 3) Confirm it is installed ONLY inside venv:
    #    python -c "import requests; print(requests.__version__)"
    # 4) Deactivate and confirm it fails outside (if not installed globally)


if __name__ == "__main__":
    main()
