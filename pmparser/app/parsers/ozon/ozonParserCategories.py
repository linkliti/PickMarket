""" Ozon Parser Module for categories """
import logging
from typing import Generator
import urllib.parse as parse
from app.parsers.ozon.ozonParser import OzonParser
from app.protos import categories_pb2 as categPB
from app.utilities.jsonUtil import toJson

log = logging.getLogger(__name__)


class OzonParserCategories(OzonParser):
  """ Ozon Parser Module for categories """

  def getRootCategories(self) -> Generator[categPB.Category, None, None]:
    """ Return name, link and empty url of parent for root categories """
    log.info('Getting categories', extra={"url": '/modal/categoryMenuRoot'})
    jString: str = self.getData(host=self.host,
                                url=self.api + '/modal/categoryMenuRoot',
                                useMobile=True)

    log.info('Converting to JSON', extra={"url": '/modal/categoryMenuRoot'})
    j: dict = toJson(jString)
    j = self.getEmbededJson(j=j["widgetStates"], keyName="categoryMenuRoot")

    log.info('Filtering JSON', extra={"url": '/modal/categoryMenuRoot'})
    for category in j["categories"]:
      title = category["name"]
      # Remove the query parameters
      tempUrl: parse.ParseResult = parse.urlparse(category["link"])
      tempUrl = tempUrl._replace(query="")
      url: str =  parse.urlunparse(components=tempUrl)
      data = categPB.Category(title=title, url=url)
      yield data

  def getSubCategories(self, categoryUrl: str) -> Generator[categPB.Category, None, None]:
    """ Return name, link and url of parent of subcategory """
    log.info('Getting categories', extra={"url": categoryUrl})
    jString = self.getData(host=self.host,
                           url=self.api + '/modal/categoryMenu' + categoryUrl,
                           useMobile=True)

    log.info('Converting to JSON', extra={"url": categoryUrl})
    j: dict = toJson(jString)
    j = self.getEmbededJson(j=j["widgetStates"], keyName="categoryMenu")

    log.info('Parsing JSON', extra={"url": categoryUrl})
    while len(j["categories"]) == 1:
      j = j["categories"][0]

    log.info('Filtering JSON', extra={"url": categoryUrl})
    for category in j["categories"]:
      title: str = category["cellTrackingInfo"]["title"]
      # Remove the query parameters
      tempUrl: parse.ParseResult = parse.urlparse(category["link"])
      tempUrl = tempUrl._replace(query="")
      url: str =  parse.urlunparse(components=tempUrl)
      parentUrl: str = categoryUrl
      data = categPB.Category(title=title, url=url, parentUrl=parentUrl)
      yield data
