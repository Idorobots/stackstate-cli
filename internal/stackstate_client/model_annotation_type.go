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

// AnnotationType the model 'AnnotationType'
type AnnotationType string

// List of AnnotationType
const (
	METRIC_STREAM_ANOMALY AnnotationType = "MetricStreamAnomaly"
	METRIC_STREAM_NO_ANOMALY AnnotationType = "MetricStreamNoAnomaly"
	ANOMALY_FEEDBACK AnnotationType = "AnomalyFeedback"
)

// All allowed values of AnnotationType enum
var AllowedAnnotationTypeEnumValues = []AnnotationType{
	"MetricStreamAnomaly",
	"MetricStreamNoAnomaly",
	"AnomalyFeedback",
}

func (v *AnnotationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AnnotationType(value)
	for _, existing := range AllowedAnnotationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AnnotationType", value)
}

// NewAnnotationTypeFromValue returns a pointer to a valid AnnotationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAnnotationTypeFromValue(v string) (*AnnotationType, error) {
	ev := AnnotationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AnnotationType: valid values are %v", v, AllowedAnnotationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AnnotationType) IsValid() bool {
	for _, existing := range AllowedAnnotationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AnnotationType value
func (v AnnotationType) Ptr() *AnnotationType {
	return &v
}

type NullableAnnotationType struct {
	value *AnnotationType
	isSet bool
}

func (v NullableAnnotationType) Get() *AnnotationType {
	return v.value
}

func (v *NullableAnnotationType) Set(val *AnnotationType) {
	v.value = val
	v.isSet = true
}

func (v NullableAnnotationType) IsSet() bool {
	return v.isSet
}

func (v *NullableAnnotationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnnotationType(val *AnnotationType) *NullableAnnotationType {
	return &NullableAnnotationType{value: val, isSet: true}
}

func (v NullableAnnotationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnnotationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

