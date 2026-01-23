"""
Day 7 Topic: Async Programming (async/await) — Theory + Hands-on Lab (Runnable)
=============================================================================

Template followed:
What & Why → Mind Map → Engineering → Java↔Python → Mistakes → Lab → Install/Run

------------------------------------------------------------
INSTALL
------------------------------------------------------------
No external packages needed (Python 3.9+)

------------------------------------------------------------
RUN
------------------------------------------------------------
python day7_async_await.py

------------------------------------------------------------
Key idea for Java devs
------------------------------------------------------------
- Async is concurrency (overlap waiting), not parallel CPU execution.
- Similar "shape" to Java CompletableFuture chains.
"""

from __future__ import annotations
import asyncio
from time import perf_counter


# ============================================================
# Simulated async I/O (like API call)
# ============================================================
async def fake_api_call(i: int, delay: float = 0.20) -> str:
    """
    WHAT: an async function returns a coroutine object.
    WHY await: yields control so the event loop can run other tasks.
    """
    await asyncio.sleep(delay)  # non-blocking sleep (simulates network wait)
    return f"api-result-{i}"


async def sequential_calls(n: int) -> list[str]:
    # Await one-by-one (slowest for many independent calls)
    results = []
    for i in range(n):
        results.append(await fake_api_call(i))
    return results


async def concurrent_calls(n: int) -> list[str]:
    # Launch all calls concurrently and await them together
    tasks = [fake_api_call(i) for i in range(n)]
    return await asyncio.gather(*tasks)


async def main_async() -> None:
    n = 20

    print("\n=== Async lab: sequential vs concurrent ===")
    start = perf_counter()
    r1 = await sequential_calls(n)
    t1 = perf_counter() - start
    print(f"Sequential await: {t1:.3f}s  results={len(r1)}")

    start = perf_counter()
    r2 = await concurrent_calls(n)
    t2 = perf_counter() - start
    print(f"asyncio.gather:   {t2:.3f}s  results={len(r2)}")

    print("\nWHY gather is faster here: calls are independent and mostly waiting (I/O).")
    print("COMMON MISTAKE: using time.sleep() in async code -> blocks event loop. Use asyncio.sleep().")


def main() -> None:
    asyncio.run(main_async())


if __name__ == "__main__":
    main()
