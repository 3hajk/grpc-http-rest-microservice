package cfg

import "time"

type Config struct {
	GRPCService GRPC `mapstructure:"GRPC"`
	HTTPService HTTP `mapstructure:"HTTP"`

	Regenerate time.Duration `mapstructure:"REGENERATE"      default:"5m"`
}

type GRPC struct {
	// gRPC server start parameters section
	// GRPCPort is TCP port to listen by gRPC server
	Port string `mapstructure:"PORT"                 default:"9090"`
}

type HTTP struct {
	// HTTP/REST gateway start parameters section
	// HTTPPort is TCP port to listen by HTTP/REST gateway
	Port string `mapstructure:"PORT"                 default:"8080"`
}
