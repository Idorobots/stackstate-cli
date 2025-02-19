/*
StackState API

This API documentation page describes the StackState server API. The StackState UI and CLI use the StackState server API to configure and query StackState.  You can use the API for similar purposes.  Each request sent to the StackState server API must be authenticated using one of the available authentication methods.   *Note that the StackState receiver API, used to send topology, telemetry and traces to StackState, is not described on this page and requires a different authentication method.*  For more information on StackState, refer to the [StackState documentation](https://docs.stackstate.com).

API version: 5.2.0
Contact: info@stackstate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackstate_api

import (
	"encoding/json"
	"fmt"
)

// Argument - struct for Argument
type Argument struct {
	ArgumentAnomalyDirectionVal          *ArgumentAnomalyDirectionVal
	ArgumentBaselineMetricStreamRef      *ArgumentBaselineMetricStreamRef
	ArgumentBooleanVal                   *ArgumentBooleanVal
	ArgumentComponentTypeRef             *ArgumentComponentTypeRef
	ArgumentDoubleVal                    *ArgumentDoubleVal
	ArgumentDurationVal                  *ArgumentDurationVal
	ArgumentEventStreamRef               *ArgumentEventStreamRef
	ArgumentLongVal                      *ArgumentLongVal
	ArgumentMetricStreamId               *ArgumentMetricStreamId
	ArgumentMetricStreamRef              *ArgumentMetricStreamRef
	ArgumentNodeIdVal                    *ArgumentNodeIdVal
	ArgumentPropagatedHealthStateVal     *ArgumentPropagatedHealthStateVal
	ArgumentQueryViewRef                 *ArgumentQueryViewRef
	ArgumentRelationTypeRef              *ArgumentRelationTypeRef
	ArgumentRunStateVal                  *ArgumentRunStateVal
	ArgumentScriptMetricQueryVal         *ArgumentScriptMetricQueryVal
	ArgumentSimpleTrainingPeriodicityVal *ArgumentSimpleTrainingPeriodicityVal
	ArgumentStateVal                     *ArgumentStateVal
	ArgumentStringVal                    *ArgumentStringVal
	ArgumentStructTypeVal                *ArgumentStructTypeVal
	ArgumentStsEventStreamVal            *ArgumentStsEventStreamVal
}

// ArgumentAnomalyDirectionValAsArgument is a convenience function that returns ArgumentAnomalyDirectionVal wrapped in Argument
func ArgumentAnomalyDirectionValAsArgument(v *ArgumentAnomalyDirectionVal) Argument {
	return Argument{
		ArgumentAnomalyDirectionVal: v,
	}
}

// ArgumentBaselineMetricStreamRefAsArgument is a convenience function that returns ArgumentBaselineMetricStreamRef wrapped in Argument
func ArgumentBaselineMetricStreamRefAsArgument(v *ArgumentBaselineMetricStreamRef) Argument {
	return Argument{
		ArgumentBaselineMetricStreamRef: v,
	}
}

// ArgumentBooleanValAsArgument is a convenience function that returns ArgumentBooleanVal wrapped in Argument
func ArgumentBooleanValAsArgument(v *ArgumentBooleanVal) Argument {
	return Argument{
		ArgumentBooleanVal: v,
	}
}

// ArgumentComponentTypeRefAsArgument is a convenience function that returns ArgumentComponentTypeRef wrapped in Argument
func ArgumentComponentTypeRefAsArgument(v *ArgumentComponentTypeRef) Argument {
	return Argument{
		ArgumentComponentTypeRef: v,
	}
}

// ArgumentDoubleValAsArgument is a convenience function that returns ArgumentDoubleVal wrapped in Argument
func ArgumentDoubleValAsArgument(v *ArgumentDoubleVal) Argument {
	return Argument{
		ArgumentDoubleVal: v,
	}
}

// ArgumentDurationValAsArgument is a convenience function that returns ArgumentDurationVal wrapped in Argument
func ArgumentDurationValAsArgument(v *ArgumentDurationVal) Argument {
	return Argument{
		ArgumentDurationVal: v,
	}
}

// ArgumentEventStreamRefAsArgument is a convenience function that returns ArgumentEventStreamRef wrapped in Argument
func ArgumentEventStreamRefAsArgument(v *ArgumentEventStreamRef) Argument {
	return Argument{
		ArgumentEventStreamRef: v,
	}
}

// ArgumentLongValAsArgument is a convenience function that returns ArgumentLongVal wrapped in Argument
func ArgumentLongValAsArgument(v *ArgumentLongVal) Argument {
	return Argument{
		ArgumentLongVal: v,
	}
}

// ArgumentMetricStreamIdAsArgument is a convenience function that returns ArgumentMetricStreamId wrapped in Argument
func ArgumentMetricStreamIdAsArgument(v *ArgumentMetricStreamId) Argument {
	return Argument{
		ArgumentMetricStreamId: v,
	}
}

// ArgumentMetricStreamRefAsArgument is a convenience function that returns ArgumentMetricStreamRef wrapped in Argument
func ArgumentMetricStreamRefAsArgument(v *ArgumentMetricStreamRef) Argument {
	return Argument{
		ArgumentMetricStreamRef: v,
	}
}

// ArgumentNodeIdValAsArgument is a convenience function that returns ArgumentNodeIdVal wrapped in Argument
func ArgumentNodeIdValAsArgument(v *ArgumentNodeIdVal) Argument {
	return Argument{
		ArgumentNodeIdVal: v,
	}
}

// ArgumentPropagatedHealthStateValAsArgument is a convenience function that returns ArgumentPropagatedHealthStateVal wrapped in Argument
func ArgumentPropagatedHealthStateValAsArgument(v *ArgumentPropagatedHealthStateVal) Argument {
	return Argument{
		ArgumentPropagatedHealthStateVal: v,
	}
}

// ArgumentQueryViewRefAsArgument is a convenience function that returns ArgumentQueryViewRef wrapped in Argument
func ArgumentQueryViewRefAsArgument(v *ArgumentQueryViewRef) Argument {
	return Argument{
		ArgumentQueryViewRef: v,
	}
}

// ArgumentRelationTypeRefAsArgument is a convenience function that returns ArgumentRelationTypeRef wrapped in Argument
func ArgumentRelationTypeRefAsArgument(v *ArgumentRelationTypeRef) Argument {
	return Argument{
		ArgumentRelationTypeRef: v,
	}
}

// ArgumentRunStateValAsArgument is a convenience function that returns ArgumentRunStateVal wrapped in Argument
func ArgumentRunStateValAsArgument(v *ArgumentRunStateVal) Argument {
	return Argument{
		ArgumentRunStateVal: v,
	}
}

// ArgumentScriptMetricQueryValAsArgument is a convenience function that returns ArgumentScriptMetricQueryVal wrapped in Argument
func ArgumentScriptMetricQueryValAsArgument(v *ArgumentScriptMetricQueryVal) Argument {
	return Argument{
		ArgumentScriptMetricQueryVal: v,
	}
}

// ArgumentSimpleTrainingPeriodicityValAsArgument is a convenience function that returns ArgumentSimpleTrainingPeriodicityVal wrapped in Argument
func ArgumentSimpleTrainingPeriodicityValAsArgument(v *ArgumentSimpleTrainingPeriodicityVal) Argument {
	return Argument{
		ArgumentSimpleTrainingPeriodicityVal: v,
	}
}

// ArgumentStateValAsArgument is a convenience function that returns ArgumentStateVal wrapped in Argument
func ArgumentStateValAsArgument(v *ArgumentStateVal) Argument {
	return Argument{
		ArgumentStateVal: v,
	}
}

// ArgumentStringValAsArgument is a convenience function that returns ArgumentStringVal wrapped in Argument
func ArgumentStringValAsArgument(v *ArgumentStringVal) Argument {
	return Argument{
		ArgumentStringVal: v,
	}
}

// ArgumentStructTypeValAsArgument is a convenience function that returns ArgumentStructTypeVal wrapped in Argument
func ArgumentStructTypeValAsArgument(v *ArgumentStructTypeVal) Argument {
	return Argument{
		ArgumentStructTypeVal: v,
	}
}

// ArgumentStsEventStreamValAsArgument is a convenience function that returns ArgumentStsEventStreamVal wrapped in Argument
func ArgumentStsEventStreamValAsArgument(v *ArgumentStsEventStreamVal) Argument {
	return Argument{
		ArgumentStsEventStreamVal: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *Argument) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'ArgumentAnomalyDirectionVal'
	if jsonDict["_type"] == "ArgumentAnomalyDirectionVal" {
		// try to unmarshal JSON data into ArgumentAnomalyDirectionVal
		err = json.Unmarshal(data, &dst.ArgumentAnomalyDirectionVal)
		if err == nil {
			return nil // data stored in dst.ArgumentAnomalyDirectionVal, return on the first match
		} else {
			dst.ArgumentAnomalyDirectionVal = nil
			return fmt.Errorf("Failed to unmarshal Argument as ArgumentAnomalyDirectionVal: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ArgumentBaselineMetricStreamRef'
	if jsonDict["_type"] == "ArgumentBaselineMetricStreamRef" {
		// try to unmarshal JSON data into ArgumentBaselineMetricStreamRef
		err = json.Unmarshal(data, &dst.ArgumentBaselineMetricStreamRef)
		if err == nil {
			return nil // data stored in dst.ArgumentBaselineMetricStreamRef, return on the first match
		} else {
			dst.ArgumentBaselineMetricStreamRef = nil
			return fmt.Errorf("Failed to unmarshal Argument as ArgumentBaselineMetricStreamRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ArgumentBooleanVal'
	if jsonDict["_type"] == "ArgumentBooleanVal" {
		// try to unmarshal JSON data into ArgumentBooleanVal
		err = json.Unmarshal(data, &dst.ArgumentBooleanVal)
		if err == nil {
			return nil // data stored in dst.ArgumentBooleanVal, return on the first match
		} else {
			dst.ArgumentBooleanVal = nil
			return fmt.Errorf("Failed to unmarshal Argument as ArgumentBooleanVal: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ArgumentComponentTypeRef'
	if jsonDict["_type"] == "ArgumentComponentTypeRef" {
		// try to unmarshal JSON data into ArgumentComponentTypeRef
		err = json.Unmarshal(data, &dst.ArgumentComponentTypeRef)
		if err == nil {
			return nil // data stored in dst.ArgumentComponentTypeRef, return on the first match
		} else {
			dst.ArgumentComponentTypeRef = nil
			return fmt.Errorf("Failed to unmarshal Argument as ArgumentComponentTypeRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ArgumentDoubleVal'
	if jsonDict["_type"] == "ArgumentDoubleVal" {
		// try to unmarshal JSON data into ArgumentDoubleVal
		err = json.Unmarshal(data, &dst.ArgumentDoubleVal)
		if err == nil {
			return nil // data stored in dst.ArgumentDoubleVal, return on the first match
		} else {
			dst.ArgumentDoubleVal = nil
			return fmt.Errorf("Failed to unmarshal Argument as ArgumentDoubleVal: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ArgumentDurationVal'
	if jsonDict["_type"] == "ArgumentDurationVal" {
		// try to unmarshal JSON data into ArgumentDurationVal
		err = json.Unmarshal(data, &dst.ArgumentDurationVal)
		if err == nil {
			return nil // data stored in dst.ArgumentDurationVal, return on the first match
		} else {
			dst.ArgumentDurationVal = nil
			return fmt.Errorf("Failed to unmarshal Argument as ArgumentDurationVal: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ArgumentEventStreamRef'
	if jsonDict["_type"] == "ArgumentEventStreamRef" {
		// try to unmarshal JSON data into ArgumentEventStreamRef
		err = json.Unmarshal(data, &dst.ArgumentEventStreamRef)
		if err == nil {
			return nil // data stored in dst.ArgumentEventStreamRef, return on the first match
		} else {
			dst.ArgumentEventStreamRef = nil
			return fmt.Errorf("Failed to unmarshal Argument as ArgumentEventStreamRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ArgumentLongVal'
	if jsonDict["_type"] == "ArgumentLongVal" {
		// try to unmarshal JSON data into ArgumentLongVal
		err = json.Unmarshal(data, &dst.ArgumentLongVal)
		if err == nil {
			return nil // data stored in dst.ArgumentLongVal, return on the first match
		} else {
			dst.ArgumentLongVal = nil
			return fmt.Errorf("Failed to unmarshal Argument as ArgumentLongVal: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ArgumentMetricStreamId'
	if jsonDict["_type"] == "ArgumentMetricStreamId" {
		// try to unmarshal JSON data into ArgumentMetricStreamId
		err = json.Unmarshal(data, &dst.ArgumentMetricStreamId)
		if err == nil {
			return nil // data stored in dst.ArgumentMetricStreamId, return on the first match
		} else {
			dst.ArgumentMetricStreamId = nil
			return fmt.Errorf("Failed to unmarshal Argument as ArgumentMetricStreamId: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ArgumentMetricStreamRef'
	if jsonDict["_type"] == "ArgumentMetricStreamRef" {
		// try to unmarshal JSON data into ArgumentMetricStreamRef
		err = json.Unmarshal(data, &dst.ArgumentMetricStreamRef)
		if err == nil {
			return nil // data stored in dst.ArgumentMetricStreamRef, return on the first match
		} else {
			dst.ArgumentMetricStreamRef = nil
			return fmt.Errorf("Failed to unmarshal Argument as ArgumentMetricStreamRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ArgumentNodeIdVal'
	if jsonDict["_type"] == "ArgumentNodeIdVal" {
		// try to unmarshal JSON data into ArgumentNodeIdVal
		err = json.Unmarshal(data, &dst.ArgumentNodeIdVal)
		if err == nil {
			return nil // data stored in dst.ArgumentNodeIdVal, return on the first match
		} else {
			dst.ArgumentNodeIdVal = nil
			return fmt.Errorf("Failed to unmarshal Argument as ArgumentNodeIdVal: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ArgumentPropagatedHealthStateVal'
	if jsonDict["_type"] == "ArgumentPropagatedHealthStateVal" {
		// try to unmarshal JSON data into ArgumentPropagatedHealthStateVal
		err = json.Unmarshal(data, &dst.ArgumentPropagatedHealthStateVal)
		if err == nil {
			return nil // data stored in dst.ArgumentPropagatedHealthStateVal, return on the first match
		} else {
			dst.ArgumentPropagatedHealthStateVal = nil
			return fmt.Errorf("Failed to unmarshal Argument as ArgumentPropagatedHealthStateVal: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ArgumentQueryViewRef'
	if jsonDict["_type"] == "ArgumentQueryViewRef" {
		// try to unmarshal JSON data into ArgumentQueryViewRef
		err = json.Unmarshal(data, &dst.ArgumentQueryViewRef)
		if err == nil {
			return nil // data stored in dst.ArgumentQueryViewRef, return on the first match
		} else {
			dst.ArgumentQueryViewRef = nil
			return fmt.Errorf("Failed to unmarshal Argument as ArgumentQueryViewRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ArgumentRelationTypeRef'
	if jsonDict["_type"] == "ArgumentRelationTypeRef" {
		// try to unmarshal JSON data into ArgumentRelationTypeRef
		err = json.Unmarshal(data, &dst.ArgumentRelationTypeRef)
		if err == nil {
			return nil // data stored in dst.ArgumentRelationTypeRef, return on the first match
		} else {
			dst.ArgumentRelationTypeRef = nil
			return fmt.Errorf("Failed to unmarshal Argument as ArgumentRelationTypeRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ArgumentRunStateVal'
	if jsonDict["_type"] == "ArgumentRunStateVal" {
		// try to unmarshal JSON data into ArgumentRunStateVal
		err = json.Unmarshal(data, &dst.ArgumentRunStateVal)
		if err == nil {
			return nil // data stored in dst.ArgumentRunStateVal, return on the first match
		} else {
			dst.ArgumentRunStateVal = nil
			return fmt.Errorf("Failed to unmarshal Argument as ArgumentRunStateVal: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ArgumentScriptMetricQueryVal'
	if jsonDict["_type"] == "ArgumentScriptMetricQueryVal" {
		// try to unmarshal JSON data into ArgumentScriptMetricQueryVal
		err = json.Unmarshal(data, &dst.ArgumentScriptMetricQueryVal)
		if err == nil {
			return nil // data stored in dst.ArgumentScriptMetricQueryVal, return on the first match
		} else {
			dst.ArgumentScriptMetricQueryVal = nil
			return fmt.Errorf("Failed to unmarshal Argument as ArgumentScriptMetricQueryVal: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ArgumentSimpleTrainingPeriodicityVal'
	if jsonDict["_type"] == "ArgumentSimpleTrainingPeriodicityVal" {
		// try to unmarshal JSON data into ArgumentSimpleTrainingPeriodicityVal
		err = json.Unmarshal(data, &dst.ArgumentSimpleTrainingPeriodicityVal)
		if err == nil {
			return nil // data stored in dst.ArgumentSimpleTrainingPeriodicityVal, return on the first match
		} else {
			dst.ArgumentSimpleTrainingPeriodicityVal = nil
			return fmt.Errorf("Failed to unmarshal Argument as ArgumentSimpleTrainingPeriodicityVal: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ArgumentStateVal'
	if jsonDict["_type"] == "ArgumentStateVal" {
		// try to unmarshal JSON data into ArgumentStateVal
		err = json.Unmarshal(data, &dst.ArgumentStateVal)
		if err == nil {
			return nil // data stored in dst.ArgumentStateVal, return on the first match
		} else {
			dst.ArgumentStateVal = nil
			return fmt.Errorf("Failed to unmarshal Argument as ArgumentStateVal: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ArgumentStringVal'
	if jsonDict["_type"] == "ArgumentStringVal" {
		// try to unmarshal JSON data into ArgumentStringVal
		err = json.Unmarshal(data, &dst.ArgumentStringVal)
		if err == nil {
			return nil // data stored in dst.ArgumentStringVal, return on the first match
		} else {
			dst.ArgumentStringVal = nil
			return fmt.Errorf("Failed to unmarshal Argument as ArgumentStringVal: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ArgumentStructTypeVal'
	if jsonDict["_type"] == "ArgumentStructTypeVal" {
		// try to unmarshal JSON data into ArgumentStructTypeVal
		err = json.Unmarshal(data, &dst.ArgumentStructTypeVal)
		if err == nil {
			return nil // data stored in dst.ArgumentStructTypeVal, return on the first match
		} else {
			dst.ArgumentStructTypeVal = nil
			return fmt.Errorf("Failed to unmarshal Argument as ArgumentStructTypeVal: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ArgumentStsEventStreamVal'
	if jsonDict["_type"] == "ArgumentStsEventStreamVal" {
		// try to unmarshal JSON data into ArgumentStsEventStreamVal
		err = json.Unmarshal(data, &dst.ArgumentStsEventStreamVal)
		if err == nil {
			return nil // data stored in dst.ArgumentStsEventStreamVal, return on the first match
		} else {
			dst.ArgumentStsEventStreamVal = nil
			return fmt.Errorf("Failed to unmarshal Argument as ArgumentStsEventStreamVal: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Argument) MarshalJSON() ([]byte, error) {
	if src.ArgumentAnomalyDirectionVal != nil {
		return json.Marshal(&src.ArgumentAnomalyDirectionVal)
	}

	if src.ArgumentBaselineMetricStreamRef != nil {
		return json.Marshal(&src.ArgumentBaselineMetricStreamRef)
	}

	if src.ArgumentBooleanVal != nil {
		return json.Marshal(&src.ArgumentBooleanVal)
	}

	if src.ArgumentComponentTypeRef != nil {
		return json.Marshal(&src.ArgumentComponentTypeRef)
	}

	if src.ArgumentDoubleVal != nil {
		return json.Marshal(&src.ArgumentDoubleVal)
	}

	if src.ArgumentDurationVal != nil {
		return json.Marshal(&src.ArgumentDurationVal)
	}

	if src.ArgumentEventStreamRef != nil {
		return json.Marshal(&src.ArgumentEventStreamRef)
	}

	if src.ArgumentLongVal != nil {
		return json.Marshal(&src.ArgumentLongVal)
	}

	if src.ArgumentMetricStreamId != nil {
		return json.Marshal(&src.ArgumentMetricStreamId)
	}

	if src.ArgumentMetricStreamRef != nil {
		return json.Marshal(&src.ArgumentMetricStreamRef)
	}

	if src.ArgumentNodeIdVal != nil {
		return json.Marshal(&src.ArgumentNodeIdVal)
	}

	if src.ArgumentPropagatedHealthStateVal != nil {
		return json.Marshal(&src.ArgumentPropagatedHealthStateVal)
	}

	if src.ArgumentQueryViewRef != nil {
		return json.Marshal(&src.ArgumentQueryViewRef)
	}

	if src.ArgumentRelationTypeRef != nil {
		return json.Marshal(&src.ArgumentRelationTypeRef)
	}

	if src.ArgumentRunStateVal != nil {
		return json.Marshal(&src.ArgumentRunStateVal)
	}

	if src.ArgumentScriptMetricQueryVal != nil {
		return json.Marshal(&src.ArgumentScriptMetricQueryVal)
	}

	if src.ArgumentSimpleTrainingPeriodicityVal != nil {
		return json.Marshal(&src.ArgumentSimpleTrainingPeriodicityVal)
	}

	if src.ArgumentStateVal != nil {
		return json.Marshal(&src.ArgumentStateVal)
	}

	if src.ArgumentStringVal != nil {
		return json.Marshal(&src.ArgumentStringVal)
	}

	if src.ArgumentStructTypeVal != nil {
		return json.Marshal(&src.ArgumentStructTypeVal)
	}

	if src.ArgumentStsEventStreamVal != nil {
		return json.Marshal(&src.ArgumentStsEventStreamVal)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Argument) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ArgumentAnomalyDirectionVal != nil {
		return obj.ArgumentAnomalyDirectionVal
	}

	if obj.ArgumentBaselineMetricStreamRef != nil {
		return obj.ArgumentBaselineMetricStreamRef
	}

	if obj.ArgumentBooleanVal != nil {
		return obj.ArgumentBooleanVal
	}

	if obj.ArgumentComponentTypeRef != nil {
		return obj.ArgumentComponentTypeRef
	}

	if obj.ArgumentDoubleVal != nil {
		return obj.ArgumentDoubleVal
	}

	if obj.ArgumentDurationVal != nil {
		return obj.ArgumentDurationVal
	}

	if obj.ArgumentEventStreamRef != nil {
		return obj.ArgumentEventStreamRef
	}

	if obj.ArgumentLongVal != nil {
		return obj.ArgumentLongVal
	}

	if obj.ArgumentMetricStreamId != nil {
		return obj.ArgumentMetricStreamId
	}

	if obj.ArgumentMetricStreamRef != nil {
		return obj.ArgumentMetricStreamRef
	}

	if obj.ArgumentNodeIdVal != nil {
		return obj.ArgumentNodeIdVal
	}

	if obj.ArgumentPropagatedHealthStateVal != nil {
		return obj.ArgumentPropagatedHealthStateVal
	}

	if obj.ArgumentQueryViewRef != nil {
		return obj.ArgumentQueryViewRef
	}

	if obj.ArgumentRelationTypeRef != nil {
		return obj.ArgumentRelationTypeRef
	}

	if obj.ArgumentRunStateVal != nil {
		return obj.ArgumentRunStateVal
	}

	if obj.ArgumentScriptMetricQueryVal != nil {
		return obj.ArgumentScriptMetricQueryVal
	}

	if obj.ArgumentSimpleTrainingPeriodicityVal != nil {
		return obj.ArgumentSimpleTrainingPeriodicityVal
	}

	if obj.ArgumentStateVal != nil {
		return obj.ArgumentStateVal
	}

	if obj.ArgumentStringVal != nil {
		return obj.ArgumentStringVal
	}

	if obj.ArgumentStructTypeVal != nil {
		return obj.ArgumentStructTypeVal
	}

	if obj.ArgumentStsEventStreamVal != nil {
		return obj.ArgumentStsEventStreamVal
	}

	// all schemas are nil
	return nil
}

type NullableArgument struct {
	value *Argument
	isSet bool
}

func (v NullableArgument) Get() *Argument {
	return v.value
}

func (v *NullableArgument) Set(val *Argument) {
	v.value = val
	v.isSet = true
}

func (v NullableArgument) IsSet() bool {
	return v.isSet
}

func (v *NullableArgument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArgument(val *Argument) *NullableArgument {
	return &NullableArgument{value: val, isSet: true}
}

func (v NullableArgument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArgument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
