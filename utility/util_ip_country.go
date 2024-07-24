/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package utility

// IPCountryName
//
// # IP国家名称
//
// IP国家名称转换；用于将IP国家名称转换为中文名称；
//
// # 参数
//   - name		国家名称(string)
//
// # 返回
//   - string	国家名称
func IPCountryName(name string) string {
	switch name {
	case "China":
		return "中国"
	default:
		return name
	}
}
