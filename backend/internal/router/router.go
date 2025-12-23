package router

import (
	"github.com/kouyou13/vue-go-todo/backend/internal/category"
	"github.com/kouyou13/vue-go-todo/backend/internal/task"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kouyou13/vue-go-todo/backend/config"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(config.CORSConfig()))
	api := r.Group("/api")

	task.RegisterRoutes(api)
	category.RegisterRoutes(api)

	return r
}
