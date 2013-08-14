package socketio

type IOEndpoint struct {
	Path  string
	Query string
}

func NewEndpoint(path, query string) *IOEndpoint {
	return &IOEndpoint{Path: path, Query: query}
}

func (e IOEndpoint) String() string {
	if e.Query != "" {
		return e.Path + "?" + e.Query
	}
	return e.Path
}
