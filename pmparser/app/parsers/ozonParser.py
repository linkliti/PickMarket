""" Ozon Parser Module """
import json
import logging
from app.parsers.baseParser import Parser

log = logging.getLogger(__name__)


def filterCategoriesJSON(j):
  """ Filter keys from JSON """
  if not isinstance(j, dict):
    return j
  nj = {}
  for key, value in j.items():
    if key in ["title", "url", "categories"]:
      nj[key] = value
    if key == "categories":
      if isinstance(value, list):
        nj[key] = [filterCategoriesJSON(element) for element in value]
      else:
        nj[key] = filterCategoriesJSON(value)
  return nj


class OzonParser(Parser):
  """ Ozon Parser """

  def __init__(self) -> None:
    self.host = "www.ozon.ru"

  def getCategories(self):
    """ Get categories and subcategories """
    log.info('Getting categories...')
    html = self.getData(host=self.host, url="/categories/")
    log.info('Received HTML')
    categoryIDs = [
      str(x['href']).rsplit('-', 1)[1].rstrip('/')
      for x in self.parseData(html=html, selector='.container a[href*="/category"]')
    ]

    categories: dict[str, list] = {"categories": []}

    for catID in categoryIDs:
      log.debug(catID)
      # Get Subcategories
      resp = self.getData(
        host=self.host,
        url='/api/composer-api.bx/_action/v2/categoryChildV3?menuId=185&categoryId=' + catID)
      j: dict = json.loads(resp)
      # Escape "data"
      j = j["data"]
      # Fix "columns" to "categories"
      j['categories'] = j.pop('columns')
      # Filter out keys
      j = filterCategoriesJSON(j)
      categories["categories"].append(j)

    return categories
