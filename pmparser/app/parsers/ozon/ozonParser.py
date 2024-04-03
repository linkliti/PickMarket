""" Ozon Parser Module """
import logging
from typing import override

from app.parsers.baseParser import Parser
from app.selen.seleniumPool import getBrowserToGetOzonJson
from app.utilities.jsonUtil import toJson

log = logging.getLogger(__name__)


class OzonParser(Parser):
  """ Ozon Parser """

  def __init__(self) -> None:
    self.host = "www.ozon.ru"
    self.api = "/api/entrypoint-api.bx/page/json/v2?url="
    super().__init__()

  @override
  def getDataViaSelenium(self, url: str) -> str:
    data: str = getBrowserToGetOzonJson(url=url)
    return data

  def getEmbededJson(self, j: dict, keyName: str) -> dict:
    """Get embeded JSON in key with keyName in name"""
    matchingKeys: list = [key for key in j.keys() if keyName in key]
    if not matchingKeys:
      log.error("Embed JSON: key not found", extra={"keyName": keyName, "json": j})
      raise Exception(f"Embed JSON: {keyName} not found")
    longestKey = max(matchingKeys, key=lambda k: len(j[k]))
    jString: str = j[longestKey]
    # log.debug("Embed JSON: %s", jString)
    jItem: dict = toJson(jString)
    return jItem
