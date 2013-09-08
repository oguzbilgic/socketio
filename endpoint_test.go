package socketio

import (
	"testing"
)

func TestParseEndpoint(t *testing.T) {
	endpoint := ParseEndpoint("/channel")

	if endpoint.Path != "/channel" {
		t.Errorf("Error")
	}

	if endpoint.Query != "" {
		t.Errorf("Error")
	}
}

func TestParseEndpointQuery(t *testing.T) {
	endpoint := ParseEndpoint("/channel?key=value")

	if endpoint.Path != "/channel" {
		t.Errorf("Error")
	}

	if endpoint.Query != "key=value" {
		t.Errorf("Error")
	}
}
