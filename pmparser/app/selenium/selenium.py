""" Selenium Module """
import logging
from seleniumbase import Driver
from seleniumbase import undetected

log = logging.getLogger(__name__)

BADLIST: list[str] = [
  'Just a moment', 'We need to make sure that you are not a robot.', 'Checking your browser'
]


def startSelenium(*args, **kwargs) -> undetected.Chrome:
  """ Start Selenium """
  driver: undetected.Chrome = Driver(*args, **kwargs)
  return driver


def checkForBlock(data: str) -> bool:
  """ Check for block. True if blocked """
  if any(sub in data for sub in BADLIST):
    log.warning("Quick page grab failed")
    return True
  return False
