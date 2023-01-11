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

// UserProfileSaveError - struct for UserProfileSaveError
type UserProfileSaveError struct {
	UserNameMismatchError *UserNameMismatchError
	UserNotFoundError     *UserNotFoundError
	UserNotLoggedInError  *UserNotLoggedInError
}

// UserNameMismatchErrorAsUserProfileSaveError is a convenience function that returns UserNameMismatchError wrapped in UserProfileSaveError
func UserNameMismatchErrorAsUserProfileSaveError(v *UserNameMismatchError) UserProfileSaveError {
	return UserProfileSaveError{
		UserNameMismatchError: v,
	}
}

// UserNotFoundErrorAsUserProfileSaveError is a convenience function that returns UserNotFoundError wrapped in UserProfileSaveError
func UserNotFoundErrorAsUserProfileSaveError(v *UserNotFoundError) UserProfileSaveError {
	return UserProfileSaveError{
		UserNotFoundError: v,
	}
}

// UserNotLoggedInErrorAsUserProfileSaveError is a convenience function that returns UserNotLoggedInError wrapped in UserProfileSaveError
func UserNotLoggedInErrorAsUserProfileSaveError(v *UserNotLoggedInError) UserProfileSaveError {
	return UserProfileSaveError{
		UserNotLoggedInError: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UserProfileSaveError) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'UserNameMismatchError'
	if jsonDict["_type"] == "UserNameMismatchError" {
		// try to unmarshal JSON data into UserNameMismatchError
		err = json.Unmarshal(data, &dst.UserNameMismatchError)
		if err == nil {
			return nil // data stored in dst.UserNameMismatchError, return on the first match
		} else {
			dst.UserNameMismatchError = nil
			return fmt.Errorf("Failed to unmarshal UserProfileSaveError as UserNameMismatchError: %s", err.Error())
		}
	}

	// check if the discriminator value is 'UserNotFoundError'
	if jsonDict["_type"] == "UserNotFoundError" {
		// try to unmarshal JSON data into UserNotFoundError
		err = json.Unmarshal(data, &dst.UserNotFoundError)
		if err == nil {
			return nil // data stored in dst.UserNotFoundError, return on the first match
		} else {
			dst.UserNotFoundError = nil
			return fmt.Errorf("Failed to unmarshal UserProfileSaveError as UserNotFoundError: %s", err.Error())
		}
	}

	// check if the discriminator value is 'UserNotLoggedInError'
	if jsonDict["_type"] == "UserNotLoggedInError" {
		// try to unmarshal JSON data into UserNotLoggedInError
		err = json.Unmarshal(data, &dst.UserNotLoggedInError)
		if err == nil {
			return nil // data stored in dst.UserNotLoggedInError, return on the first match
		} else {
			dst.UserNotLoggedInError = nil
			return fmt.Errorf("Failed to unmarshal UserProfileSaveError as UserNotLoggedInError: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UserProfileSaveError) MarshalJSON() ([]byte, error) {
	if src.UserNameMismatchError != nil {
		return json.Marshal(&src.UserNameMismatchError)
	}

	if src.UserNotFoundError != nil {
		return json.Marshal(&src.UserNotFoundError)
	}

	if src.UserNotLoggedInError != nil {
		return json.Marshal(&src.UserNotLoggedInError)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UserProfileSaveError) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.UserNameMismatchError != nil {
		return obj.UserNameMismatchError
	}

	if obj.UserNotFoundError != nil {
		return obj.UserNotFoundError
	}

	if obj.UserNotLoggedInError != nil {
		return obj.UserNotLoggedInError
	}

	// all schemas are nil
	return nil
}

type NullableUserProfileSaveError struct {
	value *UserProfileSaveError
	isSet bool
}

func (v NullableUserProfileSaveError) Get() *UserProfileSaveError {
	return v.value
}

func (v *NullableUserProfileSaveError) Set(val *UserProfileSaveError) {
	v.value = val
	v.isSet = true
}

func (v NullableUserProfileSaveError) IsSet() bool {
	return v.isSet
}

func (v *NullableUserProfileSaveError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserProfileSaveError(val *UserProfileSaveError) *NullableUserProfileSaveError {
	return &NullableUserProfileSaveError{value: val, isSet: true}
}

func (v NullableUserProfileSaveError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserProfileSaveError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
