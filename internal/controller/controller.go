package controller

import (
	"api/internal/util/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index() func(*gin.Context) {
	return func(c *gin.Context) {

		conf := config.Config

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"name":  conf.Name,
			"usage": conf.Usage,
		})
	}
}
