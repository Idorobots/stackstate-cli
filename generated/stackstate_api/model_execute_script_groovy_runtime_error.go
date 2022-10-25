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

// ExecuteScriptGroovyRuntimeError struct for ExecuteScriptGroovyRuntimeError
type ExecuteScriptGroovyRuntimeError struct {
	Type     string          `json:"_type"`
	Message  string          `json:"message"`
	Location *ScriptLocation `json:"location,omitempty"`
}

// NewExecuteScriptGroovyRuntimeError instantiates a new ExecuteScriptGroovyRuntimeError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecuteScriptGroovyRuntimeError(type_ string, message string) *ExecuteScriptGroovyRuntimeError {
	this := ExecuteScriptGroovyRuntimeError{}
	this.Type = type_
	this.Message = message
	return &this
}

// NewExecuteScriptGroovyRuntimeErrorWithDefaults instantiates a new ExecuteScriptGroovyRuntimeError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecuteScriptGroovyRuntimeErrorWithDefaults() *ExecuteScriptGroovyRuntimeError {
	this := ExecuteScriptGroovyRuntimeError{}
	return &this
}

// GetType returns the Type field value
func (o *ExecuteScriptGroovyRuntimeError) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExecuteScriptGroovyRuntimeError) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExecuteScriptGroovyRuntimeError) SetType(v string) {
	o.Type = v
}

// GetMessage returns the Message field value
func (o *ExecuteScriptGroovyRuntimeError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ExecuteScriptGroovyRuntimeError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ExecuteScriptGroovyRuntimeError) SetMessage(v string) {
	o.Message = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ExecuteScriptGroovyRuntimeError) GetLocation() ScriptLocation {
	if o == nil || o.Location == nil {
		var ret ScriptLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecuteScriptGroovyRuntimeError) GetLocationOk() (*ScriptLocation, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ExecuteScriptGroovyRuntimeError) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given ScriptLocation and assigns it to the Location field.
func (o *ExecuteScriptGroovyRuntimeError) SetLocation(v ScriptLocation) {
	o.Location = &v
}

func (o ExecuteScriptGroovyRuntimeError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	return json.Marshal(toSerialize)
}

type NullableExecuteScriptGroovyRuntimeError struct {
	value *ExecuteScriptGroovyRuntimeError
	isSet bool
}

func (v NullableExecuteScriptGroovyRuntimeError) Get() *ExecuteScriptGroovyRuntimeError {
	return v.value
}

func (v *NullableExecuteScriptGroovyRuntimeError) Set(val *ExecuteScriptGroovyRuntimeError) {
	v.value = val
	v.isSet = true
}

func (v NullableExecuteScriptGroovyRuntimeError) IsSet() bool {
	return v.isSet
}

func (v *NullableExecuteScriptGroovyRuntimeError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecuteScriptGroovyRuntimeError(val *ExecuteScriptGroovyRuntimeError) *NullableExecuteScriptGroovyRuntimeError {
	return &NullableExecuteScriptGroovyRuntimeError{value: val, isSet: true}
}

func (v NullableExecuteScriptGroovyRuntimeError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecuteScriptGroovyRuntimeError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
