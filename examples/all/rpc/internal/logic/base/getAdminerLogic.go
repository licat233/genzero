package baselogic

import (
	"context"

	"github.com/licat233/genzero/examples/all/model"
	"github.com/licat233/genzero/examples/all/rpc/admin_pb"
	"github.com/licat233/genzero/examples/all/rpc/internal/logic/dataconv"
	"github.com/licat233/genzero/examples/all/rpc/internal/svc"
	"github.com/licat233/genzero/examples/example_pkg/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAdminerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminerLogic {
	return &GetAdminerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取管理员
func (l *GetAdminerLogic) GetAdminer(in *admin_pb.GetAdminerReq) (*admin_pb.GetAdminerResp, error) {
	res, err := l.svcCtx.AdminerModel.FindById(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.GetAdminerResp{
		Adminer: dataconv.MdAdminerToPbAdminer(res),
	}, nil
}
