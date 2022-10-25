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

// EventRelation struct for EventRelation
type EventRelation struct {
	Type                string              `json:"_type"`
	Id                  int64               `json:"id"`
	RelationTypeId      int64               `json:"relationTypeId"`
	Name                *string             `json:"name,omitempty"`
	Source              EventComponent      `json:"source"`
	Target              EventComponent      `json:"target"`
	DependencyDirection DependencyDirection `json:"dependencyDirection"`
}

// NewEventRelation instantiates a new EventRelation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventRelation(type_ string, id int64, relationTypeId int64, source EventComponent, target EventComponent, dependencyDirection DependencyDirection) *EventRelation {
	this := EventRelation{}
	this.Type = type_
	this.Id = id
	this.RelationTypeId = relationTypeId
	this.Source = source
	this.Target = target
	this.DependencyDirection = dependencyDirection
	return &this
}

// NewEventRelationWithDefaults instantiates a new EventRelation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventRelationWithDefaults() *EventRelation {
	this := EventRelation{}
	return &this
}

// GetType returns the Type field value
func (o *EventRelation) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EventRelation) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EventRelation) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *EventRelation) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EventRelation) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EventRelation) SetId(v int64) {
	o.Id = v
}

// GetRelationTypeId returns the RelationTypeId field value
func (o *EventRelation) GetRelationTypeId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RelationTypeId
}

// GetRelationTypeIdOk returns a tuple with the RelationTypeId field value
// and a boolean to check if the value has been set.
func (o *EventRelation) GetRelationTypeIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelationTypeId, true
}

// SetRelationTypeId sets field value
func (o *EventRelation) SetRelationTypeId(v int64) {
	o.RelationTypeId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EventRelation) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventRelation) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EventRelation) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EventRelation) SetName(v string) {
	o.Name = &v
}

// GetSource returns the Source field value
func (o *EventRelation) GetSource() EventComponent {
	if o == nil {
		var ret EventComponent
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *EventRelation) GetSourceOk() (*EventComponent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *EventRelation) SetSource(v EventComponent) {
	o.Source = v
}

// GetTarget returns the Target field value
func (o *EventRelation) GetTarget() EventComponent {
	if o == nil {
		var ret EventComponent
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *EventRelation) GetTargetOk() (*EventComponent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *EventRelation) SetTarget(v EventComponent) {
	o.Target = v
}

// GetDependencyDirection returns the DependencyDirection field value
func (o *EventRelation) GetDependencyDirection() DependencyDirection {
	if o == nil {
		var ret DependencyDirection
		return ret
	}

	return o.DependencyDirection
}

// GetDependencyDirectionOk returns a tuple with the DependencyDirection field value
// and a boolean to check if the value has been set.
func (o *EventRelation) GetDependencyDirectionOk() (*DependencyDirection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DependencyDirection, true
}

// SetDependencyDirection sets field value
func (o *EventRelation) SetDependencyDirection(v DependencyDirection) {
	o.DependencyDirection = v
}

func (o EventRelation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["relationTypeId"] = o.RelationTypeId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["target"] = o.Target
	}
	if true {
		toSerialize["dependencyDirection"] = o.DependencyDirection
	}
	return json.Marshal(toSerialize)
}

type NullableEventRelation struct {
	value *EventRelation
	isSet bool
}

func (v NullableEventRelation) Get() *EventRelation {
	return v.value
}

func (v *NullableEventRelation) Set(val *EventRelation) {
	v.value = val
	v.isSet = true
}

func (v NullableEventRelation) IsSet() bool {
	return v.isSet
}

func (v *NullableEventRelation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventRelation(val *EventRelation) *NullableEventRelation {
	return &NullableEventRelation{value: val, isSet: true}
}

func (v NullableEventRelation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventRelation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
