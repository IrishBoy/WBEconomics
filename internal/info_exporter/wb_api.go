package info_exporter

import (
	wb "WBEconomics/internal/clients/wildberries"
	types "WBEconomics/internal/types"
	wb_types "WBEconomics/internal/types/wb_types"
	"context"
	"time"
)

// type WB_client struct {
// 	client wb.DefaultApiService
// }

func GetDetailedReportJSON(start_period time.Time, end_period time.Time, Options wb_types.SupplierReportDetailByPeriodOpts) (any, error) {
	_, detailedReport, _ := wb.GetSupplierReportDetailByPeriod(context.TODO(), start_period.String(), end_period.String(), &Options)
	// fmt.Printf("%s\n", detailedReport)
	return detailedReport, nil
}
func GetFinancialReportJSON(user int, start_period time.Time, end_period time.Time) (types.FinancialReport, error) {
	return types.FinancialReport{}, nil
}
