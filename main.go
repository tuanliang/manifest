package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB  *gorm.DB
	err error
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMySql() error {
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
func main() {
	// 连接数据库
	err := initMySql()
	if err != nil {
		panic(err)
	}
	defer DB.Close()
	// 模型绑定
	DB.AutoMigrate(&Todo{})

	r := gin.Default()

	// 模板文件引用的静态文件
	r.Static("/static", "static")
	// 哪里找模板文件index
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := r.Group("v1")
	{
		// 添加
		v1Group.POST("/todo", func(c *gin.Context) {
			// 1.从请求把数据拿出来
			var todo Todo
			c.BindJSON(&todo)
			// 2.存入数据库
			// 3.返回相应
			if err = DB.Create(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		// 查看所有的待办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			var todoList []Todo
			if err = DB.Find(&todoList).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todoList)
			}
		})
		// 查看某一个待办事项
		v1Group.GET("/todo:id", func(c *gin.Context) {

		})
		// 修改某一个代办事项
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
				return
			}
			var todo Todo
			if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
			c.BindJSON(&todo)
			if err = DB.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		// 删除某一个待办事项
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
				return
			}
			if err = DB.Where("id=?", id).Delete(Todo{}).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{id: "delete"})
			}
		})
	}
	r.Run()
}
