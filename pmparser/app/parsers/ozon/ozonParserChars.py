""" Ozon Parser Module for item parameters """
import logging
from typing import Generator

from app.parsers.ozon.ozonParser import OzonParser
from app.protos import items_pb2 as itemsPB
from app.protos import types_pb2 as typesPB
from app.utilities.jsonUtil import msgToStr, toJson

log = logging.getLogger(__name__)


class OzonParserChars(OzonParser):
  """ Ozon Parser Module for item parameters """

  def getItemChars(self, itemUrl: str) -> Generator[itemsPB.Characteristic, None, None]:
    """Get item params from Ozon"""
    itemUrl = itemUrl + "features"
    log.debug('Getting params', extra={"url": itemUrl})
    jString: str = self.getData(host=self.host, url=self.api + itemUrl)

    log.info('Converting data to JSON', extra={"url": itemUrl})
    j: dict = toJson(jString)
    j = self.getEmbededJson(j=j["widgetStates"], keyName="webCharacteristics")

    log.info('Parsing chars', extra={"url": itemUrl})
    for char in j["characteristics"]:
      for charData in ["short", "long"]:
        if charData in char:
          for charObj in char[charData]:
            name: str = charObj["name"]
            key: str = str(charObj["key"]).lower()
            value: str | int | float | list[str] | None = None
            # List
            if len(charObj["values"]) > 1:
              value = [v["text"] for v in charObj["values"]]
            else:
              # String
              value = charObj["values"][0]["text"]
              # Try to Num
              value = tryConvertStrToNum(v=str(value))
            # Determine value type
            if isinstance(value, str):
              listValue = typesPB.StringList(values=[value])
              charObj = itemsPB.Characteristic(key=key, name=name, listVal=listValue)
            elif isinstance(value, float):
              charObj = itemsPB.Characteristic(key=key, name=name, numVal=value)
            elif isinstance(value, list):
              listValue = typesPB.StringList(values=value)
              charObj = itemsPB.Characteristic(key=key, name=name, listVal=listValue)
            else:
              log.error("Invalid value type", extra={"value": value, "type": type(value)})
              raise ValueError("Invalid value type " + str(type(value)))
            log.debug("char data", extra={"char": msgToStr(msg=charObj)})
            yield charObj


def tryConvertStrToNum(v: str) -> float | str:
  """Try to convert str to number"""
  try:
    return float(v)
  except ValueError:
    return v
