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

// checks if the GoogleFlavorDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GoogleFlavorDto{}

// GoogleFlavorDto struct for GoogleFlavorDto
type GoogleFlavorDto struct {
	Name *string `json:"name,omitempty"`
	Cpu *int32 `json:"cpu,omitempty"`
	Ram *int64 `json:"ram,omitempty"`
	Description map[string]interface{} `json:"description,omitempty"`
	LinuxPrice *string `json:"linuxPrice,omitempty"`
	WindowsPrice *string `json:"windowsPrice,omitempty"`
	LinuxSpotPrice *string `json:"linuxSpotPrice,omitempty"`
	WindowsSpotPrice *string `json:"windowsSpotPrice,omitempty"`
}

// NewGoogleFlavorDto instantiates a new GoogleFlavorDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleFlavorDto() *GoogleFlavorDto {
	this := GoogleFlavorDto{}
	return &this
}

// NewGoogleFlavorDtoWithDefaults instantiates a new GoogleFlavorDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleFlavorDtoWithDefaults() *GoogleFlavorDto {
	this := GoogleFlavorDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GoogleFlavorDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleFlavorDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GoogleFlavorDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GoogleFlavorDto) SetName(v string) {
	o.Name = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *GoogleFlavorDto) GetCpu() int32 {
	if o == nil || IsNil(o.Cpu) {
		var ret int32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleFlavorDto) GetCpuOk() (*int32, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *GoogleFlavorDto) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int32 and assigns it to the Cpu field.
func (o *GoogleFlavorDto) SetCpu(v int32) {
	o.Cpu = &v
}

// GetRam returns the Ram field value if set, zero value otherwise.
func (o *GoogleFlavorDto) GetRam() int64 {
	if o == nil || IsNil(o.Ram) {
		var ret int64
		return ret
	}
	return *o.Ram
}

// GetRamOk returns a tuple with the Ram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleFlavorDto) GetRamOk() (*int64, bool) {
	if o == nil || IsNil(o.Ram) {
		return nil, false
	}
	return o.Ram, true
}

// HasRam returns a boolean if a field has been set.
func (o *GoogleFlavorDto) HasRam() bool {
	if o != nil && !IsNil(o.Ram) {
		return true
	}

	return false
}

// SetRam gets a reference to the given int64 and assigns it to the Ram field.
func (o *GoogleFlavorDto) SetRam(v int64) {
	o.Ram = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GoogleFlavorDto) GetDescription() map[string]interface{} {
	if o == nil || IsNil(o.Description) {
		var ret map[string]interface{}
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleFlavorDto) GetDescriptionOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Description) {
		return map[string]interface{}{}, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GoogleFlavorDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given map[string]interface{} and assigns it to the Description field.
func (o *GoogleFlavorDto) SetDescription(v map[string]interface{}) {
	o.Description = v
}

// GetLinuxPrice returns the LinuxPrice field value if set, zero value otherwise.
func (o *GoogleFlavorDto) GetLinuxPrice() string {
	if o == nil || IsNil(o.LinuxPrice) {
		var ret string
		return ret
	}
	return *o.LinuxPrice
}

// GetLinuxPriceOk returns a tuple with the LinuxPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleFlavorDto) GetLinuxPriceOk() (*string, bool) {
	if o == nil || IsNil(o.LinuxPrice) {
		return nil, false
	}
	return o.LinuxPrice, true
}

// HasLinuxPrice returns a boolean if a field has been set.
func (o *GoogleFlavorDto) HasLinuxPrice() bool {
	if o != nil && !IsNil(o.LinuxPrice) {
		return true
	}

	return false
}

// SetLinuxPrice gets a reference to the given string and assigns it to the LinuxPrice field.
func (o *GoogleFlavorDto) SetLinuxPrice(v string) {
	o.LinuxPrice = &v
}

// GetWindowsPrice returns the WindowsPrice field value if set, zero value otherwise.
func (o *GoogleFlavorDto) GetWindowsPrice() string {
	if o == nil || IsNil(o.WindowsPrice) {
		var ret string
		return ret
	}
	return *o.WindowsPrice
}

