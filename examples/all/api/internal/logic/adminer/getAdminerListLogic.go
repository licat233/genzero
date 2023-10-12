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

type GetAdminerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAdminerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminerListLogic {
	return &GetAdminerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdminerListLogic) GetAdminerList(req *types.GetAdminerListReq) (any, error) {
	in := &admin_pb.GetAdminerListReq{
		ListReq: &admin_pb.ListReq{
			PageSize: req.PageSize,
			Page:     req.Page,
			Keyword:  req.Keyword,
		},
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
	rpcResp, err := l.svcCtx.AdminRpc.GetAdminerList(l.ctx, in)
	if err != nil {
		//rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}
	pbList := rpcResp.Adminers
	data := dataconv.PbAdminers2ApiAdminers(pbList)

	return respx.DefaultListResp(data, rpcResp.Total, req.PageSize, req.Page, nil)
}
