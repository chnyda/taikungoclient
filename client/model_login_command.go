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

// checks if the LoginCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoginCommand{}

// LoginCommand struct for LoginCommand
type LoginCommand struct {
	Email *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
	Mode *string `json:"mode,omitempty"`
	AccessKey *string `json:"accessKey,omitempty"`
	SecretKey *string `json:"secretKey,omitempty"`
}

// NewLoginCommand instantiates a new LoginCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginCommand() *LoginCommand {
	this := LoginCommand{}
	return &this
}

// NewLoginCommandWithDefaults instantiates a new LoginCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginCommandWithDefaults() *LoginCommand {
	this := LoginCommand{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *LoginCommand) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginCommand) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *LoginCommand) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *LoginCommand) SetEmail(v string) {
	o.Email = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *LoginCommand) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginCommand) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *LoginCommand) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *LoginCommand) SetPassword(v string) {
	o.Password = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *LoginCommand) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginCommand) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *LoginCommand) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *LoginCommand) SetMode(v string) {
	o.Mode = &v
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *LoginCommand) GetAccessKey() string {
	if o == nil || IsNil(o.AccessKey) {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginCommand) GetAccessKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AccessKey) {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *LoginCommand) HasAccessKey() bool {
	if o != nil && !IsNil(o.AccessKey) {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *LoginCommand) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *LoginCommand) GetSecretKey() string {
	if o == nil || IsNil(o.SecretKey) {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginCommand) GetSecretKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SecretKey) {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *LoginCommand) HasSecretKey() bool {
	if o != nil && !IsNil(o.SecretKey) {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *LoginCommand) SetSecretKey(v string) {
	o.SecretKey = &v
}

func (o LoginCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoginCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.AccessKey) {
		toSerialize["accessKey"] = o.AccessKey
	}
	if !IsNil(o.SecretKey) {
		toSerialize["secretKey"] = o.SecretKey
	}
	return toSerialize, nil
}

type NullableLoginCommand struct {
	value *LoginCommand
	isSet bool
}

func (v NullableLoginCommand) Get() *LoginCommand {
	return v.value
}

func (v *NullableLoginCommand) Set(val *LoginCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginCommand(val *LoginCommand) *NullableLoginCommand {
	return &NullableLoginCommand{value: val, isSet: true}
}

func (v NullableLoginCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


