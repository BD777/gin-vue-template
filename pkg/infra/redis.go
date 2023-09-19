package infra

import (
	"context"
	"gin-vue-template/pkg/setting"
	"log"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	sec, err := setting.Cfg.GetSection("redis")
	if err != nil {
		log.Fatal(2, "Fail to get section 'redis': %v", err)
		return nil
	}

	addr := sec.Key("ADDR").String()
	password := sec.Key("PASSWORD").String()
	db, err := sec.Key("DB").Int()
	if err != nil {
		log.Fatal(2, "Fail to get section 'redis': %v", err)
		return nil
	}

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       int(db),
	})
	return client
}

type PrefixRedisClient struct {
	cli    *redis.Client
	prefix string
}

var RedisClient = NewRedisClient()

func NewRedisWithPrefix(prefix string) *PrefixRedisClient {
	return &PrefixRedisClient{
		cli:    RedisClient,
		prefix: prefix,
	}
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
