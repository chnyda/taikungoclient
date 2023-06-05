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

// checks if the DeleteScheduleCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteScheduleCommand{}

// DeleteScheduleCommand struct for DeleteScheduleCommand
type DeleteScheduleCommand struct {
	ProjectId *int32 `json:"projectId,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewDeleteScheduleCommand instantiates a new DeleteScheduleCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteScheduleCommand() *DeleteScheduleCommand {
	this := DeleteScheduleCommand{}
	return &this
}

// NewDeleteScheduleCommandWithDefaults instantiates a new DeleteScheduleCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteScheduleCommandWithDefaults() *DeleteScheduleCommand {
	this := DeleteScheduleCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *DeleteScheduleCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteScheduleCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *DeleteScheduleCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *DeleteScheduleCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeleteScheduleCommand) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteScheduleCommand) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeleteScheduleCommand) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeleteScheduleCommand) SetName(v string) {
	o.Name = &v
}

func (o DeleteScheduleCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteScheduleCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableDeleteScheduleCommand struct {
	value *DeleteScheduleCommand
	isSet bool
}

func (v NullableDeleteScheduleCommand) Get() *DeleteScheduleCommand {
	return v.value
}

func (v *NullableDeleteScheduleCommand) Set(val *DeleteScheduleCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteScheduleCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteScheduleCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteScheduleCommand(val *DeleteScheduleCommand) *NullableDeleteScheduleCommand {
	return &NullableDeleteScheduleCommand{value: val, isSet: true}
}

func (v NullableDeleteScheduleCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteScheduleCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


