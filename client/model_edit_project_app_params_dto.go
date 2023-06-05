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

// checks if the EditProjectAppParamsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditProjectAppParamsDto{}

// EditProjectAppParamsDto struct for EditProjectAppParamsDto
type EditProjectAppParamsDto struct {
	Key string `json:"key"`
	Value *string `json:"value,omitempty"`
}

// NewEditProjectAppParamsDto instantiates a new EditProjectAppParamsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditProjectAppParamsDto(key string) *EditProjectAppParamsDto {
	this := EditProjectAppParamsDto{}
	this.Key = key
	return &this
}

// NewEditProjectAppParamsDtoWithDefaults instantiates a new EditProjectAppParamsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditProjectAppParamsDtoWithDefaults() *EditProjectAppParamsDto {
	this := EditProjectAppParamsDto{}
	return &this
}

// GetKey returns the Key field value
func (o *EditProjectAppParamsDto) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *EditProjectAppParamsDto) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *EditProjectAppParamsDto) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EditProjectAppParamsDto) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProjectAppParamsDto) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EditProjectAppParamsDto) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *EditProjectAppParamsDto) SetValue(v string) {
	o.Value = &v
}

func (o EditProjectAppParamsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditProjectAppParamsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableEditProjectAppParamsDto struct {
	value *EditProjectAppParamsDto
	isSet bool
}

func (v NullableEditProjectAppParamsDto) Get() *EditProjectAppParamsDto {
	return v.value
}

func (v *NullableEditProjectAppParamsDto) Set(val *EditProjectAppParamsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableEditProjectAppParamsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableEditProjectAppParamsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditProjectAppParamsDto(val *EditProjectAppParamsDto) *NullableEditProjectAppParamsDto {
	return &NullableEditProjectAppParamsDto{value: val, isSet: true}
}

func (v NullableEditProjectAppParamsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditProjectAppParamsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


