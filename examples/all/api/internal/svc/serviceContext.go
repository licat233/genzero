package svc

import (
	"github.com/licat233/genzero/examples/all/api/internal/config"
	"github.com/licat233/genzero/examples/all/model"
	"github.com/licat233/genzero/examples/all/rpc/admin_pb"
)

type ServiceContext struct {
	Config   config.Config
	AdminRpc admin_pb.BaseClient
	model.AdminerModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
