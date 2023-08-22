package server

import (
	"context"
	"fmt"
	"github.com/devexps/go-bi/api/gen/go/common/conf"
	v1 "github.com/devexps/go-bi/api/gen/go/report_service/v1"
	"github.com/devexps/go-bi/pkg/topic"
	"github.com/devexps/go-bi/report_service/internal/service"
	"github.com/devexps/go-micro/transport/kafka/v2"
	"github.com/devexps/go-micro/v2/broker"
	"github.com/devexps/go-micro/v2/log"
)

// NewKafkaServer create a kafka server.
func NewKafkaServer(cfg *conf.Bootstrap, _ log.Logger,
	saverSvc service.SaverService,
) *kafka.Server {
	srv := kafka.NewServer(
		kafka.WithAddress(cfg.Server.Kafka.Addrs),
		kafka.WithCodec("json"),
		kafka.WithGlobalTracerProvider(),
		kafka.WithGlobalPropagator(),
	)

	_ = srv.RegisterSubscriber(context.Background(),
		topic.EventReportData, topic.LoggerSaverQueue, false,
		registerEventReportHandler(saverSvc.SaveEventReport),
		EventReportCreator)
	return srv
}

/////////////////////////////////////////////////////////////////////////////////////

func EventReportCreator() broker.Any { return &v1.RealTimeWarehousingData{} }

type EventReportHandler func(ctx context.Context, topic string, headers broker.Headers, msg *v1.RealTimeWarehousingData) error

func registerEventReportHandler(fnc EventReportHandler) broker.Handler {
	return func(ctx context.Context, event broker.Event) error {
		switch t := event.Message().Body.(type) {
		case *v1.RealTimeWarehousingData:
			if err := fnc(ctx, event.Topic(), event.Message().Headers, t); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported type: %T", t)
		}
		return nil
	}
}
