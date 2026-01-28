from time import perf_counter
from concurrent.futures import ThreadPoolExecutor, ProcessPoolExecutor

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

## Hari Priya
def run_threads(tasks: int, iterations: int, workers: int) -> tuple[float, int]:
    start = perf_counter()
    total = 0
    with ThreadPoolExecutor(max_workers=workers) as ex:
        futures = [ex.submit(cpu_work, iterations) for _ in range(tasks)]
        for f in futures:
            total += f.result()
    return perf_counter() - start, total

## Harivardan
def run_processes(tasks: int, iterations: int, workers: int) -> tuple[float, int]:
    start = perf_counter()
    total = 0
    with ProcessPoolExecutor(max_workers=workers) as ex:
        futures = [ex.submit(cpu_work, iterations) for _ in range(tasks)]
        for f in futures:
            total += f.result()
    return perf_counter() - start, total


def main()->None:
    tasks = 4 
    iterations = 20_00_00
    workrs = 4
    print("CPU Bound Bench Mark")
    t_seq , out1 = run_sequential(tasks,iterations)
    print(f"Sequencial : {t_seq:.3f}s Output : {out1}")

    t_thr,out2 = run_threads(tasks,iterations,workrs)
    print(f"ThreadPool({workrs}): {t_thr:.3f}s  (checksum={out2})")
    print("WHY: CPU-bound threads often don't scale in CPython due to GIL.")

    t_proc , out3 = run_processes(tasks,iterations,workrs)
    print(f"ProcessPool({workrs}): {t_proc:.3f}s  (checksum={out3})")
    print("WHY: CPU-bound processes scale well as they bypass the GIL.")    

if __name__ == "__main__" :
   main()

 