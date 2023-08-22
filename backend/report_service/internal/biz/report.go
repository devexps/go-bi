package biz

import (
	"github.com/devexps/go-bi/report_service/internal/data"
	"github.com/devexps/go-micro/v2/log"
)

type ReportUseCase interface {
}

type reportUseCase struct {
	repo data.ReportRepo
	log  *log.Helper
}

// NewReportUseCase new a Greeter use case.
func NewReportUseCase(repo data.ReportRepo, logger log.Logger) ReportUseCase {
	return &reportUseCase{repo: repo, log: log.NewHelper(logger)}
}
