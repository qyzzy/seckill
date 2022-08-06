package svc

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"seckill/service/coupon/model"
	"seckill/service/coupon/rpc/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	CouponModel model.CouponModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Silent),
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Println(err)
	}
	cache := redis.NewClient(&redis.Options{
		Addr: c.Redis.Host,
		DB:   2,
	})
	return &ServiceContext{
		Config:      c,
		CouponModel: model.NewCouponModel(conn, cache),
	}
}
