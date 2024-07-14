package logic

import (
	"gin-vue-template/pkg/infra/redis"
	"gin-vue-template/pkg/models"
)

const (
	redisPrefixLogin = "login"
)

type Logic struct {
	loginRedis *redis.PrefixRedisClient
	dao        *models.GormDAO
}

func NewLogic(redisCli *redis.PrefixRedisClient, dao *models.GormDAO) *Logic {
	return &Logic{
		loginRedis: redisCli.WithPrefix(redisPrefixLogin),
		dao:        dao,
	}
}
