package db

import (
	"command-service/config"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
)

func ConnectMysql(config config.MySql) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DbName,
	)

	gormDb, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}

	return gormDb
}
