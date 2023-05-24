package info_exporter

import (
	"WBEconomics/internal/types"
	"time"
)

func getDetailedReportJSON(user int, start_period time.Time, end_period time.Time) ([]types.DetailedOperation, error) {
	return nil, nil
}

func getFinancialReportJSON(user int, start_period time.Time, end_period time.Time) (types.FinancialReport, error) {
	return types.FinancialReport{}, nil
}
