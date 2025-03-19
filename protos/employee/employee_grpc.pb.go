// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: protos/employee.proto

package employee

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
	EmployeeService_CreateEmployee_FullMethodName = "/employee.EmployeeService/CreateEmployee"
	EmployeeService_GetEmployee_FullMethodName    = "/employee.EmployeeService/GetEmployee"
	EmployeeService_UpdateEmployee_FullMethodName = "/employee.EmployeeService/UpdateEmployee"
	EmployeeService_DeleteEmployee_FullMethodName = "/employee.EmployeeService/DeleteEmployee"
)

// EmployeeServiceClient is the client API for EmployeeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmployeeServiceClient interface {
	CreateEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*Employee, error)
	GetEmployee(ctx context.Context, in *EmployeeRequest, opts ...grpc.CallOption) (*Employee, error)
	UpdateEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*Employee, error)
	DeleteEmployee(ctx context.Context, in *EmployeeRequest, opts ...grpc.CallOption) (*StringResponse, error)
}

type employeeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmployeeServiceClient(cc grpc.ClientConnInterface) EmployeeServiceClient {
	return &employeeServiceClient{cc}
}

func (c *employeeServiceClient) CreateEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*Employee, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Employee)
	err := c.cc.Invoke(ctx, EmployeeService_CreateEmployee_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) GetEmployee(ctx context.Context, in *EmployeeRequest, opts ...grpc.CallOption) (*Employee, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Employee)
	err := c.cc.Invoke(ctx, EmployeeService_GetEmployee_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) UpdateEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*Employee, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Employee)
	err := c.cc.Invoke(ctx, EmployeeService_UpdateEmployee_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) DeleteEmployee(ctx context.Context, in *EmployeeRequest, opts ...grpc.CallOption) (*StringResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StringResponse)
	err := c.cc.Invoke(ctx, EmployeeService_DeleteEmployee_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmployeeServiceServer is the server API for EmployeeService service.
// All implementations must embed UnimplementedEmployeeServiceServer
// for forward compatibility.
type EmployeeServiceServer interface {
	CreateEmployee(context.Context, *Employee) (*Employee, error)
	GetEmployee(context.Context, *EmployeeRequest) (*Employee, error)
	UpdateEmployee(context.Context, *Employee) (*Employee, error)
	DeleteEmployee(context.Context, *EmployeeRequest) (*StringResponse, error)
	mustEmbedUnimplementedEmployeeServiceServer()
}

// UnimplementedEmployeeServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEmployeeServiceServer struct{}

func (UnimplementedEmployeeServiceServer) CreateEmployee(context.Context, *Employee) (*Employee, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEmployee not implemented")
}
func (UnimplementedEmployeeServiceServer) GetEmployee(context.Context, *EmployeeRequest) (*Employee, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmployee not implemented")
}
func (UnimplementedEmployeeServiceServer) UpdateEmployee(context.Context, *Employee) (*Employee, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmployee not implemented")
}
func (UnimplementedEmployeeServiceServer) DeleteEmployee(context.Context, *EmployeeRequest) (*StringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEmployee not implemented")
}
func (UnimplementedEmployeeServiceServer) mustEmbedUnimplementedEmployeeServiceServer() {}
func (UnimplementedEmployeeServiceServer) testEmbeddedByValue()                         {}

// UnsafeEmployeeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmployeeServiceServer will
// result in compilation errors.
type UnsafeEmployeeServiceServer interface {
	mustEmbedUnimplementedEmployeeServiceServer()
}

func RegisterEmployeeServiceServer(s grpc.ServiceRegistrar, srv EmployeeServiceServer) {
	// If the following call pancis, it indicates UnimplementedEmployeeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&EmployeeService_ServiceDesc, srv)
}

func _EmployeeService_CreateEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Employee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).CreateEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmployeeService_CreateEmployee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).CreateEmployee(ctx, req.(*Employee))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_GetEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).GetEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmployeeService_GetEmployee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).GetEmployee(ctx, req.(*EmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_UpdateEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Employee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).UpdateEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmployeeService_UpdateEmployee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).UpdateEmployee(ctx, req.(*Employee))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_DeleteEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).DeleteEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmployeeService_DeleteEmployee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).DeleteEmployee(ctx, req.(*EmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EmployeeService_ServiceDesc is the grpc.ServiceDesc for EmployeeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmployeeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "employee.EmployeeService",
	HandlerType: (*EmployeeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEmployee",
			Handler:    _EmployeeService_CreateEmployee_Handler,
		},
		{
			MethodName: "GetEmployee",
			Handler:    _EmployeeService_GetEmployee_Handler,
		},
		{
			MethodName: "UpdateEmployee",
			Handler:    _EmployeeService_UpdateEmployee_Handler,
		},
		{
			MethodName: "DeleteEmployee",
			Handler:    _EmployeeService_DeleteEmployee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/employee.proto",
}
