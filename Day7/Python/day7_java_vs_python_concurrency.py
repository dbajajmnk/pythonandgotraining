"""
Day 7 Topic: Comparison with Java Concurrency — Python equivalents + Lab (Runnable)
================================================================================

Template followed:
What & Why → Mind Map → Engineering → Java↔Python → Mistakes → Lab → Install/Run

------------------------------------------------------------
INSTALL
------------------------------------------------------------
No external packages needed (Python 3.9+)

------------------------------------------------------------
RUN
------------------------------------------------------------
python day7_java_vs_python_concurrency.py

------------------------------------------------------------
Mental model for Java devs
------------------------------------------------------------
- Java threads: parallel by default on multi-core (no GIL).
- Python (CPython) threads: great for I/O, not great for CPU-bound scaling (GIL).
- For CPU parallelism in Python: use processes (ProcessPoolExecutor).
- For massive I/O concurrency: use asyncio (event loop).
"""

from __future__ import annotations
import asyncio
from concurrent.futures import ThreadPoolExecutor, ProcessPoolExecutor
from time import perf_counter, sleep


# ============================================================
# Workloads
# ============================================================
def io_bound(i: int, delay: float = 0.10) -> str:
    """
    Simulates blocking I/O (like requests.get without async).
    In Java: you'd still use thread pools for blocking I/O, or async NIO style.
    """
    sleep(delay)
    return f"io-{i}"


async def io_bound_async(i: int, delay: float = 0.10) -> str:
    """
    Simulates non-blocking I/O (async driver).
    Python: asyncio scales well for many concurrent waits.
    """
    await asyncio.sleep(delay)
    return f"io-async-{i}"


def cpu_bound(iterations: int) -> int:
    """
    Pure Python CPU work.
    In CPython, threads won't scale well; processes can.
    """
    acc = 0
    for i in range(iterations):
        acc += (i * i) % 97
    return acc


# ============================================================
# Java-like "Executor strategy" in Python
# ============================================================
def run_io_with_threadpool(n: int = 30, workers: int = 10) -> float:
    start = perf_counter()
    with ThreadPoolExecutor(max_workers=workers) as ex:
        list(ex.map(io_bound, range(n)))
    return perf_counter() - start


async def run_io_with_asyncio(n: int = 30) -> float:
    start = perf_counter()
    await asyncio.gather(*(io_bound_async(i) for i in range(n)))
    return perf_counter() - start


def run_cpu_with_processpool(tasks: int = 6, iterations: int = 300_000, workers: int = 4) -> float:
    start = perf_counter()
    with ProcessPoolExecutor(max_workers=workers) as ex:
        list(ex.map(cpu_bound, [iterations] * tasks))
    return perf_counter() - start


def main() -> None:
    print("\n=== Java ↔ Python concurrency mapping lab ===")

    # I/O (blocking) -> ThreadPoolExecutor (similar to Java ExecutorService)
    t_threads = run_io_with_threadpool()
    print(f"I/O with ThreadPoolExecutor: {t_threads:.3f}s")

    # I/O (non-blocking) -> asyncio
    t_async = asyncio.run(run_io_with_asyncio())
    print(f"I/O with asyncio.gather:     {t_async:.3f}s")

    # CPU -> ProcessPoolExecutor
    t_cpu = run_cpu_with_processpool()
    print(f"CPU with ProcessPoolExecutor:{t_cpu:.3f}s")

    print("\nInterpretation:")
    print("- ThreadPool is great for BLOCKING I/O (sleep/network/DB waits).")
    print("- asyncio is great for NON-BLOCKING I/O (async drivers, many sockets).")
    print("- ProcessPool is best for CPU-bound pure Python work (bypasses GIL).")


if __name__ == "__main__":
    main()
