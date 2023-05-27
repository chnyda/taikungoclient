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

// checks if the NotificationSendCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationSendCommand{}

// NotificationSendCommand struct for NotificationSendCommand
type NotificationSendCommand struct {
	ProjectId *int32 `json:"projectId,omitempty"`
	ActionType *ActionType `json:"actionType,omitempty"`
	ActionStatus *ActionStatus `json:"actionStatus,omitempty"`
	ProjectType *ProjectType `json:"projectType,omitempty"`
}

// NewNotificationSendCommand instantiates a new NotificationSendCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationSendCommand() *NotificationSendCommand {
	this := NotificationSendCommand{}
	return &this
}

// NewNotificationSendCommandWithDefaults instantiates a new NotificationSendCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationSendCommandWithDefaults() *NotificationSendCommand {
	this := NotificationSendCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *NotificationSendCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSendCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *NotificationSendCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *NotificationSendCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetActionType returns the ActionType field value if set, zero value otherwise.
func (o *NotificationSendCommand) GetActionType() ActionType {
	if o == nil || IsNil(o.ActionType) {
		var ret ActionType
		return ret
	}
	return *o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSendCommand) GetActionTypeOk() (*ActionType, bool) {
	if o == nil || IsNil(o.ActionType) {
		return nil, false
	}
	return o.ActionType, true
}

// HasActionType returns a boolean if a field has been set.
func (o *NotificationSendCommand) HasActionType() bool {
	if o != nil && !IsNil(o.ActionType) {
		return true
	}

	return false
}

// SetActionType gets a reference to the given ActionType and assigns it to the ActionType field.
func (o *NotificationSendCommand) SetActionType(v ActionType) {
	o.ActionType = &v
}

// GetActionStatus returns the ActionStatus field value if set, zero value otherwise.
func (o *NotificationSendCommand) GetActionStatus() ActionStatus {
	if o == nil || IsNil(o.ActionStatus) {
		var ret ActionStatus
		return ret
	}
	return *o.ActionStatus
}

// GetActionStatusOk returns a tuple with the ActionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSendCommand) GetActionStatusOk() (*ActionStatus, bool) {
	if o == nil || IsNil(o.ActionStatus) {
		return nil, false
	}
	return o.ActionStatus, true
}

// HasActionStatus returns a boolean if a field has been set.
func (o *NotificationSendCommand) HasActionStatus() bool {
	if o != nil && !IsNil(o.ActionStatus) {
		return true
	}

	return false
}

// SetActionStatus gets a reference to the given ActionStatus and assigns it to the ActionStatus field.
func (o *NotificationSendCommand) SetActionStatus(v ActionStatus) {
	o.ActionStatus = &v
}

// GetProjectType returns the ProjectType field value if set, zero value otherwise.
func (o *NotificationSendCommand) GetProjectType() ProjectType {
	if o == nil || IsNil(o.ProjectType) {
		var ret ProjectType
		return ret
	}
	return *o.ProjectType
}

// GetProjectTypeOk returns a tuple with the ProjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSendCommand) GetProjectTypeOk() (*ProjectType, bool) {
	if o == nil || IsNil(o.ProjectType) {
		return nil, false
	}
	return o.ProjectType, true
}

// HasProjectType returns a boolean if a field has been set.
func (o *NotificationSendCommand) HasProjectType() bool {
	if o != nil && !IsNil(o.ProjectType) {
		return true
	}

	return false
}

// SetProjectType gets a reference to the given ProjectType and assigns it to the ProjectType field.
func (o *NotificationSendCommand) SetProjectType(v ProjectType) {
	o.ProjectType = &v
}

func (o NotificationSendCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationSendCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.ActionType) {
		toSerialize["actionType"] = o.ActionType
	}
	if !IsNil(o.ActionStatus) {
		toSerialize["actionStatus"] = o.ActionStatus
	}
	if !IsNil(o.ProjectType) {
		toSerialize["projectType"] = o.ProjectType
	}
	return toSerialize, nil
}

type NullableNotificationSendCommand struct {
	value *NotificationSendCommand
	isSet bool
}

func (v NullableNotificationSendCommand) Get() *NotificationSendCommand {
	return v.value
}

func (v *NullableNotificationSendCommand) Set(val *NotificationSendCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationSendCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationSendCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationSendCommand(val *NotificationSendCommand) *NullableNotificationSendCommand {
	return &NullableNotificationSendCommand{value: val, isSet: true}
}

func (v NullableNotificationSendCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationSendCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


