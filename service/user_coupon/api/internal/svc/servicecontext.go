package svc

import (
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"seckill/service/coupon/rpc/coupon"
	"seckill/service/user_coupon/api/internal/config"
	"seckill/service/user_coupon/api/internal/middleware"
	"seckill/service/user_coupon/model"
)

type ServiceContext struct {
	Config          config.Config
	Auth            rest.Middleware
	UserCouponModel model.UserCouponModel
	CouponRpc       coupon.Coupon
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
		Config:          c,
		Auth:            middleware.NewAuthMiddleware().Handle,
		UserCouponModel: model.NewUserCouponModel(conn, cache),
		CouponRpc:       coupon.NewCoupon(zrpc.MustNewClient(c.CouponRpc)),
	}
}
