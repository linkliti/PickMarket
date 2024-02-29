""" Test GRPC interaction with ozon """
#pylint: disable = unused-import, invalid-name, import-error, redefined-outer-name, unused-argument
#pylint: disable=no-member

import grpc
from app.protos import categories_pb2 as categPB
from app.protos import categories_pb2_grpc as categPBgrpc
from app.protos import items_pb2 as itemsPB
from app.protos import items_pb2_grpc as itemsPBgrpc
from app.protos import types_pb2 as typesPB

from .test_base import channel, logger


def test_grpcGetItems(logger: None, channel: grpc.Channel) -> None:
  """Test getItems via gRPC"""
  stub = itemsPBgrpc.ItemParserStub(channel=channel)
  response: itemsPB.ItemResponse = stub.GetItems(
    itemsPB.ItemsRequest(market=typesPB.Markets.OZON,
                         pageUrl='/category/ingalyatory-i-aksessuary-35039',
                         numOfPages=2))
  for item in response:  # type: ignore
    print(item)


def test_getItemChars(logger: None, channel: grpc.Channel) -> None:
  """Test getItemChars via gRPC"""
  stub = itemsPBgrpc.ItemParserStub(channel=channel)
  response: itemsPB.CharacteristicResponse = stub.GetItemCharacteristics(
    itemsPB.CharacteristicsRequest(
      market=typesPB.Markets.OZON,
      itemUrl=
      '/product/gigabyte-videokarta-geforce-rtx-3080-10-gb-gv-n3080gaming-oc-10gd-2-0-lhr-306060738'
    ))
  for char in response:  # type: ignore
    print(char)


def test_grpcGetRootCategories(logger: None, channel: grpc.Channel) -> None:
  """Test getRootCategories via gRPC"""
  stub = categPBgrpc.CategoryParserStub(channel=channel)
  response: categPB.CategoryResponse = stub.GetRootCategories(
    categPB.RootCategoriesRequest(market=typesPB.Markets.OZON))
  for category in response:  # type: ignore
    print(category)


def test_grpcGetSubCategories(logger: None, channel: grpc.Channel) -> None:
  """Test getSubCategories via gRPC"""
  stub = categPBgrpc.CategoryParserStub(channel=channel)
  response: categPB.CategoryResponse = stub.GetSubCategories(
    categPB.SubCategoriesRequest(market=typesPB.Markets.OZON,
                                 categoryUrl='/category/zhenskaya-odezhda-7501'))
  for category in response:  # type: ignore
    print(category)


def test_grpcGetFilters(logger: None, channel: grpc.Channel) -> None:
  """Test getFilters via gRPC"""
  stub = categPBgrpc.CategoryParserStub(channel=channel)
  response: categPB.FilterResponse = stub.GetCategoryFilters(
    categPB.FiltersRequest(market=typesPB.Markets.OZON,
                           categoryUrl='/category/bluzy-i-rubashki-zhenskie-7511'))
  for filt in response:  # type: ignore
    print(filt)
