#! -*- coding=utf-8 -*-
import os, sys
# sys.path.append(os.path.join(os.path.dirname(__file__), '.', 'core'))
from core.websocket.ws_server import WsServer
from template.logging_template import Logger

if __name__ == "__main__":
    WsServer().server_run()