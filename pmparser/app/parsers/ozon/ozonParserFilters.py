""" Ozon Parser Module for categories """
from typing import Generator
import logging

import json

from app.parsers.ozon.ozonParser import OzonParser
from app.parsers.baseDataclass import BaseFiltersDataclass, CheckboxesFilterItem, RangeFilter

log = logging.getLogger(__name__)


class OzonParserFilters(OzonParser):
  """ Ozon Parser Module for items """

  def getFilters(self, pageUrl: str) -> Generator[BaseFiltersDataclass, None, None]:
    """ Get filters for category """
    log.info('Getting filters: %s', pageUrl)
    jString: str = self.getData(host=self.host,
                                url=self.api + "/modal/filters" + pageUrl,
                                useMobile=True)

    log.info('Converting to JSON: %s', pageUrl)
    j: dict = json.loads(jString)
    j = self.getEmbededJson(j=j["widgetStates"], keyName="filters")

    log.info('Parsing filters: %s', pageUrl)
