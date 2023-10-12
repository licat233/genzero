package baselogic

import (
	"context"

	"github.com/licat233/genzero/examples/all/rpc/admin_pb"
	"github.com/licat233/genzero/examples/all/rpc/internal/svc"
	"github.com/licat233/genzero/examples/example_pkg/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAdminerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAdminerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAdminerLogic {
	return &DelAdminerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除管理员
func (l *DelAdminerLogic) DelAdminer(in *admin_pb.DelAdminerReq) (*admin_pb.DelAdminerResp, error) {
	if err := l.svcCtx.AdminerModel.SoftDelete(l.ctx, in.Id); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.DelAdminerResp{}, nil
}
