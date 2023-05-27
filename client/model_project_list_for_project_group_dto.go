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

// checks if the ProjectListForProjectGroupDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectListForProjectGroupDto{}

// ProjectListForProjectGroupDto struct for ProjectListForProjectGroupDto
type ProjectListForProjectGroupDto struct {
	ProjectId *int32 `json:"projectId,omitempty"`
	ProjectName NullableString `json:"projectName,omitempty"`
	IsSelected *bool `json:"isSelected,omitempty"`
}

// NewProjectListForProjectGroupDto instantiates a new ProjectListForProjectGroupDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectListForProjectGroupDto() *ProjectListForProjectGroupDto {
	this := ProjectListForProjectGroupDto{}
	return &this
}

// NewProjectListForProjectGroupDtoWithDefaults instantiates a new ProjectListForProjectGroupDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectListForProjectGroupDtoWithDefaults() *ProjectListForProjectGroupDto {
	this := ProjectListForProjectGroupDto{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *ProjectListForProjectGroupDto) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectListForProjectGroupDto) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ProjectListForProjectGroupDto) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *ProjectListForProjectGroupDto) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectListForProjectGroupDto) GetProjectName() string {
	if o == nil || IsNil(o.ProjectName.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectName.Get()
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectListForProjectGroupDto) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectName.Get(), o.ProjectName.IsSet()
}

// HasProjectName returns a boolean if a field has been set.
func (o *ProjectListForProjectGroupDto) HasProjectName() bool {
	if o != nil && o.ProjectName.IsSet() {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given NullableString and assigns it to the ProjectName field.
func (o *ProjectListForProjectGroupDto) SetProjectName(v string) {
	o.ProjectName.Set(&v)
}
// SetProjectNameNil sets the value for ProjectName to be an explicit nil
func (o *ProjectListForProjectGroupDto) SetProjectNameNil() {
	o.ProjectName.Set(nil)
}

// UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
func (o *ProjectListForProjectGroupDto) UnsetProjectName() {
	o.ProjectName.Unset()
}

// GetIsSelected returns the IsSelected field value if set, zero value otherwise.
func (o *ProjectListForProjectGroupDto) GetIsSelected() bool {
	if o == nil || IsNil(o.IsSelected) {
		var ret bool
		return ret
	}
	return *o.IsSelected
}

// GetIsSelectedOk returns a tuple with the IsSelected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectListForProjectGroupDto) GetIsSelectedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSelected) {
		return nil, false
	}
	return o.IsSelected, true
}

// HasIsSelected returns a boolean if a field has been set.
func (o *ProjectListForProjectGroupDto) HasIsSelected() bool {
	if o != nil && !IsNil(o.IsSelected) {
		return true
	}

	return false
}

// SetIsSelected gets a reference to the given bool and assigns it to the IsSelected field.
func (o *ProjectListForProjectGroupDto) SetIsSelected(v bool) {
	o.IsSelected = &v
}

func (o ProjectListForProjectGroupDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectListForProjectGroupDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.ProjectName.IsSet() {
		toSerialize["projectName"] = o.ProjectName.Get()
	}
	if !IsNil(o.IsSelected) {
		toSerialize["isSelected"] = o.IsSelected
	}
	return toSerialize, nil
}

type NullableProjectListForProjectGroupDto struct {
	value *ProjectListForProjectGroupDto
	isSet bool
}

func (v NullableProjectListForProjectGroupDto) Get() *ProjectListForProjectGroupDto {
	return v.value
}

func (v *NullableProjectListForProjectGroupDto) Set(val *ProjectListForProjectGroupDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectListForProjectGroupDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectListForProjectGroupDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectListForProjectGroupDto(val *ProjectListForProjectGroupDto) *NullableProjectListForProjectGroupDto {
	return &NullableProjectListForProjectGroupDto{value: val, isSet: true}
}

func (v NullableProjectListForProjectGroupDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectListForProjectGroupDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


