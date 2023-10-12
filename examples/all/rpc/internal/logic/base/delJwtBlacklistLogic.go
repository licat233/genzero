package baselogic

import (
	"context"

	"github.com/licat233/genzero/examples/all/rpc/admin_pb"
	"github.com/licat233/genzero/examples/all/rpc/internal/svc"
	"github.com/licat233/genzero/examples/example_pkg/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelJwtBlacklistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelJwtBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelJwtBlacklistLogic {
	return &DelJwtBlacklistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除jwt黑名单
func (l *DelJwtBlacklistLogic) DelJwtBlacklist(in *admin_pb.DelJwtBlacklistReq) (*admin_pb.DelJwtBlacklistResp, error) {
	if err := l.svcCtx.JwtBlacklistModel.Delete(l.ctx, in.Id); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.DelJwtBlacklistResp{}, nil
}
