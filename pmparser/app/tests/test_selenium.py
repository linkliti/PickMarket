""" Test Selenium Module """
# pylint: disable = unused-import, invalid-name, import-error, redefined-outer-name, unused-argument
from seleniumbase import undetected
from app.selenium.seleniumCookies import getCookies
from app.selenium.selenium import startSelenium
from app.tests.test_base import logger

def test_getCookies(logger) -> None:
  """ Test getCookies """
  driver: undetected.Chrome = startSelenium(uc=True, headless=True, mobile=True)
  cookies: str = getCookies(driver)
  print(cookies)
  assert cookies
