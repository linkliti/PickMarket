""" Categories Test File """
# pylint: disable = unused-import, invalid-name, import-error, redefined-outer-name, unused-argument line-too-long, missing-function-docstring
import logging

import pytest
import requests

from .test_base import logger

# Define the base URL for the tests
BASE_URL = "http://localhost:5004/categories"

# Define shared parameters
params = [
  ("ozon", "/category/smartfony-15502/"),
  ("ozon", "/category/tehnika-dlya-kuhni-10523/"),
]


@pytest.mark.parametrize("market, url", params)
def test_filter(logger: None, market: str, url: str) -> None:
  response: requests.Response = requests.get(f"{BASE_URL}/{market}/filter",
                                             params={"url": url},
                                             timeout=10)
  logging.debug(response.text)
  assert response.status_code == 200


@pytest.mark.parametrize("market, url", params)
def test_sub(logger: None, market: str, url: str) -> None:
  response: requests.Response = requests.get(f"{BASE_URL}/{market}/sub",
                                             params={"url": url},
                                             timeout=10)
  logging.debug(response.text)
  assert response.status_code == 200


@pytest.mark.parametrize("market", [
  ("ozon"),
])
def test_root(logger: None, market: str) -> None:
  response: requests.Response = requests.get(f"{BASE_URL}/{market}/root", timeout=10)
  logging.debug(response.text)
  assert response.status_code == 200
