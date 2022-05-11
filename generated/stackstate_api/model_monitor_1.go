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

// Monitor1 struct for Monitor1
type Monitor1 struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Identifier *string `json:"identifier,omitempty"`
	Description *string `json:"description,omitempty"`
	FunctionId int64 `json:"functionId"`
	Arguments []map[string]interface{} `json:"arguments"`
	RemediationHint *string `json:"remediationHint,omitempty"`
	TopologyMapping string `json:"topologyMapping"`
	IntervalSeconds int32 `json:"intervalSeconds"`
}

// NewMonitor1 instantiates a new Monitor1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitor1(id int64, name string, functionId int64, arguments []map[string]interface{}, topologyMapping string, intervalSeconds int32) *Monitor1 {
	this := Monitor1{}
	this.Id = id
	this.Name = name
	this.FunctionId = functionId
	this.Arguments = arguments
	this.TopologyMapping = topologyMapping
	this.IntervalSeconds = intervalSeconds
	return &this
}

// NewMonitor1WithDefaults instantiates a new Monitor1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitor1WithDefaults() *Monitor1 {
	this := Monitor1{}
	return &this
}

// GetId returns the Id field value
func (o *Monitor1) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Monitor1) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Monitor1) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Monitor1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Monitor1) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Monitor1) SetName(v string) {
	o.Name = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *Monitor1) GetIdentifier() string {
	if o == nil || o.Identifier == nil {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monitor1) GetIdentifierOk() (*string, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *Monitor1) HasIdentifier() bool {
	if o != nil && o.Identifier != nil {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *Monitor1) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Monitor1) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monitor1) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Monitor1) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Monitor1) SetDescription(v string) {
	o.Description = &v
}

// GetFunctionId returns the FunctionId field value
func (o *Monitor1) GetFunctionId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FunctionId
}

// GetFunctionIdOk returns a tuple with the FunctionId field value
// and a boolean to check if the value has been set.
func (o *Monitor1) GetFunctionIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FunctionId, true
}

// SetFunctionId sets field value
func (o *Monitor1) SetFunctionId(v int64) {
	o.FunctionId = v
}

// GetArguments returns the Arguments field value
func (o *Monitor1) GetArguments() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value
// and a boolean to check if the value has been set.
func (o *Monitor1) GetArgumentsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Arguments, true
}

// SetArguments sets field value
func (o *Monitor1) SetArguments(v []map[string]interface{}) {
	o.Arguments = v
}

// GetRemediationHint returns the RemediationHint field value if set, zero value otherwise.
func (o *Monitor1) GetRemediationHint() string {
	if o == nil || o.RemediationHint == nil {
		var ret string
		return ret
	}
	return *o.RemediationHint
}

// GetRemediationHintOk returns a tuple with the RemediationHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monitor1) GetRemediationHintOk() (*string, bool) {
	if o == nil || o.RemediationHint == nil {
		return nil, false
	}
	return o.RemediationHint, true
}

// HasRemediationHint returns a boolean if a field has been set.
func (o *Monitor1) HasRemediationHint() bool {
	if o != nil && o.RemediationHint != nil {
		return true
	}

	return false
}

// SetRemediationHint gets a reference to the given string and assigns it to the RemediationHint field.
func (o *Monitor1) SetRemediationHint(v string) {
	o.RemediationHint = &v
}

// GetTopologyMapping returns the TopologyMapping field value
func (o *Monitor1) GetTopologyMapping() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TopologyMapping
}

// GetTopologyMappingOk returns a tuple with the TopologyMapping field value
// and a boolean to check if the value has been set.
func (o *Monitor1) GetTopologyMappingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopologyMapping, true
}

// SetTopologyMapping sets field value
func (o *Monitor1) SetTopologyMapping(v string) {
	o.TopologyMapping = v
}

// GetIntervalSeconds returns the IntervalSeconds field value
func (o *Monitor1) GetIntervalSeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IntervalSeconds
}

// GetIntervalSecondsOk returns a tuple with the IntervalSeconds field value
// and a boolean to check if the value has been set.
func (o *Monitor1) GetIntervalSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntervalSeconds, true
}

// SetIntervalSeconds sets field value
func (o *Monitor1) SetIntervalSeconds(v int32) {
	o.IntervalSeconds = v
}

func (o Monitor1) MarshalJSON() ([]byte, error) {
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
		toSerialize["topologyMapping"] = o.TopologyMapping
	}
	if true {
		toSerialize["intervalSeconds"] = o.IntervalSeconds
	}
	return json.Marshal(toSerialize)
}

type NullableMonitor1 struct {
	value *Monitor1
	isSet bool
}

func (v NullableMonitor1) Get() *Monitor1 {
	return v.value
}

func (v *NullableMonitor1) Set(val *Monitor1) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitor1) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitor1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitor1(val *Monitor1) *NullableMonitor1 {
	return &NullableMonitor1{value: val, isSet: true}
}

func (v NullableMonitor1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitor1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


