package dto

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/google/uuid"
)

type UserInfoDTO struct {
	UserUUID   uuid.UUID  `json:"user_uuid"`
	Username   string     `json:"username"`
	Email      string     `json:"email"`
	Phone      string     `json:"phone"`
	Role       string     `json:"role"`
	Permission gjson.Json `json:"permission"`
	CreatedAt  gtime.Time `json:"created_at"`
	UpdatedAt  gtime.Time `json:"updated_at"`
}
