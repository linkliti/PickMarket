""" Ozon Parser Module for categories """
import logging
import re
from concurrent.futures import ThreadPoolExecutor
from itertools import islice
from typing import Any, Generator, Iterator

from app.parsers.ozon.ozonParser import OzonParser
from app.protos import items_pb2 as itemsPB
from app.protos import types_pb2 as typesPB
from app.utilities.jsonUtil import toJson

log = logging.getLogger(__name__)

TOOLARGEFILTERS: list[str] = ["Продавец"]


class OzonParserFilters(OzonParser):
  """ Ozon Parser Module for items """

  def __init__(self) -> None:
    self.categoryUrl: str = ""
    self.fKnownTypes: list[str] = [
      "cellWithSubtitleToggleCounter", "tagFilter", "rangeFilter", "multipleRangesFilter",
      "brandFilter", "colorFilter"
    ]
    self.tooLargeFilterTitles: list[str] = ["Продавец"]
    super().__init__()

  def getRootFilters(self, categoryUrl: str) -> Generator[itemsPB.Filter, None, None]:
    """ Get filters for category """
    self.categoryUrl = categoryUrl
    log.info('Getting filters: %s', self.categoryUrl)
    reqParams: dict[str, str] = {"all_filters": "t"}
    jString: str = self.getData(host=self.host,
                                url=self.api + "/modal/filters" + self.categoryUrl,
                                params=reqParams)

    log.info('Converting to JSON: %s', self.categoryUrl)
    j: dict = toJson(jString)
    j = self.getEmbededJson(j=j["widgetStates"], keyName="filters")

    log.info('Parsing filters: %s', self.categoryUrl)
    # Merging with extra filters in j["sections"][2]
    filterList: list = j["sections"][1]["filters"]
    if len(j["sections"]) > 2 and "filters" in j["sections"][2]:
      filterList.extend(j["sections"][2]["filters"])

    with ThreadPoolExecutor() as executor:
      futures: Iterator[itemsPB.Filter | None] = executor.map(self.worker, filterList)
      for res in futures:
        fdata: itemsPB.Filter | None = res
        if fdata:
          yield fdata

  def worker(self, filt) -> itemsPB.Filter | None:
    """Get filter values from OZON"""
    fw = OzonFilterWorker(filterJson=filt, categoryUrl=self.categoryUrl)
    return fw.returnFilter()


