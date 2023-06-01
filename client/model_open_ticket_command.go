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

// checks if the OpenTicketCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenTicketCommand{}

// OpenTicketCommand struct for OpenTicketCommand
type OpenTicketCommand struct {
	TicketId string `json:"ticketId"`
}

// NewOpenTicketCommand instantiates a new OpenTicketCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenTicketCommand(ticketId string) *OpenTicketCommand {
	this := OpenTicketCommand{}
	this.TicketId = ticketId
	return &this
}

// NewOpenTicketCommandWithDefaults instantiates a new OpenTicketCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenTicketCommandWithDefaults() *OpenTicketCommand {
	this := OpenTicketCommand{}
	return &this
}

// GetTicketId returns the TicketId field value
func (o *OpenTicketCommand) GetTicketId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TicketId
}

// GetTicketIdOk returns a tuple with the TicketId field value
// and a boolean to check if the value has been set.
func (o *OpenTicketCommand) GetTicketIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TicketId, true
}

// SetTicketId sets field value
func (o *OpenTicketCommand) SetTicketId(v string) {
	o.TicketId = v
}

func (o OpenTicketCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenTicketCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ticketId"] = o.TicketId
	return toSerialize, nil
}

type NullableOpenTicketCommand struct {
	value *OpenTicketCommand
	isSet bool
}

func (v NullableOpenTicketCommand) Get() *OpenTicketCommand {
	return v.value
}

func (v *NullableOpenTicketCommand) Set(val *OpenTicketCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenTicketCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenTicketCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenTicketCommand(val *OpenTicketCommand) *NullableOpenTicketCommand {
	return &NullableOpenTicketCommand{value: val, isSet: true}
}

func (v NullableOpenTicketCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenTicketCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

