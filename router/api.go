package router

import (
	. "github.com/aykikr/recipo/controller"
	"github.com/gin-gonic/gin"
)

func apiRouter(api *gin.RouterGroup) {
	api.POST("/signup", User.Create)
	api.POST("/signin", User.Login)
}
