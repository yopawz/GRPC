// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// ManageEmployeeWithGQLClient is the client API for ManageEmployeeWithGQL service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManageEmployeeWithGQLClient interface {
	GetAllEmployeesWithGQL(ctx context.Context, in *GetAllEmployeesWithGQLParams, opts ...grpc.CallOption) (*EmployeeListGQL, error)
	InsertEmployeeWithGQL(ctx context.Context, in *InsertEmployeeWithGQLParams, opts ...grpc.CallOption) (*ResponseGQL, error)
	UpdateEmployeeWithGQL(ctx context.Context, in *UpdateEmployeeWithGQLParams, opts ...grpc.CallOption) (*ResponseGQL, error)
	DeleteEmployeeWithGQL(ctx context.Context, in *DeleteEmployeeWithGQLParams, opts ...grpc.CallOption) (*ResponseGQL, error)
	GetEmployeeWithIdWithGQL(ctx context.Context, in *GetEmployeeWithIdWithGQLParams, opts ...grpc.CallOption) (*EmployeeGQL, error)
}

type manageEmployeeWithGQLClient struct {
	cc grpc.ClientConnInterface
}

func NewManageEmployeeWithGQLClient(cc grpc.ClientConnInterface) ManageEmployeeWithGQLClient {
	return &manageEmployeeWithGQLClient{cc}
}

