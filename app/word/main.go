package main

import (
	_ "proxima-vocabulary/app/word/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2" // 导入 MySQL 驱动
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	"proxima-vocabulary/app/word/internal/cmd"
)

func main() {
    var ctx = gctx.New()
    conf, err := g.Cfg("etcd").Get(ctx, "etcd.address")
    if err != nil {
        panic(err)
    }

    var address = conf.String()
    g.Log().Info(ctx, "Connecting to etcd:", address)
    registry := etcd.New(address)
    grpcx.Resolver.Register(registry)
    g.Log().Info(ctx, "Service registered to etcd")

    cmd.Main.Run(ctx)
}
