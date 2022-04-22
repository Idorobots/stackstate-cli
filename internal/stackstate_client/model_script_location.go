/*
StackState API

StackState's API specification

API version: 0.0.1
Contact: info@stackstate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackstate_client

import (
	"encoding/json"
)

// ScriptLocation struct for ScriptLocation
type ScriptLocation struct {
	Line int32 `json:"line"`
	Column int32 `json:"column"`
}

// NewScriptLocation instantiates a new ScriptLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScriptLocation(line int32, column int32) *ScriptLocation {
	this := ScriptLocation{}
	this.Line = line
	this.Column = column
	return &this
}

// NewScriptLocationWithDefaults instantiates a new ScriptLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScriptLocationWithDefaults() *ScriptLocation {
	this := ScriptLocation{}
	return &this
}

// GetLine returns the Line field value
func (o *ScriptLocation) GetLine() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Line
}

// GetLineOk returns a tuple with the Line field value
// and a boolean to check if the value has been set.
func (o *ScriptLocation) GetLineOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Line, true
}

// SetLine sets field value
func (o *ScriptLocation) SetLine(v int32) {
	o.Line = v
}

// GetColumn returns the Column field value
func (o *ScriptLocation) GetColumn() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Column
}

// GetColumnOk returns a tuple with the Column field value
// and a boolean to check if the value has been set.
func (o *ScriptLocation) GetColumnOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Column, true
}

// SetColumn sets field value
func (o *ScriptLocation) SetColumn(v int32) {
	o.Column = v
}

func (o ScriptLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["line"] = o.Line
	}
	if true {
		toSerialize["column"] = o.Column
	}
	return json.Marshal(toSerialize)
}

type NullableScriptLocation struct {
	value *ScriptLocation
	isSet bool
}

func (v NullableScriptLocation) Get() *ScriptLocation {
	return v.value
}

func (v *NullableScriptLocation) Set(val *ScriptLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableScriptLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableScriptLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScriptLocation(val *ScriptLocation) *NullableScriptLocation {
	return &NullableScriptLocation{value: val, isSet: true}
}

func (v NullableScriptLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScriptLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


