package traefik_helloworld_plugin

import (
	"context"
	"fmt"
	"net/http"
)

// Config the plugin configuration.
type Config struct {
	Username string `json:"username"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

// Example a plugin.
type Helloworld struct {
	next     http.Handler
	username string
}

// New created a new plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &Helloworld{
		next:     next,
		username: config.Username,
	}, nil
}

func (e *Helloworld) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Errorf("Hello world")
	e.next.ServeHTTP(rw, req)
}
