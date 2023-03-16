package cfg

type Config struct {
	GRPCService GRPC   `mapstructure:"GRPC"`
	HTTPService HTTP   `mapstructure:"HTTP"`
	Logger      Logger `mapstructure:"LOGGER"`
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

type Logger struct {
	// Log parameters section
	// LogLevel is global log level: Debug(-1), Info(0), Warn(1), Error(2), DPanic(3), Panic(4), Fatal(5)
	LogLevel int `mapstructure:"LEVEL"                 default:"-1"`
	// LogTimeFormat is print time format for logger e.g. 2006-01-02T15:04:05Z07:00
	LogTimeFormat string `mapstructure:"TIME_FORMAT"`
}
