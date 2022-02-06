package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jphf/go-web-init.git/internal/service"
)

type ApiController struct {
	GetService  service.GetService
	PostService service.PostService
}

func (controller ApiController) GetName() func(*gin.Context) {
	getService := controller.GetService

	return func(c *gin.Context) {
		name := getService.GetName()
		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	}
}

func (controller ApiController) GetUsage() func(*gin.Context) {
	getService := controller.GetService

	return func(c *gin.Context) {
		usage := getService.GetUsage()
		c.JSON(http.StatusOK, gin.H{
			"usage": usage,
		})
	}
}

func (controller ApiController) PostData() func(*gin.Context) {
	postService := controller.PostService

	return func(c *gin.Context) {
		name := c.PostForm("name")
		usage := c.PostForm("usage")
		postService.PostData(name, usage)

		c.JSON(http.StatusOK, gin.H{
			"name":  name,
			"usage": usage,
		})
	}
}
