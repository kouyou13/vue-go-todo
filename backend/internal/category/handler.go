package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes は category に関するルートを登録する
func RegisterRoutes(rg *gin.RouterGroup) {
	tasksRg := rg.Group("/api/categories")
	{
		tasksRg.GET("", getCategories)
		tasksRg.POST("", postCategories)
		tasksRg.PUT("/:id", putCategories)
		tasksRg.DELETE("/:id", deleteCategories)
	}
}

// getCategories はカテゴリー一覧取得
func getCategories(c *gin.Context) {
	res := GetAll()
	c.JSON(http.StatusOK, res)
}

// postCategories はカテゴリー追加
func postCategories(c *gin.Context) {
	var t Category
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created := Create(t)
	c.JSON(http.StatusCreated, created)
}

// putCategories はカテゴリー更新
func putCategories(c *gin.Context) {
	id := c.Param("id")
	var t Category
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

// deleteCategories はカテゴリー削除
func deleteCategories(c *gin.Context) {
	id := c.Param("id")
	if err := Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
