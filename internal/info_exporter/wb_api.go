package info_exporter

import (
	wb "WBEconomics/internal/clients/wildberries"
	wb_types "WBEconomics/internal/types/wb_types"
	"context"
	"time"
)

func GetDetailedReportAPI(start_period time.Time, end_period time.Time, Options wb_types.SupplierReportDetailByPeriodOpts) ([]wb_types.DetailReportItem, any, error) {
	detailedReport, responseBody, _ := wb.GetSupplierReportDetailByPeriod(context.TODO(), start_period.String(), end_period.String(), &Options)
	// fmt.Printf("%s\n", detailedReport)
	return detailedReport, responseBody, nil
}
