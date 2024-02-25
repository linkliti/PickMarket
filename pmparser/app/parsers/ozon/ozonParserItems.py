""" Ozon Parser Module for categories """
import html
import json
import logging
import re
from typing import Generator

from app.parsers.baseDataclass import BaseItemDataclass
from app.parsers.ozon.ozonParser import OzonParser

log = logging.getLogger(__name__)


class OzonParserItems(OzonParser):
  """ Ozon Parser Module for items """

  def getItems(self,
               pageUrl: str,
               query: str | None = None,
               numOfPages: int | None = None) -> Generator[BaseItemDataclass, None, None]:
    """ Get items from Ozon """
    log.debug('Getting items: %s', pageUrl)
    reqParams: dict[str, str] = {}
    if query:
      reqParams = {"text": query}
    jString: str = self.getData(host=self.host, url=self.api + pageUrl, params=reqParams)

    log.info('Converting data to JSON: %s', pageUrl)
    j: dict = json.loads(jString)

    log.info('Getting page info for items: %s', pageUrl)
    jPage = self.getEmbededJson(j=j, keyName="shared")
    totalItems: int = jPage["catalog"]["totalFound"]
    totalPages = jPage["catalog"]["totalPages"]
    if numOfPages:
      totalPages: int = min(totalPages, numOfPages)

    log.info('Item and page counts for %s: items %d, pages %d', pageUrl, totalItems, totalPages)
    for i in self.getItemsFromPage(pageUrl=pageUrl, page=1, reqParams=reqParams, jString=jString):
      yield i
    for page in range(1, totalPages + 1):
      for i in self.getItemsFromPage(pageUrl=pageUrl, page=page, reqParams=reqParams):
        yield i

  def getItemsFromPage(self,
                       pageUrl: str,
                       page: int,
                       reqParams: dict[str, str],
                       jString: str | None = None) -> Generator[BaseItemDataclass, None, None]:
    """ Get items from page """
    log.info('Getting items from page %d: %s', page, pageUrl)
    if jString is None:
      jString = self.getData(host=self.host,
                             url=self.api + pageUrl + "&page=" + str(page),
                             params=reqParams)

    log.info('Converting data to JSON: %s', pageUrl)
    j: dict = json.loads(jString)
    jItems: dict = self.getEmbededJson(j=j["widgetStates"], keyName="searchResultsV2")
    log.info('Parsing items: %s', pageUrl)
    for item in jItems["items"]:
      try:
        data: BaseItemDataclass = self.getItem(itemJson=item)
      except KeyError as e:
        log.error("Failed to get item with error: %s %s: %s", type(e), e, item)
        continue
      log.debug("Item: %s", data)
      yield data

  def getItem(self, itemJson: dict) -> BaseItemDataclass:
    """ Create an instance of BaseItemDataClass from JSON data """
    stars: float | None = None
    comments: int | None = None
    oldPrice: int | None = None
    name: str = ""
    price: int = 0
    # MainState Varibles
    for item in itemJson["mainState"]:
      atom: dict = item["atom"]
      if "textAtom" in atom:
        name = atom["textAtom"]["text"]
        name = html.unescape(name)
      elif "priceV2" in atom:
        price = int(re.sub(r"\D", "", atom["priceV2"]["price"][0]["text"]))
        if len(atom["priceV2"]["price"]) == 2:
          oldPrice = int(re.sub(r"\D", "", atom["priceV2"]["price"][1]["text"]))
      elif "labelList" in atom:
        for item in atom["labelList"]["items"]:
          if "icon" in item:
            if item["icon"]["image"] == "ic_s_star_filled_compact":
              stars = float(re.sub("<.*?>", "", item["title"]))
            elif item["icon"]["image"] == "ic_s_dialog_filled_compact":
              comments = int(re.sub(r"\D", "", item["title"]))
    # Other Variables
    url: str = itemJson["action"]["link"].split("?")[0]
    try:
      imageUrl: str = itemJson["tileImage"]["items"][0]["image"]["link"]
    except KeyError:
      imageUrl: str = itemJson["tileImage"]["items"][0]["video"]["preview"]
    isAdult: bool = itemJson["isAdult"]
    data = BaseItemDataclass(name=name,
                             url=url,
                             imageUrl=imageUrl,
                             price=price,
                             isAdult=isAdult,
                             stars=stars,
                             comments=comments,
                             oldPrice=oldPrice)
    return data
