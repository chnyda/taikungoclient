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

// checks if the TanzuFlavorsListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TanzuFlavorsListDto{}

// TanzuFlavorsListDto struct for TanzuFlavorsListDto
type TanzuFlavorsListDto struct {
	Name *string `json:"name,omitempty"`
	Cpu *int32 `json:"cpu,omitempty"`
	Ram *int64 `json:"ram,omitempty"`
}

// NewTanzuFlavorsListDto instantiates a new TanzuFlavorsListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTanzuFlavorsListDto() *TanzuFlavorsListDto {
	this := TanzuFlavorsListDto{}
	return &this
}

// NewTanzuFlavorsListDtoWithDefaults instantiates a new TanzuFlavorsListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTanzuFlavorsListDtoWithDefaults() *TanzuFlavorsListDto {
	this := TanzuFlavorsListDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TanzuFlavorsListDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TanzuFlavorsListDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TanzuFlavorsListDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TanzuFlavorsListDto) SetName(v string) {
	o.Name = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *TanzuFlavorsListDto) GetCpu() int32 {
	if o == nil || IsNil(o.Cpu) {
		var ret int32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TanzuFlavorsListDto) GetCpuOk() (*int32, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *TanzuFlavorsListDto) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int32 and assigns it to the Cpu field.
func (o *TanzuFlavorsListDto) SetCpu(v int32) {
	o.Cpu = &v
}

// GetRam returns the Ram field value if set, zero value otherwise.
func (o *TanzuFlavorsListDto) GetRam() int64 {
	if o == nil || IsNil(o.Ram) {
		var ret int64
		return ret
	}
	return *o.Ram
}

// GetRamOk returns a tuple with the Ram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TanzuFlavorsListDto) GetRamOk() (*int64, bool) {
	if o == nil || IsNil(o.Ram) {
		return nil, false
	}
	return o.Ram, true
}

// HasRam returns a boolean if a field has been set.
func (o *TanzuFlavorsListDto) HasRam() bool {
	if o != nil && !IsNil(o.Ram) {
		return true
	}

	return false
}

// SetRam gets a reference to the given int64 and assigns it to the Ram field.
func (o *TanzuFlavorsListDto) SetRam(v int64) {
	o.Ram = &v
}

func (o TanzuFlavorsListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TanzuFlavorsListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.Ram) {
		toSerialize["ram"] = o.Ram
	}
	return toSerialize, nil
}

type NullableTanzuFlavorsListDto struct {
	value *TanzuFlavorsListDto
	isSet bool
}

func (v NullableTanzuFlavorsListDto) Get() *TanzuFlavorsListDto {
	return v.value
}

func (v *NullableTanzuFlavorsListDto) Set(val *TanzuFlavorsListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTanzuFlavorsListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTanzuFlavorsListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTanzuFlavorsListDto(val *TanzuFlavorsListDto) *NullableTanzuFlavorsListDto {
	return &NullableTanzuFlavorsListDto{value: val, isSet: true}
}

func (v NullableTanzuFlavorsListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTanzuFlavorsListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


