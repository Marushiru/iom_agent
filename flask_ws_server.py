import asyncio
import websockets


async def echo(websocket, path):
    async for message in websocket:
        print('you have a new message :{}'.format(message))
        await websocket.send(message)


async def main():
    async with websockets.serve(echo, "localhost", 8888):
        await asyncio.Future()


asyncio.run(main())