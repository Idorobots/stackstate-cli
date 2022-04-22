/*
StackState API

StackState's API specification

API version: 0.0.1
Contact: info@stackstate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackstate_client

import (
	"encoding/json"
)

// NodeTypes struct for NodeTypes
type NodeTypes struct {
	NodeTypes []NodeTypesNodeTypesInner `json:"nodeTypes"`
}

// NewNodeTypes instantiates a new NodeTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeTypes(nodeTypes []NodeTypesNodeTypesInner) *NodeTypes {
	this := NodeTypes{}
	this.NodeTypes = nodeTypes
	return &this
}

// NewNodeTypesWithDefaults instantiates a new NodeTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeTypesWithDefaults() *NodeTypes {
	this := NodeTypes{}
	return &this
}

// GetNodeTypes returns the NodeTypes field value
func (o *NodeTypes) GetNodeTypes() []NodeTypesNodeTypesInner {
	if o == nil {
		var ret []NodeTypesNodeTypesInner
		return ret
	}

	return o.NodeTypes
}

// GetNodeTypesOk returns a tuple with the NodeTypes field value
// and a boolean to check if the value has been set.
func (o *NodeTypes) GetNodeTypesOk() ([]NodeTypesNodeTypesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeTypes, true
}

// SetNodeTypes sets field value
func (o *NodeTypes) SetNodeTypes(v []NodeTypesNodeTypesInner) {
	o.NodeTypes = v
}

func (o NodeTypes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["nodeTypes"] = o.NodeTypes
	}
	return json.Marshal(toSerialize)
}

type NullableNodeTypes struct {
	value *NodeTypes
	isSet bool
}

func (v NullableNodeTypes) Get() *NodeTypes {
	return v.value
}

func (v *NullableNodeTypes) Set(val *NodeTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeTypes(val *NodeTypes) *NullableNodeTypes {
	return &NullableNodeTypes{value: val, isSet: true}
}

func (v NullableNodeTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


