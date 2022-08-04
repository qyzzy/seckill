package model

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// No threshold coupon

type (
	FullDiscountCouponModel interface {
		fullDiscountCouponModel
	}

	customFullDiscountCouponModel struct {
		*defaultFullDiscountCouponModel
	}
)

func NewFullDiscountCouponModel(conn *gorm.DB, cache *redis.Client) FullDiscountCouponModel {
	return &customFullDiscountCouponModel{
		defaultFullDiscountCouponModel: newFullDiscountCouponModel(conn, cache),
	}
}
