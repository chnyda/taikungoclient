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

// checks if the ArchiveTicketCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArchiveTicketCommand{}

// ArchiveTicketCommand struct for ArchiveTicketCommand
type ArchiveTicketCommand struct {
	TicketId string `json:"ticketId"`
}

// NewArchiveTicketCommand instantiates a new ArchiveTicketCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArchiveTicketCommand(ticketId string) *ArchiveTicketCommand {
	this := ArchiveTicketCommand{}
	this.TicketId = ticketId
	return &this
}

// NewArchiveTicketCommandWithDefaults instantiates a new ArchiveTicketCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArchiveTicketCommandWithDefaults() *ArchiveTicketCommand {
	this := ArchiveTicketCommand{}
	return &this
}

// GetTicketId returns the TicketId field value
func (o *ArchiveTicketCommand) GetTicketId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TicketId
}

// GetTicketIdOk returns a tuple with the TicketId field value
// and a boolean to check if the value has been set.
func (o *ArchiveTicketCommand) GetTicketIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TicketId, true
}

// SetTicketId sets field value
func (o *ArchiveTicketCommand) SetTicketId(v string) {
	o.TicketId = v
}

func (o ArchiveTicketCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArchiveTicketCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ticketId"] = o.TicketId
	return toSerialize, nil
}

type NullableArchiveTicketCommand struct {
	value *ArchiveTicketCommand
	isSet bool
}

func (v NullableArchiveTicketCommand) Get() *ArchiveTicketCommand {
	return v.value
}

func (v *NullableArchiveTicketCommand) Set(val *ArchiveTicketCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableArchiveTicketCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableArchiveTicketCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArchiveTicketCommand(val *ArchiveTicketCommand) *NullableArchiveTicketCommand {
	return &NullableArchiveTicketCommand{value: val, isSet: true}
}

func (v NullableArchiveTicketCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArchiveTicketCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

