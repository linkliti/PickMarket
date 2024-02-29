""" Base Test File """
# pylint: disable = unused-import, invalid-name, import-error, redefined-outer-name, unused-argument line-too-long
import logging

import grpc
import pytest
from app.utilities.log import setupLogger

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
  channel: grpc.Channel = grpc.insecure_channel(target='localhost:50051')
  return channel
