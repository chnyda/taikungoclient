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

// checks if the InteractiveShellSendCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InteractiveShellSendCommand{}

// InteractiveShellSendCommand struct for InteractiveShellSendCommand
type InteractiveShellSendCommand struct {
	ProjectId int32 `json:"projectId"`
	Token string `json:"token"`
}

// NewInteractiveShellSendCommand instantiates a new InteractiveShellSendCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInteractiveShellSendCommand(projectId int32, token string) *InteractiveShellSendCommand {
	this := InteractiveShellSendCommand{}
	this.ProjectId = projectId
	this.Token = token
	return &this
}

// NewInteractiveShellSendCommandWithDefaults instantiates a new InteractiveShellSendCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInteractiveShellSendCommandWithDefaults() *InteractiveShellSendCommand {
	this := InteractiveShellSendCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *InteractiveShellSendCommand) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *InteractiveShellSendCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *InteractiveShellSendCommand) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetToken returns the Token field value
func (o *InteractiveShellSendCommand) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *InteractiveShellSendCommand) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *InteractiveShellSendCommand) SetToken(v string) {
	o.Token = v
}

func (o InteractiveShellSendCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InteractiveShellSendCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

type NullableInteractiveShellSendCommand struct {
	value *InteractiveShellSendCommand
	isSet bool
}

func (v NullableInteractiveShellSendCommand) Get() *InteractiveShellSendCommand {
	return v.value
}

func (v *NullableInteractiveShellSendCommand) Set(val *InteractiveShellSendCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableInteractiveShellSendCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableInteractiveShellSendCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInteractiveShellSendCommand(val *InteractiveShellSendCommand) *NullableInteractiveShellSendCommand {
	return &NullableInteractiveShellSendCommand{value: val, isSet: true}
}

func (v NullableInteractiveShellSendCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInteractiveShellSendCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


