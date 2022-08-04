package model

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn *gorm.DB, c *redis.Client) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c),
	}
}
