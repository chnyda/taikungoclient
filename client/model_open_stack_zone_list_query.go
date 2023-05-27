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

// checks if the OpenStackZoneListQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenStackZoneListQuery{}

// OpenStackZoneListQuery struct for OpenStackZoneListQuery
type OpenStackZoneListQuery struct {
	OpenStackUser NullableString `json:"openStackUser,omitempty"`
	OpenStackPassword NullableString `json:"openStackPassword,omitempty"`
	OpenStackUrl NullableString `json:"openStackUrl,omitempty"`
	OpenStackDomain NullableString `json:"openStackDomain,omitempty"`
	OpenStackRegion NullableString `json:"openStackRegion,omitempty"`
	ApplicationCredEnabled *bool `json:"applicationCredEnabled,omitempty"`
	IsAdmin *bool `json:"isAdmin,omitempty"`
	OpenstackProject NullableString `json:"openstackProject,omitempty"`
}

// NewOpenStackZoneListQuery instantiates a new OpenStackZoneListQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenStackZoneListQuery() *OpenStackZoneListQuery {
	this := OpenStackZoneListQuery{}
	return &this
}

// NewOpenStackZoneListQueryWithDefaults instantiates a new OpenStackZoneListQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenStackZoneListQueryWithDefaults() *OpenStackZoneListQuery {
	this := OpenStackZoneListQuery{}
	return &this
}

// GetOpenStackUser returns the OpenStackUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenStackZoneListQuery) GetOpenStackUser() string {
	if o == nil || IsNil(o.OpenStackUser.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackUser.Get()
}

// GetOpenStackUserOk returns a tuple with the OpenStackUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenStackZoneListQuery) GetOpenStackUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackUser.Get(), o.OpenStackUser.IsSet()
}

// HasOpenStackUser returns a boolean if a field has been set.
func (o *OpenStackZoneListQuery) HasOpenStackUser() bool {
	if o != nil && o.OpenStackUser.IsSet() {
		return true
	}

	return false
}

// SetOpenStackUser gets a reference to the given NullableString and assigns it to the OpenStackUser field.
func (o *OpenStackZoneListQuery) SetOpenStackUser(v string) {
	o.OpenStackUser.Set(&v)
}
// SetOpenStackUserNil sets the value for OpenStackUser to be an explicit nil
func (o *OpenStackZoneListQuery) SetOpenStackUserNil() {
	o.OpenStackUser.Set(nil)
}

// UnsetOpenStackUser ensures that no value is present for OpenStackUser, not even an explicit nil
func (o *OpenStackZoneListQuery) UnsetOpenStackUser() {
	o.OpenStackUser.Unset()
}

// GetOpenStackPassword returns the OpenStackPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenStackZoneListQuery) GetOpenStackPassword() string {
	if o == nil || IsNil(o.OpenStackPassword.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackPassword.Get()
}

// GetOpenStackPasswordOk returns a tuple with the OpenStackPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenStackZoneListQuery) GetOpenStackPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackPassword.Get(), o.OpenStackPassword.IsSet()
}

// HasOpenStackPassword returns a boolean if a field has been set.
func (o *OpenStackZoneListQuery) HasOpenStackPassword() bool {
	if o != nil && o.OpenStackPassword.IsSet() {
		return true
	}

	return false
}

// SetOpenStackPassword gets a reference to the given NullableString and assigns it to the OpenStackPassword field.
func (o *OpenStackZoneListQuery) SetOpenStackPassword(v string) {
	o.OpenStackPassword.Set(&v)
}
// SetOpenStackPasswordNil sets the value for OpenStackPassword to be an explicit nil
func (o *OpenStackZoneListQuery) SetOpenStackPasswordNil() {
	o.OpenStackPassword.Set(nil)
}

// UnsetOpenStackPassword ensures that no value is present for OpenStackPassword, not even an explicit nil
func (o *OpenStackZoneListQuery) UnsetOpenStackPassword() {
	o.OpenStackPassword.Unset()
}

// GetOpenStackUrl returns the OpenStackUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenStackZoneListQuery) GetOpenStackUrl() string {
	if o == nil || IsNil(o.OpenStackUrl.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackUrl.Get()
}

// GetOpenStackUrlOk returns a tuple with the OpenStackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenStackZoneListQuery) GetOpenStackUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackUrl.Get(), o.OpenStackUrl.IsSet()
}

