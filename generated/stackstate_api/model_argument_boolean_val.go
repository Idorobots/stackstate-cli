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

// ArgumentBooleanVal struct for ArgumentBooleanVal
type ArgumentBooleanVal struct {
	Type                string `json:"_type"`
	Id                  *int64 `json:"id,omitempty"`
	LastUpdateTimestamp *int64 `json:"lastUpdateTimestamp,omitempty"`
	Parameter           int64  `json:"parameter"`
	Value               bool   `json:"value"`
}

// NewArgumentBooleanVal instantiates a new ArgumentBooleanVal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArgumentBooleanVal(type_ string, parameter int64, value bool) *ArgumentBooleanVal {
	this := ArgumentBooleanVal{}
	this.Type = type_
	this.Parameter = parameter
	this.Value = value
	return &this
}

// NewArgumentBooleanValWithDefaults instantiates a new ArgumentBooleanVal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArgumentBooleanValWithDefaults() *ArgumentBooleanVal {
	this := ArgumentBooleanVal{}
	return &this
}

// GetType returns the Type field value
func (o *ArgumentBooleanVal) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ArgumentBooleanVal) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ArgumentBooleanVal) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ArgumentBooleanVal) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArgumentBooleanVal) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ArgumentBooleanVal) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ArgumentBooleanVal) SetId(v int64) {
	o.Id = &v
}

// GetLastUpdateTimestamp returns the LastUpdateTimestamp field value if set, zero value otherwise.
func (o *ArgumentBooleanVal) GetLastUpdateTimestamp() int64 {
	if o == nil || o.LastUpdateTimestamp == nil {
		var ret int64
		return ret
	}
	return *o.LastUpdateTimestamp
}

// GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArgumentBooleanVal) GetLastUpdateTimestampOk() (*int64, bool) {
	if o == nil || o.LastUpdateTimestamp == nil {
		return nil, false
	}
	return o.LastUpdateTimestamp, true
}

// HasLastUpdateTimestamp returns a boolean if a field has been set.
func (o *ArgumentBooleanVal) HasLastUpdateTimestamp() bool {
	if o != nil && o.LastUpdateTimestamp != nil {
		return true
	}

	return false
}

// SetLastUpdateTimestamp gets a reference to the given int64 and assigns it to the LastUpdateTimestamp field.
func (o *ArgumentBooleanVal) SetLastUpdateTimestamp(v int64) {
	o.LastUpdateTimestamp = &v
}

// GetParameter returns the Parameter field value
func (o *ArgumentBooleanVal) GetParameter() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value
// and a boolean to check if the value has been set.
func (o *ArgumentBooleanVal) GetParameterOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameter, true
}

// SetParameter sets field value
func (o *ArgumentBooleanVal) SetParameter(v int64) {
	o.Parameter = v
}

// GetValue returns the Value field value
func (o *ArgumentBooleanVal) GetValue() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ArgumentBooleanVal) GetValueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ArgumentBooleanVal) SetValue(v bool) {
	o.Value = v
}

func (o ArgumentBooleanVal) MarshalJSON() ([]byte, error) {
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
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableArgumentBooleanVal struct {
	value *ArgumentBooleanVal
	isSet bool
}

func (v NullableArgumentBooleanVal) Get() *ArgumentBooleanVal {
	return v.value
}

func (v *NullableArgumentBooleanVal) Set(val *ArgumentBooleanVal) {
	v.value = val
	v.isSet = true
}

func (v NullableArgumentBooleanVal) IsSet() bool {
	return v.isSet
}

func (v *NullableArgumentBooleanVal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArgumentBooleanVal(val *ArgumentBooleanVal) *NullableArgumentBooleanVal {
	return &NullableArgumentBooleanVal{value: val, isSet: true}
}

func (v NullableArgumentBooleanVal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArgumentBooleanVal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
