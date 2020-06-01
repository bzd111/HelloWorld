import asyncio
from asyncio import Future


async def foo():
    await asyncio.sleep(10)
    return 'Foo!'


def got_result(future: Future):
    print(f'got the result! {future.result()}')


async def hello_world():
    task = asyncio.create_task(foo())
    task.add_done_callback(got_result)
    print(task)
    await asyncio.sleep(5)
    print('Hello World!')
    await asyncio.sleep(10)
    print(task)


if __name__ == '__main__':
    asyncio.run(hello_world())
