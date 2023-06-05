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

// checks if the OpenstackSubnetListQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenstackSubnetListQuery{}

// OpenstackSubnetListQuery struct for OpenstackSubnetListQuery
type OpenstackSubnetListQuery struct {
	OpenStackUser *string `json:"openStackUser,omitempty"`
	OpenStackPassword *string `json:"openStackPassword,omitempty"`
	OpenStackUrl *string `json:"openStackUrl,omitempty"`
	OpenStackProject *string `json:"openStackProject,omitempty"`
	OpenStackProjectId *string `json:"openStackProjectId,omitempty"`
	OpenStackDomain *string `json:"openStackDomain,omitempty"`
	OpenStackRegion *string `json:"openStackRegion,omitempty"`
	ApplicationCredEnabled *bool `json:"applicationCredEnabled,omitempty"`
}

// NewOpenstackSubnetListQuery instantiates a new OpenstackSubnetListQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenstackSubnetListQuery() *OpenstackSubnetListQuery {
	this := OpenstackSubnetListQuery{}
	return &this
}

// NewOpenstackSubnetListQueryWithDefaults instantiates a new OpenstackSubnetListQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenstackSubnetListQueryWithDefaults() *OpenstackSubnetListQuery {
	this := OpenstackSubnetListQuery{}
	return &this
}

// GetOpenStackUser returns the OpenStackUser field value if set, zero value otherwise.
func (o *OpenstackSubnetListQuery) GetOpenStackUser() string {
	if o == nil || IsNil(o.OpenStackUser) {
		var ret string
		return ret
	}
	return *o.OpenStackUser
}

// GetOpenStackUserOk returns a tuple with the OpenStackUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackSubnetListQuery) GetOpenStackUserOk() (*string, bool) {
	if o == nil || IsNil(o.OpenStackUser) {
		return nil, false
	}
	return o.OpenStackUser, true
}

// HasOpenStackUser returns a boolean if a field has been set.
func (o *OpenstackSubnetListQuery) HasOpenStackUser() bool {
	if o != nil && !IsNil(o.OpenStackUser) {
		return true
	}

	return false
}

// SetOpenStackUser gets a reference to the given string and assigns it to the OpenStackUser field.
func (o *OpenstackSubnetListQuery) SetOpenStackUser(v string) {
	o.OpenStackUser = &v
}

// GetOpenStackPassword returns the OpenStackPassword field value if set, zero value otherwise.
func (o *OpenstackSubnetListQuery) GetOpenStackPassword() string {
	if o == nil || IsNil(o.OpenStackPassword) {
		var ret string
		return ret
	}
	return *o.OpenStackPassword
}

// GetOpenStackPasswordOk returns a tuple with the OpenStackPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackSubnetListQuery) GetOpenStackPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.OpenStackPassword) {
		return nil, false
	}
	return o.OpenStackPassword, true
}

// HasOpenStackPassword returns a boolean if a field has been set.
func (o *OpenstackSubnetListQuery) HasOpenStackPassword() bool {
	if o != nil && !IsNil(o.OpenStackPassword) {
		return true
	}

	return false
}

// SetOpenStackPassword gets a reference to the given string and assigns it to the OpenStackPassword field.
func (o *OpenstackSubnetListQuery) SetOpenStackPassword(v string) {
	o.OpenStackPassword = &v
}

// GetOpenStackUrl returns the OpenStackUrl field value if set, zero value otherwise.
func (o *OpenstackSubnetListQuery) GetOpenStackUrl() string {
	if o == nil || IsNil(o.OpenStackUrl) {
		var ret string
		return ret
	}
	return *o.OpenStackUrl
}

// GetOpenStackUrlOk returns a tuple with the OpenStackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackSubnetListQuery) GetOpenStackUrlOk() (*string, bool) {
	if o == nil || IsNil(o.OpenStackUrl) {
		return nil, false
	}
	return o.OpenStackUrl, true
}

// HasOpenStackUrl returns a boolean if a field has been set.
func (o *OpenstackSubnetListQuery) HasOpenStackUrl() bool {
	if o != nil && !IsNil(o.OpenStackUrl) {
		return true
	}

	return false
}

// SetOpenStackUrl gets a reference to the given string and assigns it to the OpenStackUrl field.
func (o *OpenstackSubnetListQuery) SetOpenStackUrl(v string) {
	o.OpenStackUrl = &v
}

// GetOpenStackProject returns the OpenStackProject field value if set, zero value otherwise.
func (o *OpenstackSubnetListQuery) GetOpenStackProject() string {
	if o == nil || IsNil(o.OpenStackProject) {
		var ret string
		return ret
	}
	return *o.OpenStackProject
}

// GetOpenStackProjectOk returns a tuple with the OpenStackProject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackSubnetListQuery) GetOpenStackProjectOk() (*string, bool) {
	if o == nil || IsNil(o.OpenStackProject) {
		return nil, false
	}
	return o.OpenStackProject, true
}

// HasOpenStackProject returns a boolean if a field has been set.
func (o *OpenstackSubnetListQuery) HasOpenStackProject() bool {
	if o != nil && !IsNil(o.OpenStackProject) {
		return true
	}

	return false
}

// SetOpenStackProject gets a reference to the given string and assigns it to the OpenStackProject field.
func (o *OpenstackSubnetListQuery) SetOpenStackProject(v string) {
	o.OpenStackProject = &v
}

// GetOpenStackProjectId returns the OpenStackProjectId field value if set, zero value otherwise.
func (o *OpenstackSubnetListQuery) GetOpenStackProjectId() string {
	if o == nil || IsNil(o.OpenStackProjectId) {
		var ret string
		return ret
	}
	return *o.OpenStackProjectId
}

