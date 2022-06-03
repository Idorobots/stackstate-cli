/*
StackState API

StackState's API specification

API version: 0.0.1
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
)


type ExportAnomalyApi interface {

	/*
	ExportAnomaly Export anomalies with metric history and feedback

	

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiExportAnomalyRequest
	*/
	ExportAnomaly(ctx context.Context) ApiExportAnomalyRequest

	// ExportAnomalyExecute executes the request
	//  @return []AnomalyWithContext
	ExportAnomalyExecute(r ApiExportAnomalyRequest) ([]AnomalyWithContext, *http.Response, error)
}

// ExportAnomalyApiService ExportAnomalyApi service
type ExportAnomalyApiService service

type ApiExportAnomalyRequest struct {
	ctx context.Context
	ApiService ExportAnomalyApi
	startTime *int64
	feedback *string
	endTime *int64
	history *int64
}

// Beginning of timerange of to be exported anomalies.  Timestamp in unix millis.
func (r ApiExportAnomalyRequest) StartTime(startTime int64) ApiExportAnomalyRequest {
	r.startTime = &startTime
	return r
}

// Type of filtering to do on feedback.  Filtering on feedback is currently mandatory, with only the &#39;present&#39; value being supporeted (feedback is available).
func (r ApiExportAnomalyRequest) Feedback(feedback string) ApiExportAnomalyRequest {
	r.feedback = &feedback
	return r
}

// End of timerange of to be exported anomalies.  Timestamp in unix millis.
func (r ApiExportAnomalyRequest) EndTime(endTime int64) ApiExportAnomalyRequest {
	r.endTime = &endTime
	return r
}

// Amount of historic data, leading up to the anomaly, to be exported.  Duration in unix millis.
func (r ApiExportAnomalyRequest) History(history int64) ApiExportAnomalyRequest {
	r.history = &history
	return r
}

func (r ApiExportAnomalyRequest) Execute() ([]AnomalyWithContext, *http.Response, error) {
	return r.ApiService.ExportAnomalyExecute(r)
}

/*
ExportAnomaly Export anomalies with metric history and feedback



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiExportAnomalyRequest
*/
func (a *ExportAnomalyApiService) ExportAnomaly(ctx context.Context) ApiExportAnomalyRequest {
	return ApiExportAnomalyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []AnomalyWithContext
func (a *ExportAnomalyApiService) ExportAnomalyExecute(r ApiExportAnomalyRequest) ([]AnomalyWithContext, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []AnomalyWithContext
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportAnomalyApiService.ExportAnomaly")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/anomalies/export"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.startTime == nil {
		return localVarReturnValue, nil, reportError("startTime is required and must be specified")
	}
	if r.feedback == nil {
		return localVarReturnValue, nil, reportError("feedback is required and must be specified")
	}

	localVarQueryParams.Add("startTime", parameterToString(*r.startTime, ""))
	if r.endTime != nil {
		localVarQueryParams.Add("endTime", parameterToString(*r.endTime, ""))
	}
	if r.history != nil {
		localVarQueryParams.Add("history", parameterToString(*r.history, ""))
	}
	localVarQueryParams.Add("feedback", parameterToString(*r.feedback, ""))
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
		if localVarHTTPResponse.StatusCode == 413 {
			var v TooManyAnomaliesError
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


type ExportAnomalyApiMock struct {
	ExportAnomalyCalls *[]ExportAnomalyCall
	ExportAnomalyResponse ExportAnomalyMockResponse
}	

func NewExportAnomalyApiMock() ExportAnomalyApiMock {
	xExportAnomalyCalls := make([]ExportAnomalyCall, 0)
	return ExportAnomalyApiMock {
		ExportAnomalyCalls: &xExportAnomalyCalls,
	}
}

type ExportAnomalyMockResponse struct {
	Result []AnomalyWithContext
	Response *http.Response
	Error error
}

type ExportAnomalyCall struct {
	PstartTime *int64
	Pfeedback *string
	PendTime *int64
	Phistory *int64
}


func (mock ExportAnomalyApiMock) ExportAnomaly(ctx context.Context) ApiExportAnomalyRequest {
	return ApiExportAnomalyRequest{
		ApiService: mock,
		ctx: ctx,
	}
}

func (mock ExportAnomalyApiMock) ExportAnomalyExecute(r ApiExportAnomalyRequest) ([]AnomalyWithContext, *http.Response, error) {
	p := ExportAnomalyCall {
			PstartTime: r.startTime,
			Pfeedback: r.feedback,
			PendTime: r.endTime,
			Phistory: r.history,
	}
	*mock.ExportAnomalyCalls = append(*mock.ExportAnomalyCalls, p)
	return mock.ExportAnomalyResponse.Result, mock.ExportAnomalyResponse.Response, mock.ExportAnomalyResponse.Error
}


