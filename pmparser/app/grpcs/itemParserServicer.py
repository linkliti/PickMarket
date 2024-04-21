""" Item Parser GRPC servicer """
import logging
from typing import Generator

import grpc
from app.parsers.ozon.ozonFilterMerge import OzonFilterMerge
from app.parsers.ozon.ozonParserChars import OzonParserChars
# from app.parsers.ozon.ozonParserFilters import OzonParserFilters
from app.parsers.ozon.ozonParserItems import OzonParserItems
from app.protos import items_pb2 as itemsPB
from app.protos import items_pb2_grpc as itemsPBgrpc
from app.protos import types_pb2 as typesPB

log = logging.getLogger(__name__)


class PMItemParserServicer(itemsPBgrpc.ItemParserServicer):
  """Implementation of ItemParserServicer"""

  def GetItems(self, request: itemsPB.ItemsRequest,
               context: grpc.ServicerContext) -> Generator[itemsPB.ItemResponse, None, None]:
    """Get items from marketplace"""
    gen: Generator[itemsPB.Item, None, None] | None = None
    market: typesPB.Markets = request.market
    pageUrl: str = request.pageUrl
    userQuery: str | None = request.userQuery if request.HasField('userQuery') else None
    params: str | None = request.params if request.HasField('params') else None
    numOfPages: int | None = request.numOfPages if request.HasField('numOfPages') else None
    try:
      match market:
        case typesPB.Markets.OZON:
          p = OzonParserItems()
          gen = p.getItems(pageUrl=pageUrl, query=userQuery, params=params, numOfPages=numOfPages)
      if gen:
        for item in gen:
          resp = itemsPB.ItemResponse(item=item)
          yield resp
    except Exception as e:  # pylint: disable=broad-except
      log.error("GetItems exception", extra={"error": str(e)})
      context.set_code(code=grpc.StatusCode.INTERNAL)
      context.set_details(details=str(e))
      return

  def GetItemCharacteristics(
      self, request: itemsPB.CharacteristicsRequest,
      context: grpc.ServicerContext) -> Generator[itemsPB.CharacteristicResponse, None, None]:
    """Get item characteristics"""
    gen: Generator[itemsPB.Characteristic, None, None] | None = None
    market: typesPB.Markets = request.market
    itemUrl: str = request.itemUrl
    try:
      match market:
        case typesPB.Markets.OZON:
          p = OzonParserChars()
          gen = p.getItemChars(itemUrl=itemUrl)
      if gen:
        for char in gen:
          resp = itemsPB.CharacteristicResponse(char=char)
          yield resp
    except Exception as e:  # pylint: disable=broad-except
      log.error("GetItemCharacteristics exception", extra={"error": str(e)})
      context.set_code(code=grpc.StatusCode.INTERNAL)
      context.set_details(details=str(e))
      return

  def GetCategoryFilters(
      self, request: itemsPB.FiltersRequest,
      context: grpc.ServicerContext) -> Generator[itemsPB.FilterResponse, None, None]:
    gen: Generator[itemsPB.Filter, None, None] | None = None
    market: typesPB.Markets = request.market
    categoryUrl: str = request.categoryUrl
    try:
      match market:
        case typesPB.Markets.OZON:
          # p = OzonParserFilters()
          # gen = p.getRootFilters(categoryUrl=categoryUrl)
          p = OzonFilterMerge()
          gen = p.getMergedFilters(categoryUrl=categoryUrl)
      if gen:
        for filt in gen:
          resp = itemsPB.FilterResponse(filter=filt)
          yield resp
    except Exception as e:  # pylint: disable=broad-except
      log.error("GetCategoryFilters exception", extra={"error": str(e)})
      context.set_code(code=grpc.StatusCode.INTERNAL)
      context.set_details(details=str(e))
      return
