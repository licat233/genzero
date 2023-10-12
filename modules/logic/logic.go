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
		//判断是否已使用goctl生成相关代码文件
		exists, err := tools.PathExists(config.C.Logic.Rpc.Dir)
		if err != nil {
			tools.Error("[logic] checking file failed: %s", err)
			return err
		}
		if !exists {
			tools.Warning("[logic] after completing the configuration of the .proto file, use the goctl tool to generate the rpc service file")
		} else {
			if err = l.RpcLogic.Run(); err != nil {
				tools.Error("generate rpc logic code failed.")
				return err
			} else {
				tools.Success("generate rpc logic code success.")
			}
		}
	}
	if config.C.Logic.Api.Status {
		//判断是否已使用goctl生成相关代码文件
		exists, err := tools.PathExists(config.C.Logic.Api.Dir)
		if err != nil {
			tools.Error("[logic] checking file failed: %s", err)
			return err
		}
		if !exists {
			tools.Warning("[logic] after completing the configuration of the .api file, use the goctl tool to generate the api service file")
		} else {
			if err = l.ApiLogic.Run(); err != nil {
				tools.Error("generate api logic code failed.")
				return err
			} else {
				tools.Success("generate api logic code success.")
			}
		}
	}
	// tools.Success("generate all logic code success")
	return
}
