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

// checks if the PatchCronJobCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchCronJobCommand{}

// PatchCronJobCommand struct for PatchCronJobCommand
type PatchCronJobCommand struct {
	ProjectId int32 `json:"projectId"`
	Yaml string `json:"yaml"`
	Name string `json:"name"`
	Namespace string `json:"namespace"`
}

// NewPatchCronJobCommand instantiates a new PatchCronJobCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchCronJobCommand(projectId int32, yaml string, name string, namespace string) *PatchCronJobCommand {
	this := PatchCronJobCommand{}
	this.ProjectId = projectId
	this.Yaml = yaml
	this.Name = name
	this.Namespace = namespace
	return &this
}

// NewPatchCronJobCommandWithDefaults instantiates a new PatchCronJobCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchCronJobCommandWithDefaults() *PatchCronJobCommand {
	this := PatchCronJobCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *PatchCronJobCommand) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *PatchCronJobCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *PatchCronJobCommand) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetYaml returns the Yaml field value
func (o *PatchCronJobCommand) GetYaml() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Yaml
}

// GetYamlOk returns a tuple with the Yaml field value
// and a boolean to check if the value has been set.
func (o *PatchCronJobCommand) GetYamlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Yaml, true
}

// SetYaml sets field value
func (o *PatchCronJobCommand) SetYaml(v string) {
	o.Yaml = v
}

// GetName returns the Name field value
func (o *PatchCronJobCommand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PatchCronJobCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PatchCronJobCommand) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value
func (o *PatchCronJobCommand) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *PatchCronJobCommand) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *PatchCronJobCommand) SetNamespace(v string) {
	o.Namespace = v
}

func (o PatchCronJobCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchCronJobCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["yaml"] = o.Yaml
	toSerialize["name"] = o.Name
	toSerialize["namespace"] = o.Namespace
	return toSerialize, nil
}

type NullablePatchCronJobCommand struct {
	value *PatchCronJobCommand
	isSet bool
}

func (v NullablePatchCronJobCommand) Get() *PatchCronJobCommand {
	return v.value
}

func (v *NullablePatchCronJobCommand) Set(val *PatchCronJobCommand) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchCronJobCommand) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchCronJobCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchCronJobCommand(val *PatchCronJobCommand) *NullablePatchCronJobCommand {
	return &NullablePatchCronJobCommand{value: val, isSet: true}
}

func (v NullablePatchCronJobCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchCronJobCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


