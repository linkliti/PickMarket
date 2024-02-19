""" Selenium Module for cookies """
import random
import logging
from typing import List
from mypy.test.helpers import retry_on_error
from seleniumbase import undetected

from app.selenium.selenium import checkForBlock, startSelenium

log = logging.getLogger(__name__)

ADULT_PAGE = 'https://www.ozon.ru/category/energeticheskie-napitki-9473/'


def cookiesToHeader(cookies: List[dict]) -> str:
  """ Convert cookies to header """
  return "; ".join(f"{cookie['name']}={cookie['value']}" for cookie in cookies)


def getCookies(driver: undetected.Chrome = None) -> str:
  """ Complete age verification and return cookies """
  log.debug("Cookies: Running Chrome")
  # Init driver
  if driver is None:
    driver: undetected.Chrome = startSelenium(uc=True, mobile=True)
  driver.default_get(ADULT_PAGE)
  if checkForBlock(data=driver.page_source):
    driver.get(ADULT_PAGE)

  # Generate birthdate
  day: int = random.randint(1, 28)
  month: int = random.randint(1, 12)
  year: int = random.randint(1975, 2005)
  date: str = f"{day:02d}.{month:02d}.{year}"

  # Use selenium
  driver.type("input[name='birthdate']", date)
  driver.js_click(".b237-a")
  driver.sleep(2)

  # Exporting cookies
  cookies: List[dict] = driver.get_cookies()
  # log.debug("Cookies: %s", cookies)
  cookiesHeader: str = cookiesToHeader(cookies=cookies)
  log.debug("Cookies: %s", cookiesHeader)
  return cookiesHeader
