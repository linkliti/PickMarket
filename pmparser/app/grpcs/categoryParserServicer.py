""" Category Parser GRPC servicer """
from typing import Generator

import grpc
from app.parsers.ozon.ozonParserCategories import OzonParserCategories
from app.parsers.ozon.ozonParserFilters import OzonParserFilters
from app.protos import categories_pb2 as categPB
from app.protos import categories_pb2_grpc as categPBgrpc
from app.protos import types_pb2 as typesPB


class PMCategoryParserServicer(categPBgrpc.CategoryParserServicer):
  """Implementation of CategoryParserServicer"""

  def GetRootCategories(
      self, request: categPB.RootCategoriesRequest,
      context: grpc.ServicerContext) -> Generator[categPB.CategoryResponse, None, None]:
    gen: Generator[categPB.Category, None, None] | None = None
    market: typesPB.Markets = request.market
    match market:
      case typesPB.Markets.OZON:
        p = OzonParserCategories()
        gen = p.getRootCategories()
    if gen:
      for category in gen:
        resp = categPB.CategoryResponse(category=category)
        yield resp

  def GetSubCategories(
      self, request: categPB.SubCategoriesRequest,
      context: grpc.ServicerContext) -> Generator[categPB.CategoryResponse, None, None]:
    gen: Generator[categPB.Category, None, None] | None = None
    market: typesPB.Markets = request.market
    categoryUrl: str = request.categoryUrl
    match market:
      case typesPB.Markets.OZON:
        p = OzonParserCategories()
        gen = p.getSubCategories(categoryUrl=categoryUrl)
    if gen:
      for category in gen:
        resp = categPB.CategoryResponse(category=category)
        yield resp

  def GetCategoryFilters(
      self, request: categPB.FiltersRequest,
      context: grpc.ServicerContext) -> Generator[categPB.FilterResponse, None, None]:
    gen: Generator[categPB.Filter, None, None] | None = None
    market: typesPB.Markets = request.market
    categoryUrl: str = request.categoryUrl
    match market:
      case typesPB.Markets.OZON:
        p = OzonParserFilters()
        gen = p.getRootFilters(categoryUrl=categoryUrl)
    if gen:
      for filt in gen:
        resp = categPB.FilterResponse(filter=filt)
        yield resp
