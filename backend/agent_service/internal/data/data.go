package data

import (
	"github.com/devexps/go-bi/api/gen/go/common/conf"
	"github.com/devexps/go-micro/v2/log"
)

// Data .
type Data struct {
	log *log.Helper

	// TODO wrapped database client
}

// NewData .
func NewData(cfg *conf.Bootstrap, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	d := &Data{
		log: l,
	}
	return d, func() {
		l.Info("closing the data resources")
	}, nil
}
