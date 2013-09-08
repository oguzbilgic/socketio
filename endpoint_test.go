package socketio

import (
	"testing"
)

func TestParseEndpoint(t *testing.T) {
	endpoint := ParseEndpoint("")

	if endpoint.Path != "" {
		t.Errorf("Error")
	}

	if endpoint.Query != "" {
		t.Errorf("Error")
	}
}

func TestParseEndpoint2(t *testing.T) {
	endpoint := ParseEndpoint("/channel")

	if endpoint.Path != "/channel" {
		t.Errorf("Error")
	}

	if endpoint.Query != "" {
		t.Errorf("Error")
	}
}

func TestParseEndpoint3(t *testing.T) {
	endpoint := ParseEndpoint("/channel?Key=Val")

	if endpoint.Path != "/channel" {
		t.Errorf("Error")
	}

	if endpoint.Query != "Key=Val" {
		t.Errorf("Error")
	}
}
