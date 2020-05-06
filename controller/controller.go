package controller

import (
	"github.com/gin-gonic/gin"
	"kala-v2/models"
	"net/http"
)

/*
主页
*/
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

/*
创建一个待办事项
*/
func CreateATodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	if err := models.CreateATodo(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

/*
查看所有待办事项
*/
func GetTodoList(c *gin.Context) {
	todoList, err := models.GetTodoList()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

/*
修改某一待办事项
*/
func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "获取不到ID值"})
	} else {
		var todo models.Todo
		c.BindJSON(&todo)
		err := models.UpdateATodo(id, &todo)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		} else {
			c.JSON(http.StatusOK, todo)
		}
	}
}

/*
删除某一待办事项
*/
func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "获取不到ID值"})
	} else {
		err := models.DeleteATodo(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"msg": "删除ID成功"})
		}
	}
}
