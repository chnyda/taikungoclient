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

// checks if the ProjectCommonRecordDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectCommonRecordDto{}

// ProjectCommonRecordDto struct for ProjectCommonRecordDto
type ProjectCommonRecordDto struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	ExpiredAt *string `json:"expiredAt,omitempty"`
}

// NewProjectCommonRecordDto instantiates a new ProjectCommonRecordDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectCommonRecordDto() *ProjectCommonRecordDto {
	this := ProjectCommonRecordDto{}
	return &this
}

// NewProjectCommonRecordDtoWithDefaults instantiates a new ProjectCommonRecordDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectCommonRecordDtoWithDefaults() *ProjectCommonRecordDto {
	this := ProjectCommonRecordDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProjectCommonRecordDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCommonRecordDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProjectCommonRecordDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ProjectCommonRecordDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProjectCommonRecordDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCommonRecordDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProjectCommonRecordDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProjectCommonRecordDto) SetName(v string) {
	o.Name = &v
}

// GetExpiredAt returns the ExpiredAt field value if set, zero value otherwise.
func (o *ProjectCommonRecordDto) GetExpiredAt() string {
	if o == nil || IsNil(o.ExpiredAt) {
		var ret string
		return ret
	}
	return *o.ExpiredAt
}

// GetExpiredAtOk returns a tuple with the ExpiredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCommonRecordDto) GetExpiredAtOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiredAt) {
		return nil, false
	}
	return o.ExpiredAt, true
}

// HasExpiredAt returns a boolean if a field has been set.
func (o *ProjectCommonRecordDto) HasExpiredAt() bool {
	if o != nil && !IsNil(o.ExpiredAt) {
		return true
	}

	return false
}

// SetExpiredAt gets a reference to the given string and assigns it to the ExpiredAt field.
func (o *ProjectCommonRecordDto) SetExpiredAt(v string) {
	o.ExpiredAt = &v
}

func (o ProjectCommonRecordDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectCommonRecordDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ExpiredAt) {
		toSerialize["expiredAt"] = o.ExpiredAt
	}
	return toSerialize, nil
}

type NullableProjectCommonRecordDto struct {
	value *ProjectCommonRecordDto
	isSet bool
}

func (v NullableProjectCommonRecordDto) Get() *ProjectCommonRecordDto {
	return v.value
}

func (v *NullableProjectCommonRecordDto) Set(val *ProjectCommonRecordDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectCommonRecordDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectCommonRecordDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectCommonRecordDto(val *ProjectCommonRecordDto) *NullableProjectCommonRecordDto {
	return &NullableProjectCommonRecordDto{value: val, isSet: true}
}

func (v NullableProjectCommonRecordDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectCommonRecordDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


