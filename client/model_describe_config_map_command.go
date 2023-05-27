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

// checks if the DescribeConfigMapCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeConfigMapCommand{}

// DescribeConfigMapCommand struct for DescribeConfigMapCommand
type DescribeConfigMapCommand struct {
	ProjectId int32 `json:"projectId"`
	Name string `json:"name"`
	Namespace string `json:"namespace"`
}

// NewDescribeConfigMapCommand instantiates a new DescribeConfigMapCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeConfigMapCommand(projectId int32, name string, namespace string) *DescribeConfigMapCommand {
	this := DescribeConfigMapCommand{}
	this.ProjectId = projectId
	this.Name = name
	this.Namespace = namespace
	return &this
}

// NewDescribeConfigMapCommandWithDefaults instantiates a new DescribeConfigMapCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeConfigMapCommandWithDefaults() *DescribeConfigMapCommand {
	this := DescribeConfigMapCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *DescribeConfigMapCommand) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *DescribeConfigMapCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *DescribeConfigMapCommand) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetName returns the Name field value
func (o *DescribeConfigMapCommand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DescribeConfigMapCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DescribeConfigMapCommand) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value
func (o *DescribeConfigMapCommand) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *DescribeConfigMapCommand) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *DescribeConfigMapCommand) SetNamespace(v string) {
	o.Namespace = v
}

func (o DescribeConfigMapCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeConfigMapCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["name"] = o.Name
	toSerialize["namespace"] = o.Namespace
	return toSerialize, nil
}

type NullableDescribeConfigMapCommand struct {
	value *DescribeConfigMapCommand
	isSet bool
}

func (v NullableDescribeConfigMapCommand) Get() *DescribeConfigMapCommand {
	return v.value
}

func (v *NullableDescribeConfigMapCommand) Set(val *DescribeConfigMapCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeConfigMapCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeConfigMapCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeConfigMapCommand(val *DescribeConfigMapCommand) *NullableDescribeConfigMapCommand {
	return &NullableDescribeConfigMapCommand{value: val, isSet: true}
}

func (v NullableDescribeConfigMapCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeConfigMapCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


