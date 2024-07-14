package redis

import (
	"context"
	"gin-vue-template/config"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type PrefixRedisClient struct {
	cli    *redis.Client
	prefix string
}

func NewRedisClient(cfg *config.Config) *PrefixRedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	return &PrefixRedisClient{cli: client}
}

func (p *PrefixRedisClient) WithPrefix(prefix string) *PrefixRedisClient {
	return &PrefixRedisClient{cli: p.cli, prefix: prefix}
}

func (p *PrefixRedisClient) genKey(key string) string {
	return strings.Join([]string{p.prefix, key}, "_")
}

func (p *PrefixRedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return p.cli.Set(ctx, p.genKey(key), value, expiration)
}

func (p *PrefixRedisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	return p.cli.Get(ctx, p.genKey(key))
}
