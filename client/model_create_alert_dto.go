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

// checks if the CreateAlertDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAlertDto{}

// CreateAlertDto struct for CreateAlertDto
type CreateAlertDto struct {
	Alerts []KubernetesAlertCreateDto `json:"alerts,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewCreateAlertDto instantiates a new CreateAlertDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAlertDto() *CreateAlertDto {
	this := CreateAlertDto{}
	return &this
}

// NewCreateAlertDtoWithDefaults instantiates a new CreateAlertDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAlertDtoWithDefaults() *CreateAlertDto {
	this := CreateAlertDto{}
	return &this
}

// GetAlerts returns the Alerts field value if set, zero value otherwise.
func (o *CreateAlertDto) GetAlerts() []KubernetesAlertCreateDto {
	if o == nil || IsNil(o.Alerts) {
		var ret []KubernetesAlertCreateDto
		return ret
	}
	return o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAlertDto) GetAlertsOk() ([]KubernetesAlertCreateDto, bool) {
	if o == nil || IsNil(o.Alerts) {
		return nil, false
	}
	return o.Alerts, true
}

// HasAlerts returns a boolean if a field has been set.
func (o *CreateAlertDto) HasAlerts() bool {
	if o != nil && !IsNil(o.Alerts) {
		return true
	}

	return false
}

// SetAlerts gets a reference to the given []KubernetesAlertCreateDto and assigns it to the Alerts field.
func (o *CreateAlertDto) SetAlerts(v []KubernetesAlertCreateDto) {
	o.Alerts = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CreateAlertDto) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAlertDto) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateAlertDto) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CreateAlertDto) SetStatus(v string) {
	o.Status = &v
}

func (o CreateAlertDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAlertDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alerts) {
		toSerialize["alerts"] = o.Alerts
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableCreateAlertDto struct {
	value *CreateAlertDto
	isSet bool
}

func (v NullableCreateAlertDto) Get() *CreateAlertDto {
	return v.value
}

func (v *NullableCreateAlertDto) Set(val *CreateAlertDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAlertDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAlertDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAlertDto(val *CreateAlertDto) *NullableCreateAlertDto {
	return &NullableCreateAlertDto{value: val, isSet: true}
}

func (v NullableCreateAlertDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAlertDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


