""" Ozon Parser Module for categories """
import json
import logging
from typing import Generator

from app.parsers.baseDataclass import BaseCategoryDataclass
from app.parsers.ozon.ozonParser import OzonParser

log = logging.getLogger(__name__)


class OzonParserCategories(OzonParser):
  """ Ozon Parser Module for categories """

  def getRootCategories(self) -> Generator[BaseCategoryDataclass, None, None]:
    """ Return name, link and empty url of parent for root categories """
    log.info('Getting categories: root')
    jString: str = self.getData(host=self.host,
                                url=self.api + '/modal/categoryMenuRoot',
                                useMobile=True)

    log.info('Converting to JSON: root')
    j: dict = json.loads(jString)
    j = self.getEmbededJson(j=j["widgetStates"], keyName="categoryMenuRoot")

    log.info('Filtering JSON: root')
    for category in j["categories"]:
      title = category["name"]
      url = category["link"]
      data = BaseCategoryDataclass(title=title, url=url, parent="")
      yield data

  def getSubCategories(self, categoryUrl: str) -> Generator[BaseCategoryDataclass, None, None]:
    """ Return name, link and url of parent of subcategory """
    log.info('Getting categories: %s', categoryUrl)
    jString = self.getData(host=self.host,
                           url=self.api + '/modal/categoryMenu' + categoryUrl,
                           useMobile=True)

    log.info('Converting to JSON: %s', categoryUrl)
    j: dict = json.loads(jString)
    j = self.getEmbededJson(j=j["widgetStates"], keyName="categoryMenu")

    log.info('Parsing JSON: %s', categoryUrl)
    while len(j["categories"]) == 1:
      j = j["categories"][0]

    log.info('Filtering JSON: %s', categoryUrl)
    for category in j["categories"]:
      title: str = category["name"]
      url: str = category["link"]
      parent: str = categoryUrl
      data = BaseCategoryDataclass(title=title, url=url, parent=parent)
      yield data
