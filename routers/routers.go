package routers

import (
	"manifest/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 模板文件引用的静态文件
	r.Static("/static", "static")
	// 哪里找模板文件index
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	v1Group := r.Group("v1")
	{
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)
		// 查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)

		// 修改某一个代办事项
		v1Group.PUT("/todo/:id", controller.UpdateTodo)
		// 删除某一个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}
	return r
}
