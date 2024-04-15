""" Filter merge depending on items chars """

import logging
from concurrent.futures import Future, ThreadPoolExecutor
from typing import Generator

from app.parsers.ozon.ozonParserChars import OzonParserChars
from app.parsers.ozon.ozonParserFilters import OzonParserFilters
from app.parsers.ozon.ozonParserItems import OzonParserItems
from app.protos import items_pb2 as itemsPB
from app.protos import types_pb2 as typesPB

log = logging.getLogger(__name__)


class OzonFilterMerge:
  """ Filter merge depending on items chars """

  def __init__(self) -> None:
    self.itemsParser = OzonParserItems()
    self.filterParser = OzonParserFilters()
    self.charsParser = OzonParserChars()

  def getMergedFilters(self, categoryUrl: str) -> Generator[itemsPB.Filter, None, None]:
    """ Get merged filter """
    with ThreadPoolExecutor() as executor:
      # Filters and items
      itemList: list[itemsPB.Item]
      filterList: list[itemsPB.Filter]
      charsList: list[itemsPB.Characteristic] = []
      futures: list[Future] = [
        executor.submit(self.itemsParser.getItems, pageUrl=categoryUrl, numOfPages=1),
        executor.submit(self.filterParser.getRootFilters, categoryUrl=categoryUrl)
      ]
      itemList, filterList = list(futures[0].result()), list(futures[1].result())
      log.debug("Received items and filters list",
                extra={
                  "itemsLen": len(itemList),
                  "filtersLen": len(filterList),
                  "categoryUrl": categoryUrl
                })

      # Item chars
      def fetchAndAppendTask(item: itemsPB.Item) -> None:
        chars: Generator[itemsPB.Characteristic, None,
                         None] = self.charsParser.getItemChars(itemUrl=item.url)
        charsList.extend(chars)

      with ThreadPoolExecutor(max_workers=5) as executor:
        futures = [executor.submit(fetchAndAppendTask, item) for item in itemList[:5]]
        for future in futures:
          future.result()

      log.debug("Received chars list",
                extra={
                  "charsLen": len(charsList),
                  "categoryUrl": categoryUrl
                })
      # Unique keys
      charsKeys: set[str] = {char.key.lower() for char in charsList}
      filterKeys: set[str] = {filt.key for filt in filterList}
      # Add new filter for every chars with numVal
      char: itemsPB.Characteristic
      for char in charsList:
        if char.key.lower() not in filterKeys:
          # Add numeric filter
          if char.HasField("numVal"):
            # Guess max value (4 -> 100, 65 -> 1000, etc)
            guessedMaxValue: int = 10**(len(str(int(char.numVal))) + 1)
            rangeFilter = typesPB.RangeFilter(min=0, max=guessedMaxValue)
            newFilt = itemsPB.Filter(title=char.name,
                                     key=char.key.lower(),
                                     internalType=typesPB.Filters.RANGE,
                                     rangeFilter=rangeFilter)
          # Add bool filter
          elif len(char.listVal.values) > 1 and char.listVal.values[0] in ["Да", "Нет"]:
            newFilt = itemsPB.Filter(title=char.name,
                                     key=char.key.lower(),
                                     internalType=typesPB.Filters.BOOL,
                                     boolFilter=typesPB.BoolFilter(value="t"))
          else:
            continue
          filterList += [newFilt]
          filterKeys.add(newFilt.key)
      # Yield new filters
      compatibleKeys: set[str] = charsKeys.intersection(filterKeys)
      compatibleFilters: list[itemsPB.Filter] = list(
        filter(lambda x: x.key in compatibleKeys, filterList))
      for filt in compatibleFilters:
        yield filt
