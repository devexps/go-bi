package service

import (
	v1 "github.com/devexps/go-bi/api/gen/go/report_service/v1"
	"github.com/devexps/go-bi/report_service/internal/biz"
)

type ReportService interface {
	v1.ReportServiceServer
}

type reportService struct {
	v1.UnimplementedReportServiceServer

	uc biz.ReportUseCase
}

// NewReportService new a report service.
func NewReportService(uc biz.ReportUseCase) ReportService {
	return &reportService{uc: uc}
}
