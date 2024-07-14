package api

import (
	"context"
	"gin-vue-template/pkg/constants"
	"gin-vue-template/pkg/logic"
	"gin-vue-template/pkg/models"
	"gin-vue-template/pkg/utils/ctxmeta"

	"github.com/gin-gonic/gin"
)

type API struct {
	logic *logic.Logic
	dao   *models.GormDAO
}

func NewAPI(logic *logic.Logic, dao *models.GormDAO) *API {
	return &API{
		logic: logic,
		dao:   dao,
	}
}

type GinContextHelper struct {
	c *gin.Context
}

func NewGinContextHelper(c *gin.Context) *GinContextHelper {
	return &GinContextHelper{c: c}
}

func (g *GinContextHelper) Context() context.Context {
	ctx := g.c.Request.Context()
	if logid, ok := g.c.Get(string(ctxmeta.LogIDKey)); ok {
		ctx = ctxmeta.WithLogID(ctx, logid.(string))
	}

	if loginInfo, ok := g.c.Get(constants.AuthKey); ok {
		info := loginInfo.(*logic.LoginInfo)
		ctx = ctxmeta.WithUsername(ctx, info.Username)
		ctx = ctxmeta.WithUserRole(ctx, info.Role)
	}

	return ctx
}
