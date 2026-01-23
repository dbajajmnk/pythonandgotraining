"""
Day 6 Topic: Memory Management & GC (Python) — In-Depth Theory + Hands-on Lab
============================================================================

Template followed:
What & Why → Mind Map → Engineering → Java↔Python → Mistakes → Lab

------------------------------------------------------------
INSTALL
------------------------------------------------------------
No external packages needed (Python 3.9+)

------------------------------------------------------------
RUN
------------------------------------------------------------
python day6_memory_management_gc.py
"""

from __future__ import annotations
import gc
import weakref

# ============================================================
# WHAT & WHY
# ============================================================
# WHAT:
# - Python uses reference counting + cyclic GC (generational)
#
# WHY:
# - Reference counting frees most objects immediately when refs hit zero
# - Cyclic GC handles circular references (A -> B -> A)
#
# WHEN IT MATTERS:
# - Long-running apps/services
# - Caches and big in-memory data
# - Accidental global references

# ============================================================
# HANDS-ON LAB A: Bounded cache (prevent unbounded growth)
# ============================================================
# Scenario:
# - Report service caches monthly reports.
# - Without eviction, memory grows forever.
#
# Requirement:
# - Implement a max-size cache (simple LRU-like eviction).

class BoundedCache:
    def __init__(self, max_items: int):
        self.max_items = max_items
        self._data: dict[str, str] = {}
        self._order: list[str] = []  # tracks insertion order (simple eviction)

    def set(self, key: str, value: str) -> None:
        if key in self._data:
            # refresh order
            self._order.remove(key)
        self._data[key] = value
        self._order.append(key)

        # Evict oldest if exceeding size
        if len(self._order) > self.max_items:
            oldest = self._order.pop(0)
            del self._data[oldest]

    def get(self, key: str) -> str | None:
        return self._data.get(key)

    def __len__(self) -> int:
        return len(self._data)

# ============================================================
# HANDS-ON LAB B: Circular references + GC
# ============================================================

class Node:
    def __init__(self, name: str):
        self.name = name
        self.other: "Node | None" = None

def make_cycle():
    a = Node("A")
    b = Node("B")
    a.other = b
    b.other = a
    return a, b  # creates a cycle

def main() -> None:
    cache = BoundedCache(max_items=3)
    for i in range(1, 8):
        cache.set(f"month-{i}", f"report-data-{i}")
        print(f"After insert month-{i}: cache_size={len(cache)} keys={list(cache._data.keys())}")

    # Cycle demo
    a, b = make_cycle()
    # Drop strong references (cycle remains internally)
    a_ref = weakref.ref(a)
    b_ref = weakref.ref(b)
    del a, b

    # Force GC collection (normally automatic)
    gc.collect()
    print("Cycle collected? A alive:", a_ref() is not None, "B alive:", b_ref() is not None)

if __name__ == "__main__":
    main()
