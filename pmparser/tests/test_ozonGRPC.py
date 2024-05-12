""" Test GRPC interaction with ozon """
#pylint: disable = unused-import, invalid-name, import-error, redefined-outer-name, unused-argument
#pylint: disable = no-member, line-too-long

import grpc
from app.protos import categories_pb2 as categPB
from app.protos import categories_pb2_grpc as categPBgrpc
from app.protos import items_pb2 as itemsPB
from app.protos import items_pb2_grpc as itemsPBgrpc
from app.protos import types_pb2 as typesPB

from .test_base import channel, logger


def test_grpcGetItems(logger: None, channel: grpc.Channel) -> None:
  """Test getItems via gRPC"""
  categoryUrl = "/category/smartfony-15502/"
  market: typesPB.Markets = typesPB.Markets.OZON

  stub = itemsPBgrpc.ItemParserStub(channel=channel)
  response: itemsPB.ItemResponse = stub.GetItems(
    itemsPB.ItemsRequest(market=market, pageUrl=categoryUrl, userQuery="apple", numOfPages=2))
  for item in response:  # type: ignore
    print(item)


def test_getItemChars(logger: None, channel: grpc.Channel) -> None:
  """Test getItemChars via gRPC"""
  itemUrl: str = \
    '/product/xiaomi-smartfon-redmi-note-13-pro-12-512-gb-chernyy-1410043273/'
  market: typesPB.Markets = typesPB.Markets.OZON

  stub = itemsPBgrpc.ItemParserStub(channel=channel)
  response: itemsPB.CharacteristicResponse = stub.GetItemCharacteristics(
    itemsPB.CharacteristicsRequest(market=market, itemUrl=itemUrl))
  for char in response:  # type: ignore
    print(char)


def test_grpcGetRootCategories(logger: None, channel: grpc.Channel) -> None:
  """Test getRootCategories via gRPC"""
  market: typesPB.Markets = typesPB.Markets.OZON

  stub = categPBgrpc.CategoryParserStub(channel=channel)
  response: categPB.CategoryResponse = stub.GetRootCategories(
    categPB.RootCategoriesRequest(market=market))
  for category in response:  # type: ignore
    print(category)


def test_grpcGetSubCategories(logger: None, channel: grpc.Channel) -> None:
  """Test getSubCategories via gRPC"""
  categoryUrl: str = "/category/kvadrokoptery-i-aksessuary-7159/"
  market: typesPB.Markets = typesPB.Markets.OZON

  stub = categPBgrpc.CategoryParserStub(channel=channel)
  response: categPB.CategoryResponse = stub.GetSubCategories(
    categPB.SubCategoriesRequest(market=market, categoryUrl=categoryUrl))
  print("Categories:")
  for category in response:  # type: ignore
    print(category)


def test_grpcGetFilters(logger: None, channel: grpc.Channel) -> None:
  """Test getFilters via gRPC"""
  categoryUrl: str = "/category/smartfony-15502/"
  market: typesPB.Markets = typesPB.Markets.OZON

  stub = itemsPBgrpc.ItemParserStub(channel=channel)
  response: itemsPB.FilterResponse = stub.GetCategoryFilters(
    itemsPB.FiltersRequest(market=market, categoryUrl=categoryUrl))
  for filt in response:  # type: ignore
    print(filt)
