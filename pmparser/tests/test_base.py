""" Base Test File """
# pylint: disable = unused-import, invalid-name, import-error, redefined-outer-name, unused-argument line-too-long
import pytest
from app.utilities.log import setupLogger


@pytest.fixture
def logger() -> None:
  """Logger"""
  # Clean log file
  with open(file='parser.log', mode='w', encoding="utf-8") as f:
    f.close()
  setupLogger(name='root', debug=True)
