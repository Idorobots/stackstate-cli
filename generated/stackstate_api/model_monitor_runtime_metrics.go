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

// MonitorRuntimeMetrics struct for MonitorRuntimeMetrics
type MonitorRuntimeMetrics struct {
	// Representing the count of the monitors results that are UNKNOWN and are mapped to topology.
	UnknownCount *int32 `json:"unknownCount,omitempty"`
	// Representing the count of the monitors results that are CLEAR and are mapped to topology.
	ClearCount *int32 `json:"clearCount,omitempty"`
	// Representing the count of the monitors results that are DEVIATING and are mapped to topology.
	DeviatingCount *int32 `json:"deviatingCount,omitempty"`
	// Representing the count of the monitors results that are CRITICAL and are mapped to topology.
	CriticalCount *int32 `json:"criticalCount,omitempty"`
	// Representing the epoch millis of the last monitor run.
	LastRunTimestamp *int64 `json:"lastRunTimestamp,omitempty"`
	// Representing the epoch millis of the last monitor successful run.
	LastSuccessfulRunTimestamp *int64 `json:"lastSuccessfulRunTimestamp,omitempty"`
	// Representing the epoch millis of the last monitor failed run.
	LastFailedRunTimestamp *int64 `json:"lastFailedRunTimestamp,omitempty"`
}

// NewMonitorRuntimeMetrics instantiates a new MonitorRuntimeMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorRuntimeMetrics() *MonitorRuntimeMetrics {
	this := MonitorRuntimeMetrics{}
	return &this
}

// NewMonitorRuntimeMetricsWithDefaults instantiates a new MonitorRuntimeMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorRuntimeMetricsWithDefaults() *MonitorRuntimeMetrics {
	this := MonitorRuntimeMetrics{}
	return &this
}

// GetUnknownCount returns the UnknownCount field value if set, zero value otherwise.
func (o *MonitorRuntimeMetrics) GetUnknownCount() int32 {
	if o == nil || o.UnknownCount == nil {
		var ret int32
		return ret
	}
	return *o.UnknownCount
}

// GetUnknownCountOk returns a tuple with the UnknownCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorRuntimeMetrics) GetUnknownCountOk() (*int32, bool) {
	if o == nil || o.UnknownCount == nil {
		return nil, false
	}
	return o.UnknownCount, true
}

// HasUnknownCount returns a boolean if a field has been set.
func (o *MonitorRuntimeMetrics) HasUnknownCount() bool {
	if o != nil && o.UnknownCount != nil {
		return true
	}

	return false
}

// SetUnknownCount gets a reference to the given int32 and assigns it to the UnknownCount field.
func (o *MonitorRuntimeMetrics) SetUnknownCount(v int32) {
	o.UnknownCount = &v
}

// GetClearCount returns the ClearCount field value if set, zero value otherwise.
func (o *MonitorRuntimeMetrics) GetClearCount() int32 {
	if o == nil || o.ClearCount == nil {
		var ret int32
		return ret
	}
	return *o.ClearCount
}

// GetClearCountOk returns a tuple with the ClearCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorRuntimeMetrics) GetClearCountOk() (*int32, bool) {
	if o == nil || o.ClearCount == nil {
		return nil, false
	}
	return o.ClearCount, true
}

// HasClearCount returns a boolean if a field has been set.
func (o *MonitorRuntimeMetrics) HasClearCount() bool {
	if o != nil && o.ClearCount != nil {
		return true
	}

	return false
}

// SetClearCount gets a reference to the given int32 and assigns it to the ClearCount field.
func (o *MonitorRuntimeMetrics) SetClearCount(v int32) {
	o.ClearCount = &v
}

// GetDeviatingCount returns the DeviatingCount field value if set, zero value otherwise.
func (o *MonitorRuntimeMetrics) GetDeviatingCount() int32 {
	if o == nil || o.DeviatingCount == nil {
		var ret int32
		return ret
	}
	return *o.DeviatingCount
}

// GetDeviatingCountOk returns a tuple with the DeviatingCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorRuntimeMetrics) GetDeviatingCountOk() (*int32, bool) {
	if o == nil || o.DeviatingCount == nil {
		return nil, false
	}
	return o.DeviatingCount, true
}

// HasDeviatingCount returns a boolean if a field has been set.
func (o *MonitorRuntimeMetrics) HasDeviatingCount() bool {
	if o != nil && o.DeviatingCount != nil {
		return true
	}

	return false
}

// SetDeviatingCount gets a reference to the given int32 and assigns it to the DeviatingCount field.
func (o *MonitorRuntimeMetrics) SetDeviatingCount(v int32) {
	o.DeviatingCount = &v
}

// GetCriticalCount returns the CriticalCount field value if set, zero value otherwise.
func (o *MonitorRuntimeMetrics) GetCriticalCount() int32 {
	if o == nil || o.CriticalCount == nil {
		var ret int32
		return ret
	}
	return *o.CriticalCount
}

// GetCriticalCountOk returns a tuple with the CriticalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorRuntimeMetrics) GetCriticalCountOk() (*int32, bool) {
	if o == nil || o.CriticalCount == nil {
		return nil, false
	}
	return o.CriticalCount, true
}

// HasCriticalCount returns a boolean if a field has been set.
func (o *MonitorRuntimeMetrics) HasCriticalCount() bool {
	if o != nil && o.CriticalCount != nil {
		return true
	}

	return false
}

