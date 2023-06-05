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

// checks if the OpenstackImageList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenstackImageList{}

// OpenstackImageList struct for OpenstackImageList
type OpenstackImageList struct {
	Data []CommonStringBasedDropdownDto `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewOpenstackImageList instantiates a new OpenstackImageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenstackImageList() *OpenstackImageList {
	this := OpenstackImageList{}
	return &this
}

// NewOpenstackImageListWithDefaults instantiates a new OpenstackImageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenstackImageListWithDefaults() *OpenstackImageList {
	this := OpenstackImageList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OpenstackImageList) GetData() []CommonStringBasedDropdownDto {
	if o == nil || IsNil(o.Data) {
		var ret []CommonStringBasedDropdownDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackImageList) GetDataOk() ([]CommonStringBasedDropdownDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OpenstackImageList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CommonStringBasedDropdownDto and assigns it to the Data field.
func (o *OpenstackImageList) SetData(v []CommonStringBasedDropdownDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *OpenstackImageList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackImageList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *OpenstackImageList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *OpenstackImageList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o OpenstackImageList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenstackImageList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableOpenstackImageList struct {
	value *OpenstackImageList
	isSet bool
}

func (v NullableOpenstackImageList) Get() *OpenstackImageList {
	return v.value
}

func (v *NullableOpenstackImageList) Set(val *OpenstackImageList) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenstackImageList) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenstackImageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenstackImageList(val *OpenstackImageList) *NullableOpenstackImageList {
	return &NullableOpenstackImageList{value: val, isSet: true}
}

func (v NullableOpenstackImageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenstackImageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


