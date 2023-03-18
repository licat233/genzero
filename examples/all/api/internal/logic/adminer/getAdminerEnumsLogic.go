package adminer

import (
	"context"

	"github.com/licat233/genzero/examples/all/api/internal/svc"
	"github.com/licat233/genzero/examples/all/api/internal/types"

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

func (l *GetAdminerEnumsLogic) GetAdminerEnums(req *types.GetAdminerEnumsReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
