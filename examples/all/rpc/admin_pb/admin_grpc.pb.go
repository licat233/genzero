// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: admin.proto

package admin_pb

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

// BaseClient is the client API for Base service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BaseClient interface {
	// 添加管理员
	AddAdminer(ctx context.Context, in *AddAdminerReq, opts ...grpc.CallOption) (*AddAdminerResp, error)
	// 更新管理员
	PutAdminer(ctx context.Context, in *PutAdminerReq, opts ...grpc.CallOption) (*PutAdminerResp, error)
	// 获取管理员
	GetAdminer(ctx context.Context, in *GetAdminerReq, opts ...grpc.CallOption) (*GetAdminerResp, error)
	// 删除管理员
	DelAdminer(ctx context.Context, in *DelAdminerReq, opts ...grpc.CallOption) (*DelAdminerResp, error)
	// 获取管理员列表
	GetAdminerList(ctx context.Context, in *GetAdminerListReq, opts ...grpc.CallOption) (*GetAdminerListResp, error)
	// 获取管理员枚举列表
	GetAdminerEnums(ctx context.Context, in *GetAdminerEnumsReq, opts ...grpc.CallOption) (*Enums, error)
	// 添加jwt黑名单
	AddJwtBlacklist(ctx context.Context, in *AddJwtBlacklistReq, opts ...grpc.CallOption) (*AddJwtBlacklistResp, error)
	// 更新jwt黑名单
	PutJwtBlacklist(ctx context.Context, in *PutJwtBlacklistReq, opts ...grpc.CallOption) (*PutJwtBlacklistResp, error)
	// 获取jwt黑名单
	GetJwtBlacklist(ctx context.Context, in *GetJwtBlacklistReq, opts ...grpc.CallOption) (*GetJwtBlacklistResp, error)
	// 删除jwt黑名单
	DelJwtBlacklist(ctx context.Context, in *DelJwtBlacklistReq, opts ...grpc.CallOption) (*DelJwtBlacklistResp, error)
	// 获取jwt黑名单列表
	GetJwtBlacklistList(ctx context.Context, in *GetJwtBlacklistListReq, opts ...grpc.CallOption) (*GetJwtBlacklistListResp, error)
	// 获取jwt黑名单枚举列表
	GetJwtBlacklistEnums(ctx context.Context, in *GetJwtBlacklistEnumsReq, opts ...grpc.CallOption) (*Enums, error)
}

type baseClient struct {
	cc grpc.ClientConnInterface
}

func NewBaseClient(cc grpc.ClientConnInterface) BaseClient {
	return &baseClient{cc}
}

