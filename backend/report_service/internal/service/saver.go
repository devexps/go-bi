package service

import (
	"context"
	v1 "github.com/devexps/go-bi/api/gen/go/report_service/v1"
	"github.com/devexps/go-bi/report_service/internal/data"
	"github.com/devexps/go-micro/v2/broker"
	"github.com/devexps/go-micro/v2/log"
)

type SaverService interface {
	SaveEventReport(ctx context.Context, topic string, headers broker.Headers, msg *v1.RealTimeWarehousingData) error
}

type saverService struct {
	log *log.Helper

	realtimeRepo data.RealTimeWarehousingRepo
}

// NewSaverService new a saver service.
func NewSaverService(logger log.Logger, realtimeRepo data.RealTimeWarehousingRepo) SaverService {
	l := log.NewHelper(log.With(logger, "module", "report_service/service/saver"))
	return &saverService{
		log:          l,
		realtimeRepo: realtimeRepo,
	}
}

// SaveEventReport .
func (s *saverService) SaveEventReport(ctx context.Context, topic string, headers broker.Headers, msg *v1.RealTimeWarehousingData) error {
	s.log.Infof("topic = %v, headers = %v, msg = %v", topic, headers, msg)
	return s.realtimeRepo.Create(ctx, msg)
}
