/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikuncore

import (
	"encoding/json"
)

// checks if the DeleteUserGroupCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteUserGroupCommand{}

// DeleteUserGroupCommand struct for DeleteUserGroupCommand
type DeleteUserGroupCommand struct {
	UserGroupIds []int32 `json:"userGroupIds,omitempty"`
}

// NewDeleteUserGroupCommand instantiates a new DeleteUserGroupCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteUserGroupCommand() *DeleteUserGroupCommand {
	this := DeleteUserGroupCommand{}
	return &this
}

// NewDeleteUserGroupCommandWithDefaults instantiates a new DeleteUserGroupCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteUserGroupCommandWithDefaults() *DeleteUserGroupCommand {
	this := DeleteUserGroupCommand{}
	return &this
}

// GetUserGroupIds returns the UserGroupIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeleteUserGroupCommand) GetUserGroupIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.UserGroupIds
}

// GetUserGroupIdsOk returns a tuple with the UserGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteUserGroupCommand) GetUserGroupIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.UserGroupIds) {
		return nil, false
	}
	return o.UserGroupIds, true
}

// HasUserGroupIds returns a boolean if a field has been set.
func (o *DeleteUserGroupCommand) HasUserGroupIds() bool {
	if o != nil && IsNil(o.UserGroupIds) {
		return true
	}

	return false
}

// SetUserGroupIds gets a reference to the given []int32 and assigns it to the UserGroupIds field.
func (o *DeleteUserGroupCommand) SetUserGroupIds(v []int32) {
	o.UserGroupIds = v
}

func (o DeleteUserGroupCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteUserGroupCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.UserGroupIds != nil {
		toSerialize["userGroupIds"] = o.UserGroupIds
	}
	return toSerialize, nil
}

type NullableDeleteUserGroupCommand struct {
	value *DeleteUserGroupCommand
	isSet bool
}

func (v NullableDeleteUserGroupCommand) Get() *DeleteUserGroupCommand {
	return v.value
}

func (v *NullableDeleteUserGroupCommand) Set(val *DeleteUserGroupCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteUserGroupCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteUserGroupCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteUserGroupCommand(val *DeleteUserGroupCommand) *NullableDeleteUserGroupCommand {
	return &NullableDeleteUserGroupCommand{value: val, isSet: true}
}

func (v NullableDeleteUserGroupCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteUserGroupCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

