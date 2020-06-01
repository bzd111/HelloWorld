import asyncio


async def foo():
    await asyncio.sleep(10)
    print('Foo!')


async def hello_world():
    task = asyncio.create_task(foo())
    print(task)
    await asyncio.sleep(5)
    print('Hello World!')
    await asyncio.sleep(10)
    print(task)


if __name__ == '__main__':
    asyncio.run(hello_world())
