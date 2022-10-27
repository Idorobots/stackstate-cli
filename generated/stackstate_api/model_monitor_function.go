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

// MonitorFunction struct for MonitorFunction
type MonitorFunction struct {
	Id                  int64   `json:"id"`
	Name                string  `json:"name"`
	Identifier          *string `json:"identifier,omitempty"`
	Description         *string `json:"description,omitempty"`
	LastUpdateTimestamp int64   `json:"lastUpdateTimestamp"`
}

// NewMonitorFunction instantiates a new MonitorFunction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorFunction(id int64, name string, lastUpdateTimestamp int64) *MonitorFunction {
	this := MonitorFunction{}
	this.Id = id
	this.Name = name
	this.LastUpdateTimestamp = lastUpdateTimestamp
	return &this
}

// NewMonitorFunctionWithDefaults instantiates a new MonitorFunction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorFunctionWithDefaults() *MonitorFunction {
	this := MonitorFunction{}
	return &this
}

// GetId returns the Id field value
func (o *MonitorFunction) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MonitorFunction) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MonitorFunction) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *MonitorFunction) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MonitorFunction) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MonitorFunction) SetName(v string) {
	o.Name = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *MonitorFunction) GetIdentifier() string {
	if o == nil || o.Identifier == nil {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFunction) GetIdentifierOk() (*string, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *MonitorFunction) HasIdentifier() bool {
	if o != nil && o.Identifier != nil {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *MonitorFunction) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MonitorFunction) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFunction) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MonitorFunction) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MonitorFunction) SetDescription(v string) {
	o.Description = &v
}

// GetLastUpdateTimestamp returns the LastUpdateTimestamp field value
func (o *MonitorFunction) GetLastUpdateTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.LastUpdateTimestamp
}

// GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field value
// and a boolean to check if the value has been set.
func (o *MonitorFunction) GetLastUpdateTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdateTimestamp, true
}

// SetLastUpdateTimestamp sets field value
func (o *MonitorFunction) SetLastUpdateTimestamp(v int64) {
	o.LastUpdateTimestamp = v
}

func (o MonitorFunction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Identifier != nil {
		toSerialize["identifier"] = o.Identifier
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["lastUpdateTimestamp"] = o.LastUpdateTimestamp
	}
	return json.Marshal(toSerialize)
}

type NullableMonitorFunction struct {
	value *MonitorFunction
	isSet bool
}

func (v NullableMonitorFunction) Get() *MonitorFunction {
	return v.value
}

func (v *NullableMonitorFunction) Set(val *MonitorFunction) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorFunction) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorFunction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorFunction(val *MonitorFunction) *NullableMonitorFunction {
	return &NullableMonitorFunction{value: val, isSet: true}
}

func (v NullableMonitorFunction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorFunction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
