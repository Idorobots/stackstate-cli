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

// ValueTuple struct for ValueTuple
type ValueTuple struct {
	Time int64 `json:"time"`
	Value string `json:"value"`
}

// NewValueTuple instantiates a new ValueTuple object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValueTuple(time int64, value string) *ValueTuple {
	this := ValueTuple{}
	this.Time = time
	this.Value = value
	return &this
}

// NewValueTupleWithDefaults instantiates a new ValueTuple object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValueTupleWithDefaults() *ValueTuple {
	this := ValueTuple{}
	return &this
}

// GetTime returns the Time field value
func (o *ValueTuple) GetTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *ValueTuple) GetTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *ValueTuple) SetTime(v int64) {
	o.Time = v
}

// GetValue returns the Value field value
func (o *ValueTuple) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ValueTuple) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ValueTuple) SetValue(v string) {
	o.Value = v
}

func (o ValueTuple) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["time"] = o.Time
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableValueTuple struct {
	value *ValueTuple
	isSet bool
}

func (v NullableValueTuple) Get() *ValueTuple {
	return v.value
}

func (v *NullableValueTuple) Set(val *ValueTuple) {
	v.value = val
	v.isSet = true
}

func (v NullableValueTuple) IsSet() bool {
	return v.isSet
}

func (v *NullableValueTuple) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValueTuple(val *ValueTuple) *NullableValueTuple {
	return &NullableValueTuple{value: val, isSet: true}
}

func (v NullableValueTuple) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValueTuple) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


