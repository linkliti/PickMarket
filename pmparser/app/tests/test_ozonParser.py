""" Test BaseParser Module """
# pylint: disable = unused-import, invalid-name, import-error, redefined-outer-name, unused-argument
from app.parsers.ozon.ozonParserCategories import OzonParserCategories
from app.tests.test_base import logger

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


# def test_getItems(logger: None) -> None:
#   """Test getItems"""
#   p = OzonParserItems()
#   p.getItems(pageUrl='/category/zhenskaya-odezhda-7501', query='свитер')
#   # for i in p.getItems(url='/category/zhenskaya-odezhda-7501', query='свитер'):
#   #   print(i)
