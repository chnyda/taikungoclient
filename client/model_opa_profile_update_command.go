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

// checks if the OpaProfileUpdateCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpaProfileUpdateCommand{}

// OpaProfileUpdateCommand struct for OpaProfileUpdateCommand
type OpaProfileUpdateCommand struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	ForbidNodePort NullableBool `json:"forbidNodePort,omitempty"`
	ForbidHttpIngress NullableBool `json:"forbidHttpIngress,omitempty"`
	RequireProbe NullableBool `json:"requireProbe,omitempty"`
	UniqueIngresses NullableBool `json:"uniqueIngresses,omitempty"`
	UniqueServiceSelector NullableBool `json:"uniqueServiceSelector,omitempty"`
	AllowedRepo []string `json:"allowedRepo,omitempty"`
	ForbidSpecificTags []string `json:"forbidSpecificTags,omitempty"`
	IngressWhitelist []string `json:"ingressWhitelist,omitempty"`
}

// NewOpaProfileUpdateCommand instantiates a new OpaProfileUpdateCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpaProfileUpdateCommand(id int32, name string) *OpaProfileUpdateCommand {
	this := OpaProfileUpdateCommand{}
	this.Id = id
	this.Name = name
	return &this
}

// NewOpaProfileUpdateCommandWithDefaults instantiates a new OpaProfileUpdateCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpaProfileUpdateCommandWithDefaults() *OpaProfileUpdateCommand {
	this := OpaProfileUpdateCommand{}
	return &this
}

// GetId returns the Id field value
func (o *OpaProfileUpdateCommand) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OpaProfileUpdateCommand) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OpaProfileUpdateCommand) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *OpaProfileUpdateCommand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OpaProfileUpdateCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OpaProfileUpdateCommand) SetName(v string) {
	o.Name = v
}

// GetForbidNodePort returns the ForbidNodePort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpaProfileUpdateCommand) GetForbidNodePort() bool {
	if o == nil || IsNil(o.ForbidNodePort.Get()) {
		var ret bool
		return ret
	}
	return *o.ForbidNodePort.Get()
}

// GetForbidNodePortOk returns a tuple with the ForbidNodePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpaProfileUpdateCommand) GetForbidNodePortOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ForbidNodePort.Get(), o.ForbidNodePort.IsSet()
}

// HasForbidNodePort returns a boolean if a field has been set.
func (o *OpaProfileUpdateCommand) HasForbidNodePort() bool {
	if o != nil && o.ForbidNodePort.IsSet() {
		return true
	}

	return false
}

// SetForbidNodePort gets a reference to the given NullableBool and assigns it to the ForbidNodePort field.
func (o *OpaProfileUpdateCommand) SetForbidNodePort(v bool) {
	o.ForbidNodePort.Set(&v)
}
// SetForbidNodePortNil sets the value for ForbidNodePort to be an explicit nil
func (o *OpaProfileUpdateCommand) SetForbidNodePortNil() {
	o.ForbidNodePort.Set(nil)
}

// UnsetForbidNodePort ensures that no value is present for ForbidNodePort, not even an explicit nil
func (o *OpaProfileUpdateCommand) UnsetForbidNodePort() {
	o.ForbidNodePort.Unset()
}

// GetForbidHttpIngress returns the ForbidHttpIngress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpaProfileUpdateCommand) GetForbidHttpIngress() bool {
	if o == nil || IsNil(o.ForbidHttpIngress.Get()) {
		var ret bool
		return ret
	}
	return *o.ForbidHttpIngress.Get()
}

// GetForbidHttpIngressOk returns a tuple with the ForbidHttpIngress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpaProfileUpdateCommand) GetForbidHttpIngressOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ForbidHttpIngress.Get(), o.ForbidHttpIngress.IsSet()
}

// HasForbidHttpIngress returns a boolean if a field has been set.
func (o *OpaProfileUpdateCommand) HasForbidHttpIngress() bool {
	if o != nil && o.ForbidHttpIngress.IsSet() {
		return true
	}

	return false
}

// SetForbidHttpIngress gets a reference to the given NullableBool and assigns it to the ForbidHttpIngress field.
func (o *OpaProfileUpdateCommand) SetForbidHttpIngress(v bool) {
	o.ForbidHttpIngress.Set(&v)
}
// SetForbidHttpIngressNil sets the value for ForbidHttpIngress to be an explicit nil
func (o *OpaProfileUpdateCommand) SetForbidHttpIngressNil() {
	o.ForbidHttpIngress.Set(nil)
}

