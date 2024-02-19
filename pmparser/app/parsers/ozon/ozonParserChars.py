""" Ozon Parser Module for item parameters """
import json
from typing import Generator, List
import logging

from app.parsers.ozon.ozonParser import OzonParser
from app.parsers.baseDataclass import BaseItemCharsDataclass

log = logging.getLogger(__name__)


class OzonParserChars(OzonParser):
  """ Ozon Parser Module for item parameters """

  def getItemChars(self, itemUrl: str) -> Generator[dict, None, None]:
    """Get item params from Ozon"""
    itemUrl = itemUrl + "/features"
    log.debug('Getting params: %s', itemUrl)
    jString: str = self.getData(host=self.host, url=self.api + itemUrl)

    log.info('Converting data to JSON: %s', itemUrl)
    j: dict = json.loads(jString)
    j = self.getEmbededJson(j=j["widgetStates"], keyName="webCharacteristics")

    log.info('Parsing chars: %s', itemUrl)
    for char in j["characteristics"]:
      for charData in ["short", "long"]:
        if charData in char:
          for charObj in char[charData]:
            name: str = charObj["name"]
            key: str = charObj["key"]
            value: str | int | float | List[str] = None
            # List
            if len(charObj["values"]) > 1:
              value = [v["text"] for v in charObj["values"]]
            else:
              # String
              value = charObj["values"][0]["text"]
              # Try to Float/Int
              value = tryConvertStrToNum(v=value)
            charObj = BaseItemCharsDataclass(key=key, name=name, value=value)
            log.debug("Char: %s", charObj)
            yield charObj


def tryConvertStrToNum(v: str) -> int | float | str:
  """Try to convert str to number"""
  try:
    v: float = float(v)
    v: int | float = int(v) if v.is_integer() else v
  except ValueError:
    pass
  return v
