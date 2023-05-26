package main

import (
	wb "WBEconomics/generated/wildberries"
	"WBEconomics/internal/service"
	"context"
	"fmt"
)

type API struct {
	wb service.WB
}

func NewApi(
	wb service.WB,
) *API {
	return &API{
		wb: wb,
	}
}

// func (a *API) wbApi() {

// }
func main() {
	// Create an instance of the DefaultApiService
	apiService := &wb.DefaultApiService{}

	// Set the necessary parameters
	dateFrom := "2023-02-02"
	dateTo := "2023-05-02"
	var localVarOptionals *wb.DefaultApiApiV1SupplierReportDetailByPeriodGetOpts // Optional parameters, if any
	// localVarOptionals.Limit= 10

	// Optionally create a context
	ctx := context.TODO()

	// Call the ApiV1SupplierReportDetailByPeriodGet method
	reportItems, response, err := apiService.ApiV1SupplierReportDetailByPeriodGet(ctx, dateFrom, dateTo, localVarOptionals)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Process the retrieved report items
	for _, item := range reportItems {
		fmt.Println(item)
	}

	// Access the response if needed
	fmt.Println(response)
}

// func main() {
// 	fromString := "2023-02-02"
// 	toString := "2023-05-02"
// 	from, _ := time.Parse("2006-01-02", fromString)
// 	to, _ := time.Parse("2006-01-02", toString)
// 	// client := wb.DefaultApiService{}
// 	wbClient := info_exporter.WB_client{}
// 	report, err := wbClient.GetDetailedReportJSON(from, to)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println(report)
// 	}
// }
