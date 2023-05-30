package modules

import (
	"github.com/licat233/genzero/modules/api"
	"github.com/licat233/genzero/modules/logic"
	"github.com/licat233/genzero/modules/model"
	"github.com/licat233/genzero/modules/pb"
)

type Module interface {
	Run() error
}

var _ Module = (*api.ApiModule)(nil)
var _ Module = (*pb.PbModule)(nil)
var _ Module = (*model.ModelModule)(nil)
var _ Module = (*logic.LogicModule)(nil)
