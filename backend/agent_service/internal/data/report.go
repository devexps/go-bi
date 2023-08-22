package data

import (
	"context"
	reportV1 "github.com/devexps/go-bi/api/gen/go/report_service/v1"
	"github.com/devexps/go-bi/pkg/topic"
	timeUtil "github.com/devexps/go-bi/pkg/util/time"
	"github.com/devexps/go-micro/v2/broker"
	"github.com/devexps/go-micro/v2/log"
	"time"
)

type ReportRepo interface {
	Publish(ctx context.Context, eventName string, content string) error
}

type reportRepo struct {
	data *Data
	log  *log.Helper
}

// NewReportRepo .
func NewReportRepo(data *Data, logger log.Logger) ReportRepo {
	l := log.NewHelper(log.With(logger, "module", "agent_service/data/report"))
	return &reportRepo{
		data: data,
		log:  l,
	}
}

func (r reportRepo) Publish(ctx context.Context, eventName string, content string) error {
	return r.data.kafkaBroker.Publish(topic.EventReportData, reportV1.RealTimeWarehousingData{
		EventName:  eventName,
		ReportData: content,
		CreateTime: timeUtil.UnixMilliToString(time.Now().UnixMilli()),
	}, broker.WithPublishContext(ctx))
}
