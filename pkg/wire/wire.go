//go:build wireinject

package wire

//go:generate wire
import (
	"gin-vue-template/config"
	"gin-vue-template/pkg/infra/redis"
	"gin-vue-template/pkg/logic"
	"gin-vue-template/pkg/models"
	"gin-vue-template/pkg/service"
	"gin-vue-template/routers"
	"gin-vue-template/routers/api"
	"gin-vue-template/routers/middleware"

	"github.com/google/wire"
)

func InitializeService() *service.Service {
	wire.Build(
		models.NewDAO,
		redis.NewRedisClient,
		logic.NewLogic,
		middleware.NewMiddleware,
		api.NewAPI,
		routers.NewRouter,
		config.NewConfig,
		service.NewHTTPServer,

		wire.Struct(new(service.Service), "*"),
	)
	return nil
}
