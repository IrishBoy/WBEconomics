package service

import (
	"time"

	"github.com/antihax/optional"
)

type Reports interface {
	APIGetter
	FilesGetter
	GetIncomePeriodReport(start_period time.Time, end_period time.Time, report_params IncomePeriodReportParams)
}
type APIGetter interface {
	GetDetailedReportAPI(start_period time.Time, end_period time.Time) (any, error)
}

type FilesGetter interface {
	GetDetailedReportFiles(start_period time.Time, end_period time.Time) (any, error)
}

type IncomePeriodReportParams struct {
	Source    string
	APIParams APIParams
}

type APIParams struct {
	Limit optional.Int32
	Rrdid optional.Int32
}