// UnsetForbidHttpIngress ensures that no value is present for ForbidHttpIngress, not even an explicit nil
func (o *OpaProfileUpdateCommand) UnsetForbidHttpIngress() {
	o.ForbidHttpIngress.Unset()
}

// GetRequireProbe returns the RequireProbe field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpaProfileUpdateCommand) GetRequireProbe() bool {
	if o == nil || IsNil(o.RequireProbe.Get()) {
		var ret bool
		return ret
	}
	return *o.RequireProbe.Get()
}

// GetRequireProbeOk returns a tuple with the RequireProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpaProfileUpdateCommand) GetRequireProbeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequireProbe.Get(), o.RequireProbe.IsSet()
}

// HasRequireProbe returns a boolean if a field has been set.
func (o *OpaProfileUpdateCommand) HasRequireProbe() bool {
	if o != nil && o.RequireProbe.IsSet() {
		return true
	}

	return false
}

// SetRequireProbe gets a reference to the given NullableBool and assigns it to the RequireProbe field.
func (o *OpaProfileUpdateCommand) SetRequireProbe(v bool) {
	o.RequireProbe.Set(&v)
}
// SetRequireProbeNil sets the value for RequireProbe to be an explicit nil
func (o *OpaProfileUpdateCommand) SetRequireProbeNil() {
	o.RequireProbe.Set(nil)
}

// UnsetRequireProbe ensures that no value is present for RequireProbe, not even an explicit nil
func (o *OpaProfileUpdateCommand) UnsetRequireProbe() {
	o.RequireProbe.Unset()
}

// GetUniqueIngresses returns the UniqueIngresses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpaProfileUpdateCommand) GetUniqueIngresses() bool {
	if o == nil || IsNil(o.UniqueIngresses.Get()) {
		var ret bool
		return ret
	}
	return *o.UniqueIngresses.Get()
}

// GetUniqueIngressesOk returns a tuple with the UniqueIngresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpaProfileUpdateCommand) GetUniqueIngressesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.UniqueIngresses.Get(), o.UniqueIngresses.IsSet()
}

// HasUniqueIngresses returns a boolean if a field has been set.
func (o *OpaProfileUpdateCommand) HasUniqueIngresses() bool {
	if o != nil && o.UniqueIngresses.IsSet() {
		return true
	}

	return false
}

// SetUniqueIngresses gets a reference to the given NullableBool and assigns it to the UniqueIngresses field.
func (o *OpaProfileUpdateCommand) SetUniqueIngresses(v bool) {
	o.UniqueIngresses.Set(&v)
}
// SetUniqueIngressesNil sets the value for UniqueIngresses to be an explicit nil
func (o *OpaProfileUpdateCommand) SetUniqueIngressesNil() {
	o.UniqueIngresses.Set(nil)
}

// UnsetUniqueIngresses ensures that no value is present for UniqueIngresses, not even an explicit nil
func (o *OpaProfileUpdateCommand) UnsetUniqueIngresses() {
	o.UniqueIngresses.Unset()
}

// GetUniqueServiceSelector returns the UniqueServiceSelector field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpaProfileUpdateCommand) GetUniqueServiceSelector() bool {
	if o == nil || IsNil(o.UniqueServiceSelector.Get()) {
		var ret bool
		return ret
	}
	return *o.UniqueServiceSelector.Get()
}

// GetUniqueServiceSelectorOk returns a tuple with the UniqueServiceSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpaProfileUpdateCommand) GetUniqueServiceSelectorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.UniqueServiceSelector.Get(), o.UniqueServiceSelector.IsSet()
}

// HasUniqueServiceSelector returns a boolean if a field has been set.
func (o *OpaProfileUpdateCommand) HasUniqueServiceSelector() bool {
	if o != nil && o.UniqueServiceSelector.IsSet() {
		return true
	}

	return false
}

// SetUniqueServiceSelector gets a reference to the given NullableBool and assigns it to the UniqueServiceSelector field.
func (o *OpaProfileUpdateCommand) SetUniqueServiceSelector(v bool) {
	o.UniqueServiceSelector.Set(&v)
}
// SetUniqueServiceSelectorNil sets the value for UniqueServiceSelector to be an explicit nil
func (o *OpaProfileUpdateCommand) SetUniqueServiceSelectorNil() {
	o.UniqueServiceSelector.Set(nil)
}

// UnsetUniqueServiceSelector ensures that no value is present for UniqueServiceSelector, not even an explicit nil
func (o *OpaProfileUpdateCommand) UnsetUniqueServiceSelector() {
	o.UniqueServiceSelector.Unset()
}

