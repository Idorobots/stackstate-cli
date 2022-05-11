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

// EventItemsWithTotal struct for EventItemsWithTotal
type EventItemsWithTotal struct {
	Items []TopologyEvent `json:"items"`
	Total int64 `json:"total"`
}

// NewEventItemsWithTotal instantiates a new EventItemsWithTotal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventItemsWithTotal(items []TopologyEvent, total int64) *EventItemsWithTotal {
	this := EventItemsWithTotal{}
	this.Items = items
	this.Total = total
	return &this
}

// NewEventItemsWithTotalWithDefaults instantiates a new EventItemsWithTotal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventItemsWithTotalWithDefaults() *EventItemsWithTotal {
	this := EventItemsWithTotal{}
	return &this
}

// GetItems returns the Items field value
func (o *EventItemsWithTotal) GetItems() []TopologyEvent {
	if o == nil {
		var ret []TopologyEvent
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *EventItemsWithTotal) GetItemsOk() ([]TopologyEvent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *EventItemsWithTotal) SetItems(v []TopologyEvent) {
	o.Items = v
}

// GetTotal returns the Total field value
func (o *EventItemsWithTotal) GetTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *EventItemsWithTotal) GetTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *EventItemsWithTotal) SetTotal(v int64) {
	o.Total = v
}

func (o EventItemsWithTotal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	if true {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableEventItemsWithTotal struct {
	value *EventItemsWithTotal
	isSet bool
}

func (v NullableEventItemsWithTotal) Get() *EventItemsWithTotal {
	return v.value
}

func (v *NullableEventItemsWithTotal) Set(val *EventItemsWithTotal) {
	v.value = val
	v.isSet = true
}

func (v NullableEventItemsWithTotal) IsSet() bool {
	return v.isSet
}

func (v *NullableEventItemsWithTotal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventItemsWithTotal(val *EventItemsWithTotal) *NullableEventItemsWithTotal {
	return &NullableEventItemsWithTotal{value: val, isSet: true}
}

func (v NullableEventItemsWithTotal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventItemsWithTotal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


