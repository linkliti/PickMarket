""" Selenium Module for fallbacks """
import logging
import json
from http.cookies import SimpleCookie
from bs4 import BeautifulSoup, NavigableString, Tag
from seleniumbase import undetected

from .selenium import checkForBlock, startSelenium

log = logging.getLogger(__name__)


def getDataFallback(url: str, header: dict, driver: undetected.Chrome | None = None) -> str:
  """Fallback method for getting data"""
  canBeClosed = False
  # Init driver
  if driver is None:
    ua = header.get('User-Agent', None) # type: ignore
    cookies = header.get('Cookie', None) # type: ignore
    mobile = bool(ua)
    canBeClosed = True
    driver = startSelenium(uc=True, mobile=mobile)
    if cookies:
      driver.default_get("https://www.ozon.ru")
      cookie = SimpleCookie()
      cookie.load(rawdata=cookies)
      for k, v in cookie.items():
        cookieDict: dict = {}
        cookieDict["name"] = k
        cookieDict["value"] = v.value
        driver.add_cookie(cookie_dict=cookieDict)

  # Get Page
  driver.default_get(url)
  if checkForBlock(data=driver.page_source):
    driver.get(url)
    driver.sleep(5)
  # Check for block
  data: str = driver.page_source
  if canBeClosed:
    driver.close()
  if checkForBlock(data=data):
    log.error("Failed to bypass block: %s", data)
    raise Exception("Failed to bypass block")

  # Try to convert to JSON
  try:
    soup = BeautifulSoup(markup=data, features='html.parser')
    preTag: Tag | NavigableString | None = soup.find(name='pre')
    if not preTag:
      raise Exception("Failed to parse webpage")
    jsonData: dict = json.loads(s=preTag.text)
    return json.dumps(obj=jsonData, ensure_ascii=False)
  except:
    return data
