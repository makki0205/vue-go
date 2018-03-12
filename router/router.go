package router

import (
	"github.com/gin-gonic/gin"
	"github.com/chotchy-inc/PATRAProductAPI/middleware"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/js", "./public/js")
	r.Static("/image", "./public/image")
	r.Static("/css", "./public/css")

	r.LoadHTMLGlob("view/*")
	r.NoRoute(func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	api := r.Group("/api")
	apiRouter(api)
	auth := api.Group("")
	auth.Use(middleware.Jwt("hogehoge", 3600*24*365))
	authApiRouter(auth)
	return r

}
