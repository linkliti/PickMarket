from app.protos import types_pb2 as _types_pb2
from google.rpc import status_pb2 as _status_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class RootCategoriesRequest(_message.Message):
    __slots__ = ("market",)
    MARKET_FIELD_NUMBER: _ClassVar[int]
    market: _types_pb2.Markets
    def __init__(self, market: _Optional[_Union[_types_pb2.Markets, str]] = ...) -> None: ...

class SubCategoriesRequest(_message.Message):
    __slots__ = ("market", "categoryUrl")
    MARKET_FIELD_NUMBER: _ClassVar[int]
    CATEGORYURL_FIELD_NUMBER: _ClassVar[int]
    market: _types_pb2.Markets
    categoryUrl: str
    def __init__(self, market: _Optional[_Union[_types_pb2.Markets, str]] = ..., categoryUrl: _Optional[str] = ...) -> None: ...

class CategoryResponse(_message.Message):
    __slots__ = ("category", "status")
    CATEGORY_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    category: Category
    status: _status_pb2.Status
    def __init__(self, category: _Optional[_Union[Category, _Mapping]] = ..., status: _Optional[_Union[_status_pb2.Status, _Mapping]] = ...) -> None: ...

class Category(_message.Message):
    __slots__ = ("title", "url", "parentUrl")
    TITLE_FIELD_NUMBER: _ClassVar[int]
    URL_FIELD_NUMBER: _ClassVar[int]
    PARENTURL_FIELD_NUMBER: _ClassVar[int]
    title: str
    url: str
    parentUrl: str
    def __init__(self, title: _Optional[str] = ..., url: _Optional[str] = ..., parentUrl: _Optional[str] = ...) -> None: ...
