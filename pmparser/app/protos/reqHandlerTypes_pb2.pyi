from app.protos import items_pb2 as _items_pb2
from app.protos import types_pb2 as _types_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class PrefType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    NUM: _ClassVar[PrefType]
    LIST: _ClassVar[PrefType]
NUM: PrefType
LIST: PrefType

class UserPrefs(_message.Message):
    __slots__ = ("key", "type", "numValue", "listValue")
    KEY_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    NUMVALUE_FIELD_NUMBER: _ClassVar[int]
    LISTVALUE_FIELD_NUMBER: _ClassVar[int]
    key: str
    type: PrefType
    numValue: float
    listValue: _types_pb2.StringList
    def __init__(self, key: _Optional[str] = ..., type: _Optional[_Union[PrefType, str]] = ..., numValue: _Optional[float] = ..., listValue: _Optional[_Union[_types_pb2.StringList, _Mapping]] = ...) -> None: ...

class ItemExtended(_message.Message):
    __slots__ = ("item", "chars", "totalWeight")
    ITEM_FIELD_NUMBER: _ClassVar[int]
    CHARS_FIELD_NUMBER: _ClassVar[int]
    TOTALWEIGHT_FIELD_NUMBER: _ClassVar[int]
    item: _items_pb2.Item
    chars: _containers.RepeatedCompositeFieldContainer[_items_pb2.Characteristic]
    totalWeight: int
    def __init__(self, item: _Optional[_Union[_items_pb2.Item, _Mapping]] = ..., chars: _Optional[_Iterable[_Union[_items_pb2.Characteristic, _Mapping]]] = ..., totalWeight: _Optional[int] = ...) -> None: ...
