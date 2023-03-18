package adminer

import (
	"context"

	"github.com/licat233/genzero/examples/all/api/internal/svc"
	"github.com/licat233/genzero/examples/all/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutAdminerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutAdminerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutAdminerLogic {
	return &PutAdminerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutAdminerLogic) PutAdminer(req *types.PutAdminerReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
