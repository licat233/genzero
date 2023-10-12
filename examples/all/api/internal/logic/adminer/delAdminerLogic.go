package adminer

import (
	"context"

	"github.com/licat233/genzero/examples/all/api/internal/svc"
	"github.com/licat233/genzero/examples/all/api/internal/types"
	"github.com/licat233/genzero/examples/all/rpc/admin_pb"
	"github.com/licat233/genzero/examples/example_pkg/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAdminerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelAdminerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAdminerLogic {
	return &DelAdminerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelAdminerLogic) DelAdminer(req *types.DelAdminerReq) (any, error) {
	_, err := l.svcCtx.AdminRpc.DelAdminer(l.ctx, &admin_pb.DelAdminerReq{
		Id: req.Id,
	})
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStatusResp(nil)
}
