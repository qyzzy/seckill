package model

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"strconv"
	"time"
)

var (
	cacheSeckillCouponIdPrefix = "cache:seckill:coupon:id:"
)

type (
	couponModel interface {
		CreateCoupon(ctx context.Context, data *Coupon) (int64, error)
		FindOne(ctx context.Context, id int64) (*Coupon, error)
		AddCouponTakeCount(ctx context.Context, id int64) (int64, error)
		AddCouponUsedCount(ctx context.Context, id int64) (int64, error)
	}

	defaultCouponModel struct {
		db    *gorm.DB
		cache *redis.Client
		table string
	}

	Coupon struct {
		gorm.Model
		Name string `gorm:"not null"`
		// 1: fullDiscountCoupon 2: singleProductCoupon 3: productsCoupon 10: noThresholdCoupon
		TypeId int64 `gorm:"not null"`
		// 0: not mutex, stackable with other coupon 1: mutex
		IsMutex int64 `gorm:"noy null"`
		// default all == 0
		ProductId  int64
		ShopId     int64
		CategoryId int64
		// How much money to start using
		// If it is no threshold coupon, with amount == 0
		WithAmount int64 `gorm:"not null"`
		UsedAmount int64 `gorm:"not null"`
		Stock      int64 `gorm:"not null"`
		TakeCount  int64 `gorm:"not null"`
		UsedCount  int64 `gorm:"not null"`
		// Issue coupon start time default value = created_at
		// 0: invalid 1: valid 2: enforce invalid
		Status int64 `gorm:"not null"`
		// 1 : absolute prescription, 2 : relative prescription
		ValidType int64
		// if valid type == 2
		ValidDay int64
		// if valid type == 1
		ValidStartTime time.Time
		ValidEndTime   time.Time
		//
		UpdateUserId int64 `gorm:"not null"`
	}
)

func newCouponModel(db *gorm.DB, cache *redis.Client) *defaultCouponModel {
	return &defaultCouponModel{
		db:    db,
		cache: cache,
		table: "`coupon`",
	}
}

func (m *defaultCouponModel) CreateCoupon(ctx context.Context, data *Coupon) (int64, error) {
	// Transaction
	err := m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var resp *Coupon
		err := tx.Table(m.tableName()).
			Where("name = ?", data.Name).
			First(&resp).Error

		if !(err.Error() == "record not found" || err == nil) {
			return fmt.Errorf("name existed")
		}
		err = tx.Table(m.tableName()).Create(&data).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	seckillCouponKey := fmt.Sprintf("%s%v", cacheSeckillCouponIdPrefix, data.ID)
	seckillCouponStockKey := seckillCouponKey + ":" + "stock:"
	seckillCouponTakeCountKey := seckillCouponKey + ":" + "take_count:"
	seckillCouponUsedCountKey := seckillCouponKey + ":" + "used_count:"

	// Search cache
	jsonData, err := json.Marshal(&data)
	_, err = m.cache.
		Set(ctx, seckillCouponKey, jsonData, time.Hour*24*7).Result()
	if err != nil {
		return 0, nil
	}
	_, err = m.cache.
		Set(ctx, seckillCouponStockKey, data.Stock, time.Hour*24*7).Result()
	if err != nil {
		return 0, nil
	}

	_, err = m.cache.
		Set(ctx, seckillCouponTakeCountKey, 0, time.Hour*24*7).Result()
	if err != nil {
		return 0, nil
	}

	_, err = m.cache.
		Set(ctx, seckillCouponUsedCountKey, 0, time.Hour*24*7).Result()
	if err != nil {
		return 0, nil
	}

	return int64(data.ID), nil
}

func (m *defaultCouponModel) FindOne(ctx context.Context, id int64) (*Coupon, error) {
	seckillCouponKey := fmt.Sprintf("%s%v", cacheSeckillCouponIdPrefix, id)
	var coupon *Coupon

	jsonData, err := m.cache.Get(ctx, seckillCouponKey).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			err = m.db.WithContext(ctx).Select("*").
				Where("id = ?", id).
				First(&coupon).Error
			if err != nil {
				return nil, err
			} else {
				return coupon, nil
			}
		}
		return nil, err
	}
	err = json.Unmarshal([]byte(jsonData), &coupon)
	if err != nil {
		return nil, err
	}
	return coupon, nil
}

func (m *defaultCouponModel) AddCouponTakeCount(ctx context.Context, id int64) (int64, error) {
	seckillCouponStockKey := fmt.Sprintf("%s%v:stock:", cacheSeckillCouponIdPrefix, id)
	seckillCouponTakeCountKey := fmt.Sprintf("%s%v:take_count:", cacheSeckillCouponIdPrefix, id)

	res1, err := m.cache.Get(ctx, seckillCouponStockKey).Result()
	if err != nil {
		return -1, err
	}
	res2, err := m.cache.Get(ctx, seckillCouponTakeCountKey).Result()
	if err != nil {
		return -1, err
	}

	stock, _ := strconv.Atoi(res1)
	takeCount, _ := strconv.Atoi(res2)
	if stock < takeCount+1 {
		return -1, err
	}

	_, err = m.cache.Incr(ctx, seckillCouponTakeCountKey).Result()
	if err != nil {
		return -1, err
	}

	return int64(takeCount + 1), err
}

func (m *defaultCouponModel) AddCouponUsedCount(ctx context.Context, id int64) (int64, error) {
	seckillCouponStockKey := fmt.Sprintf("%s%v:stock:", cacheSeckillCouponIdPrefix, id)
	seckillCouponUsedCountKey := fmt.Sprintf("%s%v:used_count:", cacheSeckillCouponIdPrefix, id)

	res1, err := m.cache.Get(ctx, seckillCouponStockKey).Result()
	if err != nil {
		return -1, err
	}
	res2, err := m.cache.Get(ctx, seckillCouponUsedCountKey).Result()
	if err != nil {
		return -1, err
	}

	stock, _ := strconv.Atoi(res1)
	usedCount, _ := strconv.Atoi(res2)
	if stock < usedCount+1 {
		return -1, err
	}

	_, err = m.cache.Incr(ctx, seckillCouponUsedCountKey).Result()
	if err != nil {
		return -1, err
	}

	return int64(usedCount + 1), err
}

func (m *defaultCouponModel) tableName() string {
	return m.table
}
