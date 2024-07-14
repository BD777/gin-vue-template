package ctxmeta

import (
	"context"
	"gin-vue-template/pkg/constants"
)

type ctxKey string

const (
	LogIDKey    ctxKey = "logID"
	TenantIDKey ctxKey = "tenantID"
	UsernameKey ctxKey = "username"
	UserRoleKey ctxKey = "userRole"
)

func WithLogID(ctx context.Context, logID string) context.Context {
	return context.WithValue(ctx, LogIDKey, logID)
}

func GetLogID(ctx context.Context) string {
	if logID, ok := ctx.Value(LogIDKey).(string); ok {
		return logID
	}

	return ""
}

func WithUsername(ctx context.Context, username string) context.Context {
	return context.WithValue(ctx, UsernameKey, username)
}

func GetUsername(ctx context.Context) string {
	if username, ok := ctx.Value(UsernameKey).(string); ok {
		return username
	}

	return ""
}

func WithUserRole(ctx context.Context, userRole string) context.Context {
	return context.WithValue(ctx, UserRoleKey, userRole)
}

func GetUserRole(ctx context.Context) string {
	if userRole, ok := ctx.Value(UserRoleKey).(string); ok {
		return userRole
	}

	return ""
}

func WithTenantID(ctx context.Context, tenantID int64) context.Context {
	return context.WithValue(ctx, TenantIDKey, tenantID)
}

func GetTenantID(ctx context.Context) int64 {
	if tenantID, ok := ctx.Value(TenantIDKey).(int64); ok {
		return tenantID
	}

	return constants.InvalidTenantID
}
