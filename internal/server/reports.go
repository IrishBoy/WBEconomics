package server

import (
	exporter "WBEconomics/internal/info_exporter"
	service "WBEconomics/internal/service"
	types "WBEconomics/internal/types"
	wb_types "WBEconomics/internal/types/wb_types"
	"time"
)

// func main() {

// }

func GetIncomePeriodReport(start_period time.Time, end_period time.Time, source string,
	report_params service.IncomePeriodReportParams) (any, error) {

	if validateTime := periodValidate(start_period, end_period); validateTime != nil {
		return nil, validateTime
	}
	if report_params.Source == "API" {
		options := wb_types.SupplierReportDetailByPeriodOpts{
			Limit: report_params.APIParams.Limit,
			Rrdid: report_params.APIParams.Rrdid,
		}
		report, _, err := GetIncomePeriodFromApi(start_period, end_period, options)
		if err != nil {
			return nil, err
		}
		return report, nil
	} else if report_params.Source == "FILE" {
		report, err := GetIncomePeriodFromFile(start_period, end_period)
		if err != nil {
			return nil, err
		}
		return report, nil
	}
	incomeReport := types.OutputReport{}
	return incomeReport, nil
	// continue
}

func GetIncomePeriodFromApi(start_period time.Time, end_period time.Time,
	Options wb_types.SupplierReportDetailByPeriodOpts) ([]wb_types.DetailReportItem, any, error) {
	detailedReport, responseBody, err := exporter.GetDetailedReportAPI(start_period, end_period, Options)
	if err != nil {
		return nil, nil, err
	}

	return detailedReport, responseBody, nil
}

func GetIncomePeriodFromFile(start_period time.Time, end_period time.Time,
) ([]types.DetailedOperation, error) {
	detailedReport, err := exporter.GetDetailedReportFile(start_period, end_period)
	if err != nil {
		return nil, err
	}
	return detailedReport, err
}
