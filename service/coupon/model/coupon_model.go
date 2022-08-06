package model

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type (
	CouponModel interface {
		couponModel
	}

	customCouponModel struct {
		*defaultCouponModel
	}
)

func NewCouponModel(conn *gorm.DB, cache *redis.Client) CouponModel {
	return &customCouponModel{
		defaultCouponModel: newCouponModel(conn, cache),
	}
}
