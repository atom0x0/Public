#! -*- coding=utf-8 -*-
import asyncio
import websockets
from conf import websocket_conf
from template.logging_template import Logger

log = Logger(name=__name__).logger


class WsServer(object):
    def __init__(self):
        self.port = websocket_conf.port
        self.host = websocket_conf.host

    async def server(self, websocket, path):
        log.info(f"Allow connect: {websocket}")
        while True:
            msg = await websocket.recv()
            log.info(f"> From client {msg}")
            await websocket.send(msg)

    def server_run(self):
        log.info(f"Websocket service starting ...... OK")
        log.info(f"\nAllow host: {self.host} \nAllow port: {self.port}")
        asyncio.get_event_loop().run_until_complete(websockets.serve(self.server, str(self.host), int(self.port)))
        asyncio.get_event_loop().run_forever()
