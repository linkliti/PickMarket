""" Selenium Helper Functions"""
import logging

log = logging.getLogger(__name__)

BADLIST: list[str] = [
  'Just a moment', 'We need to make sure that you are not a robot.', 'Checking your browser',
  'Один момент', 'Доступ ограничен'
]

class BlockedError(Exception):
  """Custom exception for block"""

def cookiesToHeader(cookies: list[dict]) -> str:
  """ Convert cookies to header """
  return "; ".join(f"{cookie['name']}={cookie['value']}" for cookie in cookies)


def checkForBlock(data: str) -> bool:
  """ Check for block. True if blocked """
  if any(sub in data for sub in BADLIST):
    log.warning("Quick page grab failed")
    return True
  return False
