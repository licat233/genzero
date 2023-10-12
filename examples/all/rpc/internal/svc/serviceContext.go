package svc

import (
	"github.com/licat233/genzero/examples/all/model"
	"github.com/licat233/genzero/examples/all/rpc/internal/config"
)

type ServiceContext struct {
	Config            config.Config
	AdminerModel      model.AdminerModel //管理员
	JwtBlacklistModel model.JwtBlacklistModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
