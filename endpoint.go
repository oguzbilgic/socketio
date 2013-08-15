package socketio

type Endpoint struct {
	Path  string
	Query string
}

func NewEndpoint(path, query string) *Endpoint {
	return &Endpoint{path, query}
}

func (e Endpoint) String() string {
	if e.Query != "" {
		return e.Path + "?" + e.Query
	}
	return e.Path
}
