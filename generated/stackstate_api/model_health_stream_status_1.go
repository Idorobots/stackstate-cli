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

// HealthStreamStatus1 struct for HealthStreamStatus1
type HealthStreamStatus1 struct {
	Partition int32 `json:"partition"`
	ConsistencyModel string `json:"consistencyModel"`
	RecoverMessage *string `json:"recoverMessage,omitempty"`
	GlobalErrors []HealthStreamError `json:"globalErrors,omitempty"`
	AggregateMetrics HealthStreamMetrics1 `json:"aggregateMetrics"`
	MainStreamStatus *HealthSubStreamStatus1 `json:"mainStreamStatus,omitempty"`
}

// NewHealthStreamStatus1 instantiates a new HealthStreamStatus1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthStreamStatus1(partition int32, consistencyModel string, aggregateMetrics HealthStreamMetrics1) *HealthStreamStatus1 {
	this := HealthStreamStatus1{}
	this.Partition = partition
	this.ConsistencyModel = consistencyModel
	this.AggregateMetrics = aggregateMetrics
	return &this
}

// NewHealthStreamStatus1WithDefaults instantiates a new HealthStreamStatus1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthStreamStatus1WithDefaults() *HealthStreamStatus1 {
	this := HealthStreamStatus1{}
	return &this
}

// GetPartition returns the Partition field value
func (o *HealthStreamStatus1) GetPartition() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value
// and a boolean to check if the value has been set.
func (o *HealthStreamStatus1) GetPartitionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Partition, true
}

// SetPartition sets field value
func (o *HealthStreamStatus1) SetPartition(v int32) {
	o.Partition = v
}

// GetConsistencyModel returns the ConsistencyModel field value
func (o *HealthStreamStatus1) GetConsistencyModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsistencyModel
}

// GetConsistencyModelOk returns a tuple with the ConsistencyModel field value
// and a boolean to check if the value has been set.
func (o *HealthStreamStatus1) GetConsistencyModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsistencyModel, true
}

// SetConsistencyModel sets field value
func (o *HealthStreamStatus1) SetConsistencyModel(v string) {
	o.ConsistencyModel = v
}

// GetRecoverMessage returns the RecoverMessage field value if set, zero value otherwise.
func (o *HealthStreamStatus1) GetRecoverMessage() string {
	if o == nil || o.RecoverMessage == nil {
		var ret string
		return ret
	}
	return *o.RecoverMessage
}

// GetRecoverMessageOk returns a tuple with the RecoverMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthStreamStatus1) GetRecoverMessageOk() (*string, bool) {
	if o == nil || o.RecoverMessage == nil {
		return nil, false
	}
	return o.RecoverMessage, true
}

// HasRecoverMessage returns a boolean if a field has been set.
func (o *HealthStreamStatus1) HasRecoverMessage() bool {
	if o != nil && o.RecoverMessage != nil {
		return true
	}

	return false
}

// SetRecoverMessage gets a reference to the given string and assigns it to the RecoverMessage field.
func (o *HealthStreamStatus1) SetRecoverMessage(v string) {
	o.RecoverMessage = &v
}

// GetGlobalErrors returns the GlobalErrors field value if set, zero value otherwise.
func (o *HealthStreamStatus1) GetGlobalErrors() []HealthStreamError {
	if o == nil || o.GlobalErrors == nil {
		var ret []HealthStreamError
		return ret
	}
	return o.GlobalErrors
}

// GetGlobalErrorsOk returns a tuple with the GlobalErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthStreamStatus1) GetGlobalErrorsOk() ([]HealthStreamError, bool) {
	if o == nil || o.GlobalErrors == nil {
		return nil, false
	}
	return o.GlobalErrors, true
}

// HasGlobalErrors returns a boolean if a field has been set.
func (o *HealthStreamStatus1) HasGlobalErrors() bool {
	if o != nil && o.GlobalErrors != nil {
		return true
	}

	return false
}

// SetGlobalErrors gets a reference to the given []HealthStreamError and assigns it to the GlobalErrors field.
func (o *HealthStreamStatus1) SetGlobalErrors(v []HealthStreamError) {
	o.GlobalErrors = v
}

// GetAggregateMetrics returns the AggregateMetrics field value
func (o *HealthStreamStatus1) GetAggregateMetrics() HealthStreamMetrics1 {
	if o == nil {
		var ret HealthStreamMetrics1
		return ret
	}

	return o.AggregateMetrics
}

// GetAggregateMetricsOk returns a tuple with the AggregateMetrics field value
// and a boolean to check if the value has been set.
func (o *HealthStreamStatus1) GetAggregateMetricsOk() (*HealthStreamMetrics1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AggregateMetrics, true
}

// SetAggregateMetrics sets field value
func (o *HealthStreamStatus1) SetAggregateMetrics(v HealthStreamMetrics1) {
	o.AggregateMetrics = v
}

// GetMainStreamStatus returns the MainStreamStatus field value if set, zero value otherwise.
func (o *HealthStreamStatus1) GetMainStreamStatus() HealthSubStreamStatus1 {
	if o == nil || o.MainStreamStatus == nil {
		var ret HealthSubStreamStatus1
		return ret
	}
	return *o.MainStreamStatus
}

// GetMainStreamStatusOk returns a tuple with the MainStreamStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthStreamStatus1) GetMainStreamStatusOk() (*HealthSubStreamStatus1, bool) {
	if o == nil || o.MainStreamStatus == nil {
		return nil, false
	}
	return o.MainStreamStatus, true
}

// HasMainStreamStatus returns a boolean if a field has been set.
func (o *HealthStreamStatus1) HasMainStreamStatus() bool {
	if o != nil && o.MainStreamStatus != nil {
		return true
	}

	return false
}

// SetMainStreamStatus gets a reference to the given HealthSubStreamStatus1 and assigns it to the MainStreamStatus field.
func (o *HealthStreamStatus1) SetMainStreamStatus(v HealthSubStreamStatus1) {
	o.MainStreamStatus = &v
}

func (o HealthStreamStatus1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["partition"] = o.Partition
	}
	if true {
		toSerialize["consistencyModel"] = o.ConsistencyModel
	}
	if o.RecoverMessage != nil {
		toSerialize["recoverMessage"] = o.RecoverMessage
	}
	if o.GlobalErrors != nil {
		toSerialize["globalErrors"] = o.GlobalErrors
	}
	if true {
		toSerialize["aggregateMetrics"] = o.AggregateMetrics
	}
	if o.MainStreamStatus != nil {
		toSerialize["mainStreamStatus"] = o.MainStreamStatus
	}
	return json.Marshal(toSerialize)
}

type NullableHealthStreamStatus1 struct {
	value *HealthStreamStatus1
	isSet bool
}

func (v NullableHealthStreamStatus1) Get() *HealthStreamStatus1 {
	return v.value
}

func (v *NullableHealthStreamStatus1) Set(val *HealthStreamStatus1) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthStreamStatus1) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthStreamStatus1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthStreamStatus1(val *HealthStreamStatus1) *NullableHealthStreamStatus1 {
	return &NullableHealthStreamStatus1{value: val, isSet: true}
}

func (v NullableHealthStreamStatus1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthStreamStatus1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


