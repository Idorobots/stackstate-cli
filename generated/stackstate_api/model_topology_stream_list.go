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

// TopologyStreamList struct for TopologyStreamList
type TopologyStreamList struct {
	Streams []TopologyStreamListItem `json:"streams"`
}

// NewTopologyStreamList instantiates a new TopologyStreamList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopologyStreamList(streams []TopologyStreamListItem) *TopologyStreamList {
	this := TopologyStreamList{}
	this.Streams = streams
	return &this
}

// NewTopologyStreamListWithDefaults instantiates a new TopologyStreamList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopologyStreamListWithDefaults() *TopologyStreamList {
	this := TopologyStreamList{}
	return &this
}

// GetStreams returns the Streams field value
func (o *TopologyStreamList) GetStreams() []TopologyStreamListItem {
	if o == nil {
		var ret []TopologyStreamListItem
		return ret
	}

	return o.Streams
}

// GetStreamsOk returns a tuple with the Streams field value
// and a boolean to check if the value has been set.
func (o *TopologyStreamList) GetStreamsOk() ([]TopologyStreamListItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Streams, true
}

// SetStreams sets field value
func (o *TopologyStreamList) SetStreams(v []TopologyStreamListItem) {
	o.Streams = v
}

func (o TopologyStreamList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["streams"] = o.Streams
	}
	return json.Marshal(toSerialize)
}

type NullableTopologyStreamList struct {
	value *TopologyStreamList
	isSet bool
}

func (v NullableTopologyStreamList) Get() *TopologyStreamList {
	return v.value
}

func (v *NullableTopologyStreamList) Set(val *TopologyStreamList) {
	v.value = val
	v.isSet = true
}

func (v NullableTopologyStreamList) IsSet() bool {
	return v.isSet
}

func (v *NullableTopologyStreamList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopologyStreamList(val *TopologyStreamList) *NullableTopologyStreamList {
	return &NullableTopologyStreamList{value: val, isSet: true}
}

func (v NullableTopologyStreamList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopologyStreamList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


