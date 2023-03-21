package main

import (
	"manifest/dao"
	"manifest/models"
	"manifest/routers"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// 连接数据库
	err := dao.InitMySql()
	if err != nil {
		panic(err)
	}
	defer dao.DB.Close()
	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	r := routers.SetupRouter()
	r.Run()
}
