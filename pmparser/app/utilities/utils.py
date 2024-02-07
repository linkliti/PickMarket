import os


def getEnv(var_name, default=None):
  return os.environ.get(var_name, default)