func (c *baseClient) AddAdminer(ctx context.Context, in *AddAdminerReq, opts ...grpc.CallOption) (*AddAdminerResp, error) {
	out := new(AddAdminerResp)
	err := c.cc.Invoke(ctx, "/admin_proto.Base/AddAdminer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseClient) PutAdminer(ctx context.Context, in *PutAdminerReq, opts ...grpc.CallOption) (*PutAdminerResp, error) {
	out := new(PutAdminerResp)
	err := c.cc.Invoke(ctx, "/admin_proto.Base/PutAdminer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseClient) GetAdminer(ctx context.Context, in *GetAdminerReq, opts ...grpc.CallOption) (*GetAdminerResp, error) {
	out := new(GetAdminerResp)
	err := c.cc.Invoke(ctx, "/admin_proto.Base/GetAdminer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseClient) DelAdminer(ctx context.Context, in *DelAdminerReq, opts ...grpc.CallOption) (*DelAdminerResp, error) {
	out := new(DelAdminerResp)
	err := c.cc.Invoke(ctx, "/admin_proto.Base/DelAdminer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseClient) GetAdminerList(ctx context.Context, in *GetAdminerListReq, opts ...grpc.CallOption) (*GetAdminerListResp, error) {
	out := new(GetAdminerListResp)
	err := c.cc.Invoke(ctx, "/admin_proto.Base/GetAdminerList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseClient) GetAdminerEnums(ctx context.Context, in *GetAdminerEnumsReq, opts ...grpc.CallOption) (*Enums, error) {
	out := new(Enums)
	err := c.cc.Invoke(ctx, "/admin_proto.Base/GetAdminerEnums", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseClient) AddJwtBlacklist(ctx context.Context, in *AddJwtBlacklistReq, opts ...grpc.CallOption) (*AddJwtBlacklistResp, error) {
	out := new(AddJwtBlacklistResp)
	err := c.cc.Invoke(ctx, "/admin_proto.Base/AddJwtBlacklist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseClient) PutJwtBlacklist(ctx context.Context, in *PutJwtBlacklistReq, opts ...grpc.CallOption) (*PutJwtBlacklistResp, error) {
	out := new(PutJwtBlacklistResp)
	err := c.cc.Invoke(ctx, "/admin_proto.Base/PutJwtBlacklist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseClient) GetJwtBlacklist(ctx context.Context, in *GetJwtBlacklistReq, opts ...grpc.CallOption) (*GetJwtBlacklistResp, error) {
	out := new(GetJwtBlacklistResp)
	err := c.cc.Invoke(ctx, "/admin_proto.Base/GetJwtBlacklist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseClient) DelJwtBlacklist(ctx context.Context, in *DelJwtBlacklistReq, opts ...grpc.CallOption) (*DelJwtBlacklistResp, error) {
	out := new(DelJwtBlacklistResp)
	err := c.cc.Invoke(ctx, "/admin_proto.Base/DelJwtBlacklist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseClient) GetJwtBlacklistList(ctx context.Context, in *GetJwtBlacklistListReq, opts ...grpc.CallOption) (*GetJwtBlacklistListResp, error) {
	out := new(GetJwtBlacklistListResp)
	err := c.cc.Invoke(ctx, "/admin_proto.Base/GetJwtBlacklistList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseClient) GetJwtBlacklistEnums(ctx context.Context, in *GetJwtBlacklistEnumsReq, opts ...grpc.CallOption) (*Enums, error) {
	out := new(Enums)
	err := c.cc.Invoke(ctx, "/admin_proto.Base/GetJwtBlacklistEnums", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BaseServer is the server API for Base service.
// All implementations must embed UnimplementedBaseServer
// for forward compatibility
type BaseServer interface {
	// 添加管理员
	AddAdminer(context.Context, *AddAdminerReq) (*AddAdminerResp, error)
	// 更新管理员
	PutAdminer(context.Context, *PutAdminerReq) (*PutAdminerResp, error)
	// 获取管理员
	GetAdminer(context.Context, *GetAdminerReq) (*GetAdminerResp, error)
	// 删除管理员
	DelAdminer(context.Context, *DelAdminerReq) (*DelAdminerResp, error)
	// 获取管理员列表
	GetAdminerList(context.Context, *GetAdminerListReq) (*GetAdminerListResp, error)
	// 获取管理员枚举列表
	GetAdminerEnums(context.Context, *GetAdminerEnumsReq) (*Enums, error)
	// 添加jwt黑名单
	AddJwtBlacklist(context.Context, *AddJwtBlacklistReq) (*AddJwtBlacklistResp, error)
	// 更新jwt黑名单
	PutJwtBlacklist(context.Context, *PutJwtBlacklistReq) (*PutJwtBlacklistResp, error)
	// 获取jwt黑名单
	GetJwtBlacklist(context.Context, *GetJwtBlacklistReq) (*GetJwtBlacklistResp, error)
	// 删除jwt黑名单
	DelJwtBlacklist(context.Context, *DelJwtBlacklistReq) (*DelJwtBlacklistResp, error)
	// 获取jwt黑名单列表
	GetJwtBlacklistList(context.Context, *GetJwtBlacklistListReq) (*GetJwtBlacklistListResp, error)
	// 获取jwt黑名单枚举列表
	GetJwtBlacklistEnums(context.Context, *GetJwtBlacklistEnumsReq) (*Enums, error)
	mustEmbedUnimplementedBaseServer()
}

// UnimplementedBaseServer must be embedded to have forward compatible implementations.
type UnimplementedBaseServer struct {
}

func (UnimplementedBaseServer) AddAdminer(context.Context, *AddAdminerReq) (*AddAdminerResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAdminer not implemented")
}
func (UnimplementedBaseServer) PutAdminer(context.Context, *PutAdminerReq) (*PutAdminerResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutAdminer not implemented")
}
func (UnimplementedBaseServer) GetAdminer(context.Context, *GetAdminerReq) (*GetAdminerResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdminer not implemented")
}
func (UnimplementedBaseServer) DelAdminer(context.Context, *DelAdminerReq) (*DelAdminerResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelAdminer not implemented")
}
func (UnimplementedBaseServer) GetAdminerList(context.Context, *GetAdminerListReq) (*GetAdminerListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdminerList not implemented")
}
func (UnimplementedBaseServer) GetAdminerEnums(context.Context, *GetAdminerEnumsReq) (*Enums, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdminerEnums not implemented")
}
func (UnimplementedBaseServer) AddJwtBlacklist(context.Context, *AddJwtBlacklistReq) (*AddJwtBlacklistResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddJwtBlacklist not implemented")
}
func (UnimplementedBaseServer) PutJwtBlacklist(context.Context, *PutJwtBlacklistReq) (*PutJwtBlacklistResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutJwtBlacklist not implemented")
}
func (UnimplementedBaseServer) GetJwtBlacklist(context.Context, *GetJwtBlacklistReq) (*GetJwtBlacklistResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJwtBlacklist not implemented")
}
func (UnimplementedBaseServer) DelJwtBlacklist(context.Context, *DelJwtBlacklistReq) (*DelJwtBlacklistResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelJwtBlacklist not implemented")
}
func (UnimplementedBaseServer) GetJwtBlacklistList(context.Context, *GetJwtBlacklistListReq) (*GetJwtBlacklistListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJwtBlacklistList not implemented")
}
func (UnimplementedBaseServer) GetJwtBlacklistEnums(context.Context, *GetJwtBlacklistEnumsReq) (*Enums, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJwtBlacklistEnums not implemented")
}
func (UnimplementedBaseServer) mustEmbedUnimplementedBaseServer() {}

// UnsafeBaseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BaseServer will
// result in compilation errors.
type UnsafeBaseServer interface {
	mustEmbedUnimplementedBaseServer()
}

func RegisterBaseServer(s grpc.ServiceRegistrar, srv BaseServer) {
	s.RegisterService(&Base_ServiceDesc, srv)
}

func _Base_AddAdminer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAdminerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServer).AddAdminer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_proto.Base/AddAdminer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServer).AddAdminer(ctx, req.(*AddAdminerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Base_PutAdminer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutAdminerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServer).PutAdminer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_proto.Base/PutAdminer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServer).PutAdminer(ctx, req.(*PutAdminerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Base_GetAdminer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdminerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServer).GetAdminer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_proto.Base/GetAdminer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServer).GetAdminer(ctx, req.(*GetAdminerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Base_DelAdminer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelAdminerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServer).DelAdminer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_proto.Base/DelAdminer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServer).DelAdminer(ctx, req.(*DelAdminerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Base_GetAdminerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdminerListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServer).GetAdminerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_proto.Base/GetAdminerList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServer).GetAdminerList(ctx, req.(*GetAdminerListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Base_GetAdminerEnums_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdminerEnumsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServer).GetAdminerEnums(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_proto.Base/GetAdminerEnums",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServer).GetAdminerEnums(ctx, req.(*GetAdminerEnumsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Base_AddJwtBlacklist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddJwtBlacklistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServer).AddJwtBlacklist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_proto.Base/AddJwtBlacklist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServer).AddJwtBlacklist(ctx, req.(*AddJwtBlacklistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Base_PutJwtBlacklist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutJwtBlacklistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServer).PutJwtBlacklist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_proto.Base/PutJwtBlacklist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServer).PutJwtBlacklist(ctx, req.(*PutJwtBlacklistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Base_GetJwtBlacklist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJwtBlacklistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServer).GetJwtBlacklist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_proto.Base/GetJwtBlacklist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServer).GetJwtBlacklist(ctx, req.(*GetJwtBlacklistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Base_DelJwtBlacklist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelJwtBlacklistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServer).DelJwtBlacklist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_proto.Base/DelJwtBlacklist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServer).DelJwtBlacklist(ctx, req.(*DelJwtBlacklistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Base_GetJwtBlacklistList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJwtBlacklistListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServer).GetJwtBlacklistList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_proto.Base/GetJwtBlacklistList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServer).GetJwtBlacklistList(ctx, req.(*GetJwtBlacklistListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Base_GetJwtBlacklistEnums_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJwtBlacklistEnumsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServer).GetJwtBlacklistEnums(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_proto.Base/GetJwtBlacklistEnums",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServer).GetJwtBlacklistEnums(ctx, req.(*GetJwtBlacklistEnumsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Base_ServiceDesc is the grpc.ServiceDesc for Base service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Base_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin_proto.Base",
	HandlerType: (*BaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddAdminer",
			Handler:    _Base_AddAdminer_Handler,
		},
		{
			MethodName: "PutAdminer",
			Handler:    _Base_PutAdminer_Handler,
		},
		{
			MethodName: "GetAdminer",
			Handler:    _Base_GetAdminer_Handler,
		},
		{
			MethodName: "DelAdminer",
			Handler:    _Base_DelAdminer_Handler,
		},
		{
			MethodName: "GetAdminerList",
			Handler:    _Base_GetAdminerList_Handler,
		},
		{
			MethodName: "GetAdminerEnums",
			Handler:    _Base_GetAdminerEnums_Handler,
		},
		{
			MethodName: "AddJwtBlacklist",
			Handler:    _Base_AddJwtBlacklist_Handler,
		},
		{
			MethodName: "PutJwtBlacklist",
			Handler:    _Base_PutJwtBlacklist_Handler,
		},
		{
			MethodName: "GetJwtBlacklist",
			Handler:    _Base_GetJwtBlacklist_Handler,
		},
		{
			MethodName: "DelJwtBlacklist",
			Handler:    _Base_DelJwtBlacklist_Handler,
		},
		{
			MethodName: "GetJwtBlacklistList",
			Handler:    _Base_GetJwtBlacklistList_Handler,
		},
		{
			MethodName: "GetJwtBlacklistEnums",
			Handler:    _Base_GetJwtBlacklistEnums_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.proto",
}

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminClient interface {
	// Preset methods can be deleted. At least one method needs to be defined in a service
	// 预置的方法，可以删除，一个 service 中至少需要定义一个 method
	Test(ctx context.Context, in *NilReq, opts ...grpc.CallOption) (*NilResp, error)
}

type adminClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminClient(cc grpc.ClientConnInterface) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) Test(ctx context.Context, in *NilReq, opts ...grpc.CallOption) (*NilResp, error) {
	out := new(NilResp)
	err := c.cc.Invoke(ctx, "/admin_proto.Admin/Test", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
// All implementations must embed UnimplementedAdminServer
// for forward compatibility
type AdminServer interface {
	// Preset methods can be deleted. At least one method needs to be defined in a service
	// 预置的方法，可以删除，一个 service 中至少需要定义一个 method
	Test(context.Context, *NilReq) (*NilResp, error)
	mustEmbedUnimplementedAdminServer()
}

// UnimplementedAdminServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (UnimplementedAdminServer) Test(context.Context, *NilReq) (*NilResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test not implemented")
}
func (UnimplementedAdminServer) mustEmbedUnimplementedAdminServer() {}

// UnsafeAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServer will
// result in compilation errors.
type UnsafeAdminServer interface {
	mustEmbedUnimplementedAdminServer()
}

func RegisterAdminServer(s grpc.ServiceRegistrar, srv AdminServer) {
	s.RegisterService(&Admin_ServiceDesc, srv)
}

func _Admin_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NilReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_proto.Admin/Test",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).Test(ctx, req.(*NilReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Admin_ServiceDesc is the grpc.ServiceDesc for Admin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Admin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin_proto.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Test",
			Handler:    _Admin_Test_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.proto",
}