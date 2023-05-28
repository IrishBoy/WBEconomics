package wildberries

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

// type APIClient struct {
// 	cfg    *Configuration
// 	common service // Reuse a single struct instead of allocating one for each service on the heap.

// 	// API Services

// 	DefaultApi *DefaultApiService

// 	Marketplace_Api *Marketplace_ApiService

// 	_Api *_ApiService
// }

// type service struct {
// 	client *APIClient
// }

type GenericSwaggerError struct {
	body  []byte
	error string
	model interface{}
}

func (e GenericSwaggerError) Error() string {
	return e.error
}

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

func ParseResponse(response *http.Response) ([]byte, *http.Response, error) {

	bodyText, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		fmt.Println(err)
		return nil, response, err
	}
	return bodyText, response, nil

}

func Decode(VarOut interface{}, response []byte, contentType string) (err error) {
	if strings.Contains(contentType, "application/xml") {
		if err = xml.Unmarshal(response, VarOut); err != nil {
			return err
		}
		return nil
	} else if strings.Contains(contentType, "application/json") {
		if err = json.Unmarshal(response, VarOut); err != nil {
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
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
