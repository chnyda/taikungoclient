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

// checks if the DeleteAlertCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteAlertCommand{}

// DeleteAlertCommand struct for DeleteAlertCommand
type DeleteAlertCommand struct {
	ProjectId int32 `json:"projectId"`
}

// NewDeleteAlertCommand instantiates a new DeleteAlertCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteAlertCommand(projectId int32) *DeleteAlertCommand {
	this := DeleteAlertCommand{}
	this.ProjectId = projectId
	return &this
}

// NewDeleteAlertCommandWithDefaults instantiates a new DeleteAlertCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteAlertCommandWithDefaults() *DeleteAlertCommand {
	this := DeleteAlertCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *DeleteAlertCommand) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *DeleteAlertCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *DeleteAlertCommand) SetProjectId(v int32) {
	o.ProjectId = v
}

func (o DeleteAlertCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteAlertCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

type NullableDeleteAlertCommand struct {
	value *DeleteAlertCommand
	isSet bool
}

func (v NullableDeleteAlertCommand) Get() *DeleteAlertCommand {
	return v.value
}

func (v *NullableDeleteAlertCommand) Set(val *DeleteAlertCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteAlertCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteAlertCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteAlertCommand(val *DeleteAlertCommand) *NullableDeleteAlertCommand {
	return &NullableDeleteAlertCommand{value: val, isSet: true}
}

func (v NullableDeleteAlertCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteAlertCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


