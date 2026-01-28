from time import perf_counter

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


def main()->None:
    tasks = 4 
    iterations = 20_00_00
    workrs = 4
    print("CPU Bound Bench Mark")
    t_seq , out1 = run_sequential(tasks,iterations)
    print(f"Sequencial : {t_seq:.3f}s Output : {out1}")


if __name__ == "__main__" :
   main()

 