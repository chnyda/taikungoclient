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
	"time"
)

// checks if the CDeleteBackupRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CDeleteBackupRequestDto{}

// CDeleteBackupRequestDto struct for CDeleteBackupRequestDto
type CDeleteBackupRequestDto struct {
	MetadataName *string `json:"metadataName,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	BackupName *string `json:"backupName,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	Phase *string `json:"phase,omitempty"`
}

// NewCDeleteBackupRequestDto instantiates a new CDeleteBackupRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCDeleteBackupRequestDto() *CDeleteBackupRequestDto {
	this := CDeleteBackupRequestDto{}
	return &this
}

// NewCDeleteBackupRequestDtoWithDefaults instantiates a new CDeleteBackupRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCDeleteBackupRequestDtoWithDefaults() *CDeleteBackupRequestDto {
	this := CDeleteBackupRequestDto{}
	return &this
}

// GetMetadataName returns the MetadataName field value if set, zero value otherwise.
func (o *CDeleteBackupRequestDto) GetMetadataName() string {
	if o == nil || IsNil(o.MetadataName) {
		var ret string
		return ret
	}
	return *o.MetadataName
}

// GetMetadataNameOk returns a tuple with the MetadataName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CDeleteBackupRequestDto) GetMetadataNameOk() (*string, bool) {
	if o == nil || IsNil(o.MetadataName) {
		return nil, false
	}
	return o.MetadataName, true
}

// HasMetadataName returns a boolean if a field has been set.
func (o *CDeleteBackupRequestDto) HasMetadataName() bool {
	if o != nil && !IsNil(o.MetadataName) {
		return true
	}

	return false
}

// SetMetadataName gets a reference to the given string and assigns it to the MetadataName field.
func (o *CDeleteBackupRequestDto) SetMetadataName(v string) {
	o.MetadataName = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CDeleteBackupRequestDto) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CDeleteBackupRequestDto) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CDeleteBackupRequestDto) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CDeleteBackupRequestDto) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetBackupName returns the BackupName field value if set, zero value otherwise.
func (o *CDeleteBackupRequestDto) GetBackupName() string {
	if o == nil || IsNil(o.BackupName) {
		var ret string
		return ret
	}
	return *o.BackupName
}

// GetBackupNameOk returns a tuple with the BackupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CDeleteBackupRequestDto) GetBackupNameOk() (*string, bool) {
	if o == nil || IsNil(o.BackupName) {
		return nil, false
	}
	return o.BackupName, true
}

// HasBackupName returns a boolean if a field has been set.
func (o *CDeleteBackupRequestDto) HasBackupName() bool {
	if o != nil && !IsNil(o.BackupName) {
		return true
	}

	return false
}

// SetBackupName gets a reference to the given string and assigns it to the BackupName field.
func (o *CDeleteBackupRequestDto) SetBackupName(v string) {
	o.BackupName = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *CDeleteBackupRequestDto) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CDeleteBackupRequestDto) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *CDeleteBackupRequestDto) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *CDeleteBackupRequestDto) SetNamespace(v string) {
	o.Namespace = &v
}

// GetPhase returns the Phase field value if set, zero value otherwise.
func (o *CDeleteBackupRequestDto) GetPhase() string {
	if o == nil || IsNil(o.Phase) {
		var ret string
		return ret
	}
	return *o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CDeleteBackupRequestDto) GetPhaseOk() (*string, bool) {
	if o == nil || IsNil(o.Phase) {
		return nil, false
	}
	return o.Phase, true
}

// HasPhase returns a boolean if a field has been set.
func (o *CDeleteBackupRequestDto) HasPhase() bool {
	if o != nil && !IsNil(o.Phase) {
		return true
	}

	return false
}

// SetPhase gets a reference to the given string and assigns it to the Phase field.
func (o *CDeleteBackupRequestDto) SetPhase(v string) {
	o.Phase = &v
}

func (o CDeleteBackupRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CDeleteBackupRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MetadataName) {
		toSerialize["metadataName"] = o.MetadataName
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.BackupName) {
		toSerialize["backupName"] = o.BackupName
	}
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !IsNil(o.Phase) {
		toSerialize["phase"] = o.Phase
	}
	return toSerialize, nil
}

type NullableCDeleteBackupRequestDto struct {
	value *CDeleteBackupRequestDto
	isSet bool
}

func (v NullableCDeleteBackupRequestDto) Get() *CDeleteBackupRequestDto {
	return v.value
}

func (v *NullableCDeleteBackupRequestDto) Set(val *CDeleteBackupRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCDeleteBackupRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCDeleteBackupRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCDeleteBackupRequestDto(val *CDeleteBackupRequestDto) *NullableCDeleteBackupRequestDto {
	return &NullableCDeleteBackupRequestDto{value: val, isSet: true}
}

func (v NullableCDeleteBackupRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCDeleteBackupRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


