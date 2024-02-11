""" Ozon Parser Module """
import json
import logging
from typing import Generator
from app.parsers.baseParser import Parser

log = logging.getLogger(__name__)


class OzonParser(Parser):
  """ Ozon Parser """

  def __init__(self) -> None:
    self.host = "www.ozon.ru"
    super().__init__()

  def __getEmbededJSON__(self, j: dict) -> dict:
    """Get JSON inside j["widgetStates"]["categoryMenu*"]"""
    j = next((j["widgetStates"][key] for key in j["widgetStates"] if "categoryMenu" in key), None)
    jString = self.jsonQuotesEscape(s=j)
    log.debug(jString)
    j = json.loads(jString)
    return j

  def getRootCategories(self) -> Generator[dict[str, str], None, None]:
    """ Return name, link and empty url of parent for root categories """
    log.info('Getting categories: root')
    jString: str = self.getData(
      host=self.host,
      url='/api/entrypoint-api.bx/page/json/v2?url=/modal/categoryMenuRoot',
      useMobile=True)

    log.info('Converting to JSON: root')
    j: dict = json.loads(jString)
    j = self.__getEmbededJSON__(j)

    log.info('Filtering JSON: Root')
    for category in j["categories"]:
      data: dict[str, str] = {"title": category["name"], "url": category["link"], "parent": ""}
      yield data

  def getSubCategories(self, categoryUrl: str) -> Generator[dict[str, str], None, None]:
    """ Return name, link and url of parent of subcategory """
    log.info('Getting categories: %s', categoryUrl)
    jString = self.getData(host=self.host,
                           url='/api/entrypoint-api.bx/page/json/v2?url=/modal/categoryMenu' +
                           categoryUrl,
                           useMobile=True)

    log.info('Converting to JSON: %s', categoryUrl)
    j: dict = json.loads(jString)
    j = self.__getEmbededJSON__(j)

    log.info('Parsing JSON: %s', categoryUrl)
    while len(j["categories"]) == 1:
      j = j["categories"][0]

    log.info('Filtering JSON: Root')
    for category in j["categories"]:
      data: dict[str, str] = {
        "title": category["name"],
        "url": category["link"],
        "parent": categoryUrl
      }
      yield data
