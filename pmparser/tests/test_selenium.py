""" Test Selenium Module """
# pylint: disable = unused-import, invalid-name, import-error, redefined-outer-name, unused-argument
from app.selen.selenium import startSelenium
from app.selen.seleniumCookies import getCookies
from seleniumbase import undetected

from .test_base import logger


def test_getCookies(logger) -> None:
  """ Test getCookies """
  driver: undetected.Chrome = startSelenium(uc=True, headless=True, mobile=True)
  cookies: str = getCookies(driver)
  print(cookies)
  assert cookies
