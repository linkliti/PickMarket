from app.protos import types_pb2 as _types_pb2
from google.rpc import status_pb2 as _status_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ItemsRequest(_message.Message):
    __slots__ = ("market", "pageUrl", "userQuery", "params", "numOfPages")
    MARKET_FIELD_NUMBER: _ClassVar[int]
    PAGEURL_FIELD_NUMBER: _ClassVar[int]
    USERQUERY_FIELD_NUMBER: _ClassVar[int]
    PARAMS_FIELD_NUMBER: _ClassVar[int]
    NUMOFPAGES_FIELD_NUMBER: _ClassVar[int]
    market: _types_pb2.Markets
    pageUrl: str
    userQuery: str
    params: str
    numOfPages: int
    def __init__(self, market: _Optional[_Union[_types_pb2.Markets, str]] = ..., pageUrl: _Optional[str] = ..., userQuery: _Optional[str] = ..., params: _Optional[str] = ..., numOfPages: _Optional[int] = ...) -> None: ...

class ItemResponse(_message.Message):
    __slots__ = ("item", "status")
    ITEM_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    item: Item
    status: _status_pb2.Status
    def __init__(self, item: _Optional[_Union[Item, _Mapping]] = ..., status: _Optional[_Union[_status_pb2.Status, _Mapping]] = ...) -> None: ...

class Item(_message.Message):
    __slots__ = ("name", "url", "imageUrl", "isAdult", "price", "oldPrice", "rating", "comments", "original")
    NAME_FIELD_NUMBER: _ClassVar[int]
    URL_FIELD_NUMBER: _ClassVar[int]
    IMAGEURL_FIELD_NUMBER: _ClassVar[int]
    ISADULT_FIELD_NUMBER: _ClassVar[int]
    PRICE_FIELD_NUMBER: _ClassVar[int]
    OLDPRICE_FIELD_NUMBER: _ClassVar[int]
    RATING_FIELD_NUMBER: _ClassVar[int]
    COMMENTS_FIELD_NUMBER: _ClassVar[int]
    ORIGINAL_FIELD_NUMBER: _ClassVar[int]
    name: str
    url: str
    imageUrl: str
    isAdult: bool
    price: float
    oldPrice: float
    rating: float
    comments: int
    original: bool
    def __init__(self, name: _Optional[str] = ..., url: _Optional[str] = ..., imageUrl: _Optional[str] = ..., isAdult: bool = ..., price: _Optional[float] = ..., oldPrice: _Optional[float] = ..., rating: _Optional[float] = ..., comments: _Optional[int] = ..., original: bool = ...) -> None: ...

class CharacteristicsRequest(_message.Message):
    __slots__ = ("market", "itemUrl")
    MARKET_FIELD_NUMBER: _ClassVar[int]
    ITEMURL_FIELD_NUMBER: _ClassVar[int]
    market: _types_pb2.Markets
    itemUrl: str
    def __init__(self, market: _Optional[_Union[_types_pb2.Markets, str]] = ..., itemUrl: _Optional[str] = ...) -> None: ...

class CharacteristicResponse(_message.Message):
    __slots__ = ("char", "status")
    CHAR_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    char: Characteristic
    status: _status_pb2.Status
    def __init__(self, char: _Optional[_Union[Characteristic, _Mapping]] = ..., status: _Optional[_Union[_status_pb2.Status, _Mapping]] = ...) -> None: ...

class Characteristic(_message.Message):
    __slots__ = ("key", "name", "charWeight", "maxWeight", "numVal", "listVal")
    KEY_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    CHARWEIGHT_FIELD_NUMBER: _ClassVar[int]
    MAXWEIGHT_FIELD_NUMBER: _ClassVar[int]
    NUMVAL_FIELD_NUMBER: _ClassVar[int]
    LISTVAL_FIELD_NUMBER: _ClassVar[int]
    key: str
    name: str
    charWeight: float
    maxWeight: float
    numVal: float
    listVal: _types_pb2.StringList
    def __init__(self, key: _Optional[str] = ..., name: _Optional[str] = ..., charWeight: _Optional[float] = ..., maxWeight: _Optional[float] = ..., numVal: _Optional[float] = ..., listVal: _Optional[_Union[_types_pb2.StringList, _Mapping]] = ...) -> None: ...

class FiltersRequest(_message.Message):
    __slots__ = ("market", "categoryUrl")
    MARKET_FIELD_NUMBER: _ClassVar[int]
    CATEGORYURL_FIELD_NUMBER: _ClassVar[int]
    market: _types_pb2.Markets
    categoryUrl: str
    def __init__(self, market: _Optional[_Union[_types_pb2.Markets, str]] = ..., categoryUrl: _Optional[str] = ...) -> None: ...

class FilterResponse(_message.Message):
    __slots__ = ("filter", "status")
    FILTER_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    filter: Filter
    status: _status_pb2.Status
    def __init__(self, filter: _Optional[_Union[Filter, _Mapping]] = ..., status: _Optional[_Union[_status_pb2.Status, _Mapping]] = ...) -> None: ...

class Filter(_message.Message):
    __slots__ = ("title", "key", "externalType", "internalType", "rangeFilter", "selectionFilter", "boolFilter")
    TITLE_FIELD_NUMBER: _ClassVar[int]
    KEY_FIELD_NUMBER: _ClassVar[int]
    EXTERNALTYPE_FIELD_NUMBER: _ClassVar[int]
    INTERNALTYPE_FIELD_NUMBER: _ClassVar[int]
    RANGEFILTER_FIELD_NUMBER: _ClassVar[int]
    SELECTIONFILTER_FIELD_NUMBER: _ClassVar[int]
    BOOLFILTER_FIELD_NUMBER: _ClassVar[int]
    title: str
    key: str
    externalType: str
    internalType: _types_pb2.Filters
    rangeFilter: _types_pb2.RangeFilter
    selectionFilter: _types_pb2.SelectionFilter
    boolFilter: _types_pb2.BoolFilter
    def __init__(self, title: _Optional[str] = ..., key: _Optional[str] = ..., externalType: _Optional[str] = ..., internalType: _Optional[_Union[_types_pb2.Filters, str]] = ..., rangeFilter: _Optional[_Union[_types_pb2.RangeFilter, _Mapping]] = ..., selectionFilter: _Optional[_Union[_types_pb2.SelectionFilter, _Mapping]] = ..., boolFilter: _Optional[_Union[_types_pb2.BoolFilter, _Mapping]] = ...) -> None: ...
