""" Selenium Module for marketplaces """
import json
import logging
import random

from app.selen.seleniumWorker import SeleniumWorker
from app.utilities.jsonUtil import toJson
from bs4 import BeautifulSoup, NavigableString, Tag

log = logging.getLogger(__name__)

OZON_ADULT_PAGE = 'https://www.ozon.ru/category/energeticheskie-napitki-9473/'
OZON_MAIN_PAGE = 'https://www.ozon.ru/'


class SeleniumWorkerMarkets(SeleniumWorker):
  """ Selenium worker for marketplaces """

  def __init__(self) -> None:  #pylint: disable=useless-parent-delegation
    day: int = random.randint(1, 28)
    month: int = random.randint(1, 12)
    year: int = random.randint(1975, 2005)
    self.date: str = f"{day:02d}.{month:02d}.{year}"
    super().__init__()

  def resetCookies(self) -> None:
    """ Reset cookies """
    self.driver.delete_all_cookies()
    self.driver.add_cookie(cookie_dict=\
      {"name": "adult_user_birthdate", "value": self.date})

  def setOzonAdultCookies(self) -> None:
    """ Complete age verification and return cookies """
    self.getPageSource(url=OZON_MAIN_PAGE)
    self.resetCookies()

    # self.getPageSource(url=OZON_ADULT_PAGE)
    # # Use selenium
    # self.driver.type("input[name='birthdate']", self.date)
    # self.driver.js_click("button > div:nth-child(2)")
    # self.driver.sleep(1)

  def getOzonStringJSON(self, url: str) -> str:
    """Get JSON from OZON in 'pre' tag"""
    data: str = self.getPageSource(url=url)
    try:
      soup = BeautifulSoup(markup=data, features='html.parser')
      preTag: Tag | NavigableString | None = soup.find(name='pre')
      if not preTag:
        raise Exception("Failed to parse webpage")
      jsonData: dict = toJson(s=preTag.text)
      return json.dumps(obj=jsonData, ensure_ascii=False)
    except:
      return data
