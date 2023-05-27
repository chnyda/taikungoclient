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

// checks if the DeleteBackupStorageLocationCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteBackupStorageLocationCommand{}

// DeleteBackupStorageLocationCommand struct for DeleteBackupStorageLocationCommand
type DeleteBackupStorageLocationCommand struct {
	ProjectId *int32 `json:"projectId,omitempty"`
	StorageLocation NullableString `json:"storageLocation,omitempty"`
}

// NewDeleteBackupStorageLocationCommand instantiates a new DeleteBackupStorageLocationCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteBackupStorageLocationCommand() *DeleteBackupStorageLocationCommand {
	this := DeleteBackupStorageLocationCommand{}
	return &this
}

// NewDeleteBackupStorageLocationCommandWithDefaults instantiates a new DeleteBackupStorageLocationCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteBackupStorageLocationCommandWithDefaults() *DeleteBackupStorageLocationCommand {
	this := DeleteBackupStorageLocationCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *DeleteBackupStorageLocationCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteBackupStorageLocationCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *DeleteBackupStorageLocationCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *DeleteBackupStorageLocationCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetStorageLocation returns the StorageLocation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeleteBackupStorageLocationCommand) GetStorageLocation() string {
	if o == nil || IsNil(o.StorageLocation.Get()) {
		var ret string
		return ret
	}
	return *o.StorageLocation.Get()
}

// GetStorageLocationOk returns a tuple with the StorageLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteBackupStorageLocationCommand) GetStorageLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageLocation.Get(), o.StorageLocation.IsSet()
}

// HasStorageLocation returns a boolean if a field has been set.
func (o *DeleteBackupStorageLocationCommand) HasStorageLocation() bool {
	if o != nil && o.StorageLocation.IsSet() {
		return true
	}

	return false
}

// SetStorageLocation gets a reference to the given NullableString and assigns it to the StorageLocation field.
func (o *DeleteBackupStorageLocationCommand) SetStorageLocation(v string) {
	o.StorageLocation.Set(&v)
}
// SetStorageLocationNil sets the value for StorageLocation to be an explicit nil
func (o *DeleteBackupStorageLocationCommand) SetStorageLocationNil() {
	o.StorageLocation.Set(nil)
}

// UnsetStorageLocation ensures that no value is present for StorageLocation, not even an explicit nil
func (o *DeleteBackupStorageLocationCommand) UnsetStorageLocation() {
	o.StorageLocation.Unset()
}

func (o DeleteBackupStorageLocationCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteBackupStorageLocationCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.StorageLocation.IsSet() {
		toSerialize["storageLocation"] = o.StorageLocation.Get()
	}
	return toSerialize, nil
}

type NullableDeleteBackupStorageLocationCommand struct {
	value *DeleteBackupStorageLocationCommand
	isSet bool
}

func (v NullableDeleteBackupStorageLocationCommand) Get() *DeleteBackupStorageLocationCommand {
	return v.value
}

func (v *NullableDeleteBackupStorageLocationCommand) Set(val *DeleteBackupStorageLocationCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteBackupStorageLocationCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteBackupStorageLocationCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteBackupStorageLocationCommand(val *DeleteBackupStorageLocationCommand) *NullableDeleteBackupStorageLocationCommand {
	return &NullableDeleteBackupStorageLocationCommand{value: val, isSet: true}
}

func (v NullableDeleteBackupStorageLocationCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteBackupStorageLocationCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