// GetOpenStackProjectIdOk returns a tuple with the OpenStackProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackSubnetListQuery) GetOpenStackProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenStackProjectId) {
		return nil, false
	}
	return o.OpenStackProjectId, true
}

// HasOpenStackProjectId returns a boolean if a field has been set.
func (o *OpenstackSubnetListQuery) HasOpenStackProjectId() bool {
	if o != nil && !IsNil(o.OpenStackProjectId) {
		return true
	}

	return false
}

// SetOpenStackProjectId gets a reference to the given string and assigns it to the OpenStackProjectId field.
func (o *OpenstackSubnetListQuery) SetOpenStackProjectId(v string) {
	o.OpenStackProjectId = &v
}

// GetOpenStackDomain returns the OpenStackDomain field value if set, zero value otherwise.
func (o *OpenstackSubnetListQuery) GetOpenStackDomain() string {
	if o == nil || IsNil(o.OpenStackDomain) {
		var ret string
		return ret
	}
	return *o.OpenStackDomain
}

// GetOpenStackDomainOk returns a tuple with the OpenStackDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackSubnetListQuery) GetOpenStackDomainOk() (*string, bool) {
	if o == nil || IsNil(o.OpenStackDomain) {
		return nil, false
	}
	return o.OpenStackDomain, true
}

// HasOpenStackDomain returns a boolean if a field has been set.
func (o *OpenstackSubnetListQuery) HasOpenStackDomain() bool {
	if o != nil && !IsNil(o.OpenStackDomain) {
		return true
	}

	return false
}

// SetOpenStackDomain gets a reference to the given string and assigns it to the OpenStackDomain field.
func (o *OpenstackSubnetListQuery) SetOpenStackDomain(v string) {
	o.OpenStackDomain = &v
}

// GetOpenStackRegion returns the OpenStackRegion field value if set, zero value otherwise.
func (o *OpenstackSubnetListQuery) GetOpenStackRegion() string {
	if o == nil || IsNil(o.OpenStackRegion) {
		var ret string
		return ret
	}
	return *o.OpenStackRegion
}

// GetOpenStackRegionOk returns a tuple with the OpenStackRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackSubnetListQuery) GetOpenStackRegionOk() (*string, bool) {
	if o == nil || IsNil(o.OpenStackRegion) {
		return nil, false
	}
	return o.OpenStackRegion, true
}

// HasOpenStackRegion returns a boolean if a field has been set.
func (o *OpenstackSubnetListQuery) HasOpenStackRegion() bool {
	if o != nil && !IsNil(o.OpenStackRegion) {
		return true
	}

	return false
}

// SetOpenStackRegion gets a reference to the given string and assigns it to the OpenStackRegion field.
func (o *OpenstackSubnetListQuery) SetOpenStackRegion(v string) {
	o.OpenStackRegion = &v
}

// GetApplicationCredEnabled returns the ApplicationCredEnabled field value if set, zero value otherwise.
func (o *OpenstackSubnetListQuery) GetApplicationCredEnabled() bool {
	if o == nil || IsNil(o.ApplicationCredEnabled) {
		var ret bool
		return ret
	}
	return *o.ApplicationCredEnabled
}

// GetApplicationCredEnabledOk returns a tuple with the ApplicationCredEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackSubnetListQuery) GetApplicationCredEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ApplicationCredEnabled) {
		return nil, false
	}
	return o.ApplicationCredEnabled, true
}

// HasApplicationCredEnabled returns a boolean if a field has been set.
func (o *OpenstackSubnetListQuery) HasApplicationCredEnabled() bool {
	if o != nil && !IsNil(o.ApplicationCredEnabled) {
		return true
	}

	return false
}

// SetApplicationCredEnabled gets a reference to the given bool and assigns it to the ApplicationCredEnabled field.
func (o *OpenstackSubnetListQuery) SetApplicationCredEnabled(v bool) {
	o.ApplicationCredEnabled = &v
}

func (o OpenstackSubnetListQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenstackSubnetListQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OpenStackUser) {
		toSerialize["openStackUser"] = o.OpenStackUser
	}
	if !IsNil(o.OpenStackPassword) {
		toSerialize["openStackPassword"] = o.OpenStackPassword
	}
	if !IsNil(o.OpenStackUrl) {
		toSerialize["openStackUrl"] = o.OpenStackUrl
	}
	if !IsNil(o.OpenStackProject) {
		toSerialize["openStackProject"] = o.OpenStackProject
	}
	if !IsNil(o.OpenStackProjectId) {
		toSerialize["openStackProjectId"] = o.OpenStackProjectId
	}
	if !IsNil(o.OpenStackDomain) {
		toSerialize["openStackDomain"] = o.OpenStackDomain
	}
	if !IsNil(o.OpenStackRegion) {
		toSerialize["openStackRegion"] = o.OpenStackRegion
	}
	if !IsNil(o.ApplicationCredEnabled) {
		toSerialize["applicationCredEnabled"] = o.ApplicationCredEnabled
	}
	return toSerialize, nil
}

type NullableOpenstackSubnetListQuery struct {
	value *OpenstackSubnetListQuery
	isSet bool
}

func (v NullableOpenstackSubnetListQuery) Get() *OpenstackSubnetListQuery {
	return v.value
}

func (v *NullableOpenstackSubnetListQuery) Set(val *OpenstackSubnetListQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenstackSubnetListQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenstackSubnetListQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenstackSubnetListQuery(val *OpenstackSubnetListQuery) *NullableOpenstackSubnetListQuery {
	return &NullableOpenstackSubnetListQuery{value: val, isSet: true}
}

func (v NullableOpenstackSubnetListQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenstackSubnetListQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


