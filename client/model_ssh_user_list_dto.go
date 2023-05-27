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

// checks if the SshUserListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SshUserListDto{}

// SshUserListDto struct for SshUserListDto
type SshUserListDto struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	SshPublicKey NullableString `json:"sshPublicKey,omitempty"`
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// NewSshUserListDto instantiates a new SshUserListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSshUserListDto() *SshUserListDto {
	this := SshUserListDto{}
	return &this
}

// NewSshUserListDtoWithDefaults instantiates a new SshUserListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSshUserListDtoWithDefaults() *SshUserListDto {
	this := SshUserListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SshUserListDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshUserListDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SshUserListDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *SshUserListDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SshUserListDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SshUserListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SshUserListDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *SshUserListDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *SshUserListDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *SshUserListDto) UnsetName() {
	o.Name.Unset()
}

// GetSshPublicKey returns the SshPublicKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SshUserListDto) GetSshPublicKey() string {
	if o == nil || IsNil(o.SshPublicKey.Get()) {
		var ret string
		return ret
	}
	return *o.SshPublicKey.Get()
}

// GetSshPublicKeyOk returns a tuple with the SshPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SshUserListDto) GetSshPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SshPublicKey.Get(), o.SshPublicKey.IsSet()
}

// HasSshPublicKey returns a boolean if a field has been set.
func (o *SshUserListDto) HasSshPublicKey() bool {
	if o != nil && o.SshPublicKey.IsSet() {
		return true
	}

	return false
}

// SetSshPublicKey gets a reference to the given NullableString and assigns it to the SshPublicKey field.
func (o *SshUserListDto) SetSshPublicKey(v string) {
	o.SshPublicKey.Set(&v)
}
// SetSshPublicKeyNil sets the value for SshPublicKey to be an explicit nil
func (o *SshUserListDto) SetSshPublicKeyNil() {
	o.SshPublicKey.Set(nil)
}

// UnsetSshPublicKey ensures that no value is present for SshPublicKey, not even an explicit nil
func (o *SshUserListDto) UnsetSshPublicKey() {
	o.SshPublicKey.Unset()
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *SshUserListDto) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshUserListDto) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *SshUserListDto) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *SshUserListDto) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

func (o SshUserListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SshUserListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.SshPublicKey.IsSet() {
		toSerialize["sshPublicKey"] = o.SshPublicKey.Get()
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["isDeleted"] = o.IsDeleted
	}
	return toSerialize, nil
}

type NullableSshUserListDto struct {
	value *SshUserListDto
	isSet bool
}

func (v NullableSshUserListDto) Get() *SshUserListDto {
	return v.value
}

func (v *NullableSshUserListDto) Set(val *SshUserListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSshUserListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSshUserListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSshUserListDto(val *SshUserListDto) *NullableSshUserListDto {
	return &NullableSshUserListDto{value: val, isSet: true}
}

func (v NullableSshUserListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSshUserListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


