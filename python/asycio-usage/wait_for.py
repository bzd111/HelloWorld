import asyncio


async def foo(n):
    await asyncio.sleep(10)
    print(f'n: {n}!')


async def main():
    try:
        await asyncio.wait_for(foo(1), timeout=5)
    except asyncio.TimeoutError as err:
        print(err)  # no information
        print('timeout!')


if __name__ == '__main__':
    asyncio.run(main())
