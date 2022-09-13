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

// Envelope struct for Envelope
type Envelope struct {
	Status string `json:"status"`
	Data *Data `json:"data,omitempty"`
	ErrorType *string `json:"errorType,omitempty"`
	Error *string `json:"error,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}

// NewEnvelope instantiates a new Envelope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvelope(status string) *Envelope {
	this := Envelope{}
	this.Status = status
	return &this
}

// NewEnvelopeWithDefaults instantiates a new Envelope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvelopeWithDefaults() *Envelope {
	this := Envelope{}
	return &this
}

// GetStatus returns the Status field value
func (o *Envelope) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Envelope) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Envelope) SetStatus(v string) {
	o.Status = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Envelope) GetData() Data {
	if o == nil || o.Data == nil {
		var ret Data
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Envelope) GetDataOk() (*Data, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Envelope) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given Data and assigns it to the Data field.
func (o *Envelope) SetData(v Data) {
	o.Data = &v
}

// GetErrorType returns the ErrorType field value if set, zero value otherwise.
func (o *Envelope) GetErrorType() string {
	if o == nil || o.ErrorType == nil {
		var ret string
		return ret
	}
	return *o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Envelope) GetErrorTypeOk() (*string, bool) {
	if o == nil || o.ErrorType == nil {
		return nil, false
	}
	return o.ErrorType, true
}

// HasErrorType returns a boolean if a field has been set.
func (o *Envelope) HasErrorType() bool {
	if o != nil && o.ErrorType != nil {
		return true
	}

	return false
}

// SetErrorType gets a reference to the given string and assigns it to the ErrorType field.
func (o *Envelope) SetErrorType(v string) {
	o.ErrorType = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *Envelope) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Envelope) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *Envelope) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *Envelope) SetError(v string) {
	o.Error = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *Envelope) GetWarnings() []string {
	if o == nil || o.Warnings == nil {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Envelope) GetWarningsOk() ([]string, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *Envelope) HasWarnings() bool {
	if o != nil && o.Warnings != nil {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *Envelope) SetWarnings(v []string) {
	o.Warnings = v
}

func (o Envelope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.ErrorType != nil {
		toSerialize["errorType"] = o.ErrorType
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	return json.Marshal(toSerialize)
}

type NullableEnvelope struct {
	value *Envelope
	isSet bool
}

func (v NullableEnvelope) Get() *Envelope {
	return v.value
}

func (v *NullableEnvelope) Set(val *Envelope) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvelope) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvelope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvelope(val *Envelope) *NullableEnvelope {
	return &NullableEnvelope{value: val, isSet: true}
}

func (v NullableEnvelope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvelope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


