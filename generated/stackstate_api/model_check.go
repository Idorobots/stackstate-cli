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

// Check struct for Check
type Check struct {
	Type string `json:"_type"`
	Arguments []Argument `json:"arguments"`
	Description *string `json:"description,omitempty"`
	Function int64 `json:"function"`
	Id *int64 `json:"id,omitempty"`
	LastUpdateTimestamp *int64 `json:"lastUpdateTimestamp,omitempty"`
	Name string `json:"name"`
	RemediationHint *string `json:"remediationHint,omitempty"`
	State *CheckState `json:"state,omitempty"`
	SyncCreated bool `json:"syncCreated"`
}

// NewCheck instantiates a new Check object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheck(type_ string, arguments []Argument, function int64, name string, syncCreated bool) *Check {
	this := Check{}
	this.Type = type_
	this.Arguments = arguments
	this.Function = function
	this.Name = name
	this.SyncCreated = syncCreated
	return &this
}

// NewCheckWithDefaults instantiates a new Check object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckWithDefaults() *Check {
	this := Check{}
	return &this
}

// GetType returns the Type field value
func (o *Check) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Check) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Check) SetType(v string) {
	o.Type = v
}

// GetArguments returns the Arguments field value
func (o *Check) GetArguments() []Argument {
	if o == nil {
		var ret []Argument
		return ret
	}

	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value
// and a boolean to check if the value has been set.
func (o *Check) GetArgumentsOk() ([]Argument, bool) {
	if o == nil {
		return nil, false
	}
	return o.Arguments, true
}

// SetArguments sets field value
func (o *Check) SetArguments(v []Argument) {
	o.Arguments = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Check) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Check) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Check) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Check) SetDescription(v string) {
	o.Description = &v
}

// GetFunction returns the Function field value
func (o *Check) GetFunction() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Function
}

// GetFunctionOk returns a tuple with the Function field value
// and a boolean to check if the value has been set.
func (o *Check) GetFunctionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Function, true
}

// SetFunction sets field value
func (o *Check) SetFunction(v int64) {
	o.Function = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Check) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Check) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Check) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Check) SetId(v int64) {
	o.Id = &v
}

// GetLastUpdateTimestamp returns the LastUpdateTimestamp field value if set, zero value otherwise.
func (o *Check) GetLastUpdateTimestamp() int64 {
	if o == nil || o.LastUpdateTimestamp == nil {
		var ret int64
		return ret
	}
	return *o.LastUpdateTimestamp
}

// GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Check) GetLastUpdateTimestampOk() (*int64, bool) {
	if o == nil || o.LastUpdateTimestamp == nil {
		return nil, false
	}
	return o.LastUpdateTimestamp, true
}

// HasLastUpdateTimestamp returns a boolean if a field has been set.
func (o *Check) HasLastUpdateTimestamp() bool {
	if o != nil && o.LastUpdateTimestamp != nil {
		return true
	}

	return false
}

// SetLastUpdateTimestamp gets a reference to the given int64 and assigns it to the LastUpdateTimestamp field.
func (o *Check) SetLastUpdateTimestamp(v int64) {
	o.LastUpdateTimestamp = &v
}

// GetName returns the Name field value
func (o *Check) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Check) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Check) SetName(v string) {
	o.Name = v
}

// GetRemediationHint returns the RemediationHint field value if set, zero value otherwise.
func (o *Check) GetRemediationHint() string {
	if o == nil || o.RemediationHint == nil {
		var ret string
		return ret
	}
	return *o.RemediationHint
}

// GetRemediationHintOk returns a tuple with the RemediationHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Check) GetRemediationHintOk() (*string, bool) {
	if o == nil || o.RemediationHint == nil {
		return nil, false
	}
	return o.RemediationHint, true
}

// HasRemediationHint returns a boolean if a field has been set.
func (o *Check) HasRemediationHint() bool {
	if o != nil && o.RemediationHint != nil {
		return true
	}

	return false
}

// SetRemediationHint gets a reference to the given string and assigns it to the RemediationHint field.
func (o *Check) SetRemediationHint(v string) {
	o.RemediationHint = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Check) GetState() CheckState {
	if o == nil || o.State == nil {
		var ret CheckState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Check) GetStateOk() (*CheckState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Check) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given CheckState and assigns it to the State field.
func (o *Check) SetState(v CheckState) {
	o.State = &v
}

// GetSyncCreated returns the SyncCreated field value
func (o *Check) GetSyncCreated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SyncCreated
}

// GetSyncCreatedOk returns a tuple with the SyncCreated field value
// and a boolean to check if the value has been set.
func (o *Check) GetSyncCreatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SyncCreated, true
}

// SetSyncCreated sets field value
func (o *Check) SetSyncCreated(v bool) {
	o.SyncCreated = v
}

func (o Check) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if true {
		toSerialize["arguments"] = o.Arguments
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["function"] = o.Function
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdateTimestamp != nil {
		toSerialize["lastUpdateTimestamp"] = o.LastUpdateTimestamp
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.RemediationHint != nil {
		toSerialize["remediationHint"] = o.RemediationHint
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["syncCreated"] = o.SyncCreated
	}
	return json.Marshal(toSerialize)
}

type NullableCheck struct {
	value *Check
	isSet bool
}

func (v NullableCheck) Get() *Check {
	return v.value
}

func (v *NullableCheck) Set(val *Check) {
	v.value = val
	v.isSet = true
}

func (v NullableCheck) IsSet() bool {
	return v.isSet
}

func (v *NullableCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheck(val *Check) *NullableCheck {
	return &NullableCheck{value: val, isSet: true}
}

func (v NullableCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


