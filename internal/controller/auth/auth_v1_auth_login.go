package auth

import (
	"bamboo-service/api/auth/v1"
	"bamboo-service/internal/model/dto"
	"context"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/google/uuid"
)

func (c *ControllerV1) AuthLogin(ctx context.Context, req *v1.AuthLoginReq) (res *v1.AuthLoginRes, err error) {

	response := dto.UserInfoDTO{
		UserUUID:   uuid.UUID{},
		Username:   "",
		Email:      "",
		Phone:      "",
		Role:       "",
		Permission: gjson.Json{},
		CreatedAt:  gtime.Time{},
		UpdatedAt:  gtime.Time{},
	}

	return &v1.AuthLoginRes{
		ResponseDTO: bresult.SuccessHasData(ctx, "验证码获取成功", &response),
	}, nil
}
