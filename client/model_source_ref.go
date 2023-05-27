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

// checks if the SourceRef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceRef{}

// SourceRef struct for SourceRef
type SourceRef struct {
	Kind NullableString `json:"kind,omitempty"`
	Name NullableString `json:"name,omitempty"`
}

// NewSourceRef instantiates a new SourceRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceRef() *SourceRef {
	this := SourceRef{}
	return &this
}

// NewSourceRefWithDefaults instantiates a new SourceRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceRefWithDefaults() *SourceRef {
	this := SourceRef{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceRef) GetKind() string {
	if o == nil || IsNil(o.Kind.Get()) {
		var ret string
		return ret
	}
	return *o.Kind.Get()
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceRef) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Kind.Get(), o.Kind.IsSet()
}

// HasKind returns a boolean if a field has been set.
func (o *SourceRef) HasKind() bool {
	if o != nil && o.Kind.IsSet() {
		return true
	}

	return false
}

// SetKind gets a reference to the given NullableString and assigns it to the Kind field.
func (o *SourceRef) SetKind(v string) {
	o.Kind.Set(&v)
}
// SetKindNil sets the value for Kind to be an explicit nil
func (o *SourceRef) SetKindNil() {
	o.Kind.Set(nil)
}

// UnsetKind ensures that no value is present for Kind, not even an explicit nil
func (o *SourceRef) UnsetKind() {
	o.Kind.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceRef) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceRef) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SourceRef) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *SourceRef) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *SourceRef) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *SourceRef) UnsetName() {
	o.Name.Unset()
}

func (o SourceRef) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Kind.IsSet() {
		toSerialize["kind"] = o.Kind.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullableSourceRef struct {
	value *SourceRef
	isSet bool
}

func (v NullableSourceRef) Get() *SourceRef {
	return v.value
}

func (v *NullableSourceRef) Set(val *SourceRef) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceRef) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceRef(val *SourceRef) *NullableSourceRef {
	return &NullableSourceRef{value: val, isSet: true}
}

func (v NullableSourceRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


