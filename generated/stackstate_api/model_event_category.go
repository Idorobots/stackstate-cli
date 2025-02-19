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

// EventCategory the model 'EventCategory'
type EventCategory string

// List of EventCategory
const (
	EVENTCATEGORY_CHANGES    EventCategory = "Changes"
	EVENTCATEGORY_ALERTS     EventCategory = "Alerts"
	EVENTCATEGORY_ANOMALIES  EventCategory = "Anomalies"
	EVENTCATEGORY_ACTIVITIES EventCategory = "Activities"
	EVENTCATEGORY_OTHERS     EventCategory = "Others"
)

// All allowed values of EventCategory enum
var AllowedEventCategoryEnumValues = []EventCategory{
	"Changes",
	"Alerts",
	"Anomalies",
	"Activities",
	"Others",
}

func (v *EventCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventCategory(value)
	for _, existing := range AllowedEventCategoryEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventCategory", value)
}

// NewEventCategoryFromValue returns a pointer to a valid EventCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventCategoryFromValue(v string) (*EventCategory, error) {
	ev := EventCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventCategory: valid values are %v", v, AllowedEventCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventCategory) IsValid() bool {
	for _, existing := range AllowedEventCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventCategory value
func (v EventCategory) Ptr() *EventCategory {
	return &v
}

type NullableEventCategory struct {
	value *EventCategory
	isSet bool
}

func (v NullableEventCategory) Get() *EventCategory {
	return v.value
}

func (v *NullableEventCategory) Set(val *EventCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableEventCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableEventCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventCategory(val *EventCategory) *NullableEventCategory {
	return &NullableEventCategory{value: val, isSet: true}
}

func (v NullableEventCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
