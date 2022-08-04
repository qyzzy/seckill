package model

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
)

var (
	cacheSeckillFullDiscountCouponStock = "cache:seckill:user:phoneNumber:"
)

type (
	fullDiscountCouponModel interface {
		Insert(ctx context.Context, data *FullDiscountCoupon) (int64, error)
		FindOneById(ctx context.Context, id int64) (*FullDiscountCoupon, error)
	}

	defaultFullDiscountCouponModel struct {
		conn  *gorm.DB
		cache *redis.Client
		table string
	}

	FullDiscountCoupon struct {
		gorm.Model
		Name           string    `gorm:"not null"`
		WithAmount     int       `gorm:"not null"`
		UsedAmount     int       `gorm:"not null"`
		Stock          int       `gorm:"not null"`
		StartTime      time.Time `gorm:"not null"`
		EndTime        time.Time `gorm:"not n ull"`
		Status         int       `gorm:"not null"`
		ValidStartTime time.Time `gorm:"not null"`
		ValidEndTime   time.Time `gorm:"not null"`
	}
)

func newFullDiscountCouponModel(conn *gorm.DB, cache *redis.Client) *defaultFullDiscountCouponModel {
	return &defaultFullDiscountCouponModel{
		conn:  conn,
		cache: cache,
		table: "`full_discount_coupon`",
	}
}

func (m *defaultFullDiscountCouponModel) Insert(ctx context.Context, data *FullDiscountCoupon) (int64, error) {
	err := m.conn.WithContext(ctx).
		Table(m.table).Create(&data).Error
	if err != nil {
		return 0, err
	}
	return int64(data.ID), err
}

func (m *defaultFullDiscountCouponModel) FindOneById(ctx context.Context, id int64) (*FullDiscountCoupon, error) {
	var fullDiscountCoupon *FullDiscountCoupon
	err := m.conn.WithContext(ctx).Table(m.table).Select("*").Where("id = ?", id).First(&fullDiscountCoupon)
	if err != nil {
		return nil, nil
	}
	return fullDiscountCoupon, nil
}

func (m *defaultFullDiscountCouponModel) DeductStockByCache(ctx context.Context, id int64) (int, error) {
	seckillCouponKey := fmt.Sprintf("%s%v", cacheSeckillFullDiscountCouponStock, id)
	res, err := m.cache.Get(ctx, seckillCouponKey).Result()
	if err != nil {
		log.Println(err)
		return -1, err
	}
	stock, _ := strconv.Atoi(res)
	m.cache.Decr(ctx, seckillCouponKey)
	return stock - 1, err
}
