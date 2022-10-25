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
)

// AnomalyWithContext struct for AnomalyWithContext
type AnomalyWithContext struct {
	Anomaly  Annotation    `json:"anomaly"`
	Data     []Point       `json:"data"`
	Feedback *FeedbackData `json:"feedback,omitempty"`
}

// NewAnomalyWithContext instantiates a new AnomalyWithContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnomalyWithContext(anomaly Annotation, data []Point) *AnomalyWithContext {
	this := AnomalyWithContext{}
	this.Anomaly = anomaly
	this.Data = data
	return &this
}

// NewAnomalyWithContextWithDefaults instantiates a new AnomalyWithContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnomalyWithContextWithDefaults() *AnomalyWithContext {
	this := AnomalyWithContext{}
	return &this
}

// GetAnomaly returns the Anomaly field value
func (o *AnomalyWithContext) GetAnomaly() Annotation {
	if o == nil {
		var ret Annotation
		return ret
	}

	return o.Anomaly
}

// GetAnomalyOk returns a tuple with the Anomaly field value
// and a boolean to check if the value has been set.
func (o *AnomalyWithContext) GetAnomalyOk() (*Annotation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Anomaly, true
}

// SetAnomaly sets field value
func (o *AnomalyWithContext) SetAnomaly(v Annotation) {
	o.Anomaly = v
}

// GetData returns the Data field value
func (o *AnomalyWithContext) GetData() []Point {
	if o == nil {
		var ret []Point
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AnomalyWithContext) GetDataOk() ([]Point, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AnomalyWithContext) SetData(v []Point) {
	o.Data = v
}

// GetFeedback returns the Feedback field value if set, zero value otherwise.
func (o *AnomalyWithContext) GetFeedback() FeedbackData {
	if o == nil || o.Feedback == nil {
		var ret FeedbackData
		return ret
	}
	return *o.Feedback
}

// GetFeedbackOk returns a tuple with the Feedback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnomalyWithContext) GetFeedbackOk() (*FeedbackData, bool) {
	if o == nil || o.Feedback == nil {
		return nil, false
	}
	return o.Feedback, true
}

// HasFeedback returns a boolean if a field has been set.
func (o *AnomalyWithContext) HasFeedback() bool {
	if o != nil && o.Feedback != nil {
		return true
	}

	return false
}

// SetFeedback gets a reference to the given FeedbackData and assigns it to the Feedback field.
func (o *AnomalyWithContext) SetFeedback(v FeedbackData) {
	o.Feedback = &v
}

func (o AnomalyWithContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["anomaly"] = o.Anomaly
	}
	if true {
		toSerialize["data"] = o.Data
	}
	if o.Feedback != nil {
		toSerialize["feedback"] = o.Feedback
	}
	return json.Marshal(toSerialize)
}

type NullableAnomalyWithContext struct {
	value *AnomalyWithContext
	isSet bool
}

func (v NullableAnomalyWithContext) Get() *AnomalyWithContext {
	return v.value
}

func (v *NullableAnomalyWithContext) Set(val *AnomalyWithContext) {
	v.value = val
	v.isSet = true
}

func (v NullableAnomalyWithContext) IsSet() bool {
	return v.isSet
}

func (v *NullableAnomalyWithContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnomalyWithContext(val *AnomalyWithContext) *NullableAnomalyWithContext {
	return &NullableAnomalyWithContext{value: val, isSet: true}
}

func (v NullableAnomalyWithContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnomalyWithContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
