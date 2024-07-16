/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package task

import (
	"bamboo-service/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
)

// dogeCloudTask
//
// # 多吉云 Token 更新任务
//
// 多吉云 Token 更新任务，用于定时更新多吉云 Token
// 每小时更新一次，更新成功后缓存数据
// 缓存数据有效期为该令牌的有效期时限
func (t *task) dogeCloudTask() {
	gtimer.AddTimes(t.ctx, time.Hour, 1, func(_ context.Context) {
		bucket, err := service.DogeCloud().GetAccessTokenAPI(t.ctx)
		if err != nil {
			g.Log().Error(t.ctx, err.Error())
		}
		// 缓存数据重写
		err = g.Redis().HMSet(t.ctx, "global:dc:bucket", gconv.Map(bucket))
		if err != nil {
			g.Log().Error(t.ctx, err.Error())
		}
		// 设置过期时间
		_, err = g.Redis().ExpireAt(t.ctx, "global:dc:bucket", bucket.ExpiredAt.Time)
		if err != nil {
			g.Log().Error(t.ctx, err.Error())
		}
		g.Log().Notice(t.ctx, "[TASK] 多吉云 Token 更新成功")
	})
}
