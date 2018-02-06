package router

import (
	"github.com/gin-gonic/gin"
	. "github.com/makki0205/vue-go/controller"
	"github.com/makki0205/vue-go/middleware"
)

func apiRouter(api *gin.RouterGroup) {
	api.POST("/signup", User.Create)
	api.POST("/signin", middleware.Login)

}

func authApiRouter(auth *gin.RouterGroup) {
	auth.GET("/hoge", func(c *gin.Context) {
		c.JSON(200, "ok")
	})
}
