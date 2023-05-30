package logic

import (
	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/modules/logic/internal/apilogic"
	"github.com/licat233/genzero/modules/logic/internal/rpclogic"
	"github.com/licat233/genzero/tools"
)

type LogicModule struct {
	RpcLogic *rpclogic.RpcLogic
	ApiLogic *apilogic.ApiLogic
}

func New() *LogicModule {
	return &LogicModule{
		RpcLogic: rpclogic.New(),
		ApiLogic: apilogic.New(),
	}
}

func (l *LogicModule) Run() (err error) {
	//务必先执行 Rpc
	if config.C.Logic.Rpc.Status {
		if err = l.RpcLogic.Run(); err != nil {
			tools.Error("generate rpc logic code failed.")
			return err
		} else {
			tools.Success("generate rpc logic code success.")
		}
	}
	if config.C.Logic.Api.Status {
		if err = l.ApiLogic.Run(); err != nil {
			tools.Error("generate api logic code failed.")
			return err
		} else {
			tools.Success("generate api logic code success.")
		}
	}
	// tools.Success("generate all logic code success")
	return
}