// HasOpenStackUrl returns a boolean if a field has been set.
func (o *OpenStackZoneListQuery) HasOpenStackUrl() bool {
	if o != nil && o.OpenStackUrl.IsSet() {
		return true
	}

	return false
}

// SetOpenStackUrl gets a reference to the given NullableString and assigns it to the OpenStackUrl field.
func (o *OpenStackZoneListQuery) SetOpenStackUrl(v string) {
	o.OpenStackUrl.Set(&v)
}
// SetOpenStackUrlNil sets the value for OpenStackUrl to be an explicit nil
func (o *OpenStackZoneListQuery) SetOpenStackUrlNil() {
	o.OpenStackUrl.Set(nil)
}

// UnsetOpenStackUrl ensures that no value is present for OpenStackUrl, not even an explicit nil
func (o *OpenStackZoneListQuery) UnsetOpenStackUrl() {
	o.OpenStackUrl.Unset()
}

// GetOpenStackDomain returns the OpenStackDomain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenStackZoneListQuery) GetOpenStackDomain() string {
	if o == nil || IsNil(o.OpenStackDomain.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackDomain.Get()
}

// GetOpenStackDomainOk returns a tuple with the OpenStackDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenStackZoneListQuery) GetOpenStackDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackDomain.Get(), o.OpenStackDomain.IsSet()
}

// HasOpenStackDomain returns a boolean if a field has been set.
func (o *OpenStackZoneListQuery) HasOpenStackDomain() bool {
	if o != nil && o.OpenStackDomain.IsSet() {
		return true
	}

	return false
}

// SetOpenStackDomain gets a reference to the given NullableString and assigns it to the OpenStackDomain field.
func (o *OpenStackZoneListQuery) SetOpenStackDomain(v string) {
	o.OpenStackDomain.Set(&v)
}
// SetOpenStackDomainNil sets the value for OpenStackDomain to be an explicit nil
func (o *OpenStackZoneListQuery) SetOpenStackDomainNil() {
	o.OpenStackDomain.Set(nil)
}

// UnsetOpenStackDomain ensures that no value is present for OpenStackDomain, not even an explicit nil
func (o *OpenStackZoneListQuery) UnsetOpenStackDomain() {
	o.OpenStackDomain.Unset()
}

// GetOpenStackRegion returns the OpenStackRegion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenStackZoneListQuery) GetOpenStackRegion() string {
	if o == nil || IsNil(o.OpenStackRegion.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackRegion.Get()
}

// GetOpenStackRegionOk returns a tuple with the OpenStackRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenStackZoneListQuery) GetOpenStackRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackRegion.Get(), o.OpenStackRegion.IsSet()
}

// HasOpenStackRegion returns a boolean if a field has been set.
func (o *OpenStackZoneListQuery) HasOpenStackRegion() bool {
	if o != nil && o.OpenStackRegion.IsSet() {
		return true
	}

	return false
}

// SetOpenStackRegion gets a reference to the given NullableString and assigns it to the OpenStackRegion field.
func (o *OpenStackZoneListQuery) SetOpenStackRegion(v string) {
	o.OpenStackRegion.Set(&v)
}
// SetOpenStackRegionNil sets the value for OpenStackRegion to be an explicit nil
func (o *OpenStackZoneListQuery) SetOpenStackRegionNil() {
	o.OpenStackRegion.Set(nil)
}

// UnsetOpenStackRegion ensures that no value is present for OpenStackRegion, not even an explicit nil
func (o *OpenStackZoneListQuery) UnsetOpenStackRegion() {
	o.OpenStackRegion.Unset()
}

// GetApplicationCredEnabled returns the ApplicationCredEnabled field value if set, zero value otherwise.
func (o *OpenStackZoneListQuery) GetApplicationCredEnabled() bool {
	if o == nil || IsNil(o.ApplicationCredEnabled) {
		var ret bool
		return ret
	}
	return *o.ApplicationCredEnabled
}

// GetApplicationCredEnabledOk returns a tuple with the ApplicationCredEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenStackZoneListQuery) GetApplicationCredEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ApplicationCredEnabled) {
		return nil, false
	}
	return o.ApplicationCredEnabled, true
}

// HasApplicationCredEnabled returns a boolean if a field has been set.
func (o *OpenStackZoneListQuery) HasApplicationCredEnabled() bool {
	if o != nil && !IsNil(o.ApplicationCredEnabled) {
		return true
	}

	return false
}

