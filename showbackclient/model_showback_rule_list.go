/*
Taikun - Showback API

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@itera.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikunshowback

import (
	"encoding/json"
)

// checks if the ShowbackRuleList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShowbackRuleList{}

// ShowbackRuleList struct for ShowbackRuleList
type ShowbackRuleList struct {
	Data []ShowbackRulesListDto `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewShowbackRuleList instantiates a new ShowbackRuleList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShowbackRuleList() *ShowbackRuleList {
	this := ShowbackRuleList{}
	return &this
}

// NewShowbackRuleListWithDefaults instantiates a new ShowbackRuleList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShowbackRuleListWithDefaults() *ShowbackRuleList {
	this := ShowbackRuleList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ShowbackRuleList) GetData() []ShowbackRulesListDto {
	if o == nil || IsNil(o.Data) {
		var ret []ShowbackRulesListDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShowbackRuleList) GetDataOk() ([]ShowbackRulesListDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ShowbackRuleList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ShowbackRulesListDto and assigns it to the Data field.
func (o *ShowbackRuleList) SetData(v []ShowbackRulesListDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ShowbackRuleList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShowbackRuleList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ShowbackRuleList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *ShowbackRuleList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o ShowbackRuleList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShowbackRuleList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableShowbackRuleList struct {
	value *ShowbackRuleList
	isSet bool
}

func (v NullableShowbackRuleList) Get() *ShowbackRuleList {
	return v.value
}

func (v *NullableShowbackRuleList) Set(val *ShowbackRuleList) {
	v.value = val
	v.isSet = true
}

func (v NullableShowbackRuleList) IsSet() bool {
	return v.isSet
}

func (v *NullableShowbackRuleList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShowbackRuleList(val *ShowbackRuleList) *NullableShowbackRuleList {
	return &NullableShowbackRuleList{value: val, isSet: true}
}

func (v NullableShowbackRuleList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShowbackRuleList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


