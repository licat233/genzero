package baselogic

import (
	"context"

	"github.com/licat233/genzero/examples/all/model"
	"github.com/licat233/genzero/examples/all/rpc/admin_pb"
	"github.com/licat233/genzero/examples/all/rpc/internal/logic/dataconv"
	"github.com/licat233/genzero/examples/all/rpc/internal/svc"
	"github.com/licat233/genzero/examples/example_pkg/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetJwtBlacklistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetJwtBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetJwtBlacklistLogic {
	return &GetJwtBlacklistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取jwt黑名单
func (l *GetJwtBlacklistLogic) GetJwtBlacklist(in *admin_pb.GetJwtBlacklistReq) (*admin_pb.GetJwtBlacklistResp, error) {
	res, err := l.svcCtx.JwtBlacklistModel.FindById(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.GetJwtBlacklistResp{
		JwtBlacklist: dataconv.MdJwtBlacklistToPbJwtBlacklist(res),
	}, nil
}
