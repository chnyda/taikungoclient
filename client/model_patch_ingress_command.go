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

// checks if the PatchIngressCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchIngressCommand{}

// PatchIngressCommand struct for PatchIngressCommand
type PatchIngressCommand struct {
	ProjectId int32 `json:"projectId"`
	Yaml string `json:"yaml"`
	Name string `json:"name"`
	Namespace string `json:"namespace"`
}

// NewPatchIngressCommand instantiates a new PatchIngressCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchIngressCommand(projectId int32, yaml string, name string, namespace string) *PatchIngressCommand {
	this := PatchIngressCommand{}
	this.ProjectId = projectId
	this.Yaml = yaml
	this.Name = name
	this.Namespace = namespace
	return &this
}

// NewPatchIngressCommandWithDefaults instantiates a new PatchIngressCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchIngressCommandWithDefaults() *PatchIngressCommand {
	this := PatchIngressCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *PatchIngressCommand) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *PatchIngressCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *PatchIngressCommand) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetYaml returns the Yaml field value
func (o *PatchIngressCommand) GetYaml() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Yaml
}

// GetYamlOk returns a tuple with the Yaml field value
// and a boolean to check if the value has been set.
func (o *PatchIngressCommand) GetYamlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Yaml, true
}

// SetYaml sets field value
func (o *PatchIngressCommand) SetYaml(v string) {
	o.Yaml = v
}

// GetName returns the Name field value
func (o *PatchIngressCommand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PatchIngressCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PatchIngressCommand) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value
func (o *PatchIngressCommand) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *PatchIngressCommand) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *PatchIngressCommand) SetNamespace(v string) {
	o.Namespace = v
}

func (o PatchIngressCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchIngressCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["yaml"] = o.Yaml
	toSerialize["name"] = o.Name
	toSerialize["namespace"] = o.Namespace
	return toSerialize, nil
}

type NullablePatchIngressCommand struct {
	value *PatchIngressCommand
	isSet bool
}

func (v NullablePatchIngressCommand) Get() *PatchIngressCommand {
	return v.value
}

func (v *NullablePatchIngressCommand) Set(val *PatchIngressCommand) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchIngressCommand) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchIngressCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchIngressCommand(val *PatchIngressCommand) *NullablePatchIngressCommand {
	return &NullablePatchIngressCommand{value: val, isSet: true}
}

func (v NullablePatchIngressCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchIngressCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


