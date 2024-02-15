""" Selenium Module for fallbacks """
import logging
import json
from bs4 import BeautifulSoup, Tag
from seleniumbase import undetected

from app.selenium.selenium import checkForBlock, startSelenium

log = logging.getLogger(__name__)


def getDataFallback(url: str, header: dict[dict], driver: undetected.Chrome = None) -> str:
  """Fallback method for getting data"""
  canBeClosed = False
  ua = header.get('User-Agent', None)
  cookies = header.get('Cookie', None)
  mobile = bool(ua)
  # Init driver
  if driver is None:
    canBeClosed = True
    driver: undetected.Chrome = startSelenium(uc=True, headless=True, mobile=mobile)
    if cookies:
      driver.add_cookie(header)

  # Get Page
  driver.default_get(url)
  if checkForBlock(data=driver.page_source):
    driver.get(url)

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
    preTag: Tag = soup.find(name='pre')
    data = json.loads(s=preTag.text)
    return json.dumps(obj=data, ensure_ascii=False)
  except:
    return data
