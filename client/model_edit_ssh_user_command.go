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

// checks if the EditSshUserCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditSshUserCommand{}

// EditSshUserCommand struct for EditSshUserCommand
type EditSshUserCommand struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	SshPublicKey string `json:"sshPublicKey"`
	AccessProfileId int32 `json:"accessProfileId"`
}

// NewEditSshUserCommand instantiates a new EditSshUserCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditSshUserCommand(id int32, name string, sshPublicKey string, accessProfileId int32) *EditSshUserCommand {
	this := EditSshUserCommand{}
	this.Id = id
	this.Name = name
	this.SshPublicKey = sshPublicKey
	this.AccessProfileId = accessProfileId
	return &this
}

// NewEditSshUserCommandWithDefaults instantiates a new EditSshUserCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditSshUserCommandWithDefaults() *EditSshUserCommand {
	this := EditSshUserCommand{}
	return &this
}

// GetId returns the Id field value
func (o *EditSshUserCommand) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EditSshUserCommand) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EditSshUserCommand) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *EditSshUserCommand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EditSshUserCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EditSshUserCommand) SetName(v string) {
	o.Name = v
}

// GetSshPublicKey returns the SshPublicKey field value
func (o *EditSshUserCommand) GetSshPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SshPublicKey
}

// GetSshPublicKeyOk returns a tuple with the SshPublicKey field value
// and a boolean to check if the value has been set.
func (o *EditSshUserCommand) GetSshPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SshPublicKey, true
}

// SetSshPublicKey sets field value
func (o *EditSshUserCommand) SetSshPublicKey(v string) {
	o.SshPublicKey = v
}

// GetAccessProfileId returns the AccessProfileId field value
func (o *EditSshUserCommand) GetAccessProfileId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccessProfileId
}

// GetAccessProfileIdOk returns a tuple with the AccessProfileId field value
// and a boolean to check if the value has been set.
func (o *EditSshUserCommand) GetAccessProfileIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessProfileId, true
}

// SetAccessProfileId sets field value
func (o *EditSshUserCommand) SetAccessProfileId(v int32) {
	o.AccessProfileId = v
}

func (o EditSshUserCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditSshUserCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["sshPublicKey"] = o.SshPublicKey
	toSerialize["accessProfileId"] = o.AccessProfileId
	return toSerialize, nil
}

type NullableEditSshUserCommand struct {
	value *EditSshUserCommand
	isSet bool
}

func (v NullableEditSshUserCommand) Get() *EditSshUserCommand {
	return v.value
}

func (v *NullableEditSshUserCommand) Set(val *EditSshUserCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableEditSshUserCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableEditSshUserCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditSshUserCommand(val *EditSshUserCommand) *NullableEditSshUserCommand {
	return &NullableEditSshUserCommand{value: val, isSet: true}
}

func (v NullableEditSshUserCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditSshUserCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


