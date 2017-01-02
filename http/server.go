package http

import (
	"net"
	"net/http"
	"net/url"
)

// DefaultAddr is the default bind address.
const DefaultAddr = ":3000"

// Server represents an HTTP server.
type Server struct {
	ln net.Listener

	// Handler to serve.
	Handler *Handler

	// Bind address to open.
	Addr string
}
