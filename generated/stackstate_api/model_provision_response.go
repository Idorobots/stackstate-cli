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
)

// ProvisionResponse struct for ProvisionResponse
type ProvisionResponse struct {
	Id                  *int64  `json:"id,omitempty"`
	Status              *string `json:"status,omitempty"`
	Name                *string `json:"name,omitempty"`
	StackPackVersion    *string `json:"stackPackVersion,omitempty"`
	LastUpdateTimestamp *int64  `json:"lastUpdateTimestamp,omitempty"`
}

// NewProvisionResponse instantiates a new ProvisionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisionResponse() *ProvisionResponse {
	this := ProvisionResponse{}
	return &this
}

// NewProvisionResponseWithDefaults instantiates a new ProvisionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisionResponseWithDefaults() *ProvisionResponse {
	this := ProvisionResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProvisionResponse) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionResponse) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProvisionResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ProvisionResponse) SetId(v int64) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ProvisionResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ProvisionResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ProvisionResponse) SetStatus(v string) {
	o.Status = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProvisionResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProvisionResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProvisionResponse) SetName(v string) {
	o.Name = &v
}

// GetStackPackVersion returns the StackPackVersion field value if set, zero value otherwise.
func (o *ProvisionResponse) GetStackPackVersion() string {
	if o == nil || o.StackPackVersion == nil {
		var ret string
		return ret
	}
	return *o.StackPackVersion
}

// GetStackPackVersionOk returns a tuple with the StackPackVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionResponse) GetStackPackVersionOk() (*string, bool) {
	if o == nil || o.StackPackVersion == nil {
		return nil, false
	}
	return o.StackPackVersion, true
}

// HasStackPackVersion returns a boolean if a field has been set.
func (o *ProvisionResponse) HasStackPackVersion() bool {
	if o != nil && o.StackPackVersion != nil {
		return true
	}

	return false
}

// SetStackPackVersion gets a reference to the given string and assigns it to the StackPackVersion field.
func (o *ProvisionResponse) SetStackPackVersion(v string) {
	o.StackPackVersion = &v
}

// GetLastUpdateTimestamp returns the LastUpdateTimestamp field value if set, zero value otherwise.
func (o *ProvisionResponse) GetLastUpdateTimestamp() int64 {
	if o == nil || o.LastUpdateTimestamp == nil {
		var ret int64
		return ret
	}
	return *o.LastUpdateTimestamp
}

// GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionResponse) GetLastUpdateTimestampOk() (*int64, bool) {
	if o == nil || o.LastUpdateTimestamp == nil {
		return nil, false
	}
	return o.LastUpdateTimestamp, true
}

// HasLastUpdateTimestamp returns a boolean if a field has been set.
func (o *ProvisionResponse) HasLastUpdateTimestamp() bool {
	if o != nil && o.LastUpdateTimestamp != nil {
		return true
	}

	return false
}

// SetLastUpdateTimestamp gets a reference to the given int64 and assigns it to the LastUpdateTimestamp field.
func (o *ProvisionResponse) SetLastUpdateTimestamp(v int64) {
	o.LastUpdateTimestamp = &v
}

func (o ProvisionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.StackPackVersion != nil {
		toSerialize["stackPackVersion"] = o.StackPackVersion
	}
	if o.LastUpdateTimestamp != nil {
		toSerialize["lastUpdateTimestamp"] = o.LastUpdateTimestamp
	}
	return json.Marshal(toSerialize)
}

type NullableProvisionResponse struct {
	value *ProvisionResponse
	isSet bool
}

func (v NullableProvisionResponse) Get() *ProvisionResponse {
	return v.value
}

func (v *NullableProvisionResponse) Set(val *ProvisionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisionResponse(val *ProvisionResponse) *NullableProvisionResponse {
	return &NullableProvisionResponse{value: val, isSet: true}
}

func (v NullableProvisionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
