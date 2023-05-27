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

// checks if the AmazonAvailabilityZonesCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmazonAvailabilityZonesCommand{}

// AmazonAvailabilityZonesCommand struct for AmazonAvailabilityZonesCommand
type AmazonAvailabilityZonesCommand struct {
	Region NullableString `json:"region,omitempty"`
	AwsAccessKeyId NullableString `json:"awsAccessKeyId,omitempty"`
	AwsSecretAccessKey NullableString `json:"awsSecretAccessKey,omitempty"`
	CloudId NullableInt32 `json:"cloudId,omitempty"`
}

// NewAmazonAvailabilityZonesCommand instantiates a new AmazonAvailabilityZonesCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonAvailabilityZonesCommand() *AmazonAvailabilityZonesCommand {
	this := AmazonAvailabilityZonesCommand{}
	return &this
}

// NewAmazonAvailabilityZonesCommandWithDefaults instantiates a new AmazonAvailabilityZonesCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonAvailabilityZonesCommandWithDefaults() *AmazonAvailabilityZonesCommand {
	this := AmazonAvailabilityZonesCommand{}
	return &this
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AmazonAvailabilityZonesCommand) GetRegion() string {
	if o == nil || IsNil(o.Region.Get()) {
		var ret string
		return ret
	}
	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AmazonAvailabilityZonesCommand) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// HasRegion returns a boolean if a field has been set.
func (o *AmazonAvailabilityZonesCommand) HasRegion() bool {
	if o != nil && o.Region.IsSet() {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NullableString and assigns it to the Region field.
func (o *AmazonAvailabilityZonesCommand) SetRegion(v string) {
	o.Region.Set(&v)
}
// SetRegionNil sets the value for Region to be an explicit nil
func (o *AmazonAvailabilityZonesCommand) SetRegionNil() {
	o.Region.Set(nil)
}

// UnsetRegion ensures that no value is present for Region, not even an explicit nil
func (o *AmazonAvailabilityZonesCommand) UnsetRegion() {
	o.Region.Unset()
}

// GetAwsAccessKeyId returns the AwsAccessKeyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AmazonAvailabilityZonesCommand) GetAwsAccessKeyId() string {
	if o == nil || IsNil(o.AwsAccessKeyId.Get()) {
		var ret string
		return ret
	}
	return *o.AwsAccessKeyId.Get()
}

// GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AmazonAvailabilityZonesCommand) GetAwsAccessKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AwsAccessKeyId.Get(), o.AwsAccessKeyId.IsSet()
}

// HasAwsAccessKeyId returns a boolean if a field has been set.
func (o *AmazonAvailabilityZonesCommand) HasAwsAccessKeyId() bool {
	if o != nil && o.AwsAccessKeyId.IsSet() {
		return true
	}

	return false
}

// SetAwsAccessKeyId gets a reference to the given NullableString and assigns it to the AwsAccessKeyId field.
func (o *AmazonAvailabilityZonesCommand) SetAwsAccessKeyId(v string) {
	o.AwsAccessKeyId.Set(&v)
}
// SetAwsAccessKeyIdNil sets the value for AwsAccessKeyId to be an explicit nil
func (o *AmazonAvailabilityZonesCommand) SetAwsAccessKeyIdNil() {
	o.AwsAccessKeyId.Set(nil)
}

// UnsetAwsAccessKeyId ensures that no value is present for AwsAccessKeyId, not even an explicit nil
func (o *AmazonAvailabilityZonesCommand) UnsetAwsAccessKeyId() {
	o.AwsAccessKeyId.Unset()
}

// GetAwsSecretAccessKey returns the AwsSecretAccessKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AmazonAvailabilityZonesCommand) GetAwsSecretAccessKey() string {
	if o == nil || IsNil(o.AwsSecretAccessKey.Get()) {
		var ret string
		return ret
	}
	return *o.AwsSecretAccessKey.Get()
}

// GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AmazonAvailabilityZonesCommand) GetAwsSecretAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AwsSecretAccessKey.Get(), o.AwsSecretAccessKey.IsSet()
}

// HasAwsSecretAccessKey returns a boolean if a field has been set.
func (o *AmazonAvailabilityZonesCommand) HasAwsSecretAccessKey() bool {
	if o != nil && o.AwsSecretAccessKey.IsSet() {
		return true
	}

	return false
}

// SetAwsSecretAccessKey gets a reference to the given NullableString and assigns it to the AwsSecretAccessKey field.
func (o *AmazonAvailabilityZonesCommand) SetAwsSecretAccessKey(v string) {
	o.AwsSecretAccessKey.Set(&v)
}
// SetAwsSecretAccessKeyNil sets the value for AwsSecretAccessKey to be an explicit nil
func (o *AmazonAvailabilityZonesCommand) SetAwsSecretAccessKeyNil() {
	o.AwsSecretAccessKey.Set(nil)
}

// UnsetAwsSecretAccessKey ensures that no value is present for AwsSecretAccessKey, not even an explicit nil
func (o *AmazonAvailabilityZonesCommand) UnsetAwsSecretAccessKey() {
	o.AwsSecretAccessKey.Unset()
}

// GetCloudId returns the CloudId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AmazonAvailabilityZonesCommand) GetCloudId() int32 {
	if o == nil || IsNil(o.CloudId.Get()) {
		var ret int32
		return ret
	}
	return *o.CloudId.Get()
}

// GetCloudIdOk returns a tuple with the CloudId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AmazonAvailabilityZonesCommand) GetCloudIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CloudId.Get(), o.CloudId.IsSet()
}

// HasCloudId returns a boolean if a field has been set.
func (o *AmazonAvailabilityZonesCommand) HasCloudId() bool {
	if o != nil && o.CloudId.IsSet() {
		return true
	}

	return false
}

// SetCloudId gets a reference to the given NullableInt32 and assigns it to the CloudId field.
func (o *AmazonAvailabilityZonesCommand) SetCloudId(v int32) {
	o.CloudId.Set(&v)
}
// SetCloudIdNil sets the value for CloudId to be an explicit nil
func (o *AmazonAvailabilityZonesCommand) SetCloudIdNil() {
	o.CloudId.Set(nil)
}

// UnsetCloudId ensures that no value is present for CloudId, not even an explicit nil
func (o *AmazonAvailabilityZonesCommand) UnsetCloudId() {
	o.CloudId.Unset()
}

func (o AmazonAvailabilityZonesCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmazonAvailabilityZonesCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Region.IsSet() {
		toSerialize["region"] = o.Region.Get()
	}
	if o.AwsAccessKeyId.IsSet() {
		toSerialize["awsAccessKeyId"] = o.AwsAccessKeyId.Get()
	}
	if o.AwsSecretAccessKey.IsSet() {
		toSerialize["awsSecretAccessKey"] = o.AwsSecretAccessKey.Get()
	}
	if o.CloudId.IsSet() {
		toSerialize["cloudId"] = o.CloudId.Get()
	}
	return toSerialize, nil
}

type NullableAmazonAvailabilityZonesCommand struct {
	value *AmazonAvailabilityZonesCommand
	isSet bool
}

func (v NullableAmazonAvailabilityZonesCommand) Get() *AmazonAvailabilityZonesCommand {
	return v.value
}

func (v *NullableAmazonAvailabilityZonesCommand) Set(val *AmazonAvailabilityZonesCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonAvailabilityZonesCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonAvailabilityZonesCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonAvailabilityZonesCommand(val *AmazonAvailabilityZonesCommand) *NullableAmazonAvailabilityZonesCommand {
	return &NullableAmazonAvailabilityZonesCommand{value: val, isSet: true}
}

func (v NullableAmazonAvailabilityZonesCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonAvailabilityZonesCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


