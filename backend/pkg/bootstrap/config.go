package bootstrap

import (
	"github.com/devexps/go-bi/api/gen/go/common/conf"
	"github.com/devexps/go-micro/v2/config"
	"github.com/devexps/go-micro/v2/config/file"
)

type ConfigType string

// NewConfigProvider creates a configuration
func NewConfigProvider(configPath string) config.Config {
	return config.New(
		config.WithSource(
			NewFileConfigSource(configPath),
		),
	)
}

// NewFileConfigSource creates a local file configuration source
func NewFileConfigSource(filePath string) config.Source {
	return file.NewSource(filePath)
}

// LoadBootstrapConfig loader bootstrap configuration
func LoadBootstrapConfig(configPath string) *conf.Bootstrap {
	cfg := NewConfigProvider(configPath)
	if err := cfg.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := cfg.Scan(&bc); err != nil {
		panic(err)
	}

	if bc.Server == nil {
		bc.Server = &conf.Server{}
		_ = cfg.Scan(&bc.Server)
	}

	if bc.Data == nil {
		bc.Data = &conf.Data{}
		_ = cfg.Scan(&bc.Data)
	}

	if bc.Auth == nil {
		bc.Auth = &conf.Auth{}
		_ = cfg.Scan(&bc.Auth)
	}

	if bc.Trace == nil {
		bc.Trace = &conf.Tracer{}
		_ = cfg.Scan(&bc.Trace)
	}

	if bc.Logger == nil {
		bc.Logger = &conf.Logger{}
		_ = cfg.Scan(&bc.Logger)
	}

	if bc.Registry == nil {
		bc.Registry = &conf.Registry{}
		_ = cfg.Scan(&bc.Registry)
	}

	return &bc
}
