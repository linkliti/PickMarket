""" Set up logger """
import logging


def setupLogger(name, debug) -> logging.Logger:
  """ Set up logger """
  log = logging.getLogger(name)
  if debug:
    log.setLevel(logging.DEBUG)
  else:
    log.setLevel(logging.INFO)
  logFormat = logging.Formatter('[%(asctime)s][%(levelname)s]: [%(name)s] %(message)s')
  logFile = logging.FileHandler('parser.log')
  logFile.setFormatter(logFormat)
  log.addHandler(logFile)
  return log
