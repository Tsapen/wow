package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/caarlos0/env/v9"

	"github.com/Tsapen/wow/internal/wow"
)

type serverEnvs struct {
	RootDir string `env:"WOW_ROOT_DIR"`
	Config  string `env:"WOW_SERVER_CONFIG"`
}

type clientEnvs struct {
	RootDir string `env:"WOW_ROOT_DIR"`
	Config  string `env:"WOW_CLIENT_CONFIG"`
}

type ServerConfig struct {
	Address            string        `json:"address"`
	ConnectionMaxCount int           `json:"connections_max_count"`
	Timeout            time.Duration `json:"timeout"`
}

type ClientConfig struct {
	Address string        `json:"address"`
	Timeout time.Duration `json:"timeout"`
}

func GetForServer() (*ServerConfig, error) {
	envs := new(serverEnvs)
	if err := env.Parse(envs); err != nil {
		return nil, fmt.Errorf("failed to get envs: %w", err)
	}

	cfg := new(ServerConfig)
	if err := readFromEnv(path.Join(envs.RootDir, envs.Config), cfg); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	return cfg, nil
}

func GetForClient() (*ClientConfig, error) {
	envs := new(clientEnvs)
	if err := env.Parse(envs); err != nil {
		return nil, fmt.Errorf("failed to get envs: %w", err)
	}

	cfg := new(ClientConfig)
	if err := readFromEnv(path.Join(envs.RootDir, envs.Config), cfg); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	return cfg, nil
}

func readFromEnv(filepath string, receiver any) (err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("open file %s: %w", filepath, err)
	}

	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			err = wow.HandleErrPair(fmt.Errorf("close file: %w", closeErr), err)
		}
	}()

	if err = json.NewDecoder(file).Decode(receiver); err != nil {
		return fmt.Errorf("decode file: %w", err)
	}

	return
}
