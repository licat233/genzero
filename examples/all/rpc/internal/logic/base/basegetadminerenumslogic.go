package baselogic

import (
	"context"

	"github.com/licat233/genzero/examples/all/rpc/admin_pb"
	"github.com/licat233/genzero/examples/all/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BaseGetAdminerEnumsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseGetAdminerEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseGetAdminerEnumsLogic {
	return &BaseGetAdminerEnumsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取管理员枚举列表
func (l *BaseGetAdminerEnumsLogic) BaseGetAdminerEnums(in *admin_pb.GetAdminerEnumsReq) (*admin_pb.Enums, error) {
	// todo: add your logic here and delete this line

	return &admin_pb.Enums{}, nil
}
