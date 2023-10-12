// Code generated by goctl. DO NOT EDIT.
// Source: admin.proto

package server

import (
	"context"

	"github.com/licat233/genzero/examples/all/rpc/admin_pb"
	"github.com/licat233/genzero/examples/all/rpc/internal/logic/base"
	"github.com/licat233/genzero/examples/all/rpc/internal/svc"
)

type BaseServer struct {
	svcCtx *svc.ServiceContext
	admin_pb.UnimplementedBaseServer
}

func NewBaseServer(svcCtx *svc.ServiceContext) *BaseServer {
	return &BaseServer{
		svcCtx: svcCtx,
	}
}

// 添加管理员
func (s *BaseServer) AddAdminer(ctx context.Context, in *admin_pb.AddAdminerReq) (*admin_pb.AddAdminerResp, error) {
	l := baselogic.NewAddAdminerLogic(ctx, s.svcCtx)
	return l.AddAdminer(in)
}

// 更新管理员
func (s *BaseServer) PutAdminer(ctx context.Context, in *admin_pb.PutAdminerReq) (*admin_pb.PutAdminerResp, error) {
	l := baselogic.NewPutAdminerLogic(ctx, s.svcCtx)
	return l.PutAdminer(in)
}

// 获取管理员
func (s *BaseServer) GetAdminer(ctx context.Context, in *admin_pb.GetAdminerReq) (*admin_pb.GetAdminerResp, error) {
	l := baselogic.NewGetAdminerLogic(ctx, s.svcCtx)
	return l.GetAdminer(in)
}

// 删除管理员
func (s *BaseServer) DelAdminer(ctx context.Context, in *admin_pb.DelAdminerReq) (*admin_pb.DelAdminerResp, error) {
	l := baselogic.NewDelAdminerLogic(ctx, s.svcCtx)
	return l.DelAdminer(in)
}

// 获取管理员列表
func (s *BaseServer) GetAdminerList(ctx context.Context, in *admin_pb.GetAdminerListReq) (*admin_pb.GetAdminerListResp, error) {
	l := baselogic.NewGetAdminerListLogic(ctx, s.svcCtx)
	return l.GetAdminerList(in)
}

// 获取管理员枚举列表
func (s *BaseServer) GetAdminerEnums(ctx context.Context, in *admin_pb.GetAdminerEnumsReq) (*admin_pb.Enums, error) {
	l := baselogic.NewGetAdminerEnumsLogic(ctx, s.svcCtx)
	return l.GetAdminerEnums(in)
}

// 添加jwt黑名单
func (s *BaseServer) AddJwtBlacklist(ctx context.Context, in *admin_pb.AddJwtBlacklistReq) (*admin_pb.AddJwtBlacklistResp, error) {
	l := baselogic.NewAddJwtBlacklistLogic(ctx, s.svcCtx)
	return l.AddJwtBlacklist(in)
}

// 更新jwt黑名单
func (s *BaseServer) PutJwtBlacklist(ctx context.Context, in *admin_pb.PutJwtBlacklistReq) (*admin_pb.PutJwtBlacklistResp, error) {
	l := baselogic.NewPutJwtBlacklistLogic(ctx, s.svcCtx)
	return l.PutJwtBlacklist(in)
}

// 获取jwt黑名单
func (s *BaseServer) GetJwtBlacklist(ctx context.Context, in *admin_pb.GetJwtBlacklistReq) (*admin_pb.GetJwtBlacklistResp, error) {
	l := baselogic.NewGetJwtBlacklistLogic(ctx, s.svcCtx)
	return l.GetJwtBlacklist(in)
}

// 删除jwt黑名单
func (s *BaseServer) DelJwtBlacklist(ctx context.Context, in *admin_pb.DelJwtBlacklistReq) (*admin_pb.DelJwtBlacklistResp, error) {
	l := baselogic.NewDelJwtBlacklistLogic(ctx, s.svcCtx)
	return l.DelJwtBlacklist(in)
}

// 获取jwt黑名单列表
func (s *BaseServer) GetJwtBlacklistList(ctx context.Context, in *admin_pb.GetJwtBlacklistListReq) (*admin_pb.GetJwtBlacklistListResp, error) {
	l := baselogic.NewGetJwtBlacklistListLogic(ctx, s.svcCtx)
	return l.GetJwtBlacklistList(in)
}

// 获取jwt黑名单枚举列表
func (s *BaseServer) GetJwtBlacklistEnums(ctx context.Context, in *admin_pb.GetJwtBlacklistEnumsReq) (*admin_pb.Enums, error) {
	l := baselogic.NewGetJwtBlacklistEnumsLogic(ctx, s.svcCtx)
	return l.GetJwtBlacklistEnums(in)
}
