/*
StackState API

This API documentation page describes the StackState server API. The StackState UI and CLI use the StackState server API to configure and query StackState.  You can use the API for similar purposes.  Each request sent to the StackState server API must be authenticated using one of the available authentication methods.   *Note that the StackState receiver API, used to send topology, telemetry and traces to StackState, is not described on this page and requires a different authentication method.*  For more information on StackState, refer to the [StackState documentation](https://docs.stackstate.com).

API version: 5.2.0
Contact: info@stackstate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackstate_api

import (
	"encoding/json"
)

// ArgumentMetricStreamId struct for ArgumentMetricStreamId
type ArgumentMetricStreamId struct {
	Type                string `json:"_type"`
	Id                  *int64 `json:"id,omitempty"`
	LastUpdateTimestamp *int64 `json:"lastUpdateTimestamp,omitempty"`
	Parameter           int64  `json:"parameter"`
	Stream              int64  `json:"stream"`
}

// NewArgumentMetricStreamId instantiates a new ArgumentMetricStreamId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArgumentMetricStreamId(type_ string, parameter int64, stream int64) *ArgumentMetricStreamId {
	this := ArgumentMetricStreamId{}
	this.Type = type_
	this.Parameter = parameter
	this.Stream = stream
	return &this
}

// NewArgumentMetricStreamIdWithDefaults instantiates a new ArgumentMetricStreamId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArgumentMetricStreamIdWithDefaults() *ArgumentMetricStreamId {
	this := ArgumentMetricStreamId{}
	return &this
}

// GetType returns the Type field value
func (o *ArgumentMetricStreamId) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ArgumentMetricStreamId) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ArgumentMetricStreamId) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ArgumentMetricStreamId) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArgumentMetricStreamId) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ArgumentMetricStreamId) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ArgumentMetricStreamId) SetId(v int64) {
	o.Id = &v
}

// GetLastUpdateTimestamp returns the LastUpdateTimestamp field value if set, zero value otherwise.
func (o *ArgumentMetricStreamId) GetLastUpdateTimestamp() int64 {
	if o == nil || o.LastUpdateTimestamp == nil {
		var ret int64
		return ret
	}
	return *o.LastUpdateTimestamp
}

// GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArgumentMetricStreamId) GetLastUpdateTimestampOk() (*int64, bool) {
	if o == nil || o.LastUpdateTimestamp == nil {
		return nil, false
	}
	return o.LastUpdateTimestamp, true
}

// HasLastUpdateTimestamp returns a boolean if a field has been set.
func (o *ArgumentMetricStreamId) HasLastUpdateTimestamp() bool {
	if o != nil && o.LastUpdateTimestamp != nil {
		return true
	}

	return false
}

// SetLastUpdateTimestamp gets a reference to the given int64 and assigns it to the LastUpdateTimestamp field.
func (o *ArgumentMetricStreamId) SetLastUpdateTimestamp(v int64) {
	o.LastUpdateTimestamp = &v
}

// GetParameter returns the Parameter field value
func (o *ArgumentMetricStreamId) GetParameter() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value
// and a boolean to check if the value has been set.
func (o *ArgumentMetricStreamId) GetParameterOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameter, true
}

// SetParameter sets field value
func (o *ArgumentMetricStreamId) SetParameter(v int64) {
	o.Parameter = v
}

// GetStream returns the Stream field value
func (o *ArgumentMetricStreamId) GetStream() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Stream
}

// GetStreamOk returns a tuple with the Stream field value
// and a boolean to check if the value has been set.
func (o *ArgumentMetricStreamId) GetStreamOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stream, true
}

// SetStream sets field value
func (o *ArgumentMetricStreamId) SetStream(v int64) {
	o.Stream = v
}

func (o ArgumentMetricStreamId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdateTimestamp != nil {
		toSerialize["lastUpdateTimestamp"] = o.LastUpdateTimestamp
	}
	if true {
		toSerialize["parameter"] = o.Parameter
	}
	if true {
		toSerialize["stream"] = o.Stream
	}
	return json.Marshal(toSerialize)
}

type NullableArgumentMetricStreamId struct {
	value *ArgumentMetricStreamId
	isSet bool
}

func (v NullableArgumentMetricStreamId) Get() *ArgumentMetricStreamId {
	return v.value
}

func (v *NullableArgumentMetricStreamId) Set(val *ArgumentMetricStreamId) {
	v.value = val
	v.isSet = true
}

func (v NullableArgumentMetricStreamId) IsSet() bool {
	return v.isSet
}

func (v *NullableArgumentMetricStreamId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArgumentMetricStreamId(val *ArgumentMetricStreamId) *NullableArgumentMetricStreamId {
	return &NullableArgumentMetricStreamId{value: val, isSet: true}
}

func (v NullableArgumentMetricStreamId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArgumentMetricStreamId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
