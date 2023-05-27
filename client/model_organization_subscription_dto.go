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
	"time"
)

// checks if the OrganizationSubscriptionDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationSubscriptionDto{}

// OrganizationSubscriptionDto struct for OrganizationSubscriptionDto
type OrganizationSubscriptionDto struct {
	Id *int32 `json:"id,omitempty"`
	OrganizationId *int32 `json:"organizationId,omitempty"`
	OrganizationName NullableString `json:"organizationName,omitempty"`
	SubscriptionId *int32 `json:"subscriptionId,omitempty"`
	StripeSubscriptionId NullableString `json:"stripeSubscriptionId,omitempty"`
	SubscriptionType NullableString `json:"subscriptionType,omitempty"`
	SubscriptionName NullableString `json:"subscriptionName,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty"`
	EndDate NullableTime `json:"endDate,omitempty"`
	Invoices []InvoiceDto `json:"invoices,omitempty"`
}

// NewOrganizationSubscriptionDto instantiates a new OrganizationSubscriptionDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationSubscriptionDto() *OrganizationSubscriptionDto {
	this := OrganizationSubscriptionDto{}
	return &this
}

// NewOrganizationSubscriptionDtoWithDefaults instantiates a new OrganizationSubscriptionDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationSubscriptionDtoWithDefaults() *OrganizationSubscriptionDto {
	this := OrganizationSubscriptionDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationSubscriptionDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSubscriptionDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationSubscriptionDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *OrganizationSubscriptionDto) SetId(v int32) {
	o.Id = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *OrganizationSubscriptionDto) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId) {
		var ret int32
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSubscriptionDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *OrganizationSubscriptionDto) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int32 and assigns it to the OrganizationId field.
func (o *OrganizationSubscriptionDto) SetOrganizationId(v int32) {
	o.OrganizationId = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationSubscriptionDto) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationName.Get()
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationSubscriptionDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationName.Get(), o.OrganizationName.IsSet()
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *OrganizationSubscriptionDto) HasOrganizationName() bool {
	if o != nil && o.OrganizationName.IsSet() {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given NullableString and assigns it to the OrganizationName field.
func (o *OrganizationSubscriptionDto) SetOrganizationName(v string) {
	o.OrganizationName.Set(&v)
}
// SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil
func (o *OrganizationSubscriptionDto) SetOrganizationNameNil() {
	o.OrganizationName.Set(nil)
}

// UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
func (o *OrganizationSubscriptionDto) UnsetOrganizationName() {
	o.OrganizationName.Unset()
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *OrganizationSubscriptionDto) GetSubscriptionId() int32 {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret int32
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSubscriptionDto) GetSubscriptionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *OrganizationSubscriptionDto) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given int32 and assigns it to the SubscriptionId field.
func (o *OrganizationSubscriptionDto) SetSubscriptionId(v int32) {
	o.SubscriptionId = &v
}

// GetStripeSubscriptionId returns the StripeSubscriptionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationSubscriptionDto) GetStripeSubscriptionId() string {
	if o == nil || IsNil(o.StripeSubscriptionId.Get()) {
		var ret string
		return ret
	}
	return *o.StripeSubscriptionId.Get()
}

// GetStripeSubscriptionIdOk returns a tuple with the StripeSubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationSubscriptionDto) GetStripeSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StripeSubscriptionId.Get(), o.StripeSubscriptionId.IsSet()
}

// HasStripeSubscriptionId returns a boolean if a field has been set.
func (o *OrganizationSubscriptionDto) HasStripeSubscriptionId() bool {
	if o != nil && o.StripeSubscriptionId.IsSet() {
		return true
	}

	return false
}

// SetStripeSubscriptionId gets a reference to the given NullableString and assigns it to the StripeSubscriptionId field.
func (o *OrganizationSubscriptionDto) SetStripeSubscriptionId(v string) {
	o.StripeSubscriptionId.Set(&v)
}
// SetStripeSubscriptionIdNil sets the value for StripeSubscriptionId to be an explicit nil
func (o *OrganizationSubscriptionDto) SetStripeSubscriptionIdNil() {
	o.StripeSubscriptionId.Set(nil)
}

// UnsetStripeSubscriptionId ensures that no value is present for StripeSubscriptionId, not even an explicit nil
func (o *OrganizationSubscriptionDto) UnsetStripeSubscriptionId() {
	o.StripeSubscriptionId.Unset()
}

// GetSubscriptionType returns the SubscriptionType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationSubscriptionDto) GetSubscriptionType() string {
	if o == nil || IsNil(o.SubscriptionType.Get()) {
		var ret string
		return ret
	}
	return *o.SubscriptionType.Get()
}

// GetSubscriptionTypeOk returns a tuple with the SubscriptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationSubscriptionDto) GetSubscriptionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionType.Get(), o.SubscriptionType.IsSet()
}

