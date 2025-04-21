package cmd

import (
	"context"
	"proxima-vocabulary/app/word/internal/controller/words"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"google.golang.org/grpc"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "word grpc service",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			c := grpcx.Server.NewConfig()
			
			// 从配置文件读取 gRPC 地址
			address, err := g.Cfg().Get(ctx, "grpc.server.address")
			if err != nil {
				return err
			}
			c.Address = address.String()
			
			c.Options = append(c.Options, []grpc.ServerOption{
				grpcx.Server.ChainUnary(
					grpcx.Server.UnaryValidate,
				)}...,
			)
			s := grpcx.Server.New(c)
			words.Register(s)
			s.Run()
			return nil
		},
	}
)