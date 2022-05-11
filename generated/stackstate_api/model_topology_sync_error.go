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

// TopologySyncError struct for TopologySyncError
type TopologySyncError struct {
	Message string `json:"message"`
}

// NewTopologySyncError instantiates a new TopologySyncError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopologySyncError(message string) *TopologySyncError {
	this := TopologySyncError{}
	this.Message = message
	return &this
}

// NewTopologySyncErrorWithDefaults instantiates a new TopologySyncError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopologySyncErrorWithDefaults() *TopologySyncError {
	this := TopologySyncError{}
	return &this
}

// GetMessage returns the Message field value
func (o *TopologySyncError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *TopologySyncError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *TopologySyncError) SetMessage(v string) {
	o.Message = v
}

func (o TopologySyncError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableTopologySyncError struct {
	value *TopologySyncError
	isSet bool
}

func (v NullableTopologySyncError) Get() *TopologySyncError {
	return v.value
}

func (v *NullableTopologySyncError) Set(val *TopologySyncError) {
	v.value = val
	v.isSet = true
}

func (v NullableTopologySyncError) IsSet() bool {
	return v.isSet
}

func (v *NullableTopologySyncError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopologySyncError(val *TopologySyncError) *NullableTopologySyncError {
	return &NullableTopologySyncError{value: val, isSet: true}
}

func (v NullableTopologySyncError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopologySyncError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


