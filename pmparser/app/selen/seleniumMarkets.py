""" Selenium Module for marketplaces """
import json
import random
import logging
from bs4 import BeautifulSoup, NavigableString, Tag

from app.selen.seleniumWorker import SeleniumWorker

log = logging.getLogger(__name__)

OZON_ADULT_PAGE = 'https://www.ozon.ru/category/energeticheskie-napitki-9473/'


class SeleniumWorkerMarkets(SeleniumWorker):
  """ Selenium worker for marketplaces """

  def __init__(self) -> None:
    super().__init__()

  def setOzonAdultCookies(self) -> None:
    """ Complete age verification and return cookies """
    self.getPageSource(url=OZON_ADULT_PAGE)
    # Generate birthdate
    day: int = random.randint(1, 28)
    month: int = random.randint(1, 12)
    year: int = random.randint(1975, 2005)
    date: str = f"{day:02d}.{month:02d}.{year}"
    # Use selenium
    self.driver.type("input[name='birthdate']", date)
    self.driver.js_click("button > div:nth-child(2)")
    self.driver.sleep(1)

  def getOzonStringJSON(self, url: str) -> str:
    """Get JSON from OZON in 'pre' tag"""
    data: str = self.getPageSource(url=url)
    try:
      soup = BeautifulSoup(markup=data, features='html.parser')
      preTag: Tag | NavigableString | None = soup.find(name='pre')
      if not preTag:
        raise Exception("Failed to parse webpage")
      jsonData: dict = json.loads(s=preTag.text)
      return json.dumps(obj=jsonData, ensure_ascii=False)
    except:
      return data
