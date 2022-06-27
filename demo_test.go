package my_plugin_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ShaunVyxw/my_plugin"
)

func TestDemo(t *testing.T) {
	cfg := my_plugin.CreateConfig()

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := my_plugin.New(ctx, next, cfg, "my_plugin")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	assertHeader(t, req, "X-Origin-URI", "/test")
	assertHeader(t, req, "X-Forwarded-Host", "localhost")
	assertHeader(t, req, "X-Original-METHOD", "GET")
	assertHeader(t, req, "XX-Forwarded-Proto", "http")
}

func assertHeader(t *testing.T, req *http.Request, key, expected string) {
	t.Helper()

	if req.Header.Get(key) != expected {
		t.Errorf("invalid header value: %s", req.Header.Get(key))
	}
}