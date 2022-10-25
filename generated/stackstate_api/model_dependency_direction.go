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
	"fmt"
)

// DependencyDirection the model 'DependencyDirection'
type DependencyDirection string

// List of DependencyDirection
const (
	DEPENDENCYDIRECTION_ONE_WAY DependencyDirection = "one-way"
	DEPENDENCYDIRECTION_NONE    DependencyDirection = "none"
	DEPENDENCYDIRECTION_BOTH    DependencyDirection = "both"
)

// All allowed values of DependencyDirection enum
var AllowedDependencyDirectionEnumValues = []DependencyDirection{
	"one-way",
	"none",
	"both",
}

func (v *DependencyDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DependencyDirection(value)
	for _, existing := range AllowedDependencyDirectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DependencyDirection", value)
}

// NewDependencyDirectionFromValue returns a pointer to a valid DependencyDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDependencyDirectionFromValue(v string) (*DependencyDirection, error) {
	ev := DependencyDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DependencyDirection: valid values are %v", v, AllowedDependencyDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DependencyDirection) IsValid() bool {
	for _, existing := range AllowedDependencyDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DependencyDirection value
func (v DependencyDirection) Ptr() *DependencyDirection {
	return &v
}

type NullableDependencyDirection struct {
	value *DependencyDirection
	isSet bool
}

func (v NullableDependencyDirection) Get() *DependencyDirection {
	return v.value
}

func (v *NullableDependencyDirection) Set(val *DependencyDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableDependencyDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableDependencyDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDependencyDirection(val *DependencyDirection) *NullableDependencyDirection {
	return &NullableDependencyDirection{value: val, isSet: true}
}

func (v NullableDependencyDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDependencyDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
