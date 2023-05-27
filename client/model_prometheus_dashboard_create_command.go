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

// checks if the PrometheusDashboardCreateCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrometheusDashboardCreateCommand{}

// PrometheusDashboardCreateCommand struct for PrometheusDashboardCreateCommand
type PrometheusDashboardCreateCommand struct {
	ProjectId *int32 `json:"projectId,omitempty"`
	Name string `json:"name"`
	Expression string `json:"expression"`
	Description string `json:"description"`
	CategoryName string `json:"categoryName"`
}

// NewPrometheusDashboardCreateCommand instantiates a new PrometheusDashboardCreateCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrometheusDashboardCreateCommand(name string, expression string, description string, categoryName string) *PrometheusDashboardCreateCommand {
	this := PrometheusDashboardCreateCommand{}
	this.Name = name
	this.Expression = expression
	this.Description = description
	this.CategoryName = categoryName
	return &this
}

// NewPrometheusDashboardCreateCommandWithDefaults instantiates a new PrometheusDashboardCreateCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrometheusDashboardCreateCommandWithDefaults() *PrometheusDashboardCreateCommand {
	this := PrometheusDashboardCreateCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *PrometheusDashboardCreateCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusDashboardCreateCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *PrometheusDashboardCreateCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *PrometheusDashboardCreateCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetName returns the Name field value
func (o *PrometheusDashboardCreateCommand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PrometheusDashboardCreateCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PrometheusDashboardCreateCommand) SetName(v string) {
	o.Name = v
}

// GetExpression returns the Expression field value
func (o *PrometheusDashboardCreateCommand) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *PrometheusDashboardCreateCommand) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *PrometheusDashboardCreateCommand) SetExpression(v string) {
	o.Expression = v
}

// GetDescription returns the Description field value
func (o *PrometheusDashboardCreateCommand) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *PrometheusDashboardCreateCommand) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *PrometheusDashboardCreateCommand) SetDescription(v string) {
	o.Description = v
}

// GetCategoryName returns the CategoryName field value
func (o *PrometheusDashboardCreateCommand) GetCategoryName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CategoryName
}

// GetCategoryNameOk returns a tuple with the CategoryName field value
// and a boolean to check if the value has been set.
func (o *PrometheusDashboardCreateCommand) GetCategoryNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CategoryName, true
}

// SetCategoryName sets field value
func (o *PrometheusDashboardCreateCommand) SetCategoryName(v string) {
	o.CategoryName = v
}

func (o PrometheusDashboardCreateCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrometheusDashboardCreateCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	toSerialize["name"] = o.Name
	toSerialize["expression"] = o.Expression
	toSerialize["description"] = o.Description
	toSerialize["categoryName"] = o.CategoryName
	return toSerialize, nil
}

type NullablePrometheusDashboardCreateCommand struct {
	value *PrometheusDashboardCreateCommand
	isSet bool
}

func (v NullablePrometheusDashboardCreateCommand) Get() *PrometheusDashboardCreateCommand {
	return v.value
}

func (v *NullablePrometheusDashboardCreateCommand) Set(val *PrometheusDashboardCreateCommand) {
	v.value = val
	v.isSet = true
}

func (v NullablePrometheusDashboardCreateCommand) IsSet() bool {
	return v.isSet
}

func (v *NullablePrometheusDashboardCreateCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrometheusDashboardCreateCommand(val *PrometheusDashboardCreateCommand) *NullablePrometheusDashboardCreateCommand {
	return &NullablePrometheusDashboardCreateCommand{value: val, isSet: true}
}

func (v NullablePrometheusDashboardCreateCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrometheusDashboardCreateCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


