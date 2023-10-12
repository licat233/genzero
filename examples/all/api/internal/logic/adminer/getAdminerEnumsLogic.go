package adminer

import (
	"context"

	"github.com/licat233/genzero/examples/all/api/internal/svc"
	"github.com/licat233/genzero/examples/all/api/internal/types"
	"github.com/licat233/genzero/examples/all/rpc/admin_pb"
	"github.com/licat233/genzero/examples/example_pkg/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminerEnumsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAdminerEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminerEnumsLogic {
	return &GetAdminerEnumsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdminerEnumsLogic) GetAdminerEnums(req *types.GetAdminerEnumsReq) (any, error) {
	rpcResp, err := l.svcCtx.AdminRpc.GetAdminerEnums(l.ctx, &admin_pb.GetAdminerEnumsReq{
		ParentId: req.ParentId,
	})
	if err != nil {
		return nil, err
	}
	data := make([]*types.Enum, 0)
	for _, v := range rpcResp.Enums {
		data = append(data, &types.Enum{
			Label: v.Label,
			Value: v.Value,
		})
	}

	return respx.DefaultSingleResp(data, nil)
}
