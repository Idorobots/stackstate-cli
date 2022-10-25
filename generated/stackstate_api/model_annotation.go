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

// Annotation struct for Annotation
type Annotation struct {
	Id                string          `json:"id"`
	Name              string          `json:"name"`
	Reference         Reference       `json:"reference"`
	Identifiers       []string        `json:"identifiers"`
	Description       string          `json:"description"`
	AnnotationType    AnnotationType  `json:"annotationType"`
	EventTimeInterval TimeRange       `json:"eventTimeInterval"`
	ProcessedTime     int64           `json:"processedTime"`
	CreatedTime       int64           `json:"createdTime"`
	Tags              []string        `json:"tags"`
	Data              *AnnotationData `json:"data,omitempty"`
}

// NewAnnotation instantiates a new Annotation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnnotation(id string, name string, reference Reference, identifiers []string, description string, annotationType AnnotationType, eventTimeInterval TimeRange, processedTime int64, createdTime int64, tags []string) *Annotation {
	this := Annotation{}
	this.Id = id
	this.Name = name
	this.Reference = reference
	this.Identifiers = identifiers
	this.Description = description
	this.AnnotationType = annotationType
	this.EventTimeInterval = eventTimeInterval
	this.ProcessedTime = processedTime
	this.CreatedTime = createdTime
	this.Tags = tags
	return &this
}

// NewAnnotationWithDefaults instantiates a new Annotation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnnotationWithDefaults() *Annotation {
	this := Annotation{}
	return &this
}

// GetId returns the Id field value
func (o *Annotation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Annotation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Annotation) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Annotation) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Annotation) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Annotation) SetName(v string) {
	o.Name = v
}

// GetReference returns the Reference field value
func (o *Annotation) GetReference() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *Annotation) GetReferenceOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *Annotation) SetReference(v Reference) {
	o.Reference = v
}

// GetIdentifiers returns the Identifiers field value
func (o *Annotation) GetIdentifiers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Identifiers
}

// GetIdentifiersOk returns a tuple with the Identifiers field value
// and a boolean to check if the value has been set.
func (o *Annotation) GetIdentifiersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identifiers, true
}

// SetIdentifiers sets field value
func (o *Annotation) SetIdentifiers(v []string) {
	o.Identifiers = v
}

// GetDescription returns the Description field value
func (o *Annotation) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Annotation) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Annotation) SetDescription(v string) {
	o.Description = v
}

// GetAnnotationType returns the AnnotationType field value
func (o *Annotation) GetAnnotationType() AnnotationType {
	if o == nil {
		var ret AnnotationType
		return ret
	}

	return o.AnnotationType
}

// GetAnnotationTypeOk returns a tuple with the AnnotationType field value
// and a boolean to check if the value has been set.
func (o *Annotation) GetAnnotationTypeOk() (*AnnotationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnnotationType, true
}

// SetAnnotationType sets field value
func (o *Annotation) SetAnnotationType(v AnnotationType) {
	o.AnnotationType = v
}

// GetEventTimeInterval returns the EventTimeInterval field value
func (o *Annotation) GetEventTimeInterval() TimeRange {
	if o == nil {
		var ret TimeRange
		return ret
	}

	return o.EventTimeInterval
}

// GetEventTimeIntervalOk returns a tuple with the EventTimeInterval field value
// and a boolean to check if the value has been set.
func (o *Annotation) GetEventTimeIntervalOk() (*TimeRange, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTimeInterval, true
}

// SetEventTimeInterval sets field value
func (o *Annotation) SetEventTimeInterval(v TimeRange) {
	o.EventTimeInterval = v
}

// GetProcessedTime returns the ProcessedTime field value
func (o *Annotation) GetProcessedTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ProcessedTime
}

// GetProcessedTimeOk returns a tuple with the ProcessedTime field value
// and a boolean to check if the value has been set.
func (o *Annotation) GetProcessedTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessedTime, true
}

// SetProcessedTime sets field value
func (o *Annotation) SetProcessedTime(v int64) {
	o.ProcessedTime = v
}

// GetCreatedTime returns the CreatedTime field value
func (o *Annotation) GetCreatedTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *Annotation) GetCreatedTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *Annotation) SetCreatedTime(v int64) {
	o.CreatedTime = v
}

// GetTags returns the Tags field value
func (o *Annotation) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *Annotation) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *Annotation) SetTags(v []string) {
	o.Tags = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Annotation) GetData() AnnotationData {
	if o == nil || o.Data == nil {
		var ret AnnotationData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Annotation) GetDataOk() (*AnnotationData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Annotation) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given AnnotationData and assigns it to the Data field.
func (o *Annotation) SetData(v AnnotationData) {
	o.Data = &v
}

func (o Annotation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["reference"] = o.Reference
	}
	if true {
		toSerialize["identifiers"] = o.Identifiers
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["annotationType"] = o.AnnotationType
	}
	if true {
		toSerialize["eventTimeInterval"] = o.EventTimeInterval
	}
	if true {
		toSerialize["processedTime"] = o.ProcessedTime
	}
	if true {
		toSerialize["createdTime"] = o.CreatedTime
	}
	if true {
		toSerialize["tags"] = o.Tags
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAnnotation struct {
	value *Annotation
	isSet bool
}

func (v NullableAnnotation) Get() *Annotation {
	return v.value
}

func (v *NullableAnnotation) Set(val *Annotation) {
	v.value = val
	v.isSet = true
}

func (v NullableAnnotation) IsSet() bool {
	return v.isSet
}

func (v *NullableAnnotation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnnotation(val *Annotation) *NullableAnnotation {
	return &NullableAnnotation{value: val, isSet: true}
}

func (v NullableAnnotation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnnotation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
