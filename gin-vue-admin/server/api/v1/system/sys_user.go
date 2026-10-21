package system

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/radish-go/gin-vue-admin/server/global"
	"github.com/radish-go/gin-vue-admin/server/model/common/response"
	"github.com/radish-go/gin-vue-admin/server/model/system"
	svc "github.com/radish-go/gin-vue-admin/server/service/system"
	"github.com/radish-go/gin-vue-admin/server/utils"
)

// UserApi 用户相关接口（gva 约定的 Api 层只做参数绑定与响应封装）
type UserApi struct{}

var userService = new(svc.UserService)

// Login 登录
// @Summary 用户登录
// @Tags User
// @Accept json
// @Produce json
// @Success 0 {object} response.Response{data=system.LoginResp}
// @Router /user/login [post]
func (UserApi) Login(c *gin.Context) {
	var req system.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误："+err.Error(), c)
		return
	}
	resp, err := userService.Login(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(resp, c)
}

// Register 注册
func (UserApi) Register(c *gin.Context) {
	var req system.RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误："+err.Error(), c)
		return
	}
	if err := userService.Register(req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("注册成功", c)
}

// GetUserInfo 获取当前登录用户信息
func (UserApi) GetUserInfo(c *gin.Context) {
	id, ok := c.Get("userID")
	if !ok {
		response.FailWithCode(response.AUTHERROR, "未登录", c)
		return
	}
	info, err := userService.GetUserInfo(id.(uint))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(info, c)
}

// ListUsers 用户列表
func (UserApi) ListUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	list, total, err := userService.ListUsers(page, pageSize)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(gin.H{"list": list, "total": total, "page": page, "pageSize": pageSize}, c)
}

// Logout 退出登录（JWT 无状态，前端删除 token 即可）
func (UserApi) Logout(c *gin.Context) {
	global.GVA_LOG.Info("用户退出登录", zap.String("username", c.GetString("username")))
	response.OkWithMessage("退出成功", c)
}

// RefreshToken 续签 token
func (UserApi) RefreshToken(c *gin.Context) {
	id, _ := c.Get("userID")
	userID, _ := id.(uint)
	token, err := utils.GenerateToken(userID, c.GetString("username"))
	if err != nil {
		response.FailWithMessage("token 续签失败", c)
		return
	}
	response.OkWithData(gin.H{"token": token}, c)
}
