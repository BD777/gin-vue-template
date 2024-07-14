package api

import (
	"gin-vue-template/pkg/constants"
	"gin-vue-template/pkg/e"
	"gin-vue-template/pkg/logic"
	"gin-vue-template/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *API) ListAdmins(c *gin.Context) {
	ctx := NewGinContextHelper(c).Context()

	info, ok := c.Get(constants.AuthKey)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}
	loginInfo := info.(*logic.LoginInfo)

	if constants.Role(loginInfo.Role) != constants.RoleAdmin {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	pagination := &models.Pagination{}
	if err := c.ShouldBindQuery(pagination); err != nil {
		e.GinR(c, e.RetCodeInvalidParams, nil)
		return
	}

	admins := a.dao.ListAdmins(ctx, &models.ListRequest{
		QueryCount: true,
		Pagination: pagination,
	})

	e.GinOK(c, admins)
}
