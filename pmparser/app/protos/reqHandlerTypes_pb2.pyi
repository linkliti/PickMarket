from app.protos import items_pb2 as _items_pb2
from app.protos import types_pb2 as _types_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ItemsRequestWithPrefs(_message.Message):
    __slots__ = ("request", "prefs")
    class PrefsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: UserPref
        def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[UserPref, _Mapping]] = ...) -> None: ...
    REQUEST_FIELD_NUMBER: _ClassVar[int]
    PREFS_FIELD_NUMBER: _ClassVar[int]
    request: _items_pb2.ItemsRequest
    prefs: _containers.MessageMap[str, UserPref]
    def __init__(self, request: _Optional[_Union[_items_pb2.ItemsRequest, _Mapping]] = ..., prefs: _Optional[_Mapping[str, UserPref]] = ...) -> None: ...

class UserPref(_message.Message):
    __slots__ = ("priority", "numVal", "listVal")
    PRIORITY_FIELD_NUMBER: _ClassVar[int]
    NUMVAL_FIELD_NUMBER: _ClassVar[int]
    LISTVAL_FIELD_NUMBER: _ClassVar[int]
    priority: int
    numVal: float
    listVal: _types_pb2.StringList
    def __init__(self, priority: _Optional[int] = ..., numVal: _Optional[float] = ..., listVal: _Optional[_Union[_types_pb2.StringList, _Mapping]] = ...) -> None: ...

class ItemExtended(_message.Message):
    __slots__ = ("item", "similar", "totalWeight", "chars")
    ITEM_FIELD_NUMBER: _ClassVar[int]
    SIMILAR_FIELD_NUMBER: _ClassVar[int]
    TOTALWEIGHT_FIELD_NUMBER: _ClassVar[int]
    CHARS_FIELD_NUMBER: _ClassVar[int]
    item: _items_pb2.Item
    similar: _containers.RepeatedCompositeFieldContainer[_items_pb2.Item]
    totalWeight: float
    chars: _containers.RepeatedCompositeFieldContainer[_items_pb2.Characteristic]
    def __init__(self, item: _Optional[_Union[_items_pb2.Item, _Mapping]] = ..., similar: _Optional[_Iterable[_Union[_items_pb2.Item, _Mapping]]] = ..., totalWeight: _Optional[float] = ..., chars: _Optional[_Iterable[_Union[_items_pb2.Characteristic, _Mapping]]] = ...) -> None: ...
