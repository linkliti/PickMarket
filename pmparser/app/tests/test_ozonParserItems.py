""" Test BaseParser Module """
# pylint: disable = unused-import, invalid-name, import-error, redefined-outer-name, unused-argument
from app.tests.test_base import logger
from app.parsers.ozon.ozonParserItems import OzonParserItems


def test_getItems(logger: None) -> None:
  """Test getItems"""
  p = OzonParserItems()
  for i in p.getItems(
      pageUrl=
      '/category/ingalyatory-i-aksessuary-35039/?country=2&currency_price=118.000;1500.000&text=небулайзер&type=33239' #pylint: disable=line-too-long
  ):
    print(i)


def test_getItemsFewPages(logger: None) -> None:
  """Test getItems with few pages"""
  p = OzonParserItems()
  for i in p.getItems(pageUrl='/category/energeticheskie-napitki-9473', numOfPages=2):
    print(i)
