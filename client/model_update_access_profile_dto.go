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

// checks if the UpdateAccessProfileDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAccessProfileDto{}

// UpdateAccessProfileDto struct for UpdateAccessProfileDto
type UpdateAccessProfileDto struct {
	Name *string `json:"name,omitempty"`
	HttpProxy *string `json:"httpProxy,omitempty"`
}

// NewUpdateAccessProfileDto instantiates a new UpdateAccessProfileDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAccessProfileDto() *UpdateAccessProfileDto {
	this := UpdateAccessProfileDto{}
	return &this
}

// NewUpdateAccessProfileDtoWithDefaults instantiates a new UpdateAccessProfileDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAccessProfileDtoWithDefaults() *UpdateAccessProfileDto {
	this := UpdateAccessProfileDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateAccessProfileDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccessProfileDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateAccessProfileDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateAccessProfileDto) SetName(v string) {
	o.Name = &v
}

// GetHttpProxy returns the HttpProxy field value if set, zero value otherwise.
func (o *UpdateAccessProfileDto) GetHttpProxy() string {
	if o == nil || IsNil(o.HttpProxy) {
		var ret string
		return ret
	}
	return *o.HttpProxy
}

// GetHttpProxyOk returns a tuple with the HttpProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccessProfileDto) GetHttpProxyOk() (*string, bool) {
	if o == nil || IsNil(o.HttpProxy) {
		return nil, false
	}
	return o.HttpProxy, true
}

// HasHttpProxy returns a boolean if a field has been set.
func (o *UpdateAccessProfileDto) HasHttpProxy() bool {
	if o != nil && !IsNil(o.HttpProxy) {
		return true
	}

	return false
}

// SetHttpProxy gets a reference to the given string and assigns it to the HttpProxy field.
func (o *UpdateAccessProfileDto) SetHttpProxy(v string) {
	o.HttpProxy = &v
}

func (o UpdateAccessProfileDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAccessProfileDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.HttpProxy) {
		toSerialize["httpProxy"] = o.HttpProxy
	}
	return toSerialize, nil
}

type NullableUpdateAccessProfileDto struct {
	value *UpdateAccessProfileDto
	isSet bool
}

func (v NullableUpdateAccessProfileDto) Get() *UpdateAccessProfileDto {
	return v.value
}

func (v *NullableUpdateAccessProfileDto) Set(val *UpdateAccessProfileDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAccessProfileDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAccessProfileDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAccessProfileDto(val *UpdateAccessProfileDto) *NullableUpdateAccessProfileDto {
	return &NullableUpdateAccessProfileDto{value: val, isSet: true}
}

func (v NullableUpdateAccessProfileDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAccessProfileDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


