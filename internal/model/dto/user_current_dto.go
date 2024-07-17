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

// UserCurrentDTO
//
// # 用户当前信息
//
// 用户当前信息，用于用户当前信息展示；
//
// # 参数
//   - UUID				用户唯一令牌(string)
//   - Username			用户名(string)
//   - Phone			用户手机号(string)
//   - Email			用户邮箱(string)
//   - Ruuid			角色组(string)
//   - Vuuid			会员主键(string)
//   - Description	个人简述(string)
type UserCurrentDTO struct {
	UUID        string `json:"uuid"         orm:"uuid"         dc:"用户唯一令牌"` // 用户唯一令牌
	Username    string `json:"username"     orm:"username"     dc:"用户名"`    // 用户名
	Phone       string `json:"phone"        orm:"phone"        dc:"用户手机号"`  // 用户手机号
	Email       string `json:"email"        orm:"email"        dc:"用户邮箱"`   // 用户邮箱
	Ruuid       string `json:"ruuid"        orm:"ruuid"        dc:"角色组"`    // 角色组
	Vuuid       string `json:"vuuid"        orm:"vuuid"        dc:"会员主键"`   // 会员主键
	Description string `json:"description"  orm:"description"  dc:"个人简述"`   // 个人简述
}
