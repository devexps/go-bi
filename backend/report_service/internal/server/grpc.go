package server

import (
	"github.com/devexps/go-bi/api/gen/go/common/conf"
	v1 "github.com/devexps/go-bi/api/gen/go/report_service/v1"
	"github.com/devexps/go-bi/pkg/bootstrap"
	"github.com/devexps/go-bi/report_service/internal/service"
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/middleware/logging"
	"github.com/devexps/go-micro/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(cfg *conf.Bootstrap, logger log.Logger,
	reportSvc service.ReportService,
) *grpc.Server {
	srv := bootstrap.CreateGrpcServer(cfg, logging.Server(logger))
	v1.RegisterReportServiceServer(srv, reportSvc)
	return srv
}
