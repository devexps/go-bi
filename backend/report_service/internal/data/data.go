package data

import (
	"database/sql"
	"github.com/devexps/go-bi/api/gen/go/common/conf"
	"github.com/devexps/go-micro/v2/log"

	// init sql driver
	_ "github.com/ClickHouse/clickhouse-go/v2"
)

// Data .
type Data struct {
	log *log.Helper

	db *sql.DB
}

// NewData .
func NewData(logger log.Logger, db *sql.DB) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "report_service/data"))
	d := &Data{
		log: l,
		db:  db,
	}
	return d, func() {
		l.Info("closing the data resources")
		if err := d.db.Close(); err != nil {
			l.Error(err)
		}
	}, nil
}

// NewClickHouseClient creates new ent client
func NewClickHouseClient(logger log.Logger, cfg *conf.Bootstrap) *sql.DB {
	l := log.NewHelper(log.With(logger, "module", "report_service/data/clickhouse"))

	client, err := sql.Open(cfg.Data.Database.Driver, cfg.Data.Database.Source)
	if err != nil {
		l.Fatalf("failed opening connection to db: %v", err)
	}
	if err := client.Ping(); err != nil {
		l.Fatalf("failed pinging to db: %v", err)
	}
	return client
}
