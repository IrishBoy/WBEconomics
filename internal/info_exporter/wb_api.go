package info_exporter

import (
	wb "WBEconomics/generated/wildberries"
	"WBEconomics/internal/types"
	"context"
	"time"
)

type WB_client struct {
	client wb.DefaultApiService
}

func (w WB_client) GetDetailedReportJSON(start_period time.Time, end_period time.Time) (any, error) {
	detailedReport, _, _ := w.client.ApiV1SupplierReportDetailByPeriodGet(context.TODO(), start_period.String(), end_period.String(), nil)
	// detailedReport := "34123"
	print(detailedReport)
	return detailedReport, nil
}

func GetFinancialReportJSON(user int, start_period time.Time, end_period time.Time) (types.FinancialReport, error) {
	return types.FinancialReport{}, nil
}
