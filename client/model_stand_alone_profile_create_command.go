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

// checks if the StandAloneProfileCreateCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandAloneProfileCreateCommand{}

// StandAloneProfileCreateCommand struct for StandAloneProfileCreateCommand
type StandAloneProfileCreateCommand struct {
	Name string `json:"name"`
	PublicKey string `json:"publicKey"`
	SecurityGroups []StandAloneProfileSecurityGroupDto `json:"securityGroups,omitempty"`
	OrganizationId NullableInt32 `json:"organizationId,omitempty"`
}

// NewStandAloneProfileCreateCommand instantiates a new StandAloneProfileCreateCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandAloneProfileCreateCommand(name string, publicKey string) *StandAloneProfileCreateCommand {
	this := StandAloneProfileCreateCommand{}
	this.Name = name
	this.PublicKey = publicKey
	return &this
}

// NewStandAloneProfileCreateCommandWithDefaults instantiates a new StandAloneProfileCreateCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandAloneProfileCreateCommandWithDefaults() *StandAloneProfileCreateCommand {
	this := StandAloneProfileCreateCommand{}
	return &this
}

// GetName returns the Name field value
func (o *StandAloneProfileCreateCommand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StandAloneProfileCreateCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StandAloneProfileCreateCommand) SetName(v string) {
	o.Name = v
}

// GetPublicKey returns the PublicKey field value
func (o *StandAloneProfileCreateCommand) GetPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *StandAloneProfileCreateCommand) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *StandAloneProfileCreateCommand) SetPublicKey(v string) {
	o.PublicKey = v
}

// GetSecurityGroups returns the SecurityGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandAloneProfileCreateCommand) GetSecurityGroups() []StandAloneProfileSecurityGroupDto {
	if o == nil {
		var ret []StandAloneProfileSecurityGroupDto
		return ret
	}
	return o.SecurityGroups
}

// GetSecurityGroupsOk returns a tuple with the SecurityGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandAloneProfileCreateCommand) GetSecurityGroupsOk() ([]StandAloneProfileSecurityGroupDto, bool) {
	if o == nil || IsNil(o.SecurityGroups) {
		return nil, false
	}
	return o.SecurityGroups, true
}

// HasSecurityGroups returns a boolean if a field has been set.
func (o *StandAloneProfileCreateCommand) HasSecurityGroups() bool {
	if o != nil && IsNil(o.SecurityGroups) {
		return true
	}

	return false
}

// SetSecurityGroups gets a reference to the given []StandAloneProfileSecurityGroupDto and assigns it to the SecurityGroups field.
func (o *StandAloneProfileCreateCommand) SetSecurityGroups(v []StandAloneProfileSecurityGroupDto) {
	o.SecurityGroups = v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandAloneProfileCreateCommand) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret int32
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandAloneProfileCreateCommand) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *StandAloneProfileCreateCommand) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableInt32 and assigns it to the OrganizationId field.
func (o *StandAloneProfileCreateCommand) SetOrganizationId(v int32) {
	o.OrganizationId.Set(&v)
}
// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *StandAloneProfileCreateCommand) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *StandAloneProfileCreateCommand) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

func (o StandAloneProfileCreateCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StandAloneProfileCreateCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["publicKey"] = o.PublicKey
	if o.SecurityGroups != nil {
		toSerialize["securityGroups"] = o.SecurityGroups
	}
	if o.OrganizationId.IsSet() {
		toSerialize["organizationId"] = o.OrganizationId.Get()
	}
	return toSerialize, nil
}

type NullableStandAloneProfileCreateCommand struct {
	value *StandAloneProfileCreateCommand
	isSet bool
}

func (v NullableStandAloneProfileCreateCommand) Get() *StandAloneProfileCreateCommand {
	return v.value
}

func (v *NullableStandAloneProfileCreateCommand) Set(val *StandAloneProfileCreateCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableStandAloneProfileCreateCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableStandAloneProfileCreateCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandAloneProfileCreateCommand(val *StandAloneProfileCreateCommand) *NullableStandAloneProfileCreateCommand {
	return &NullableStandAloneProfileCreateCommand{value: val, isSet: true}
}

func (v NullableStandAloneProfileCreateCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandAloneProfileCreateCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

