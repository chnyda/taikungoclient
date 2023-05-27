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

// checks if the DescribeNodeCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeNodeCommand{}

// DescribeNodeCommand struct for DescribeNodeCommand
type DescribeNodeCommand struct {
	ProjectId int32 `json:"projectId"`
	Name string `json:"name"`
}

// NewDescribeNodeCommand instantiates a new DescribeNodeCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeNodeCommand(projectId int32, name string) *DescribeNodeCommand {
	this := DescribeNodeCommand{}
	this.ProjectId = projectId
	this.Name = name
	return &this
}

// NewDescribeNodeCommandWithDefaults instantiates a new DescribeNodeCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeNodeCommandWithDefaults() *DescribeNodeCommand {
	this := DescribeNodeCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *DescribeNodeCommand) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *DescribeNodeCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *DescribeNodeCommand) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetName returns the Name field value
func (o *DescribeNodeCommand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DescribeNodeCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DescribeNodeCommand) SetName(v string) {
	o.Name = v
}

func (o DescribeNodeCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeNodeCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableDescribeNodeCommand struct {
	value *DescribeNodeCommand
	isSet bool
}

func (v NullableDescribeNodeCommand) Get() *DescribeNodeCommand {
	return v.value
}

func (v *NullableDescribeNodeCommand) Set(val *DescribeNodeCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeNodeCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeNodeCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeNodeCommand(val *DescribeNodeCommand) *NullableDescribeNodeCommand {
	return &NullableDescribeNodeCommand{value: val, isSet: true}
}

func (v NullableDescribeNodeCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeNodeCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


