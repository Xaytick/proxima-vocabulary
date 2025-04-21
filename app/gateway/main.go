package main

import (
	_ "proxima-vocabulary/app/gateway/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"proxima-vocabulary/app/gateway/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
