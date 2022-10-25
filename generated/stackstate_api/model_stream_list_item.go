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

// StreamListItem struct for StreamListItem
type StreamListItem struct {
	Urn              string `json:"urn"`
	ConsistencyModel string `json:"consistencyModel"`
	SubStreams       int32  `json:"subStreams"`
}

// NewStreamListItem instantiates a new StreamListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamListItem(urn string, consistencyModel string, subStreams int32) *StreamListItem {
	this := StreamListItem{}
	this.Urn = urn
	this.ConsistencyModel = consistencyModel
	this.SubStreams = subStreams
	return &this
}

// NewStreamListItemWithDefaults instantiates a new StreamListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamListItemWithDefaults() *StreamListItem {
	this := StreamListItem{}
	return &this
}

// GetUrn returns the Urn field value
func (o *StreamListItem) GetUrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Urn
}

// GetUrnOk returns a tuple with the Urn field value
// and a boolean to check if the value has been set.
func (o *StreamListItem) GetUrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Urn, true
}

// SetUrn sets field value
func (o *StreamListItem) SetUrn(v string) {
	o.Urn = v
}

// GetConsistencyModel returns the ConsistencyModel field value
func (o *StreamListItem) GetConsistencyModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsistencyModel
}

// GetConsistencyModelOk returns a tuple with the ConsistencyModel field value
// and a boolean to check if the value has been set.
func (o *StreamListItem) GetConsistencyModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsistencyModel, true
}

// SetConsistencyModel sets field value
func (o *StreamListItem) SetConsistencyModel(v string) {
	o.ConsistencyModel = v
}

// GetSubStreams returns the SubStreams field value
func (o *StreamListItem) GetSubStreams() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SubStreams
}

// GetSubStreamsOk returns a tuple with the SubStreams field value
// and a boolean to check if the value has been set.
func (o *StreamListItem) GetSubStreamsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubStreams, true
}

// SetSubStreams sets field value
func (o *StreamListItem) SetSubStreams(v int32) {
	o.SubStreams = v
}

func (o StreamListItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["urn"] = o.Urn
	}
	if true {
		toSerialize["consistencyModel"] = o.ConsistencyModel
	}
	if true {
		toSerialize["subStreams"] = o.SubStreams
	}
	return json.Marshal(toSerialize)
}

type NullableStreamListItem struct {
	value *StreamListItem
	isSet bool
}

func (v NullableStreamListItem) Get() *StreamListItem {
	return v.value
}

func (v *NullableStreamListItem) Set(val *StreamListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamListItem(val *StreamListItem) *NullableStreamListItem {
	return &NullableStreamListItem{value: val, isSet: true}
}

func (v NullableStreamListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
