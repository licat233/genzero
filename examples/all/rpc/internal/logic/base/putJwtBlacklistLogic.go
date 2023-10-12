package baselogic

import (
	"context"

	"github.com/licat233/genzero/examples/all/rpc/admin_pb"
	"github.com/licat233/genzero/examples/all/rpc/internal/logic/dataconv"
	"github.com/licat233/genzero/examples/all/rpc/internal/svc"
	"github.com/licat233/genzero/examples/example_pkg/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutJwtBlacklistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutJwtBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutJwtBlacklistLogic {
	return &PutJwtBlacklistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新jwt黑名单
func (l *PutJwtBlacklistLogic) PutJwtBlacklist(in *admin_pb.PutJwtBlacklistReq) (*admin_pb.PutJwtBlacklistResp, error) {
	data := dataconv.PbJwtBlacklistToMdJwtBlacklist(in.JwtBlacklist)
	if err := l.svcCtx.JwtBlacklistModel.Update(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.PutJwtBlacklistResp{}, nil
}
