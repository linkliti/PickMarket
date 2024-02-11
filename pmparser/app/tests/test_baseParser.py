""" Test BaseParser Module """
# pylint: disable = unused-import, invalid-name, import-error, redefined-outer-name, unused-argument
from app.utilities.log import setupLogger
from app.parsers.ozonParser import OzonParser
import pytest

@pytest.fixture
def logger() -> None:
  """Logger"""
  setupLogger(name='root', debug=True)

def test_getRootCategories(logger) -> None:
  """Test getRootCategories"""
  p = OzonParser()
  for i in p.getRootCategories():
    print(i)

def test_getSubCategories(logger) -> None:
  """Test getRootCategories"""
  p = OzonParser()
  for i in p.getSubCategories('/category/zhenskaya-odezhda-7501'):
    print(i)
