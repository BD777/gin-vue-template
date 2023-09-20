package api

import (
	"gin-vue-template/constants"
	"gin-vue-template/pkg/e"
	"gin-vue-template/pkg/logic"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLoginInfo(c *gin.Context) {
	ctx := c.Request.Context()

	token, err := c.Cookie(constants.TokenKey)
	if err != nil {
		c.JSON(http.StatusOK, e.R(e.NOT_LOGIN, nil))
		return
	}

	info, err := logic.GetLoginInfo(ctx, token)
	if err != nil {
		log.Printf("get login info failed: %v", err)
		c.JSON(http.StatusOK, e.R(e.NOT_LOGIN, nil))
		return
	}

	c.JSON(http.StatusOK, e.R(e.OK, info))
}

func Login(c *gin.Context) {
	ctx := c.Request.Context()

	var data struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		log.Printf("bind json failed: %v", err)
		c.JSON(http.StatusOK, e.R(e.INVALID_PARAMS, nil))
		return
	}

	info, err := logic.Login(ctx, data.Username, data.Password)
	if err != nil {
		log.Printf("login failed: %v", err)
		c.JSON(http.StatusOK, e.R(e.LOGIN_FAIL, nil))
		return
	}

	c.SetCookie(constants.TokenKey, info.Token, 3600*24, "/", "localhost", false, true)
	c.JSON(http.StatusOK, e.R(e.OK, info))
}
