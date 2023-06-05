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

// checks if the AdminUserCreateCommandPassword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminUserCreateCommandPassword{}

// AdminUserCreateCommandPassword struct for AdminUserCreateCommandPassword
type AdminUserCreateCommandPassword struct {
}

// NewAdminUserCreateCommandPassword instantiates a new AdminUserCreateCommandPassword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminUserCreateCommandPassword() *AdminUserCreateCommandPassword {
	this := AdminUserCreateCommandPassword{}
	return &this
}

// NewAdminUserCreateCommandPasswordWithDefaults instantiates a new AdminUserCreateCommandPassword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminUserCreateCommandPasswordWithDefaults() *AdminUserCreateCommandPassword {
	this := AdminUserCreateCommandPassword{}
	return &this
}

func (o AdminUserCreateCommandPassword) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminUserCreateCommandPassword) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableAdminUserCreateCommandPassword struct {
	value *AdminUserCreateCommandPassword
	isSet bool
}

func (v NullableAdminUserCreateCommandPassword) Get() *AdminUserCreateCommandPassword {
	return v.value
}

func (v *NullableAdminUserCreateCommandPassword) Set(val *AdminUserCreateCommandPassword) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminUserCreateCommandPassword) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminUserCreateCommandPassword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminUserCreateCommandPassword(val *AdminUserCreateCommandPassword) *NullableAdminUserCreateCommandPassword {
	return &NullableAdminUserCreateCommandPassword{value: val, isSet: true}
}

func (v NullableAdminUserCreateCommandPassword) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminUserCreateCommandPassword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

