package server

import (
	"context"
	"strings"
	"testing"

	pluginv1 "github.com/Silo-Server/silo-plugin-sdk/pkg/pluginproto/silo/plugin/v1"
)

func TestSetupPageNeedsNoBackend(t *testing.T) {
	response, err := New(nil).Handle(context.Background(), &pluginv1.HandleHTTPRequest{
		Method: "GET",
		Path:   "/v1/setup",
	})
	if err != nil || response.GetStatusCode() != 200 {
		t.Fatalf("setup page: response=%v error=%v", response, err)
	}
	if !strings.HasPrefix(response.GetHeaders()["Content-Type"], "text/html") {
		t.Fatalf("setup page content type = %q", response.GetHeaders()["Content-Type"])
	}
	if !strings.Contains(string(response.GetBody()), "Comic Pages") {
		t.Fatal("setup page body is missing its title")
	}
}

func TestSetupPageIsGetOnly(t *testing.T) {
	response, _ := New(nil).Handle(context.Background(), &pluginv1.HandleHTTPRequest{
		Method: "POST",
		Path:   "/v1/setup",
		Body:   []byte("{}"),
	})
	if response.GetStatusCode() == 200 {
		t.Fatal("POST /v1/setup must not serve the page")
	}
}
