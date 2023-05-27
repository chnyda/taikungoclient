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

// checks if the KubesprayListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubesprayListDto{}

// KubesprayListDto struct for KubesprayListDto
type KubesprayListDto struct {
	Id *int32 `json:"id,omitempty"`
	Version NullableString `json:"version,omitempty"`
	KubernetesVersion NullableString `json:"kubernetesVersion,omitempty"`
	IsDeprecated *bool `json:"isDeprecated,omitempty"`
}

// NewKubesprayListDto instantiates a new KubesprayListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubesprayListDto() *KubesprayListDto {
	this := KubesprayListDto{}
	return &this
}

// NewKubesprayListDtoWithDefaults instantiates a new KubesprayListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubesprayListDtoWithDefaults() *KubesprayListDto {
	this := KubesprayListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KubesprayListDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubesprayListDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KubesprayListDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KubesprayListDto) SetId(v int32) {
	o.Id = &v
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubesprayListDto) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubesprayListDto) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *KubesprayListDto) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *KubesprayListDto) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *KubesprayListDto) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *KubesprayListDto) UnsetVersion() {
	o.Version.Unset()
}

// GetKubernetesVersion returns the KubernetesVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubesprayListDto) GetKubernetesVersion() string {
	if o == nil || IsNil(o.KubernetesVersion.Get()) {
		var ret string
		return ret
	}
	return *o.KubernetesVersion.Get()
}

// GetKubernetesVersionOk returns a tuple with the KubernetesVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubesprayListDto) GetKubernetesVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KubernetesVersion.Get(), o.KubernetesVersion.IsSet()
}

// HasKubernetesVersion returns a boolean if a field has been set.
func (o *KubesprayListDto) HasKubernetesVersion() bool {
	if o != nil && o.KubernetesVersion.IsSet() {
		return true
	}

	return false
}

// SetKubernetesVersion gets a reference to the given NullableString and assigns it to the KubernetesVersion field.
func (o *KubesprayListDto) SetKubernetesVersion(v string) {
	o.KubernetesVersion.Set(&v)
}
// SetKubernetesVersionNil sets the value for KubernetesVersion to be an explicit nil
func (o *KubesprayListDto) SetKubernetesVersionNil() {
	o.KubernetesVersion.Set(nil)
}

// UnsetKubernetesVersion ensures that no value is present for KubernetesVersion, not even an explicit nil
func (o *KubesprayListDto) UnsetKubernetesVersion() {
	o.KubernetesVersion.Unset()
}

// GetIsDeprecated returns the IsDeprecated field value if set, zero value otherwise.
func (o *KubesprayListDto) GetIsDeprecated() bool {
	if o == nil || IsNil(o.IsDeprecated) {
		var ret bool
		return ret
	}
	return *o.IsDeprecated
}

// GetIsDeprecatedOk returns a tuple with the IsDeprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubesprayListDto) GetIsDeprecatedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeprecated) {
		return nil, false
	}
	return o.IsDeprecated, true
}

// HasIsDeprecated returns a boolean if a field has been set.
func (o *KubesprayListDto) HasIsDeprecated() bool {
	if o != nil && !IsNil(o.IsDeprecated) {
		return true
	}

	return false
}

// SetIsDeprecated gets a reference to the given bool and assigns it to the IsDeprecated field.
func (o *KubesprayListDto) SetIsDeprecated(v bool) {
	o.IsDeprecated = &v
}

func (o KubesprayListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubesprayListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if o.KubernetesVersion.IsSet() {
		toSerialize["kubernetesVersion"] = o.KubernetesVersion.Get()
	}
	if !IsNil(o.IsDeprecated) {
		toSerialize["isDeprecated"] = o.IsDeprecated
	}
	return toSerialize, nil
}

type NullableKubesprayListDto struct {
	value *KubesprayListDto
	isSet bool
}

func (v NullableKubesprayListDto) Get() *KubesprayListDto {
	return v.value
}

func (v *NullableKubesprayListDto) Set(val *KubesprayListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableKubesprayListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableKubesprayListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubesprayListDto(val *KubesprayListDto) *NullableKubesprayListDto {
	return &NullableKubesprayListDto{value: val, isSet: true}
}

func (v NullableKubesprayListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubesprayListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


