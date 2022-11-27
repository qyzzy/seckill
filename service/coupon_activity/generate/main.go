package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"seckill/service/coupon_activity/model"
)

func main() {
	conn, err := gorm.Open(mysql.Open("root:45568769zby@tcp(127.0.0.1:3306)/seckill"), &gorm.Config{
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

	_ = conn.AutoMigrate(
		&model.CouponActivity{},
	)
}
