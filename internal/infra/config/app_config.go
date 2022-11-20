//go:build wireinject
// +build wireinject

package config

import (
	"remoteworkout/internal/infra/web"

	"github.com/google/wire"
)

type AppContext struct {
	Server *web.HttpServer
}

func ProvideHttpServer() *web.HttpServer {
	wire.Build(web.CreateHttpServer)
	return &web.HttpServer{}
}

func ProvideAppContext() *AppContext {
	wire.Build(ContextSet)
	return &AppContext{}
}

var ContextSet = wire.NewSet(
	ProvideHttpServer,
	wire.Struct(new(AppContext), "*"),
)
