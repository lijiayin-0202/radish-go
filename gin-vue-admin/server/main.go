package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/radish-go/gin-vue-admin/server/core"
	"github.com/radish-go/gin-vue-admin/server/global"
	"github.com/radish-go/gin-vue-admin/server/initialize"
)

// radish-go / gin-vue-admin 子项目：后台管理脚手架服务端入口
func main() {
	configPath := flag.String("c", "config.yaml", "配置文件路径")
	flag.Parse()

	// 1. 加载配置
	global.GVA_CONFIG = core.LoadConfig(*configPath)
	// 2. 初始化日志
	global.GVA_LOG = core.NewLogger(global.GVA_CONFIG.Zap)
	// 3. 初始化 MySQL（连接失败不阻断启动，降级为内置演示账号）
	global.GVA_DB = initialize.GormMysql()
	// 4. 初始化 Redis（未配置则跳过）
	global.GVA_REDIS = initialize.Redis()
	// 5. 注册路由
	engine := initialize.Routers()

	runServer(global.GVA_CONFIG.System.Addr, engine)
}

// runServer 启动 HTTP 服务并支持优雅退出
func runServer(addr string, engine *gin.Engine) {
	srv := &http.Server{
		Addr:    addr,
		Handler: engine,
	}

	go func() {
		global.GVA_LOG.Info("gin-vue-admin server started", zap.String("addr", addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.GVA_LOG.Info("shutting down gin-vue-admin server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.GVA_LOG.Error("server forced to shutdown", zap.Error(err))
	}
	global.GVA_LOG.Info("gin-vue-admin server exited")
}
