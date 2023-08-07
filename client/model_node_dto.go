/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikuncore

import (
	"encoding/json"
)

// checks if the NodeDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeDto{}

// NodeDto struct for NodeDto
type NodeDto struct {
	MetadataName interface{} `json:"metadataName,omitempty"`
	KubeletReady interface{} `json:"kubeletReady,omitempty"`
	KubeletSufficient interface{} `json:"kubeletSufficient,omitempty"`
	KubeletDiskPressure interface{} `json:"kubeletDiskPressure,omitempty"`
	KubeletMemory interface{} `json:"kubeletMemory,omitempty"`
}

// NewNodeDto instantiates a new NodeDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeDto() *NodeDto {
	this := NodeDto{}
	return &this
}

// NewNodeDtoWithDefaults instantiates a new NodeDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeDtoWithDefaults() *NodeDto {
	this := NodeDto{}
	return &this
}

// GetMetadataName returns the MetadataName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NodeDto) GetMetadataName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MetadataName
}

// GetMetadataNameOk returns a tuple with the MetadataName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodeDto) GetMetadataNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MetadataName) {
		return nil, false
	}
	return &o.MetadataName, true
}

// HasMetadataName returns a boolean if a field has been set.
func (o *NodeDto) HasMetadataName() bool {
	if o != nil && IsNil(o.MetadataName) {
		return true
	}

	return false
}

// SetMetadataName gets a reference to the given interface{} and assigns it to the MetadataName field.
func (o *NodeDto) SetMetadataName(v interface{}) {
	o.MetadataName = v
}

// GetKubeletReady returns the KubeletReady field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NodeDto) GetKubeletReady() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.KubeletReady
}

// GetKubeletReadyOk returns a tuple with the KubeletReady field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodeDto) GetKubeletReadyOk() (*interface{}, bool) {
	if o == nil || IsNil(o.KubeletReady) {
		return nil, false
	}
	return &o.KubeletReady, true
}

// HasKubeletReady returns a boolean if a field has been set.
func (o *NodeDto) HasKubeletReady() bool {
	if o != nil && IsNil(o.KubeletReady) {
		return true
	}

	return false
}

// SetKubeletReady gets a reference to the given interface{} and assigns it to the KubeletReady field.
func (o *NodeDto) SetKubeletReady(v interface{}) {
	o.KubeletReady = v
}

// GetKubeletSufficient returns the KubeletSufficient field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NodeDto) GetKubeletSufficient() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.KubeletSufficient
}

// GetKubeletSufficientOk returns a tuple with the KubeletSufficient field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodeDto) GetKubeletSufficientOk() (*interface{}, bool) {
	if o == nil || IsNil(o.KubeletSufficient) {
		return nil, false
	}
	return &o.KubeletSufficient, true
}

// HasKubeletSufficient returns a boolean if a field has been set.
func (o *NodeDto) HasKubeletSufficient() bool {
	if o != nil && IsNil(o.KubeletSufficient) {
		return true
	}

	return false
}

// SetKubeletSufficient gets a reference to the given interface{} and assigns it to the KubeletSufficient field.
func (o *NodeDto) SetKubeletSufficient(v interface{}) {
	o.KubeletSufficient = v
}

// GetKubeletDiskPressure returns the KubeletDiskPressure field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NodeDto) GetKubeletDiskPressure() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.KubeletDiskPressure
}

// GetKubeletDiskPressureOk returns a tuple with the KubeletDiskPressure field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodeDto) GetKubeletDiskPressureOk() (*interface{}, bool) {
	if o == nil || IsNil(o.KubeletDiskPressure) {
		return nil, false
	}
	return &o.KubeletDiskPressure, true
}

// HasKubeletDiskPressure returns a boolean if a field has been set.
func (o *NodeDto) HasKubeletDiskPressure() bool {
	if o != nil && IsNil(o.KubeletDiskPressure) {
		return true
	}

	return false
}

// SetKubeletDiskPressure gets a reference to the given interface{} and assigns it to the KubeletDiskPressure field.
func (o *NodeDto) SetKubeletDiskPressure(v interface{}) {
	o.KubeletDiskPressure = v
}

// GetKubeletMemory returns the KubeletMemory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NodeDto) GetKubeletMemory() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.KubeletMemory
}

// GetKubeletMemoryOk returns a tuple with the KubeletMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodeDto) GetKubeletMemoryOk() (*interface{}, bool) {
	if o == nil || IsNil(o.KubeletMemory) {
		return nil, false
	}
	return &o.KubeletMemory, true
}

// HasKubeletMemory returns a boolean if a field has been set.
func (o *NodeDto) HasKubeletMemory() bool {
	if o != nil && IsNil(o.KubeletMemory) {
		return true
	}

	return false
}

// SetKubeletMemory gets a reference to the given interface{} and assigns it to the KubeletMemory field.
func (o *NodeDto) SetKubeletMemory(v interface{}) {
	o.KubeletMemory = v
}

func (o NodeDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MetadataName != nil {
		toSerialize["metadataName"] = o.MetadataName
	}
	if o.KubeletReady != nil {
		toSerialize["kubeletReady"] = o.KubeletReady
	}
	if o.KubeletSufficient != nil {
		toSerialize["kubeletSufficient"] = o.KubeletSufficient
	}
	if o.KubeletDiskPressure != nil {
		toSerialize["kubeletDiskPressure"] = o.KubeletDiskPressure
	}
	if o.KubeletMemory != nil {
		toSerialize["kubeletMemory"] = o.KubeletMemory
	}
	return toSerialize, nil
}

type NullableNodeDto struct {
	value *NodeDto
	isSet bool
}

func (v NullableNodeDto) Get() *NodeDto {
	return v.value
}

func (v *NullableNodeDto) Set(val *NodeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeDto(val *NodeDto) *NullableNodeDto {
	return &NullableNodeDto{value: val, isSet: true}
}

func (v NullableNodeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

