// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v4.25.1
// source: app/protos/reqHandlerTypes.proto

package parser

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ItemsRequestWithPrefs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *ItemsRequest        `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Prefs   map[string]*UserPref `protobuf:"bytes,2,rep,name=prefs,proto3" json:"prefs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ItemsRequestWithPrefs) Reset() {
	*x = ItemsRequestWithPrefs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_protos_reqHandlerTypes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemsRequestWithPrefs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemsRequestWithPrefs) ProtoMessage() {}

func (x *ItemsRequestWithPrefs) ProtoReflect() protoreflect.Message {
	mi := &file_app_protos_reqHandlerTypes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemsRequestWithPrefs.ProtoReflect.Descriptor instead.
func (*ItemsRequestWithPrefs) Descriptor() ([]byte, []int) {
	return file_app_protos_reqHandlerTypes_proto_rawDescGZIP(), []int{0}
}

func (x *ItemsRequestWithPrefs) GetRequest() *ItemsRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *ItemsRequestWithPrefs) GetPrefs() map[string]*UserPref {
	if x != nil {
		return x.Prefs
	}
	return nil
}

type UserPref struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Priority int32 `protobuf:"varint,1,opt,name=priority,proto3" json:"priority,omitempty"`
	// Types that are assignable to Value:
	//	*UserPref_NumVal
	//	*UserPref_ListVal
	Value isUserPref_Value `protobuf_oneof:"value"`
}

func (x *UserPref) Reset() {
	*x = UserPref{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_protos_reqHandlerTypes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPref) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPref) ProtoMessage() {}

func (x *UserPref) ProtoReflect() protoreflect.Message {
	mi := &file_app_protos_reqHandlerTypes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPref.ProtoReflect.Descriptor instead.
func (*UserPref) Descriptor() ([]byte, []int) {
	return file_app_protos_reqHandlerTypes_proto_rawDescGZIP(), []int{1}
}

func (x *UserPref) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (m *UserPref) GetValue() isUserPref_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *UserPref) GetNumVal() float64 {
	if x, ok := x.GetValue().(*UserPref_NumVal); ok {
		return x.NumVal
	}
	return 0
}

func (x *UserPref) GetListVal() *StringList {
	if x, ok := x.GetValue().(*UserPref_ListVal); ok {
		return x.ListVal
	}
	return nil
}

type isUserPref_Value interface {
	isUserPref_Value()
}

type UserPref_NumVal struct {
	NumVal float64 `protobuf:"fixed64,2,opt,name=numVal,proto3,oneof"`
}

type UserPref_ListVal struct {
	ListVal *StringList `protobuf:"bytes,3,opt,name=listVal,proto3,oneof"`
}

func (*UserPref_NumVal) isUserPref_Value() {}

func (*UserPref_ListVal) isUserPref_Value() {}

