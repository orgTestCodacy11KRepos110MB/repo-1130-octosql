// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package plugins

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

// DatasourceClient is the client API for Datasource service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatasourceClient interface {
	GetTable(ctx context.Context, in *GetTableRequest, opts ...grpc.CallOption) (*GetTableResponse, error)
	PushDownPredicates(ctx context.Context, in *PushDownPredicatesRequest, opts ...grpc.CallOption) (*PushDownPredicatesResponse, error)
	Materialize(ctx context.Context, in *MaterializeRequest, opts ...grpc.CallOption) (*MaterializeResponse, error)
	Metadata(ctx context.Context, in *MetadataRequest, opts ...grpc.CallOption) (*MetadataResponse, error)
}

type datasourceClient struct {
	cc grpc.ClientConnInterface
}

func NewDatasourceClient(cc grpc.ClientConnInterface) DatasourceClient {
	return &datasourceClient{cc}
}

func (c *datasourceClient) GetTable(ctx context.Context, in *GetTableRequest, opts ...grpc.CallOption) (*GetTableResponse, error) {
	out := new(GetTableResponse)
	err := c.cc.Invoke(ctx, "/plugins.Datasource/GetTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasourceClient) PushDownPredicates(ctx context.Context, in *PushDownPredicatesRequest, opts ...grpc.CallOption) (*PushDownPredicatesResponse, error) {
	out := new(PushDownPredicatesResponse)
	err := c.cc.Invoke(ctx, "/plugins.Datasource/PushDownPredicates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasourceClient) Materialize(ctx context.Context, in *MaterializeRequest, opts ...grpc.CallOption) (*MaterializeResponse, error) {
	out := new(MaterializeResponse)
	err := c.cc.Invoke(ctx, "/plugins.Datasource/Materialize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasourceClient) Metadata(ctx context.Context, in *MetadataRequest, opts ...grpc.CallOption) (*MetadataResponse, error) {
	out := new(MetadataResponse)
	err := c.cc.Invoke(ctx, "/plugins.Datasource/Metadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatasourceServer is the server API for Datasource service.
// All implementations must embed UnimplementedDatasourceServer
// for forward compatibility
type DatasourceServer interface {
	GetTable(context.Context, *GetTableRequest) (*GetTableResponse, error)
	PushDownPredicates(context.Context, *PushDownPredicatesRequest) (*PushDownPredicatesResponse, error)
	Materialize(context.Context, *MaterializeRequest) (*MaterializeResponse, error)
	Metadata(context.Context, *MetadataRequest) (*MetadataResponse, error)
	mustEmbedUnimplementedDatasourceServer()
}

// UnimplementedDatasourceServer must be embedded to have forward compatible implementations.
type UnimplementedDatasourceServer struct {
}

func (UnimplementedDatasourceServer) GetTable(context.Context, *GetTableRequest) (*GetTableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTable not implemented")
}
func (UnimplementedDatasourceServer) PushDownPredicates(context.Context, *PushDownPredicatesRequest) (*PushDownPredicatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushDownPredicates not implemented")
}
func (UnimplementedDatasourceServer) Materialize(context.Context, *MaterializeRequest) (*MaterializeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Materialize not implemented")
}
func (UnimplementedDatasourceServer) Metadata(context.Context, *MetadataRequest) (*MetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Metadata not implemented")
}
func (UnimplementedDatasourceServer) mustEmbedUnimplementedDatasourceServer() {}

// UnsafeDatasourceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatasourceServer will
// result in compilation errors.
type UnsafeDatasourceServer interface {
	mustEmbedUnimplementedDatasourceServer()
}

func RegisterDatasourceServer(s grpc.ServiceRegistrar, srv DatasourceServer) {
	s.RegisterService(&Datasource_ServiceDesc, srv)
}

func _Datasource_GetTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasourceServer).GetTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugins.Datasource/GetTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasourceServer).GetTable(ctx, req.(*GetTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Datasource_PushDownPredicates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushDownPredicatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasourceServer).PushDownPredicates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugins.Datasource/PushDownPredicates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasourceServer).PushDownPredicates(ctx, req.(*PushDownPredicatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Datasource_Materialize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MaterializeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasourceServer).Materialize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugins.Datasource/Materialize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasourceServer).Materialize(ctx, req.(*MaterializeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Datasource_Metadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasourceServer).Metadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugins.Datasource/Metadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasourceServer).Metadata(ctx, req.(*MetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Datasource_ServiceDesc is the grpc.ServiceDesc for Datasource service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Datasource_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "plugins.Datasource",
	HandlerType: (*DatasourceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTable",
			Handler:    _Datasource_GetTable_Handler,
		},
		{
			MethodName: "PushDownPredicates",
			Handler:    _Datasource_PushDownPredicates_Handler,
		},
		{
			MethodName: "Materialize",
			Handler:    _Datasource_Materialize_Handler,
		},
		{
			MethodName: "Metadata",
			Handler:    _Datasource_Metadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugins.proto",
}

// ExecutionDatasourceClient is the client API for ExecutionDatasource service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExecutionDatasourceClient interface {
	Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (ExecutionDatasource_RunClient, error)
}

type executionDatasourceClient struct {
	cc grpc.ClientConnInterface
}

func NewExecutionDatasourceClient(cc grpc.ClientConnInterface) ExecutionDatasourceClient {
	return &executionDatasourceClient{cc}
}

func (c *executionDatasourceClient) Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (ExecutionDatasource_RunClient, error) {
	stream, err := c.cc.NewStream(ctx, &ExecutionDatasource_ServiceDesc.Streams[0], "/plugins.ExecutionDatasource/Run", opts...)
	if err != nil {
		return nil, err
	}
	x := &executionDatasourceRunClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ExecutionDatasource_RunClient interface {
	Recv() (*RunResponseMessage, error)
	grpc.ClientStream
}

type executionDatasourceRunClient struct {
	grpc.ClientStream
}

func (x *executionDatasourceRunClient) Recv() (*RunResponseMessage, error) {
	m := new(RunResponseMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExecutionDatasourceServer is the server API for ExecutionDatasource service.
// All implementations must embed UnimplementedExecutionDatasourceServer
// for forward compatibility
type ExecutionDatasourceServer interface {
	Run(*RunRequest, ExecutionDatasource_RunServer) error
	mustEmbedUnimplementedExecutionDatasourceServer()
}

// UnimplementedExecutionDatasourceServer must be embedded to have forward compatible implementations.
type UnimplementedExecutionDatasourceServer struct {
}

func (UnimplementedExecutionDatasourceServer) Run(*RunRequest, ExecutionDatasource_RunServer) error {
	return status.Errorf(codes.Unimplemented, "method Run not implemented")
}
func (UnimplementedExecutionDatasourceServer) mustEmbedUnimplementedExecutionDatasourceServer() {}

// UnsafeExecutionDatasourceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExecutionDatasourceServer will
// result in compilation errors.
type UnsafeExecutionDatasourceServer interface {
	mustEmbedUnimplementedExecutionDatasourceServer()
}

func RegisterExecutionDatasourceServer(s grpc.ServiceRegistrar, srv ExecutionDatasourceServer) {
	s.RegisterService(&ExecutionDatasource_ServiceDesc, srv)
}

func _ExecutionDatasource_Run_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RunRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExecutionDatasourceServer).Run(m, &executionDatasourceRunServer{stream})
}

type ExecutionDatasource_RunServer interface {
	Send(*RunResponseMessage) error
	grpc.ServerStream
}

type executionDatasourceRunServer struct {
	grpc.ServerStream
}

func (x *executionDatasourceRunServer) Send(m *RunResponseMessage) error {
	return x.ServerStream.SendMsg(m)
}

// ExecutionDatasource_ServiceDesc is the grpc.ServiceDesc for ExecutionDatasource service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExecutionDatasource_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "plugins.ExecutionDatasource",
	HandlerType: (*ExecutionDatasourceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Run",
			Handler:       _ExecutionDatasource_Run_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "plugins.proto",
}
