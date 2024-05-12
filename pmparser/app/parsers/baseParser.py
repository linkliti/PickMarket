""" Base parser class for marketplaces """
import logging
from urllib.parse import urlencode

from app.selen.seleniumPool import getBrowserToGetPageSource
from bs4 import BeautifulSoup, ResultSet, Tag

log = logging.getLogger(__name__)


class Parser():
  """ Base parser class for marketplaces """

  def __init__(self) -> None:
    self.mobileUA = "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Mobile Safari/537.36"  #pylint: disable = line-too-long

  def getData(self, host: str, url: str, params: dict[str, str] | None = None) -> str:
    """ Get any data from url """
    if params:
      url += "?" + urlencode(params)
    data = self.getDataViaSelenium(url="https://" + host + url)
    return data


  def getDataViaSelenium(self, url: str) -> str:
    """ Get data via selenium """
    data = getBrowserToGetPageSource(url=url)
    return data

  def htmlStringToTags(self, html: str, selector: str) -> ResultSet[Tag]:
    """ Parse any data from html using CSS selector """
    soup = BeautifulSoup(markup=html, features="html.parser")
    tags: ResultSet[Tag] = soup.select(selector=selector)
    return tags

  def jsonQuotesEscape(self, s: str) -> str:
    """ Escape quotes in JSON string """
    return s.replace('\\"', '"')

  def priceClean(self, price: str) -> int:
    """Remove non-digit characters and convert to integer """
    return int(''.join(filter(str.isdigit, price)))
