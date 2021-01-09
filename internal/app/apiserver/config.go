package apiserver

import "github.com/Dukastlik/avitomx-api.git/internal/app/products"

// Config ...
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml: "log_level"`
	Products *products.Config
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Products: products.NewConfig(),
	}
}
