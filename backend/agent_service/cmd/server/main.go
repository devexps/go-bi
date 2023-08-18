package main

import (
	"github.com/devexps/go-bi/pkg/bootstrap"
	"github.com/devexps/go-bi/pkg/service"
	"github.com/devexps/go-micro/v2"
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/registry"
	"github.com/devexps/go-micro/v2/transport/http"
)

// go build -ldflags "-X main.Version=x.y.z"

var (
	Version string
	Service = bootstrap.NewServiceInfo(
		service.AgentService,
		Version,
		"",
	)
)

func newApp(ll log.Logger, hs *http.Server, rr registry.Registrar) *micro.App {
	return micro.New(
		micro.ID(Service.GetInstanceId()),
		micro.Name(Service.Name),
		micro.Version(Service.Version),
		micro.Metadata(Service.Metadata),
		micro.Logger(ll),
		micro.Server(
			hs,
		),
		micro.Registrar(rr),
	)
}

func main() {
	cfg, ll, reg := bootstrap.Bootstrap(Service)

	app, cleanup, err := initApp(ll, reg, cfg)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
