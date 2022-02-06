package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jphf/go-web-init.git/internal/util/config"
)

type PageController struct{}

func (controller PageController) Index() func(*gin.Context) {
	return func(c *gin.Context) {

		conf := config.Config

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"name":  conf.Name,
			"usage": conf.Usage,
		})
	}
}
