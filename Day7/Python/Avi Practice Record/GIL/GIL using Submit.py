import time
from concurrent.futures import ThreadPoolExecutor, as_completed

def cpu_task(n):
    """A CPU-heavy loop that requires the GIL."""
    count = 0
    for i in range(n):
        count += i
    return count

def run_example():
    iterations = 10_000_000
    start_time = time.perf_counter()

    # Create the thread pool
    with ThreadPoolExecutor(max_workers=4) as executor:
        # 1. Submit each task individually and store the returned 'Future' objects
        # submit returns immediately without waiting for the task to finish.
        futures = [executor.submit(cpu_task, iterations) for _ in range(4)]
        
        print("Tasks submitted. Waiting for results...")

        # 2. Process results as they complete using as_completed()
        # This helper function yields futures as soon as they finish, regardless of submission order.
        for future in as_completed(futures):
            # 3. Retrieve the actual result using .result()
            # This call will block ONLY for the specific task that just finished.
            result = future.result()
            print(f"Task finished with result: {result}")

    total_time = time.perf_counter() - start_time
    print(f"\nTotal time taken: {total_time:.2f} seconds")

if __name__ == "__main__":
    run_example()