// SetApplicationCredEnabled gets a reference to the given bool and assigns it to the ApplicationCredEnabled field.
func (o *OpenStackZoneListQuery) SetApplicationCredEnabled(v bool) {
	o.ApplicationCredEnabled = &v
}

// GetIsAdmin returns the IsAdmin field value if set, zero value otherwise.
func (o *OpenStackZoneListQuery) GetIsAdmin() bool {
	if o == nil || IsNil(o.IsAdmin) {
		var ret bool
		return ret
	}
	return *o.IsAdmin
}

// GetIsAdminOk returns a tuple with the IsAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenStackZoneListQuery) GetIsAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAdmin) {
		return nil, false
	}
	return o.IsAdmin, true
}

// HasIsAdmin returns a boolean if a field has been set.
func (o *OpenStackZoneListQuery) HasIsAdmin() bool {
	if o != nil && !IsNil(o.IsAdmin) {
		return true
	}

	return false
}

// SetIsAdmin gets a reference to the given bool and assigns it to the IsAdmin field.
func (o *OpenStackZoneListQuery) SetIsAdmin(v bool) {
	o.IsAdmin = &v
}

// GetOpenstackProject returns the OpenstackProject field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenStackZoneListQuery) GetOpenstackProject() string {
	if o == nil || IsNil(o.OpenstackProject.Get()) {
		var ret string
		return ret
	}
	return *o.OpenstackProject.Get()
}

// GetOpenstackProjectOk returns a tuple with the OpenstackProject field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenStackZoneListQuery) GetOpenstackProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenstackProject.Get(), o.OpenstackProject.IsSet()
}

// HasOpenstackProject returns a boolean if a field has been set.
func (o *OpenStackZoneListQuery) HasOpenstackProject() bool {
	if o != nil && o.OpenstackProject.IsSet() {
		return true
	}

	return false
}

// SetOpenstackProject gets a reference to the given NullableString and assigns it to the OpenstackProject field.
func (o *OpenStackZoneListQuery) SetOpenstackProject(v string) {
	o.OpenstackProject.Set(&v)
}
// SetOpenstackProjectNil sets the value for OpenstackProject to be an explicit nil
func (o *OpenStackZoneListQuery) SetOpenstackProjectNil() {
	o.OpenstackProject.Set(nil)
}

// UnsetOpenstackProject ensures that no value is present for OpenstackProject, not even an explicit nil
func (o *OpenStackZoneListQuery) UnsetOpenstackProject() {
	o.OpenstackProject.Unset()
}

func (o OpenStackZoneListQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenStackZoneListQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.OpenStackUser.IsSet() {
		toSerialize["openStackUser"] = o.OpenStackUser.Get()
	}
	if o.OpenStackPassword.IsSet() {
		toSerialize["openStackPassword"] = o.OpenStackPassword.Get()
	}
	if o.OpenStackUrl.IsSet() {
		toSerialize["openStackUrl"] = o.OpenStackUrl.Get()
	}
	if o.OpenStackDomain.IsSet() {
		toSerialize["openStackDomain"] = o.OpenStackDomain.Get()
	}
	if o.OpenStackRegion.IsSet() {
		toSerialize["openStackRegion"] = o.OpenStackRegion.Get()
	}
	if !IsNil(o.ApplicationCredEnabled) {
		toSerialize["applicationCredEnabled"] = o.ApplicationCredEnabled
	}
	if !IsNil(o.IsAdmin) {
		toSerialize["isAdmin"] = o.IsAdmin
	}
	if o.OpenstackProject.IsSet() {
		toSerialize["openstackProject"] = o.OpenstackProject.Get()
	}
	return toSerialize, nil
}

type NullableOpenStackZoneListQuery struct {
	value *OpenStackZoneListQuery
	isSet bool
}

func (v NullableOpenStackZoneListQuery) Get() *OpenStackZoneListQuery {
	return v.value
}

func (v *NullableOpenStackZoneListQuery) Set(val *OpenStackZoneListQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenStackZoneListQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenStackZoneListQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenStackZoneListQuery(val *OpenStackZoneListQuery) *NullableOpenStackZoneListQuery {
	return &NullableOpenStackZoneListQuery{value: val, isSet: true}
}

func (v NullableOpenStackZoneListQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenStackZoneListQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


