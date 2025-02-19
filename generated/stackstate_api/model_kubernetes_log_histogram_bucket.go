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

// KubernetesLogHistogramBucket struct for KubernetesLogHistogramBucket
type KubernetesLogHistogramBucket struct {
	// Total logs record count in the bucket.
	Count int64 `json:"count"`
	// The bucket initial timestamp.
	StartTime int32 `json:"startTime"`
	// The bucket final timestamp.
	EndTime int32 `json:"endTime"`
}

// NewKubernetesLogHistogramBucket instantiates a new KubernetesLogHistogramBucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesLogHistogramBucket(count int64, startTime int32, endTime int32) *KubernetesLogHistogramBucket {
	this := KubernetesLogHistogramBucket{}
	this.Count = count
	this.StartTime = startTime
	this.EndTime = endTime
	return &this
}

// NewKubernetesLogHistogramBucketWithDefaults instantiates a new KubernetesLogHistogramBucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesLogHistogramBucketWithDefaults() *KubernetesLogHistogramBucket {
	this := KubernetesLogHistogramBucket{}
	return &this
}

// GetCount returns the Count field value
func (o *KubernetesLogHistogramBucket) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *KubernetesLogHistogramBucket) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *KubernetesLogHistogramBucket) SetCount(v int64) {
	o.Count = v
}

// GetStartTime returns the StartTime field value
func (o *KubernetesLogHistogramBucket) GetStartTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *KubernetesLogHistogramBucket) GetStartTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *KubernetesLogHistogramBucket) SetStartTime(v int32) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *KubernetesLogHistogramBucket) GetEndTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *KubernetesLogHistogramBucket) GetEndTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *KubernetesLogHistogramBucket) SetEndTime(v int32) {
	o.EndTime = v
}

func (o KubernetesLogHistogramBucket) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["count"] = o.Count
	}
	if true {
		toSerialize["startTime"] = o.StartTime
	}
	if true {
		toSerialize["endTime"] = o.EndTime
	}
	return json.Marshal(toSerialize)
}

type NullableKubernetesLogHistogramBucket struct {
	value *KubernetesLogHistogramBucket
	isSet bool
}

func (v NullableKubernetesLogHistogramBucket) Get() *KubernetesLogHistogramBucket {
	return v.value
}

func (v *NullableKubernetesLogHistogramBucket) Set(val *KubernetesLogHistogramBucket) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesLogHistogramBucket) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesLogHistogramBucket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesLogHistogramBucket(val *KubernetesLogHistogramBucket) *NullableKubernetesLogHistogramBucket {
	return &NullableKubernetesLogHistogramBucket{value: val, isSet: true}
}

func (v NullableKubernetesLogHistogramBucket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesLogHistogramBucket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
