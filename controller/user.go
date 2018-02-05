package controller

import (
	"net/http"

	"github.com/aykikr/recipo/model"
	"github.com/aykikr/recipo/service"
	"github.com/gin-gonic/gin"
)

var User = userimpl{}

type userimpl struct {
}

func (u *userimpl) Create(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	user = service.User.Store(user)
	json(user, c)
}

func (u *userimpl) Login(c *gin.Context) {
	var req model.User
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	user, ok := service.User.Login(req.Email, req.Password)
	if !ok {
		BatRequest("ログイン失敗", c)
	}
	token := service.Token.Create(user.ID)
	json(token, c)
}