class OzonFilterWorker(OzonParser):
  """Get filter values from OZON"""

  def __init__(self, filterJson: dict, categoryUrl: str) -> None:
    super().__init__()
    self.j: dict = filterJson
    self.categoryUrl: str = categoryUrl
    # Defaults
    self.title: str = ""
    self.key: str = ""
    self.externalType: str = self.j["type"]
    self.internalType: typesPB.Filters | None = None
    self.rangeFilter: typesPB.RangeFilter | None = None
    self.selectionFilter: typesPB.SelectionFilter | None = None
    self.boolFilter: typesPB.BoolFilter | None = None
    self.getFilterValues()

  def getFilterValues(self) -> None:
    """ Get filter values """
    if self.externalType == "multipleRangesFilter":
      self.externalType = "rangeFilter"
      self.j = self.j["multipleRangesFilter"]

    self.title: str = self.j[self.externalType]["title"]
    log.debug('filter data', extra={"externalType": self.externalType, "title": self.title})

    match self.externalType:
      case "colorFilter":
        self.selectionFilter = self.getColorFilterValues()
        self.internalType = self.determineSelectionType(isRadio=self.selectionFilter.isRadio)
        return
      case "brandFilter":
        self.selectionFilter = self.getBrandFilterValues()
        self.internalType = self.determineSelectionType(isRadio=self.selectionFilter.isRadio)
        return
      case "tagFilter":
        self.selectionFilter = self.getTagFilterValues()
        self.internalType = self.determineSelectionType(isRadio=self.selectionFilter.isRadio)
        return
      case "rangeFilter":
        self.rangeFilter = self.getRangeFilterValues()
        self.internalType = typesPB.Filters.RANGE
        return
      case "cellWithSubtitleToggleCounter":
        self.boolFilter = self.getBoolFilterValues()
        self.internalType = typesPB.Filters.BOOL
        return
      case "cellWithSubtitleCounter":
        self.selectionFilter = self.getCollapsedTagFilterValues()
        self.internalType = self.determineSelectionType(isRadio=self.selectionFilter.isRadio)
        return
      case _:
        log.warning('Unknown filter type: %s', self.externalType)
        return

  def returnFilter(self) -> itemsPB.Filter | None:
    """Return filter"""
    if self.internalType is None:
      return None
    filt: itemsPB.Filter | None = None
    kwargs: dict[str, Any] = {
      'title': self.title,
      'key': self.key,
      'externalType': self.externalType,
      'internalType': self.internalType
    }
    if self.rangeFilter:
      kwargs['rangeFilter'] = self.rangeFilter
    elif self.selectionFilter:
      kwargs['selectionFilter'] = self.selectionFilter
    elif self.boolFilter:
      kwargs['boolFilter'] = self.boolFilter
    filt: itemsPB.Filter | None = itemsPB.Filter(**kwargs)
    return filt

  def determineSelectionType(self, isRadio) -> typesPB.Filters:
    """Determine selection type"""
    return typesPB.Filters.SELECTIONRADIO if isRadio else typesPB.Filters.SELECTION

  def getColorFilterValues(self) -> typesPB.SelectionFilter:
    """Get color values for filter"""
    isRadio = False
    link: str = self.j[self.externalType]["colors"][0]["action"]["link"]
    self.key = link.split("&")[1].split("=")[0]
    items: list[typesPB.SelectionFilterItem] = []
    # Colors
    j: dict = self.getJsonWithMoreValues()
    for section in j["sections"]:
      for item in section["values"]:
        data: dict = item["cellWithSubtitle24IconCheckboxRadioCounter"]
        text: str = data["title"]
        value: str = data["action"]["params"]["value"]
        o = typesPB.SelectionFilterItem(text=text, value=value)
        items.append(o)
    return typesPB.SelectionFilter(isRadio=isRadio, items=items)

  def getBrandFilterValues(self) -> typesPB.SelectionFilter:
    """Get brand values for filter"""
    isRadio = bool(
      re.search(
        pattern=r"-radio-filter$",
        string=self.j[self.externalType]["roundedCells"][0]["testInfo"]["automatizationId"]))
    link: str = self.j[self.externalType]["roundedCells"][0]["action"]["link"]
    self.key = link.split("&")[1].split("=")[0]
    items: list[typesPB.SelectionFilterItem] = []
    # Popular brands
    for tag in self.j[self.externalType]["roundedCells"]:
      text: str = tag["title"]
      link = tag["action"]["link"]
      value: str = link.split("&")[-1].split("=")[-1]
      o = typesPB.SelectionFilterItem(text=text, value=value)
      items.append(o)
    # Other Brands
    if "bottomCell" in self.j[self.externalType]:
      j: dict = self.getJsonWithMoreValues()
      sections = islice(j["sections"], 1, None)
      for section in sections:
        for item in section["values"]:
          data: dict = item["cellWithSubtitleCheckboxRadioCounter"]
          text: str = data["title"]
          value: str = data["action"]["params"]["value"]
          o = typesPB.SelectionFilterItem(text=text, value=value)
          items.append(o)
    return typesPB.SelectionFilter(isRadio=isRadio, items=items)

  def getTagFilterValues(self) -> typesPB.SelectionFilter:
    """Get tag values for filter"""
    isRadio = bool(
      re.search(
        pattern=r"-radio-filter$",
        string=self.j[self.externalType]["tags"][0]["tag"]["testInfo"]["automatizationId"]))
    # Key
    link: str = self.j[self.externalType]["tags"][0]["tag"]["action"]["link"]
    self.key = link.split("&")[1].split("=")[0]
    items: list[typesPB.SelectionFilterItem] = []
    # "All tags" button present
    if "rightButton" in self.j[self.externalType] and self.title not in TOOLARGEFILTERS:
      j: dict = self.getJsonWithMoreValues()
      for section in j["sections"]:
        for item in section["values"]:
          data: dict = item[item["type"]]
          text: str = data["title"]
          value: str = data["action"]["params"]["value"]
          o = typesPB.SelectionFilterItem(text=text, value=value)
          items.append(o)
    # No "All tags" button present OR filter is too large
    for tag in self.j[self.externalType]["tags"]:
      text: str = tag["tag"]["text"]
      link = tag["tag"]["action"]["link"]
      value: str = link.split("&")[-1].split("=")[-1]
      o = typesPB.SelectionFilterItem(text=text, value=value)
      items.append(o)
    return typesPB.SelectionFilter(isRadio=isRadio, items=items)

  def getRangeFilterValues(self) -> typesPB.RangeFilter:
    """Get range values for filter"""
    link: str = self.j[self.externalType]["action"]["link"]
    self.key = link.split("&")[1].split("=")[0]
    # Min-max values
    minValue: int = self.j[self.externalType]["minValue"]
    maxValue: int = self.j[self.externalType]["maxValue"]
    return typesPB.RangeFilter(min=minValue, max=maxValue)

  def getBoolFilterValues(self) -> typesPB.BoolFilter:
    """Get bool values for filter"""
    # Key
    link: str = self.j[self.externalType]["action"]["link"]
    self.key = link.split("&")[1].split("=")[0]
    # Boolean value ('t')
    value: str = link.split("&")[-1].split("=")[-1]
    # Return
    return typesPB.BoolFilter(value=value)

  def getCollapsedTagFilterValues(self) -> typesPB.SelectionFilter:
    """Get collapsed tag values for filter"""
    isRadio = False
    # Key
    link: str = self.j[self.externalType]["action"]["link"]
    self.key = link.split("&")[1].split("=")[1]
    items: list[typesPB.SelectionFilterItem] = []
    # Get tags
    j: dict = self.getJsonWithMoreValues()
    for section in j["sections"]:
      for item in section["values"]:
        data: dict = item[item["type"]]
        text: str = data["title"]
        value: str = data["action"]["params"]["value"]
        o = typesPB.SelectionFilterItem(text=text, value=value)
        items.append(o)
    return typesPB.SelectionFilter(isRadio=isRadio, items=items)

  def getJsonWithMoreValues(self) -> dict:
    """Get values for filter from OZON"""
    log.info('Getting more filter values for key: %s @ %s', self.key, self.categoryUrl)
    url: str = self.api + "/modal/filterValues?all_filters=t&filter=" + self.key +\
      "&search_uri=" + self.categoryUrl
    jString: str = self.getData(host=self.host, url=url)
    log.info('Converting to JSON: %s', url)
    j: dict = toJson(jString)
    j = self.getEmbededJson(j=j["widgetStates"], keyName="filterValues")
    return j
