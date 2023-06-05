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

// checks if the KubernetesQuotaListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesQuotaListDto{}

// KubernetesQuotaListDto struct for KubernetesQuotaListDto
type KubernetesQuotaListDto struct {
	SumOfCpu *float64 `json:"sumOfCpu,omitempty"`
	SumOfRamInGb *float64 `json:"sumOfRamInGb,omitempty"`
	SumOfCpuUsage *float64 `json:"sumOfCpuUsage,omitempty"`
	SumOfRamUsage *float64 `json:"sumOfRamUsage,omitempty"`
	PodsCapacity *int32 `json:"podsCapacity,omitempty"`
	PodsTotalCount *int32 `json:"podsTotalCount,omitempty"`
}

// NewKubernetesQuotaListDto instantiates a new KubernetesQuotaListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesQuotaListDto() *KubernetesQuotaListDto {
	this := KubernetesQuotaListDto{}
	return &this
}

// NewKubernetesQuotaListDtoWithDefaults instantiates a new KubernetesQuotaListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesQuotaListDtoWithDefaults() *KubernetesQuotaListDto {
	this := KubernetesQuotaListDto{}
	return &this
}

// GetSumOfCpu returns the SumOfCpu field value if set, zero value otherwise.
func (o *KubernetesQuotaListDto) GetSumOfCpu() float64 {
	if o == nil || IsNil(o.SumOfCpu) {
		var ret float64
		return ret
	}
	return *o.SumOfCpu
}

// GetSumOfCpuOk returns a tuple with the SumOfCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesQuotaListDto) GetSumOfCpuOk() (*float64, bool) {
	if o == nil || IsNil(o.SumOfCpu) {
		return nil, false
	}
	return o.SumOfCpu, true
}

// HasSumOfCpu returns a boolean if a field has been set.
func (o *KubernetesQuotaListDto) HasSumOfCpu() bool {
	if o != nil && !IsNil(o.SumOfCpu) {
		return true
	}

	return false
}

// SetSumOfCpu gets a reference to the given float64 and assigns it to the SumOfCpu field.
func (o *KubernetesQuotaListDto) SetSumOfCpu(v float64) {
	o.SumOfCpu = &v
}

// GetSumOfRamInGb returns the SumOfRamInGb field value if set, zero value otherwise.
func (o *KubernetesQuotaListDto) GetSumOfRamInGb() float64 {
	if o == nil || IsNil(o.SumOfRamInGb) {
		var ret float64
		return ret
	}
	return *o.SumOfRamInGb
}

// GetSumOfRamInGbOk returns a tuple with the SumOfRamInGb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesQuotaListDto) GetSumOfRamInGbOk() (*float64, bool) {
	if o == nil || IsNil(o.SumOfRamInGb) {
		return nil, false
	}
	return o.SumOfRamInGb, true
}

// HasSumOfRamInGb returns a boolean if a field has been set.
func (o *KubernetesQuotaListDto) HasSumOfRamInGb() bool {
	if o != nil && !IsNil(o.SumOfRamInGb) {
		return true
	}

	return false
}

// SetSumOfRamInGb gets a reference to the given float64 and assigns it to the SumOfRamInGb field.
func (o *KubernetesQuotaListDto) SetSumOfRamInGb(v float64) {
	o.SumOfRamInGb = &v
}

// GetSumOfCpuUsage returns the SumOfCpuUsage field value if set, zero value otherwise.
func (o *KubernetesQuotaListDto) GetSumOfCpuUsage() float64 {
	if o == nil || IsNil(o.SumOfCpuUsage) {
		var ret float64
		return ret
	}
	return *o.SumOfCpuUsage
}

// GetSumOfCpuUsageOk returns a tuple with the SumOfCpuUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesQuotaListDto) GetSumOfCpuUsageOk() (*float64, bool) {
	if o == nil || IsNil(o.SumOfCpuUsage) {
		return nil, false
	}
	return o.SumOfCpuUsage, true
}

// HasSumOfCpuUsage returns a boolean if a field has been set.
func (o *KubernetesQuotaListDto) HasSumOfCpuUsage() bool {
	if o != nil && !IsNil(o.SumOfCpuUsage) {
		return true
	}

	return false
}

// SetSumOfCpuUsage gets a reference to the given float64 and assigns it to the SumOfCpuUsage field.
func (o *KubernetesQuotaListDto) SetSumOfCpuUsage(v float64) {
	o.SumOfCpuUsage = &v
}

