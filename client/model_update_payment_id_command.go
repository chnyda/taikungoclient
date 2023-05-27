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

// checks if the UpdatePaymentIdCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePaymentIdCommand{}

// UpdatePaymentIdCommand struct for UpdatePaymentIdCommand
type UpdatePaymentIdCommand struct {
	PaymentMethodId NullableString `json:"paymentMethodId,omitempty"`
	PaymentIntentId NullableString `json:"paymentIntentId,omitempty"`
}

// NewUpdatePaymentIdCommand instantiates a new UpdatePaymentIdCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePaymentIdCommand() *UpdatePaymentIdCommand {
	this := UpdatePaymentIdCommand{}
	return &this
}

// NewUpdatePaymentIdCommandWithDefaults instantiates a new UpdatePaymentIdCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePaymentIdCommandWithDefaults() *UpdatePaymentIdCommand {
	this := UpdatePaymentIdCommand{}
	return &this
}

// GetPaymentMethodId returns the PaymentMethodId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdatePaymentIdCommand) GetPaymentMethodId() string {
	if o == nil || IsNil(o.PaymentMethodId.Get()) {
		var ret string
		return ret
	}
	return *o.PaymentMethodId.Get()
}

// GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdatePaymentIdCommand) GetPaymentMethodIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentMethodId.Get(), o.PaymentMethodId.IsSet()
}

// HasPaymentMethodId returns a boolean if a field has been set.
func (o *UpdatePaymentIdCommand) HasPaymentMethodId() bool {
	if o != nil && o.PaymentMethodId.IsSet() {
		return true
	}

	return false
}

// SetPaymentMethodId gets a reference to the given NullableString and assigns it to the PaymentMethodId field.
func (o *UpdatePaymentIdCommand) SetPaymentMethodId(v string) {
	o.PaymentMethodId.Set(&v)
}
// SetPaymentMethodIdNil sets the value for PaymentMethodId to be an explicit nil
func (o *UpdatePaymentIdCommand) SetPaymentMethodIdNil() {
	o.PaymentMethodId.Set(nil)
}

// UnsetPaymentMethodId ensures that no value is present for PaymentMethodId, not even an explicit nil
func (o *UpdatePaymentIdCommand) UnsetPaymentMethodId() {
	o.PaymentMethodId.Unset()
}

// GetPaymentIntentId returns the PaymentIntentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdatePaymentIdCommand) GetPaymentIntentId() string {
	if o == nil || IsNil(o.PaymentIntentId.Get()) {
		var ret string
		return ret
	}
	return *o.PaymentIntentId.Get()
}

// GetPaymentIntentIdOk returns a tuple with the PaymentIntentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdatePaymentIdCommand) GetPaymentIntentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentIntentId.Get(), o.PaymentIntentId.IsSet()
}

// HasPaymentIntentId returns a boolean if a field has been set.
func (o *UpdatePaymentIdCommand) HasPaymentIntentId() bool {
	if o != nil && o.PaymentIntentId.IsSet() {
		return true
	}

	return false
}

// SetPaymentIntentId gets a reference to the given NullableString and assigns it to the PaymentIntentId field.
func (o *UpdatePaymentIdCommand) SetPaymentIntentId(v string) {
	o.PaymentIntentId.Set(&v)
}
// SetPaymentIntentIdNil sets the value for PaymentIntentId to be an explicit nil
func (o *UpdatePaymentIdCommand) SetPaymentIntentIdNil() {
	o.PaymentIntentId.Set(nil)
}

// UnsetPaymentIntentId ensures that no value is present for PaymentIntentId, not even an explicit nil
func (o *UpdatePaymentIdCommand) UnsetPaymentIntentId() {
	o.PaymentIntentId.Unset()
}

func (o UpdatePaymentIdCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePaymentIdCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentMethodId.IsSet() {
		toSerialize["paymentMethodId"] = o.PaymentMethodId.Get()
	}
	if o.PaymentIntentId.IsSet() {
		toSerialize["paymentIntentId"] = o.PaymentIntentId.Get()
	}
	return toSerialize, nil
}

type NullableUpdatePaymentIdCommand struct {
	value *UpdatePaymentIdCommand
	isSet bool
}

func (v NullableUpdatePaymentIdCommand) Get() *UpdatePaymentIdCommand {
	return v.value
}

func (v *NullableUpdatePaymentIdCommand) Set(val *UpdatePaymentIdCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePaymentIdCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePaymentIdCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePaymentIdCommand(val *UpdatePaymentIdCommand) *NullableUpdatePaymentIdCommand {
	return &NullableUpdatePaymentIdCommand{value: val, isSet: true}
}

func (v NullableUpdatePaymentIdCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePaymentIdCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


