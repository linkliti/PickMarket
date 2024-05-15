""" Base Test File """
# pylint: disable = unused-import, invalid-name, import-error, redefined-outer-name, unused-argument line-too-long
import logging
from typing import Generator

import pytest
from app.utilities.log import setupLogger


@pytest.fixture
def logger() -> None:
  """Logger"""
  setupLogger(name='test', debug=True, filename="logs/parserTest.log")
