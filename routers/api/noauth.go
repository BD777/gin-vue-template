package api

import (
	"gin-vue-template/pkg/constants"
	"gin-vue-template/pkg/e"
	"gin-vue-template/pkg/infra/logger"

	"github.com/gin-gonic/gin"
)

func (a *API) GetLoginInfo(c *gin.Context) {
	ctx := NewGinContextHelper(c).Context()

	token, err := c.Cookie(constants.TokenKey)
	if err != nil {
		e.GinR(c, e.RetCodeNotLogin, nil)
		return
	}

	info, err := a.logic.GetLoginInfo(ctx, token)
	if err != nil {
		logger.Error(ctx, "get login info failed: %v", err)
		e.GinR(c, e.RetCodeNotLogin, nil)
		return
	}

	e.GinOK(c, info)
}

func (a *API) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var data struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		logger.Error(ctx, "bind json failed: %v", err)
		e.GinR(c, e.RetCodeInvalidParams, nil)
		return
	}

	info, err := a.logic.Login(ctx, data.Username, data.Password)
	if err != nil {
		logger.Error(ctx, "login failed: %v", err)
		e.GinR(c, e.RetCodeLoginFail, nil)
		return
	}

	c.SetCookie(constants.TokenKey, info.Token, 3600*24, "/", "localhost", false, true)
	e.GinOK(c, info)
}
