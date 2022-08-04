package svc

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"seckill/service/coupon/api/internal/config"
	"seckill/service/coupon/model"
)

type ServiceContext struct {
	Config                  config.Config
	FullDiscountCouponModel model.FullDiscountCouponModel
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
	cache := redis.NewClient(&redis.Options{
		Addr: c.Redis.Host,
		DB:   1,
	})
	if err != nil {
		log.Println("Gorm open db failed")
	}
	return &ServiceContext{
		Config:                  c,
		FullDiscountCouponModel: model.NewFullDiscountCouponModel(conn, cache),
	}
}
