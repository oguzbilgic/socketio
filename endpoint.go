package socketio

import (
	"strings"
)

// Endpoint is a socket.io endpoint.
type Endpoint struct {
	Path  string
	Query string
}

// NewEndpoint returns a new Endpoint with the given path and query.
func NewEndpoint(path, query string) *Endpoint {
	return &Endpoint{path, query}
}

// ParseEndpoint parses the given rawn endpoint into an Endpoint.
func ParseEndpoint(rawEndpoint string) *Endpoint {
	parts := strings.SplitN(rawEndpoint, "?", 2)

	if len(parts) == 1 {
		return &Endpoint{parts[0], ""}
	}

	return &Endpoint{parts[0], parts[1]}
}

// String returns the string representation of the endpoint.
func (e Endpoint) String() string {
	if e.Query != "" {
		return e.Path + "?" + e.Query
	}
	return e.Path
}
