//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/mezidia/architecture_labs/tree/main/docs/Lab3/server/menu"
)

func ComposeApiServer(port HttpPortNumber) (*MenuApiServer, error) {
	wire.Build(
		NewDbConnection,
		menu.Providers,
		wire.Struct(new(MenuApiServer), "Port", "MenuHandler"),
	)
	return nil, nil
}
