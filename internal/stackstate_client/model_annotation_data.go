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

// AnnotationData - struct for AnnotationData
type AnnotationData struct {
	FeedbackData *FeedbackData
	GenericAnnotationData *GenericAnnotationData
	MetricStreamAnomalyData *MetricStreamAnomalyData
	MetricStreamNoAnomalyData *MetricStreamNoAnomalyData
}

// FeedbackDataAsAnnotationData is a convenience function that returns FeedbackData wrapped in AnnotationData
func FeedbackDataAsAnnotationData(v *FeedbackData) AnnotationData {
	return AnnotationData{ FeedbackData: v}
}

// GenericAnnotationDataAsAnnotationData is a convenience function that returns GenericAnnotationData wrapped in AnnotationData
func GenericAnnotationDataAsAnnotationData(v *GenericAnnotationData) AnnotationData {
	return AnnotationData{ GenericAnnotationData: v}
}

// MetricStreamAnomalyDataAsAnnotationData is a convenience function that returns MetricStreamAnomalyData wrapped in AnnotationData
func MetricStreamAnomalyDataAsAnnotationData(v *MetricStreamAnomalyData) AnnotationData {
	return AnnotationData{ MetricStreamAnomalyData: v}
}

// MetricStreamNoAnomalyDataAsAnnotationData is a convenience function that returns MetricStreamNoAnomalyData wrapped in AnnotationData
func MetricStreamNoAnomalyDataAsAnnotationData(v *MetricStreamNoAnomalyData) AnnotationData {
	return AnnotationData{ MetricStreamNoAnomalyData: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AnnotationData) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'FeedbackData'
	if jsonDict["_type"] == "FeedbackData" {
		// try to unmarshal JSON data into FeedbackData
		err = json.Unmarshal(data, &dst.FeedbackData)
		if err == nil {
			return nil // data stored in dst.FeedbackData, return on the first match
		} else {
			dst.FeedbackData = nil
			return fmt.Errorf("Failed to unmarshal AnnotationData as FeedbackData: %s", err.Error())
		}
	}

	// check if the discriminator value is 'GenericAnnotationData'
	if jsonDict["_type"] == "GenericAnnotationData" {
		// try to unmarshal JSON data into GenericAnnotationData
		err = json.Unmarshal(data, &dst.GenericAnnotationData)
		if err == nil {
			return nil // data stored in dst.GenericAnnotationData, return on the first match
		} else {
			dst.GenericAnnotationData = nil
			return fmt.Errorf("Failed to unmarshal AnnotationData as GenericAnnotationData: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MetricStreamAnomalyData'
	if jsonDict["_type"] == "MetricStreamAnomalyData" {
		// try to unmarshal JSON data into MetricStreamAnomalyData
		err = json.Unmarshal(data, &dst.MetricStreamAnomalyData)
		if err == nil {
			return nil // data stored in dst.MetricStreamAnomalyData, return on the first match
		} else {
			dst.MetricStreamAnomalyData = nil
			return fmt.Errorf("Failed to unmarshal AnnotationData as MetricStreamAnomalyData: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MetricStreamNoAnomalyData'
	if jsonDict["_type"] == "MetricStreamNoAnomalyData" {
		// try to unmarshal JSON data into MetricStreamNoAnomalyData
		err = json.Unmarshal(data, &dst.MetricStreamNoAnomalyData)
		if err == nil {
			return nil // data stored in dst.MetricStreamNoAnomalyData, return on the first match
		} else {
			dst.MetricStreamNoAnomalyData = nil
			return fmt.Errorf("Failed to unmarshal AnnotationData as MetricStreamNoAnomalyData: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AnnotationData) MarshalJSON() ([]byte, error) {
	if src.FeedbackData != nil {
		return json.Marshal(&src.FeedbackData)
	}

	if src.GenericAnnotationData != nil {
		return json.Marshal(&src.GenericAnnotationData)
	}

	if src.MetricStreamAnomalyData != nil {
		return json.Marshal(&src.MetricStreamAnomalyData)
	}

	if src.MetricStreamNoAnomalyData != nil {
		return json.Marshal(&src.MetricStreamNoAnomalyData)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AnnotationData) GetActualInstance() (interface{}) {
	if obj.FeedbackData != nil {
		return obj.FeedbackData
	}

	if obj.GenericAnnotationData != nil {
		return obj.GenericAnnotationData
	}

	if obj.MetricStreamAnomalyData != nil {
		return obj.MetricStreamAnomalyData
	}

	if obj.MetricStreamNoAnomalyData != nil {
		return obj.MetricStreamNoAnomalyData
	}

	// all schemas are nil
	return nil
}

type NullableAnnotationData struct {
	value *AnnotationData
	isSet bool
}

func (v NullableAnnotationData) Get() *AnnotationData {
	return v.value
}

func (v *NullableAnnotationData) Set(val *AnnotationData) {
	v.value = val
	v.isSet = true
}

func (v NullableAnnotationData) IsSet() bool {
	return v.isSet
}

func (v *NullableAnnotationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnnotationData(val *AnnotationData) *NullableAnnotationData {
	return &NullableAnnotationData{value: val, isSet: true}
}

func (v NullableAnnotationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnnotationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


