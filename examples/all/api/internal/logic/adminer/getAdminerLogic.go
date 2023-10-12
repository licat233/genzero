package adminer

import (
	"context"

	"github.com/licat233/genzero/examples/all/api/internal/logic/dataconv"
	"github.com/licat233/genzero/examples/all/api/internal/svc"
	"github.com/licat233/genzero/examples/all/api/internal/types"
	"github.com/licat233/genzero/examples/all/rpc/admin_pb"
	"github.com/licat233/genzero/examples/example_pkg/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAdminerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminerLogic {
	return &GetAdminerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdminerLogic) GetAdminer(req *types.GetAdminerReq) (any, error) {
	rpcResp, err := l.svcCtx.AdminRpc.GetAdminer(l.ctx, &admin_pb.GetAdminerReq{
		Id: req.Id,
	})
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}
	data := dataconv.PbAdminerToApiAdminer(rpcResp.Adminer)

	return respx.DefaultSingleResp(data, nil)
}