// SetCriticalCount gets a reference to the given int32 and assigns it to the CriticalCount field.
func (o *MonitorRuntimeMetrics) SetCriticalCount(v int32) {
	o.CriticalCount = &v
}

// GetLastRunTimestamp returns the LastRunTimestamp field value if set, zero value otherwise.
func (o *MonitorRuntimeMetrics) GetLastRunTimestamp() int64 {
	if o == nil || o.LastRunTimestamp == nil {
		var ret int64
		return ret
	}
	return *o.LastRunTimestamp
}

// GetLastRunTimestampOk returns a tuple with the LastRunTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorRuntimeMetrics) GetLastRunTimestampOk() (*int64, bool) {
	if o == nil || o.LastRunTimestamp == nil {
		return nil, false
	}
	return o.LastRunTimestamp, true
}

// HasLastRunTimestamp returns a boolean if a field has been set.
func (o *MonitorRuntimeMetrics) HasLastRunTimestamp() bool {
	if o != nil && o.LastRunTimestamp != nil {
		return true
	}

	return false
}

// SetLastRunTimestamp gets a reference to the given int64 and assigns it to the LastRunTimestamp field.
func (o *MonitorRuntimeMetrics) SetLastRunTimestamp(v int64) {
	o.LastRunTimestamp = &v
}

// GetLastSuccessfulRunTimestamp returns the LastSuccessfulRunTimestamp field value if set, zero value otherwise.
func (o *MonitorRuntimeMetrics) GetLastSuccessfulRunTimestamp() int64 {
	if o == nil || o.LastSuccessfulRunTimestamp == nil {
		var ret int64
		return ret
	}
	return *o.LastSuccessfulRunTimestamp
}

// GetLastSuccessfulRunTimestampOk returns a tuple with the LastSuccessfulRunTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorRuntimeMetrics) GetLastSuccessfulRunTimestampOk() (*int64, bool) {
	if o == nil || o.LastSuccessfulRunTimestamp == nil {
		return nil, false
	}
	return o.LastSuccessfulRunTimestamp, true
}

// HasLastSuccessfulRunTimestamp returns a boolean if a field has been set.
func (o *MonitorRuntimeMetrics) HasLastSuccessfulRunTimestamp() bool {
	if o != nil && o.LastSuccessfulRunTimestamp != nil {
		return true
	}

	return false
}

// SetLastSuccessfulRunTimestamp gets a reference to the given int64 and assigns it to the LastSuccessfulRunTimestamp field.
func (o *MonitorRuntimeMetrics) SetLastSuccessfulRunTimestamp(v int64) {
	o.LastSuccessfulRunTimestamp = &v
}

// GetLastFailedRunTimestamp returns the LastFailedRunTimestamp field value if set, zero value otherwise.
func (o *MonitorRuntimeMetrics) GetLastFailedRunTimestamp() int64 {
	if o == nil || o.LastFailedRunTimestamp == nil {
		var ret int64
		return ret
	}
	return *o.LastFailedRunTimestamp
}

// GetLastFailedRunTimestampOk returns a tuple with the LastFailedRunTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorRuntimeMetrics) GetLastFailedRunTimestampOk() (*int64, bool) {
	if o == nil || o.LastFailedRunTimestamp == nil {
		return nil, false
	}
	return o.LastFailedRunTimestamp, true
}

// HasLastFailedRunTimestamp returns a boolean if a field has been set.
func (o *MonitorRuntimeMetrics) HasLastFailedRunTimestamp() bool {
	if o != nil && o.LastFailedRunTimestamp != nil {
		return true
	}

	return false
}

// SetLastFailedRunTimestamp gets a reference to the given int64 and assigns it to the LastFailedRunTimestamp field.
func (o *MonitorRuntimeMetrics) SetLastFailedRunTimestamp(v int64) {
	o.LastFailedRunTimestamp = &v
}

func (o MonitorRuntimeMetrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnknownCount != nil {
		toSerialize["unknownCount"] = o.UnknownCount
	}
	if o.ClearCount != nil {
		toSerialize["clearCount"] = o.ClearCount
	}
	if o.DeviatingCount != nil {
		toSerialize["deviatingCount"] = o.DeviatingCount
	}
	if o.CriticalCount != nil {
		toSerialize["criticalCount"] = o.CriticalCount
	}
	if o.LastRunTimestamp != nil {
		toSerialize["lastRunTimestamp"] = o.LastRunTimestamp
	}
	if o.LastSuccessfulRunTimestamp != nil {
		toSerialize["lastSuccessfulRunTimestamp"] = o.LastSuccessfulRunTimestamp
	}
	if o.LastFailedRunTimestamp != nil {
		toSerialize["lastFailedRunTimestamp"] = o.LastFailedRunTimestamp
	}
	return json.Marshal(toSerialize)
}

type NullableMonitorRuntimeMetrics struct {
	value *MonitorRuntimeMetrics
	isSet bool
}

func (v NullableMonitorRuntimeMetrics) Get() *MonitorRuntimeMetrics {
	return v.value
}

func (v *NullableMonitorRuntimeMetrics) Set(val *MonitorRuntimeMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorRuntimeMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorRuntimeMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorRuntimeMetrics(val *MonitorRuntimeMetrics) *NullableMonitorRuntimeMetrics {
	return &NullableMonitorRuntimeMetrics{value: val, isSet: true}
}

func (v NullableMonitorRuntimeMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorRuntimeMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
