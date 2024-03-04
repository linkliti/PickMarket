""" Selenium Worker Module """
import logging
import os
from app.selen.seleniumHelpers import checkForBlock, cookiesToHeader
from seleniumbase import Driver, undetected

log = logging.getLogger(__name__)

DEBUG = bool(os.environ.get('DEBUG', False))


class SeleniumWorker():
  """Base selenium worker"""

  def __init__(self, *args, **kwargs) -> None:
    presetArgs: dict[str, bool] = {
      'uc': True,
      'mobile': True,
      'block_images': True,
      'do_not_track': True,
      'dark_mode': True,
      'uc_subprocess': True,
    }
    if DEBUG:
      presetArgs['headed'] = True
    else:
      presetArgs['headless'] = True
    self.driver: undetected.Chrome = Driver(*args, **presetArgs, **kwargs)

  def getPageSource(self, url: str) -> str:
    """GET page source from url using browser"""
    log.debug("Getting page source: %s", url)
    # Get Page
    self.driver.default_get(url)
    if checkForBlock(data=self.driver.page_source):
      self.driver.get(url)
    # Check for block
    data: str = self.driver.page_source
    if checkForBlock(data=data):
      log.error("Failed to bypass block: %s", data)
      raise Exception("Failed to bypass block")
    return data

  def exportCookiesToStr(self) -> str:
    """ Export cookies as string """
    cookies: list[dict] = self.driver.get_cookies()
    # log.debug("Cookies: %s", cookies)
    cookiesHeader: str = cookiesToHeader(cookies=cookies)
    log.debug("Cookies: %s", cookiesHeader)
    return cookiesHeader
