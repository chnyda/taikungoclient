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

// checks if the BackupCredentialsSearchList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupCredentialsSearchList{}

// BackupCredentialsSearchList struct for BackupCredentialsSearchList
type BackupCredentialsSearchList struct {
	Data []CommonSearchResponseData `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewBackupCredentialsSearchList instantiates a new BackupCredentialsSearchList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupCredentialsSearchList() *BackupCredentialsSearchList {
	this := BackupCredentialsSearchList{}
	return &this
}

// NewBackupCredentialsSearchListWithDefaults instantiates a new BackupCredentialsSearchList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupCredentialsSearchListWithDefaults() *BackupCredentialsSearchList {
	this := BackupCredentialsSearchList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BackupCredentialsSearchList) GetData() []CommonSearchResponseData {
	if o == nil || IsNil(o.Data) {
		var ret []CommonSearchResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupCredentialsSearchList) GetDataOk() ([]CommonSearchResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BackupCredentialsSearchList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CommonSearchResponseData and assigns it to the Data field.
func (o *BackupCredentialsSearchList) SetData(v []CommonSearchResponseData) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *BackupCredentialsSearchList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupCredentialsSearchList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *BackupCredentialsSearchList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *BackupCredentialsSearchList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o BackupCredentialsSearchList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackupCredentialsSearchList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableBackupCredentialsSearchList struct {
	value *BackupCredentialsSearchList
	isSet bool
}

func (v NullableBackupCredentialsSearchList) Get() *BackupCredentialsSearchList {
	return v.value
}

func (v *NullableBackupCredentialsSearchList) Set(val *BackupCredentialsSearchList) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupCredentialsSearchList) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupCredentialsSearchList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupCredentialsSearchList(val *BackupCredentialsSearchList) *NullableBackupCredentialsSearchList {
	return &NullableBackupCredentialsSearchList{value: val, isSet: true}
}

func (v NullableBackupCredentialsSearchList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupCredentialsSearchList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


