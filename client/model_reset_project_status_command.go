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

// checks if the ResetProjectStatusCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResetProjectStatusCommand{}

// ResetProjectStatusCommand struct for ResetProjectStatusCommand
type ResetProjectStatusCommand struct {
	ProjectId *int32 `json:"projectId,omitempty"`
	Status *ProjectStatus `json:"status,omitempty"`
}

// NewResetProjectStatusCommand instantiates a new ResetProjectStatusCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResetProjectStatusCommand() *ResetProjectStatusCommand {
	this := ResetProjectStatusCommand{}
	return &this
}

// NewResetProjectStatusCommandWithDefaults instantiates a new ResetProjectStatusCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResetProjectStatusCommandWithDefaults() *ResetProjectStatusCommand {
	this := ResetProjectStatusCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *ResetProjectStatusCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetProjectStatusCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ResetProjectStatusCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *ResetProjectStatusCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ResetProjectStatusCommand) GetStatus() ProjectStatus {
	if o == nil || IsNil(o.Status) {
		var ret ProjectStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetProjectStatusCommand) GetStatusOk() (*ProjectStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ResetProjectStatusCommand) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ProjectStatus and assigns it to the Status field.
func (o *ResetProjectStatusCommand) SetStatus(v ProjectStatus) {
	o.Status = &v
}

func (o ResetProjectStatusCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResetProjectStatusCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableResetProjectStatusCommand struct {
	value *ResetProjectStatusCommand
	isSet bool
}

func (v NullableResetProjectStatusCommand) Get() *ResetProjectStatusCommand {
	return v.value
}

func (v *NullableResetProjectStatusCommand) Set(val *ResetProjectStatusCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableResetProjectStatusCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableResetProjectStatusCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResetProjectStatusCommand(val *ResetProjectStatusCommand) *NullableResetProjectStatusCommand {
	return &NullableResetProjectStatusCommand{value: val, isSet: true}
}

func (v NullableResetProjectStatusCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResetProjectStatusCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


