from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Markets(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    OZON: _ClassVar[Markets]

class Filters(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    RANGE: _ClassVar[Filters]
    SELECTION: _ClassVar[Filters]
    SELECTIONRADIO: _ClassVar[Filters]
    BOOL: _ClassVar[Filters]
OZON: Markets
RANGE: Filters
SELECTION: Filters
SELECTIONRADIO: Filters
BOOL: Filters

class RangeFilter(_message.Message):
    __slots__ = ("min", "max")
    MIN_FIELD_NUMBER: _ClassVar[int]
    MAX_FIELD_NUMBER: _ClassVar[int]
    min: float
    max: float
    def __init__(self, min: _Optional[float] = ..., max: _Optional[float] = ...) -> None: ...

class SelectionFilterItem(_message.Message):
    __slots__ = ("text", "value")
    TEXT_FIELD_NUMBER: _ClassVar[int]
    VALUE_FIELD_NUMBER: _ClassVar[int]
    text: str
    value: str
    def __init__(self, text: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...

class SelectionFilter(_message.Message):
    __slots__ = ("isRadio", "items")
    ISRADIO_FIELD_NUMBER: _ClassVar[int]
    ITEMS_FIELD_NUMBER: _ClassVar[int]
    isRadio: bool
    items: _containers.RepeatedCompositeFieldContainer[SelectionFilterItem]
    def __init__(self, isRadio: bool = ..., items: _Optional[_Iterable[_Union[SelectionFilterItem, _Mapping]]] = ...) -> None: ...

class BoolFilter(_message.Message):
    __slots__ = ("value",)
    VALUE_FIELD_NUMBER: _ClassVar[int]
    value: str
    def __init__(self, value: _Optional[str] = ...) -> None: ...

class StringList(_message.Message):
    __slots__ = ("values",)
    VALUES_FIELD_NUMBER: _ClassVar[int]
    values: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, values: _Optional[_Iterable[str]] = ...) -> None: ...
