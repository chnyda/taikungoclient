/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@itera.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikuncore

import (
	"encoding/json"
)

// checks if the UserExistCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserExistCommand{}

// UserExistCommand struct for UserExistCommand
type UserExistCommand struct {
	Email NullableString `json:"email,omitempty"`
	UserName NullableString `json:"userName,omitempty"`
}

// NewUserExistCommand instantiates a new UserExistCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserExistCommand() *UserExistCommand {
	this := UserExistCommand{}
	return &this
}

// NewUserExistCommandWithDefaults instantiates a new UserExistCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserExistCommandWithDefaults() *UserExistCommand {
	this := UserExistCommand{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserExistCommand) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserExistCommand) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *UserExistCommand) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *UserExistCommand) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *UserExistCommand) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *UserExistCommand) UnsetEmail() {
	o.Email.Unset()
}

// GetUserName returns the UserName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserExistCommand) GetUserName() string {
	if o == nil || IsNil(o.UserName.Get()) {
		var ret string
		return ret
	}
	return *o.UserName.Get()
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserExistCommand) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserName.Get(), o.UserName.IsSet()
}

// HasUserName returns a boolean if a field has been set.
func (o *UserExistCommand) HasUserName() bool {
	if o != nil && o.UserName.IsSet() {
		return true
	}

	return false
}

// SetUserName gets a reference to the given NullableString and assigns it to the UserName field.
func (o *UserExistCommand) SetUserName(v string) {
	o.UserName.Set(&v)
}
// SetUserNameNil sets the value for UserName to be an explicit nil
func (o *UserExistCommand) SetUserNameNil() {
	o.UserName.Set(nil)
}

// UnsetUserName ensures that no value is present for UserName, not even an explicit nil
func (o *UserExistCommand) UnsetUserName() {
	o.UserName.Unset()
}

func (o UserExistCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserExistCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.UserName.IsSet() {
		toSerialize["userName"] = o.UserName.Get()
	}
	return toSerialize, nil
}

type NullableUserExistCommand struct {
	value *UserExistCommand
	isSet bool
}

func (v NullableUserExistCommand) Get() *UserExistCommand {
	return v.value
}

func (v *NullableUserExistCommand) Set(val *UserExistCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableUserExistCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableUserExistCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserExistCommand(val *UserExistCommand) *NullableUserExistCommand {
	return &NullableUserExistCommand{value: val, isSet: true}
}

func (v NullableUserExistCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserExistCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


