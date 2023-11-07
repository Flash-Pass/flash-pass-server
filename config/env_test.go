package config

import (
	"context"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	config, err := LoadEnv(context.Background())

	if err != nil {
		panic(err)
	}

	t.Log(config)
}
