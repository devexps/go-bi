package bootstrap

import (
	"github.com/devexps/go-bi/api/gen/go/common/conf"
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/registry"
)

// Bootstrap App
func Bootstrap(serviceInfo *ServiceInfo) (*conf.Bootstrap, log.Logger, registry.Registrar) {
	// inject command flags
	Flags := NewCommandFlags()
	Flags.Init()

	// load configs
	cfg := LoadBootstrapConfig(Flags.Conf)
	if cfg == nil {
		panic("load config failed")
	}

	// init logger
	ll := NewLoggerProvider(cfg.Logger, serviceInfo)

	// init registrar
	reg := NewRegistry(cfg.Registry)

	// init tracer
	err := NewTracerProvider(cfg.Trace, serviceInfo)
	if err != nil {
		panic(err)
	}

	return cfg, ll, reg
}
