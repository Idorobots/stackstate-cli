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

type ExportApi interface {

	/*
		ExportSettings Export settings

		Export StackState Templated JSON (STJ) setting nodes.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiExportSettingsRequest
	*/
	ExportSettings(ctx context.Context) ApiExportSettingsRequest

	// ExportSettingsExecute executes the request
	//  @return string
	ExportSettingsExecute(r ApiExportSettingsRequest) (string, *http.Response, error)
}

// ExportApiService ExportApi service
type ExportApiService service

type ApiExportSettingsRequest struct {
	ctx        context.Context
	ApiService ExportApi
	export     *Export
}

func (r ApiExportSettingsRequest) Export(export Export) ApiExportSettingsRequest {
	r.export = &export
	return r
}

func (r ApiExportSettingsRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.ExportSettingsExecute(r)
}

/*
ExportSettings Export settings

Export StackState Templated JSON (STJ) setting nodes.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiExportSettingsRequest
*/
func (a *ExportApiService) ExportSettings(ctx context.Context) ApiExportSettingsRequest {
	return ApiExportSettingsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return string
func (a *ExportApiService) ExportSettingsExecute(r ApiExportSettingsRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportApiService.ExportSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/export"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.export == nil {
		return localVarReturnValue, nil, reportError("export is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.export
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

type ExportApiMock struct {
	ExportSettingsCalls    *[]ExportSettingsCall
	ExportSettingsResponse ExportSettingsMockResponse
}

func NewExportApiMock() ExportApiMock {
	xExportSettingsCalls := make([]ExportSettingsCall, 0)
	return ExportApiMock{
		ExportSettingsCalls: &xExportSettingsCalls,
	}
}

type ExportSettingsMockResponse struct {
	Result   string
	Response *http.Response
	Error    error
}

type ExportSettingsCall struct {
	Pexport *Export
}

func (mock ExportApiMock) ExportSettings(ctx context.Context) ApiExportSettingsRequest {
	return ApiExportSettingsRequest{
		ApiService: mock,
		ctx:        ctx,
	}
}

func (mock ExportApiMock) ExportSettingsExecute(r ApiExportSettingsRequest) (string, *http.Response, error) {
	p := ExportSettingsCall{
		Pexport: r.export,
	}
	*mock.ExportSettingsCalls = append(*mock.ExportSettingsCalls, p)
	return mock.ExportSettingsResponse.Result, mock.ExportSettingsResponse.Response, mock.ExportSettingsResponse.Error
}
