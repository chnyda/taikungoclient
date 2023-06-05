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

// checks if the BillingCredentialsSearchList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingCredentialsSearchList{}

// BillingCredentialsSearchList struct for BillingCredentialsSearchList
type BillingCredentialsSearchList struct {
	Data []CommonSearchResponseData `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewBillingCredentialsSearchList instantiates a new BillingCredentialsSearchList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingCredentialsSearchList() *BillingCredentialsSearchList {
	this := BillingCredentialsSearchList{}
	return &this
}

// NewBillingCredentialsSearchListWithDefaults instantiates a new BillingCredentialsSearchList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingCredentialsSearchListWithDefaults() *BillingCredentialsSearchList {
	this := BillingCredentialsSearchList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BillingCredentialsSearchList) GetData() []CommonSearchResponseData {
	if o == nil || IsNil(o.Data) {
		var ret []CommonSearchResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingCredentialsSearchList) GetDataOk() ([]CommonSearchResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BillingCredentialsSearchList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CommonSearchResponseData and assigns it to the Data field.
func (o *BillingCredentialsSearchList) SetData(v []CommonSearchResponseData) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *BillingCredentialsSearchList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingCredentialsSearchList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *BillingCredentialsSearchList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *BillingCredentialsSearchList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o BillingCredentialsSearchList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingCredentialsSearchList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableBillingCredentialsSearchList struct {
	value *BillingCredentialsSearchList
	isSet bool
}

func (v NullableBillingCredentialsSearchList) Get() *BillingCredentialsSearchList {
	return v.value
}

func (v *NullableBillingCredentialsSearchList) Set(val *BillingCredentialsSearchList) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingCredentialsSearchList) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingCredentialsSearchList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingCredentialsSearchList(val *BillingCredentialsSearchList) *NullableBillingCredentialsSearchList {
	return &NullableBillingCredentialsSearchList{value: val, isSet: true}
}

func (v NullableBillingCredentialsSearchList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingCredentialsSearchList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


