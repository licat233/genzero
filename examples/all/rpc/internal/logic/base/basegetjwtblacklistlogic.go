package baselogic

import (
	"context"

	"github.com/licat233/genzero/examples/all/rpc/admin_pb"
	"github.com/licat233/genzero/examples/all/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BaseGetJwtBlacklistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseGetJwtBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseGetJwtBlacklistLogic {
	return &BaseGetJwtBlacklistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取jwt黑名单
func (l *BaseGetJwtBlacklistLogic) BaseGetJwtBlacklist(in *admin_pb.GetJwtBlacklistReq) (*admin_pb.GetJwtBlacklistResp, error) {
	// todo: add your logic here and delete this line

	return &admin_pb.GetJwtBlacklistResp{}, nil
}