func (c *manageEmployeeWithGQLClient) GetAllEmployeesWithGQL(ctx context.Context, in *GetAllEmployeesWithGQLParams, opts ...grpc.CallOption) (*EmployeeListGQL, error) {
	out := new(EmployeeListGQL)
	err := c.cc.Invoke(ctx, "/proto.ManageEmployeeWithGQL/GetAllEmployeesWithGQL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageEmployeeWithGQLClient) InsertEmployeeWithGQL(ctx context.Context, in *InsertEmployeeWithGQLParams, opts ...grpc.CallOption) (*ResponseGQL, error) {
	out := new(ResponseGQL)
	err := c.cc.Invoke(ctx, "/proto.ManageEmployeeWithGQL/InsertEmployeeWithGQL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageEmployeeWithGQLClient) UpdateEmployeeWithGQL(ctx context.Context, in *UpdateEmployeeWithGQLParams, opts ...grpc.CallOption) (*ResponseGQL, error) {
	out := new(ResponseGQL)
	err := c.cc.Invoke(ctx, "/proto.ManageEmployeeWithGQL/UpdateEmployeeWithGQL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageEmployeeWithGQLClient) DeleteEmployeeWithGQL(ctx context.Context, in *DeleteEmployeeWithGQLParams, opts ...grpc.CallOption) (*ResponseGQL, error) {
	out := new(ResponseGQL)
	err := c.cc.Invoke(ctx, "/proto.ManageEmployeeWithGQL/DeleteEmployeeWithGQL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageEmployeeWithGQLClient) GetEmployeeWithIdWithGQL(ctx context.Context, in *GetEmployeeWithIdWithGQLParams, opts ...grpc.CallOption) (*EmployeeGQL, error) {
	out := new(EmployeeGQL)
	err := c.cc.Invoke(ctx, "/proto.ManageEmployeeWithGQL/GetEmployeeWithIdWithGQL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManageEmployeeWithGQLServer is the server API for ManageEmployeeWithGQL service.
// All implementations must embed UnimplementedManageEmployeeWithGQLServer
// for forward compatibility
type ManageEmployeeWithGQLServer interface {
	GetAllEmployeesWithGQL(context.Context, *GetAllEmployeesWithGQLParams) (*EmployeeListGQL, error)
	InsertEmployeeWithGQL(context.Context, *InsertEmployeeWithGQLParams) (*ResponseGQL, error)
	UpdateEmployeeWithGQL(context.Context, *UpdateEmployeeWithGQLParams) (*ResponseGQL, error)
	DeleteEmployeeWithGQL(context.Context, *DeleteEmployeeWithGQLParams) (*ResponseGQL, error)
	GetEmployeeWithIdWithGQL(context.Context, *GetEmployeeWithIdWithGQLParams) (*EmployeeGQL, error)
	mustEmbedUnimplementedManageEmployeeWithGQLServer()
}

// UnimplementedManageEmployeeWithGQLServer must be embedded to have forward compatible implementations.
type UnimplementedManageEmployeeWithGQLServer struct {
}

func (UnimplementedManageEmployeeWithGQLServer) GetAllEmployeesWithGQL(context.Context, *GetAllEmployeesWithGQLParams) (*EmployeeListGQL, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllEmployeesWithGQL not implemented")
}
func (UnimplementedManageEmployeeWithGQLServer) InsertEmployeeWithGQL(context.Context, *InsertEmployeeWithGQLParams) (*ResponseGQL, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertEmployeeWithGQL not implemented")
}
func (UnimplementedManageEmployeeWithGQLServer) UpdateEmployeeWithGQL(context.Context, *UpdateEmployeeWithGQLParams) (*ResponseGQL, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmployeeWithGQL not implemented")
}
func (UnimplementedManageEmployeeWithGQLServer) DeleteEmployeeWithGQL(context.Context, *DeleteEmployeeWithGQLParams) (*ResponseGQL, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEmployeeWithGQL not implemented")
}
func (UnimplementedManageEmployeeWithGQLServer) GetEmployeeWithIdWithGQL(context.Context, *GetEmployeeWithIdWithGQLParams) (*EmployeeGQL, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmployeeWithIdWithGQL not implemented")
}
func (UnimplementedManageEmployeeWithGQLServer) mustEmbedUnimplementedManageEmployeeWithGQLServer() {}

// UnsafeManageEmployeeWithGQLServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManageEmployeeWithGQLServer will
// result in compilation errors.
type UnsafeManageEmployeeWithGQLServer interface {
	mustEmbedUnimplementedManageEmployeeWithGQLServer()
}

func RegisterManageEmployeeWithGQLServer(s grpc.ServiceRegistrar, srv ManageEmployeeWithGQLServer) {
	s.RegisterService(&ManageEmployeeWithGQL_ServiceDesc, srv)
}

func _ManageEmployeeWithGQL_GetAllEmployeesWithGQL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllEmployeesWithGQLParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageEmployeeWithGQLServer).GetAllEmployeesWithGQL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ManageEmployeeWithGQL/GetAllEmployeesWithGQL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageEmployeeWithGQLServer).GetAllEmployeesWithGQL(ctx, req.(*GetAllEmployeesWithGQLParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManageEmployeeWithGQL_InsertEmployeeWithGQL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertEmployeeWithGQLParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageEmployeeWithGQLServer).InsertEmployeeWithGQL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ManageEmployeeWithGQL/InsertEmployeeWithGQL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageEmployeeWithGQLServer).InsertEmployeeWithGQL(ctx, req.(*InsertEmployeeWithGQLParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManageEmployeeWithGQL_UpdateEmployeeWithGQL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEmployeeWithGQLParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageEmployeeWithGQLServer).UpdateEmployeeWithGQL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ManageEmployeeWithGQL/UpdateEmployeeWithGQL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageEmployeeWithGQLServer).UpdateEmployeeWithGQL(ctx, req.(*UpdateEmployeeWithGQLParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManageEmployeeWithGQL_DeleteEmployeeWithGQL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEmployeeWithGQLParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageEmployeeWithGQLServer).DeleteEmployeeWithGQL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ManageEmployeeWithGQL/DeleteEmployeeWithGQL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageEmployeeWithGQLServer).DeleteEmployeeWithGQL(ctx, req.(*DeleteEmployeeWithGQLParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManageEmployeeWithGQL_GetEmployeeWithIdWithGQL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEmployeeWithIdWithGQLParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageEmployeeWithGQLServer).GetEmployeeWithIdWithGQL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ManageEmployeeWithGQL/GetEmployeeWithIdWithGQL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageEmployeeWithGQLServer).GetEmployeeWithIdWithGQL(ctx, req.(*GetEmployeeWithIdWithGQLParams))
	}
	return interceptor(ctx, in, info, handler)
}

// ManageEmployeeWithGQL_ServiceDesc is the grpc.ServiceDesc for ManageEmployeeWithGQL service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ManageEmployeeWithGQL_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ManageEmployeeWithGQL",
	HandlerType: (*ManageEmployeeWithGQLServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllEmployeesWithGQL",
			Handler:    _ManageEmployeeWithGQL_GetAllEmployeesWithGQL_Handler,
		},
		{
			MethodName: "InsertEmployeeWithGQL",
			Handler:    _ManageEmployeeWithGQL_InsertEmployeeWithGQL_Handler,
		},
		{
			MethodName: "UpdateEmployeeWithGQL",
			Handler:    _ManageEmployeeWithGQL_UpdateEmployeeWithGQL_Handler,
		},
		{
			MethodName: "DeleteEmployeeWithGQL",
			Handler:    _ManageEmployeeWithGQL_DeleteEmployeeWithGQL_Handler,
		},
		{
			MethodName: "GetEmployeeWithIdWithGQL",
			Handler:    _ManageEmployeeWithGQL_GetEmployeeWithIdWithGQL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/gql/gql.proto",
}