// HasSubscriptionType returns a boolean if a field has been set.
func (o *OrganizationSubscriptionDto) HasSubscriptionType() bool {
	if o != nil && o.SubscriptionType.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionType gets a reference to the given NullableString and assigns it to the SubscriptionType field.
func (o *OrganizationSubscriptionDto) SetSubscriptionType(v string) {
	o.SubscriptionType.Set(&v)
}
// SetSubscriptionTypeNil sets the value for SubscriptionType to be an explicit nil
func (o *OrganizationSubscriptionDto) SetSubscriptionTypeNil() {
	o.SubscriptionType.Set(nil)
}

// UnsetSubscriptionType ensures that no value is present for SubscriptionType, not even an explicit nil
func (o *OrganizationSubscriptionDto) UnsetSubscriptionType() {
	o.SubscriptionType.Unset()
}

// GetSubscriptionName returns the SubscriptionName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationSubscriptionDto) GetSubscriptionName() string {
	if o == nil || IsNil(o.SubscriptionName.Get()) {
		var ret string
		return ret
	}
	return *o.SubscriptionName.Get()
}

// GetSubscriptionNameOk returns a tuple with the SubscriptionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationSubscriptionDto) GetSubscriptionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionName.Get(), o.SubscriptionName.IsSet()
}

// HasSubscriptionName returns a boolean if a field has been set.
func (o *OrganizationSubscriptionDto) HasSubscriptionName() bool {
	if o != nil && o.SubscriptionName.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionName gets a reference to the given NullableString and assigns it to the SubscriptionName field.
func (o *OrganizationSubscriptionDto) SetSubscriptionName(v string) {
	o.SubscriptionName.Set(&v)
}
// SetSubscriptionNameNil sets the value for SubscriptionName to be an explicit nil
func (o *OrganizationSubscriptionDto) SetSubscriptionNameNil() {
	o.SubscriptionName.Set(nil)
}

// UnsetSubscriptionName ensures that no value is present for SubscriptionName, not even an explicit nil
func (o *OrganizationSubscriptionDto) UnsetSubscriptionName() {
	o.SubscriptionName.Unset()
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *OrganizationSubscriptionDto) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSubscriptionDto) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *OrganizationSubscriptionDto) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *OrganizationSubscriptionDto) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationSubscriptionDto) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationSubscriptionDto) GetEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *OrganizationSubscriptionDto) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableTime and assigns it to the EndDate field.
func (o *OrganizationSubscriptionDto) SetEndDate(v time.Time) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *OrganizationSubscriptionDto) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *OrganizationSubscriptionDto) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetInvoices returns the Invoices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationSubscriptionDto) GetInvoices() []InvoiceDto {
	if o == nil {
		var ret []InvoiceDto
		return ret
	}
	return o.Invoices
}

// GetInvoicesOk returns a tuple with the Invoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationSubscriptionDto) GetInvoicesOk() ([]InvoiceDto, bool) {
	if o == nil || IsNil(o.Invoices) {
		return nil, false
	}
	return o.Invoices, true
}

// HasInvoices returns a boolean if a field has been set.
func (o *OrganizationSubscriptionDto) HasInvoices() bool {
	if o != nil && IsNil(o.Invoices) {
		return true
	}

	return false
}

// SetInvoices gets a reference to the given []InvoiceDto and assigns it to the Invoices field.
func (o *OrganizationSubscriptionDto) SetInvoices(v []InvoiceDto) {
	o.Invoices = v
}

func (o OrganizationSubscriptionDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationSubscriptionDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if o.OrganizationName.IsSet() {
		toSerialize["organizationName"] = o.OrganizationName.Get()
	}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if o.StripeSubscriptionId.IsSet() {
		toSerialize["stripeSubscriptionId"] = o.StripeSubscriptionId.Get()
	}
	if o.SubscriptionType.IsSet() {
		toSerialize["subscriptionType"] = o.SubscriptionType.Get()
	}
	if o.SubscriptionName.IsSet() {
		toSerialize["subscriptionName"] = o.SubscriptionName.Get()
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if o.EndDate.IsSet() {
		toSerialize["endDate"] = o.EndDate.Get()
	}
	if o.Invoices != nil {
		toSerialize["invoices"] = o.Invoices
	}
	return toSerialize, nil
}

type NullableOrganizationSubscriptionDto struct {
	value *OrganizationSubscriptionDto
	isSet bool
}

func (v NullableOrganizationSubscriptionDto) Get() *OrganizationSubscriptionDto {
	return v.value
}

func (v *NullableOrganizationSubscriptionDto) Set(val *OrganizationSubscriptionDto) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationSubscriptionDto) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationSubscriptionDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationSubscriptionDto(val *OrganizationSubscriptionDto) *NullableOrganizationSubscriptionDto {
	return &NullableOrganizationSubscriptionDto{value: val, isSet: true}
}

func (v NullableOrganizationSubscriptionDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationSubscriptionDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


