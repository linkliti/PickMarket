""" Test GRPC interaction with ozon """
# pylint: disable=unused-import, invalid-name, import-error, redefined-outer-name, unused-argument
# pylint: disable=no-member, line-too-long

import logging

import grpc
import pytest
from app.protos import categories_pb2 as categPB
from app.protos import categories_pb2_grpc as categPBgrpc
from app.protos import items_pb2 as itemsPB
from app.protos import items_pb2_grpc as itemsPBgrpc
from app.protos import types_pb2 as typesPB

from .test_base import logger

OZON = typesPB.Markets.OZON


@pytest.fixture
def grpc_channel(request) -> grpc.Channel:
  """Connect to the appropriate gRPC service based on the request param"""
  target = request.param
  return grpc.insecure_channel(target)


# Parameters for category tests
category_params = [
  (OZON, "/category/kvadrokoptery-i-aksessuary-7159/"),
  (OZON, "/category/smartfony-15502/"),
  (OZON, "/category/noutbuki-15692/"),
]

# Parameters for item tests
item_params = [
  (OZON, "/category/kvadrokoptery-i-aksessuary-7159/", "", 1),
  (OZON, "/category/smartfony-15502/", "apple", 2),
  (OZON, "/category/noutbuki-15692/", "dell", 3),
]

# Parameters for item characteristics tests
item_char_params = [
  (OZON, '/product/xiaomi-smartfon-redmi-note-13-pro-12-512-gb-chernyy-1410043273/'),
  (OZON, '/product/samsung-televizor-qe43q60cau-43-4k-uhd-chernyy-1492516784/'),
]

# Parameters for root categories tests
root_category_params = [
  (OZON),
]


@pytest.mark.parametrize("grpc_channel", ['localhost:5000', 'localhost:5003'], indirect=True)
@pytest.mark.parametrize("market, categoryUrl, userQuery, numOfPages", item_params)
def test_grpcGetItems(logger: None, grpc_channel: grpc.Channel, market: typesPB.Markets,
                      categoryUrl: str, userQuery: str, numOfPages: int) -> None:
  """Test getItems via gRPC"""
  stub = itemsPBgrpc.ItemParserStub(channel=grpc_channel)
  response: itemsPB.ItemResponse = stub.GetItems(
    itemsPB.ItemsRequest(market=market,
                         pageUrl=categoryUrl,
                         userQuery=userQuery,
                         numOfPages=numOfPages))
  for item in response:  # type: ignore
    logging.debug(item)


@pytest.mark.parametrize("grpc_channel", ['localhost:5000', 'localhost:5003'], indirect=True)
@pytest.mark.parametrize("market, itemUrl", item_char_params)
def test_getItemChars(logger: None, grpc_channel: grpc.Channel, market: typesPB.Markets,
                      itemUrl: str) -> None:
  """Test getItemChars via gRPC"""
  stub = itemsPBgrpc.ItemParserStub(channel=grpc_channel)
  response: itemsPB.CharacteristicResponse = stub.GetItemCharacteristics(
    itemsPB.CharacteristicsRequest(market=market, itemUrl=itemUrl))
  for char in response:  # type: ignore
    logging.debug(char)


@pytest.mark.parametrize("grpc_channel", ['localhost:5000', 'localhost:5002'], indirect=True)
@pytest.mark.parametrize("market", root_category_params)
def test_grpcGetRootCategories(logger: None, grpc_channel: grpc.Channel,
                               market: typesPB.Markets) -> None:
  """Test getRootCategories via gRPC"""
  stub = categPBgrpc.CategoryParserStub(channel=grpc_channel)
  response: categPB.CategoryResponse = stub.GetRootCategories(
    categPB.RootCategoriesRequest(market=market))
  for category in response:  # type: ignore
    logging.debug(category)


@pytest.mark.parametrize("grpc_channel", ['localhost:5000', 'localhost:5002'], indirect=True)
@pytest.mark.parametrize("market, categoryUrl", category_params)
def test_grpcGetSubCategories(logger: None, grpc_channel: grpc.Channel, market: typesPB.Markets,
                              categoryUrl: str) -> None:
  """Test getSubCategories via gRPC"""
  stub = categPBgrpc.CategoryParserStub(channel=grpc_channel)
  response: categPB.CategoryResponse = stub.GetSubCategories(
    categPB.SubCategoriesRequest(market=market, categoryUrl=categoryUrl))
  logging.debug("Categories:")
  for category in response:  # type: ignore
    logging.debug(category)


@pytest.mark.parametrize("grpc_channel", ['localhost:5000', 'localhost:5003'], indirect=True)
@pytest.mark.parametrize("market, categoryUrl", category_params)
def test_grpcGetFilters(logger: None, grpc_channel: grpc.Channel, market: typesPB.Markets,
                        categoryUrl: str) -> None:
  """Test getFilters via gRPC"""
  stub = itemsPBgrpc.ItemParserStub(channel=grpc_channel)
  response: itemsPB.FilterResponse = stub.GetCategoryFilters(
    itemsPB.FiltersRequest(market=market, categoryUrl=categoryUrl))
  for filt in response:  # type: ignore
    logging.debug(filt)
