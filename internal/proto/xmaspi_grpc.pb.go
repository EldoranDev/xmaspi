// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: xmaspi.proto

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

// XmasPIClient is the client API for XmasPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type XmasPIClient interface {
	SetLed(ctx context.Context, in *SetLedRequest, opts ...grpc.CallOption) (*SetLedResponse, error)
	SetAnimation(ctx context.Context, in *SetAnimationRequest, opts ...grpc.CallOption) (*SetAnimationResponse, error)
	SetStatic(ctx context.Context, in *SetStaticRequest, opts ...grpc.CallOption) (*SetStaticResponse, error)
	GetControllerInfo(ctx context.Context, in *ControllerInfoRequest, opts ...grpc.CallOption) (*ControllerInfoResponse, error)
	Render(ctx context.Context, in *RenderRequest, opts ...grpc.CallOption) (*RenderResponse, error)
	GetLedCount(ctx context.Context, in *GetLedCountRequest, opts ...grpc.CallOption) (*GetLedCountResponse, error)
	GetAnimations(ctx context.Context, in *GetAnimationsRequest, opts ...grpc.CallOption) (*GetAnimationsResponse, error)
	GetStatics(ctx context.Context, in *GetStaticsRequest, opts ...grpc.CallOption) (*GetStaticsResponse, error)
}

type xmasPIClient struct {
	cc grpc.ClientConnInterface
}

func NewXmasPIClient(cc grpc.ClientConnInterface) XmasPIClient {
	return &xmasPIClient{cc}
}

