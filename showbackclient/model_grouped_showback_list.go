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

// checks if the GroupedShowbackList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupedShowbackList{}

// GroupedShowbackList struct for GroupedShowbackList
type GroupedShowbackList struct {
	Projects []GroupedShowbackSummaryInfos `json:"projects,omitempty"`
	Credentials []GroupedShowbackSummaryInfos `json:"credentials,omitempty"`
	ByLabelValues []GroupedShowbackSummaryInfos `json:"byLabelValues,omitempty"`
}

// NewGroupedShowbackList instantiates a new GroupedShowbackList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupedShowbackList() *GroupedShowbackList {
	this := GroupedShowbackList{}
	return &this
}

// NewGroupedShowbackListWithDefaults instantiates a new GroupedShowbackList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupedShowbackListWithDefaults() *GroupedShowbackList {
	this := GroupedShowbackList{}
	return &this
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *GroupedShowbackList) GetProjects() []GroupedShowbackSummaryInfos {
	if o == nil || IsNil(o.Projects) {
		var ret []GroupedShowbackSummaryInfos
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupedShowbackList) GetProjectsOk() ([]GroupedShowbackSummaryInfos, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *GroupedShowbackList) HasProjects() bool {
	if o != nil && !IsNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []GroupedShowbackSummaryInfos and assigns it to the Projects field.
func (o *GroupedShowbackList) SetProjects(v []GroupedShowbackSummaryInfos) {
	o.Projects = v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *GroupedShowbackList) GetCredentials() []GroupedShowbackSummaryInfos {
	if o == nil || IsNil(o.Credentials) {
		var ret []GroupedShowbackSummaryInfos
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupedShowbackList) GetCredentialsOk() ([]GroupedShowbackSummaryInfos, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *GroupedShowbackList) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given []GroupedShowbackSummaryInfos and assigns it to the Credentials field.
func (o *GroupedShowbackList) SetCredentials(v []GroupedShowbackSummaryInfos) {
	o.Credentials = v
}

// GetByLabelValues returns the ByLabelValues field value if set, zero value otherwise.
func (o *GroupedShowbackList) GetByLabelValues() []GroupedShowbackSummaryInfos {
	if o == nil || IsNil(o.ByLabelValues) {
		var ret []GroupedShowbackSummaryInfos
		return ret
	}
	return o.ByLabelValues
}

// GetByLabelValuesOk returns a tuple with the ByLabelValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupedShowbackList) GetByLabelValuesOk() ([]GroupedShowbackSummaryInfos, bool) {
	if o == nil || IsNil(o.ByLabelValues) {
		return nil, false
	}
	return o.ByLabelValues, true
}

// HasByLabelValues returns a boolean if a field has been set.
func (o *GroupedShowbackList) HasByLabelValues() bool {
	if o != nil && !IsNil(o.ByLabelValues) {
		return true
	}

	return false
}

// SetByLabelValues gets a reference to the given []GroupedShowbackSummaryInfos and assigns it to the ByLabelValues field.
func (o *GroupedShowbackList) SetByLabelValues(v []GroupedShowbackSummaryInfos) {
	o.ByLabelValues = v
}

func (o GroupedShowbackList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupedShowbackList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Projects) {
		toSerialize["projects"] = o.Projects
	}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.ByLabelValues) {
		toSerialize["byLabelValues"] = o.ByLabelValues
	}
	return toSerialize, nil
}

type NullableGroupedShowbackList struct {
	value *GroupedShowbackList
	isSet bool
}

func (v NullableGroupedShowbackList) Get() *GroupedShowbackList {
	return v.value
}

func (v *NullableGroupedShowbackList) Set(val *GroupedShowbackList) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupedShowbackList) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupedShowbackList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupedShowbackList(val *GroupedShowbackList) *NullableGroupedShowbackList {
	return &NullableGroupedShowbackList{value: val, isSet: true}
}

func (v NullableGroupedShowbackList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupedShowbackList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


