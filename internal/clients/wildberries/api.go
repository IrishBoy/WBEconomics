package wildberries

import (
	wb_types "WBEconomics/internal/types/wb_types"
	"context"
	"fmt"
	"net/url"
	"strings"
)

func GetSupplierReportDetailByPeriod(ctx context.Context, dateFrom string, dateTo string,
	localVarOptionals *wb_types.SupplierReportDetailByPeriodOpts) ([]wb_types.DetailReportItem, any, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		lovarVarPath        = "/api/v1/supplier/reportDetailByPeriod"
		localVarContentType = "application/json"
	)
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("dateFrom", parameterToString(dateFrom, ""))
	localVarQueryParams.Add("dateTo", parameterToString(dateTo, ""))

	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}

	if localVarOptionals != nil && localVarOptionals.Rrdid.IsSet() {
		localVarQueryParams.Add("rrdid", parameterToString(localVarOptionals.Rrdid.Value(), ""))
	}
	// return req, nil
	r, err := PrepareRequest(ctx, localVarHttpMethod, lovarVarPath,
		localVarOptionals, localVarContentType, localVarQueryParams)
	if err != nil {
		return nil, nil, err

	}
	resp, err := MakeRequest(*r)
	if err != nil {
		fmt.Println("Error in MakeRequest")
		return nil, nil, err
	}
	response, err := ParseResponse(resp)
	// defer resp.Body.Close()
	// bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error in ParseResponse")
		return nil, nil, err
	}

	return nil, response, nil
}
