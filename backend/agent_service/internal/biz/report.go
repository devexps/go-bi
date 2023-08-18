package biz

import (
	"github.com/devexps/go-bi/agent_service/internal/data"
	"github.com/devexps/go-micro/v2/log"
)

type ReportUseCase interface {
}

type reportUseCase struct {
	repo data.ReportRepo
	log  *log.Helper
}

// NewReportUseCase new a report use case.
func NewReportUseCase(repo data.ReportRepo, logger log.Logger) ReportUseCase {
	l := log.NewHelper(log.With(logger, "module", "agent_service/usecase/report"))
	return &reportUseCase{
		repo: repo,
		log:  l,
	}
}
