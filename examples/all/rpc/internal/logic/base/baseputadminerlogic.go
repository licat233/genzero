package baselogic

import (
	"context"

	"github.com/licat233/genzero/examples/all/rpc/admin_pb"
	"github.com/licat233/genzero/examples/all/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BasePutAdminerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBasePutAdminerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BasePutAdminerLogic {
	return &BasePutAdminerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新管理员
func (l *BasePutAdminerLogic) BasePutAdminer(in *admin_pb.PutAdminerReq) (*admin_pb.PutAdminerResp, error) {
	// todo: add your logic here and delete this line

	return &admin_pb.PutAdminerResp{}, nil
}
