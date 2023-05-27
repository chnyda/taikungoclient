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

// checks if the ProxmoxFlavorData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProxmoxFlavorData{}

// ProxmoxFlavorData struct for ProxmoxFlavorData
type ProxmoxFlavorData struct {
	Name NullableString `json:"name,omitempty"`
	Cpu *int32 `json:"cpu,omitempty"`
	Ram *int64 `json:"ram,omitempty"`
}

// NewProxmoxFlavorData instantiates a new ProxmoxFlavorData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxmoxFlavorData() *ProxmoxFlavorData {
	this := ProxmoxFlavorData{}
	return &this
}

// NewProxmoxFlavorDataWithDefaults instantiates a new ProxmoxFlavorData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxmoxFlavorDataWithDefaults() *ProxmoxFlavorData {
	this := ProxmoxFlavorData{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxmoxFlavorData) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxmoxFlavorData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ProxmoxFlavorData) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ProxmoxFlavorData) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ProxmoxFlavorData) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ProxmoxFlavorData) UnsetName() {
	o.Name.Unset()
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *ProxmoxFlavorData) GetCpu() int32 {
	if o == nil || IsNil(o.Cpu) {
		var ret int32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxmoxFlavorData) GetCpuOk() (*int32, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *ProxmoxFlavorData) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int32 and assigns it to the Cpu field.
func (o *ProxmoxFlavorData) SetCpu(v int32) {
	o.Cpu = &v
}

// GetRam returns the Ram field value if set, zero value otherwise.
func (o *ProxmoxFlavorData) GetRam() int64 {
	if o == nil || IsNil(o.Ram) {
		var ret int64
		return ret
	}
	return *o.Ram
}

// GetRamOk returns a tuple with the Ram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxmoxFlavorData) GetRamOk() (*int64, bool) {
	if o == nil || IsNil(o.Ram) {
		return nil, false
	}
	return o.Ram, true
}

// HasRam returns a boolean if a field has been set.
func (o *ProxmoxFlavorData) HasRam() bool {
	if o != nil && !IsNil(o.Ram) {
		return true
	}

	return false
}

// SetRam gets a reference to the given int64 and assigns it to the Ram field.
func (o *ProxmoxFlavorData) SetRam(v int64) {
	o.Ram = &v
}

func (o ProxmoxFlavorData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProxmoxFlavorData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.Ram) {
		toSerialize["ram"] = o.Ram
	}
	return toSerialize, nil
}

type NullableProxmoxFlavorData struct {
	value *ProxmoxFlavorData
	isSet bool
}

func (v NullableProxmoxFlavorData) Get() *ProxmoxFlavorData {
	return v.value
}

func (v *NullableProxmoxFlavorData) Set(val *ProxmoxFlavorData) {
	v.value = val
	v.isSet = true
}

func (v NullableProxmoxFlavorData) IsSet() bool {
	return v.isSet
}

func (v *NullableProxmoxFlavorData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxmoxFlavorData(val *ProxmoxFlavorData) *NullableProxmoxFlavorData {
	return &NullableProxmoxFlavorData{value: val, isSet: true}
}

func (v NullableProxmoxFlavorData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxmoxFlavorData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


