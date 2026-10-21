package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/radish-go/gin-vue-admin/server/global"
	"github.com/radish-go/gin-vue-admin/server/middleware"
	systemRouter "github.com/radish-go/gin-vue-admin/server/router/system"
)

// Routers 装配 gin 引擎：公共路由 + JWT 私有路由
func Routers() *gin.Engine {
	if global.GVA_CONFIG.System.Env == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.Cors())

	prefix := global.GVA_CONFIG.System.Prefix
	public := r.Group(prefix)
	private := r.Group(prefix)
	private.Use(middleware.JWTAuth())

	// 各业务模块的路由在此注册
	systemRouter.RegisterRouters(public, private)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ok", "data": gin.H{
			"service": "gin-vue-admin",
			"env":     global.GVA_CONFIG.System.Env,
		}})
	})
	return r
}
