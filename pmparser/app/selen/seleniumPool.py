""" Selenium Browser Pool """
import queue
from app.selen.seleniumMarkets import SeleniumWorkerMarkets
from app.selen.seleniumHelpers import BlockedError

browserQueue = queue.Queue()


def getBrowserToGetPageSource(url) -> str:
  """ Take browser from queue"""
  browserWorker: SeleniumWorkerMarkets = browserQueue.get()
  try:
    data: str = browserWorker.getPageSource(url)
    return data
  # except BlockedError:
  #   browserWorker.setOzonAdultCookies()
  #   return getBrowserToGetOzonJson(str)
  finally:
    browserQueue.put(item=browserWorker)


def getBrowserToGetOzonJson(url) -> str:
  """ Take browser from queue"""
  browserWorker: SeleniumWorkerMarkets = browserQueue.get()
  try:
    data: str = browserWorker.getOzonStringJSON(url)
    browserWorker.resetCookies()
    return data
  except BlockedError:
    browserWorker.setOzonAdultCookies()
    # Last attempt
    return browserWorker.getOzonStringJSON(url)
  finally:
    browserQueue.put(item=browserWorker)
