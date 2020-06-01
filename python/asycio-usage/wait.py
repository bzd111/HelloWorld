import asyncio
from random import randrange


async def foo(n):
    s = randrange(5)
    print(f'{n} will sleep for: {s} seconds')
    await asyncio.sleep(s)
    print(f'n: {n}!')


async def main():
    tasks = [foo(1), foo(2), foo(3)]
    done, pending = await asyncio.wait(tasks, return_when=asyncio.FIRST_COMPLETED)
    # result = await asyncio.wait(tasks)
    print(done)
    print(pending)


if __name__ == '__main__':
    asyncio.run(main())
