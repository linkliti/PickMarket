""" Test Ozon Parsing """
# pylint: disable = unused-import, invalid-name, import-error, redefined-outer-name, unused-argument, line-too-long
from app.parsers.ozon.ozonParserCategories import OzonParserCategories
from app.parsers.ozon.ozonParserChars import OzonParserChars
from app.parsers.ozon.ozonParserFilters import OzonParserFilters
from app.parsers.ozon.ozonParserItems import OzonParserItems

from .test_base import logger


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
  # url = '/category/videokarty-15721'
  url = '/category/bluzy-i-rubashki-zhenskie-7511'
  for i in p.getRootFilters(categoryUrl=url):
    print(i)


def test_getItems(logger: None) -> None:
  """Test getItems"""
  p = OzonParserItems()
  for i in p.getItems(
      pageUrl=
      '/category/ingalyatory-i-aksessuary-35039/?country=2&currency_price=118.000;1500.000&text=небулайзер&type=33239'
  ):
    print(i)


def test_getItemsFewPages(logger: None) -> None:
  """Test getItems with few pages"""
  p = OzonParserItems()
  for i in p.getItems(pageUrl='/category/ingalyatory-i-aksessuary-35039', numOfPages=2):
    print(i)


def test_getItemChars(logger: None) -> None:
  """Test getItemChars"""
  p = OzonParserChars()
  url = '/product/gigabyte-videokarta-geforce-rtx-3080-10-gb-gv-n3080gaming-oc-10gd-2-0-lhr-306060738'
  # url = '/product/monster-energy-nitro-super-dry-irlandiya-12-sht-h-500-ml-1181224292'
  for i in p.getItemChars(itemUrl=url):
    print(i)
