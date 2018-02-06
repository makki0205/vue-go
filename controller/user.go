package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/makki0205/vue-go/model"
	"github.com/makki0205/vue-go/service"
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
	Json(user, c)
}
