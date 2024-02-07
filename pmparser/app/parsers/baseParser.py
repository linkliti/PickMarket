""" Base parser class for marketplace """
from http import client
import logging

from bs4 import BeautifulSoup, ResultSet, Tag

log = logging.getLogger(__name__)


class Parser():
  """ Base parser class for marketplaces """

  def getData(self, host: str, url: str, header: dict = None) -> str:
    """ Get any data from url """
    if header is None:
      header = {}
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
