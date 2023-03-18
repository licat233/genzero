package baselogic

import (
	"context"

	"github.com/licat233/genzero/examples/all/rpc/admin_pb"
	"github.com/licat233/genzero/examples/all/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BaseGetJwtBlacklistListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseGetJwtBlacklistListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseGetJwtBlacklistListLogic {
	return &BaseGetJwtBlacklistListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取jwt黑名单列表
func (l *BaseGetJwtBlacklistListLogic) BaseGetJwtBlacklistList(in *admin_pb.GetJwtBlacklistListReq) (*admin_pb.GetJwtBlacklistListResp, error) {
	// todo: add your logic here and delete this line

	return &admin_pb.GetJwtBlacklistListResp{}, nil
}
