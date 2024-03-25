""" Set up logger """
import logging
import sys


def setupLogger(name, debug, filename='parser.log') -> logging.Logger:
  """ Set up logger """
  log: logging.Logger = logging.getLogger(name=name)
  log.setLevel(level=logging.DEBUG if debug else logging.INFO)

  logFormat = logging.Formatter(
    fmt='time=%(asctime)s level=%(levelname)s pid=%(process)d name="%(name)s" msg="%(message)s"',
    datefmt="%Y-%m-%dT%H:%M:%S%z")

  if debug:
    stdoutHandler = logging.StreamHandler(stream=sys.stdout)
    stdoutHandler.setFormatter(fmt=logFormat)
    log.addHandler(hdlr=stdoutHandler)

  fileHandler = logging.FileHandler(filename=filename)
  fileHandler.setFormatter(fmt=logFormat)
  log.addHandler(hdlr=fileHandler)

  return log
