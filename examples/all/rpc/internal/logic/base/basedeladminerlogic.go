package baselogic

import (
	"context"

	"github.com/licat233/genzero/examples/all/rpc/admin_pb"
	"github.com/licat233/genzero/examples/all/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BaseDelAdminerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseDelAdminerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseDelAdminerLogic {
	return &BaseDelAdminerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除管理员
func (l *BaseDelAdminerLogic) BaseDelAdminer(in *admin_pb.DelAdminerReq) (*admin_pb.DelAdminerResp, error) {
	// todo: add your logic here and delete this line

	return &admin_pb.DelAdminerResp{}, nil
}
