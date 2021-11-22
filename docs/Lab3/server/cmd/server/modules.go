//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	mn "github.com/mezidia/architecture_labs/tree/main/docs/Lab3/server/menu"
)

// ComposeApiServer will create an instance of CharApiServer according to providers defined in this file.
func ComposeApiServer(port HttpPortNumber) (*MenuApiServer, error) {
	wire.Build(
		// DB connection provider (defined in main.go).
		NewDbConnection,
		// Add providers from channels package.
		mn.Providers,
		// Provide ChatApiServer instantiating the structure and injecting channels handler and port number.
		wire.Struct(new(MenuApiServer), "Port", "ChannelsHandler"),
	)
	return nil, nil
}
