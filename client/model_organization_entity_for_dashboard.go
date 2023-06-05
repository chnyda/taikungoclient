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

// checks if the OrganizationEntityForDashboard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationEntityForDashboard{}

// OrganizationEntityForDashboard struct for OrganizationEntityForDashboard
type OrganizationEntityForDashboard struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Users *int32 `json:"users,omitempty"`
}

// NewOrganizationEntityForDashboard instantiates a new OrganizationEntityForDashboard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationEntityForDashboard() *OrganizationEntityForDashboard {
	this := OrganizationEntityForDashboard{}
	return &this
}

// NewOrganizationEntityForDashboardWithDefaults instantiates a new OrganizationEntityForDashboard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationEntityForDashboardWithDefaults() *OrganizationEntityForDashboard {
	this := OrganizationEntityForDashboard{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationEntityForDashboard) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationEntityForDashboard) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationEntityForDashboard) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *OrganizationEntityForDashboard) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationEntityForDashboard) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationEntityForDashboard) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationEntityForDashboard) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationEntityForDashboard) SetName(v string) {
	o.Name = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *OrganizationEntityForDashboard) GetUsers() int32 {
	if o == nil || IsNil(o.Users) {
		var ret int32
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationEntityForDashboard) GetUsersOk() (*int32, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *OrganizationEntityForDashboard) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given int32 and assigns it to the Users field.
func (o *OrganizationEntityForDashboard) SetUsers(v int32) {
	o.Users = &v
}

func (o OrganizationEntityForDashboard) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationEntityForDashboard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableOrganizationEntityForDashboard struct {
	value *OrganizationEntityForDashboard
	isSet bool
}

func (v NullableOrganizationEntityForDashboard) Get() *OrganizationEntityForDashboard {
	return v.value
}

func (v *NullableOrganizationEntityForDashboard) Set(val *OrganizationEntityForDashboard) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationEntityForDashboard) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationEntityForDashboard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationEntityForDashboard(val *OrganizationEntityForDashboard) *NullableOrganizationEntityForDashboard {
	return &NullableOrganizationEntityForDashboard{value: val, isSet: true}
}

func (v NullableOrganizationEntityForDashboard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationEntityForDashboard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


