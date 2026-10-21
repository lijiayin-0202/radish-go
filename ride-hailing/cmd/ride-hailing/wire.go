//go:build wireinject
// +build wireinject

// 该文件只在执行 wire 生成时使用，正常编译会被 build tag 排除。
// 重新生成命令：cd cmd/ride-hailing && wire
package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"github.com/radish-go/ride-hailing/internal/biz"
	"github.com/radish-go/ride-hailing/internal/conf"
	"github.com/radish-go/ride-hailing/internal/data"
	"github.com/radish-go/ride-hailing/internal/server"
	"github.com/radish-go/ride-hailing/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
