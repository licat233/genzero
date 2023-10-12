package adminer

import (
	"context"

	"github.com/licat233/genzero/examples/all/api/internal/svc"
	"github.com/licat233/genzero/examples/all/api/internal/types"
	"github.com/licat233/genzero/examples/all/rpc/admin_pb"
	"github.com/licat233/genzero/examples/example_pkg/respx"

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

func (l *PutAdminerLogic) PutAdminer(req *types.PutAdminerReq) (any, error) {
	in := &admin_pb.PutAdminerReq{
		Adminer: &admin_pb.Adminer{
			Id:           req.Id,
			Uuid:         req.Uuid,
			Name:         req.Name,
			Avatar:       req.Avatar,
			Passport:     req.Passport,
			Password:     req.Password,
			Email:        req.Email,
			Resume:       req.Resume,
			Status:       req.Status,
			IsSuperAdmin: req.IsSuperAdmin,
			LoginCount:   req.LoginCount,
			LastLogin:    req.LastLogin,
		},
	}
	if _, err := l.svcCtx.AdminRpc.PutAdminer(l.ctx, in); err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStatusResp(nil)
}
