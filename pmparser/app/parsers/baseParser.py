""" Base parser class for marketplace """
from http import client
import logging
import urllib.parse

from bs4 import BeautifulSoup, ResultSet, Tag
from seleniumbase import undetected

from app.selenium.seleniumFallback import getDataFallback
from app.selenium.selenium import checkForBlock

log = logging.getLogger(__name__)


class Parser():
  """ Base parser class for marketplaces """

  def __init__(self) -> None:
    self.mobileUA = "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Mobile Safari/537.36"  #pylint: disable = line-too-long

  def getData(self,
              host: str,
              url: str,
              header: dict[str, str] = None,
              params: dict[str, str] = None,
              useMobile=False,
              driver: undetected.Chrome = None) -> str:
    """ Get any data from url """
    if header is None:
      header = {}
    if params is None:
      params = {}
    if useMobile:
      header["User-Agent"] = self.mobileUA
    conn = client.HTTPSConnection(host=host)
    url = url + '/?' + urllib.parse.urlencode(query=params) if params else url
    conn.request(method="GET", url=url, body='', headers=header)
    response: client.HTTPResponse = conn.getresponse()
    result: bytes = response.read()
    data: str = result.decode(encoding="utf-8")
    if checkForBlock(data=data):
      log.warning("Triggered Cloudflare protection. Getting data via selenium")
      data = getDataFallback(url="https://" + host + url, header=header, driver=driver)
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
