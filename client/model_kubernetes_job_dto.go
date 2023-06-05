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

// checks if the KubernetesJobDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesJobDto{}

// KubernetesJobDto struct for KubernetesJobDto
type KubernetesJobDto struct {
	MetadataName *string `json:"metadataName,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	Age *string `json:"age,omitempty"`
	Completions *int32 `json:"completions,omitempty"`
	Conditions *string `json:"conditions,omitempty"`
}

// NewKubernetesJobDto instantiates a new KubernetesJobDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesJobDto() *KubernetesJobDto {
	this := KubernetesJobDto{}
	return &this
}

// NewKubernetesJobDtoWithDefaults instantiates a new KubernetesJobDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesJobDtoWithDefaults() *KubernetesJobDto {
	this := KubernetesJobDto{}
	return &this
}

// GetMetadataName returns the MetadataName field value if set, zero value otherwise.
func (o *KubernetesJobDto) GetMetadataName() string {
	if o == nil || IsNil(o.MetadataName) {
		var ret string
		return ret
	}
	return *o.MetadataName
}

// GetMetadataNameOk returns a tuple with the MetadataName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesJobDto) GetMetadataNameOk() (*string, bool) {
	if o == nil || IsNil(o.MetadataName) {
		return nil, false
	}
	return o.MetadataName, true
}

// HasMetadataName returns a boolean if a field has been set.
func (o *KubernetesJobDto) HasMetadataName() bool {
	if o != nil && !IsNil(o.MetadataName) {
		return true
	}

	return false
}

// SetMetadataName gets a reference to the given string and assigns it to the MetadataName field.
func (o *KubernetesJobDto) SetMetadataName(v string) {
	o.MetadataName = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *KubernetesJobDto) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesJobDto) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *KubernetesJobDto) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *KubernetesJobDto) SetNamespace(v string) {
	o.Namespace = &v
}

// GetAge returns the Age field value if set, zero value otherwise.
func (o *KubernetesJobDto) GetAge() string {
	if o == nil || IsNil(o.Age) {
		var ret string
		return ret
	}
	return *o.Age
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesJobDto) GetAgeOk() (*string, bool) {
	if o == nil || IsNil(o.Age) {
		return nil, false
	}
	return o.Age, true
}

// HasAge returns a boolean if a field has been set.
func (o *KubernetesJobDto) HasAge() bool {
	if o != nil && !IsNil(o.Age) {
		return true
	}

	return false
}

// SetAge gets a reference to the given string and assigns it to the Age field.
func (o *KubernetesJobDto) SetAge(v string) {
	o.Age = &v
}

// GetCompletions returns the Completions field value if set, zero value otherwise.
func (o *KubernetesJobDto) GetCompletions() int32 {
	if o == nil || IsNil(o.Completions) {
		var ret int32
		return ret
	}
	return *o.Completions
}

// GetCompletionsOk returns a tuple with the Completions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesJobDto) GetCompletionsOk() (*int32, bool) {
	if o == nil || IsNil(o.Completions) {
		return nil, false
	}
	return o.Completions, true
}

// HasCompletions returns a boolean if a field has been set.
func (o *KubernetesJobDto) HasCompletions() bool {
	if o != nil && !IsNil(o.Completions) {
		return true
	}

	return false
}

// SetCompletions gets a reference to the given int32 and assigns it to the Completions field.
func (o *KubernetesJobDto) SetCompletions(v int32) {
	o.Completions = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *KubernetesJobDto) GetConditions() string {
	if o == nil || IsNil(o.Conditions) {
		var ret string
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesJobDto) GetConditionsOk() (*string, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *KubernetesJobDto) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given string and assigns it to the Conditions field.
func (o *KubernetesJobDto) SetConditions(v string) {
	o.Conditions = &v
}

func (o KubernetesJobDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesJobDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MetadataName) {
		toSerialize["metadataName"] = o.MetadataName
	}
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !IsNil(o.Age) {
		toSerialize["age"] = o.Age
	}
	if !IsNil(o.Completions) {
		toSerialize["completions"] = o.Completions
	}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	return toSerialize, nil
}

type NullableKubernetesJobDto struct {
	value *KubernetesJobDto
	isSet bool
}

func (v NullableKubernetesJobDto) Get() *KubernetesJobDto {
	return v.value
}

func (v *NullableKubernetesJobDto) Set(val *KubernetesJobDto) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesJobDto) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesJobDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesJobDto(val *KubernetesJobDto) *NullableKubernetesJobDto {
	return &NullableKubernetesJobDto{value: val, isSet: true}
}

func (v NullableKubernetesJobDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesJobDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


