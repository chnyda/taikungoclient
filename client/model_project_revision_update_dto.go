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

// checks if the ProjectRevisionUpdateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectRevisionUpdateDto{}

// ProjectRevisionUpdateDto struct for ProjectRevisionUpdateDto
type ProjectRevisionUpdateDto struct {
	CloudCredentialRevision *int32 `json:"cloudCredentialRevision,omitempty"`
	OpaProfileRevision *int32 `json:"opaProfileRevision,omitempty"`
	AccessProfileRevision *int32 `json:"accessProfileRevision,omitempty"`
}

// NewProjectRevisionUpdateDto instantiates a new ProjectRevisionUpdateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectRevisionUpdateDto() *ProjectRevisionUpdateDto {
	this := ProjectRevisionUpdateDto{}
	return &this
}

// NewProjectRevisionUpdateDtoWithDefaults instantiates a new ProjectRevisionUpdateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectRevisionUpdateDtoWithDefaults() *ProjectRevisionUpdateDto {
	this := ProjectRevisionUpdateDto{}
	return &this
}

// GetCloudCredentialRevision returns the CloudCredentialRevision field value if set, zero value otherwise.
func (o *ProjectRevisionUpdateDto) GetCloudCredentialRevision() int32 {
	if o == nil || IsNil(o.CloudCredentialRevision) {
		var ret int32
		return ret
	}
	return *o.CloudCredentialRevision
}

// GetCloudCredentialRevisionOk returns a tuple with the CloudCredentialRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRevisionUpdateDto) GetCloudCredentialRevisionOk() (*int32, bool) {
	if o == nil || IsNil(o.CloudCredentialRevision) {
		return nil, false
	}
	return o.CloudCredentialRevision, true
}

// HasCloudCredentialRevision returns a boolean if a field has been set.
func (o *ProjectRevisionUpdateDto) HasCloudCredentialRevision() bool {
	if o != nil && !IsNil(o.CloudCredentialRevision) {
		return true
	}

	return false
}

// SetCloudCredentialRevision gets a reference to the given int32 and assigns it to the CloudCredentialRevision field.
func (o *ProjectRevisionUpdateDto) SetCloudCredentialRevision(v int32) {
	o.CloudCredentialRevision = &v
}

// GetOpaProfileRevision returns the OpaProfileRevision field value if set, zero value otherwise.
func (o *ProjectRevisionUpdateDto) GetOpaProfileRevision() int32 {
	if o == nil || IsNil(o.OpaProfileRevision) {
		var ret int32
		return ret
	}
	return *o.OpaProfileRevision
}

// GetOpaProfileRevisionOk returns a tuple with the OpaProfileRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRevisionUpdateDto) GetOpaProfileRevisionOk() (*int32, bool) {
	if o == nil || IsNil(o.OpaProfileRevision) {
		return nil, false
	}
	return o.OpaProfileRevision, true
}

// HasOpaProfileRevision returns a boolean if a field has been set.
func (o *ProjectRevisionUpdateDto) HasOpaProfileRevision() bool {
	if o != nil && !IsNil(o.OpaProfileRevision) {
		return true
	}

	return false
}

// SetOpaProfileRevision gets a reference to the given int32 and assigns it to the OpaProfileRevision field.
func (o *ProjectRevisionUpdateDto) SetOpaProfileRevision(v int32) {
	o.OpaProfileRevision = &v
}

// GetAccessProfileRevision returns the AccessProfileRevision field value if set, zero value otherwise.
func (o *ProjectRevisionUpdateDto) GetAccessProfileRevision() int32 {
	if o == nil || IsNil(o.AccessProfileRevision) {
		var ret int32
		return ret
	}
	return *o.AccessProfileRevision
}

// GetAccessProfileRevisionOk returns a tuple with the AccessProfileRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRevisionUpdateDto) GetAccessProfileRevisionOk() (*int32, bool) {
	if o == nil || IsNil(o.AccessProfileRevision) {
		return nil, false
	}
	return o.AccessProfileRevision, true
}

// HasAccessProfileRevision returns a boolean if a field has been set.
func (o *ProjectRevisionUpdateDto) HasAccessProfileRevision() bool {
	if o != nil && !IsNil(o.AccessProfileRevision) {
		return true
	}

	return false
}

// SetAccessProfileRevision gets a reference to the given int32 and assigns it to the AccessProfileRevision field.
func (o *ProjectRevisionUpdateDto) SetAccessProfileRevision(v int32) {
	o.AccessProfileRevision = &v
}

func (o ProjectRevisionUpdateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectRevisionUpdateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudCredentialRevision) {
		toSerialize["cloudCredentialRevision"] = o.CloudCredentialRevision
	}
	if !IsNil(o.OpaProfileRevision) {
		toSerialize["opaProfileRevision"] = o.OpaProfileRevision
	}
	if !IsNil(o.AccessProfileRevision) {
		toSerialize["accessProfileRevision"] = o.AccessProfileRevision
	}
	return toSerialize, nil
}

type NullableProjectRevisionUpdateDto struct {
	value *ProjectRevisionUpdateDto
	isSet bool
}

func (v NullableProjectRevisionUpdateDto) Get() *ProjectRevisionUpdateDto {
	return v.value
}

func (v *NullableProjectRevisionUpdateDto) Set(val *ProjectRevisionUpdateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectRevisionUpdateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectRevisionUpdateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectRevisionUpdateDto(val *ProjectRevisionUpdateDto) *NullableProjectRevisionUpdateDto {
	return &NullableProjectRevisionUpdateDto{value: val, isSet: true}
}

func (v NullableProjectRevisionUpdateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectRevisionUpdateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


