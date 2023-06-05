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

// checks if the UnbindFlavorFromProjectCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnbindFlavorFromProjectCommand{}

// UnbindFlavorFromProjectCommand struct for UnbindFlavorFromProjectCommand
type UnbindFlavorFromProjectCommand struct {
	Ids []int32 `json:"ids,omitempty"`
}

// NewUnbindFlavorFromProjectCommand instantiates a new UnbindFlavorFromProjectCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnbindFlavorFromProjectCommand() *UnbindFlavorFromProjectCommand {
	this := UnbindFlavorFromProjectCommand{}
	return &this
}

// NewUnbindFlavorFromProjectCommandWithDefaults instantiates a new UnbindFlavorFromProjectCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnbindFlavorFromProjectCommandWithDefaults() *UnbindFlavorFromProjectCommand {
	this := UnbindFlavorFromProjectCommand{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *UnbindFlavorFromProjectCommand) GetIds() []int32 {
	if o == nil || IsNil(o.Ids) {
		var ret []int32
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnbindFlavorFromProjectCommand) GetIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *UnbindFlavorFromProjectCommand) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []int32 and assigns it to the Ids field.
func (o *UnbindFlavorFromProjectCommand) SetIds(v []int32) {
	o.Ids = v
}

func (o UnbindFlavorFromProjectCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnbindFlavorFromProjectCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	return toSerialize, nil
}

type NullableUnbindFlavorFromProjectCommand struct {
	value *UnbindFlavorFromProjectCommand
	isSet bool
}

func (v NullableUnbindFlavorFromProjectCommand) Get() *UnbindFlavorFromProjectCommand {
	return v.value
}

func (v *NullableUnbindFlavorFromProjectCommand) Set(val *UnbindFlavorFromProjectCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableUnbindFlavorFromProjectCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableUnbindFlavorFromProjectCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnbindFlavorFromProjectCommand(val *UnbindFlavorFromProjectCommand) *NullableUnbindFlavorFromProjectCommand {
	return &NullableUnbindFlavorFromProjectCommand{value: val, isSet: true}
}

func (v NullableUnbindFlavorFromProjectCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnbindFlavorFromProjectCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


