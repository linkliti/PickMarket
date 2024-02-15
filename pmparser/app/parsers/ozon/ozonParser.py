""" Ozon Parser Module """

import logging
import json
from app.parsers.baseParser import Parser

log = logging.getLogger(__name__)


class OzonParser(Parser):
  """ Ozon Parser """

  def __init__(self) -> None:
    self.host = "www.ozon.ru"
    self.api = "/api/entrypoint-api.bx/page/json/v2?url="
    super().__init__()

  def getEmbededJson(self, j: dict, keyName: str) -> dict:
    """Get embeded JSON in key with keyName in name"""
    jItem: dict = next((j[key] for key in j if keyName in key), None)
    if not jItem:
      log.error("Embed JSON: %s not found: %s", keyName, j)
      raise Exception(f"Embed JSON: {keyName} not found")
    jString = self.jsonQuotesEscape(s=jItem)
    log.debug("Embed JSON: %s",jString)
    j = json.loads(jString)
    return j