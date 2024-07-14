package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-vue-template/config"
	"gin-vue-template/routers/api"
	"gin-vue-template/routers/middleware"
)

func initFrontendFiles(r *gin.Engine) {
	r.StaticFile("/", "./frontend/dist/index.html")
	r.StaticFile("/favicon.ico", "./frontend/dist/favicon.ico")
	r.StaticFS("/js", http.Dir("./frontend/dist/js"))
	r.StaticFS("/css", http.Dir("./frontend/dist/css"))
}

func NewRouter(cfg *config.Config, mw *middleware.Middleware, a *api.API) *gin.Engine {
	r := gin.New()
	initFrontendFiles(r)

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(cfg.App.RunMode)

	r.Use(mw.InitMiddleware())

	apiGroup := r.Group("/api")

	noauthAPI := apiGroup.Group("/public")
	{
		noauthAPI.GET("/logininfo", a.GetLoginInfo)
		noauthAPI.POST("/login", a.Login)
	}

	commonAPI := apiGroup.Group("/common")
	commonAPI.Use(mw.AuthMiddleware())
	{
		commonAPI.GET("/pageauth", a.GetPageAuth)
	}

	settingAdminAPI := apiGroup.Group("/setting/admin")
	settingAdminAPI.Use(mw.AuthMiddleware())
	{
		settingAdminAPI.GET("/list", a.ListAdmins)
	}

	return r
}
