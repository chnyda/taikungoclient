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

// checks if the IpAddressRangeListCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpAddressRangeListCommand{}

// IpAddressRangeListCommand struct for IpAddressRangeListCommand
type IpAddressRangeListCommand struct {
	IpAddress NullableString `json:"ipAddress,omitempty"`
	NetMask *int32 `json:"netMask,omitempty"`
	Gateway NullableString `json:"gateway,omitempty"`
}

// NewIpAddressRangeListCommand instantiates a new IpAddressRangeListCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpAddressRangeListCommand() *IpAddressRangeListCommand {
	this := IpAddressRangeListCommand{}
	return &this
}

// NewIpAddressRangeListCommandWithDefaults instantiates a new IpAddressRangeListCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpAddressRangeListCommandWithDefaults() *IpAddressRangeListCommand {
	this := IpAddressRangeListCommand{}
	return &this
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpAddressRangeListCommand) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress.Get()) {
		var ret string
		return ret
	}
	return *o.IpAddress.Get()
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpAddressRangeListCommand) GetIpAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpAddress.Get(), o.IpAddress.IsSet()
}

// HasIpAddress returns a boolean if a field has been set.
func (o *IpAddressRangeListCommand) HasIpAddress() bool {
	if o != nil && o.IpAddress.IsSet() {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given NullableString and assigns it to the IpAddress field.
func (o *IpAddressRangeListCommand) SetIpAddress(v string) {
	o.IpAddress.Set(&v)
}
// SetIpAddressNil sets the value for IpAddress to be an explicit nil
func (o *IpAddressRangeListCommand) SetIpAddressNil() {
	o.IpAddress.Set(nil)
}

// UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
func (o *IpAddressRangeListCommand) UnsetIpAddress() {
	o.IpAddress.Unset()
}

// GetNetMask returns the NetMask field value if set, zero value otherwise.
func (o *IpAddressRangeListCommand) GetNetMask() int32 {
	if o == nil || IsNil(o.NetMask) {
		var ret int32
		return ret
	}
	return *o.NetMask
}

// GetNetMaskOk returns a tuple with the NetMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpAddressRangeListCommand) GetNetMaskOk() (*int32, bool) {
	if o == nil || IsNil(o.NetMask) {
		return nil, false
	}
	return o.NetMask, true
}

// HasNetMask returns a boolean if a field has been set.
func (o *IpAddressRangeListCommand) HasNetMask() bool {
	if o != nil && !IsNil(o.NetMask) {
		return true
	}

	return false
}

// SetNetMask gets a reference to the given int32 and assigns it to the NetMask field.
func (o *IpAddressRangeListCommand) SetNetMask(v int32) {
	o.NetMask = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpAddressRangeListCommand) GetGateway() string {
	if o == nil || IsNil(o.Gateway.Get()) {
		var ret string
		return ret
	}
	return *o.Gateway.Get()
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpAddressRangeListCommand) GetGatewayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gateway.Get(), o.Gateway.IsSet()
}

// HasGateway returns a boolean if a field has been set.
func (o *IpAddressRangeListCommand) HasGateway() bool {
	if o != nil && o.Gateway.IsSet() {
		return true
	}

	return false
}

// SetGateway gets a reference to the given NullableString and assigns it to the Gateway field.
func (o *IpAddressRangeListCommand) SetGateway(v string) {
	o.Gateway.Set(&v)
}
// SetGatewayNil sets the value for Gateway to be an explicit nil
func (o *IpAddressRangeListCommand) SetGatewayNil() {
	o.Gateway.Set(nil)
}

// UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
func (o *IpAddressRangeListCommand) UnsetGateway() {
	o.Gateway.Unset()
}

func (o IpAddressRangeListCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpAddressRangeListCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IpAddress.IsSet() {
		toSerialize["ipAddress"] = o.IpAddress.Get()
	}
	if !IsNil(o.NetMask) {
		toSerialize["netMask"] = o.NetMask
	}
	if o.Gateway.IsSet() {
		toSerialize["gateway"] = o.Gateway.Get()
	}
	return toSerialize, nil
}

type NullableIpAddressRangeListCommand struct {
	value *IpAddressRangeListCommand
	isSet bool
}

func (v NullableIpAddressRangeListCommand) Get() *IpAddressRangeListCommand {
	return v.value
}

func (v *NullableIpAddressRangeListCommand) Set(val *IpAddressRangeListCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableIpAddressRangeListCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableIpAddressRangeListCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpAddressRangeListCommand(val *IpAddressRangeListCommand) *NullableIpAddressRangeListCommand {
	return &NullableIpAddressRangeListCommand{value: val, isSet: true}
}

func (v NullableIpAddressRangeListCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpAddressRangeListCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


