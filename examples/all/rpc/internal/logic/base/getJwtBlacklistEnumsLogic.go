package baselogic

import (
	"context"

	"github.com/licat233/genzero/examples/all/rpc/admin_pb"
	"github.com/licat233/genzero/examples/all/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetJwtBlacklistEnumsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetJwtBlacklistEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetJwtBlacklistEnumsLogic {
	return &GetJwtBlacklistEnumsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取jwt黑名单枚举列表
func (l *GetJwtBlacklistEnumsLogic) GetJwtBlacklistEnums(in *admin_pb.GetJwtBlacklistEnumsReq) (*admin_pb.Enums, error) {
	// todo: add your logic here and delete this line

	return &admin_pb.Enums{}, nil
}
