package config

import (
	"os"
	"testing"
)

func TestDevConfig(t *testing.T) {
	_ = os.Setenv("FRAME_ENV", "DEV")
	conf := GetConfig()
	expectedAddr := ":3003"
	gotEnv := conf.Environment
	gotAddr := conf.Addr

	if gotEnv != EnvType.Dev {
		t.Error("Expected: env", EnvType.Dev, "got:", gotEnv)
	}

	if expectedAddr != gotAddr {
		t.Error("Expected: Addr", expectedAddr, "got:", gotAddr)
	}
}

func TestProdConfig(t *testing.T) {
	_ = os.Setenv("FRAME_ENV", "PROD")
	conf := GetConfig()
	expectedAddr := ":3004"
	gotEnv := conf.Environment
	gotAddr := conf.Addr

	if gotEnv != EnvType.Prod {
		t.Error("Expected: env", EnvType.Prod, "got:", gotEnv)
	}

	if expectedAddr != gotAddr {
		t.Error("Expected: Addr", expectedAddr, "got:", gotAddr)
	}
}
