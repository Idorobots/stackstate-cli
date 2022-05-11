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

// HealthSubStreamError - struct for HealthSubStreamError
type HealthSubStreamError struct {
	HealthStreamNotFound *HealthStreamNotFound
	HealthSubStreamNotFound *HealthSubStreamNotFound
}

// HealthStreamNotFoundAsHealthSubStreamError is a convenience function that returns HealthStreamNotFound wrapped in HealthSubStreamError
func HealthStreamNotFoundAsHealthSubStreamError(v *HealthStreamNotFound) HealthSubStreamError {
	return HealthSubStreamError{
		HealthStreamNotFound: v,
	}
}

// HealthSubStreamNotFoundAsHealthSubStreamError is a convenience function that returns HealthSubStreamNotFound wrapped in HealthSubStreamError
func HealthSubStreamNotFoundAsHealthSubStreamError(v *HealthSubStreamNotFound) HealthSubStreamError {
	return HealthSubStreamError{
		HealthSubStreamNotFound: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *HealthSubStreamError) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'HealthStreamNotFound'
	if jsonDict["_type"] == "HealthStreamNotFound" {
		// try to unmarshal JSON data into HealthStreamNotFound
		err = json.Unmarshal(data, &dst.HealthStreamNotFound)
		if err == nil {
			return nil // data stored in dst.HealthStreamNotFound, return on the first match
		} else {
			dst.HealthStreamNotFound = nil
			return fmt.Errorf("Failed to unmarshal HealthSubStreamError as HealthStreamNotFound: %s", err.Error())
		}
	}

	// check if the discriminator value is 'HealthSubStreamNotFound'
	if jsonDict["_type"] == "HealthSubStreamNotFound" {
		// try to unmarshal JSON data into HealthSubStreamNotFound
		err = json.Unmarshal(data, &dst.HealthSubStreamNotFound)
		if err == nil {
			return nil // data stored in dst.HealthSubStreamNotFound, return on the first match
		} else {
			dst.HealthSubStreamNotFound = nil
			return fmt.Errorf("Failed to unmarshal HealthSubStreamError as HealthSubStreamNotFound: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src HealthSubStreamError) MarshalJSON() ([]byte, error) {
	if src.HealthStreamNotFound != nil {
		return json.Marshal(&src.HealthStreamNotFound)
	}

	if src.HealthSubStreamNotFound != nil {
		return json.Marshal(&src.HealthSubStreamNotFound)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *HealthSubStreamError) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.HealthStreamNotFound != nil {
		return obj.HealthStreamNotFound
	}

	if obj.HealthSubStreamNotFound != nil {
		return obj.HealthSubStreamNotFound
	}

	// all schemas are nil
	return nil
}

type NullableHealthSubStreamError struct {
	value *HealthSubStreamError
	isSet bool
}

func (v NullableHealthSubStreamError) Get() *HealthSubStreamError {
	return v.value
}

func (v *NullableHealthSubStreamError) Set(val *HealthSubStreamError) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthSubStreamError) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthSubStreamError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthSubStreamError(val *HealthSubStreamError) *NullableHealthSubStreamError {
	return &NullableHealthSubStreamError{value: val, isSet: true}
}

func (v NullableHealthSubStreamError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthSubStreamError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


