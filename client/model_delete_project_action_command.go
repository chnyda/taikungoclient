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

// checks if the DeleteProjectActionCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteProjectActionCommand{}

// DeleteProjectActionCommand struct for DeleteProjectActionCommand
type DeleteProjectActionCommand struct {
	ProjectId *int32 `json:"projectId,omitempty"`
}

// NewDeleteProjectActionCommand instantiates a new DeleteProjectActionCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteProjectActionCommand() *DeleteProjectActionCommand {
	this := DeleteProjectActionCommand{}
	return &this
}

// NewDeleteProjectActionCommandWithDefaults instantiates a new DeleteProjectActionCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteProjectActionCommandWithDefaults() *DeleteProjectActionCommand {
	this := DeleteProjectActionCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *DeleteProjectActionCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteProjectActionCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *DeleteProjectActionCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *DeleteProjectActionCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

func (o DeleteProjectActionCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteProjectActionCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	return toSerialize, nil
}

type NullableDeleteProjectActionCommand struct {
	value *DeleteProjectActionCommand
	isSet bool
}

func (v NullableDeleteProjectActionCommand) Get() *DeleteProjectActionCommand {
	return v.value
}

func (v *NullableDeleteProjectActionCommand) Set(val *DeleteProjectActionCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteProjectActionCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteProjectActionCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteProjectActionCommand(val *DeleteProjectActionCommand) *NullableDeleteProjectActionCommand {
	return &NullableDeleteProjectActionCommand{value: val, isSet: true}
}

func (v NullableDeleteProjectActionCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteProjectActionCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


