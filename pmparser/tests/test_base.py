""" Base Test File """
# pylint: disable = unused-import, invalid-name, import-error, redefined-outer-name, unused-argument line-too-long
from concurrent.futures import ThreadPoolExecutor
import logging

import grpc
import pytest
from app.utilities.log import setupLogger
from app.selen.seleniumPool import browserQueue
from app.selen.seleniumMarkets import SeleniumWorkerMarkets

log = logging.getLogger(__name__)


@pytest.fixture
def logger() -> None:
  """Logger"""
  # Clean log file
  with open(file='parser.log', mode='w', encoding="utf-8") as f:
    f.close()
  setupLogger(name='root', debug=True, filename="parserTest.log")


@pytest.fixture
def channel() -> grpc.Channel:
  """Connect to server via gRPC"""
  channel: grpc.Channel = grpc.insecure_channel(target='localhost:5003')
  return channel


@pytest.fixture
def browserPool():
  """Browser pool fixture"""
  worker = SeleniumWorkerMarkets()
  browserQueue.put(item=worker)
  setCookies(worker)
  log.info("Browsers started. Worker count: %d", 1)


def setCookies(worker: SeleniumWorkerMarkets) -> None:
  """ Call setOzonAdultCookies """
  worker.setOzonAdultCookies()
