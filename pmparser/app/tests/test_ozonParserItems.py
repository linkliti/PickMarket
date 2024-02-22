""" Test BaseParser Module """
# pylint: disable = unused-import, invalid-name, import-error, redefined-outer-name, unused-argument line-too-long
from app.tests.test_base import logger
from app.parsers.ozon.ozonParserItems import OzonParserItems
from app.parsers.ozon.ozonParserChars import OzonParserChars


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
