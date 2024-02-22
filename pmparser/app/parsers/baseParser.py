""" Base parser class for marketplaces """
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
    self.cookie = "cf_clearance=g8bca34ZHYc705FTqaq5YfF2URcjAM7McM4FK0Elnh4-1708084673-1.0-AT+6R/u0TY4o9yBEaVbmlnCp+zYjPCLPbG1NZFBh/Dd/h07xSwP2E8BL2cX5DJrYWeZnoynMdxwYM3XYmJUD//c=; __cf_bm=SRz9b8fH.6rK_bT7zsyB0Q3NTwjlMEmNjlRfylnTuyU-1708084673-1.0-AeiDAsqW8jsJLd4qKAIUDQoPtU6UK+oYZlBkIdoxH/ufMepNFSlgpmsGkTSfufADpheopCgsahxq3q4metQeEt8=; abt_data=2393367cd561d510462a9052ff2e5f28:92686c9af86e7ceb65296b17cf8268a7744b9be176414aa6d9a78d41394000679ec8780d938336e950912a5b5e9efbecd162c3214434a4fd210b23933096685f70b5896ad61045de28df1bf76c84240bf1a488af488548b63b12cf55a7470f2c07e436f972823d3b9bb88373dd4a70e682e540176b96fffc6840c3dabba833f513bcbc198b4dbfe9d288e1d8a12469b9e731e46e83b5f181a0099264477558fcccab6a196021b2f95a01dae127a53297959f6cdb4fe41864b6b040600aece10809b081b69585765dd9e18c8660feb3dd422dcbabae565f8770ec7ba0d99ddaec7ca1719cb4c0d9b574d06dd0b8dca485940d9b672e5e71cf4681745342fad826; __Secure-ext_xcid=3afebd8b99605920a03ebb7508fcef9a; xcid=3afebd8b99605920a03ebb7508fcef9a; __Secure-refresh-token=4.0.jAeUZVG2SWK8BASi5GVySg.8.Acq-vqwktmoPzfcae2jNPgtA6d9QejwxoScoCzi98IqApcv_kQI_TcIacS5zbgSSwQ..20240216135753.c7JVSMqtsthpNcZfjTj8Jb_H1xfh-HGPlcZ58_TxHeo; __Secure-access-token=4.0.jAeUZVG2SWK8BASi5GVySg.8.Acq-vqwktmoPzfcae2jNPgtA6d9QejwxoScoCzi98IqApcv_kQI_TcIacS5zbgSSwQ..20240216135753.MUyDqrBkAfuap9z5JUDrD0Xy-mN951-yi4eOopbZWjU; adult_user_birthdate=2004-12-02; is_adult_confirmed=true; __Secure-user-id=0; __Secure-ab-group=8"  #pylint: disable = line-too-long

  def getData(self,
              host: str,
              url: str,
              header: dict[str, str] | None = None,
              params: dict[str, str] | None = None,
              useMobile=False,
              driver: undetected.Chrome | None = None) -> str:
    """ Get any data from url """
    if header is None:
      header = {}
    if params is None:
      params = {}
    if useMobile:
      header["User-Agent"] = self.mobileUA
    if self.cookie:
      header["Cookie"] = self.cookie
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
