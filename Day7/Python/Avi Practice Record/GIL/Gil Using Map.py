import time
from concurrent.futures import ThreadPoolExecutor

def cpu_work(n):
    count=0
    for i in range(n):
        count +=1
    return count

def run_task():
    iterations=10_000_000
    start_time=time.perf_counter()

    with ThreadPoolExecutor(max_workers=4) as executor:
        list(executor.map(cpu_work,[iterations]*4))
        # try and use submit() method here in place of map()

    print(f"Time taken: {time.perf_counter()-start_time:.2f} seconds")

if __name__=="__main__":
    run_task()


