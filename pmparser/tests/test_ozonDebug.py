""" Debug JSON files from OZON """
# pylint: disable = unused-import, invalid-name, import-error, redefined-outer-name, unused-argument line-too-long
import json
import logging
import os
from urllib.parse import quote

from app.parsers.baseParser import Parser
from app.parsers.ozon.ozonParser import OzonParser
from app.selen.selenium import startSelenium
from seleniumbase import undetected

from .test_base import logger

log = logging.getLogger(__name__)

PAGEAPI = "/api/entrypoint-api.bx/page/json/v2?url="
urllist: list[dict[str, str | bool]] = [
{
  "tag": "RootCategories",
  "url": PAGEAPI + "/modal/categoryMenuRoot",
  "type": "json",
  "key": "categoryMenuRoot",
  "mobile": True
},
{
  "tag": "LVL2Categories",
  "url": PAGEAPI + "/modal/categoryMenu/category/zhenskaya-odezhda-7501",
  "type": "json",
  "key": "categoryMenu",
  "mobile": True
},
{
  "tag": "LVL3PlusCategories",
  "url": PAGEAPI + "/modal/categoryMenu/category/odezhda-obuv-i-aksessuary-7500",
  "type": "json",
  "key": "categoryMenu",
  "mobile": True
},
{
  "tag": "ItemParams",
  "url": PAGEAPI + "/product/monster-energy-nitro-super-dry-irlandiya-12-sht-h-500-ml-1181224292/features",
  "type": "json",
  "key": "webCharacteristics",
  "mobile": True
},
{
  "tag": "ItemsList",
  "url": PAGEAPI + "/category/zhenskaya-odezhda-7501/?text=свитер",
  "type": "json",
  "key": "searchResultsV2",
  "mobile": True
},
{
  "tag": "FiltersCategory",
  "url": PAGEAPI + "/modal/filters/category/videokarty-15721/?all_filters=t",
  "type": "json",
  "key": "filters",
  "mobile": True
},
{
  "tag": "FilterValues",
  "url": PAGEAPI + "/modal/filterValues?filter=gpuseries&search_uri=/category/videokarty-15721/",
  "type": "json",
  "key": "filterValues",
  "mobile": True
}]

urllist = [{
  "tag": "FilterBrands",
  "url": PAGEAPI + "/modal/filterValues?filter=brand&search_uri=/category/videokarty-15721/",
  "type": "json",
  "key": "filterValues",
  "mobile": True
}]
# bottomCell rightButton

def test_getAllJsons(logger: None):
  """ Get All JSON from urlList and save to files in .temp/jsons folder """
  # https://www.ozon.ru
  # log: logging.Logger = logging.getLogger(__name__)
  p = Parser()
  os.makedirs("./.temp/jsons", exist_ok=True)
  # driver: undetected.Chrome = startSelenium(uc=True, mobile=False)
  mobile: undetected.Chrome = startSelenium(uc=True, mobile=True)
  for item in urllist:
    print(item)
    encodedUrl: str = quote(string=str(item["url"]), safe=':/?+=')
    # if item["mobile"]:
    data: str = p.getData(host="www.ozon.ru", url=encodedUrl, useMobile=True, driver=mobile)
    # else:
    #   data: str = p.getData(host="www.ozon.ru", url=encodedUrl, driver=driver)
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
      js = p.getEmbededJson(j=j['widgetStates'], keyName=str(item['key']).rsplit(sep='.', maxsplit=1)[-1])
      with open(file=f"./.temp/parsed/{item['tag']}.{item['type']}", mode="w", encoding="utf-8") as f:
        f.write(json.dumps(js, ensure_ascii=False))
