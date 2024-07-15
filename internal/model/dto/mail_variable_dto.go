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

// MailVariableDTO
//
// # 邮件变量 DTO
//
// 邮件变量 DTO，用于邮件变量的数据传输对象；
//
// # 参数
//   - Key		变量名(string)
//   - Value	变量值(string)
type MailVariableDTO struct {
	Key   string `json:"value_name" summary:"变量名"`
	Value string `json:"value" summary:"变量值"`
}
