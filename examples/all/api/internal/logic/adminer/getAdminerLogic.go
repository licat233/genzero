package adminer

import (
	"context"

	"github.com/licat233/genzero/examples/all/api/internal/svc"
	"github.com/licat233/genzero/examples/all/api/internal/types"

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

func (l *GetAdminerLogic) GetAdminer(req *types.GetAdminerReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
