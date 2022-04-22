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

// ServerVersion struct for ServerVersion
type ServerVersion struct {
	Major int32 `json:"major"`
	Minor int32 `json:"minor"`
	Patch int32 `json:"patch"`
	Diff string `json:"diff"`
	Commit string `json:"commit"`
	IsDev bool `json:"isDev"`
}

// NewServerVersion instantiates a new ServerVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerVersion(major int32, minor int32, patch int32, diff string, commit string, isDev bool) *ServerVersion {
	this := ServerVersion{}
	this.Major = major
	this.Minor = minor
	this.Patch = patch
	this.Diff = diff
	this.Commit = commit
	this.IsDev = isDev
	return &this
}

// NewServerVersionWithDefaults instantiates a new ServerVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerVersionWithDefaults() *ServerVersion {
	this := ServerVersion{}
	return &this
}

// GetMajor returns the Major field value
func (o *ServerVersion) GetMajor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Major
}

// GetMajorOk returns a tuple with the Major field value
// and a boolean to check if the value has been set.
func (o *ServerVersion) GetMajorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Major, true
}

// SetMajor sets field value
func (o *ServerVersion) SetMajor(v int32) {
	o.Major = v
}

// GetMinor returns the Minor field value
func (o *ServerVersion) GetMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Minor
}

// GetMinorOk returns a tuple with the Minor field value
// and a boolean to check if the value has been set.
func (o *ServerVersion) GetMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Minor, true
}

// SetMinor sets field value
func (o *ServerVersion) SetMinor(v int32) {
	o.Minor = v
}

// GetPatch returns the Patch field value
func (o *ServerVersion) GetPatch() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Patch
}

// GetPatchOk returns a tuple with the Patch field value
// and a boolean to check if the value has been set.
func (o *ServerVersion) GetPatchOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patch, true
}

// SetPatch sets field value
func (o *ServerVersion) SetPatch(v int32) {
	o.Patch = v
}

// GetDiff returns the Diff field value
func (o *ServerVersion) GetDiff() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Diff
}

// GetDiffOk returns a tuple with the Diff field value
// and a boolean to check if the value has been set.
func (o *ServerVersion) GetDiffOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Diff, true
}

// SetDiff sets field value
func (o *ServerVersion) SetDiff(v string) {
	o.Diff = v
}

// GetCommit returns the Commit field value
func (o *ServerVersion) GetCommit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Commit
}

// GetCommitOk returns a tuple with the Commit field value
// and a boolean to check if the value has been set.
func (o *ServerVersion) GetCommitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Commit, true
}

// SetCommit sets field value
func (o *ServerVersion) SetCommit(v string) {
	o.Commit = v
}

// GetIsDev returns the IsDev field value
func (o *ServerVersion) GetIsDev() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDev
}

// GetIsDevOk returns a tuple with the IsDev field value
// and a boolean to check if the value has been set.
func (o *ServerVersion) GetIsDevOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDev, true
}

// SetIsDev sets field value
func (o *ServerVersion) SetIsDev(v bool) {
	o.IsDev = v
}

func (o ServerVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["major"] = o.Major
	}
	if true {
		toSerialize["minor"] = o.Minor
	}
	if true {
		toSerialize["patch"] = o.Patch
	}
	if true {
		toSerialize["diff"] = o.Diff
	}
	if true {
		toSerialize["commit"] = o.Commit
	}
	if true {
		toSerialize["isDev"] = o.IsDev
	}
	return json.Marshal(toSerialize)
}

type NullableServerVersion struct {
	value *ServerVersion
	isSet bool
}

func (v NullableServerVersion) Get() *ServerVersion {
	return v.value
}

func (v *NullableServerVersion) Set(val *ServerVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableServerVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableServerVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerVersion(val *ServerVersion) *NullableServerVersion {
	return &NullableServerVersion{value: val, isSet: true}
}

func (v NullableServerVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


