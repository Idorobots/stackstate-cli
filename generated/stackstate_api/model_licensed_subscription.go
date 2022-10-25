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

// LicensedSubscription struct for LicensedSubscription
type LicensedSubscription struct {
	Type         string       `json:"_type"`
	Subscription Subscription `json:"subscription"`
}

// NewLicensedSubscription instantiates a new LicensedSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicensedSubscription(type_ string, subscription Subscription) *LicensedSubscription {
	this := LicensedSubscription{}
	this.Type = type_
	this.Subscription = subscription
	return &this
}

// NewLicensedSubscriptionWithDefaults instantiates a new LicensedSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicensedSubscriptionWithDefaults() *LicensedSubscription {
	this := LicensedSubscription{}
	return &this
}

// GetType returns the Type field value
func (o *LicensedSubscription) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LicensedSubscription) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LicensedSubscription) SetType(v string) {
	o.Type = v
}

// GetSubscription returns the Subscription field value
func (o *LicensedSubscription) GetSubscription() Subscription {
	if o == nil {
		var ret Subscription
		return ret
	}

	return o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value
// and a boolean to check if the value has been set.
func (o *LicensedSubscription) GetSubscriptionOk() (*Subscription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscription, true
}

// SetSubscription sets field value
func (o *LicensedSubscription) SetSubscription(v Subscription) {
	o.Subscription = v
}

func (o LicensedSubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if true {
		toSerialize["subscription"] = o.Subscription
	}
	return json.Marshal(toSerialize)
}

type NullableLicensedSubscription struct {
	value *LicensedSubscription
	isSet bool
}

func (v NullableLicensedSubscription) Get() *LicensedSubscription {
	return v.value
}

func (v *NullableLicensedSubscription) Set(val *LicensedSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableLicensedSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableLicensedSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicensedSubscription(val *LicensedSubscription) *NullableLicensedSubscription {
	return &NullableLicensedSubscription{value: val, isSet: true}
}

func (v NullableLicensedSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicensedSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
