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

// checks if the EditProjectAppParamCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditProjectAppParamCommand{}

// EditProjectAppParamCommand struct for EditProjectAppParamCommand
type EditProjectAppParamCommand struct {
	ProjectAppId int32 `json:"projectAppId"`
	Parameters []EditProjectAppParamsDto `json:"parameters,omitempty"`
}

// NewEditProjectAppParamCommand instantiates a new EditProjectAppParamCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditProjectAppParamCommand(projectAppId int32) *EditProjectAppParamCommand {
	this := EditProjectAppParamCommand{}
	this.ProjectAppId = projectAppId
	return &this
}

// NewEditProjectAppParamCommandWithDefaults instantiates a new EditProjectAppParamCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditProjectAppParamCommandWithDefaults() *EditProjectAppParamCommand {
	this := EditProjectAppParamCommand{}
	return &this
}

// GetProjectAppId returns the ProjectAppId field value
func (o *EditProjectAppParamCommand) GetProjectAppId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectAppId
}

// GetProjectAppIdOk returns a tuple with the ProjectAppId field value
// and a boolean to check if the value has been set.
func (o *EditProjectAppParamCommand) GetProjectAppIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectAppId, true
}

// SetProjectAppId sets field value
func (o *EditProjectAppParamCommand) SetProjectAppId(v int32) {
	o.ProjectAppId = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *EditProjectAppParamCommand) GetParameters() []EditProjectAppParamsDto {
	if o == nil || IsNil(o.Parameters) {
		var ret []EditProjectAppParamsDto
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProjectAppParamCommand) GetParametersOk() ([]EditProjectAppParamsDto, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *EditProjectAppParamCommand) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []EditProjectAppParamsDto and assigns it to the Parameters field.
func (o *EditProjectAppParamCommand) SetParameters(v []EditProjectAppParamsDto) {
	o.Parameters = v
}

func (o EditProjectAppParamCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditProjectAppParamCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectAppId"] = o.ProjectAppId
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	return toSerialize, nil
}

type NullableEditProjectAppParamCommand struct {
	value *EditProjectAppParamCommand
	isSet bool
}

func (v NullableEditProjectAppParamCommand) Get() *EditProjectAppParamCommand {
	return v.value
}

func (v *NullableEditProjectAppParamCommand) Set(val *EditProjectAppParamCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableEditProjectAppParamCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableEditProjectAppParamCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditProjectAppParamCommand(val *EditProjectAppParamCommand) *NullableEditProjectAppParamCommand {
	return &NullableEditProjectAppParamCommand{value: val, isSet: true}
}

func (v NullableEditProjectAppParamCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditProjectAppParamCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


