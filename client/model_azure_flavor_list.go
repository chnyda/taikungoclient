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

// checks if the AzureFlavorList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureFlavorList{}

// AzureFlavorList struct for AzureFlavorList
type AzureFlavorList struct {
	Data []AzureFlavorsWithPriceDto `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewAzureFlavorList instantiates a new AzureFlavorList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureFlavorList() *AzureFlavorList {
	this := AzureFlavorList{}
	return &this
}

// NewAzureFlavorListWithDefaults instantiates a new AzureFlavorList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureFlavorListWithDefaults() *AzureFlavorList {
	this := AzureFlavorList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureFlavorList) GetData() []AzureFlavorsWithPriceDto {
	if o == nil {
		var ret []AzureFlavorsWithPriceDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureFlavorList) GetDataOk() ([]AzureFlavorsWithPriceDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AzureFlavorList) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AzureFlavorsWithPriceDto and assigns it to the Data field.
func (o *AzureFlavorList) SetData(v []AzureFlavorsWithPriceDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *AzureFlavorList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureFlavorList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *AzureFlavorList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *AzureFlavorList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o AzureFlavorList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureFlavorList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableAzureFlavorList struct {
	value *AzureFlavorList
	isSet bool
}

func (v NullableAzureFlavorList) Get() *AzureFlavorList {
	return v.value
}

func (v *NullableAzureFlavorList) Set(val *AzureFlavorList) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureFlavorList) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureFlavorList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureFlavorList(val *AzureFlavorList) *NullableAzureFlavorList {
	return &NullableAzureFlavorList{value: val, isSet: true}
}

func (v NullableAzureFlavorList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureFlavorList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


