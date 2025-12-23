package router

import (
	"github.com/kouyou13/vue-go-todo/backend/internal/category"
	"github.com/kouyou13/vue-go-todo/backend/internal/task"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")

	task.RegisterRoutes(api)
	category.RegisterRoutes(api)

	return r
}
