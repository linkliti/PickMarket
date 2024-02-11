""" Base parser class for marketplace """
from http import client
import logging

from bs4 import BeautifulSoup, ResultSet, Tag

log = logging.getLogger(__name__)


class Parser():
  """ Base parser class for marketplaces """
  def __init__(self) -> None:
    self.mobileUA = "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Mobile Safari/537.36" #pylint: disable = line-too-long

  def getData(self, host: str, url: str, header: dict = None, useMobile = False) -> str:
    """ Get any data from url """
    if header is None:
      header = {}
    if useMobile:
      header["user-agent"] = self.mobileUA
    conn = client.HTTPSConnection(host=host)
    conn.request(method="GET", url=url, body='', headers=header)
    response: client.HTTPResponse = conn.getresponse()
    result: bytes = response.read()
    return result.decode(encoding="utf-8")

  def parseData(self, html: str, selector: str) -> list:
    """ Parse any data from html using CSS selector """
    soup = BeautifulSoup(markup=html, features="html.parser")
    data: ResultSet[Tag] = soup.select(selector=selector)
    return data

  def jsonQuotesEscape(self, s: str) -> str:
    """ Escape quotes in JSON string """
    return s.replace('\\"', '"')
