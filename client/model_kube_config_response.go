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

// checks if the KubeConfigResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubeConfigResponse{}

// KubeConfigResponse struct for KubeConfigResponse
type KubeConfigResponse struct {
	Data NullableString `json:"data,omitempty"`
}

// NewKubeConfigResponse instantiates a new KubeConfigResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubeConfigResponse() *KubeConfigResponse {
	this := KubeConfigResponse{}
	return &this
}

// NewKubeConfigResponseWithDefaults instantiates a new KubeConfigResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubeConfigResponseWithDefaults() *KubeConfigResponse {
	this := KubeConfigResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubeConfigResponse) GetData() string {
	if o == nil || IsNil(o.Data.Get()) {
		var ret string
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubeConfigResponse) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *KubeConfigResponse) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableString and assigns it to the Data field.
func (o *KubeConfigResponse) SetData(v string) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *KubeConfigResponse) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *KubeConfigResponse) UnsetData() {
	o.Data.Unset()
}

func (o KubeConfigResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubeConfigResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	return toSerialize, nil
}

type NullableKubeConfigResponse struct {
	value *KubeConfigResponse
	isSet bool
}

func (v NullableKubeConfigResponse) Get() *KubeConfigResponse {
	return v.value
}

func (v *NullableKubeConfigResponse) Set(val *KubeConfigResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKubeConfigResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKubeConfigResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubeConfigResponse(val *KubeConfigResponse) *NullableKubeConfigResponse {
	return &NullableKubeConfigResponse{value: val, isSet: true}
}

func (v NullableKubeConfigResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubeConfigResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