type ItemExtended struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item        *Item             `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Similar     []*Item           `protobuf:"bytes,2,rep,name=similar,proto3" json:"similar,omitempty"`
	TotalWeight float64           `protobuf:"fixed64,3,opt,name=totalWeight,proto3" json:"totalWeight,omitempty"`
	Chars       []*Characteristic `protobuf:"bytes,4,rep,name=chars,proto3" json:"chars,omitempty"`
}

func (x *ItemExtended) Reset() {
	*x = ItemExtended{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_protos_reqHandlerTypes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemExtended) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemExtended) ProtoMessage() {}

func (x *ItemExtended) ProtoReflect() protoreflect.Message {
	mi := &file_app_protos_reqHandlerTypes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemExtended.ProtoReflect.Descriptor instead.
func (*ItemExtended) Descriptor() ([]byte, []int) {
	return file_app_protos_reqHandlerTypes_proto_rawDescGZIP(), []int{2}
}

func (x *ItemExtended) GetItem() *Item {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *ItemExtended) GetSimilar() []*Item {
	if x != nil {
		return x.Similar
	}
	return nil
}

func (x *ItemExtended) GetTotalWeight() float64 {
	if x != nil {
		return x.TotalWeight
	}
	return 0
}

func (x *ItemExtended) GetChars() []*Characteristic {
	if x != nil {
		return x.Chars
	}
	return nil
}

var File_app_protos_reqHandlerTypes_proto protoreflect.FileDescriptor

var file_app_protos_reqHandlerTypes_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x72, 0x65, 0x71,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x16,
	0x61, 0x70, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf,
	0x01, 0x0a, 0x15, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x57,
	0x69, 0x74, 0x68, 0x50, 0x72, 0x65, 0x66, 0x73, 0x12, 0x32, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x05,
	0x70, 0x72, 0x65, 0x66, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x50, 0x72, 0x65, 0x66, 0x73, 0x2e, 0x50,
	0x72, 0x65, 0x66, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x70, 0x72, 0x65, 0x66, 0x73,
	0x1a, 0x4e, 0x0a, 0x0a, 0x50, 0x72, 0x65, 0x66, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x72, 0x65, 0x66, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x7d, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x56,
	0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x56,
	0x61, 0x6c, 0x12, 0x32, 0x0a, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x07, 0x6c,
	0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0xb4, 0x01, 0x0a, 0x0c, 0x49, 0x74, 0x65, 0x6d, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64,
	0x12, 0x24, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x2a, 0x0a, 0x07, 0x73, 0x69, 0x6d, 0x69, 0x6c, 0x61,
	0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x07, 0x73, 0x69, 0x6d, 0x69, 0x6c,
	0x61, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x57, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x57, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x72, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x69, 0x73, 0x74, 0x69, 0x63, 0x52,
	0x05, 0x63, 0x68, 0x61, 0x72, 0x73, 0x42, 0x0f, 0x5a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_protos_reqHandlerTypes_proto_rawDescOnce sync.Once
	file_app_protos_reqHandlerTypes_proto_rawDescData = file_app_protos_reqHandlerTypes_proto_rawDesc
)

func file_app_protos_reqHandlerTypes_proto_rawDescGZIP() []byte {
	file_app_protos_reqHandlerTypes_proto_rawDescOnce.Do(func() {
		file_app_protos_reqHandlerTypes_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_protos_reqHandlerTypes_proto_rawDescData)
	})
	return file_app_protos_reqHandlerTypes_proto_rawDescData
}

var file_app_protos_reqHandlerTypes_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_app_protos_reqHandlerTypes_proto_goTypes = []interface{}{
	(*ItemsRequestWithPrefs)(nil), // 0: app.protos.ItemsRequestWithPrefs
	(*UserPref)(nil),              // 1: app.protos.UserPref
	(*ItemExtended)(nil),          // 2: app.protos.ItemExtended
	nil,                           // 3: app.protos.ItemsRequestWithPrefs.PrefsEntry
	(*ItemsRequest)(nil),          // 4: app.protos.ItemsRequest
	(*StringList)(nil),            // 5: app.protos.StringList
	(*Item)(nil),                  // 6: app.protos.Item
	(*Characteristic)(nil),        // 7: app.protos.Characteristic
}
var file_app_protos_reqHandlerTypes_proto_depIdxs = []int32{
	4, // 0: app.protos.ItemsRequestWithPrefs.request:type_name -> app.protos.ItemsRequest
	3, // 1: app.protos.ItemsRequestWithPrefs.prefs:type_name -> app.protos.ItemsRequestWithPrefs.PrefsEntry
	5, // 2: app.protos.UserPref.listVal:type_name -> app.protos.StringList
	6, // 3: app.protos.ItemExtended.item:type_name -> app.protos.Item
	6, // 4: app.protos.ItemExtended.similar:type_name -> app.protos.Item
	7, // 5: app.protos.ItemExtended.chars:type_name -> app.protos.Characteristic
	1, // 6: app.protos.ItemsRequestWithPrefs.PrefsEntry.value:type_name -> app.protos.UserPref
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_app_protos_reqHandlerTypes_proto_init() }
func file_app_protos_reqHandlerTypes_proto_init() {
	if File_app_protos_reqHandlerTypes_proto != nil {
		return
	}
	file_app_protos_items_proto_init()
	file_app_protos_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_app_protos_reqHandlerTypes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemsRequestWithPrefs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_app_protos_reqHandlerTypes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPref); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_app_protos_reqHandlerTypes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemExtended); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_app_protos_reqHandlerTypes_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*UserPref_NumVal)(nil),
		(*UserPref_ListVal)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_app_protos_reqHandlerTypes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_protos_reqHandlerTypes_proto_goTypes,
		DependencyIndexes: file_app_protos_reqHandlerTypes_proto_depIdxs,
		MessageInfos:      file_app_protos_reqHandlerTypes_proto_msgTypes,
	}.Build()
	File_app_protos_reqHandlerTypes_proto = out.File
	file_app_protos_reqHandlerTypes_proto_rawDesc = nil
	file_app_protos_reqHandlerTypes_proto_goTypes = nil
	file_app_protos_reqHandlerTypes_proto_depIdxs = nil
}
