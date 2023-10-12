package main

import (
	"flag"
	"fmt"

	"github.com/licat233/genzero/examples/all/rpc/admin_pb"
	"github.com/licat233/genzero/examples/all/rpc/internal/config"
	adminServer "github.com/licat233/genzero/examples/all/rpc/internal/server/admin"
	baseServer "github.com/licat233/genzero/examples/all/rpc/internal/server/base"
	"github.com/licat233/genzero/examples/all/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/admin.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		admin_pb.RegisterBaseServer(grpcServer, baseServer.NewBaseServer(ctx))
		admin_pb.RegisterAdminServer(grpcServer, adminServer.NewAdminServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
