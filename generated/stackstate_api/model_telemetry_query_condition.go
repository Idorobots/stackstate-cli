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

// TelemetryQueryCondition struct for TelemetryQueryCondition
type TelemetryQueryCondition struct {
	Key string `json:"key"`
	Value *string `json:"value,omitempty"`
}

// NewTelemetryQueryCondition instantiates a new TelemetryQueryCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryQueryCondition(key string) *TelemetryQueryCondition {
	this := TelemetryQueryCondition{}
	this.Key = key
	return &this
}

// NewTelemetryQueryConditionWithDefaults instantiates a new TelemetryQueryCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryQueryConditionWithDefaults() *TelemetryQueryCondition {
	this := TelemetryQueryCondition{}
	return &this
}

// GetKey returns the Key field value
func (o *TelemetryQueryCondition) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *TelemetryQueryCondition) GetKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *TelemetryQueryCondition) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TelemetryQueryCondition) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryQueryCondition) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TelemetryQueryCondition) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *TelemetryQueryCondition) SetValue(v string) {
	o.Value = &v
}

func (o TelemetryQueryCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableTelemetryQueryCondition struct {
	value *TelemetryQueryCondition
	isSet bool
}

func (v NullableTelemetryQueryCondition) Get() *TelemetryQueryCondition {
	return v.value
}

func (v *NullableTelemetryQueryCondition) Set(val *TelemetryQueryCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryQueryCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryQueryCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryQueryCondition(val *TelemetryQueryCondition) *NullableTelemetryQueryCondition {
	return &NullableTelemetryQueryCondition{value: val, isSet: true}
}

func (v NullableTelemetryQueryCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryQueryCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