func (c *xmasPIClient) SetLed(ctx context.Context, in *SetLedRequest, opts ...grpc.CallOption) (*SetLedResponse, error) {
	out := new(SetLedResponse)
	err := c.cc.Invoke(ctx, "/XmasPI/SetLed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xmasPIClient) SetAnimation(ctx context.Context, in *SetAnimationRequest, opts ...grpc.CallOption) (*SetAnimationResponse, error) {
	out := new(SetAnimationResponse)
	err := c.cc.Invoke(ctx, "/XmasPI/SetAnimation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xmasPIClient) SetStatic(ctx context.Context, in *SetStaticRequest, opts ...grpc.CallOption) (*SetStaticResponse, error) {
	out := new(SetStaticResponse)
	err := c.cc.Invoke(ctx, "/XmasPI/SetStatic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xmasPIClient) GetControllerInfo(ctx context.Context, in *ControllerInfoRequest, opts ...grpc.CallOption) (*ControllerInfoResponse, error) {
	out := new(ControllerInfoResponse)
	err := c.cc.Invoke(ctx, "/XmasPI/GetControllerInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xmasPIClient) Render(ctx context.Context, in *RenderRequest, opts ...grpc.CallOption) (*RenderResponse, error) {
	out := new(RenderResponse)
	err := c.cc.Invoke(ctx, "/XmasPI/Render", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xmasPIClient) GetLedCount(ctx context.Context, in *GetLedCountRequest, opts ...grpc.CallOption) (*GetLedCountResponse, error) {
	out := new(GetLedCountResponse)
	err := c.cc.Invoke(ctx, "/XmasPI/GetLedCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xmasPIClient) GetAnimations(ctx context.Context, in *GetAnimationsRequest, opts ...grpc.CallOption) (*GetAnimationsResponse, error) {
	out := new(GetAnimationsResponse)
	err := c.cc.Invoke(ctx, "/XmasPI/GetAnimations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xmasPIClient) GetStatics(ctx context.Context, in *GetStaticsRequest, opts ...grpc.CallOption) (*GetStaticsResponse, error) {
	out := new(GetStaticsResponse)
	err := c.cc.Invoke(ctx, "/XmasPI/GetStatics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// XmasPIServer is the server API for XmasPI service.
// All implementations must embed UnimplementedXmasPIServer
// for forward compatibility
type XmasPIServer interface {
	SetLed(context.Context, *SetLedRequest) (*SetLedResponse, error)
	SetAnimation(context.Context, *SetAnimationRequest) (*SetAnimationResponse, error)
	SetStatic(context.Context, *SetStaticRequest) (*SetStaticResponse, error)
	GetControllerInfo(context.Context, *ControllerInfoRequest) (*ControllerInfoResponse, error)
	Render(context.Context, *RenderRequest) (*RenderResponse, error)
	GetLedCount(context.Context, *GetLedCountRequest) (*GetLedCountResponse, error)
	GetAnimations(context.Context, *GetAnimationsRequest) (*GetAnimationsResponse, error)
	GetStatics(context.Context, *GetStaticsRequest) (*GetStaticsResponse, error)
	mustEmbedUnimplementedXmasPIServer()
}

// UnimplementedXmasPIServer must be embedded to have forward compatible implementations.
type UnimplementedXmasPIServer struct {
}

func (UnimplementedXmasPIServer) SetLed(context.Context, *SetLedRequest) (*SetLedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetLed not implemented")
}
func (UnimplementedXmasPIServer) SetAnimation(context.Context, *SetAnimationRequest) (*SetAnimationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAnimation not implemented")
}
func (UnimplementedXmasPIServer) SetStatic(context.Context, *SetStaticRequest) (*SetStaticResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetStatic not implemented")
}
func (UnimplementedXmasPIServer) GetControllerInfo(context.Context, *ControllerInfoRequest) (*ControllerInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetControllerInfo not implemented")
}
func (UnimplementedXmasPIServer) Render(context.Context, *RenderRequest) (*RenderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Render not implemented")
}
func (UnimplementedXmasPIServer) GetLedCount(context.Context, *GetLedCountRequest) (*GetLedCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLedCount not implemented")
}
func (UnimplementedXmasPIServer) GetAnimations(context.Context, *GetAnimationsRequest) (*GetAnimationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnimations not implemented")
}
func (UnimplementedXmasPIServer) GetStatics(context.Context, *GetStaticsRequest) (*GetStaticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatics not implemented")
}
func (UnimplementedXmasPIServer) mustEmbedUnimplementedXmasPIServer() {}

// UnsafeXmasPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to XmasPIServer will
// result in compilation errors.
type UnsafeXmasPIServer interface {
	mustEmbedUnimplementedXmasPIServer()
}

func RegisterXmasPIServer(s grpc.ServiceRegistrar, srv XmasPIServer) {
	s.RegisterService(&XmasPI_ServiceDesc, srv)
}

func _XmasPI_SetLed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetLedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XmasPIServer).SetLed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/XmasPI/SetLed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XmasPIServer).SetLed(ctx, req.(*SetLedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XmasPI_SetAnimation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAnimationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XmasPIServer).SetAnimation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/XmasPI/SetAnimation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XmasPIServer).SetAnimation(ctx, req.(*SetAnimationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XmasPI_SetStatic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetStaticRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XmasPIServer).SetStatic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/XmasPI/SetStatic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XmasPIServer).SetStatic(ctx, req.(*SetStaticRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XmasPI_GetControllerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ControllerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XmasPIServer).GetControllerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/XmasPI/GetControllerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XmasPIServer).GetControllerInfo(ctx, req.(*ControllerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XmasPI_Render_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XmasPIServer).Render(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/XmasPI/Render",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XmasPIServer).Render(ctx, req.(*RenderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XmasPI_GetLedCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLedCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XmasPIServer).GetLedCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/XmasPI/GetLedCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XmasPIServer).GetLedCount(ctx, req.(*GetLedCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XmasPI_GetAnimations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAnimationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XmasPIServer).GetAnimations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/XmasPI/GetAnimations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XmasPIServer).GetAnimations(ctx, req.(*GetAnimationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XmasPI_GetStatics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStaticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XmasPIServer).GetStatics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/XmasPI/GetStatics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XmasPIServer).GetStatics(ctx, req.(*GetStaticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// XmasPI_ServiceDesc is the grpc.ServiceDesc for XmasPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var XmasPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "XmasPI",
	HandlerType: (*XmasPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetLed",
			Handler:    _XmasPI_SetLed_Handler,
		},
		{
			MethodName: "SetAnimation",
			Handler:    _XmasPI_SetAnimation_Handler,
		},
		{
			MethodName: "SetStatic",
			Handler:    _XmasPI_SetStatic_Handler,
		},
		{
			MethodName: "GetControllerInfo",
			Handler:    _XmasPI_GetControllerInfo_Handler,
		},
		{
			MethodName: "Render",
			Handler:    _XmasPI_Render_Handler,
		},
		{
			MethodName: "GetLedCount",
			Handler:    _XmasPI_GetLedCount_Handler,
		},
		{
			MethodName: "GetAnimations",
			Handler:    _XmasPI_GetAnimations_Handler,
		},
		{
			MethodName: "GetStatics",
			Handler:    _XmasPI_GetStatics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "xmaspi.proto",
}
