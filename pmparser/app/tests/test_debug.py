""" Debug """
# pylint: disable = unused-import, invalid-name, import-error, redefined-outer-name, unused-argument line-too-long
import os

from urllib.parse import quote
import json
import logging
from seleniumbase import undetected
from app.parsers.baseParser import Parser
from app.tests.test_base import logger
from app.selenium.selenium import startSelenium
from app.parsers.ozon.ozonParser import OzonParser
log = logging.getLogger(__name__)


urllist: list[dict[str, str]] = [
{
  "tag": "RootCategories",
  "url": "/api/entrypoint-api.bx/page/json/v2?url=/modal/categoryMenuRoot",
  "type": "json",
  "key": "categoryMenuRoot",
  "mobile": True
},
{
  "tag": "LVL2Categories",
  "url": "/api/entrypoint-api.bx/page/json/v2?url=/modal/categoryMenu/category/zhenskaya-odezhda-7501",
  "type": "json",
  "key": "categoryMenu",
  "mobile": True
},
{
  "tag": "LVL3PlusCategories",
  "url": "/api/entrypoint-api.bx/page/json/v2?url=/modal/categoryMenu/category/odezhda-obuv-i-aksessuary-7500",
  "type": "json",
  "key": "categoryMenu",
  "mobile": True
},
# {
#   "tag": "Prediction",
#   "url": "/search/?text=TPCell",
#   "type": "html",
#   "key": "location.replace(...)",
#   "mobile": False
# },
{
  "tag": "FiltersSearch",
  "url": "/api/entrypoint-api.bx/page/json/v2?url=/modal/allFilters/search/?text=энергетики+Monster+Nitro",
  "type": "json",
  "key": "filtersDesktop",
  "mobile": False
},
{
  "tag": "FiltersCategory",
  "url": "/api/entrypoint-api.bx/page/json/v2?url=/modal/allFilters/category/svitery-dzhempery-i-kardigany-zhenskie-7537",
  "type": "json",
  "key": "filtersDesktop",
  "mobile": False
},
{
  "tag": "ItemParams",
  "url": "/api/entrypoint-api.bx/page/json/v2?url=/product/monster-energy-nitro-super-dry-irlandiya-12-sht-h-500-ml-1181224292/features",
  "type": "json",
  "key": "webCharacteristics",
  "mobile": True
},
{
  "tag": "ItemsList",
  "url": "/api/entrypoint-api.bx/page/json/v2?url=/category/zhenskaya-odezhda-7501/?text=свитер",
  "type": "json",
  "key": "searchResultsV2",
  "mobile": True
}]

def test_getAllJsons(logger: None):
  """ Get All JSON from urlList and save to files in .temp/jsons folder """
  # https://www.ozon.ru
  # log: logging.Logger = logging.getLogger(__name__)
  p = Parser()
  os.makedirs("./.temp/jsons", exist_ok=True)
  driver: undetected.Chrome = startSelenium(uc=True, headless=True, mobile=False)
  mobile: undetected.Chrome = startSelenium(uc=True, headless=True, mobile=True)
  for item in urllist:
    print(item)
    encodedUrl: str = quote(item["url"], safe=':/?+=')
    if item["mobile"]:
      data: str = p.getData(host="www.ozon.ru", url=encodedUrl, driver=mobile)
    else:
      data: str = p.getData(host="www.ozon.ru", url=encodedUrl, driver=driver)
    with open(file=f"./.temp/jsons/{item["tag"]}.{item["type"]}", mode="w", encoding="utf-8") as f:
      f.write(data)

def test_extractLocalJsons(logger: None):
  """ Extract jsons using keys and save to files in .temp/parsed folder """
  p = OzonParser()
  os.makedirs("./.temp/parsed", exist_ok=True)
  for item in urllist:
    if item["key"] == "???":
      continue
    with open(file=f"./.temp/jsons/{item['tag']}.{item['type']}", mode="r", encoding="utf-8") as f:
      data: str = f.read()
      j: dict = json.loads(data)
      # Getting right JS key
      js = p.getEmbededJson(j=j['widgetStates'], keyName=item['key'].rsplit(sep='.', maxsplit=1)[-1])
      with open(file=f"./.temp/parsed/{item['tag']}.{item['type']}", mode="w", encoding="utf-8") as f:
        f.write(json.dumps(js, ensure_ascii=False))
