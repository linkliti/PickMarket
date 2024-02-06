from bs4 import BeautifulSoup
from seleniumbase import Driver
from seleniumbase import undetected
from selenium.webdriver.common.keys import Keys
from selenium.webdriver import DesiredCapabilities
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
import logging
import sys


def get_categories(driver: undetected.Chrome, url: str, selector: str) -> list:
  logging.debug("getting categories from:" + url)
  driver.default_get(url)
  html = driver.page_source
  while 'Please, enable JavaScript to continue' not in html:
    logging.debug("anti-bot triggered, refreshing...")
    driver.get(url)
    html = driver.page_source
  soup = BeautifulSoup(html, 'html.parser')
  categories = soup.select(selector)
  subcategories = []
  for category in categories:
    subcategories += get_categories(
      driver, "https://www.ozon.ru" + category['href'],
      '[data-widget="filtersDesktop"] div[style="margin-left:16px;"] a[href*="/category"]'
    )
  return [[category.text, "https://www.ozon.ru" + category['href']] for category in categories + subcategories]



def server():
  logging.basicConfig(stream=sys.stdout, level=logging.DEBUG)
  # chrome
  caps = DesiredCapabilities().CHROME
  caps["pageLoadStrategy"] = "eager"
  driver: undetected.Chrome = Driver(uc=True, headless=True, mobile=False, cap_string=caps)
  # driver.get("https://www.ozon.ru")
  logging.debug("chrome started")
  try:
    print(
      get_categories(driver, 'https://www.ozon.ru/category/',
                     'a[href*="/category"]'))
    driver.sleep(10)
  finally:
    driver.quit()
