""" Server Module """
import logging
import json

from app.parsers.ozonParser import OzonParser

log = logging.getLogger(__name__)


def server() -> None:
  """ Server """
  c: dict[str, list] = OzonParser().getCategories()
  with open(file='categories.json', mode='w', encoding="utf-8") as f:
    json.dump(c, f, ensure_ascii=False)
