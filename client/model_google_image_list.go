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

// checks if the GoogleImageList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GoogleImageList{}

// GoogleImageList struct for GoogleImageList
type GoogleImageList struct {
	Data []CommonStringBasedDropdownDto `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewGoogleImageList instantiates a new GoogleImageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleImageList() *GoogleImageList {
	this := GoogleImageList{}
	return &this
}

// NewGoogleImageListWithDefaults instantiates a new GoogleImageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleImageListWithDefaults() *GoogleImageList {
	this := GoogleImageList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GoogleImageList) GetData() []CommonStringBasedDropdownDto {
	if o == nil || IsNil(o.Data) {
		var ret []CommonStringBasedDropdownDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleImageList) GetDataOk() ([]CommonStringBasedDropdownDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GoogleImageList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CommonStringBasedDropdownDto and assigns it to the Data field.
func (o *GoogleImageList) SetData(v []CommonStringBasedDropdownDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *GoogleImageList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleImageList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *GoogleImageList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *GoogleImageList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o GoogleImageList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GoogleImageList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableGoogleImageList struct {
	value *GoogleImageList
	isSet bool
}

func (v NullableGoogleImageList) Get() *GoogleImageList {
	return v.value
}

func (v *NullableGoogleImageList) Set(val *GoogleImageList) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleImageList) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleImageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleImageList(val *GoogleImageList) *NullableGoogleImageList {
	return &NullableGoogleImageList{value: val, isSet: true}
}

func (v NullableGoogleImageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleImageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


