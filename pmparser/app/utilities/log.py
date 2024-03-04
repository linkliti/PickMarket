""" Set up logger """
import logging
import sys


def setupLogger(name, debug, filename='parser.log') -> logging.Logger:
  """ Set up logger """
  log: logging.Logger = logging.getLogger(name=name)
  if debug:
    log.setLevel(level=logging.DEBUG)
    logFile = logging.StreamHandler(stream=sys.stdout)
  else:
    log.setLevel(level=logging.INFO)
    logFile = logging.FileHandler(filename=filename)
  logFormat = logging.Formatter(
    fmt='time=%(asctime)s level=%(levelname)s pid=%(process)d name=%(name)s msg=%(message)s')
  logFile.setFormatter(fmt=logFormat)
  log.addHandler(hdlr=logFile)
  return log
