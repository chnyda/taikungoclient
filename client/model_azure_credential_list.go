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

// checks if the AzureCredentialList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureCredentialList{}

// AzureCredentialList struct for AzureCredentialList
type AzureCredentialList struct {
	Data []AzureCredentialsListDto `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewAzureCredentialList instantiates a new AzureCredentialList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureCredentialList() *AzureCredentialList {
	this := AzureCredentialList{}
	return &this
}

// NewAzureCredentialListWithDefaults instantiates a new AzureCredentialList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureCredentialListWithDefaults() *AzureCredentialList {
	this := AzureCredentialList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AzureCredentialList) GetData() []AzureCredentialsListDto {
	if o == nil || IsNil(o.Data) {
		var ret []AzureCredentialsListDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCredentialList) GetDataOk() ([]AzureCredentialsListDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AzureCredentialList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AzureCredentialsListDto and assigns it to the Data field.
func (o *AzureCredentialList) SetData(v []AzureCredentialsListDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *AzureCredentialList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCredentialList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *AzureCredentialList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *AzureCredentialList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o AzureCredentialList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureCredentialList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableAzureCredentialList struct {
	value *AzureCredentialList
	isSet bool
}

func (v NullableAzureCredentialList) Get() *AzureCredentialList {
	return v.value
}

func (v *NullableAzureCredentialList) Set(val *AzureCredentialList) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureCredentialList) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureCredentialList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureCredentialList(val *AzureCredentialList) *NullableAzureCredentialList {
	return &NullableAzureCredentialList{value: val, isSet: true}
}

func (v NullableAzureCredentialList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureCredentialList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


