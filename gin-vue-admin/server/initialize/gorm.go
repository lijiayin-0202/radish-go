package initialize

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/radish-go/gin-vue-admin/server/global"
	"github.com/radish-go/gin-vue-admin/server/model/system"
)

// GormMysql 初始化 MySQL 连接。
// 未启用或连接失败时返回 nil，服务降级为内置演示账号，保证无数据库环境也能启动。
func GormMysql() *gorm.DB {
	cfg := global.GVA_CONFIG.Mysql
	if !cfg.Enable || cfg.Host == "" {
		global.GVA_LOG.Warn("MySQL 未启用，已切换到内置演示账号模式（admin / 123456）")
		return nil
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Dbname)

	level := logger.Info
	if global.GVA_CONFIG.System.Env == "release" {
		level = logger.Warn
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(level)})
	if err != nil {
		global.GVA_LOG.Error("MySQL 连接失败，降级为演示账号模式", zapError(err))
		return nil
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := db.AutoMigrate(&system.SysUser{}); err != nil {
		global.GVA_LOG.Error("自动迁移失败", zapError(err))
	}
	seedAdmin(db)
	return db
}

// seedAdmin 首次启动时写入默认管理员
func seedAdmin(db *gorm.DB) {
	var count int64
	db.Model(&system.SysUser{}).Count(&count)
	if count > 0 {
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	db.Create(&system.SysUser{
		Username: "admin",
		Password: string(hash),
		Nickname: "超级管理员",
		Role:     "admin",
	})
	global.GVA_LOG.Info("已初始化默认管理员 admin / 123456")
}