// GetAllowedRepo returns the AllowedRepo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpaProfileUpdateCommand) GetAllowedRepo() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AllowedRepo
}

// GetAllowedRepoOk returns a tuple with the AllowedRepo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpaProfileUpdateCommand) GetAllowedRepoOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedRepo) {
		return nil, false
	}
	return o.AllowedRepo, true
}

// HasAllowedRepo returns a boolean if a field has been set.
func (o *OpaProfileUpdateCommand) HasAllowedRepo() bool {
	if o != nil && IsNil(o.AllowedRepo) {
		return true
	}

	return false
}

// SetAllowedRepo gets a reference to the given []string and assigns it to the AllowedRepo field.
func (o *OpaProfileUpdateCommand) SetAllowedRepo(v []string) {
	o.AllowedRepo = v
}

// GetForbidSpecificTags returns the ForbidSpecificTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpaProfileUpdateCommand) GetForbidSpecificTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ForbidSpecificTags
}

// GetForbidSpecificTagsOk returns a tuple with the ForbidSpecificTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpaProfileUpdateCommand) GetForbidSpecificTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.ForbidSpecificTags) {
		return nil, false
	}
	return o.ForbidSpecificTags, true
}

// HasForbidSpecificTags returns a boolean if a field has been set.
func (o *OpaProfileUpdateCommand) HasForbidSpecificTags() bool {
	if o != nil && IsNil(o.ForbidSpecificTags) {
		return true
	}

	return false
}

// SetForbidSpecificTags gets a reference to the given []string and assigns it to the ForbidSpecificTags field.
func (o *OpaProfileUpdateCommand) SetForbidSpecificTags(v []string) {
	o.ForbidSpecificTags = v
}

// GetIngressWhitelist returns the IngressWhitelist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpaProfileUpdateCommand) GetIngressWhitelist() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IngressWhitelist
}

// GetIngressWhitelistOk returns a tuple with the IngressWhitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpaProfileUpdateCommand) GetIngressWhitelistOk() ([]string, bool) {
	if o == nil || IsNil(o.IngressWhitelist) {
		return nil, false
	}
	return o.IngressWhitelist, true
}

// HasIngressWhitelist returns a boolean if a field has been set.
func (o *OpaProfileUpdateCommand) HasIngressWhitelist() bool {
	if o != nil && IsNil(o.IngressWhitelist) {
		return true
	}

	return false
}

// SetIngressWhitelist gets a reference to the given []string and assigns it to the IngressWhitelist field.
func (o *OpaProfileUpdateCommand) SetIngressWhitelist(v []string) {
	o.IngressWhitelist = v
}

func (o OpaProfileUpdateCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpaProfileUpdateCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.ForbidNodePort.IsSet() {
		toSerialize["forbidNodePort"] = o.ForbidNodePort.Get()
	}
	if o.ForbidHttpIngress.IsSet() {
		toSerialize["forbidHttpIngress"] = o.ForbidHttpIngress.Get()
	}
	if o.RequireProbe.IsSet() {
		toSerialize["requireProbe"] = o.RequireProbe.Get()
	}
	if o.UniqueIngresses.IsSet() {
		toSerialize["uniqueIngresses"] = o.UniqueIngresses.Get()
	}
	if o.UniqueServiceSelector.IsSet() {
		toSerialize["uniqueServiceSelector"] = o.UniqueServiceSelector.Get()
	}
	if o.AllowedRepo != nil {
		toSerialize["allowedRepo"] = o.AllowedRepo
	}
	if o.ForbidSpecificTags != nil {
		toSerialize["forbidSpecificTags"] = o.ForbidSpecificTags
	}
	if o.IngressWhitelist != nil {
		toSerialize["ingressWhitelist"] = o.IngressWhitelist
	}
	return toSerialize, nil
}

type NullableOpaProfileUpdateCommand struct {
	value *OpaProfileUpdateCommand
	isSet bool
}

func (v NullableOpaProfileUpdateCommand) Get() *OpaProfileUpdateCommand {
	return v.value
}

func (v *NullableOpaProfileUpdateCommand) Set(val *OpaProfileUpdateCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableOpaProfileUpdateCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableOpaProfileUpdateCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpaProfileUpdateCommand(val *OpaProfileUpdateCommand) *NullableOpaProfileUpdateCommand {
	return &NullableOpaProfileUpdateCommand{value: val, isSet: true}
}

func (v NullableOpaProfileUpdateCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpaProfileUpdateCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


