package baselogic

import (
	"context"

	"github.com/licat233/genzero/examples/all/rpc/admin_pb"
	"github.com/licat233/genzero/examples/all/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BasePutJwtBlacklistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBasePutJwtBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BasePutJwtBlacklistLogic {
	return &BasePutJwtBlacklistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新jwt黑名单
func (l *BasePutJwtBlacklistLogic) BasePutJwtBlacklist(in *admin_pb.PutJwtBlacklistReq) (*admin_pb.PutJwtBlacklistResp, error) {
	// todo: add your logic here and delete this line

	return &admin_pb.PutJwtBlacklistResp{}, nil
}
