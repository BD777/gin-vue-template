package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

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

	adminAPI := apiGroup.Group("/admin")
	{
		adminAPI.GET("/pageauth", api.GetPageAuth)
		adminAPI.GET("/logininfo", api.GetLoginInfo)
		adminAPI.POST("/login", api.Login)
	}
	// TODO: middleware to check login status

	return r
}
