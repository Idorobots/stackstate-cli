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

// ArgumentEventStreamRef struct for ArgumentEventStreamRef
type ArgumentEventStreamRef struct {
	Type                string `json:"_type"`
	Id                  *int64 `json:"id,omitempty"`
	LastUpdateTimestamp *int64 `json:"lastUpdateTimestamp,omitempty"`
	Parameter           int64  `json:"parameter"`
	Stream              int64  `json:"stream"`
}

// NewArgumentEventStreamRef instantiates a new ArgumentEventStreamRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArgumentEventStreamRef(type_ string, parameter int64, stream int64) *ArgumentEventStreamRef {
	this := ArgumentEventStreamRef{}
	this.Type = type_
	this.Parameter = parameter
	this.Stream = stream
	return &this
}

// NewArgumentEventStreamRefWithDefaults instantiates a new ArgumentEventStreamRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArgumentEventStreamRefWithDefaults() *ArgumentEventStreamRef {
	this := ArgumentEventStreamRef{}
	return &this
}

// GetType returns the Type field value
func (o *ArgumentEventStreamRef) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ArgumentEventStreamRef) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ArgumentEventStreamRef) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ArgumentEventStreamRef) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArgumentEventStreamRef) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ArgumentEventStreamRef) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ArgumentEventStreamRef) SetId(v int64) {
	o.Id = &v
}

// GetLastUpdateTimestamp returns the LastUpdateTimestamp field value if set, zero value otherwise.
func (o *ArgumentEventStreamRef) GetLastUpdateTimestamp() int64 {
	if o == nil || o.LastUpdateTimestamp == nil {
		var ret int64
		return ret
	}
	return *o.LastUpdateTimestamp
}

// GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArgumentEventStreamRef) GetLastUpdateTimestampOk() (*int64, bool) {
	if o == nil || o.LastUpdateTimestamp == nil {
		return nil, false
	}
	return o.LastUpdateTimestamp, true
}

// HasLastUpdateTimestamp returns a boolean if a field has been set.
func (o *ArgumentEventStreamRef) HasLastUpdateTimestamp() bool {
	if o != nil && o.LastUpdateTimestamp != nil {
		return true
	}

	return false
}

// SetLastUpdateTimestamp gets a reference to the given int64 and assigns it to the LastUpdateTimestamp field.
func (o *ArgumentEventStreamRef) SetLastUpdateTimestamp(v int64) {
	o.LastUpdateTimestamp = &v
}

// GetParameter returns the Parameter field value
func (o *ArgumentEventStreamRef) GetParameter() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value
// and a boolean to check if the value has been set.
func (o *ArgumentEventStreamRef) GetParameterOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameter, true
}

// SetParameter sets field value
func (o *ArgumentEventStreamRef) SetParameter(v int64) {
	o.Parameter = v
}

// GetStream returns the Stream field value
func (o *ArgumentEventStreamRef) GetStream() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Stream
}

// GetStreamOk returns a tuple with the Stream field value
// and a boolean to check if the value has been set.
func (o *ArgumentEventStreamRef) GetStreamOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stream, true
}

// SetStream sets field value
func (o *ArgumentEventStreamRef) SetStream(v int64) {
	o.Stream = v
}

func (o ArgumentEventStreamRef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdateTimestamp != nil {
		toSerialize["lastUpdateTimestamp"] = o.LastUpdateTimestamp
	}
	if true {
		toSerialize["parameter"] = o.Parameter
	}
	if true {
		toSerialize["stream"] = o.Stream
	}
	return json.Marshal(toSerialize)
}

type NullableArgumentEventStreamRef struct {
	value *ArgumentEventStreamRef
	isSet bool
}

func (v NullableArgumentEventStreamRef) Get() *ArgumentEventStreamRef {
	return v.value
}

func (v *NullableArgumentEventStreamRef) Set(val *ArgumentEventStreamRef) {
	v.value = val
	v.isSet = true
}

func (v NullableArgumentEventStreamRef) IsSet() bool {
	return v.isSet
}

func (v *NullableArgumentEventStreamRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArgumentEventStreamRef(val *ArgumentEventStreamRef) *NullableArgumentEventStreamRef {
	return &NullableArgumentEventStreamRef{value: val, isSet: true}
}

func (v NullableArgumentEventStreamRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArgumentEventStreamRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
