package data

import (
	"github.com/devexps/go-micro/v2/log"
)

type ReportRepo interface {
}

type reportRepo struct {
	data *Data
	log  *log.Helper
}

// NewReportRepo .
func NewReportRepo(data *Data, logger log.Logger) ReportRepo {
	return &reportRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
