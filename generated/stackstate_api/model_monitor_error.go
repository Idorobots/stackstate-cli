/*
StackState API

StackState's API specification

API version: 0.0.1
Contact: info@stackstate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackstate_api

import (
	"encoding/json"
)

// MonitorError struct for MonitorError
type MonitorError struct {
	Error string `json:"error"`
	Count int32  `json:"count"`
}

// NewMonitorError instantiates a new MonitorError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorError(error_ string, count int32) *MonitorError {
	this := MonitorError{}
	this.Error = error_
	this.Count = count
	return &this
}

// NewMonitorErrorWithDefaults instantiates a new MonitorError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorErrorWithDefaults() *MonitorError {
	this := MonitorError{}
	return &this
}

// GetError returns the Error field value
func (o *MonitorError) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *MonitorError) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *MonitorError) SetError(v string) {
	o.Error = v
}

// GetCount returns the Count field value
func (o *MonitorError) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *MonitorError) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *MonitorError) SetCount(v int32) {
	o.Count = v
}

func (o MonitorError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["error"] = o.Error
	}
	if true {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableMonitorError struct {
	value *MonitorError
	isSet bool
}

func (v NullableMonitorError) Get() *MonitorError {
	return v.value
}

func (v *NullableMonitorError) Set(val *MonitorError) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorError) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorError(val *MonitorError) *NullableMonitorError {
	return &NullableMonitorError{value: val, isSet: true}
}

func (v NullableMonitorError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
