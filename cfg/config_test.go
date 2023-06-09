package cfg

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Defaults(t *testing.T) {
	restoreEnv(t, "GRPC_PORT")

	_ = os.Unsetenv("GRPC_PORT")

	cfg, err := Read()
	assert.NoError(t, err)
	assert.Equal(t, "9090", cfg.GRPCService.Port)
}

func Test_ManualSetup(t *testing.T) {
	restoreEnv(t, "GRPC_PORT")

	value := "5090"
	t.Setenv("GRPC_PORT", value)

	cfg, err := Read()
	assert.NoError(t, err)
	assert.Equal(t, value, cfg.GRPCService.Port)
}

func TestLoader_Defaults(t *testing.T) {
	restoreEnv(t, "GRPC_PORT")

	value := "9091"
	t.Setenv("GRPC_PORT", value)

	var cfg Config
	err := read(&cfg)
	assert.NoError(t, err)
	assert.Equal(t, value, cfg.GRPCService.Port)
}

//nolint:thelper
func restoreEnv(t *testing.T, key string) {
	revertValue := os.Getenv(key)
	t.Cleanup(func() {
		if err := os.Setenv(key, revertValue); err != nil {
			t.Fatal(err)
		}
	})
}
