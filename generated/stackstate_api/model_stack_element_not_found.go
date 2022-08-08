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

// StackElementNotFound struct for StackElementNotFound
type StackElementNotFound struct {
	Type string `json:"_type"`
	ObjectType string `json:"objectType"`
	ObjectId string `json:"objectId"`
	Message string `json:"message"`
}

// NewStackElementNotFound instantiates a new StackElementNotFound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStackElementNotFound(type_ string, objectType string, objectId string, message string) *StackElementNotFound {
	this := StackElementNotFound{}
	this.Type = type_
	this.ObjectType = objectType
	this.ObjectId = objectId
	this.Message = message
	return &this
}

// NewStackElementNotFoundWithDefaults instantiates a new StackElementNotFound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStackElementNotFoundWithDefaults() *StackElementNotFound {
	this := StackElementNotFound{}
	return &this
}

// GetType returns the Type field value
func (o *StackElementNotFound) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StackElementNotFound) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StackElementNotFound) SetType(v string) {
	o.Type = v
}

// GetObjectType returns the ObjectType field value
func (o *StackElementNotFound) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StackElementNotFound) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StackElementNotFound) SetObjectType(v string) {
	o.ObjectType = v
}

// GetObjectId returns the ObjectId field value
func (o *StackElementNotFound) GetObjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value
// and a boolean to check if the value has been set.
func (o *StackElementNotFound) GetObjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectId, true
}

// SetObjectId sets field value
func (o *StackElementNotFound) SetObjectId(v string) {
	o.ObjectId = v
}

// GetMessage returns the Message field value
func (o *StackElementNotFound) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *StackElementNotFound) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *StackElementNotFound) SetMessage(v string) {
	o.Message = v
}

func (o StackElementNotFound) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if true {
		toSerialize["objectType"] = o.ObjectType
	}
	if true {
		toSerialize["objectId"] = o.ObjectId
	}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableStackElementNotFound struct {
	value *StackElementNotFound
	isSet bool
}

func (v NullableStackElementNotFound) Get() *StackElementNotFound {
	return v.value
}

func (v *NullableStackElementNotFound) Set(val *StackElementNotFound) {
	v.value = val
	v.isSet = true
}

func (v NullableStackElementNotFound) IsSet() bool {
	return v.isSet
}

func (v *NullableStackElementNotFound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStackElementNotFound(val *StackElementNotFound) *NullableStackElementNotFound {
	return &NullableStackElementNotFound{value: val, isSet: true}
}

func (v NullableStackElementNotFound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStackElementNotFound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


