package server

import (
	"api/internal/controller"

	"github.com/gin-gonic/gin"
)

func CreateServer() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("", controller.Index())

	r.GET("api/getName", controller.GetName())
	r.GET("api/getUsage", controller.GetUsage())

	r.POST("api/post", controller.PostData())

	return r
}
