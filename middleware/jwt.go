package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/makki0205/gojwt"
	"github.com/chotchy-inc/PATRAProductAPI/controller"
	"github.com/chotchy-inc/PATRAProductAPI/model"
	"github.com/chotchy-inc/PATRAProductAPI/service"
)

func Jwt(salt string, exp int) gin.HandlerFunc {
	jwt.SetSalt(salt)
	jwt.SetExp(exp)
	return func(c *gin.Context) {
		token := c.Query("token")
		claims, err := jwt.Decode(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"err": err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("user_id", claims["id"])
		c.Set("email", claims["email"])
		c.Next()
	}
}

func Login(c *gin.Context) {
	var req model.User
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	user, ok := service.User.Login(req.Email, req.Password)
	if !ok {
		controller.BatRequest("ログイン失敗", c)
	}
	claims := map[string]string{
		"id":    strconv.Itoa(int(user.ID)),
		"email": user.Email,
	}
	token := jwt.Generate(claims)
	controller.Json(gin.H{"token": token}, c)
}
