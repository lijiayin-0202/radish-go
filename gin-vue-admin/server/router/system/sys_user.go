package system

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/radish-go/gin-vue-admin/server/api/v1/system"
)

// UserRouter 用户路由
type UserRouter struct{}

// InitPublicRouter 无需登录即可访问的路由
func (UserRouter) InitPublicRouter(r *gin.RouterGroup) {
	api := v1.UserApi{}
	r.POST("/user/login", api.Login)
	r.POST("/user/register", api.Register)
}

// InitPrivateRouter 需要 JWT 鉴权的路由
func (UserRouter) InitPrivateRouter(r *gin.RouterGroup) {
	api := v1.UserApi{}
	r.GET("/user/info", api.GetUserInfo)
	r.GET("/user/list", api.ListUsers)
	r.POST("/user/logout", api.Logout)
	r.GET("/user/refresh", api.RefreshToken)
}

// RegisterRouters 供 initialize/router.go 统一装配
func RegisterRouters(public, private *gin.RouterGroup) {
	router := UserRouter{}
	router.InitPublicRouter(public)
	router.InitPrivateRouter(private)
}
