// Code generated by "requestgen -method DELETE -url /api/v3/order -type CancelOrderRequest -responseType .Order"; DO NOT EDIT.

package v3

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/OvictorVieira/bbgo/pkg/exchange/max/maxapi"
	"net/url"
	"reflect"
	"regexp"
)

func (c *CancelOrderRequest) Id(id uint64) *CancelOrderRequest {
	c.id = &id
	return c
}

func (c *CancelOrderRequest) ClientOrderID(clientOrderID string) *CancelOrderRequest {
	c.clientOrderID = &clientOrderID
	return c
}

// GetQueryParameters builds and checks the query parameters and returns url.Values
func (c *CancelOrderRequest) GetQueryParameters() (url.Values, error) {
	var params = map[string]interface{}{}

	query := url.Values{}
	for _k, _v := range params {
		query.Add(_k, fmt.Sprintf("%v", _v))
	}

	return query, nil
}

// GetParameters builds and checks the parameters and return the result in a map object
func (c *CancelOrderRequest) GetParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}
	// check id field -> json key id
	if c.id != nil {
		id := *c.id

		// assign parameter of id
		params["id"] = id
	} else {
	}
	// check clientOrderID field -> json key client_oid
	if c.clientOrderID != nil {
		clientOrderID := *c.clientOrderID

		// assign parameter of clientOrderID
		params["client_oid"] = clientOrderID
	} else {
	}

	return params, nil
}

// GetParametersQuery converts the parameters from GetParameters into the url.Values format
func (c *CancelOrderRequest) GetParametersQuery() (url.Values, error) {
	query := url.Values{}

	params, err := c.GetParameters()
	if err != nil {
		return query, err
	}

	for _k, _v := range params {
		if c.isVarSlice(_v) {
			c.iterateSlice(_v, func(it interface{}) {
				query.Add(_k+"[]", fmt.Sprintf("%v", it))
			})
		} else {
			query.Add(_k, fmt.Sprintf("%v", _v))
		}
	}

	return query, nil
}

// GetParametersJSON converts the parameters from GetParameters into the JSON format
func (c *CancelOrderRequest) GetParametersJSON() ([]byte, error) {
	params, err := c.GetParameters()
	if err != nil {
		return nil, err
	}

	return json.Marshal(params)
}

// GetSlugParameters builds and checks the slug parameters and return the result in a map object
func (c *CancelOrderRequest) GetSlugParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}

	return params, nil
}

func (c *CancelOrderRequest) applySlugsToUrl(url string, slugs map[string]string) string {
	for _k, _v := range slugs {
		needleRE := regexp.MustCompile(":" + _k + "\\b")
		url = needleRE.ReplaceAllString(url, _v)
	}

	return url
}

func (c *CancelOrderRequest) iterateSlice(slice interface{}, _f func(it interface{})) {
	sliceValue := reflect.ValueOf(slice)
	for _i := 0; _i < sliceValue.Len(); _i++ {
		it := sliceValue.Index(_i).Interface()
		_f(it)
	}
}

func (c *CancelOrderRequest) isVarSlice(_v interface{}) bool {
	rt := reflect.TypeOf(_v)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	}
	return false
}

func (c *CancelOrderRequest) GetSlugsMap() (map[string]string, error) {
	slugs := map[string]string{}
	params, err := c.GetSlugParameters()
	if err != nil {
		return slugs, nil
	}

	for _k, _v := range params {
		slugs[_k] = fmt.Sprintf("%v", _v)
	}

	return slugs, nil
}

func (c *CancelOrderRequest) Do(ctx context.Context) (*max.Order, error) {

	params, err := c.GetParameters()
	if err != nil {
		return nil, err
	}
	query := url.Values{}

	apiURL := "/api/v3/order"

	req, err := c.client.NewAuthenticatedRequest(ctx, "DELETE", apiURL, query, params)
	if err != nil {
		return nil, err
	}

	response, err := c.client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var apiResponse max.Order
	if err := response.DecodeJSON(&apiResponse); err != nil {
		return nil, err
	}
	return &apiResponse, nil
}
