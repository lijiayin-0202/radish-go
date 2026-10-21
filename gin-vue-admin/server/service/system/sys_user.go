package system

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/radish-go/gin-vue-admin/server/global"
	"github.com/radish-go/gin-vue-admin/server/model/system"
	"github.com/radish-go/gin-vue-admin/server/utils"
)

// demoUser 无数据库环境下的内置演示账号
var demoUser = system.SysUser{
	ID:       1,
	Username: "admin",
	Nickname: "超级管理员",
	Role:     "admin",
}

var demoPasswordHash = mustHash("123456")

func mustHash(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return pwd
	}
	return string(hash)
}

type UserService struct{}

// Login 校验账号密码并签发 token
func (s *UserService) Login(req system.LoginReq) (*system.LoginResp, error) {
	user, err := s.findUser(req.Username)
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		return nil, errors.New("用户名或密码错误")
	}
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		return nil, errors.New("token 签发失败")
	}
	return &system.LoginResp{
		Token: token,
		User:  toUserInfo(*user),
	}, nil
}

// Register 注册新用户
func (s *UserService) Register(req system.RegisterReq) error {
	if global.GVA_DB == nil {
		return errors.New("演示模式下不支持注册，请先配置 MySQL")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	nickname := req.Nickname
	if nickname == "" {
		nickname = req.Username
	}
	return global.GVA_DB.Create(&system.SysUser{
		Username: req.Username,
		Password: string(hash),
		Nickname: nickname,
		Phone:    req.Phone,
		Role:     "user",
	}).Error
}

// GetUserInfo 按 ID 查询用户信息
func (s *UserService) GetUserInfo(id uint) (*system.UserInfo, error) {
	if global.GVA_DB == nil {
		if id == demoUser.ID {
			info := toUserInfo(demoUser)
			return &info, nil
		}
		return nil, errors.New("用户不存在")
	}
	user := system.SysUser{}
	if err := global.GVA_DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}
	info := toUserInfo(user)
	return &info, nil
}

// ListUsers 用户列表（分页）
func (s *UserService) ListUsers(page, pageSize int) ([]system.UserInfo, int64, error) {
	if global.GVA_DB == nil {
		return []system.UserInfo{toUserInfo(demoUser)}, 1, nil
	}
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}
	var users []system.SysUser
	var total int64
	global.GVA_DB.Model(&system.SysUser{}).Count(&total)
	if err := global.GVA_DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	list := make([]system.UserInfo, 0, len(users))
	for i := range users {
		list = append(list, toUserInfo(users[i]))
	}
	return list, total, nil
}

func (s *UserService) findUser(username string) (*system.SysUser, error) {
	if global.GVA_DB == nil {
		if username == demoUser.Username {
			u := demoUser
			u.Password = demoPasswordHash
			return &u, nil
		}
		return nil, errors.New("用户不存在")
	}
	user := system.SysUser{}
	if err := global.GVA_DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func toUserInfo(u system.SysUser) system.UserInfo {
	return system.UserInfo{
		ID:       u.ID,
		Username: u.Username,
		Nickname: u.Nickname,
		Phone:    u.Phone,
		Role:     u.Role,
	}
}
