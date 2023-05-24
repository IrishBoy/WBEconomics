package info_exporter

import (
	"WBEconomics/internal/types"
	"time"
)

func getDetailedReportFile(user int, start_period time.Time, end_period time.Time) ([]types.DetailedOperation, error) {
	return nil, nil
}

func getFinancialReportFile(user int, start_period time.Time, end_period time.Time) (types.FinancialReport, error) {
	return types.FinancialReport{}, nil
}
