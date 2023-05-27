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

// checks if the GoogleFlavorList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GoogleFlavorList{}

// GoogleFlavorList struct for GoogleFlavorList
type GoogleFlavorList struct {
	Data []GoogleFlavorDto `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewGoogleFlavorList instantiates a new GoogleFlavorList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleFlavorList() *GoogleFlavorList {
	this := GoogleFlavorList{}
	return &this
}

// NewGoogleFlavorListWithDefaults instantiates a new GoogleFlavorList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleFlavorListWithDefaults() *GoogleFlavorList {
	this := GoogleFlavorList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoogleFlavorList) GetData() []GoogleFlavorDto {
	if o == nil {
		var ret []GoogleFlavorDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoogleFlavorList) GetDataOk() ([]GoogleFlavorDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GoogleFlavorList) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GoogleFlavorDto and assigns it to the Data field.
func (o *GoogleFlavorList) SetData(v []GoogleFlavorDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *GoogleFlavorList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleFlavorList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *GoogleFlavorList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *GoogleFlavorList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o GoogleFlavorList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GoogleFlavorList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableGoogleFlavorList struct {
	value *GoogleFlavorList
	isSet bool
}

func (v NullableGoogleFlavorList) Get() *GoogleFlavorList {
	return v.value
}

func (v *NullableGoogleFlavorList) Set(val *GoogleFlavorList) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleFlavorList) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleFlavorList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleFlavorList(val *GoogleFlavorList) *NullableGoogleFlavorList {
	return &NullableGoogleFlavorList{value: val, isSet: true}
}

func (v NullableGoogleFlavorList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleFlavorList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


