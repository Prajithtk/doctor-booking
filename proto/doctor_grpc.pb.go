// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: proto/doctor.proto

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

const (
	DoctorService_AddDoctor_FullMethodName         = "/doctor.DoctorService/AddDoctor"
	DoctorService_RemoveDoctor_FullMethodName      = "/doctor.DoctorService/RemoveDoctor"
	DoctorService_ListDoctors_FullMethodName       = "/doctor.DoctorService/ListDoctors"
	DoctorService_CheckAvailability_FullMethodName = "/doctor.DoctorService/CheckAvailability"
)

// DoctorServiceClient is the client API for DoctorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DoctorServiceClient interface {
	AddDoctor(ctx context.Context, in *AddDoctorRequest, opts ...grpc.CallOption) (*DoctorResponse, error)
	RemoveDoctor(ctx context.Context, in *RemoveDoctorRequest, opts ...grpc.CallOption) (*RemoveDoctorResponse, error)
	ListDoctors(ctx context.Context, in *ListDoctorsRequest, opts ...grpc.CallOption) (*ListDoctorsResponse, error)
	CheckAvailability(ctx context.Context, in *AvailabilityRequest, opts ...grpc.CallOption) (*AvailabilityResponse, error)
}

type doctorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDoctorServiceClient(cc grpc.ClientConnInterface) DoctorServiceClient {
	return &doctorServiceClient{cc}
}

func (c *doctorServiceClient) AddDoctor(ctx context.Context, in *AddDoctorRequest, opts ...grpc.CallOption) (*DoctorResponse, error) {
	out := new(DoctorResponse)
	err := c.cc.Invoke(ctx, DoctorService_AddDoctor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) RemoveDoctor(ctx context.Context, in *RemoveDoctorRequest, opts ...grpc.CallOption) (*RemoveDoctorResponse, error) {
	out := new(RemoveDoctorResponse)
	err := c.cc.Invoke(ctx, DoctorService_RemoveDoctor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) ListDoctors(ctx context.Context, in *ListDoctorsRequest, opts ...grpc.CallOption) (*ListDoctorsResponse, error) {
	out := new(ListDoctorsResponse)
	err := c.cc.Invoke(ctx, DoctorService_ListDoctors_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) CheckAvailability(ctx context.Context, in *AvailabilityRequest, opts ...grpc.CallOption) (*AvailabilityResponse, error) {
	out := new(AvailabilityResponse)
	err := c.cc.Invoke(ctx, DoctorService_CheckAvailability_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DoctorServiceServer is the server API for DoctorService service.
// All implementations must embed UnimplementedDoctorServiceServer
// for forward compatibility
type DoctorServiceServer interface {
	AddDoctor(context.Context, *AddDoctorRequest) (*DoctorResponse, error)
	RemoveDoctor(context.Context, *RemoveDoctorRequest) (*RemoveDoctorResponse, error)
	ListDoctors(context.Context, *ListDoctorsRequest) (*ListDoctorsResponse, error)
	CheckAvailability(context.Context, *AvailabilityRequest) (*AvailabilityResponse, error)
	mustEmbedUnimplementedDoctorServiceServer()
}

// UnimplementedDoctorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDoctorServiceServer struct {
}

func (UnimplementedDoctorServiceServer) AddDoctor(context.Context, *AddDoctorRequest) (*DoctorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDoctor not implemented")
}
func (UnimplementedDoctorServiceServer) RemoveDoctor(context.Context, *RemoveDoctorRequest) (*RemoveDoctorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDoctor not implemented")
}
func (UnimplementedDoctorServiceServer) ListDoctors(context.Context, *ListDoctorsRequest) (*ListDoctorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDoctors not implemented")
}
func (UnimplementedDoctorServiceServer) CheckAvailability(context.Context, *AvailabilityRequest) (*AvailabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAvailability not implemented")
}
func (UnimplementedDoctorServiceServer) mustEmbedUnimplementedDoctorServiceServer() {}

// UnsafeDoctorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DoctorServiceServer will
// result in compilation errors.
type UnsafeDoctorServiceServer interface {
	mustEmbedUnimplementedDoctorServiceServer()
}

func RegisterDoctorServiceServer(s grpc.ServiceRegistrar, srv DoctorServiceServer) {
	s.RegisterService(&DoctorService_ServiceDesc, srv)
}

func _DoctorService_AddDoctor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDoctorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).AddDoctor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_AddDoctor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).AddDoctor(ctx, req.(*AddDoctorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_RemoveDoctor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveDoctorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).RemoveDoctor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_RemoveDoctor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).RemoveDoctor(ctx, req.(*RemoveDoctorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_ListDoctors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDoctorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).ListDoctors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_ListDoctors_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).ListDoctors(ctx, req.(*ListDoctorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_CheckAvailability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AvailabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).CheckAvailability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_CheckAvailability_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).CheckAvailability(ctx, req.(*AvailabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DoctorService_ServiceDesc is the grpc.ServiceDesc for DoctorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DoctorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "doctor.DoctorService",
	HandlerType: (*DoctorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddDoctor",
			Handler:    _DoctorService_AddDoctor_Handler,
		},
		{
			MethodName: "RemoveDoctor",
			Handler:    _DoctorService_RemoveDoctor_Handler,
		},
		{
			MethodName: "ListDoctors",
			Handler:    _DoctorService_ListDoctors_Handler,
		},
		{
			MethodName: "CheckAvailability",
			Handler:    _DoctorService_CheckAvailability_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/doctor.proto",
}
