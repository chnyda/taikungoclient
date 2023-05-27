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

// checks if the DeleteBackupCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteBackupCommand{}

// DeleteBackupCommand struct for DeleteBackupCommand
type DeleteBackupCommand struct {
	ProjectId *int32 `json:"projectId,omitempty"`
	Name NullableString `json:"name,omitempty"`
}

// NewDeleteBackupCommand instantiates a new DeleteBackupCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteBackupCommand() *DeleteBackupCommand {
	this := DeleteBackupCommand{}
	return &this
}

// NewDeleteBackupCommandWithDefaults instantiates a new DeleteBackupCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteBackupCommandWithDefaults() *DeleteBackupCommand {
	this := DeleteBackupCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *DeleteBackupCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteBackupCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *DeleteBackupCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *DeleteBackupCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeleteBackupCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteBackupCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DeleteBackupCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DeleteBackupCommand) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *DeleteBackupCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DeleteBackupCommand) UnsetName() {
	o.Name.Unset()
}

func (o DeleteBackupCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteBackupCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullableDeleteBackupCommand struct {
	value *DeleteBackupCommand
	isSet bool
}

func (v NullableDeleteBackupCommand) Get() *DeleteBackupCommand {
	return v.value
}

func (v *NullableDeleteBackupCommand) Set(val *DeleteBackupCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteBackupCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteBackupCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteBackupCommand(val *DeleteBackupCommand) *NullableDeleteBackupCommand {
	return &NullableDeleteBackupCommand{value: val, isSet: true}
}

func (v NullableDeleteBackupCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteBackupCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


