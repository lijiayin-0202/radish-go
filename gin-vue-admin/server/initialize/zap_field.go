package initialize

import "go.uber.org/zap"

// zapError 统一包装错误字段
func zapError(err error) zap.Field {
	return zap.Error(err)
}
