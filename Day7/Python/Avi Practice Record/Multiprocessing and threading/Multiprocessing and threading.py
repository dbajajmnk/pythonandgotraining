

import threading
import multiprocessing
import time



# I/O TASK 
def io_task(name):
    print(f"{name} started")
    time.sleep(3)  
    print(f"{name} finished")



# CPU TASK 
def cpu_task(n):
    result = 0
    for i in range(10_000_000):
        result += i*i % 5
    print(f"CPU task {n} done")



# THREAD DEMO

def threading_demo():
    print("\nTHREADING DEMO (I/O tasks)")

    start = time.time()

    t1 = threading.Thread(target=io_task, args=("Task1",))
    t2 = threading.Thread(target=io_task, args=("Task2",))

    t1.start()
    t2.start()

    t1.join()
    t2.join()

    print("Time taken:", time.time() - start)



# MULTIPROCESSING DEMO

def multiprocessing_demo():
    print("\nMULTIPROCESSING DEMO (CPU tasks)")

    start = time.time()

    p1 = multiprocessing.Process(target=cpu_task, args=(1,))
    p2 = multiprocessing.Process(target=cpu_task, args=(2,))

    p1.start()
    p2.start()

    p1.join()
    p2.join()

    print("Time taken:", time.time() - start)



# MAIN

if __name__ == "__main__":
    threading_demo()
    multiprocessing_demo()