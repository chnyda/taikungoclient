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

// checks if the DaemonSetDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DaemonSetDto{}

// DaemonSetDto struct for DaemonSetDto
type DaemonSetDto struct {
	MetadataName *string `json:"metadataName,omitempty"`
	Status *string `json:"status,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	Age *string `json:"age,omitempty"`
}

// NewDaemonSetDto instantiates a new DaemonSetDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDaemonSetDto() *DaemonSetDto {
	this := DaemonSetDto{}
	return &this
}

// NewDaemonSetDtoWithDefaults instantiates a new DaemonSetDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDaemonSetDtoWithDefaults() *DaemonSetDto {
	this := DaemonSetDto{}
	return &this
}

// GetMetadataName returns the MetadataName field value if set, zero value otherwise.
func (o *DaemonSetDto) GetMetadataName() string {
	if o == nil || IsNil(o.MetadataName) {
		var ret string
		return ret
	}
	return *o.MetadataName
}

// GetMetadataNameOk returns a tuple with the MetadataName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaemonSetDto) GetMetadataNameOk() (*string, bool) {
	if o == nil || IsNil(o.MetadataName) {
		return nil, false
	}
	return o.MetadataName, true
}

// HasMetadataName returns a boolean if a field has been set.
func (o *DaemonSetDto) HasMetadataName() bool {
	if o != nil && !IsNil(o.MetadataName) {
		return true
	}

	return false
}

// SetMetadataName gets a reference to the given string and assigns it to the MetadataName field.
func (o *DaemonSetDto) SetMetadataName(v string) {
	o.MetadataName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DaemonSetDto) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaemonSetDto) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DaemonSetDto) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DaemonSetDto) SetStatus(v string) {
	o.Status = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *DaemonSetDto) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaemonSetDto) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *DaemonSetDto) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *DaemonSetDto) SetNamespace(v string) {
	o.Namespace = &v
}

// GetAge returns the Age field value if set, zero value otherwise.
func (o *DaemonSetDto) GetAge() string {
	if o == nil || IsNil(o.Age) {
		var ret string
		return ret
	}
	return *o.Age
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaemonSetDto) GetAgeOk() (*string, bool) {
	if o == nil || IsNil(o.Age) {
		return nil, false
	}
	return o.Age, true
}

// HasAge returns a boolean if a field has been set.
func (o *DaemonSetDto) HasAge() bool {
	if o != nil && !IsNil(o.Age) {
		return true
	}

	return false
}

// SetAge gets a reference to the given string and assigns it to the Age field.
func (o *DaemonSetDto) SetAge(v string) {
	o.Age = &v
}

func (o DaemonSetDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DaemonSetDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MetadataName) {
		toSerialize["metadataName"] = o.MetadataName
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !IsNil(o.Age) {
		toSerialize["age"] = o.Age
	}
	return toSerialize, nil
}

type NullableDaemonSetDto struct {
	value *DaemonSetDto
	isSet bool
}

func (v NullableDaemonSetDto) Get() *DaemonSetDto {
	return v.value
}

func (v *NullableDaemonSetDto) Set(val *DaemonSetDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDaemonSetDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDaemonSetDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDaemonSetDto(val *DaemonSetDto) *NullableDaemonSetDto {
	return &NullableDaemonSetDto{value: val, isSet: true}
}

func (v NullableDaemonSetDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDaemonSetDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


