package baselogic

import (
	"context"
	"time"

	"github.com/licat233/genzero/examples/all/model"
	"github.com/licat233/genzero/examples/all/rpc/admin_pb"
	"github.com/licat233/genzero/examples/all/rpc/internal/svc"
	"github.com/licat233/genzero/examples/example_pkg/errorx"
	"github.com/licat233/genzero/examples/example_pkg/uniqueid"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddJwtBlacklistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddJwtBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddJwtBlacklistLogic {
	return &AddJwtBlacklistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加jwt黑名单
func (l *AddJwtBlacklistLogic) AddJwtBlacklist(in *admin_pb.AddJwtBlacklistReq) (*admin_pb.AddJwtBlacklistResp, error) {
	data := &model.JwtBlacklist{
		AdminerId: in.AdminerId,
		Uuid:      uniqueid.NewUUID(), // 这里的uniqueid包，请自己定义
		Token:     in.Token,
		Platform:  in.Platform,
		Ip:        in.Ip,
		ExpireAt:  time.Unix(in.ExpireAt, 0).Local(),
	}
	if _, err := l.svcCtx.JwtBlacklistModel.Insert(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.AddJwtBlacklistResp{}, nil
}
