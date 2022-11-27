package model

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	userCouponModel interface {
		CreateUserCoupon(ctx context.Context, data *UserCoupon) (string, int64, error)
	}

	defaultUserCouponModel struct {
		db    *gorm.DB
		cache *redis.Client
		table string
	}

	UserCoupon struct {
		gorm.Model
		// coupon order id
		Uuid     uuid.UUID `gorm:"not null"`
		UserId   int64     `gorm:"not null"`
		CouponId int64     `gorm:"not null"`
		// 0: used 1: not use
		Status int64 `gorm:"not null"`
	}
)

func newUserCouponModel(conn *gorm.DB, c *redis.Client) *defaultUserCouponModel {
	return &defaultUserCouponModel{
		db:    conn,
		cache: c,
		table: "`user_coupon`",
	}
}

func (m *defaultUserCouponModel) CreateUserCoupon(ctx context.Context, data *UserCoupon) (string, int64, error) {
	err := m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var resp *UserCoupon
		err := tx.Table(m.tableName()).
			Where("user_id = ? and coupon_id = ?", data.UserId, data.CouponId).
			First(&resp).Error
		if !(err.Error() == "record not found" || err == nil) {
			return fmt.Errorf("coupon existed")
		}

		newUuid := uuid.New()
		data.Uuid = newUuid
		data.Status = 1
		err = tx.Table(m.tableName()).Create(&data).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return "", 0, err
	}
	return data.Uuid.String(), data.Status, nil
}

func (m *defaultUserCouponModel) tableName() string {
	return m.table
}
