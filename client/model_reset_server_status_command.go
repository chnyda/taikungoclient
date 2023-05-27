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

// checks if the ResetServerStatusCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResetServerStatusCommand{}

// ResetServerStatusCommand struct for ResetServerStatusCommand
type ResetServerStatusCommand struct {
	ProjectId int32 `json:"projectId"`
	ServerIds []int32 `json:"serverIds,omitempty"`
	Status *CloudStatus `json:"status,omitempty"`
}

// NewResetServerStatusCommand instantiates a new ResetServerStatusCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResetServerStatusCommand(projectId int32) *ResetServerStatusCommand {
	this := ResetServerStatusCommand{}
	this.ProjectId = projectId
	return &this
}

// NewResetServerStatusCommandWithDefaults instantiates a new ResetServerStatusCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResetServerStatusCommandWithDefaults() *ResetServerStatusCommand {
	this := ResetServerStatusCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *ResetServerStatusCommand) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *ResetServerStatusCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *ResetServerStatusCommand) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetServerIds returns the ServerIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResetServerStatusCommand) GetServerIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.ServerIds
}

// GetServerIdsOk returns a tuple with the ServerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResetServerStatusCommand) GetServerIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.ServerIds) {
		return nil, false
	}
	return o.ServerIds, true
}

// HasServerIds returns a boolean if a field has been set.
func (o *ResetServerStatusCommand) HasServerIds() bool {
	if o != nil && IsNil(o.ServerIds) {
		return true
	}

	return false
}

// SetServerIds gets a reference to the given []int32 and assigns it to the ServerIds field.
func (o *ResetServerStatusCommand) SetServerIds(v []int32) {
	o.ServerIds = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ResetServerStatusCommand) GetStatus() CloudStatus {
	if o == nil || IsNil(o.Status) {
		var ret CloudStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetServerStatusCommand) GetStatusOk() (*CloudStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ResetServerStatusCommand) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given CloudStatus and assigns it to the Status field.
func (o *ResetServerStatusCommand) SetStatus(v CloudStatus) {
	o.Status = &v
}

func (o ResetServerStatusCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResetServerStatusCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	if o.ServerIds != nil {
		toSerialize["serverIds"] = o.ServerIds
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableResetServerStatusCommand struct {
	value *ResetServerStatusCommand
	isSet bool
}

func (v NullableResetServerStatusCommand) Get() *ResetServerStatusCommand {
	return v.value
}

func (v *NullableResetServerStatusCommand) Set(val *ResetServerStatusCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableResetServerStatusCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableResetServerStatusCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResetServerStatusCommand(val *ResetServerStatusCommand) *NullableResetServerStatusCommand {
	return &NullableResetServerStatusCommand{value: val, isSet: true}
}

func (v NullableResetServerStatusCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResetServerStatusCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


