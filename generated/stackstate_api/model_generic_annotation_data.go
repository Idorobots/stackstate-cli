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

// GenericAnnotationData struct for GenericAnnotationData
type GenericAnnotationData struct {
	Type string `json:"_type"`
	Annotation map[string]interface{} `json:"annotation"`
}

// NewGenericAnnotationData instantiates a new GenericAnnotationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericAnnotationData(type_ string, annotation map[string]interface{}) *GenericAnnotationData {
	this := GenericAnnotationData{}
	this.Type = type_
	this.Annotation = annotation
	return &this
}

// NewGenericAnnotationDataWithDefaults instantiates a new GenericAnnotationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericAnnotationDataWithDefaults() *GenericAnnotationData {
	this := GenericAnnotationData{}
	return &this
}

// GetType returns the Type field value
func (o *GenericAnnotationData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GenericAnnotationData) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GenericAnnotationData) SetType(v string) {
	o.Type = v
}

// GetAnnotation returns the Annotation field value
func (o *GenericAnnotationData) GetAnnotation() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Annotation
}

// GetAnnotationOk returns a tuple with the Annotation field value
// and a boolean to check if the value has been set.
func (o *GenericAnnotationData) GetAnnotationOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Annotation, true
}

// SetAnnotation sets field value
func (o *GenericAnnotationData) SetAnnotation(v map[string]interface{}) {
	o.Annotation = v
}

func (o GenericAnnotationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if true {
		toSerialize["annotation"] = o.Annotation
	}
	return json.Marshal(toSerialize)
}

type NullableGenericAnnotationData struct {
	value *GenericAnnotationData
	isSet bool
}

func (v NullableGenericAnnotationData) Get() *GenericAnnotationData {
	return v.value
}

func (v *NullableGenericAnnotationData) Set(val *GenericAnnotationData) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericAnnotationData) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericAnnotationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericAnnotationData(val *GenericAnnotationData) *NullableGenericAnnotationData {
	return &NullableGenericAnnotationData{value: val, isSet: true}
}

func (v NullableGenericAnnotationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericAnnotationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


