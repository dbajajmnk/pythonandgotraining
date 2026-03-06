import asyncio
async def task(name):
    print("starting task :", name)
    await asyncio.sleep(2)
    print("done task :", name)
    return name

async def main():
    result=await asyncio.gather(
        task("A"),
        task("B"),
        task("C")
    )
    print("Extra")
    print("result: ",result)

asyncio.run(main())
