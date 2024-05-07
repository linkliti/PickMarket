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
            guessedMaxValue: int = self.guessMaxValue(char.numVal)
            rangeFilter = typesPB.RangeFilter(min=0, max=guessedMaxValue)
            newFilt = itemsPB.Filter(title=char.name,
                                     key=char.key.lower(),
                                     internalType=typesPB.Filters.RANGE,
                                     rangeFilter=rangeFilter)
            filterList += [newFilt]
            filterKeys.add(newFilt.key)
      # Matching keys and filters
      compatibleKeys: set[str] = charsKeys.intersection(filterKeys)
      compatibleFilters: list[itemsPB.Filter] = list(
        filter(lambda x: x.key in compatibleKeys, filterList))
      # Extra filters
      pmFilters: list[itemsPB.Filter] = self.appendExtraFilters(itemList=itemList)
      # Yield all filters
      for filt in compatibleFilters + pmFilters:
        yield filt

  def guessMaxValue(self, num: float) -> int:
    """ Guess max value """
    return 10**(len(str(int(num))) + 1)

  def appendExtraFilters(self, itemList: list[itemsPB.Item]) -> list[itemsPB.Filter]:
    """ Append extra filters """
    pmFilters: list[itemsPB.Filter] = []
    maxComments: float = max(map(lambda x: x.comments, itemList), default=0)
    maxOldPrice: float = max(map(lambda x: x.oldPrice, itemList), default=0)
    maxPrice: float = max(map(lambda x: x.price, itemList), default=0)
    commentsFilter = itemsPB.Filter(title="Отзывы",
                                    key="pm_reviews",
                                    internalType=typesPB.Filters.RANGE,
                                    rangeFilter=typesPB.RangeFilter(
                                      min=0, max=self.guessMaxValue(maxComments)))
    ratingFilter = itemsPB.Filter(title="Рейтинг",
                                  key="pm_rating",
                                  internalType=typesPB.Filters.RANGE,
                                  rangeFilter=typesPB.RangeFilter(min=0, max=5))
    oldPriceFilter = itemsPB.Filter(title="Старая цена",
                                    key="pm_oldprice",
                                    internalType=typesPB.Filters.RANGE,
                                    rangeFilter=typesPB.RangeFilter(
                                      min=0, max=self.guessMaxValue(maxOldPrice)))
    priceFilter = itemsPB.Filter(title="Цена",
                                 key="pm_price",
                                 internalType=typesPB.Filters.RANGE,
                                 rangeFilter=typesPB.RangeFilter(min=0,
                                                                 max=self.guessMaxValue(maxPrice)))
    isOriginalFilter = itemsPB.Filter(title="Оригинал",
                                      key="pm_isoriginal",
                                      internalType=typesPB.Filters.BOOL,
                                      boolFilter=typesPB.BoolFilter(value="t"))
    isAdultFilter = itemsPB.Filter(title="Для взрослых",
                                   key="pm_isadult",
                                   internalType=typesPB.Filters.BOOL,
                                   boolFilter=typesPB.BoolFilter(value="t"))
    pmFilters += [
      commentsFilter, ratingFilter, oldPriceFilter, priceFilter, isOriginalFilter, isAdultFilter
    ]
    return pmFilters
