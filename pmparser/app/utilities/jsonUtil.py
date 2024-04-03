""" JSON Utilities """
import logging
import json

log = logging.getLogger(__name__)

def toJson(s: str) -> dict:
  """ Convert string to JSON """
  try:
    return json.loads(s)
  except Exception as e:
    log.error("failed to convert to JSON", extra={"string": s})
    raise e
