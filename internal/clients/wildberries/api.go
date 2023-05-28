package wildberries

import (
	wb_types "WBEconomics/internal/types/wb_types"
	"context"
	"fmt"
	"net/url"
	"strings"
)

// type WBDefaultApiService service

func GetSupplierReportDetailByPeriod(ctx context.Context, dateFrom string, dateTo string,
	localVarOptionals *wb_types.SupplierReportDetailByPeriodOpts) ([]wb_types.DetailReportItem, any, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		lovarVarPath        = "/api/v1/supplier/reportDetailByPeriod"
		localVarContentType = "application/json"
		localVarReturnValue []wb_types.DetailReportItem
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
	responseBody, httpresponse, err := ParseResponse(resp)
	if err != nil {
		fmt.Println("Error in ParseResponse")
		return nil, nil, err
	}

	if httpresponse.StatusCode < 300 {
		err = Decode(&localVarReturnValue, responseBody, httpresponse.Header.Get("Content-Type"))
		if err != nil {
			return localVarReturnValue, httpresponse, err
		}
	}

	if httpresponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  responseBody,
			error: httpresponse.Status,
		}
		if httpresponse.StatusCode == 200 {
			var v []wb_types.DetailReportItem
			err = Decode(&v, responseBody, httpresponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, httpresponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, httpresponse, newErr
		}
		if httpresponse.StatusCode == 401 {
			var v wb_types.InlineResponse401_2
			err = Decode(&v, responseBody, httpresponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, httpresponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, httpresponse, newErr
		}
		return localVarReturnValue, httpresponse, newErr
	}

	return localVarReturnValue, responseBody, nil
}
