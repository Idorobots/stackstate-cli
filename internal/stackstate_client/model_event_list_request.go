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

// EventListRequest struct for EventListRequest
type EventListRequest struct {
	StartTimestampMs int32 `json:"startTimestampMs"`
	EndTimestampMs int32 `json:"endTimestampMs"`
	TopologyQuery string `json:"topologyQuery"`
	Limit int32 `json:"limit"`
	PlayHeadTimestampMs *int32 `json:"playHeadTimestampMs,omitempty"`
	RootCauseMode *RootCauseMode `json:"rootCauseMode,omitempty"`
	EventTypes []string `json:"eventTypes,omitempty"`
	EventTags []string `json:"eventTags,omitempty"`
	EventCategories []EventCategory `json:"eventCategories,omitempty"`
	EventSources []string `json:"eventSources,omitempty"`
	Cursor *EventCursor `json:"cursor,omitempty"`
}

// NewEventListRequest instantiates a new EventListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventListRequest(startTimestampMs int32, endTimestampMs int32, topologyQuery string, limit int32) *EventListRequest {
	this := EventListRequest{}
	this.StartTimestampMs = startTimestampMs
	this.EndTimestampMs = endTimestampMs
	this.TopologyQuery = topologyQuery
	this.Limit = limit
	return &this
}

// NewEventListRequestWithDefaults instantiates a new EventListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventListRequestWithDefaults() *EventListRequest {
	this := EventListRequest{}
	return &this
}

// GetStartTimestampMs returns the StartTimestampMs field value
func (o *EventListRequest) GetStartTimestampMs() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StartTimestampMs
}

// GetStartTimestampMsOk returns a tuple with the StartTimestampMs field value
// and a boolean to check if the value has been set.
func (o *EventListRequest) GetStartTimestampMsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTimestampMs, true
}

// SetStartTimestampMs sets field value
func (o *EventListRequest) SetStartTimestampMs(v int32) {
	o.StartTimestampMs = v
}

// GetEndTimestampMs returns the EndTimestampMs field value
func (o *EventListRequest) GetEndTimestampMs() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EndTimestampMs
}

// GetEndTimestampMsOk returns a tuple with the EndTimestampMs field value
// and a boolean to check if the value has been set.
func (o *EventListRequest) GetEndTimestampMsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTimestampMs, true
}

// SetEndTimestampMs sets field value
func (o *EventListRequest) SetEndTimestampMs(v int32) {
	o.EndTimestampMs = v
}

// GetTopologyQuery returns the TopologyQuery field value
func (o *EventListRequest) GetTopologyQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TopologyQuery
}

// GetTopologyQueryOk returns a tuple with the TopologyQuery field value
// and a boolean to check if the value has been set.
func (o *EventListRequest) GetTopologyQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopologyQuery, true
}

// SetTopologyQuery sets field value
func (o *EventListRequest) SetTopologyQuery(v string) {
	o.TopologyQuery = v
}

// GetLimit returns the Limit field value
func (o *EventListRequest) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *EventListRequest) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *EventListRequest) SetLimit(v int32) {
	o.Limit = v
}

// GetPlayHeadTimestampMs returns the PlayHeadTimestampMs field value if set, zero value otherwise.
func (o *EventListRequest) GetPlayHeadTimestampMs() int32 {
	if o == nil || o.PlayHeadTimestampMs == nil {
		var ret int32
		return ret
	}
	return *o.PlayHeadTimestampMs
}

// GetPlayHeadTimestampMsOk returns a tuple with the PlayHeadTimestampMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventListRequest) GetPlayHeadTimestampMsOk() (*int32, bool) {
	if o == nil || o.PlayHeadTimestampMs == nil {
		return nil, false
	}
	return o.PlayHeadTimestampMs, true
}

// HasPlayHeadTimestampMs returns a boolean if a field has been set.
func (o *EventListRequest) HasPlayHeadTimestampMs() bool {
	if o != nil && o.PlayHeadTimestampMs != nil {
		return true
	}

	return false
}

// SetPlayHeadTimestampMs gets a reference to the given int32 and assigns it to the PlayHeadTimestampMs field.
func (o *EventListRequest) SetPlayHeadTimestampMs(v int32) {
	o.PlayHeadTimestampMs = &v
}

// GetRootCauseMode returns the RootCauseMode field value if set, zero value otherwise.
func (o *EventListRequest) GetRootCauseMode() RootCauseMode {
	if o == nil || o.RootCauseMode == nil {
		var ret RootCauseMode
		return ret
	}
	return *o.RootCauseMode
}

// GetRootCauseModeOk returns a tuple with the RootCauseMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventListRequest) GetRootCauseModeOk() (*RootCauseMode, bool) {
	if o == nil || o.RootCauseMode == nil {
		return nil, false
	}
	return o.RootCauseMode, true
}

// HasRootCauseMode returns a boolean if a field has been set.
func (o *EventListRequest) HasRootCauseMode() bool {
	if o != nil && o.RootCauseMode != nil {
		return true
	}

	return false
}

// SetRootCauseMode gets a reference to the given RootCauseMode and assigns it to the RootCauseMode field.
func (o *EventListRequest) SetRootCauseMode(v RootCauseMode) {
	o.RootCauseMode = &v
}

