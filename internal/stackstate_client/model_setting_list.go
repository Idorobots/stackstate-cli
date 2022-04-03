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

// SettingList struct for SettingList
type SettingList struct {
	Nodes *[]Setting `json:"nodes,omitempty"`
}

// NewSettingList instantiates a new SettingList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingList() *SettingList {
	this := SettingList{}
	return &this
}

// NewSettingListWithDefaults instantiates a new SettingList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingListWithDefaults() *SettingList {
	this := SettingList{}
	return &this
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *SettingList) GetNodes() []Setting {
	if o == nil || o.Nodes == nil {
		var ret []Setting
		return ret
	}
	return *o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingList) GetNodesOk() (*[]Setting, bool) {
	if o == nil || o.Nodes == nil {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *SettingList) HasNodes() bool {
	if o != nil && o.Nodes != nil {
		return true
	}

	return false
}

// SetNodes gets a reference to the given []Setting and assigns it to the Nodes field.
func (o *SettingList) SetNodes(v []Setting) {
	o.Nodes = &v
}

func (o SettingList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Nodes != nil {
		toSerialize["nodes"] = o.Nodes
	}
	return json.Marshal(toSerialize)
}

type NullableSettingList struct {
	value *SettingList
	isSet bool
}

func (v NullableSettingList) Get() *SettingList {
	return v.value
}

func (v *NullableSettingList) Set(val *SettingList) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingList) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingList(val *SettingList) *NullableSettingList {
	return &NullableSettingList{value: val, isSet: true}
}

func (v NullableSettingList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
