/*
StackState API

This API documentation page describes the StackState server API. The StackState UI and CLI use the StackState server API to configure and query StackState.  You can use the API for similar purposes.  Each request sent to the StackState server API must be authenticated using one of the available authentication methods.   *Note that the StackState receiver API, used to send topology, telemetry and traces to StackState, is not described on this page and requires a different authentication method.*  For more information on StackState, refer to the [StackState documentation](https://docs.stackstate.com).

API version: 5.2.0
Contact: info@stackstate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackstate_api

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type ProblemApi interface {

	/*
		GetProblemCausingEvents List possible events which led to the problem

		Resulting events are ordered by likeness to be an actual problem cause

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param problemId The problem id number.
		@return ApiGetProblemCausingEventsRequest
	*/
	GetProblemCausingEvents(ctx context.Context, problemId int64) ApiGetProblemCausingEventsRequest

	// GetProblemCausingEventsExecute executes the request
	//  @return GetCausingEventsResult
	GetProblemCausingEventsExecute(r ApiGetProblemCausingEventsRequest) (*GetCausingEventsResult, *http.Response, error)
}

// ProblemApiService ProblemApi service
type ProblemApiService service

type ApiGetProblemCausingEventsRequest struct {
	ctx          context.Context
	ApiService   ProblemApi
	problemId    int64
	topologyTime *int32
	limit        *int32
}

// A timestamp at which resources will be queried. If not given the resources are quired for current time.
func (r ApiGetProblemCausingEventsRequest) TopologyTime(topologyTime int32) ApiGetProblemCausingEventsRequest {
	r.topologyTime = &topologyTime
	return r
}

// Maximum number of resources to be returned in result.
func (r ApiGetProblemCausingEventsRequest) Limit(limit int32) ApiGetProblemCausingEventsRequest {
	r.limit = &limit
	return r
}

func (r ApiGetProblemCausingEventsRequest) Execute() (*GetCausingEventsResult, *http.Response, error) {
	return r.ApiService.GetProblemCausingEventsExecute(r)
}

/*
GetProblemCausingEvents List possible events which led to the problem

Resulting events are ordered by likeness to be an actual problem cause

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param problemId The problem id number.
 @return ApiGetProblemCausingEventsRequest
*/
func (a *ProblemApiService) GetProblemCausingEvents(ctx context.Context, problemId int64) ApiGetProblemCausingEventsRequest {
	return ApiGetProblemCausingEventsRequest{
		ApiService: a,
		ctx:        ctx,
		problemId:  problemId,
	}
}

// Execute executes the request
//  @return GetCausingEventsResult
func (a *ProblemApiService) GetProblemCausingEventsExecute(r ApiGetProblemCausingEventsRequest) (*GetCausingEventsResult, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetCausingEventsResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProblemApiService.GetProblemCausingEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/problems/{problemId}/causing-events"
	localVarPath = strings.Replace(localVarPath, "{"+"problemId"+"}", url.PathEscape(parameterToString(r.problemId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.topologyTime != nil {
		localVarQueryParams.Add("topologyTime", parameterToString(*r.topologyTime, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Token"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ServiceBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-ServiceBearer"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ServiceToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v GetCausingEventsBadRequest
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v GetCausingEventsNotFound
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GenericErrorsResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v GetCausingEventsServiceUnaivailable
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// ---------------------------------------------
// ------------------ MOCKS --------------------
// ---------------------------------------------

type ProblemApiMock struct {
	GetProblemCausingEventsCalls    *[]GetProblemCausingEventsCall
	GetProblemCausingEventsResponse GetProblemCausingEventsMockResponse
}

func NewProblemApiMock() ProblemApiMock {
	xGetProblemCausingEventsCalls := make([]GetProblemCausingEventsCall, 0)
	return ProblemApiMock{
		GetProblemCausingEventsCalls: &xGetProblemCausingEventsCalls,
	}
}

type GetProblemCausingEventsMockResponse struct {
	Result   GetCausingEventsResult
	Response *http.Response
	Error    error
}

type GetProblemCausingEventsCall struct {
	PproblemId    int64
	PtopologyTime *int32
	Plimit        *int32
}

func (mock ProblemApiMock) GetProblemCausingEvents(ctx context.Context, problemId int64) ApiGetProblemCausingEventsRequest {
	return ApiGetProblemCausingEventsRequest{
		ApiService: mock,
		ctx:        ctx,
		problemId:  problemId,
	}
}

func (mock ProblemApiMock) GetProblemCausingEventsExecute(r ApiGetProblemCausingEventsRequest) (*GetCausingEventsResult, *http.Response, error) {
	p := GetProblemCausingEventsCall{
		PproblemId:    r.problemId,
		PtopologyTime: r.topologyTime,
		Plimit:        r.limit,
	}
	*mock.GetProblemCausingEventsCalls = append(*mock.GetProblemCausingEventsCalls, p)
	return &mock.GetProblemCausingEventsResponse.Result, mock.GetProblemCausingEventsResponse.Response, mock.GetProblemCausingEventsResponse.Error
}