// GetEventTypes returns the EventTypes field value if set, zero value otherwise.
func (o *EventListRequest) GetEventTypes() []string {
	if o == nil || o.EventTypes == nil {
		var ret []string
		return ret
	}
	return o.EventTypes
}

// GetEventTypesOk returns a tuple with the EventTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventListRequest) GetEventTypesOk() ([]string, bool) {
	if o == nil || o.EventTypes == nil {
		return nil, false
	}
	return o.EventTypes, true
}

// HasEventTypes returns a boolean if a field has been set.
func (o *EventListRequest) HasEventTypes() bool {
	if o != nil && o.EventTypes != nil {
		return true
	}

	return false
}

// SetEventTypes gets a reference to the given []string and assigns it to the EventTypes field.
func (o *EventListRequest) SetEventTypes(v []string) {
	o.EventTypes = v
}

// GetEventTags returns the EventTags field value if set, zero value otherwise.
func (o *EventListRequest) GetEventTags() []string {
	if o == nil || o.EventTags == nil {
		var ret []string
		return ret
	}
	return o.EventTags
}

// GetEventTagsOk returns a tuple with the EventTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventListRequest) GetEventTagsOk() ([]string, bool) {
	if o == nil || o.EventTags == nil {
		return nil, false
	}
	return o.EventTags, true
}

// HasEventTags returns a boolean if a field has been set.
func (o *EventListRequest) HasEventTags() bool {
	if o != nil && o.EventTags != nil {
		return true
	}

	return false
}

// SetEventTags gets a reference to the given []string and assigns it to the EventTags field.
func (o *EventListRequest) SetEventTags(v []string) {
	o.EventTags = v
}

// GetEventCategories returns the EventCategories field value if set, zero value otherwise.
func (o *EventListRequest) GetEventCategories() []EventCategory {
	if o == nil || o.EventCategories == nil {
		var ret []EventCategory
		return ret
	}
	return o.EventCategories
}

// GetEventCategoriesOk returns a tuple with the EventCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventListRequest) GetEventCategoriesOk() ([]EventCategory, bool) {
	if o == nil || o.EventCategories == nil {
		return nil, false
	}
	return o.EventCategories, true
}

// HasEventCategories returns a boolean if a field has been set.
func (o *EventListRequest) HasEventCategories() bool {
	if o != nil && o.EventCategories != nil {
		return true
	}

	return false
}

// SetEventCategories gets a reference to the given []EventCategory and assigns it to the EventCategories field.
func (o *EventListRequest) SetEventCategories(v []EventCategory) {
	o.EventCategories = v
}

// GetEventSources returns the EventSources field value if set, zero value otherwise.
func (o *EventListRequest) GetEventSources() []string {
	if o == nil || o.EventSources == nil {
		var ret []string
		return ret
	}
	return o.EventSources
}

// GetEventSourcesOk returns a tuple with the EventSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventListRequest) GetEventSourcesOk() ([]string, bool) {
	if o == nil || o.EventSources == nil {
		return nil, false
	}
	return o.EventSources, true
}

// HasEventSources returns a boolean if a field has been set.
func (o *EventListRequest) HasEventSources() bool {
	if o != nil && o.EventSources != nil {
		return true
	}

	return false
}

// SetEventSources gets a reference to the given []string and assigns it to the EventSources field.
func (o *EventListRequest) SetEventSources(v []string) {
	o.EventSources = v
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *EventListRequest) GetCursor() EventCursor {
	if o == nil || o.Cursor == nil {
		var ret EventCursor
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventListRequest) GetCursorOk() (*EventCursor, bool) {
	if o == nil || o.Cursor == nil {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *EventListRequest) HasCursor() bool {
	if o != nil && o.Cursor != nil {
		return true
	}

	return false
}

// SetCursor gets a reference to the given EventCursor and assigns it to the Cursor field.
func (o *EventListRequest) SetCursor(v EventCursor) {
	o.Cursor = &v
}

func (o EventListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["startTimestampMs"] = o.StartTimestampMs
	}
	if true {
		toSerialize["endTimestampMs"] = o.EndTimestampMs
	}
	if true {
		toSerialize["topologyQuery"] = o.TopologyQuery
	}
	if true {
		toSerialize["limit"] = o.Limit
	}
	if o.PlayHeadTimestampMs != nil {
		toSerialize["playHeadTimestampMs"] = o.PlayHeadTimestampMs
	}
	if o.RootCauseMode != nil {
		toSerialize["rootCauseMode"] = o.RootCauseMode
	}
	if o.EventTypes != nil {
		toSerialize["eventTypes"] = o.EventTypes
	}
	if o.EventTags != nil {
		toSerialize["eventTags"] = o.EventTags
	}
	if o.EventCategories != nil {
		toSerialize["eventCategories"] = o.EventCategories
	}
	if o.EventSources != nil {
		toSerialize["eventSources"] = o.EventSources
	}
	if o.Cursor != nil {
		toSerialize["cursor"] = o.Cursor
	}
	return json.Marshal(toSerialize)
}

type NullableEventListRequest struct {
	value *EventListRequest
	isSet bool
}

func (v NullableEventListRequest) Get() *EventListRequest {
	return v.value
}

func (v *NullableEventListRequest) Set(val *EventListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEventListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEventListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventListRequest(val *EventListRequest) *NullableEventListRequest {
	return &NullableEventListRequest{value: val, isSet: true}
}

func (v NullableEventListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


