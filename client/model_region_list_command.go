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

// checks if the RegionListCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegionListCommand{}

// RegionListCommand struct for RegionListCommand
type RegionListCommand struct {
	AwsAccessKeyId *string `json:"awsAccessKeyId,omitempty"`
	AwsSecretAccessKey *string `json:"awsSecretAccessKey,omitempty"`
}

// NewRegionListCommand instantiates a new RegionListCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionListCommand() *RegionListCommand {
	this := RegionListCommand{}
	return &this
}

// NewRegionListCommandWithDefaults instantiates a new RegionListCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionListCommandWithDefaults() *RegionListCommand {
	this := RegionListCommand{}
	return &this
}

// GetAwsAccessKeyId returns the AwsAccessKeyId field value if set, zero value otherwise.
func (o *RegionListCommand) GetAwsAccessKeyId() string {
	if o == nil || IsNil(o.AwsAccessKeyId) {
		var ret string
		return ret
	}
	return *o.AwsAccessKeyId
}

// GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionListCommand) GetAwsAccessKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.AwsAccessKeyId) {
		return nil, false
	}
	return o.AwsAccessKeyId, true
}

// HasAwsAccessKeyId returns a boolean if a field has been set.
func (o *RegionListCommand) HasAwsAccessKeyId() bool {
	if o != nil && !IsNil(o.AwsAccessKeyId) {
		return true
	}

	return false
}

// SetAwsAccessKeyId gets a reference to the given string and assigns it to the AwsAccessKeyId field.
func (o *RegionListCommand) SetAwsAccessKeyId(v string) {
	o.AwsAccessKeyId = &v
}

// GetAwsSecretAccessKey returns the AwsSecretAccessKey field value if set, zero value otherwise.
func (o *RegionListCommand) GetAwsSecretAccessKey() string {
	if o == nil || IsNil(o.AwsSecretAccessKey) {
		var ret string
		return ret
	}
	return *o.AwsSecretAccessKey
}

// GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionListCommand) GetAwsSecretAccessKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AwsSecretAccessKey) {
		return nil, false
	}
	return o.AwsSecretAccessKey, true
}

// HasAwsSecretAccessKey returns a boolean if a field has been set.
func (o *RegionListCommand) HasAwsSecretAccessKey() bool {
	if o != nil && !IsNil(o.AwsSecretAccessKey) {
		return true
	}

	return false
}

// SetAwsSecretAccessKey gets a reference to the given string and assigns it to the AwsSecretAccessKey field.
func (o *RegionListCommand) SetAwsSecretAccessKey(v string) {
	o.AwsSecretAccessKey = &v
}

func (o RegionListCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegionListCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AwsAccessKeyId) {
		toSerialize["awsAccessKeyId"] = o.AwsAccessKeyId
	}
	if !IsNil(o.AwsSecretAccessKey) {
		toSerialize["awsSecretAccessKey"] = o.AwsSecretAccessKey
	}
	return toSerialize, nil
}

type NullableRegionListCommand struct {
	value *RegionListCommand
	isSet bool
}

func (v NullableRegionListCommand) Get() *RegionListCommand {
	return v.value
}

func (v *NullableRegionListCommand) Set(val *RegionListCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionListCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionListCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionListCommand(val *RegionListCommand) *NullableRegionListCommand {
	return &NullableRegionListCommand{value: val, isSet: true}
}

func (v NullableRegionListCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionListCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


