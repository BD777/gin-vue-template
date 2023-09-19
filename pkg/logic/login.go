package logic

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"gin-vue-template/models"
	"gin-vue-template/pkg/infra"
	"log"
	"time"
)

var loginRedis = infra.NewRedisWithPrefix("login")

type LoginInfo struct {
	Username  string `json:"username"`
	Role      string `json:"role"`
	Token     string `json:"token"`
	LoginTime int64  `json:"login_time"`
}

func (l *LoginInfo) MarshalBinary() ([]byte, error) {
	return json.Marshal(l)
}

func GetLoginInfo(ctx context.Context, token string) (info *LoginInfo, err error) {
	value, err := loginRedis.Get(ctx, token).Result()
	if err != nil {
		return nil, fmt.Errorf("get login info from redis failed: %w", err)
	}

	info = &LoginInfo{}
	err = json.Unmarshal([]byte(value), info)
	if err != nil {
		return nil, fmt.Errorf("unmarshal login info failed: %w", err)
	}

	return info, nil
}

func genToken(info *LoginInfo) string {
	jsonBytes, err := json.Marshal(info)
	if err != nil {
		log.Fatalf("marshal login info failed: %v", err)
		return ""
	}

	h := sha1.New()
	h.Write(jsonBytes)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func passwordHash(password string) string {
	const salt = "1afwom@!2fawf"
	h := sha1.New()
	h.Write([]byte(password + salt))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Login(ctx context.Context, username, password string) (info *LoginInfo, err error) {
	admin := models.GetAdmin(username, passwordHash(password))
	if admin == nil {
		return nil, fmt.Errorf("invalid username or password")
	}

	info = &LoginInfo{
		Username:  admin.Username,
		Role:      admin.Role,
		LoginTime: time.Now().Unix(),
	}
	info.Token = genToken(info)

	err = loginRedis.Set(ctx, info.Token, info, time.Hour*24).Err()
	if err != nil {
		return nil, fmt.Errorf("set login info to redis failed: %w", err)
	}

	return info, nil
}
