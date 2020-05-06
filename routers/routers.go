package routers

import (
	"github.com/gin-gonic/gin"
	"kala-v2/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 告诉gin框架静态文件位置
	r.Static("/static", "static")
	// 告诉gin框架html文件位置
	r.LoadHTMLGlob("templates/*")
	// 添加主页路由
	r.GET("/", controller.IndexHandler)
	// v2 版本
	// 待办事项管理
	v1Group := r.Group("v1")
	{
		// 添加一个办事项
		v1Group.POST("/todo", controller.CreateATodo)
		// 查看所有待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		// 修改某一待办事项
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		// 删除某一待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}

	return r
}
