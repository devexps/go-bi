// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/devexps/go-bi/agent_service/internal/biz"
	"github.com/devexps/go-bi/agent_service/internal/data"
	"github.com/devexps/go-bi/agent_service/internal/server"
	"github.com/devexps/go-bi/agent_service/internal/service"
	"github.com/devexps/go-bi/api/gen/go/common/conf"
	"github.com/devexps/go-micro/v2"
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/registry"
)

// Injectors from wire.go:

// initApp init micro application.
func initApp(logger log.Logger, registrar registry.Registrar, bootstrap *conf.Bootstrap) (*micro.App, func(), error) {
	dataData, cleanup, err := data.NewData(bootstrap, logger)
	if err != nil {
		return nil, nil, err
	}
	reportRepo := data.NewReportRepo(dataData, logger)
	reportUseCase := biz.NewReportUseCase(reportRepo, logger)
	reportService := service.NewReportService(logger, reportUseCase)
	httpServer := server.NewHTTPServer(bootstrap, logger, reportService)
	app := newApp(logger, httpServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}