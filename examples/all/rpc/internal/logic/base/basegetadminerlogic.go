package baselogic

import (
	"context"

	"github.com/licat233/genzero/examples/all/rpc/admin_pb"
	"github.com/licat233/genzero/examples/all/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BaseGetAdminerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseGetAdminerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseGetAdminerLogic {
	return &BaseGetAdminerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取管理员
func (l *BaseGetAdminerLogic) BaseGetAdminer(in *admin_pb.GetAdminerReq) (*admin_pb.GetAdminerResp, error) {
	// todo: add your logic here and delete this line

	return &admin_pb.GetAdminerResp{}, nil
}
