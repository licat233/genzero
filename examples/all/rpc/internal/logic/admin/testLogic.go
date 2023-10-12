package adminlogic

import (
	"context"

	"github.com/licat233/genzero/examples/all/rpc/admin_pb"
	"github.com/licat233/genzero/examples/all/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestLogic {
	return &TestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Preset methods can be deleted. At least one method needs to be defined in a service
func (l *TestLogic) Test(in *admin_pb.NilReq) (*admin_pb.NilResp, error) {
	// todo: add your logic here and delete this line

	return &admin_pb.NilResp{}, nil
}
