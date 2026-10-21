package core

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config 后台管理脚手架的完整配置，对应 config.yaml
type Config struct {
	System System `yaml:"system"`
	JWT    JWT    `yaml:"jwt"`
	Mysql  Mysql  `yaml:"mysql"`
	Redis  Redis  `yaml:"redis"`
	Zap    Zap    `yaml:"zap"`
}

type System struct {
	// Env: debug / release / test
	Env    string `yaml:"env"`
	Addr   string `yaml:"addr"`
	Prefix string `yaml:"prefix"`
}

type JWT struct {
	SigningKey  string `yaml:"signing-key"`
	ExpiresTime string `yaml:"expires-time"`
	Issuer      string `yaml:"issuer"`
}

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	// Enable 为 false 时不连接数据库，使用内置演示账号（便于无环境的同学一键启动）
	Enable bool `yaml:"enable"`
}

type Redis struct {
	Enable   bool   `yaml:"enable"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type Zap struct {
	Level    string `yaml:"level"`
	Format   string `yaml:"format"`
	Prefix   string `yaml:"prefix"`
	Director string `yaml:"director"`
	// ShowLine 是否显示行号
	ShowLine bool `yaml:"show-line"`
	// Stacktrace 日志级别以上输出堆栈
	Stacktrace string `yaml:"stacktrace"`
}

// LoadConfig 从 path 读取 yaml 配置；文件不存在时回落到默认配置（保证 go run 可直接启动）
func LoadConfig(path string) *Config {
	cfg := defaultConfig()

	if p := os.Getenv("RADISH_GVA_CONFIG"); p != "" {
		path = p
	}

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("[config] 读取 %s 失败(%v)，使用默认配置\n", path, err)
		return cfg
	}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		fmt.Printf("[config] 解析 %s 失败(%v)，使用默认配置\n", path, err)
		return defaultConfig()
	}
	fillDefault(cfg)
	return cfg
}

func defaultConfig() *Config {
	return &Config{
		System: System{Env: "debug", Addr: ":8888", Prefix: "/api"},
		JWT:    JWT{SigningKey: "radish-go", ExpiresTime: "24h", Issuer: "radish-go"},
		Mysql:  Mysql{Enable: false, Host: "127.0.0.1", Port: 3306, Username: "root", Password: "root", Dbname: "radish_gva"},
		Redis:  Redis{Enable: false, Host: "127.0.0.1", Port: 6379, DB: 0},
		Zap:    Zap{Level: "info", Format: "console", Prefix: "[radish-gva]", Director: "logs", ShowLine: true, Stacktrace: "error"},
	}
}

func fillDefault(cfg *Config) {
	d := defaultConfig()
	if cfg.System.Addr == "" {
		cfg.System.Addr = d.System.Addr
	}
	if cfg.System.Prefix == "" {
		cfg.System.Prefix = d.System.Prefix
	}
	if cfg.JWT.SigningKey == "" {
		cfg.JWT.SigningKey = d.JWT.SigningKey
	}
	if cfg.JWT.ExpiresTime == "" {
		cfg.JWT.ExpiresTime = d.JWT.ExpiresTime
	}
	if cfg.Zap.Level == "" {
		cfg.Zap.Level = d.Zap.Level
	}
}
