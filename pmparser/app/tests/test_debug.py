""" Debug """
# pylint: disable = unused-import, invalid-name, import-error, redefined-outer-name, unused-argument line-too-long
import os

from seleniumbase import undetected
from app.parsers.baseParser import Parser
from app.tests.test_base import logger
from app.selenium.selenium import startSelenium
from urllib.parse import quote

def test_getAllJsons(logger):
  """ Get All JSON from urlList and save to files in .temp/jsons folder """
  # https://www.ozon.ru
  urlList = {
    # Root Categories -> categoryMenuRoot
    "RootCategories":
      "/api/entrypoint-api.bx/page/json/v2?url=/modal/categoryMenuRoot",
    # LVL 2 Categories -> categoryMenu
    "LVL2Categories":
      "/api/entrypoint-api.bx/page/json/v2?url=/modal/categoryMenu/category/zhenskaya-odezhda-7501",
    # LVL 3+ Categories -> categoryMenu
    "LVL3PlusCategories":
      "/api/entrypoint-api.bx/page/json/v2?url=/modal/categoryMenu/category/odezhda-obuv-i-aksessuary-7500",
    # Search category prediction -> HTML -> location.replace(...)
    "Prediction":
      "/search/?text=TPCell",
    # Filters from search -> ???
    "FiltersSearch":
      "/api/entrypoint-api.bx/page/json/v2?url=/modal/allFilters/search/?text=энергетики+Monster+Nitro",
    # Filters from category -> ???
    "FiltersCategory":
      "/api/entrypoint-api.bx/page/json/v2?url=/modal/allFilters/category/svitery-dzhempery-i-kardigany-zhenskie-7537/",
    # Item params -> webCharacteristics
    "ItemParams":
      "/api/entrypoint-api.bx/page/json/v2?url=/product/monster-energy-nitro-super-dry-irlandiya-12-sht-h-500-ml-1181224292/features",
    # Items -> searchResultsV2
    "ItemsList":
      "/api/entrypoint-api.bx/page/json/v2?url=/category/zhenskaya-odezhda-7501/?text=свитер",
  }

  # log: logging.Logger = logging.getLogger(__name__)
  p = Parser()
  os.makedirs("./.temp/jsons", exist_ok=True)
  driver: undetected.Chrome = startSelenium(uc=True, headless=True, mobile=True)
  for key, url in urlList.items():
    print("key: ", key, " url: ", url)
    encodedUrl: str = quote(url, safe=':/?+=')
    data = p.getData(host="www.ozon.ru", url=encodedUrl, useMobile=True, driver=driver)
    with open(file=f"./.temp/jsons/{key}.json", mode="w", encoding="utf-8") as f:
      f.write(data)
