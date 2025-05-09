package main

import (
	_ "bamboo-service/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/os/gctx"

	"bamboo-service/internal/cmd"
)

func main() {
	initCtx := gctx.GetInitCtx()

	err := gtime.SetTimeZone("Asia/Shanghai")
	if err != nil {
		panic(err)
	}

	cmd.Main.Run(initCtx)
}
