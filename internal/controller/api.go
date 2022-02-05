package controller

import (
	"api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetName() func(*gin.Context) {
	return func(c *gin.Context) {
		name := service.GetName()
		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	}
}

func GetUsage() func(*gin.Context) {
	return func(c *gin.Context) {
		usage := service.GetUsage()
		c.JSON(http.StatusOK, gin.H{
			"usage": usage,
		})
	}
}

func PostData() func(*gin.Context) {
	return func(c *gin.Context) {
		name := c.PostForm("name")
		usage := c.PostForm("usage")
		service.PostData(name, usage)

		c.JSON(http.StatusOK, gin.H{
			"name":  name,
			"usage": usage,
		})
	}
}
