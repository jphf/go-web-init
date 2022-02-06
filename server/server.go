package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jphf/go-web-init.git/internal/controller"
	"github.com/jphf/go-web-init.git/internal/service"
)

func CreateServer() *gin.Engine {

	getService := service.GetService{}
	postService := service.PostService{}

	pageController := controller.PageController{}
	apiController := controller.ApiController{
		GetService:  getService,
		PostService: postService,
	}

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("", pageController.Index())

	r.GET("api/getName", apiController.GetName())
	r.GET("api/getUsage", apiController.GetUsage())

	r.POST("api/post", apiController.PostData())

	return r
}
