package socketio

import (
	"strings"
)

type Endpoint struct {
	Path  string
	Query string
}

func NewEndpoint(path, query string) *Endpoint {
	return &Endpoint{path, query}
}

func ParseEndpoint(rawEndpoint string) *Endpoint {
	parts := strings.SplitN(rawEndpoint, "?", 2)

	if len(parts) == 1 {
		return &Endpoint{parts[0], ""}
	}

	return &Endpoint{parts[0], parts[1]}
}

func (e Endpoint) String() string {
	return e.Path + "?" + e.Query
}
