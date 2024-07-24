/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package dto

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// AlbumInfoDTO
//
// # 图库信息
//
// 图库信息数据传输对象
//
// # 返回
//   - AlbumUuid		图库ID(string)
//   - User				用户(string)
//   - Name				图库名称(string)
//   - DisplayName		图库展示名称(string)
//   - Description		图库描述(string)
//   - Cover			封面(string)
//   - Pattern			当前模式(string)
//   - Visible			是否对外可见(bool)
//   - MatchAddress		匹配地址(*gjson.Json)
//   - CreatedAt		创建时间(*gtime.Time)
//   - UpdatedAt		更新时间(*gtime.Time)
type AlbumInfoDTO struct {
	AlbumUuid    string      `json:"album_uuid" dc:"图库ID"`
	User         string      `json:"user" dc:"用户"`
	Name         string      `json:"name" dc:"图库名称"`
	DisplayName  string      `json:"display_name" dc:"图库展示名称"`
	Description  string      `json:"description" dc:"图库描述"`
	Cover        string      `json:"cover" dc:"封面"`
	Pattern      string      `json:"pattern" dc:"当前模式"`
	Visible      bool        `json:"visible" dc:"是否对外可见"`
	MatchAddress *gjson.Json `json:"match_address" dc:"匹配地址"`
	CreatedAt    *gtime.Time `json:"created_at" dc:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updated_at" dc:"更新时间"`
}
