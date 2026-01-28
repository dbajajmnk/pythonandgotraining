"""
Day 7 Topic: GIL Explained (important for Java devs) — Theory + Lab (Runnable)
============================================================================

Template followed:
		hy → Mind Map → Engineering → Java↔Python → Mistakes → Lab → Install/Run

------------------------------------------------------------
INSTALL
------------------------------------------------------------
No external packages needed (Python 3.9+)

------------------------------------------------------------
RUN
------------------------------------------------------------
python day7_gil_explained.py

------------------------------------------------------------
NOTE (IMPORTANT)
------------------------------------------------------------
- The GIL is a CPython implementation detail.
- Most people mean CPython when they say "Python".
- Timing results depend on your machine, OS, CPU cores, and Python build.
"""
from concurrent.futures import ThreadPoolExecutor, ProcessPoolExecutor
from time import perf_counter


# ============================================================
# WHAT / WHY (quick recap)
# ============================================================
# WHAT:
# - GIL = lock that allows only one thread to run Python bytecode at a time (in CPython).
#
# WHY:
# - CPython uses reference counting for memory management.
# - The GIL simplifies thread-safety for refcount updates and many C extensions.
#
# WHEN it matters:
# - CPU-bound pure-Python code using threads: threads interleave, but don't run in parallel.
#
# WHAT to do instead:
# - For CPU-bound parallelism: multiprocessing / ProcessPoolExecutor
# - For I/O-bound concurrency: threads or async


# ============================================================
# CPU-bound work (pure Python)
# ============================================================
def cpu_work(iterations: int) -> int:
    """
    CPU-heavy loop in pure Python.
    WHY: This stays under the GIL because it's Python bytecode.
    """
    acc = 0
    for i in range(iterations):
        acc += (i * i) % 97
    return acc
    


def run_sequential(tasks: int, iterations: int) -> tuple[float, int]:
    start = perf_counter()
    total = 0
    for _ in range(tasks):
        total += cpu_work(iterations)
    return perf_counter() - start, total


def run_threads(tasks: int, iterations: int, workers: int) -> tuple[float, int]:
    start = perf_counter()
    total = 0
    with ThreadPoolExecutor(max_workers=workers) as ex:
        futures = [ex.submit(cpu_work, iterations) for _ in range(tasks)]
        for f in futures:
            total += f.result()
    return perf_counter() - start, total


def run_processes(tasks: int, iterations: int, workers: int) -> tuple[float, int]:
    start = perf_counter()
    total = 0
    with ProcessPoolExecutor(max_workers=workers) as ex:
        futures = [ex.submit(cpu_work, iterations) for _ in range(tasks)]
        for f in futures:
            total += f.result()
    return perf_counter() - start, total


def main() -> None:
    # Tune these numbers for your laptop if needed:
    tasks = 4
    iterations = 2_000_00  # 200k (keeps demo quick; increase for clearer differences)
    workers = 4

    print("\n=== CPU-bound benchmark (pure Python) ===")
    t_seq, out1 = run_sequential(tasks, iterations)
    print(f"Sequential:      {t_seq:.3f}s  (checksum={out1})")

    t_thr, out2 = run_threads(tasks, iterations, workers)
    print(f"ThreadPool({workers}): {t_thr:.3f}s  (checksum={out2})")
    print("WHY: CPU-bound threads often don't scale in CPython due to GIL.")

    # Multiprocessing needs the __main__ guard (this file has it).
    t_pro, out3 = run_processes(tasks, iterations, workers)
    print(f"ProcessPool({workers}): {t_pro:.3f}s  (checksum={out3})")
    print("WHY: separate processes bypass the GIL; can use multiple CPU cores.")

    print("\nInterpretation tips:")
    print("- If your environment has limited cores, process speedup may be small.")
    print("- If iterations is too small, overhead dominates; increase iterations for clarity.")


if __name__ == "__main__":
    main()
