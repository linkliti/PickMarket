""" Test BaseParser Module """
# pylint: disable = unused-import, invalid-name, import-error
from app.parsers.baseParser import Parser


def testGetData() -> None:
  """Test getData"""
  p = Parser()
  html: str = p.getData(host="www.ozon.ru", url="/categories/")
  print(html)
  file = open(file="./.temp/test.html", mode='w', encoding="utf-8")
  file.write(html)
  assert html != ""
