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

// Monitor struct for Monitor
type Monitor struct {
	Id                  int64                     `json:"id"`
	Name                string                    `json:"name"`
	Identifier          *string                   `json:"identifier,omitempty"`
	Description         *string                   `json:"description,omitempty"`
	FunctionId          int64                     `json:"functionId"`
	Arguments           []map[string]interface{}  `json:"arguments"`
	RemediationHint     *string                   `json:"remediationHint,omitempty"`
	IntervalSeconds     int32                     `json:"intervalSeconds"`
	Tags                []string                  `json:"tags"`
	Status              MonitorStatusValue        `json:"status"`
	RuntimeStatus       MonitorRuntimeStatusValue `json:"runtimeStatus"`
	LastUpdateTimestamp int64                     `json:"lastUpdateTimestamp"`
}

// NewMonitor instantiates a new Monitor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitor(id int64, name string, functionId int64, arguments []map[string]interface{}, intervalSeconds int32, tags []string, status MonitorStatusValue, runtimeStatus MonitorRuntimeStatusValue, lastUpdateTimestamp int64) *Monitor {
	this := Monitor{}
	this.Id = id
	this.Name = name
	this.FunctionId = functionId
	this.Arguments = arguments
	this.IntervalSeconds = intervalSeconds
	this.Tags = tags
	this.Status = status
	this.RuntimeStatus = runtimeStatus
	this.LastUpdateTimestamp = lastUpdateTimestamp
	return &this
}

// NewMonitorWithDefaults instantiates a new Monitor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorWithDefaults() *Monitor {
	this := Monitor{}
	return &this
}

// GetId returns the Id field value
func (o *Monitor) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Monitor) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Monitor) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Monitor) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Monitor) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Monitor) SetName(v string) {
	o.Name = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *Monitor) GetIdentifier() string {
	if o == nil || o.Identifier == nil {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monitor) GetIdentifierOk() (*string, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *Monitor) HasIdentifier() bool {
	if o != nil && o.Identifier != nil {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *Monitor) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Monitor) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monitor) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Monitor) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Monitor) SetDescription(v string) {
	o.Description = &v
}

// GetFunctionId returns the FunctionId field value
func (o *Monitor) GetFunctionId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FunctionId
}

// GetFunctionIdOk returns a tuple with the FunctionId field value
// and a boolean to check if the value has been set.
func (o *Monitor) GetFunctionIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FunctionId, true
}

// SetFunctionId sets field value
func (o *Monitor) SetFunctionId(v int64) {
	o.FunctionId = v
}

// GetArguments returns the Arguments field value
func (o *Monitor) GetArguments() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value
// and a boolean to check if the value has been set.
func (o *Monitor) GetArgumentsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Arguments, true
}

// SetArguments sets field value
func (o *Monitor) SetArguments(v []map[string]interface{}) {
	o.Arguments = v
}

// GetRemediationHint returns the RemediationHint field value if set, zero value otherwise.
func (o *Monitor) GetRemediationHint() string {
	if o == nil || o.RemediationHint == nil {
		var ret string
		return ret
	}
	return *o.RemediationHint
}

// GetRemediationHintOk returns a tuple with the RemediationHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monitor) GetRemediationHintOk() (*string, bool) {
	if o == nil || o.RemediationHint == nil {
		return nil, false
	}
	return o.RemediationHint, true
}

// HasRemediationHint returns a boolean if a field has been set.
func (o *Monitor) HasRemediationHint() bool {
	if o != nil && o.RemediationHint != nil {
		return true
	}

	return false
}

// SetRemediationHint gets a reference to the given string and assigns it to the RemediationHint field.
func (o *Monitor) SetRemediationHint(v string) {
	o.RemediationHint = &v
}

// GetIntervalSeconds returns the IntervalSeconds field value
func (o *Monitor) GetIntervalSeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IntervalSeconds
}

// GetIntervalSecondsOk returns a tuple with the IntervalSeconds field value
// and a boolean to check if the value has been set.
func (o *Monitor) GetIntervalSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntervalSeconds, true
}

// SetIntervalSeconds sets field value
func (o *Monitor) SetIntervalSeconds(v int32) {
	o.IntervalSeconds = v
}

// GetTags returns the Tags field value
func (o *Monitor) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *Monitor) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *Monitor) SetTags(v []string) {
	o.Tags = v
}

// GetStatus returns the Status field value
func (o *Monitor) GetStatus() MonitorStatusValue {
	if o == nil {
		var ret MonitorStatusValue
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Monitor) GetStatusOk() (*MonitorStatusValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Monitor) SetStatus(v MonitorStatusValue) {
	o.Status = v
}

// GetRuntimeStatus returns the RuntimeStatus field value
func (o *Monitor) GetRuntimeStatus() MonitorRuntimeStatusValue {
	if o == nil {
		var ret MonitorRuntimeStatusValue
		return ret
	}

	return o.RuntimeStatus
}

// GetRuntimeStatusOk returns a tuple with the RuntimeStatus field value
// and a boolean to check if the value has been set.
func (o *Monitor) GetRuntimeStatusOk() (*MonitorRuntimeStatusValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuntimeStatus, true
}

// SetRuntimeStatus sets field value
func (o *Monitor) SetRuntimeStatus(v MonitorRuntimeStatusValue) {
	o.RuntimeStatus = v
}

// GetLastUpdateTimestamp returns the LastUpdateTimestamp field value
func (o *Monitor) GetLastUpdateTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.LastUpdateTimestamp
}

// GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field value
// and a boolean to check if the value has been set.
func (o *Monitor) GetLastUpdateTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdateTimestamp, true
}

// SetLastUpdateTimestamp sets field value
func (o *Monitor) SetLastUpdateTimestamp(v int64) {
	o.LastUpdateTimestamp = v
}

func (o Monitor) MarshalJSON() ([]byte, error) {
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
		toSerialize["functionId"] = o.FunctionId
	}
	if true {
		toSerialize["arguments"] = o.Arguments
	}
	if o.RemediationHint != nil {
		toSerialize["remediationHint"] = o.RemediationHint
	}
	if true {
		toSerialize["intervalSeconds"] = o.IntervalSeconds
	}
	if true {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["runtimeStatus"] = o.RuntimeStatus
	}
	if true {
		toSerialize["lastUpdateTimestamp"] = o.LastUpdateTimestamp
	}
	return json.Marshal(toSerialize)
}

type NullableMonitor struct {
	value *Monitor
	isSet bool
}

func (v NullableMonitor) Get() *Monitor {
	return v.value
}

func (v *NullableMonitor) Set(val *Monitor) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitor) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitor(val *Monitor) *NullableMonitor {
	return &NullableMonitor{value: val, isSet: true}
}

func (v NullableMonitor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
