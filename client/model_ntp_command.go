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

// checks if the NtpCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NtpCommand{}

// NtpCommand struct for NtpCommand
type NtpCommand struct {
	Address string `json:"address"`
}

// NewNtpCommand instantiates a new NtpCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNtpCommand(address string) *NtpCommand {
	this := NtpCommand{}
	this.Address = address
	return &this
}

// NewNtpCommandWithDefaults instantiates a new NtpCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNtpCommandWithDefaults() *NtpCommand {
	this := NtpCommand{}
	return &this
}

// GetAddress returns the Address field value
func (o *NtpCommand) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *NtpCommand) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *NtpCommand) SetAddress(v string) {
	o.Address = v
}

func (o NtpCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NtpCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	return toSerialize, nil
}

type NullableNtpCommand struct {
	value *NtpCommand
	isSet bool
}

func (v NullableNtpCommand) Get() *NtpCommand {
	return v.value
}

func (v *NullableNtpCommand) Set(val *NtpCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableNtpCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableNtpCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNtpCommand(val *NtpCommand) *NullableNtpCommand {
	return &NullableNtpCommand{value: val, isSet: true}
}

func (v NullableNtpCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNtpCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


