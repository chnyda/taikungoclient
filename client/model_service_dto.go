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

// checks if the ServiceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceDto{}

// ServiceDto struct for ServiceDto
type ServiceDto struct {
	MetadataName *string `json:"metadataName,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	Age *string `json:"age,omitempty"`
	Type *string `json:"type,omitempty"`
	Ip map[string]interface{} `json:"ip,omitempty"`
}

// NewServiceDto instantiates a new ServiceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceDto() *ServiceDto {
	this := ServiceDto{}
	return &this
}

// NewServiceDtoWithDefaults instantiates a new ServiceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceDtoWithDefaults() *ServiceDto {
	this := ServiceDto{}
	return &this
}

// GetMetadataName returns the MetadataName field value if set, zero value otherwise.
func (o *ServiceDto) GetMetadataName() string {
	if o == nil || IsNil(o.MetadataName) {
		var ret string
		return ret
	}
	return *o.MetadataName
}

// GetMetadataNameOk returns a tuple with the MetadataName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDto) GetMetadataNameOk() (*string, bool) {
	if o == nil || IsNil(o.MetadataName) {
		return nil, false
	}
	return o.MetadataName, true
}

// HasMetadataName returns a boolean if a field has been set.
func (o *ServiceDto) HasMetadataName() bool {
	if o != nil && !IsNil(o.MetadataName) {
		return true
	}

	return false
}

// SetMetadataName gets a reference to the given string and assigns it to the MetadataName field.
func (o *ServiceDto) SetMetadataName(v string) {
	o.MetadataName = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *ServiceDto) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDto) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *ServiceDto) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *ServiceDto) SetNamespace(v string) {
	o.Namespace = &v
}

// GetAge returns the Age field value if set, zero value otherwise.
func (o *ServiceDto) GetAge() string {
	if o == nil || IsNil(o.Age) {
		var ret string
		return ret
	}
	return *o.Age
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDto) GetAgeOk() (*string, bool) {
	if o == nil || IsNil(o.Age) {
		return nil, false
	}
	return o.Age, true
}

// HasAge returns a boolean if a field has been set.
func (o *ServiceDto) HasAge() bool {
	if o != nil && !IsNil(o.Age) {
		return true
	}

	return false
}

// SetAge gets a reference to the given string and assigns it to the Age field.
func (o *ServiceDto) SetAge(v string) {
	o.Age = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ServiceDto) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDto) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ServiceDto) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ServiceDto) SetType(v string) {
	o.Type = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *ServiceDto) GetIp() map[string]interface{} {
	if o == nil || IsNil(o.Ip) {
		var ret map[string]interface{}
		return ret
	}
	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDto) GetIpOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Ip) {
		return map[string]interface{}{}, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *ServiceDto) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given map[string]interface{} and assigns it to the Ip field.
func (o *ServiceDto) SetIp(v map[string]interface{}) {
	o.Ip = v
}

func (o ServiceDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceDto) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	return toSerialize, nil
}

type NullableServiceDto struct {
	value *ServiceDto
	isSet bool
}

func (v NullableServiceDto) Get() *ServiceDto {
	return v.value
}

func (v *NullableServiceDto) Set(val *ServiceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceDto(val *ServiceDto) *NullableServiceDto {
	return &NullableServiceDto{value: val, isSet: true}
}

func (v NullableServiceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


