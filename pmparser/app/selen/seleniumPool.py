""" Selenium Browser Pool """
import queue
from app.selen.seleniumMarkets import SeleniumWorkerMarkets

browserQueue = queue.Queue()


def getBrowserToGetPageSource(url) -> str:
  """ Take browser from queue"""
  browserWorker: SeleniumWorkerMarkets = browserQueue.get()
  try:
    data: str = browserWorker.getPageSource(url)
    return data
  finally:
    browserQueue.put(item=browserWorker)


def getBrowserToGetOzonJson(url) -> str:
  """ Take browser from queue"""
  browserWorker: SeleniumWorkerMarkets = browserQueue.get()
  try:
    data: str = browserWorker.getOzonStringJSON(url)
    return data
  finally:
    browserQueue.put(item=browserWorker)
