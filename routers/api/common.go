package api

import (
	"gin-vue-template/pkg/constants"
	"gin-vue-template/pkg/e"
	"gin-vue-template/pkg/logic"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *API) GetPageAuth(c *gin.Context) {
	info, ok := c.Get(constants.AuthKey)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}
	loginInfo := info.(*logic.LoginInfo)

	resp := make(map[string]bool)
	for _, a := range constants.PageAuths {
		if _, ok := a.Roles[constants.Role(loginInfo.Role)]; ok {
			resp[a.Page] = true
		}
	}

	e.GinOK(c, resp)
}
