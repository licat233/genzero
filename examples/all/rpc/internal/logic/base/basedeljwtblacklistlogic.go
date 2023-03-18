package baselogic

import (
	"context"

	"github.com/licat233/genzero/examples/all/rpc/admin_pb"
	"github.com/licat233/genzero/examples/all/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BaseDelJwtBlacklistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseDelJwtBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseDelJwtBlacklistLogic {
	return &BaseDelJwtBlacklistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除jwt黑名单
func (l *BaseDelJwtBlacklistLogic) BaseDelJwtBlacklist(in *admin_pb.DelJwtBlacklistReq) (*admin_pb.DelJwtBlacklistResp, error) {
	// todo: add your logic here and delete this line

	return &admin_pb.DelJwtBlacklistResp{}, nil
}
