package dao

import (
	"github.com/jinzhu/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitMySql() error {
	dsn := "root:root@tcp(127.0.0.1:3306)/gotest1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = DB.DB().Ping()
	if err != nil {
		return err
	}
	return nil
}
