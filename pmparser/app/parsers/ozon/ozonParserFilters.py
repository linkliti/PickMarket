""" Ozon Parser Module for categories """
import json
import logging
import re
from itertools import islice
from typing import Generator, List

from app.parsers.baseDataclass import (BaseFiltersDataclass, BoolFilter, RangeFilter,
                                       SelectionFilter, SelectionFilterItem)
from app.parsers.ozon.ozonParser import OzonParser

log = logging.getLogger(__name__)


class OzonParserFilters(OzonParser):
  """ Ozon Parser Module for items """

  def __init__(self) -> None:
    self.pageUrl: str = ""
    self.fKnownTypes: list[str] = [
      "cellWithSubtitleToggleCounter", "tagFilter", "rangeFilter", "multipleRangesFilter",
      "brandFilter", "colorFilter"
    ]
    self.tooLargeFilterTitles: list[str] = ["Продавец"]
    super().__init__()

  def getRootFilters(self, pageUrl: str) -> Generator[BaseFiltersDataclass, None, None]:
    """ Get filters for category """
    self.pageUrl = pageUrl
    log.info('Getting filters: %s', self.pageUrl)
    jString: str = self.getData(host=self.host,
                                url=self.api + "/modal/filters" + self.pageUrl + "/?all_filters=t",
                                useMobile=True)

    log.info('Converting to JSON: %s', self.pageUrl)
    j: dict = json.loads(jString)
    j = self.getEmbededJson(j=j["widgetStates"], keyName="filters")

    log.info('Parsing filters: %s', self.pageUrl)
    for filt in j["sections"][1]["filters"]:
      data: BaseFiltersDataclass | Generator | None = self.getFilter(filt)
      if isinstance(data, BaseFiltersDataclass):
        yield data
      elif isinstance(data, Generator):
        yield from data

  def getFilter(self, fJSON: dict) -> BaseFiltersDataclass | Generator | None:
    """ Get filter """
    fType: str = fJSON["type"]
    fKey: str = ""
    internalType: str = ""
    fData: SelectionFilter | BoolFilter | RangeFilter | None = None
    # Unknown filter
    if fType not in self.fKnownTypes:
      log.warning('Unknown filter type: %s', fType)
      return None

    # MultiFilter fix for price: prefer rangeFilter
    if fType == "multipleRangesFilter":
      fType = "rangeFilter"
      fJSON = fJSON["multipleRangesFilter"]

    # Title
    fTitle: str = fJSON[fType]["title"]
    log.debug('Filter type: %s %s', fType, fTitle)
    link: str = ""

    if fType == "colorFilter":
      fKey: str = ""
      isRadio = False
      items: List[SelectionFilterItem] = []
      log.debug(json.dumps(fJSON))
      # Key
      link = fJSON[fType]["colors"][0]["action"]["link"]
      fKey = link.split("&")[1].split("=")[0]
      # "All colors" button present
      # if "rightButton" in fJSON[fType]:
      for tag in self.getColorFilterValues(fKey=fKey):
        items.append(SelectionFilterItem(text=tag.text, value=tag.value))
      # No "All colors" button present
      # else:
      #   for tag in fJSON[fType]["colors"]:
      #     text: str = tag["hex"]
      #     link = tag["action"]["link"]
      #     value: str = link.split("&")[-1].split("=")[-1]
      #     items.append(SelectionFilterItem(text=text, value=value))
      # Return
      internalType: str = "selectionFilter"
      fData = SelectionFilter(isRadio=isRadio, items=items)

    elif fType == "cellWithSubtitleToggleCounter":
      # Key
      link = fJSON[fType]["action"]["link"]
      fKey = link.split("&")[1].split("=")[0]
      # Boolean value ('t')
      value: str = link.split("&")[-1].split("=")[-1]
      # Return
      internalType: str = "boolFilter"
      fData = BoolFilter(value=value)

    elif fType == "brandFilter":
      isRadio = bool(
        re.search(pattern=r"-radio-filter$",
                  string=fJSON[fType]["roundedCells"][0]["testInfo"]["automatizationId"]))
      items: List[SelectionFilterItem] = []
      # Key
      link = fJSON[fType]["roundedCells"][0]["action"]["link"]
      fKey = link.split("&")[1].split("=")[0]
      # Popular brands
      for tag in fJSON[fType]["roundedCells"]:
        text: str = tag["title"]
        link = tag["action"]["link"]
        value: str = link.split("&")[-1].split("=")[-1]
        items.append(SelectionFilterItem(text=text, value=value))
      # Other brands
      if "bottomCell" in fJSON[fType] and not fTitle in self.tooLargeFilterTitles:
        for tag in self.getBrandFilterValues(fKey=fKey):
          items.append(SelectionFilterItem(text=tag.text, value=tag.value))
      # Return
      internalType: str = "selectionFilter" + ("Radio" if isRadio else "Checkbox")
      fData = SelectionFilter(isRadio=isRadio, items=items)

    elif fType == "tagFilter":
      isRadio = bool(
        re.search(pattern=r"-radio-filter$",
                  string=fJSON[fType]["tags"][0]["tag"]["testInfo"]["automatizationId"]))
      items: List[SelectionFilterItem] = []
      # Key
      link = fJSON[fType]["tags"][0]["tag"]["action"]["link"]
      fKey = link.split("&")[1].split("=")[0]
      # "All tags" button present
      if "rightButton" in fJSON[fType] and not fTitle in self.tooLargeFilterTitles:
        for tag in self.getTagFilterValues(fKey=fKey):
          items.append(tag)
      # No "All tags" button present OR filter is too large
      else:
        for tag in fJSON[fType]["tags"]:
          text: str = tag["tag"]["text"]
          link = tag["tag"]["action"]["link"]
          value: str = link.split("&")[-1].split("=")[-1]
          items.append(SelectionFilterItem(text=text, value=value))
      # Return
      internalType: str = "selectionFilter" + ("Radio" if isRadio else "Checkbox")
      fData = SelectionFilter(isRadio=isRadio, items=items)

    elif fType == "rangeFilter":
      # Key
      link = fJSON[fType]["action"]["link"]
      fKey = link.split("&")[1].split("=")[0]
      # Min-max values
      minValue: int = fJSON[fType]["minValue"]
      maxValue: int = fJSON[fType]["maxValue"]
      # Return
      internalType = "rangeFilter"
      fData = RangeFilter(min=minValue, max=maxValue)

    # BaseFilter Object
    if not fData:
      raise Exception("fData is None")
    baseFilter = BaseFiltersDataclass(title=fTitle,
                                      key=fKey,
                                      externalType=fType,
                                      internalType=internalType,
                                      data=fData)
    log.debug('Filter: %s', baseFilter)
    return baseFilter

  def getTagFilterValues(self, fKey: str) -> Generator[SelectionFilterItem, None, None]:
    """Get values for tag filter"""
    log.info('Getting filter values for key: %s %s', self.pageUrl, fKey)
    url: str = self.api + "/modal/filterValues?all_filters=t&filter=" + fKey +\
      "&search_uri=" + self.pageUrl
    jString: str = self.getData(host=self.host, url=url, useMobile=True)

    log.info('Converting to JSON: %s', url)
    j: dict = json.loads(jString)
    j = self.getEmbededJson(j=j["widgetStates"], keyName="filterValues")
    for section in j["sections"]:
      for item in section["values"]:
        data: dict = item[item["type"]]
        text: str = data["title"]
        value: str = data["action"]["params"]["value"]
        yield SelectionFilterItem(text=text, value=value)

  def getBrandFilterValues(self, fKey: str) -> Generator[SelectionFilterItem, None, None]:
    """Get values for brand filter"""
    log.info('Getting filter values for key: %s %s', self.pageUrl, fKey)
    url: str = self.api + "/modal/filterValues?all_filters=t&filter=" + fKey +\
      "&search_uri=" + self.pageUrl
    jString: str = self.getData(host=self.host, url=url, useMobile=True)

    log.info('Converting to JSON: %s', url)
    j: dict = json.loads(jString)
    j = self.getEmbededJson(j=j["widgetStates"], keyName="filterValues")
    sections = islice(j["sections"], 1, None)
    for section in sections:
      for item in section["values"]:
        data: dict = item["cellWithSubtitleCheckboxRadioCounter"]
        text: str = data["title"]
        value: str = data["action"]["params"]["value"]
        yield SelectionFilterItem(text=text, value=value)

  def getColorFilterValues(self, fKey: str) -> Generator[SelectionFilterItem, None, None]:
    """Get values for color filter"""
    log.info('Getting filter values for key: %s %s', self.pageUrl, fKey)
    url: str = self.api + "/modal/filterValues?all_filters=t&filter=" + fKey +\
      "&search_uri=" + self.pageUrl
    jString: str = self.getData(host=self.host, url=url, useMobile=True)

    log.info('Converting to JSON: %s', url)
    j: dict = json.loads(jString)
    j = self.getEmbededJson(j=j["widgetStates"], keyName="filterValues")
    sections: dict = j["sections"]
    for section in sections:
      for item in section["values"]:
        data: dict = item["cellWithSubtitle24IconCheckboxRadioCounter"]
        text: str = data["title"]
        value: str = data["action"]["params"]["value"]
        yield SelectionFilterItem(text=text, value=value)
