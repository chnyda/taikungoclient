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

// checks if the AwsBlockDeviceMappingsCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsBlockDeviceMappingsCommand{}

// AwsBlockDeviceMappingsCommand struct for AwsBlockDeviceMappingsCommand
type AwsBlockDeviceMappingsCommand struct {
	ProjectId *int32 `json:"projectId,omitempty"`
	ImageId NullableString `json:"imageId,omitempty"`
}

// NewAwsBlockDeviceMappingsCommand instantiates a new AwsBlockDeviceMappingsCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsBlockDeviceMappingsCommand() *AwsBlockDeviceMappingsCommand {
	this := AwsBlockDeviceMappingsCommand{}
	return &this
}

// NewAwsBlockDeviceMappingsCommandWithDefaults instantiates a new AwsBlockDeviceMappingsCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsBlockDeviceMappingsCommandWithDefaults() *AwsBlockDeviceMappingsCommand {
	this := AwsBlockDeviceMappingsCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *AwsBlockDeviceMappingsCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsBlockDeviceMappingsCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *AwsBlockDeviceMappingsCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *AwsBlockDeviceMappingsCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetImageId returns the ImageId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsBlockDeviceMappingsCommand) GetImageId() string {
	if o == nil || IsNil(o.ImageId.Get()) {
		var ret string
		return ret
	}
	return *o.ImageId.Get()
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsBlockDeviceMappingsCommand) GetImageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageId.Get(), o.ImageId.IsSet()
}

// HasImageId returns a boolean if a field has been set.
func (o *AwsBlockDeviceMappingsCommand) HasImageId() bool {
	if o != nil && o.ImageId.IsSet() {
		return true
	}

	return false
}

// SetImageId gets a reference to the given NullableString and assigns it to the ImageId field.
func (o *AwsBlockDeviceMappingsCommand) SetImageId(v string) {
	o.ImageId.Set(&v)
}
// SetImageIdNil sets the value for ImageId to be an explicit nil
func (o *AwsBlockDeviceMappingsCommand) SetImageIdNil() {
	o.ImageId.Set(nil)
}

// UnsetImageId ensures that no value is present for ImageId, not even an explicit nil
func (o *AwsBlockDeviceMappingsCommand) UnsetImageId() {
	o.ImageId.Unset()
}

func (o AwsBlockDeviceMappingsCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsBlockDeviceMappingsCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.ImageId.IsSet() {
		toSerialize["imageId"] = o.ImageId.Get()
	}
	return toSerialize, nil
}

type NullableAwsBlockDeviceMappingsCommand struct {
	value *AwsBlockDeviceMappingsCommand
	isSet bool
}

func (v NullableAwsBlockDeviceMappingsCommand) Get() *AwsBlockDeviceMappingsCommand {
	return v.value
}

func (v *NullableAwsBlockDeviceMappingsCommand) Set(val *AwsBlockDeviceMappingsCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsBlockDeviceMappingsCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsBlockDeviceMappingsCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsBlockDeviceMappingsCommand(val *AwsBlockDeviceMappingsCommand) *NullableAwsBlockDeviceMappingsCommand {
	return &NullableAwsBlockDeviceMappingsCommand{value: val, isSet: true}
}

func (v NullableAwsBlockDeviceMappingsCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsBlockDeviceMappingsCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