// GetWindowsPriceOk returns a tuple with the WindowsPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleFlavorDto) GetWindowsPriceOk() (*string, bool) {
	if o == nil || IsNil(o.WindowsPrice) {
		return nil, false
	}
	return o.WindowsPrice, true
}

// HasWindowsPrice returns a boolean if a field has been set.
func (o *GoogleFlavorDto) HasWindowsPrice() bool {
	if o != nil && !IsNil(o.WindowsPrice) {
		return true
	}

	return false
}

// SetWindowsPrice gets a reference to the given string and assigns it to the WindowsPrice field.
func (o *GoogleFlavorDto) SetWindowsPrice(v string) {
	o.WindowsPrice = &v
}

// GetLinuxSpotPrice returns the LinuxSpotPrice field value if set, zero value otherwise.
func (o *GoogleFlavorDto) GetLinuxSpotPrice() string {
	if o == nil || IsNil(o.LinuxSpotPrice) {
		var ret string
		return ret
	}
	return *o.LinuxSpotPrice
}

// GetLinuxSpotPriceOk returns a tuple with the LinuxSpotPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleFlavorDto) GetLinuxSpotPriceOk() (*string, bool) {
	if o == nil || IsNil(o.LinuxSpotPrice) {
		return nil, false
	}
	return o.LinuxSpotPrice, true
}

// HasLinuxSpotPrice returns a boolean if a field has been set.
func (o *GoogleFlavorDto) HasLinuxSpotPrice() bool {
	if o != nil && !IsNil(o.LinuxSpotPrice) {
		return true
	}

	return false
}

// SetLinuxSpotPrice gets a reference to the given string and assigns it to the LinuxSpotPrice field.
func (o *GoogleFlavorDto) SetLinuxSpotPrice(v string) {
	o.LinuxSpotPrice = &v
}

// GetWindowsSpotPrice returns the WindowsSpotPrice field value if set, zero value otherwise.
func (o *GoogleFlavorDto) GetWindowsSpotPrice() string {
	if o == nil || IsNil(o.WindowsSpotPrice) {
		var ret string
		return ret
	}
	return *o.WindowsSpotPrice
}

// GetWindowsSpotPriceOk returns a tuple with the WindowsSpotPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleFlavorDto) GetWindowsSpotPriceOk() (*string, bool) {
	if o == nil || IsNil(o.WindowsSpotPrice) {
		return nil, false
	}
	return o.WindowsSpotPrice, true
}

// HasWindowsSpotPrice returns a boolean if a field has been set.
func (o *GoogleFlavorDto) HasWindowsSpotPrice() bool {
	if o != nil && !IsNil(o.WindowsSpotPrice) {
		return true
	}

	return false
}

// SetWindowsSpotPrice gets a reference to the given string and assigns it to the WindowsSpotPrice field.
func (o *GoogleFlavorDto) SetWindowsSpotPrice(v string) {
	o.WindowsSpotPrice = &v
}

func (o GoogleFlavorDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GoogleFlavorDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.Ram) {
		toSerialize["ram"] = o.Ram
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.LinuxPrice) {
		toSerialize["linuxPrice"] = o.LinuxPrice
	}
	if !IsNil(o.WindowsPrice) {
		toSerialize["windowsPrice"] = o.WindowsPrice
	}
	if !IsNil(o.LinuxSpotPrice) {
		toSerialize["linuxSpotPrice"] = o.LinuxSpotPrice
	}
	if !IsNil(o.WindowsSpotPrice) {
		toSerialize["windowsSpotPrice"] = o.WindowsSpotPrice
	}
	return toSerialize, nil
}

type NullableGoogleFlavorDto struct {
	value *GoogleFlavorDto
	isSet bool
}

func (v NullableGoogleFlavorDto) Get() *GoogleFlavorDto {
	return v.value
}

func (v *NullableGoogleFlavorDto) Set(val *GoogleFlavorDto) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleFlavorDto) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleFlavorDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleFlavorDto(val *GoogleFlavorDto) *NullableGoogleFlavorDto {
	return &NullableGoogleFlavorDto{value: val, isSet: true}
}

func (v NullableGoogleFlavorDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleFlavorDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


