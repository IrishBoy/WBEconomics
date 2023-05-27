package wildberries

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NJRCI6ImE4MTY3ZjM1LTliZTMtNDZmOC05ODIyLTExMTZkMzc0Nzg3YiJ9.PE-HiXefpEBQ-tthksfkkvl5BxJBpiOkmiOL_nKJIY8"
var wbBaseUrl = "https://statistics-api.wildberries.ru"

func PrepareRequest(ctx context.Context, VarHttpMethod string,
	VarPath string, VarOptionals any, VarContentType string, QueryParams url.Values) (*http.Request, error) {
	url, err := url.Parse(wbBaseUrl + VarPath)
	url.RawQuery = QueryParams.Encode()
	req, err := http.NewRequest(VarHttpMethod, url.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("accept", VarContentType)
	req.Header.Set("Authorization", token)
	return req, nil
}

func MakeRequest(request http.Request) (*http.Response, error) {
	client := &http.Client{}
	response, err := client.Do(&request)

	return response, err
}

func ParseResponse(response *http.Response) (string, error) {
	defer response.Body.Close()
	bodyText, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(bodyText), nil

}

func parameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	}

	return fmt.Sprintf("%v", obj)
}
