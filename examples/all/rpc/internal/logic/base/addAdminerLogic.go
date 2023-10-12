package baselogic

import (
	"context"
	"database/sql"
	"time"

	"github.com/licat233/genzero/examples/all/model"
	"github.com/licat233/genzero/examples/all/rpc/admin_pb"
	"github.com/licat233/genzero/examples/all/rpc/internal/svc"
	"github.com/licat233/genzero/examples/example_pkg/errorx"
	"github.com/licat233/genzero/examples/example_pkg/uniqueid"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAdminerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddAdminerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAdminerLogic {
	return &AddAdminerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加管理员
func (l *AddAdminerLogic) AddAdminer(in *admin_pb.AddAdminerReq) (*admin_pb.AddAdminerResp, error) {
	data := &model.Adminer{
		Uuid:         uniqueid.NewUUID(), // 这里的uniqueid包，请自己定义
		Name:         in.Name,
		Avatar:       in.Avatar,
		Passport:     in.Passport,
		Password:     in.Password,
		Email:        in.Email,
		Resume:       sql.NullString{Valid: true, String: in.Resume},
		Status:       in.Status,
		IsSuperAdmin: in.IsSuperAdmin,
		LoginCount:   in.LoginCount,
		LastLogin:    time.Unix(in.LastLogin, 0).Local(),
	}
	if _, err := l.svcCtx.AdminerModel.Insert(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.AddAdminerResp{}, nil
}
