package service

import (
	"context"
	"github.com/devexps/go-bi/agent_service/internal/biz"
	v1 "github.com/devexps/go-bi/api/gen/go/agent_service/v1"
	"github.com/devexps/go-micro/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ReportService interface {
	v1.ReportServiceServer
}

type reportService struct {
	v1.UnimplementedReportServiceServer

	log *log.Helper
	uc  biz.ReportUseCase
}

// NewReportService new a report service.
func NewReportService(logger log.Logger, uc biz.ReportUseCase) ReportService {
	l := log.NewHelper(log.With(logger, "module", "agent_service/service/report"))
	return &reportService{
		log: l,
		uc:  uc,
	}
}

func (s *reportService) PostReport(context.Context, *v1.PostReportRequest) (*v1.PostReportReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostReport not implemented")
}
