package model

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type (
	UserCouponModel interface {
		userCouponModel
	}

	customUserCouponModel struct {
		*defaultUserCouponModel
	}
)

func NewUserCouponModel(conn *gorm.DB, c *redis.Client) UserCouponModel {
	return &customUserCouponModel{
		defaultUserCouponModel: newUserCouponModel(conn, c),
	}
}
