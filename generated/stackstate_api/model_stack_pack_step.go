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

// StackPackStep struct for StackPackStep
type StackPackStep struct {
	Display *string             `json:"display,omitempty"`
	Name    *string             `json:"name,omitempty"`
	Value   *StackPackStepValue `json:"value,omitempty"`
}

// NewStackPackStep instantiates a new StackPackStep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStackPackStep() *StackPackStep {
	this := StackPackStep{}
	return &this
}

// NewStackPackStepWithDefaults instantiates a new StackPackStep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStackPackStepWithDefaults() *StackPackStep {
	this := StackPackStep{}
	return &this
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *StackPackStep) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackPackStep) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *StackPackStep) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *StackPackStep) SetDisplay(v string) {
	o.Display = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StackPackStep) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackPackStep) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StackPackStep) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StackPackStep) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *StackPackStep) GetValue() StackPackStepValue {
	if o == nil || o.Value == nil {
		var ret StackPackStepValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackPackStep) GetValueOk() (*StackPackStepValue, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *StackPackStep) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given StackPackStepValue and assigns it to the Value field.
func (o *StackPackStep) SetValue(v StackPackStepValue) {
	o.Value = &v
}

func (o StackPackStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Display != nil {
		toSerialize["display"] = o.Display
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableStackPackStep struct {
	value *StackPackStep
	isSet bool
}

func (v NullableStackPackStep) Get() *StackPackStep {
	return v.value
}

func (v *NullableStackPackStep) Set(val *StackPackStep) {
	v.value = val
	v.isSet = true
}

func (v NullableStackPackStep) IsSet() bool {
	return v.isSet
}

func (v *NullableStackPackStep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStackPackStep(val *StackPackStep) *NullableStackPackStep {
	return &NullableStackPackStep{value: val, isSet: true}
}

func (v NullableStackPackStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStackPackStep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
