""" Set up logger """
import json
import logging
import sys
from datetime import datetime
from typing import Any

from pythonjsonlogger import jsonlogger


class CustomJsonFormatter(jsonlogger.JsonFormatter):
  """ Custom logger formatter """

  def add_fields(self, log_record: dict[str, Any], record: logging.LogRecord,
                 message_dict: dict[str, Any]) -> None:
    super().add_fields(log_record=log_record, record=record, message_dict=message_dict)
    # Remove default field
    del log_record['taskName']
    # Update fields
    log_record['time'] = datetime.fromtimestamp(timestamp=record.created).isoformat()
    log_record['level'] = record.levelname
    log_record['process'] = record.process
    log_record['name'] = record.name
    log_record['msg'] = record.getMessage()


def jsonSerializer(obj: Any, *args, **kwargs) -> str:
  """JSON serializer without spaces"""
  return json.dumps(obj=obj, separators=(',', ':'), *args, **kwargs)


def setupLogger(name, debug, filename='pmparser.log') -> None:
  """ Setup logger """
  logger: logging.Logger = logging.getLogger(name=name)
  logger.setLevel(level=logging.DEBUG if debug else logging.INFO)

  # Create formatter
  formatter = CustomJsonFormatter(fmt="%(time)s %(level)s %(process)s %(name)s %(msg)s",
                                  json_ensure_ascii=False,
                                  json_serializer=jsonSerializer)

  # Create a file handler and add it to logger
  fileHandler = logging.FileHandler(filename=filename, encoding='utf-8')
  fileHandler.setFormatter(fmt=formatter)
  logger.addHandler(hdlr=fileHandler)

  # If debug is True, add a stream handler to logger
  if debug:
    streamHandler = logging.StreamHandler(stream=sys.stdout)
    streamHandler.setFormatter(fmt=formatter)
    logger.addHandler(hdlr=streamHandler)
