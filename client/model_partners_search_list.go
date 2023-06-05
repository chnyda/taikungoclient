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

// checks if the PartnersSearchList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartnersSearchList{}

// PartnersSearchList struct for PartnersSearchList
type PartnersSearchList struct {
	Data []PartnersSearchResponseData `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewPartnersSearchList instantiates a new PartnersSearchList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnersSearchList() *PartnersSearchList {
	this := PartnersSearchList{}
	return &this
}

// NewPartnersSearchListWithDefaults instantiates a new PartnersSearchList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnersSearchListWithDefaults() *PartnersSearchList {
	this := PartnersSearchList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PartnersSearchList) GetData() []PartnersSearchResponseData {
	if o == nil || IsNil(o.Data) {
		var ret []PartnersSearchResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnersSearchList) GetDataOk() ([]PartnersSearchResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PartnersSearchList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []PartnersSearchResponseData and assigns it to the Data field.
func (o *PartnersSearchList) SetData(v []PartnersSearchResponseData) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *PartnersSearchList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnersSearchList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PartnersSearchList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *PartnersSearchList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o PartnersSearchList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PartnersSearchList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullablePartnersSearchList struct {
	value *PartnersSearchList
	isSet bool
}

func (v NullablePartnersSearchList) Get() *PartnersSearchList {
	return v.value
}

func (v *NullablePartnersSearchList) Set(val *PartnersSearchList) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnersSearchList) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnersSearchList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnersSearchList(val *PartnersSearchList) *NullablePartnersSearchList {
	return &NullablePartnersSearchList{value: val, isSet: true}
}

func (v NullablePartnersSearchList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnersSearchList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


