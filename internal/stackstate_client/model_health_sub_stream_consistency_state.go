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
	"fmt"
)

// HealthSubStreamConsistencyState - struct for HealthSubStreamConsistencyState
type HealthSubStreamConsistencyState struct {
	HealthSubStreamExpiry *HealthSubStreamExpiry
	HealthSubStreamSnapshot *HealthSubStreamSnapshot
	HealthSubStreamTransactionalIncrements *HealthSubStreamTransactionalIncrements
}

// HealthSubStreamExpiryAsHealthSubStreamConsistencyState is a convenience function that returns HealthSubStreamExpiry wrapped in HealthSubStreamConsistencyState
func HealthSubStreamExpiryAsHealthSubStreamConsistencyState(v *HealthSubStreamExpiry) HealthSubStreamConsistencyState {
	return HealthSubStreamConsistencyState{ HealthSubStreamExpiry: v}
}

// HealthSubStreamSnapshotAsHealthSubStreamConsistencyState is a convenience function that returns HealthSubStreamSnapshot wrapped in HealthSubStreamConsistencyState
func HealthSubStreamSnapshotAsHealthSubStreamConsistencyState(v *HealthSubStreamSnapshot) HealthSubStreamConsistencyState {
	return HealthSubStreamConsistencyState{ HealthSubStreamSnapshot: v}
}

// HealthSubStreamTransactionalIncrementsAsHealthSubStreamConsistencyState is a convenience function that returns HealthSubStreamTransactionalIncrements wrapped in HealthSubStreamConsistencyState
func HealthSubStreamTransactionalIncrementsAsHealthSubStreamConsistencyState(v *HealthSubStreamTransactionalIncrements) HealthSubStreamConsistencyState {
	return HealthSubStreamConsistencyState{ HealthSubStreamTransactionalIncrements: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *HealthSubStreamConsistencyState) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'HealthSubStreamExpiry'
	if jsonDict["_type"] == "HealthSubStreamExpiry" {
		// try to unmarshal JSON data into HealthSubStreamExpiry
		err = json.Unmarshal(data, &dst.HealthSubStreamExpiry)
		if err == nil {
			return nil // data stored in dst.HealthSubStreamExpiry, return on the first match
		} else {
			dst.HealthSubStreamExpiry = nil
			return fmt.Errorf("Failed to unmarshal HealthSubStreamConsistencyState as HealthSubStreamExpiry: %s", err.Error())
		}
	}

	// check if the discriminator value is 'HealthSubStreamSnapshot'
	if jsonDict["_type"] == "HealthSubStreamSnapshot" {
		// try to unmarshal JSON data into HealthSubStreamSnapshot
		err = json.Unmarshal(data, &dst.HealthSubStreamSnapshot)
		if err == nil {
			return nil // data stored in dst.HealthSubStreamSnapshot, return on the first match
		} else {
			dst.HealthSubStreamSnapshot = nil
			return fmt.Errorf("Failed to unmarshal HealthSubStreamConsistencyState as HealthSubStreamSnapshot: %s", err.Error())
		}
	}

	// check if the discriminator value is 'HealthSubStreamTransactionalIncrements'
	if jsonDict["_type"] == "HealthSubStreamTransactionalIncrements" {
		// try to unmarshal JSON data into HealthSubStreamTransactionalIncrements
		err = json.Unmarshal(data, &dst.HealthSubStreamTransactionalIncrements)
		if err == nil {
			return nil // data stored in dst.HealthSubStreamTransactionalIncrements, return on the first match
		} else {
			dst.HealthSubStreamTransactionalIncrements = nil
			return fmt.Errorf("Failed to unmarshal HealthSubStreamConsistencyState as HealthSubStreamTransactionalIncrements: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src HealthSubStreamConsistencyState) MarshalJSON() ([]byte, error) {
	if src.HealthSubStreamExpiry != nil {
		return json.Marshal(&src.HealthSubStreamExpiry)
	}

	if src.HealthSubStreamSnapshot != nil {
		return json.Marshal(&src.HealthSubStreamSnapshot)
	}

	if src.HealthSubStreamTransactionalIncrements != nil {
		return json.Marshal(&src.HealthSubStreamTransactionalIncrements)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *HealthSubStreamConsistencyState) GetActualInstance() (interface{}) {
	if obj.HealthSubStreamExpiry != nil {
		return obj.HealthSubStreamExpiry
	}

	if obj.HealthSubStreamSnapshot != nil {
		return obj.HealthSubStreamSnapshot
	}

	if obj.HealthSubStreamTransactionalIncrements != nil {
		return obj.HealthSubStreamTransactionalIncrements
	}

	// all schemas are nil
	return nil
}

type NullableHealthSubStreamConsistencyState struct {
	value *HealthSubStreamConsistencyState
	isSet bool
}

func (v NullableHealthSubStreamConsistencyState) Get() *HealthSubStreamConsistencyState {
	return v.value
}

func (v *NullableHealthSubStreamConsistencyState) Set(val *HealthSubStreamConsistencyState) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthSubStreamConsistencyState) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthSubStreamConsistencyState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthSubStreamConsistencyState(val *HealthSubStreamConsistencyState) *NullableHealthSubStreamConsistencyState {
	return &NullableHealthSubStreamConsistencyState{value: val, isSet: true}
}

func (v NullableHealthSubStreamConsistencyState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthSubStreamConsistencyState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


