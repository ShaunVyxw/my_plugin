package my_plugin


import (
	"context"
	"net/http"
)

// Config the plugin configuration.
type Config struct {
	Enable bool
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Enable: true,
	}
}

// Demo a Demo plugin.
type Demo struct {
	next     http.Handler
	name     string
}

// New created a new Demo plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if config.Enable {
		return &Demo{
		}, nil
	} else {
		return nil, nil
	}
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
