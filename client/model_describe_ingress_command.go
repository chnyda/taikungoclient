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

// checks if the DescribeIngressCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeIngressCommand{}

// DescribeIngressCommand struct for DescribeIngressCommand
type DescribeIngressCommand struct {
	ProjectId int32 `json:"projectId"`
	Name string `json:"name"`
	Namespace string `json:"namespace"`
}

// NewDescribeIngressCommand instantiates a new DescribeIngressCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeIngressCommand(projectId int32, name string, namespace string) *DescribeIngressCommand {
	this := DescribeIngressCommand{}
	this.ProjectId = projectId
	this.Name = name
	this.Namespace = namespace
	return &this
}

// NewDescribeIngressCommandWithDefaults instantiates a new DescribeIngressCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeIngressCommandWithDefaults() *DescribeIngressCommand {
	this := DescribeIngressCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *DescribeIngressCommand) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *DescribeIngressCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *DescribeIngressCommand) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetName returns the Name field value
func (o *DescribeIngressCommand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DescribeIngressCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DescribeIngressCommand) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value
func (o *DescribeIngressCommand) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *DescribeIngressCommand) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *DescribeIngressCommand) SetNamespace(v string) {
	o.Namespace = v
}

func (o DescribeIngressCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeIngressCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["name"] = o.Name
	toSerialize["namespace"] = o.Namespace
	return toSerialize, nil
}

type NullableDescribeIngressCommand struct {
	value *DescribeIngressCommand
	isSet bool
}

func (v NullableDescribeIngressCommand) Get() *DescribeIngressCommand {
	return v.value
}

func (v *NullableDescribeIngressCommand) Set(val *DescribeIngressCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeIngressCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeIngressCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeIngressCommand(val *DescribeIngressCommand) *NullableDescribeIngressCommand {
	return &NullableDescribeIngressCommand{value: val, isSet: true}
}

func (v NullableDescribeIngressCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeIngressCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


