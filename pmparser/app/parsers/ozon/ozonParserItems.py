""" Ozon Parser Module for categories """
import html
import json
import logging
import re
from typing import Generator

from app.parsers.ozon.ozonParser import OzonParser
from app.protos import items_pb2 as itemsPB
from app.utilities.jsonUtil import msgToStr, toJson

log = logging.getLogger(__name__)


class OzonParserItems(OzonParser):
  """ Ozon Parser Module for items """

  def getItems(self,
               pageUrl: str,
               query: str | None = None,
               params: str | None = None,
               numOfPages: int | None = None) -> Generator[itemsPB.Item, None, None]:
    """ Get items from Ozon """
    # Request
    log.debug('Getting items', extra={"pageUrl": pageUrl})
    reqParams: dict[str, str] = {}
    if params:
      paramsDict: dict[str, str] = json.loads(params)
      reqParams = dict(paramsDict.items())
    if query:
      reqParams["text"] = query
    jString: str = self.getData(host=self.host, url=self.api + pageUrl, params=reqParams)
    # Convert data to JSON
    log.info('Converting data to JSON', extra={"pageUrl": pageUrl})
    j: dict = toJson(jString)
    log.info('Getting page info for items', extra={"pageUrl": pageUrl})
    jPage = self.getEmbededJson(j=j, keyName="shared")
    # Get items page info
    totalItems: int = jPage["catalog"]["totalFound"]
    totalPages = jPage["catalog"]["totalPages"]
    if numOfPages:
      totalPages: int = min(totalPages, numOfPages)
    log.info("Items page info", extra=\
      {"pageUrl": pageUrl, "totalItems": totalItems, "totalPages": totalPages})
    # Get items from first downloaded page
    for i in self.getItemsFromPage(pageUrl=pageUrl, page=1, reqParams=reqParams, jString=jString):
      yield i
    # Get items from other pages
    for page in range(2, totalPages + 1):
      for i in self.getItemsFromPage(pageUrl=pageUrl, page=page, reqParams=reqParams):
        yield i

  def getItemsFromPage(self,
                       pageUrl: str,
                       page: int,
                       reqParams: dict[str, str],
                       jString: str | None = None) -> Generator[itemsPB.Item, None, None]:
    """ Get items from page """
    log.info('Getting items from page', extra={"pageUrl": pageUrl, "page": page})
    if jString is None:  # First page
      reqParams["page"] = str(page)
      jString = self.getData(host=self.host, url=self.api + pageUrl, params=reqParams)
    # Convert data to JSON
    log.info('Converting data to JSON', extra={"pageUrl": pageUrl, "page": page})
    j: dict = toJson(jString)
    jItems: dict = self.getEmbededJson(j=j["widgetStates"],
                                       keyName=["tileGrid2", "searchResultsV2"])
    # Parsing
    log.info('Parsing items', extra={"pageUrl": pageUrl, "page": page})
    for item in jItems["items"]:
      try:
        # Get item data
        data: itemsPB.Item = self.getItem(itemJson=item)
      except KeyError as e:
        log.error("Failed to get item with error", extra={"error": e})
        continue
      log.debug("Item data", extra={"data": msgToStr(msg=data)})
      yield data

  def getItem(self, itemJson: dict) -> itemsPB.Item:
    """ Create an instance of BaseItemDataClass from JSON data """
    stars: float | None = None
    comments: int | None = None
    oldPrice: int | None = None
    original: bool = False
    name: str = ""
    price: int = 0
    # MainState Varibles
    for item in itemJson["mainState"]:
      try:
        atom: dict = item["atom"]
      except KeyError:
        atom: dict = item
      if "textAtom" in atom:
        name = atom["textAtom"]["text"]
        name = html.unescape(name)
      # Price structure
      elif "priceV2" in atom:
        price = int(re.sub(r"\D", "", atom["priceV2"]["price"][0]["text"]))
        if len(atom["priceV2"]["price"]) == 2:
          oldPrice = int(re.sub(r"\D", "", atom["priceV2"]["price"][1]["text"]))
      # Rating + isOriginal structure
      elif "labelList" in atom:
        for item in atom["labelList"]["items"]:
          if "icon" in item:
            if item["icon"]["image"] == "ic_s_star_filled_compact":
              stars = float(re.sub("<.*?>", "", item["title"]))
            elif item["icon"]["image"] == "ic_s_dialog_filled_compact":
              comments = int(re.sub(r"\D", "", item["title"]))
            elif item["icon"]["image"] == "ic_s_confirmed_filled_compact":
              original = True
    # Other Variables
    url: str = str(itemJson['action']['link']).split('?', maxsplit=1)[0]
    try:
      imageUrl: str = itemJson["tileImage"]["items"][0]["image"]["link"]
    except KeyError:
      imageUrl: str = itemJson["tileImage"]["items"][0]["video"]["preview"]
    isAdult: bool = itemJson["isAdult"]
    data = itemsPB.Item(name=name,
                        url=url,
                        imageUrl=imageUrl,
                        price=price,
                        isAdult=isAdult,
                        rating=stars,
                        comments=comments,
                        oldPrice=oldPrice,
                        original=original)
    return data
