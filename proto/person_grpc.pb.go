// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: proto/person.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PersonService_Create_FullMethodName = "/personservice.PersonService/Create"
	PersonService_Read_FullMethodName   = "/personservice.PersonService/Read"
	PersonService_Update_FullMethodName = "/personservice.PersonService/Update"
	PersonService_Delete_FullMethodName = "/personservice.PersonService/Delete"
)

// PersonServiceClient is the client API for PersonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PersonServiceClient interface {
	Create(ctx context.Context, in *CreatePersonRequest, opts ...grpc.CallOption) (*PersonProfileResponse, error)
	Read(ctx context.Context, in *SinglePersonRequest, opts ...grpc.CallOption) (*PersonProfileResponse, error)
	Update(ctx context.Context, in *UpdatePersonRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
	Delete(ctx context.Context, in *SinglePersonRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
}

type personServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPersonServiceClient(cc grpc.ClientConnInterface) PersonServiceClient {
	return &personServiceClient{cc}
}

func (c *personServiceClient) Create(ctx context.Context, in *CreatePersonRequest, opts ...grpc.CallOption) (*PersonProfileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PersonProfileResponse)
	err := c.cc.Invoke(ctx, PersonService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personServiceClient) Read(ctx context.Context, in *SinglePersonRequest, opts ...grpc.CallOption) (*PersonProfileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PersonProfileResponse)
	err := c.cc.Invoke(ctx, PersonService_Read_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personServiceClient) Update(ctx context.Context, in *UpdatePersonRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, PersonService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personServiceClient) Delete(ctx context.Context, in *SinglePersonRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, PersonService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PersonServiceServer is the server API for PersonService service.
// All implementations must embed UnimplementedPersonServiceServer
// for forward compatibility.
type PersonServiceServer interface {
	Create(context.Context, *CreatePersonRequest) (*PersonProfileResponse, error)
	Read(context.Context, *SinglePersonRequest) (*PersonProfileResponse, error)
	Update(context.Context, *UpdatePersonRequest) (*SuccessResponse, error)
	Delete(context.Context, *SinglePersonRequest) (*SuccessResponse, error)
	mustEmbedUnimplementedPersonServiceServer()
}

// UnimplementedPersonServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPersonServiceServer struct{}

func (UnimplementedPersonServiceServer) Create(context.Context, *CreatePersonRequest) (*PersonProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPersonServiceServer) Read(context.Context, *SinglePersonRequest) (*PersonProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedPersonServiceServer) Update(context.Context, *UpdatePersonRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPersonServiceServer) Delete(context.Context, *SinglePersonRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPersonServiceServer) mustEmbedUnimplementedPersonServiceServer() {}
func (UnimplementedPersonServiceServer) testEmbeddedByValue()                       {}

// UnsafePersonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PersonServiceServer will
// result in compilation errors.
type UnsafePersonServiceServer interface {
	mustEmbedUnimplementedPersonServiceServer()
}

func RegisterPersonServiceServer(s grpc.ServiceRegistrar, srv PersonServiceServer) {
	// If the following call pancis, it indicates UnimplementedPersonServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PersonService_ServiceDesc, srv)
}

func _PersonService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PersonService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).Create(ctx, req.(*CreatePersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SinglePersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PersonService_Read_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).Read(ctx, req.(*SinglePersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PersonService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).Update(ctx, req.(*UpdatePersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SinglePersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PersonService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).Delete(ctx, req.(*SinglePersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PersonService_ServiceDesc is the grpc.ServiceDesc for PersonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PersonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "personservice.PersonService",
	HandlerType: (*PersonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PersonService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _PersonService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PersonService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PersonService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/person.proto",
}
