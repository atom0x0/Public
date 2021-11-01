#! -*- coding: utf-8 -*-
import os
import sys
import logging
from time import strftime

# 默认日志格式
DEFAULT_LOG_FMT = "%(asctime)s %(filename)s [line:%(lineno)d] %(levelname)s: %(message)s"
# 默认时间格式
DEFUALT_LOG_DATEFMT = "%Y-%m-%d %H:%M:%S"
# 输出日志路径
LOG_OUT_PATH = "/".join(os.path.dirname(__file__).split("/")[:-1]) + "/logs/"


class Logger(object):
    def __init__(self, name=__name__):
        self._logger = logging.getLogger(name)
        self.DEFAULT_LOG_FILENAME = "{0}{1}.log".format(LOG_OUT_PATH, strftime("%Y-%m-%d"))
        self.formatter = logging.Formatter(fmt=DEFAULT_LOG_FMT, datefmt=DEFUALT_LOG_DATEFMT)
        self._logger.addHandler(self._get_file_handler(self.DEFAULT_LOG_FILENAME))
        # self._logger.addHandler(self._get_console_handler())
        self._logger.setLevel(logging.INFO)

    def _get_file_handler(self, filename):
        filehandler = logging.FileHandler(filename, encoding="utf-8")
        filehandler.setFormatter(self.formatter)
        return filehandler

    def _get_console_handler(self):
        console_handler = logging.StreamHandler(sys.stdout)
        console_handler.setFormatter(self.formatter)
        return console_handler

    @property
    def logger(self):
        return self._logger
