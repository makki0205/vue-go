package router

import (
	"github.com/gin-gonic/gin"
	. "github.com/makki0205/vue-go/controller"
	"github.com/makki0205/vue-go/middleware"
)

func apiRouter(api *gin.RouterGroup) {
	api.POST("/signup", User.Create)
	api.POST("/signin", middleware.Login)
	auth := api.Group("")
	auth.Use(middleware.Jwt("hogehoge", 3600*24*365))
}
