""" Set up logger """
import logging
import sys
from pythonjsonlogger import jsonlogger


def setupLogger(name, debug, filename='pmparser.log') -> None:
  """ Setup logger """
  logger = logging.getLogger(name)
  logger.setLevel(logging.DEBUG if debug else logging.INFO)

  # Create formatter
  formatter = jsonlogger.JsonFormatter(
    fmt='%(asctime)s %(levelname)s %(process)s %(name)s %(message)s', datefmt='%Y-%m-%dT%H:%M:%S')

  # Create a file handler and add it to logger
  fileHandler = logging.FileHandler(filename)
  fileHandler.setFormatter(formatter)
  logger.addHandler(fileHandler)

  # If debug is True, add a stream handler to logger
  if debug:
    streamHandler = logging.StreamHandler(sys.stdout)
    streamHandler.setFormatter(formatter)
    logger.addHandler(streamHandler)
