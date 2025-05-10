package dto

// UserInfoDTO 表示用户信息的数据传输对象。
//
// 包含用户的基础信息，如唯一标识符、用户名、邮箱、手机号码、角色及昵称，用于数据传输过程中的封装与解析。
type UserInfoDTO struct {
	UserUuid string `json:"user_uuid"          orm:"user_uuid"          description:"用户唯一标识符"` // 用户唯一标识符
	Username string `json:"username"           orm:"username"           description:"用户名"`     // 用户名
	Email    string `json:"email"              orm:"email"              description:"电子邮箱"`    // 电子邮箱
	Phone    string `json:"phone"              orm:"phone"              description:"手机号码"`    // 手机号码
	Role     string `json:"role"               orm:"role"               description:"用户角色"`    // 用户角色
	Nickname string `json:"nickname"           orm:"nickname"           description:"用户昵称"`    // 用户昵称
}
