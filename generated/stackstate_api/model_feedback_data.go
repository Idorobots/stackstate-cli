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

// FeedbackData struct for FeedbackData
type FeedbackData struct {
	Type string `json:"_type"`
	Subject string `json:"subject"`
	Thumbsup []string `json:"thumbsup"`
	Thumbsdown []string `json:"thumbsdown"`
	Comments []FeedbackComment `json:"comments"`
	Query *AnnotationMetricQuery `json:"query,omitempty"`
}

// NewFeedbackData instantiates a new FeedbackData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedbackData(type_ string, subject string, thumbsup []string, thumbsdown []string, comments []FeedbackComment) *FeedbackData {
	this := FeedbackData{}
	this.Type = type_
	this.Subject = subject
	this.Thumbsup = thumbsup
	this.Thumbsdown = thumbsdown
	this.Comments = comments
	return &this
}

// NewFeedbackDataWithDefaults instantiates a new FeedbackData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedbackDataWithDefaults() *FeedbackData {
	this := FeedbackData{}
	return &this
}

// GetType returns the Type field value
func (o *FeedbackData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FeedbackData) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FeedbackData) SetType(v string) {
	o.Type = v
}

// GetSubject returns the Subject field value
func (o *FeedbackData) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *FeedbackData) GetSubjectOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *FeedbackData) SetSubject(v string) {
	o.Subject = v
}

// GetThumbsup returns the Thumbsup field value
func (o *FeedbackData) GetThumbsup() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Thumbsup
}

// GetThumbsupOk returns a tuple with the Thumbsup field value
// and a boolean to check if the value has been set.
func (o *FeedbackData) GetThumbsupOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Thumbsup, true
}

// SetThumbsup sets field value
func (o *FeedbackData) SetThumbsup(v []string) {
	o.Thumbsup = v
}

// GetThumbsdown returns the Thumbsdown field value
func (o *FeedbackData) GetThumbsdown() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Thumbsdown
}

// GetThumbsdownOk returns a tuple with the Thumbsdown field value
// and a boolean to check if the value has been set.
func (o *FeedbackData) GetThumbsdownOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Thumbsdown, true
}

// SetThumbsdown sets field value
func (o *FeedbackData) SetThumbsdown(v []string) {
	o.Thumbsdown = v
}

// GetComments returns the Comments field value
func (o *FeedbackData) GetComments() []FeedbackComment {
	if o == nil {
		var ret []FeedbackComment
		return ret
	}

	return o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value
// and a boolean to check if the value has been set.
func (o *FeedbackData) GetCommentsOk() (*[]FeedbackComment, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Comments, true
}

// SetComments sets field value
func (o *FeedbackData) SetComments(v []FeedbackComment) {
	o.Comments = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *FeedbackData) GetQuery() AnnotationMetricQuery {
	if o == nil || o.Query == nil {
		var ret AnnotationMetricQuery
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackData) GetQueryOk() (*AnnotationMetricQuery, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *FeedbackData) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given AnnotationMetricQuery and assigns it to the Query field.
func (o *FeedbackData) SetQuery(v AnnotationMetricQuery) {
	o.Query = &v
}

func (o FeedbackData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if true {
		toSerialize["subject"] = o.Subject
	}
	if true {
		toSerialize["thumbsup"] = o.Thumbsup
	}
	if true {
		toSerialize["thumbsdown"] = o.Thumbsdown
	}
	if true {
		toSerialize["comments"] = o.Comments
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	return json.Marshal(toSerialize)
}

type NullableFeedbackData struct {
	value *FeedbackData
	isSet bool
}

func (v NullableFeedbackData) Get() *FeedbackData {
	return v.value
}

func (v *NullableFeedbackData) Set(val *FeedbackData) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedbackData) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedbackData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedbackData(val *FeedbackData) *NullableFeedbackData {
	return &NullableFeedbackData{value: val, isSet: true}
}

func (v NullableFeedbackData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedbackData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


