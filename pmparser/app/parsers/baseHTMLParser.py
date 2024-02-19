""" Base HTML parser """
import re
from typing import Any
from bs4 import ResultSet, Tag


class BaseItemHelper:
  """ Base HTML parser """

  def __init__(self, tag: Tag) -> None:
    self.tag: Tag = tag

  def getText(self, sel: str) -> str:
    """ Return text of HTML tag. None if tag not found """
    return self.tag.select(selector=sel)[0].text if self.tag.select(selector=sel) else None

  def getKey(self, sel: str, key: str) -> str:
    """ Return key of HTML tag. None if tag not found """
    return self.tag.select(selector=sel)[0].get(key=key) if self.tag.select(selector=sel) else None

  def getInt(self, sel: str) -> int:
    """ Return int in text of HTML tag. None if tag not found """
    r: ResultSet[Tag] = self.tag.select(selector=sel)
    if r:
      return int(''.join(filter(str.isdigit, r[0].text)))
    return None

  def getFloat(self, sel: str) -> float:
    """ Return float in text of HTML tag. None if tag not found """
    r: ResultSet[Tag] = self.tag.select(selector=sel)
    if r:
      numbers: list[Any] = re.findall(pattern=r"[-+]?\d*\.\d+|\d+", string=r[0].text)
      return float(''.join(numbers))
    return None
