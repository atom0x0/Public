#! -*- coding=utf-8 -*-
import asyncio
import websockets
import json
from websocket import create_connection

class WsClient(object):
    def __init__(self):
        self.port = 7001
        self.host = "localhost"

    async def client_run(self):
        async with websockets.connect(f"ws://{self.host}:{self.port}") as ws:
            while True:
                msg = await asyncio.get_event_loop().run_in_executor(None, lambda: input("Input: "))
                print(f"> Send {msg}")
                await ws.send(msg)
                recv = await ws.recv()
                print(f"< From server {recv}")

    async def client_run_ex(self):
        ws = create_connection(f"ws://{self.host}:{self.port}")
        while True:
            msg = await asyncio.get_event_loop().run_in_executor(None, lambda: input("Input: "))
            ws.send(msg)
            print(f"> Send {msg}")
            recv = ws.recv()
            print(f"< From server {recv}")
        

asyncio.get_event_loop().run_until_complete(WsClient().client_run())
asyncio.get_event_loop().run_forever()
