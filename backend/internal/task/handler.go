package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes は task に関するルートを登録する
func RegisterRoutes(rg *gin.RouterGroup) {
	tasksRg := rg.Group("/tasks")
	{
		tasksRg.GET("", getTasks)
		tasksRg.POST("", postTask)
		tasksRg.PUT("/:id", putTask)
		tasksRg.DELETE("/:id", deleteTask)
	}
}

// getTasks はタスク一覧取得
func getTasks(c *gin.Context) {
	title := c.Query("title")
	res := GetAll(title)
	c.JSON(http.StatusOK, res)
}

// postTask はタスク追加
func postTask(c *gin.Context) {
	var t Task
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created := Create(t)
	c.JSON(http.StatusCreated, created)
}

// putTask はタスク更新
func putTask(c *gin.Context) {
	id := c.Param("id")
	var t Task
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, err := Update(id, t)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// deleteTask はタスク削除
func deleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
