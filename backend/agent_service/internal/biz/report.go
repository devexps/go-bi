package biz

import (
	"context"
	"github.com/devexps/go-bi/agent_service/internal/data"
	v1 "github.com/devexps/go-bi/api/gen/go/agent_service/v1"
	"github.com/devexps/go-micro/v2/log"
)

type ReportUseCase interface {
	PostReport(ctx context.Context, req *v1.PostReportRequest) error
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

func (r reportUseCase) PostReport(ctx context.Context, req *v1.PostReportRequest) error {
	return r.repo.Publish(ctx, req.GetEventName(), req.GetContent())
}
