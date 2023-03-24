package core

import (
	"github.com/licat233/genzero/core/api"
	"github.com/licat233/genzero/core/logic"
	"github.com/licat233/genzero/core/model"
	"github.com/licat233/genzero/core/pb"
)

type Core interface {
	Run() error
}

var _ (Core) = (*api.ApiCore)(nil)
var _ (Core) = (*pb.PbCore)(nil)
var _ (Core) = (*model.ModelCore)(nil)
var _ (Core) = (*logic.LogicCore)(nil)
