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

// PointInner struct for PointInner
type PointInner struct {
	float64 *float64
	int64 *int64
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PointInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into float64
	err = json.Unmarshal(data, &dst.float64);
	if err == nil {
		jsonfloat64, _ := json.Marshal(dst.float64)
		if string(jsonfloat64) == "{}" { // empty struct
			dst.float64 = nil
		} else {
			return nil // data stored in dst.float64, return on the first match
		}
	} else {
		dst.float64 = nil
	}

	// try to unmarshal JSON data into int64
	err = json.Unmarshal(data, &dst.int64);
	if err == nil {
		jsonint64, _ := json.Marshal(dst.int64)
		if string(jsonint64) == "{}" { // empty struct
			dst.int64 = nil
		} else {
			return nil // data stored in dst.int64, return on the first match
		}
	} else {
		dst.int64 = nil
	}

	return fmt.Errorf("Data failed to match schemas in anyOf(PointInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PointInner) MarshalJSON() ([]byte, error) {
	if src.float64 != nil {
		return json.Marshal(&src.float64)
	}

	if src.int64 != nil {
		return json.Marshal(&src.int64)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePointInner struct {
	value *PointInner
	isSet bool
}

func (v NullablePointInner) Get() *PointInner {
	return v.value
}

func (v *NullablePointInner) Set(val *PointInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePointInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePointInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePointInner(val *PointInner) *NullablePointInner {
	return &NullablePointInner{value: val, isSet: true}
}

func (v NullablePointInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePointInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


