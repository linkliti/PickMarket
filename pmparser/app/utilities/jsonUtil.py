""" JSON Utilities """
import json
import logging

import google.protobuf.text_format as text_format
from google.protobuf import message

log = logging.getLogger(__name__)


def toJson(s: str) -> dict:
  """ Convert string to JSON """
  try:
    return json.loads(s)
  except Exception as e:
    log.error("failed to convert to JSON", extra={"string": s})
    raise e


def msgToStr(msg: message.Message) -> str:
  """ Convert protobuf message to string """
  return text_format.MessageToString(message=msg, as_utf8=True, as_one_line=True)
