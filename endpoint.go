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

	if len(parts) == 0 {
		return &Endpoint{"", ""}
	} else if len(parts) == 1 {
		return &Endpoint{parts[0], ""}
	} else {
		return &Endpoint{parts[0], parts[1]}
	}
}

func (e Endpoint) String() string {
	if e.Query != "" {
		return e.Path + "?" + e.Query
	}
	return e.Path
}
