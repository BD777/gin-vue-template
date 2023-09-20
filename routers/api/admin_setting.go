package api

import (
	"gin-vue-template/constants"
	"gin-vue-template/models"
	"gin-vue-template/pkg/e"
	"gin-vue-template/pkg/logic"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListAdmins(c *gin.Context) {
	ctx := c.Request.Context()

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

	helper := NewGinContextHelper(c)
	pagination := helper.GetPagination()
	admins := models.ListAdmins(ctx, &models.ListRequest{
		QueryCount: true,
		Pagination: pagination,
	})
	c.JSON(http.StatusOK, e.R(e.OK, admins))
}
