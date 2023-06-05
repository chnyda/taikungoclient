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

// checks if the KubernetesProfilesEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesProfilesEntity{}

// KubernetesProfilesEntity struct for KubernetesProfilesEntity
type KubernetesProfilesEntity struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	TaikunLBEnabled *bool `json:"taikunLBEnabled,omitempty"`
}

// NewKubernetesProfilesEntity instantiates a new KubernetesProfilesEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesProfilesEntity() *KubernetesProfilesEntity {
	this := KubernetesProfilesEntity{}
	return &this
}

// NewKubernetesProfilesEntityWithDefaults instantiates a new KubernetesProfilesEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesProfilesEntityWithDefaults() *KubernetesProfilesEntity {
	this := KubernetesProfilesEntity{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KubernetesProfilesEntity) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesProfilesEntity) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KubernetesProfilesEntity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KubernetesProfilesEntity) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KubernetesProfilesEntity) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesProfilesEntity) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KubernetesProfilesEntity) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KubernetesProfilesEntity) SetName(v string) {
	o.Name = &v
}

// GetTaikunLBEnabled returns the TaikunLBEnabled field value if set, zero value otherwise.
func (o *KubernetesProfilesEntity) GetTaikunLBEnabled() bool {
	if o == nil || IsNil(o.TaikunLBEnabled) {
		var ret bool
		return ret
	}
	return *o.TaikunLBEnabled
}

// GetTaikunLBEnabledOk returns a tuple with the TaikunLBEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesProfilesEntity) GetTaikunLBEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.TaikunLBEnabled) {
		return nil, false
	}
	return o.TaikunLBEnabled, true
}

// HasTaikunLBEnabled returns a boolean if a field has been set.
func (o *KubernetesProfilesEntity) HasTaikunLBEnabled() bool {
	if o != nil && !IsNil(o.TaikunLBEnabled) {
		return true
	}

	return false
}

// SetTaikunLBEnabled gets a reference to the given bool and assigns it to the TaikunLBEnabled field.
func (o *KubernetesProfilesEntity) SetTaikunLBEnabled(v bool) {
	o.TaikunLBEnabled = &v
}

func (o KubernetesProfilesEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesProfilesEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TaikunLBEnabled) {
		toSerialize["taikunLBEnabled"] = o.TaikunLBEnabled
	}
	return toSerialize, nil
}

type NullableKubernetesProfilesEntity struct {
	value *KubernetesProfilesEntity
	isSet bool
}

func (v NullableKubernetesProfilesEntity) Get() *KubernetesProfilesEntity {
	return v.value
}

func (v *NullableKubernetesProfilesEntity) Set(val *KubernetesProfilesEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesProfilesEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesProfilesEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesProfilesEntity(val *KubernetesProfilesEntity) *NullableKubernetesProfilesEntity {
	return &NullableKubernetesProfilesEntity{value: val, isSet: true}
}

func (v NullableKubernetesProfilesEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesProfilesEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


