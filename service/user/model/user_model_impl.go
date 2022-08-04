package model

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	cacheSeckillUserPhoneNumber = "cache:seckill:user:phoneNumber:"
)

type (
	userModel interface {
		RegisterByPhoneNumber(ctx context.Context, data *User) (int64, int64, int64, error)
		Login(ctx context.Context, phoneNumber, password string) (int64, int64, error)
	}

	defaultUserModel struct {
		db    *gorm.DB
		cache *redis.Client
		table string
	}

	User struct {
		gorm.Model
		Name        string `gorm:"not null"`
		PhoneNumber string `gorm:"not null"`
		Email       string
		QqNumber    string
		AuthorityId int64  `gorm:"not null"`
		Password    string `gorm:"not null"`
		Age         int64  `gorm:"not null"`
		Gender      int64  `gorm:"not null"`
		Avatar      string `gorm:"not null"`
		Status      int64  `gorm:"not null"`
	}
)

func newUserModel(conn *gorm.DB, c *redis.Client) *defaultUserModel {
	return &defaultUserModel{
		db:    conn,
		cache: c,
		table: "`user`",
	}
}

func (m *defaultUserModel) RegisterByPhoneNumber(ctx context.Context, data *User) (int64, int64, int64, error) {
	seckillUserKey := fmt.Sprintf("%s%v", cacheSeckillUserPhoneNumber, data.PhoneNumber)
	res, err := m.cache.Get(ctx, seckillUserKey).Result()
	if res != "" {
		return 0, 0, 0, fmt.Errorf("phone number registered")
	}
	if err.Error() != "redis: nil" {
		return 0, 0, 0, err
	}
	err = m.db.Transaction(func(tx *gorm.DB) error {
		var resp *User
		err := m.db.Table(m.tableName()).Select("phone_number").
			Where("phone_number = ?", data.PhoneNumber).
			First(&resp).Error
		if err.Error() != "record not found" {
			return err
		}
		if resp.PhoneNumber == data.PhoneNumber {
			return fmt.Errorf("phone number registerd")
		}
		err = m.db.Table(m.tableName()).Create(&data).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, 0, 0, err
	}
	res, err = m.cache.Set(ctx, seckillUserKey, data.ID, 60000).Result()
	if err != nil {
		return 0, 0, 0, err
	}
	return int64(data.ID), data.AuthorityId, data.Status, nil
}

func (m *defaultUserModel) Login(ctx context.Context, phoneNumber, password string) (int64, int64, error) {
	var user *User
	err := m.db.WithContext(ctx).Table(m.tableName()).
		Select("id, password, authority_id").
		Where("phone_number = ?", phoneNumber).First(&user).Error
	if err != nil {
		return 0, 0, fmt.Errorf("user not register")
	}

	if user.Password != password {
		return 0, 0, fmt.Errorf("password wrong")
	}

	return int64(user.ID), user.AuthorityId, nil
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
