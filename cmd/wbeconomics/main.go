package main

import (
	"WBEconomics/internal/info_exporter"
	"WBEconomics/internal/types/wb_types"
	"fmt"
	"time"
)

// type API struct {
// 	wb service.WB
// }

// func NewApi(
// 	wb service.WB,
// ) *API {
// 	return &API{
// 		wb: wb,
// 	}
// }

// func (a *API) wbApi() {

// }
func main() {
	fromString := "2023-02-02"
	toString := "2023-05-02"
	from, _ := time.Parse("2006-01-02", fromString)
	to, _ := time.Parse("2006-01-02", toString)
	opt := wb_types.SupplierReportDetailByPeriodOpts{}
	report, _, _ := info_exporter.GetDetailedReportAPI(from, to, opt)

	// fmt.Printf("Response: %s\n", response)
	fmt.Printf("Report: %+v\n", report[0])

	// client := &http.Client{}
	// req, err := http.NewRequest("GET", "https://statistics-api.wildberries.ru/api/v1/supplier/reportDetailByPeriod?dateFrom=2019-06-20&limit=100000&dateTo=2023-06-20", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// req.Header.Set("accept", "application/json")
	// req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NJRCI6ImE4MTY3ZjM1LTliZTMtNDZmOC05ODIyLTExMTZkMzc0Nzg3YiJ9.PE-HiXefpEBQ-tthksfkkvl5BxJBpiOkmiOL_nKJIY8")
	// resp, err := client.Do(req)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer resp.Body.Close()
	// bodyText, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s\n", bodyText)
}
