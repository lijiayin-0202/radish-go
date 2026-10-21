package global

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/radish-go/gin-vue-admin/server/core"
)

// 全局单例变量，统一在 main 中初始化，其余包只读引用
var (
	GVA_CONFIG *core.Config
	GVA_LOG    *zap.Logger
	GVA_DB     *gorm.DB      // 可能为 nil（未启用或未连上 MySQL 时降级）
	GVA_REDIS  *redis.Client // 可能为 nil（未启用 Redis）
)
