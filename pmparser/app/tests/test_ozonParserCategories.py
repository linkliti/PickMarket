""" Test BaseParser Module """
# pylint: disable = unused-import, invalid-name, import-error, redefined-outer-name, unused-argument
from app.parsers.ozon.ozonParserCategories import OzonParserCategories
from app.tests.test_base import logger
from app.parsers.ozon.ozonParserFilters import OzonParserFilters


def test_getRootCategories(logger: None) -> None:
  """Test getRootCategories"""
  p = OzonParserCategories()
  for i in p.getRootCategories():
    print(i)
  assert True


def test_getSubCategories(logger: None) -> None:
  """Test getRootCategories"""
  p = OzonParserCategories()
  for i in p.getSubCategories(categoryUrl='/category/zhenskaya-odezhda-7501'):
    print(i)
  assert True


def test_getFilters(logger: None) -> None:
  """Test getFilters"""
  p = OzonParserFilters()
  url = '/category/videokarty-15721'
  for i in p.getFilters(pageUrl=url):
    print(i)
  assert True
