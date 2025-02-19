/*
StackState API

This API documentation page describes the StackState server API. The StackState UI and CLI use the StackState server API to configure and query StackState.  You can use the API for similar purposes.  Each request sent to the StackState server API must be authenticated using one of the available authentication methods.   *Note that the StackState receiver API, used to send topology, telemetry and traces to StackState, is not described on this page and requires a different authentication method.*  For more information on StackState, refer to the [StackState documentation](https://docs.stackstate.com).

API version: 5.2.0
Contact: info@stackstate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackstate_api

import (
	"encoding/json"
)

// CheckStateAcknowledgement struct for CheckStateAcknowledgement
type CheckStateAcknowledgement struct {
	Type                  string `json:"_type"`
	Id                    *int64 `json:"id,omitempty"`
	LastUpdateTimeStamp   *int64 `json:"lastUpdateTimeStamp,omitempty"`
	AcknowledgedTimestamp int64  `json:"acknowledgedTimestamp"`
	Message               string `json:"message"`
}

// NewCheckStateAcknowledgement instantiates a new CheckStateAcknowledgement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckStateAcknowledgement(type_ string, acknowledgedTimestamp int64, message string) *CheckStateAcknowledgement {
	this := CheckStateAcknowledgement{}
	this.Type = type_
	this.AcknowledgedTimestamp = acknowledgedTimestamp
	this.Message = message
	return &this
}

// NewCheckStateAcknowledgementWithDefaults instantiates a new CheckStateAcknowledgement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckStateAcknowledgementWithDefaults() *CheckStateAcknowledgement {
	this := CheckStateAcknowledgement{}
	return &this
}

// GetType returns the Type field value
func (o *CheckStateAcknowledgement) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CheckStateAcknowledgement) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CheckStateAcknowledgement) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CheckStateAcknowledgement) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckStateAcknowledgement) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CheckStateAcknowledgement) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *CheckStateAcknowledgement) SetId(v int64) {
	o.Id = &v
}

// GetLastUpdateTimeStamp returns the LastUpdateTimeStamp field value if set, zero value otherwise.
func (o *CheckStateAcknowledgement) GetLastUpdateTimeStamp() int64 {
	if o == nil || o.LastUpdateTimeStamp == nil {
		var ret int64
		return ret
	}
	return *o.LastUpdateTimeStamp
}

// GetLastUpdateTimeStampOk returns a tuple with the LastUpdateTimeStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckStateAcknowledgement) GetLastUpdateTimeStampOk() (*int64, bool) {
	if o == nil || o.LastUpdateTimeStamp == nil {
		return nil, false
	}
	return o.LastUpdateTimeStamp, true
}

// HasLastUpdateTimeStamp returns a boolean if a field has been set.
func (o *CheckStateAcknowledgement) HasLastUpdateTimeStamp() bool {
	if o != nil && o.LastUpdateTimeStamp != nil {
		return true
	}

	return false
}

// SetLastUpdateTimeStamp gets a reference to the given int64 and assigns it to the LastUpdateTimeStamp field.
func (o *CheckStateAcknowledgement) SetLastUpdateTimeStamp(v int64) {
	o.LastUpdateTimeStamp = &v
}

// GetAcknowledgedTimestamp returns the AcknowledgedTimestamp field value
func (o *CheckStateAcknowledgement) GetAcknowledgedTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AcknowledgedTimestamp
}

// GetAcknowledgedTimestampOk returns a tuple with the AcknowledgedTimestamp field value
// and a boolean to check if the value has been set.
func (o *CheckStateAcknowledgement) GetAcknowledgedTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcknowledgedTimestamp, true
}

// SetAcknowledgedTimestamp sets field value
func (o *CheckStateAcknowledgement) SetAcknowledgedTimestamp(v int64) {
	o.AcknowledgedTimestamp = v
}

// GetMessage returns the Message field value
func (o *CheckStateAcknowledgement) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *CheckStateAcknowledgement) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *CheckStateAcknowledgement) SetMessage(v string) {
	o.Message = v
}

func (o CheckStateAcknowledgement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdateTimeStamp != nil {
		toSerialize["lastUpdateTimeStamp"] = o.LastUpdateTimeStamp
	}
	if true {
		toSerialize["acknowledgedTimestamp"] = o.AcknowledgedTimestamp
	}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableCheckStateAcknowledgement struct {
	value *CheckStateAcknowledgement
	isSet bool
}

func (v NullableCheckStateAcknowledgement) Get() *CheckStateAcknowledgement {
	return v.value
}

func (v *NullableCheckStateAcknowledgement) Set(val *CheckStateAcknowledgement) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckStateAcknowledgement) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckStateAcknowledgement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckStateAcknowledgement(val *CheckStateAcknowledgement) *NullableCheckStateAcknowledgement {
	return &NullableCheckStateAcknowledgement{value: val, isSet: true}
}

func (v NullableCheckStateAcknowledgement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckStateAcknowledgement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
