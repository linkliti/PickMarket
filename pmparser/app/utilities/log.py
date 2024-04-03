""" Set up logger """
import logging
import sys
from pythonjsonlogger import jsonlogger


def setupLogger(name, debug, filename='pmparser.log') -> None:
  """ Setup logger """
  logger: logging.Logger = logging.getLogger(name=name)
  logger.setLevel(level=logging.DEBUG if debug else logging.INFO)

  # Create formatter
  formatter = jsonlogger.JsonFormatter(
    fmt="%(asctime)s %(levelname)s %(process)s %(name)s %(message)s", datefmt='%Y-%m-%dT%H:%M:%S')

  # Create a file handler and add it to logger
  fileHandler = logging.FileHandler(filename=filename, encoding='utf-8')
  fileHandler.setFormatter(fmt=formatter)
  logger.addHandler(hdlr=fileHandler)

  # If debug is True, add a stream handler to logger
  if debug:
    streamHandler = logging.StreamHandler(stream=sys.stdout)
    streamHandler.setFormatter(fmt=formatter)
    logger.addHandler(hdlr=streamHandler)
