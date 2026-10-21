package system

import "time"

// SysUser 后台管理用户表
type SysUser struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Username  string    `gorm:"uniqueIndex;size:64;not null" json:"username"`
	Password  string    `gorm:"size:128;not null" json:"-"`
	Nickname  string    `gorm:"size:64" json:"nickname"`
	Phone     string    `gorm:"size:20" json:"phone"`
	Role      string    `gorm:"size:32;default:user" json:"role"`
}

func (SysUser) TableName() string { return "sys_users" }

// ------------------------------ DTO ------------------------------

// LoginReq 登录入参
type LoginReq struct {
	Username string `json:"username" binding:"required" example:"admin"`
	Password string `json:"password" binding:"required" example:"123456"`
}

// RegisterReq 注册入参
type RegisterReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
}

// UserInfo 脱敏后的用户信息
type UserInfo struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}

// LoginResp 登录出参
type LoginResp struct {
	Token string   `json:"token"`
	User  UserInfo `json:"user"`
}
