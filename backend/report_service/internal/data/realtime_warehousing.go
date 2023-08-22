package data

import (
	"context"
	v1 "github.com/devexps/go-bi/api/gen/go/report_service/v1"
	"github.com/devexps/go-micro/v2/log"
)

type RealTimeWarehousingRepo interface {
	Create(ctx context.Context, msg *v1.RealTimeWarehousingData) error
}

type realTimeWarehousingRepo struct {
	data *Data
	log  *log.Helper
}

// NewRealTimeWarehousingRepo .
func NewRealTimeWarehousingRepo(data *Data, logger log.Logger) RealTimeWarehousingRepo {
	l := log.NewHelper(log.With(logger, "module", "report_service/repo/realtime_warehousing"))
	return &realTimeWarehousingRepo{
		data: data,
		log:  l,
	}
}

// Create .
func (r *realTimeWarehousingRepo) Create(ctx context.Context, msg *v1.RealTimeWarehousingData) error {
	query := `INSERT INTO realtime_warehousing (event_name, report_data, create_time) VALUES (?, ?, ?)`
	_, err := r.data.db.ExecContext(ctx, query, msg.GetEventName(), msg.GetReportData(), msg.GetCreateTime())
	if err != nil {
		r.log.Error(err)
		return err
	}
	return nil
}
