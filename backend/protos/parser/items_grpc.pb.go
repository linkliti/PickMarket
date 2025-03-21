// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: app/protos/items.proto

package parser

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ItemParser_GetItems_FullMethodName               = "/app.protos.ItemParser/GetItems"
	ItemParser_GetItemCharacteristics_FullMethodName = "/app.protos.ItemParser/GetItemCharacteristics"
	ItemParser_GetCategoryFilters_FullMethodName     = "/app.protos.ItemParser/GetCategoryFilters"
)

// ItemParserClient is the client API for ItemParser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ItemParserClient interface {
	GetItems(ctx context.Context, in *ItemsRequest, opts ...grpc.CallOption) (ItemParser_GetItemsClient, error)
	GetItemCharacteristics(ctx context.Context, in *CharacteristicsRequest, opts ...grpc.CallOption) (ItemParser_GetItemCharacteristicsClient, error)
	GetCategoryFilters(ctx context.Context, in *FiltersRequest, opts ...grpc.CallOption) (ItemParser_GetCategoryFiltersClient, error)
}

type itemParserClient struct {
	cc grpc.ClientConnInterface
}

func NewItemParserClient(cc grpc.ClientConnInterface) ItemParserClient {
	return &itemParserClient{cc}
}

func (c *itemParserClient) GetItems(ctx context.Context, in *ItemsRequest, opts ...grpc.CallOption) (ItemParser_GetItemsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ItemParser_ServiceDesc.Streams[0], ItemParser_GetItems_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &itemParserGetItemsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ItemParser_GetItemsClient interface {
	Recv() (*ItemResponse, error)
	grpc.ClientStream
}

type itemParserGetItemsClient struct {
	grpc.ClientStream
}

func (x *itemParserGetItemsClient) Recv() (*ItemResponse, error) {
	m := new(ItemResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *itemParserClient) GetItemCharacteristics(ctx context.Context, in *CharacteristicsRequest, opts ...grpc.CallOption) (ItemParser_GetItemCharacteristicsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ItemParser_ServiceDesc.Streams[1], ItemParser_GetItemCharacteristics_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &itemParserGetItemCharacteristicsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ItemParser_GetItemCharacteristicsClient interface {
	Recv() (*CharacteristicResponse, error)
	grpc.ClientStream
}

type itemParserGetItemCharacteristicsClient struct {
	grpc.ClientStream
}

func (x *itemParserGetItemCharacteristicsClient) Recv() (*CharacteristicResponse, error) {
	m := new(CharacteristicResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *itemParserClient) GetCategoryFilters(ctx context.Context, in *FiltersRequest, opts ...grpc.CallOption) (ItemParser_GetCategoryFiltersClient, error) {
	stream, err := c.cc.NewStream(ctx, &ItemParser_ServiceDesc.Streams[2], ItemParser_GetCategoryFilters_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &itemParserGetCategoryFiltersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ItemParser_GetCategoryFiltersClient interface {
	Recv() (*FilterResponse, error)
	grpc.ClientStream
}

type itemParserGetCategoryFiltersClient struct {
	grpc.ClientStream
}

func (x *itemParserGetCategoryFiltersClient) Recv() (*FilterResponse, error) {
	m := new(FilterResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ItemParserServer is the server API for ItemParser service.
// All implementations must embed UnimplementedItemParserServer
// for forward compatibility
type ItemParserServer interface {
	GetItems(*ItemsRequest, ItemParser_GetItemsServer) error
	GetItemCharacteristics(*CharacteristicsRequest, ItemParser_GetItemCharacteristicsServer) error
	GetCategoryFilters(*FiltersRequest, ItemParser_GetCategoryFiltersServer) error
	mustEmbedUnimplementedItemParserServer()
}

// UnimplementedItemParserServer must be embedded to have forward compatible implementations.
type UnimplementedItemParserServer struct {
}

func (UnimplementedItemParserServer) GetItems(*ItemsRequest, ItemParser_GetItemsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetItems not implemented")
}
func (UnimplementedItemParserServer) GetItemCharacteristics(*CharacteristicsRequest, ItemParser_GetItemCharacteristicsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetItemCharacteristics not implemented")
}
func (UnimplementedItemParserServer) GetCategoryFilters(*FiltersRequest, ItemParser_GetCategoryFiltersServer) error {
	return status.Errorf(codes.Unimplemented, "method GetCategoryFilters not implemented")
}
func (UnimplementedItemParserServer) mustEmbedUnimplementedItemParserServer() {}

// UnsafeItemParserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ItemParserServer will
// result in compilation errors.
type UnsafeItemParserServer interface {
	mustEmbedUnimplementedItemParserServer()
}

func RegisterItemParserServer(s grpc.ServiceRegistrar, srv ItemParserServer) {
	s.RegisterService(&ItemParser_ServiceDesc, srv)
}

func _ItemParser_GetItems_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ItemsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ItemParserServer).GetItems(m, &itemParserGetItemsServer{stream})
}

type ItemParser_GetItemsServer interface {
	Send(*ItemResponse) error
	grpc.ServerStream
}

type itemParserGetItemsServer struct {
	grpc.ServerStream
}

func (x *itemParserGetItemsServer) Send(m *ItemResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ItemParser_GetItemCharacteristics_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CharacteristicsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ItemParserServer).GetItemCharacteristics(m, &itemParserGetItemCharacteristicsServer{stream})
}

type ItemParser_GetItemCharacteristicsServer interface {
	Send(*CharacteristicResponse) error
	grpc.ServerStream
}

type itemParserGetItemCharacteristicsServer struct {
	grpc.ServerStream
}

func (x *itemParserGetItemCharacteristicsServer) Send(m *CharacteristicResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ItemParser_GetCategoryFilters_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FiltersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ItemParserServer).GetCategoryFilters(m, &itemParserGetCategoryFiltersServer{stream})
}

type ItemParser_GetCategoryFiltersServer interface {
	Send(*FilterResponse) error
	grpc.ServerStream
}

type itemParserGetCategoryFiltersServer struct {
	grpc.ServerStream
}

func (x *itemParserGetCategoryFiltersServer) Send(m *FilterResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ItemParser_ServiceDesc is the grpc.ServiceDesc for ItemParser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ItemParser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "app.protos.ItemParser",
	HandlerType: (*ItemParserServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetItems",
			Handler:       _ItemParser_GetItems_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetItemCharacteristics",
			Handler:       _ItemParser_GetItemCharacteristics_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetCategoryFilters",
			Handler:       _ItemParser_GetCategoryFilters_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "app/protos/items.proto",
}
