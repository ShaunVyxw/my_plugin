package my_plugin


import (
	"context"
	"net/http"
)

// Config the plugin configuration.
type Config struct {
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
	}
}

// Demo a Demo plugin.
type Demo struct {
	next     http.Handler
	name     string
}

// New created a new Demo plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &Demo{
	}, nil
}

func (a *Demo) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	xOriginUri := req.RequestURI
	xForwardedHost := req.Host
	xOriginalMethod := req.Method
	xForwardedProto := "http"
	if req.TLS != nil {
		xForwardedProto = "https"
	}

	req.Header.Set("X-Origin-URI", xOriginUri)
	req.Header.Set("X-Forwarded-Host", xForwardedHost)
	req.Header.Set("X-Original-METHOD", xOriginalMethod)
	req.Header.Set("X-Forwarded-Proto", xForwardedProto)


	a.next.ServeHTTP(rw, req)
}
