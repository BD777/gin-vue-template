package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-vue-template/middleware"
	"gin-vue-template/pkg/setting"
	"gin-vue-template/routers/api"
)

func initFrontendFiles(r *gin.Engine) {
	r.StaticFile("/", "./frontend/dist/index.html")
	r.StaticFile("/favicon.ico", "./frontend/dist/favicon.ico")
	r.StaticFS("/js", http.Dir("./frontend/dist/js"))
	r.StaticFS("/css", http.Dir("./frontend/dist/css"))
}

func InitRouter() *gin.Engine {
	r := gin.New()
	initFrontendFiles(r)

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	apiGroup := r.Group("/api")

	noauthAPI := apiGroup.Group("/public")
	{
		noauthAPI.GET("/logininfo", api.GetLoginInfo)
		noauthAPI.POST("/login", api.Login)
	}

	commonAPI := apiGroup.Group("/common")
	commonAPI.Use(middleware.AuthMiddleware())
	{
		commonAPI.GET("/pageauth", api.GetPageAuth)
	}

	settingAdminAPI := apiGroup.Group("/setting/admin")
	settingAdminAPI.Use(middleware.AuthMiddleware())
	{
		settingAdminAPI.GET("/list", api.ListAdmins)
	}

	return r
}
