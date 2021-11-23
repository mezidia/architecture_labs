package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mezidia/architecture_labs/tree/main/docs/Lab3/server/menu"
)

type HttpPortNumber int

// ChatApiServer configures necessary handlers and starts listening on a configured port.
type MenuApiServer struct {
	Port HttpPortNumber

	MenuHandler menu.HttpHandlerFunc

	server *http.Server
}

// Start will set all handlers and start listening.
// If this methods succeeds, it does not return until server is shut down.
// Returned error will never be nil.
func (s *MenuApiServer) Start() error {
	if s.MenuHandler == nil {
		return fmt.Errorf("channels HTTP handler is not defined - cannot start")
	}
	if s.Port == 0 {
		return fmt.Errorf("port is not defined")
	}

	handler := new(http.ServeMux)
	handler.HandleFunc("/menu", s.MenuHandler)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

// Stops will shut down previously started HTTP server.
func (s *MenuApiServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}
	return s.server.Shutdown(context.Background())
}
