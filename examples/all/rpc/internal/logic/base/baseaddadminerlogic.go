package baselogic

import (
	"context"

	"github.com/licat233/genzero/examples/all/rpc/admin_pb"
	"github.com/licat233/genzero/examples/all/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BaseAddAdminerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseAddAdminerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseAddAdminerLogic {
	return &BaseAddAdminerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加管理员
func (l *BaseAddAdminerLogic) BaseAddAdminer(in *admin_pb.AddAdminerReq) (*admin_pb.AddAdminerResp, error) {
	// todo: add your logic here and delete this line

	return &admin_pb.AddAdminerResp{}, nil
}
