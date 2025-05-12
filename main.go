package main

import (
	"bamboo-service/internal/config/cmd"
	"bamboo-service/internal/config/setup"
	_ "bamboo-service/internal/logic"
	_ "bamboo-service/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
)

func main() {
	initCtx := gctx.GetInitCtx()

	err := gtime.SetTimeZone("Asia/Shanghai")
	if err != nil {
		panic(err)
	}

	su := setup.New(initCtx)
	su.CheckUserSuperAdminExist()
	su.GetVariableSetup()

	cmd.Main.Run(initCtx)
}
