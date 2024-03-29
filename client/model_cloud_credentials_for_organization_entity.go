/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikuncore

import (
	"encoding/json"
)

// checks if the CloudCredentialsForOrganizationEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudCredentialsForOrganizationEntity{}

// CloudCredentialsForOrganizationEntity struct for CloudCredentialsForOrganizationEntity
type CloudCredentialsForOrganizationEntity struct {
	Id *int32 `json:"id,omitempty"`
	Projects []CommonDropdownDto `json:"projects,omitempty"`
	FullName NullableString `json:"fullName,omitempty"`
	CloudType *CloudType `json:"cloudType,omitempty"`
	IsDefault *bool `json:"isDefault,omitempty"`
}

// NewCloudCredentialsForOrganizationEntity instantiates a new CloudCredentialsForOrganizationEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudCredentialsForOrganizationEntity() *CloudCredentialsForOrganizationEntity {
	this := CloudCredentialsForOrganizationEntity{}
	return &this
}

// NewCloudCredentialsForOrganizationEntityWithDefaults instantiates a new CloudCredentialsForOrganizationEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudCredentialsForOrganizationEntityWithDefaults() *CloudCredentialsForOrganizationEntity {
	this := CloudCredentialsForOrganizationEntity{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CloudCredentialsForOrganizationEntity) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudCredentialsForOrganizationEntity) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CloudCredentialsForOrganizationEntity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CloudCredentialsForOrganizationEntity) SetId(v int32) {
	o.Id = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudCredentialsForOrganizationEntity) GetProjects() []CommonDropdownDto {
	if o == nil {
		var ret []CommonDropdownDto
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudCredentialsForOrganizationEntity) GetProjectsOk() ([]CommonDropdownDto, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *CloudCredentialsForOrganizationEntity) HasProjects() bool {
	if o != nil && IsNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []CommonDropdownDto and assigns it to the Projects field.
func (o *CloudCredentialsForOrganizationEntity) SetProjects(v []CommonDropdownDto) {
	o.Projects = v
}

// GetFullName returns the FullName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudCredentialsForOrganizationEntity) GetFullName() string {
	if o == nil || IsNil(o.FullName.Get()) {
		var ret string
		return ret
	}
	return *o.FullName.Get()
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudCredentialsForOrganizationEntity) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FullName.Get(), o.FullName.IsSet()
}

// HasFullName returns a boolean if a field has been set.
func (o *CloudCredentialsForOrganizationEntity) HasFullName() bool {
	if o != nil && o.FullName.IsSet() {
		return true
	}

	return false
}

// SetFullName gets a reference to the given NullableString and assigns it to the FullName field.
func (o *CloudCredentialsForOrganizationEntity) SetFullName(v string) {
	o.FullName.Set(&v)
}
// SetFullNameNil sets the value for FullName to be an explicit nil
func (o *CloudCredentialsForOrganizationEntity) SetFullNameNil() {
	o.FullName.Set(nil)
}

// UnsetFullName ensures that no value is present for FullName, not even an explicit nil
func (o *CloudCredentialsForOrganizationEntity) UnsetFullName() {
	o.FullName.Unset()
}

// GetCloudType returns the CloudType field value if set, zero value otherwise.
func (o *CloudCredentialsForOrganizationEntity) GetCloudType() CloudType {
	if o == nil || IsNil(o.CloudType) {
		var ret CloudType
		return ret
	}
	return *o.CloudType
}

// GetCloudTypeOk returns a tuple with the CloudType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudCredentialsForOrganizationEntity) GetCloudTypeOk() (*CloudType, bool) {
	if o == nil || IsNil(o.CloudType) {
		return nil, false
	}
	return o.CloudType, true
}

// HasCloudType returns a boolean if a field has been set.
func (o *CloudCredentialsForOrganizationEntity) HasCloudType() bool {
	if o != nil && !IsNil(o.CloudType) {
		return true
	}

	return false
}

// SetCloudType gets a reference to the given CloudType and assigns it to the CloudType field.
func (o *CloudCredentialsForOrganizationEntity) SetCloudType(v CloudType) {
	o.CloudType = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *CloudCredentialsForOrganizationEntity) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudCredentialsForOrganizationEntity) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *CloudCredentialsForOrganizationEntity) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *CloudCredentialsForOrganizationEntity) SetIsDefault(v bool) {
	o.IsDefault = &v
}

func (o CloudCredentialsForOrganizationEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudCredentialsForOrganizationEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	if o.FullName.IsSet() {
		toSerialize["fullName"] = o.FullName.Get()
	}
	if !IsNil(o.CloudType) {
		toSerialize["cloudType"] = o.CloudType
	}
	if !IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	return toSerialize, nil
}

type NullableCloudCredentialsForOrganizationEntity struct {
	value *CloudCredentialsForOrganizationEntity
	isSet bool
}

func (v NullableCloudCredentialsForOrganizationEntity) Get() *CloudCredentialsForOrganizationEntity {
	return v.value
}

func (v *NullableCloudCredentialsForOrganizationEntity) Set(val *CloudCredentialsForOrganizationEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudCredentialsForOrganizationEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudCredentialsForOrganizationEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudCredentialsForOrganizationEntity(val *CloudCredentialsForOrganizationEntity) *NullableCloudCredentialsForOrganizationEntity {
	return &NullableCloudCredentialsForOrganizationEntity{value: val, isSet: true}
}

func (v NullableCloudCredentialsForOrganizationEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudCredentialsForOrganizationEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


