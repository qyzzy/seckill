package model

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type (
	couponActivityModel interface {
		CreateCouponActivity(ctx context.Context, data *CouponActivity) (int64, error)
	}

	defaultCouponActivityModel struct {
		db    *gorm.DB
		cache *redis.Client
		table string
	}

	CouponActivity struct {
		gorm.Model
		Name         string `gorm:"not null"`
		CouponTypeId int64  `gorm:"not null"`
	}
)

func newCouponActivityModel(db *gorm.DB, cache *redis.Client) *defaultCouponActivityModel {
	return &defaultCouponActivityModel{
		db:    db,
		cache: cache,
		table: "`coupon_activity`",
	}
}

func (m *defaultCouponActivityModel) CreateCouponActivity(ctx context.Context, data *CouponActivity) (int64, error) {
	err := m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var resp *CouponActivity
		err := m.db.Table(m.tableName()).
			Where("name = ?", data.Name).
			First(&resp).Error
		if err.Error() != "record not found" {
			return err
		}
		if resp.Name == data.Name {
			return fmt.Errorf("activity existed")
		}
		err = m.db.Table(m.tableName()).Create(&data).Error
		if err != nil {
			return nil
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return int64(data.ID), nil
}

func (m *defaultCouponActivityModel) tableName() string {
	return m.table
}
