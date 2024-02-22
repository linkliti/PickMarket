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
  """Range filter"""
  min: int
  max: int


@dataclass
class SelectionFilterItem:
  """Selection filter item"""
  text: str
  value: str


@dataclass
class SelectionFilter:
  """Selection filter"""
  isRadio: bool
  items: List[SelectionFilterItem]


@dataclass
class BoolFilter:
  """Boolean filter"""
  value: str = 't'


@dataclass_json
@dataclass
class BaseFiltersDataclass:
  """Dataclass for storing filters"""
  title: str
  key: str
  externalType: str
  internalType: str
  data: RangeFilter | SelectionFilter | BoolFilter