// GetSumOfRamUsage returns the SumOfRamUsage field value if set, zero value otherwise.
func (o *KubernetesQuotaListDto) GetSumOfRamUsage() float64 {
	if o == nil || IsNil(o.SumOfRamUsage) {
		var ret float64
		return ret
	}
	return *o.SumOfRamUsage
}

// GetSumOfRamUsageOk returns a tuple with the SumOfRamUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesQuotaListDto) GetSumOfRamUsageOk() (*float64, bool) {
	if o == nil || IsNil(o.SumOfRamUsage) {
		return nil, false
	}
	return o.SumOfRamUsage, true
}

// HasSumOfRamUsage returns a boolean if a field has been set.
func (o *KubernetesQuotaListDto) HasSumOfRamUsage() bool {
	if o != nil && !IsNil(o.SumOfRamUsage) {
		return true
	}

	return false
}

// SetSumOfRamUsage gets a reference to the given float64 and assigns it to the SumOfRamUsage field.
func (o *KubernetesQuotaListDto) SetSumOfRamUsage(v float64) {
	o.SumOfRamUsage = &v
}

// GetPodsCapacity returns the PodsCapacity field value if set, zero value otherwise.
func (o *KubernetesQuotaListDto) GetPodsCapacity() int32 {
	if o == nil || IsNil(o.PodsCapacity) {
		var ret int32
		return ret
	}
	return *o.PodsCapacity
}

// GetPodsCapacityOk returns a tuple with the PodsCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesQuotaListDto) GetPodsCapacityOk() (*int32, bool) {
	if o == nil || IsNil(o.PodsCapacity) {
		return nil, false
	}
	return o.PodsCapacity, true
}

// HasPodsCapacity returns a boolean if a field has been set.
func (o *KubernetesQuotaListDto) HasPodsCapacity() bool {
	if o != nil && !IsNil(o.PodsCapacity) {
		return true
	}

	return false
}

// SetPodsCapacity gets a reference to the given int32 and assigns it to the PodsCapacity field.
func (o *KubernetesQuotaListDto) SetPodsCapacity(v int32) {
	o.PodsCapacity = &v
}

// GetPodsTotalCount returns the PodsTotalCount field value if set, zero value otherwise.
func (o *KubernetesQuotaListDto) GetPodsTotalCount() int32 {
	if o == nil || IsNil(o.PodsTotalCount) {
		var ret int32
		return ret
	}
	return *o.PodsTotalCount
}

// GetPodsTotalCountOk returns a tuple with the PodsTotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesQuotaListDto) GetPodsTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PodsTotalCount) {
		return nil, false
	}
	return o.PodsTotalCount, true
}

// HasPodsTotalCount returns a boolean if a field has been set.
func (o *KubernetesQuotaListDto) HasPodsTotalCount() bool {
	if o != nil && !IsNil(o.PodsTotalCount) {
		return true
	}

	return false
}

// SetPodsTotalCount gets a reference to the given int32 and assigns it to the PodsTotalCount field.
func (o *KubernetesQuotaListDto) SetPodsTotalCount(v int32) {
	o.PodsTotalCount = &v
}

func (o KubernetesQuotaListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesQuotaListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SumOfCpu) {
		toSerialize["sumOfCpu"] = o.SumOfCpu
	}
	if !IsNil(o.SumOfRamInGb) {
		toSerialize["sumOfRamInGb"] = o.SumOfRamInGb
	}
	if !IsNil(o.SumOfCpuUsage) {
		toSerialize["sumOfCpuUsage"] = o.SumOfCpuUsage
	}
	if !IsNil(o.SumOfRamUsage) {
		toSerialize["sumOfRamUsage"] = o.SumOfRamUsage
	}
	if !IsNil(o.PodsCapacity) {
		toSerialize["podsCapacity"] = o.PodsCapacity
	}
	if !IsNil(o.PodsTotalCount) {
		toSerialize["podsTotalCount"] = o.PodsTotalCount
	}
	return toSerialize, nil
}

type NullableKubernetesQuotaListDto struct {
	value *KubernetesQuotaListDto
	isSet bool
}

func (v NullableKubernetesQuotaListDto) Get() *KubernetesQuotaListDto {
	return v.value
}

func (v *NullableKubernetesQuotaListDto) Set(val *KubernetesQuotaListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesQuotaListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesQuotaListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesQuotaListDto(val *KubernetesQuotaListDto) *NullableKubernetesQuotaListDto {
	return &NullableKubernetesQuotaListDto{value: val, isSet: true}
}

func (v NullableKubernetesQuotaListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesQuotaListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


