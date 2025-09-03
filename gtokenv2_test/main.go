package main

import (
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	_ "gtokenv2_test/internal/logic"

	"gtokenv2_test/internal/cmd"
	_ "gtokenv2_test/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
