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

// checks if the AdminUsersUpdateEmailCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminUsersUpdateEmailCommand{}

// AdminUsersUpdateEmailCommand struct for AdminUsersUpdateEmailCommand
type AdminUsersUpdateEmailCommand struct {
	Id string `json:"id"`
	Email NullableString `json:"email,omitempty"`
}

// NewAdminUsersUpdateEmailCommand instantiates a new AdminUsersUpdateEmailCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminUsersUpdateEmailCommand(id string) *AdminUsersUpdateEmailCommand {
	this := AdminUsersUpdateEmailCommand{}
	this.Id = id
	return &this
}

// NewAdminUsersUpdateEmailCommandWithDefaults instantiates a new AdminUsersUpdateEmailCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminUsersUpdateEmailCommandWithDefaults() *AdminUsersUpdateEmailCommand {
	this := AdminUsersUpdateEmailCommand{}
	return &this
}

// GetId returns the Id field value
func (o *AdminUsersUpdateEmailCommand) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AdminUsersUpdateEmailCommand) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AdminUsersUpdateEmailCommand) SetId(v string) {
	o.Id = v
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdminUsersUpdateEmailCommand) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdminUsersUpdateEmailCommand) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *AdminUsersUpdateEmailCommand) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *AdminUsersUpdateEmailCommand) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *AdminUsersUpdateEmailCommand) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *AdminUsersUpdateEmailCommand) UnsetEmail() {
	o.Email.Unset()
}

func (o AdminUsersUpdateEmailCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminUsersUpdateEmailCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	return toSerialize, nil
}

type NullableAdminUsersUpdateEmailCommand struct {
	value *AdminUsersUpdateEmailCommand
	isSet bool
}

func (v NullableAdminUsersUpdateEmailCommand) Get() *AdminUsersUpdateEmailCommand {
	return v.value
}

func (v *NullableAdminUsersUpdateEmailCommand) Set(val *AdminUsersUpdateEmailCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminUsersUpdateEmailCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminUsersUpdateEmailCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminUsersUpdateEmailCommand(val *AdminUsersUpdateEmailCommand) *NullableAdminUsersUpdateEmailCommand {
	return &NullableAdminUsersUpdateEmailCommand{value: val, isSet: true}
}

func (v NullableAdminUsersUpdateEmailCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminUsersUpdateEmailCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


