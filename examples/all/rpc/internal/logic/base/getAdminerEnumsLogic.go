package baselogic

import (
	"context"

	"github.com/licat233/genzero/examples/all/model"
	"github.com/licat233/genzero/examples/all/rpc/admin_pb"
	"github.com/licat233/genzero/examples/all/rpc/internal/svc"
	"github.com/licat233/genzero/examples/example_pkg/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminerEnumsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAdminerEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminerEnumsLogic {
	return &GetAdminerEnumsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取管理员枚举列表
func (l *GetAdminerEnumsLogic) GetAdminerEnums(in *admin_pb.GetAdminerEnumsReq) (*admin_pb.Enums, error) {
	list, _, err := l.svcCtx.AdminerModel.FindList(l.ctx, -1, -1, "", nil)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	enums := []*admin_pb.Enum{}
	for _, item := range list {
		enums = append(enums, &admin_pb.Enum{
			Label: item.Name,
			Value: item.Id,
		})
	}

	return &admin_pb.Enums{
		Enums: enums,
	}, nil
}
