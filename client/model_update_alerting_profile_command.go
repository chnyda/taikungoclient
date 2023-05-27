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

// checks if the UpdateAlertingProfileCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAlertingProfileCommand{}

// UpdateAlertingProfileCommand struct for UpdateAlertingProfileCommand
type UpdateAlertingProfileCommand struct {
	Id int32 `json:"id"`
	Name NullableString `json:"name,omitempty"`
	SlackConfigurationId NullableInt32 `json:"slackConfigurationId,omitempty"`
	OrganizationId NullableInt32 `json:"organizationId,omitempty"`
	Reminder *AlertingReminder `json:"reminder,omitempty"`
}

// NewUpdateAlertingProfileCommand instantiates a new UpdateAlertingProfileCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAlertingProfileCommand(id int32) *UpdateAlertingProfileCommand {
	this := UpdateAlertingProfileCommand{}
	this.Id = id
	return &this
}

// NewUpdateAlertingProfileCommandWithDefaults instantiates a new UpdateAlertingProfileCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAlertingProfileCommandWithDefaults() *UpdateAlertingProfileCommand {
	this := UpdateAlertingProfileCommand{}
	return &this
}

// GetId returns the Id field value
func (o *UpdateAlertingProfileCommand) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateAlertingProfileCommand) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateAlertingProfileCommand) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateAlertingProfileCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateAlertingProfileCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UpdateAlertingProfileCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UpdateAlertingProfileCommand) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *UpdateAlertingProfileCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UpdateAlertingProfileCommand) UnsetName() {
	o.Name.Unset()
}

// GetSlackConfigurationId returns the SlackConfigurationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateAlertingProfileCommand) GetSlackConfigurationId() int32 {
	if o == nil || IsNil(o.SlackConfigurationId.Get()) {
		var ret int32
		return ret
	}
	return *o.SlackConfigurationId.Get()
}

// GetSlackConfigurationIdOk returns a tuple with the SlackConfigurationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateAlertingProfileCommand) GetSlackConfigurationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SlackConfigurationId.Get(), o.SlackConfigurationId.IsSet()
}

// HasSlackConfigurationId returns a boolean if a field has been set.
func (o *UpdateAlertingProfileCommand) HasSlackConfigurationId() bool {
	if o != nil && o.SlackConfigurationId.IsSet() {
		return true
	}

	return false
}

// SetSlackConfigurationId gets a reference to the given NullableInt32 and assigns it to the SlackConfigurationId field.
func (o *UpdateAlertingProfileCommand) SetSlackConfigurationId(v int32) {
	o.SlackConfigurationId.Set(&v)
}
// SetSlackConfigurationIdNil sets the value for SlackConfigurationId to be an explicit nil
func (o *UpdateAlertingProfileCommand) SetSlackConfigurationIdNil() {
	o.SlackConfigurationId.Set(nil)
}

// UnsetSlackConfigurationId ensures that no value is present for SlackConfigurationId, not even an explicit nil
func (o *UpdateAlertingProfileCommand) UnsetSlackConfigurationId() {
	o.SlackConfigurationId.Unset()
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateAlertingProfileCommand) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret int32
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateAlertingProfileCommand) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *UpdateAlertingProfileCommand) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableInt32 and assigns it to the OrganizationId field.
func (o *UpdateAlertingProfileCommand) SetOrganizationId(v int32) {
	o.OrganizationId.Set(&v)
}
// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *UpdateAlertingProfileCommand) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *UpdateAlertingProfileCommand) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

// GetReminder returns the Reminder field value if set, zero value otherwise.
func (o *UpdateAlertingProfileCommand) GetReminder() AlertingReminder {
	if o == nil || IsNil(o.Reminder) {
		var ret AlertingReminder
		return ret
	}
	return *o.Reminder
}

// GetReminderOk returns a tuple with the Reminder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAlertingProfileCommand) GetReminderOk() (*AlertingReminder, bool) {
	if o == nil || IsNil(o.Reminder) {
		return nil, false
	}
	return o.Reminder, true
}

// HasReminder returns a boolean if a field has been set.
func (o *UpdateAlertingProfileCommand) HasReminder() bool {
	if o != nil && !IsNil(o.Reminder) {
		return true
	}

	return false
}

// SetReminder gets a reference to the given AlertingReminder and assigns it to the Reminder field.
func (o *UpdateAlertingProfileCommand) SetReminder(v AlertingReminder) {
	o.Reminder = &v
}

func (o UpdateAlertingProfileCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAlertingProfileCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.SlackConfigurationId.IsSet() {
		toSerialize["slackConfigurationId"] = o.SlackConfigurationId.Get()
	}
	if o.OrganizationId.IsSet() {
		toSerialize["organizationId"] = o.OrganizationId.Get()
	}
	if !IsNil(o.Reminder) {
		toSerialize["reminder"] = o.Reminder
	}
	return toSerialize, nil
}

type NullableUpdateAlertingProfileCommand struct {
	value *UpdateAlertingProfileCommand
	isSet bool
}

func (v NullableUpdateAlertingProfileCommand) Get() *UpdateAlertingProfileCommand {
	return v.value
}

func (v *NullableUpdateAlertingProfileCommand) Set(val *UpdateAlertingProfileCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAlertingProfileCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAlertingProfileCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAlertingProfileCommand(val *UpdateAlertingProfileCommand) *NullableUpdateAlertingProfileCommand {
	return &NullableUpdateAlertingProfileCommand{value: val, isSet: true}
}

func (v NullableUpdateAlertingProfileCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAlertingProfileCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


