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

// checks if the PartnersSearchCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartnersSearchCommand{}

// PartnersSearchCommand struct for PartnersSearchCommand
type PartnersSearchCommand struct {
	Limit *int32 `json:"limit,omitempty"`
	Offset *int32 `json:"offset,omitempty"`
	SearchTerm string `json:"searchTerm"`
}

// NewPartnersSearchCommand instantiates a new PartnersSearchCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnersSearchCommand(searchTerm string) *PartnersSearchCommand {
	this := PartnersSearchCommand{}
	this.SearchTerm = searchTerm
	return &this
}

// NewPartnersSearchCommandWithDefaults instantiates a new PartnersSearchCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnersSearchCommandWithDefaults() *PartnersSearchCommand {
	this := PartnersSearchCommand{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *PartnersSearchCommand) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnersSearchCommand) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *PartnersSearchCommand) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *PartnersSearchCommand) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *PartnersSearchCommand) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnersSearchCommand) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *PartnersSearchCommand) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *PartnersSearchCommand) SetOffset(v int32) {
	o.Offset = &v
}

// GetSearchTerm returns the SearchTerm field value
func (o *PartnersSearchCommand) GetSearchTerm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SearchTerm
}

// GetSearchTermOk returns a tuple with the SearchTerm field value
// and a boolean to check if the value has been set.
func (o *PartnersSearchCommand) GetSearchTermOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SearchTerm, true
}

// SetSearchTerm sets field value
func (o *PartnersSearchCommand) SetSearchTerm(v string) {
	o.SearchTerm = v
}

func (o PartnersSearchCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PartnersSearchCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	toSerialize["searchTerm"] = o.SearchTerm
	return toSerialize, nil
}

type NullablePartnersSearchCommand struct {
	value *PartnersSearchCommand
	isSet bool
}

func (v NullablePartnersSearchCommand) Get() *PartnersSearchCommand {
	return v.value
}

func (v *NullablePartnersSearchCommand) Set(val *PartnersSearchCommand) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnersSearchCommand) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnersSearchCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnersSearchCommand(val *PartnersSearchCommand) *NullablePartnersSearchCommand {
	return &NullablePartnersSearchCommand{value: val, isSet: true}
}

func (v NullablePartnersSearchCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnersSearchCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


