package server

import (
	"github.com/devexps/go-bi/agent_service/internal/service"
	v1 "github.com/devexps/go-bi/api/gen/go/agent_service/v1"
	"github.com/devexps/go-bi/api/gen/go/common/conf"
	"github.com/devexps/go-bi/pkg/bootstrap"
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/middleware/logging"
	"github.com/devexps/go-micro/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(cfg *conf.Bootstrap, logger log.Logger,
	reportSvc service.ReportService,
) *http.Server {
	srv := bootstrap.CreateHttpServer(cfg, logging.Server(logger))
	v1.RegisterReportServiceHTTPServer(srv, reportSvc)
	return srv
}
