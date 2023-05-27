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

// checks if the OpenstackVolumeQuotaDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenstackVolumeQuotaDto{}

// OpenstackVolumeQuotaDto struct for OpenstackVolumeQuotaDto
type OpenstackVolumeQuotaDto struct {
	MaxTotalVolumeSize *int64 `json:"maxTotalVolumeSize,omitempty"`
	UsedVolumeSize *int64 `json:"usedVolumeSize,omitempty"`
	MaxCountVolumeSize *int64 `json:"maxCountVolumeSize,omitempty"`
	CountVolumeSize *int64 `json:"countVolumeSize,omitempty"`
}

// NewOpenstackVolumeQuotaDto instantiates a new OpenstackVolumeQuotaDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenstackVolumeQuotaDto() *OpenstackVolumeQuotaDto {
	this := OpenstackVolumeQuotaDto{}
	return &this
}

// NewOpenstackVolumeQuotaDtoWithDefaults instantiates a new OpenstackVolumeQuotaDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenstackVolumeQuotaDtoWithDefaults() *OpenstackVolumeQuotaDto {
	this := OpenstackVolumeQuotaDto{}
	return &this
}

// GetMaxTotalVolumeSize returns the MaxTotalVolumeSize field value if set, zero value otherwise.
func (o *OpenstackVolumeQuotaDto) GetMaxTotalVolumeSize() int64 {
	if o == nil || IsNil(o.MaxTotalVolumeSize) {
		var ret int64
		return ret
	}
	return *o.MaxTotalVolumeSize
}

// GetMaxTotalVolumeSizeOk returns a tuple with the MaxTotalVolumeSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackVolumeQuotaDto) GetMaxTotalVolumeSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxTotalVolumeSize) {
		return nil, false
	}
	return o.MaxTotalVolumeSize, true
}

// HasMaxTotalVolumeSize returns a boolean if a field has been set.
func (o *OpenstackVolumeQuotaDto) HasMaxTotalVolumeSize() bool {
	if o != nil && !IsNil(o.MaxTotalVolumeSize) {
		return true
	}

	return false
}

// SetMaxTotalVolumeSize gets a reference to the given int64 and assigns it to the MaxTotalVolumeSize field.
func (o *OpenstackVolumeQuotaDto) SetMaxTotalVolumeSize(v int64) {
	o.MaxTotalVolumeSize = &v
}

// GetUsedVolumeSize returns the UsedVolumeSize field value if set, zero value otherwise.
func (o *OpenstackVolumeQuotaDto) GetUsedVolumeSize() int64 {
	if o == nil || IsNil(o.UsedVolumeSize) {
		var ret int64
		return ret
	}
	return *o.UsedVolumeSize
}

// GetUsedVolumeSizeOk returns a tuple with the UsedVolumeSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackVolumeQuotaDto) GetUsedVolumeSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.UsedVolumeSize) {
		return nil, false
	}
	return o.UsedVolumeSize, true
}

// HasUsedVolumeSize returns a boolean if a field has been set.
func (o *OpenstackVolumeQuotaDto) HasUsedVolumeSize() bool {
	if o != nil && !IsNil(o.UsedVolumeSize) {
		return true
	}

	return false
}

// SetUsedVolumeSize gets a reference to the given int64 and assigns it to the UsedVolumeSize field.
func (o *OpenstackVolumeQuotaDto) SetUsedVolumeSize(v int64) {
	o.UsedVolumeSize = &v
}

// GetMaxCountVolumeSize returns the MaxCountVolumeSize field value if set, zero value otherwise.
func (o *OpenstackVolumeQuotaDto) GetMaxCountVolumeSize() int64 {
	if o == nil || IsNil(o.MaxCountVolumeSize) {
		var ret int64
		return ret
	}
	return *o.MaxCountVolumeSize
}

// GetMaxCountVolumeSizeOk returns a tuple with the MaxCountVolumeSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackVolumeQuotaDto) GetMaxCountVolumeSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxCountVolumeSize) {
		return nil, false
	}
	return o.MaxCountVolumeSize, true
}

// HasMaxCountVolumeSize returns a boolean if a field has been set.
func (o *OpenstackVolumeQuotaDto) HasMaxCountVolumeSize() bool {
	if o != nil && !IsNil(o.MaxCountVolumeSize) {
		return true
	}

	return false
}

// SetMaxCountVolumeSize gets a reference to the given int64 and assigns it to the MaxCountVolumeSize field.
func (o *OpenstackVolumeQuotaDto) SetMaxCountVolumeSize(v int64) {
	o.MaxCountVolumeSize = &v
}

// GetCountVolumeSize returns the CountVolumeSize field value if set, zero value otherwise.
func (o *OpenstackVolumeQuotaDto) GetCountVolumeSize() int64 {
	if o == nil || IsNil(o.CountVolumeSize) {
		var ret int64
		return ret
	}
	return *o.CountVolumeSize
}

// GetCountVolumeSizeOk returns a tuple with the CountVolumeSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackVolumeQuotaDto) GetCountVolumeSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.CountVolumeSize) {
		return nil, false
	}
	return o.CountVolumeSize, true
}

// HasCountVolumeSize returns a boolean if a field has been set.
func (o *OpenstackVolumeQuotaDto) HasCountVolumeSize() bool {
	if o != nil && !IsNil(o.CountVolumeSize) {
		return true
	}

	return false
}

// SetCountVolumeSize gets a reference to the given int64 and assigns it to the CountVolumeSize field.
func (o *OpenstackVolumeQuotaDto) SetCountVolumeSize(v int64) {
	o.CountVolumeSize = &v
}

func (o OpenstackVolumeQuotaDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenstackVolumeQuotaDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxTotalVolumeSize) {
		toSerialize["maxTotalVolumeSize"] = o.MaxTotalVolumeSize
	}
	if !IsNil(o.UsedVolumeSize) {
		toSerialize["usedVolumeSize"] = o.UsedVolumeSize
	}
	if !IsNil(o.MaxCountVolumeSize) {
		toSerialize["maxCountVolumeSize"] = o.MaxCountVolumeSize
	}
	if !IsNil(o.CountVolumeSize) {
		toSerialize["countVolumeSize"] = o.CountVolumeSize
	}
	return toSerialize, nil
}

type NullableOpenstackVolumeQuotaDto struct {
	value *OpenstackVolumeQuotaDto
	isSet bool
}

func (v NullableOpenstackVolumeQuotaDto) Get() *OpenstackVolumeQuotaDto {
	return v.value
}

func (v *NullableOpenstackVolumeQuotaDto) Set(val *OpenstackVolumeQuotaDto) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenstackVolumeQuotaDto) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenstackVolumeQuotaDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenstackVolumeQuotaDto(val *OpenstackVolumeQuotaDto) *NullableOpenstackVolumeQuotaDto {
	return &NullableOpenstackVolumeQuotaDto{value: val, isSet: true}
}

func (v NullableOpenstackVolumeQuotaDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenstackVolumeQuotaDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


