package data

import (
	"github.com/devexps/go-bi/api/gen/go/common/conf"
	"github.com/devexps/go-micro/broker/kafka/v2"
	"github.com/devexps/go-micro/v2/broker"
	"github.com/devexps/go-micro/v2/log"
)

// Data .
type Data struct {
	log *log.Helper

	kafkaBroker broker.Broker
}

// NewData .
func NewData(cfg *conf.Bootstrap, logger log.Logger, kafkaBroker broker.Broker) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "agent_service/data"))
	d := &Data{
		log:         l,
		kafkaBroker: kafkaBroker,
	}
	return d, func() {
		l.Info("closing the data resources")
		if err := d.kafkaBroker.Disconnect(); err != nil {
			l.Errorf("disconnect kafka broker error: %v", err.Error())
		}
	}, nil
}

// NewKafkaBroker .
func NewKafkaBroker(cfg *conf.Bootstrap, logger log.Logger) broker.Broker {
	l := log.NewHelper(log.With(logger, "module", "agent_service/data/kafka"))

	b := kafka.NewBroker(
		broker.WithAddress(cfg.Data.Kafka.Addrs...),
		broker.WithCodec(cfg.Data.Kafka.Codec),
		broker.WithGlobalTracerProvider(),
		broker.WithGlobalPropagator(),
		kafka.WithAllowAutoTopicCreation(true),
	)
	if b == nil {
		l.Fatalf("new kafka broker failed")
	}
	if err := b.Init(); err != nil {
		l.Fatalf("init kafka broker error: %v", err.Error())
	}
	if err := b.Connect(); err != nil {
		l.Fatalf("connect kafka broker error: %v", err.Error())
	}
	return b
}
