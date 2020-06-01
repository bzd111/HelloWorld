import asyncio


async def foo(n):
    await asyncio.sleep(3)
    print(f'n: {n}!')


async def main():
    tasks = [foo(1), foo(2), foo(3)]
    await asyncio.gather(*tasks)


if __name__ == '__main__':
    asyncio.run(main())
