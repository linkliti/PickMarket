""" Ozon Parser Module for categories """
from typing import Generator
import logging


from app.parsers.ozon.ozonParser import OzonParser
from app.parsers.baseItem import BaseItemDataClass

log = logging.getLogger(__name__)


class OzonParserItems(OzonParser):
  """ Ozon Parser Module for items """

  # def getItems(self, pageUrl: str, query: str = None) -> Generator[BaseItemDataClass, None, None]:
  #   """ Get items from Ozon """
  #   log.debug('Getting items: %s', pageUrl)
  #   reqParams: dict[str, str] = {"text": query}
  #   jString: str = self.getData(host=self.host, url=self.api + pageUrl, params=reqParams)

  #   log.debug('Parsing items JSON: %s', pageUrl)
  #   log.debug(jString)
    

  # def getItems(self, url: str, query: str = None) -> Generator[BaseItemDataClass, None, None]:
  #   """ Get items from Ozon """
  #   log.debug('Getting items: %s', url)
  #   params: dict[str, str] = {"text": query}
  #   html = self.getData(host=self.host, url=url, params=params)

  #   log.debug('Parsing items: %s', url)
  #   items: ResultSet[Tag] = self.htmlStringToTags(html=html,
  #                                                 selector="#paginatorContent > div > div > div")
  #   log.debug("Got %s items", len(items))
  #   log.debug(html)
  #   for item in items:
  #     t = BaseItemHelper(tag=item)
  #     item = BaseItemDataClass(
  #       name=t.getText(sel="a > div > span"),
  #       url=t.getKey(sel="a:nth-child(1)", key="href"),
  #       imageUrl=t.getKey(sel="img", key="src"),
  #       price=t.getInt(sel="div:nth-child(1) > div:nth-child(1) > span:nth-child(1)"),
  #       oldPrice=t.getInt(sel="div:nth-child(1) > div:nth-child(1) > span:nth-child(2)"),
  #       stars=t.getFloat(sel="div:nth-child(4) span:nth-child(1) > span"),
  #       comments=t.getInt(sel="span:nth-child(2) > span"))
  #     log.debug("Got item: %s", item)
  #     yield item.to_json() #pylint: disable=no-member
