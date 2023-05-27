/*
Taikun - Showback API

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@itera.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikunshowback

import (
	"encoding/json"
	"time"
)

// checks if the CreateShowbackSummaryCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateShowbackSummaryCommand{}

// CreateShowbackSummaryCommand struct for CreateShowbackSummaryCommand
type CreateShowbackSummaryCommand struct {
	BeginApply *time.Time `json:"beginApply,omitempty"`
	Price *float64 `json:"price,omitempty"`
	ShowbackRuleId *int32 `json:"showbackRuleId,omitempty"`
	ProjectId *int32 `json:"projectId,omitempty"`
	ByLabelValue *string `json:"byLabelValue,omitempty"`
}

// NewCreateShowbackSummaryCommand instantiates a new CreateShowbackSummaryCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateShowbackSummaryCommand() *CreateShowbackSummaryCommand {
	this := CreateShowbackSummaryCommand{}
	return &this
}

// NewCreateShowbackSummaryCommandWithDefaults instantiates a new CreateShowbackSummaryCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateShowbackSummaryCommandWithDefaults() *CreateShowbackSummaryCommand {
	this := CreateShowbackSummaryCommand{}
	return &this
}

// GetBeginApply returns the BeginApply field value if set, zero value otherwise.
func (o *CreateShowbackSummaryCommand) GetBeginApply() time.Time {
	if o == nil || IsNil(o.BeginApply) {
		var ret time.Time
		return ret
	}
	return *o.BeginApply
}

// GetBeginApplyOk returns a tuple with the BeginApply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShowbackSummaryCommand) GetBeginApplyOk() (*time.Time, bool) {
	if o == nil || IsNil(o.BeginApply) {
		return nil, false
	}
	return o.BeginApply, true
}

// HasBeginApply returns a boolean if a field has been set.
func (o *CreateShowbackSummaryCommand) HasBeginApply() bool {
	if o != nil && !IsNil(o.BeginApply) {
		return true
	}

	return false
}

// SetBeginApply gets a reference to the given time.Time and assigns it to the BeginApply field.
func (o *CreateShowbackSummaryCommand) SetBeginApply(v time.Time) {
	o.BeginApply = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *CreateShowbackSummaryCommand) GetPrice() float64 {
	if o == nil || IsNil(o.Price) {
		var ret float64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShowbackSummaryCommand) GetPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *CreateShowbackSummaryCommand) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float64 and assigns it to the Price field.
func (o *CreateShowbackSummaryCommand) SetPrice(v float64) {
	o.Price = &v
}

// GetShowbackRuleId returns the ShowbackRuleId field value if set, zero value otherwise.
func (o *CreateShowbackSummaryCommand) GetShowbackRuleId() int32 {
	if o == nil || IsNil(o.ShowbackRuleId) {
		var ret int32
		return ret
	}
	return *o.ShowbackRuleId
}

// GetShowbackRuleIdOk returns a tuple with the ShowbackRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShowbackSummaryCommand) GetShowbackRuleIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ShowbackRuleId) {
		return nil, false
	}
	return o.ShowbackRuleId, true
}

// HasShowbackRuleId returns a boolean if a field has been set.
func (o *CreateShowbackSummaryCommand) HasShowbackRuleId() bool {
	if o != nil && !IsNil(o.ShowbackRuleId) {
		return true
	}

	return false
}

// SetShowbackRuleId gets a reference to the given int32 and assigns it to the ShowbackRuleId field.
func (o *CreateShowbackSummaryCommand) SetShowbackRuleId(v int32) {
	o.ShowbackRuleId = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *CreateShowbackSummaryCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShowbackSummaryCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *CreateShowbackSummaryCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *CreateShowbackSummaryCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetByLabelValue returns the ByLabelValue field value if set, zero value otherwise.
func (o *CreateShowbackSummaryCommand) GetByLabelValue() string {
	if o == nil || IsNil(o.ByLabelValue) {
		var ret string
		return ret
	}
	return *o.ByLabelValue
}

// GetByLabelValueOk returns a tuple with the ByLabelValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShowbackSummaryCommand) GetByLabelValueOk() (*string, bool) {
	if o == nil || IsNil(o.ByLabelValue) {
		return nil, false
	}
	return o.ByLabelValue, true
}

// HasByLabelValue returns a boolean if a field has been set.
func (o *CreateShowbackSummaryCommand) HasByLabelValue() bool {
	if o != nil && !IsNil(o.ByLabelValue) {
		return true
	}

	return false
}

// SetByLabelValue gets a reference to the given string and assigns it to the ByLabelValue field.
func (o *CreateShowbackSummaryCommand) SetByLabelValue(v string) {
	o.ByLabelValue = &v
}

func (o CreateShowbackSummaryCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateShowbackSummaryCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BeginApply) {
		toSerialize["beginApply"] = o.BeginApply
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.ShowbackRuleId) {
		toSerialize["showbackRuleId"] = o.ShowbackRuleId
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.ByLabelValue) {
		toSerialize["byLabelValue"] = o.ByLabelValue
	}
	return toSerialize, nil
}

type NullableCreateShowbackSummaryCommand struct {
	value *CreateShowbackSummaryCommand
	isSet bool
}

func (v NullableCreateShowbackSummaryCommand) Get() *CreateShowbackSummaryCommand {
	return v.value
}

func (v *NullableCreateShowbackSummaryCommand) Set(val *CreateShowbackSummaryCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateShowbackSummaryCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateShowbackSummaryCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateShowbackSummaryCommand(val *CreateShowbackSummaryCommand) *NullableCreateShowbackSummaryCommand {
	return &NullableCreateShowbackSummaryCommand{value: val, isSet: true}
}

func (v NullableCreateShowbackSummaryCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateShowbackSummaryCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


