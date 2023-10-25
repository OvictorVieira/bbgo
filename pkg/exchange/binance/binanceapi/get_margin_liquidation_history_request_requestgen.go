// Code generated by "requestgen -method GET -url /sapi/v1/margin/forceLiquidationRec -type GetMarginLiquidationHistoryRequest -responseType .RowsResponse -responseDataField Rows -responseDataType []MarginLiquidationRecord"; DO NOT EDIT.

package binanceapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

func (g *GetMarginLiquidationHistoryRequest) IsolatedSymbol(isolatedSymbol string) *GetMarginLiquidationHistoryRequest {
	g.isolatedSymbol = &isolatedSymbol
	return g
}

func (g *GetMarginLiquidationHistoryRequest) StartTime(startTime time.Time) *GetMarginLiquidationHistoryRequest {
	g.startTime = &startTime
	return g
}

func (g *GetMarginLiquidationHistoryRequest) EndTime(endTime time.Time) *GetMarginLiquidationHistoryRequest {
	g.endTime = &endTime
	return g
}

func (g *GetMarginLiquidationHistoryRequest) Size(size int) *GetMarginLiquidationHistoryRequest {
	g.size = &size
	return g
}

func (g *GetMarginLiquidationHistoryRequest) Current(current int) *GetMarginLiquidationHistoryRequest {
	g.current = &current
	return g
}

// GetQueryParameters builds and checks the query parameters and returns url.Values
func (g *GetMarginLiquidationHistoryRequest) GetQueryParameters() (url.Values, error) {
	var params = map[string]interface{}{}

	query := url.Values{}
	for _k, _v := range params {
		query.Add(_k, fmt.Sprintf("%v", _v))
	}

	return query, nil
}

// GetParameters builds and checks the parameters and return the result in a map object
func (g *GetMarginLiquidationHistoryRequest) GetParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}
	// check isolatedSymbol field -> json key isolatedSymbol
	if g.isolatedSymbol != nil {
		isolatedSymbol := *g.isolatedSymbol

		// assign parameter of isolatedSymbol
		params["isolatedSymbol"] = isolatedSymbol
	} else {
	}
	// check startTime field -> json key startTime
	if g.startTime != nil {
		startTime := *g.startTime

		// assign parameter of startTime
		// convert time.Time to milliseconds time stamp
		params["startTime"] = strconv.FormatInt(startTime.UnixNano()/int64(time.Millisecond), 10)
	} else {
	}
	// check endTime field -> json key endTime
	if g.endTime != nil {
		endTime := *g.endTime

		// assign parameter of endTime
		// convert time.Time to milliseconds time stamp
		params["endTime"] = strconv.FormatInt(endTime.UnixNano()/int64(time.Millisecond), 10)
	} else {
	}
	// check size field -> json key size
	if g.size != nil {
		size := *g.size

		// assign parameter of size
		params["size"] = size
	} else {
	}
	// check current field -> json key current
	if g.current != nil {
		current := *g.current

		// assign parameter of current
		params["current"] = current
	} else {
	}

	return params, nil
}

// GetParametersQuery converts the parameters from GetParameters into the url.Values format
func (g *GetMarginLiquidationHistoryRequest) GetParametersQuery() (url.Values, error) {
	query := url.Values{}

	params, err := g.GetParameters()
	if err != nil {
		return query, err
	}

	for _k, _v := range params {
		if g.isVarSlice(_v) {
			g.iterateSlice(_v, func(it interface{}) {
				query.Add(_k+"[]", fmt.Sprintf("%v", it))
			})
		} else {
			query.Add(_k, fmt.Sprintf("%v", _v))
		}
	}

	return query, nil
}

// GetParametersJSON converts the parameters from GetParameters into the JSON format
func (g *GetMarginLiquidationHistoryRequest) GetParametersJSON() ([]byte, error) {
	params, err := g.GetParameters()
	if err != nil {
		return nil, err
	}

	return json.Marshal(params)
}

// GetSlugParameters builds and checks the slug parameters and return the result in a map object
func (g *GetMarginLiquidationHistoryRequest) GetSlugParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}

	return params, nil
}

func (g *GetMarginLiquidationHistoryRequest) applySlugsToUrl(url string, slugs map[string]string) string {
	for _k, _v := range slugs {
		needleRE := regexp.MustCompile(":" + _k + "\\b")
		url = needleRE.ReplaceAllString(url, _v)
	}

	return url
}

func (g *GetMarginLiquidationHistoryRequest) iterateSlice(slice interface{}, _f func(it interface{})) {
	sliceValue := reflect.ValueOf(slice)
	for _i := 0; _i < sliceValue.Len(); _i++ {
		it := sliceValue.Index(_i).Interface()
		_f(it)
	}
}

func (g *GetMarginLiquidationHistoryRequest) isVarSlice(_v interface{}) bool {
	rt := reflect.TypeOf(_v)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	}
	return false
}

func (g *GetMarginLiquidationHistoryRequest) GetSlugsMap() (map[string]string, error) {
	slugs := map[string]string{}
	params, err := g.GetSlugParameters()
	if err != nil {
		return slugs, nil
	}

	for _k, _v := range params {
		slugs[_k] = fmt.Sprintf("%v", _v)
	}

	return slugs, nil
}

func (g *GetMarginLiquidationHistoryRequest) Do(ctx context.Context) ([]MarginLiquidationRecord, error) {

	// empty params for GET operation
	var params interface{}
	query, err := g.GetParametersQuery()
	if err != nil {
		return nil, err
	}

	apiURL := "/sapi/v1/margin/forceLiquidationRec"

	req, err := g.client.NewAuthenticatedRequest(ctx, "GET", apiURL, query, params)
	if err != nil {
		return nil, err
	}

	response, err := g.client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var apiResponse RowsResponse
	if err := response.DecodeJSON(&apiResponse); err != nil {
		return nil, err
	}
	var data []MarginLiquidationRecord
	if err := json.Unmarshal(apiResponse.Rows, &data); err != nil {
		return nil, err
	}
	return data, nil
}
