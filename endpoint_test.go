package socketio

import (
	"testing"
)

func TestParseEndpointEmpty(t *testing.T) {
	endpoint := ParseEndpoint("")

	if str := endpoint.String(); str != "" {
		t.Errorf("Endpoint.String() = %s ; want empty", str)
	}

	if path := endpoint.Path; path != "" {
		t.Errorf("Endpoint.Path = %s ; want empty", path)
	}

	if query := endpoint.Query; query != "" {
		t.Errorf("Endpoint.Query = %s ; want empty", query)
	}
}

func TestParseEndpoint(t *testing.T) {
	endpoint := ParseEndpoint("/channel")

	if str := endpoint.String(); str != "/channel" {
		t.Errorf("Endpoint.String() = %s ; want /channel", str)
	}

	if path := endpoint.Path; path != "/channel" {
		t.Errorf("Endpoint.Path = %s ; want /channel", path)
	}

	if query := endpoint.Query; query != "" {
		t.Errorf("Endpoint.Query = %s ; want empty", query)
	}
}

func TestParseEndpointQuery(t *testing.T) {
	endpoint := ParseEndpoint("/channel?key=value")

	if str := endpoint.String(); str != "/channel?key=value" {
		t.Errorf("Endpoint.String() = %s ; want /channel", str)
	}

	if path := endpoint.Path; path != "/channel" {
		t.Errorf("Endpoint.Path = %s ; want /channel", path)
	}

	if query := endpoint.Query; query != "key=value" {
		t.Errorf("Endpoint.Query = %s ; want empty", query)
	}
}
