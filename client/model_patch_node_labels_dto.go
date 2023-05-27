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

// checks if the PatchNodeLabelsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchNodeLabelsDto{}

// PatchNodeLabelsDto struct for PatchNodeLabelsDto
type PatchNodeLabelsDto struct {
	Key string `json:"key"`
	Value NullableString `json:"value,omitempty"`
	Mode string `json:"mode"`
}

// NewPatchNodeLabelsDto instantiates a new PatchNodeLabelsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchNodeLabelsDto(key string, mode string) *PatchNodeLabelsDto {
	this := PatchNodeLabelsDto{}
	this.Key = key
	this.Mode = mode
	return &this
}

// NewPatchNodeLabelsDtoWithDefaults instantiates a new PatchNodeLabelsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchNodeLabelsDtoWithDefaults() *PatchNodeLabelsDto {
	this := PatchNodeLabelsDto{}
	return &this
}

// GetKey returns the Key field value
func (o *PatchNodeLabelsDto) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *PatchNodeLabelsDto) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *PatchNodeLabelsDto) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchNodeLabelsDto) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchNodeLabelsDto) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *PatchNodeLabelsDto) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *PatchNodeLabelsDto) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *PatchNodeLabelsDto) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *PatchNodeLabelsDto) UnsetValue() {
	o.Value.Unset()
}

// GetMode returns the Mode field value
func (o *PatchNodeLabelsDto) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *PatchNodeLabelsDto) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *PatchNodeLabelsDto) SetMode(v string) {
	o.Mode = v
}

func (o PatchNodeLabelsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchNodeLabelsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	toSerialize["mode"] = o.Mode
	return toSerialize, nil
}

type NullablePatchNodeLabelsDto struct {
	value *PatchNodeLabelsDto
	isSet bool
}

func (v NullablePatchNodeLabelsDto) Get() *PatchNodeLabelsDto {
	return v.value
}

func (v *NullablePatchNodeLabelsDto) Set(val *PatchNodeLabelsDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchNodeLabelsDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchNodeLabelsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchNodeLabelsDto(val *PatchNodeLabelsDto) *NullablePatchNodeLabelsDto {
	return &NullablePatchNodeLabelsDto{value: val, isSet: true}
}

func (v NullablePatchNodeLabelsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchNodeLabelsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


