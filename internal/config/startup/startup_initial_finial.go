/*
 * ***********************************************************
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ***********************************************************
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ***********************************************************
 */

package startup

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/glog"
)

// InitialFinialStartup
//
// # 系统准备完成
//
// 系统初始化完毕，释放资源以及显示 logo
//
// # 请求
//   - ctx			上下文(context.Context)
func InitialFinialStartup(ctx context.Context) {
	glog.Noticef(ctx, "[STARTUP] 系统准备完成")
	fmt.Println("\033[1;35m" + `
	   _  ___           ____             _        
	  | |/_(_)__ ____  / __/__ _____  __(_)______ 
	 _>  </ / _ ` + "`" + `/ _ \_\ \/ -_) __/ |/ / / __/ -_)
	/_/|_/_/\_,_/\___/___/\__/_/  |___/_/\__/\__/ `)
	fmt.Println("\033[32m	   ::: XiaoService :::\033[33m	   	v1.0.0")
	fmt.Println("\033[0m")
}
