import asyncio

async def task(name):
    print("starting task:", name)
    await asyncio.sleep(2)
    print("done task:", name)
    return name

async def main():
    tasks = [task(i) for i in ["A", "B", "C"]]

    result = await asyncio.gather(*tasks)

    print("result:", result)

asyncio.run(main())