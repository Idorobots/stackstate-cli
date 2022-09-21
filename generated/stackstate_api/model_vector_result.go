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

// VectorResult struct for VectorResult
type VectorResult struct {
	Metric map[string]string `json:"metric"`
	Values ValueTuple `json:"values"`
}

// NewVectorResult instantiates a new VectorResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVectorResult(metric map[string]string, values ValueTuple) *VectorResult {
	this := VectorResult{}
	this.Metric = metric
	this.Values = values
	return &this
}

// NewVectorResultWithDefaults instantiates a new VectorResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVectorResultWithDefaults() *VectorResult {
	this := VectorResult{}
	return &this
}

// GetMetric returns the Metric field value
func (o *VectorResult) GetMetric() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *VectorResult) GetMetricOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value
func (o *VectorResult) SetMetric(v map[string]string) {
	o.Metric = v
}

// GetValues returns the Values field value
func (o *VectorResult) GetValues() ValueTuple {
	if o == nil {
		var ret ValueTuple
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *VectorResult) GetValuesOk() (*ValueTuple, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value
func (o *VectorResult) SetValues(v ValueTuple) {
	o.Values = v
}

func (o VectorResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["metric"] = o.Metric
	}
	if true {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableVectorResult struct {
	value *VectorResult
	isSet bool
}

func (v NullableVectorResult) Get() *VectorResult {
	return v.value
}

func (v *NullableVectorResult) Set(val *VectorResult) {
	v.value = val
	v.isSet = true
}

func (v NullableVectorResult) IsSet() bool {
	return v.isSet
}

func (v *NullableVectorResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVectorResult(val *VectorResult) *NullableVectorResult {
	return &NullableVectorResult{value: val, isSet: true}
}

func (v NullableVectorResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVectorResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


