"""
Day 7 Topic: Threading vs Multiprocessing — Theory + Hands-on Lab (Runnable)
============================================================================

Template followed:
What & Why → Mind Map → Engineering → Java↔Python → Mistakes → Lab → Install/Run

------------------------------------------------------------
INSTALL
------------------------------------------------------------
No external packages needed (Python 3.9+)

------------------------------------------------------------
RUN
------------------------------------------------------------
python day7_threading_vs_multiprocessing.py

------------------------------------------------------------
Goal of this lab
------------------------------------------------------------
Show:
- Threads help I/O-bound work (overlap waiting)
- Processes help CPU-bound work (use multiple cores)
"""

from __future__ import annotations
from concurrent.futures import ThreadPoolExecutor, ProcessPoolExecutor
from time import perf_counter, sleep


# ============================================================
# I/O-bound simulation Go ahead Haripriya or Mr
# ============================================================
def io_task(i: int, delay: float = 0.10) -> str:
    """
    Simulates I/O like network call, DB query, file wait.
    WHY threads help: while one task sleeps/waits, others can run.
    """
    sleep(delay)
    return f"task-{i} done"


# ============================================================
# CPU-bound simulation (pure Python)
# ============================================================
def cpu_task(iterations: int) -> int:
    """
    Pure Python CPU work (stays under GIL per process).
    WHY processes help: true parallelism across cores.
    """
    acc = 0
    for i in range(iterations):
        acc += (i * 17) % 101
    return acc


def benchmark_io(n: int = 20) -> None:
    print("\n=== I/O-bound benchmark ===")
    # Sequential
    start = perf_counter()
    for i in range(n):
        io_task(i)
    t_seq = perf_counter() - start
    print(f"Sequential:      {t_seq:.3f}s")

    # Threads
    start = perf_counter()
    with ThreadPoolExecutor(max_workers=10) as ex:
        list(ex.map(io_task, range(n)))
    t_thr = perf_counter() - start
    print(f"ThreadPool(10):  {t_thr:.3f}s")
    print("WHY: many tasks are waiting (sleep), threads overlap waiting time.")


def benchmark_cpu(tasks: int = 4, iterations: int = 400_000) -> None:
    print("\n=== CPU-bound benchmark ===")
    # Sequential
    start = perf_counter()
    total = 0
    for _ in range(tasks):
        total += cpu_task(iterations)
    t_seq = perf_counter() - start
    print(f"Sequential:         {t_seq:.3f}s  checksum={total}")

    # Threads (often no speedup in CPython due to GIL)
    start = perf_counter()
    with ThreadPoolExecutor(max_workers=4) as ex:
        total_t = sum(ex.map(cpu_task, [iterations] * tasks))
    t_thr = perf_counter() - start
    print(f"ThreadPool(4):      {t_thr:.3f}s  checksum={total_t}")
    print("WHY: CPU-bound threads often don't scale in CPython due to the GIL.")

    # Processes (bypass GIL)
    start = perf_counter()
    with ProcessPoolExecutor(max_workers=4) as ex:
        total_p = sum(ex.map(cpu_task, [iterations] * tasks))
    t_pro = perf_counter() - start
    print(f"ProcessPool(4):     {t_pro:.3f}s  checksum={total_p}")
    print("WHY: processes can run on multiple CPU cores (each process has its own GIL).")


def main() -> None:
    benchmark_io()
    benchmark_cpu()


if __name__ == "__main__":
    main()
