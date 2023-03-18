package baselogic

import (
	"context"

	"github.com/licat233/genzero/examples/all/rpc/admin_pb"
	"github.com/licat233/genzero/examples/all/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BaseAddJwtBlacklistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseAddJwtBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseAddJwtBlacklistLogic {
	return &BaseAddJwtBlacklistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加jwt黑名单
func (l *BaseAddJwtBlacklistLogic) BaseAddJwtBlacklist(in *admin_pb.AddJwtBlacklistReq) (*admin_pb.AddJwtBlacklistResp, error) {
	// todo: add your logic here and delete this line

	return &admin_pb.AddJwtBlacklistResp{}, nil
}
