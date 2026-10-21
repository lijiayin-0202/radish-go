package initialize

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/radish-go/gin-vue-admin/server/global"
)

// Redis 初始化 Redis 客户端，未启用时返回 nil
func Redis() *redis.Client {
	cfg := global.GVA_CONFIG.Redis
	if !cfg.Enable {
		global.GVA_LOG.Warn("Redis 未启用，跳过连接")
		return nil
	}

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := client.Ping(ctx).Err(); err != nil {
		global.GVA_LOG.Error("Redis 连接失败", zapError(err))
		return nil
	}
	global.GVA_LOG.Info("Redis 连接成功")
	return client
}
