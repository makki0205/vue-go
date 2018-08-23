package router

import (
	"github.com/gin-gonic/gin"
)

func apiRouter(api *gin.RouterGroup) {

}

func authApiRouter(auth *gin.RouterGroup) {
	auth.GET("/hoge", func(c *gin.Context) {
		c.JSON(200, "ok")
	})
}
