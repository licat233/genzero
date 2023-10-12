package baselogic

import (
	"context"

	"github.com/licat233/genzero/examples/all/rpc/admin_pb"
	"github.com/licat233/genzero/examples/all/rpc/internal/logic/dataconv"
	"github.com/licat233/genzero/examples/all/rpc/internal/svc"
	"github.com/licat233/genzero/examples/example_pkg/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetJwtBlacklistListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetJwtBlacklistListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetJwtBlacklistListLogic {
	return &GetJwtBlacklistListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取jwt黑名单列表
func (l *GetJwtBlacklistListLogic) GetJwtBlacklistList(in *admin_pb.GetJwtBlacklistListReq) (*admin_pb.GetJwtBlacklistListResp, error) {
	pageSize, page, keyword := dataconv.ListReqParams(in.ListReq)
	list, total, err := l.svcCtx.JwtBlacklistModel.FindList(l.ctx, pageSize, page, keyword, dataconv.PbJwtBlacklistToMdJwtBlacklist(in.JwtBlacklist))
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.GetJwtBlacklistListResp{
		JwtBlacklists: dataconv.MdJwtBlacklists2PbJwtBlacklists(list),
		Total:         total,
	}, nil
}
