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

// checks if the PartnerRecordDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartnerRecordDto{}

// PartnerRecordDto struct for PartnerRecordDto
type PartnerRecordDto struct {
	ImageUrl *string `json:"imageUrl,omitempty"`
	Id *int32 `json:"id,omitempty"`
	PaymentEnabled *bool `json:"paymentEnabled,omitempty"`
	AllowRegistration *bool `json:"allowRegistration,omitempty"`
}

// NewPartnerRecordDto instantiates a new PartnerRecordDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerRecordDto() *PartnerRecordDto {
	this := PartnerRecordDto{}
	return &this
}

// NewPartnerRecordDtoWithDefaults instantiates a new PartnerRecordDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerRecordDtoWithDefaults() *PartnerRecordDto {
	this := PartnerRecordDto{}
	return &this
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *PartnerRecordDto) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerRecordDto) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *PartnerRecordDto) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *PartnerRecordDto) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PartnerRecordDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerRecordDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PartnerRecordDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PartnerRecordDto) SetId(v int32) {
	o.Id = &v
}

// GetPaymentEnabled returns the PaymentEnabled field value if set, zero value otherwise.
func (o *PartnerRecordDto) GetPaymentEnabled() bool {
	if o == nil || IsNil(o.PaymentEnabled) {
		var ret bool
		return ret
	}
	return *o.PaymentEnabled
}

// GetPaymentEnabledOk returns a tuple with the PaymentEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerRecordDto) GetPaymentEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PaymentEnabled) {
		return nil, false
	}
	return o.PaymentEnabled, true
}

// HasPaymentEnabled returns a boolean if a field has been set.
func (o *PartnerRecordDto) HasPaymentEnabled() bool {
	if o != nil && !IsNil(o.PaymentEnabled) {
		return true
	}

	return false
}

// SetPaymentEnabled gets a reference to the given bool and assigns it to the PaymentEnabled field.
func (o *PartnerRecordDto) SetPaymentEnabled(v bool) {
	o.PaymentEnabled = &v
}

// GetAllowRegistration returns the AllowRegistration field value if set, zero value otherwise.
func (o *PartnerRecordDto) GetAllowRegistration() bool {
	if o == nil || IsNil(o.AllowRegistration) {
		var ret bool
		return ret
	}
	return *o.AllowRegistration
}

// GetAllowRegistrationOk returns a tuple with the AllowRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerRecordDto) GetAllowRegistrationOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowRegistration) {
		return nil, false
	}
	return o.AllowRegistration, true
}

// HasAllowRegistration returns a boolean if a field has been set.
func (o *PartnerRecordDto) HasAllowRegistration() bool {
	if o != nil && !IsNil(o.AllowRegistration) {
		return true
	}

	return false
}

// SetAllowRegistration gets a reference to the given bool and assigns it to the AllowRegistration field.
func (o *PartnerRecordDto) SetAllowRegistration(v bool) {
	o.AllowRegistration = &v
}

func (o PartnerRecordDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PartnerRecordDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImageUrl) {
		toSerialize["imageUrl"] = o.ImageUrl
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.PaymentEnabled) {
		toSerialize["paymentEnabled"] = o.PaymentEnabled
	}
	if !IsNil(o.AllowRegistration) {
		toSerialize["allowRegistration"] = o.AllowRegistration
	}
	return toSerialize, nil
}

type NullablePartnerRecordDto struct {
	value *PartnerRecordDto
	isSet bool
}

func (v NullablePartnerRecordDto) Get() *PartnerRecordDto {
	return v.value
}

func (v *NullablePartnerRecordDto) Set(val *PartnerRecordDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerRecordDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerRecordDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerRecordDto(val *PartnerRecordDto) *NullablePartnerRecordDto {
	return &NullablePartnerRecordDto{value: val, isSet: true}
}

func (v NullablePartnerRecordDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerRecordDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


