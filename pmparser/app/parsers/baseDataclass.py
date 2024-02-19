""" Base Dataclasses """
import logging
from dataclasses import dataclass
from typing import List
from dataclasses_json import dataclass_json

log: logging.Logger = logging.getLogger(__name__)


@dataclass_json
@dataclass
class BaseCategoryDataclass:
  """Dataclass for storing category data"""
  title: str
  url: str
  parent: str | None = None


@dataclass_json
@dataclass
class BaseItemDataclass:
  """ Dataclass for storing item data"""
  name: str
  url: str
  imageUrl: str
  price: int
  isAdult: bool
  stars: float | None = None
  comments: int | None = None
  oldPrice: int | None = None


@dataclass_json
@dataclass
class BaseItemCharsDataclass:
  """Dataclass for storing item characteristics"""
  key: str
  name: str
  value: str | int | float | List[str]

@dataclass
class RangeFilter:
  """ Range filter """
  isRadio: bool
  min: int
  max: int

@dataclass
class CheckboxesFilterItem:
  """ Checkboxes filter items """
  key: str
  title: str

@dataclass_json
@dataclass
class BaseFiltersDataclass:
  """Dataclass for storing filters"""
  key: str
  title: str
  type: str
  data: RangeFilter | List[CheckboxesFilterItem]
