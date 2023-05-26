package service

import (
	"time"
)

type WB interface {
	GetDetailedReportJSON(start_period time.Time, end_period time.Time) (any, error)
}
