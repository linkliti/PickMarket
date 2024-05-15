""" Items Test File """
# pylint: disable = unused-import, invalid-name, import-error, redefined-outer-name, unused-argument, line-too-long, missing-function-docstring
import logging
from typing import Any, Dict, List, Optional, Tuple, Union

import app.protos.items_pb2 as itemsPB
import app.protos.reqHandlerTypes_pb2 as reqTypesPB
import app.protos.types_pb2 as typesPB
import pytest
import requests
from google.protobuf.json_format import MessageToDict, ParseDict

from .test_base import logger

# Define the base URL for the tests
BASE_URL = "http://localhost:5004/calc"
OZON = typesPB.Markets.OZON

# Define shared parameters
request_params: list = [
  (OZON, "/category/smartfony-15502/", "Xiaomi", "", 1, [
    ("brand", "listVal", ["Xiaomi"], 1),
    ("pm_price", "numVal", 0, 1),
  ]),
  (OZON, "/category/naushniki-15547/", "", "", 3, [
    ("brand", "listVal", ["Apple"], 5),
    ("color", "listVal", ['Белый', 'Синий'], 3),
    ("pm_isoriginal", "listVal", ["Да"], 5),
    ("pm_rating", "numVal", 2.5, 5),
  ]),
]


def create_value(t: str, value: list[str] | float):
  if t == "listVal":
    return {"values": value}
  elif t == "numVal":
    return value
  else:
    raise ValueError(f"Invalid type: {t}")


@pytest.mark.parametrize("request_data", request_params)
def test_list(logger: None, request_data: Tuple[typesPB.Markets, str, str, str, int,
                                                list]) -> None:
  # Unpack request_data
  market, pageUrl, userQuery, params, numOfPages, prefs_params = request_data

  # Construct the request body
  request = itemsPB.ItemsRequest(market=market,
                                 pageUrl=pageUrl,
                                 userQuery=userQuery,
                                 params=params,
                                 numOfPages=numOfPages)

  prefs: Dict[str, reqTypesPB.UserPref] = {
    key: ParseDict({
      "priority": priority,
      t: create_value(t, value)
    }, reqTypesPB.UserPref()) for key, t, value, priority in prefs_params
  }

  body = reqTypesPB.ItemsRequestWithPrefs(request=request, prefs=prefs)
  bodyJSON: Dict[str, Any] = MessageToDict(body)
  logging.debug(bodyJSON)
  response: requests.Response = requests.post(f"{BASE_URL}/ozon/list",
                                              json=bodyJSON,
                                              timeout=30)
  logging.debug(response.text)
  assert response.status_code == 200
