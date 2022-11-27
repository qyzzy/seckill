package model

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type (
	CouponActivityModel interface {
		couponActivityModel
	}

	customCouponActivityModel struct {
		*defaultCouponActivityModel
	}
)

func NewCouponActivityModel(conn *gorm.DB, cache *redis.Client) CouponActivityModel {
	return &customCouponActivityModel{
		defaultCouponActivityModel: newCouponActivityModel(conn, cache),
	}
}
