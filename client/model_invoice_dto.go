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

// checks if the InvoiceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceDto{}

// InvoiceDto struct for InvoiceDto
type InvoiceDto struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	DocumentNumber NullableString `json:"documentNumber,omitempty"`
	OrganizationSubscriptionId *int32 `json:"organizationSubscriptionId,omitempty"`
	IsPaid *bool `json:"isPaid,omitempty"`
	RequiredPaymentAction *bool `json:"requiredPaymentAction,omitempty"`
	StripeInvoiceId NullableString `json:"stripeInvoiceId,omitempty"`
	Price *float64 `json:"price,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty"`
	EndDate *time.Time `json:"endDate,omitempty"`
	DueDate *time.Time `json:"dueDate,omitempty"`
}

// NewInvoiceDto instantiates a new InvoiceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceDto() *InvoiceDto {
	this := InvoiceDto{}
	return &this
}

// NewInvoiceDtoWithDefaults instantiates a new InvoiceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceDtoWithDefaults() *InvoiceDto {
	this := InvoiceDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InvoiceDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InvoiceDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *InvoiceDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *InvoiceDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *InvoiceDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *InvoiceDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *InvoiceDto) UnsetName() {
	o.Name.Unset()
}

// GetDocumentNumber returns the DocumentNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceDto) GetDocumentNumber() string {
	if o == nil || IsNil(o.DocumentNumber.Get()) {
		var ret string
		return ret
	}
	return *o.DocumentNumber.Get()
}

// GetDocumentNumberOk returns a tuple with the DocumentNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceDto) GetDocumentNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocumentNumber.Get(), o.DocumentNumber.IsSet()
}

// HasDocumentNumber returns a boolean if a field has been set.
func (o *InvoiceDto) HasDocumentNumber() bool {
	if o != nil && o.DocumentNumber.IsSet() {
		return true
	}

	return false
}

// SetDocumentNumber gets a reference to the given NullableString and assigns it to the DocumentNumber field.
func (o *InvoiceDto) SetDocumentNumber(v string) {
	o.DocumentNumber.Set(&v)
}
// SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil
func (o *InvoiceDto) SetDocumentNumberNil() {
	o.DocumentNumber.Set(nil)
}

// UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
func (o *InvoiceDto) UnsetDocumentNumber() {
	o.DocumentNumber.Unset()
}

// GetOrganizationSubscriptionId returns the OrganizationSubscriptionId field value if set, zero value otherwise.
func (o *InvoiceDto) GetOrganizationSubscriptionId() int32 {
	if o == nil || IsNil(o.OrganizationSubscriptionId) {
		var ret int32
		return ret
	}
	return *o.OrganizationSubscriptionId
}

// GetOrganizationSubscriptionIdOk returns a tuple with the OrganizationSubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDto) GetOrganizationSubscriptionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrganizationSubscriptionId) {
		return nil, false
	}
	return o.OrganizationSubscriptionId, true
}

// HasOrganizationSubscriptionId returns a boolean if a field has been set.
func (o *InvoiceDto) HasOrganizationSubscriptionId() bool {
	if o != nil && !IsNil(o.OrganizationSubscriptionId) {
		return true
	}

	return false
}

// SetOrganizationSubscriptionId gets a reference to the given int32 and assigns it to the OrganizationSubscriptionId field.
func (o *InvoiceDto) SetOrganizationSubscriptionId(v int32) {
	o.OrganizationSubscriptionId = &v
}

// GetIsPaid returns the IsPaid field value if set, zero value otherwise.
func (o *InvoiceDto) GetIsPaid() bool {
	if o == nil || IsNil(o.IsPaid) {
		var ret bool
		return ret
	}
	return *o.IsPaid
}

// GetIsPaidOk returns a tuple with the IsPaid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDto) GetIsPaidOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPaid) {
		return nil, false
	}
	return o.IsPaid, true
}

// HasIsPaid returns a boolean if a field has been set.
func (o *InvoiceDto) HasIsPaid() bool {
	if o != nil && !IsNil(o.IsPaid) {
		return true
	}

	return false
}

// SetIsPaid gets a reference to the given bool and assigns it to the IsPaid field.
func (o *InvoiceDto) SetIsPaid(v bool) {
	o.IsPaid = &v
}

// GetRequiredPaymentAction returns the RequiredPaymentAction field value if set, zero value otherwise.
func (o *InvoiceDto) GetRequiredPaymentAction() bool {
	if o == nil || IsNil(o.RequiredPaymentAction) {
		var ret bool
		return ret
	}
	return *o.RequiredPaymentAction
}

// GetRequiredPaymentActionOk returns a tuple with the RequiredPaymentAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDto) GetRequiredPaymentActionOk() (*bool, bool) {
	if o == nil || IsNil(o.RequiredPaymentAction) {
		return nil, false
	}
	return o.RequiredPaymentAction, true
}

// HasRequiredPaymentAction returns a boolean if a field has been set.
func (o *InvoiceDto) HasRequiredPaymentAction() bool {
	if o != nil && !IsNil(o.RequiredPaymentAction) {
		return true
	}

	return false
}

// SetRequiredPaymentAction gets a reference to the given bool and assigns it to the RequiredPaymentAction field.
func (o *InvoiceDto) SetRequiredPaymentAction(v bool) {
	o.RequiredPaymentAction = &v
}

// GetStripeInvoiceId returns the StripeInvoiceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceDto) GetStripeInvoiceId() string {
	if o == nil || IsNil(o.StripeInvoiceId.Get()) {
		var ret string
		return ret
	}
	return *o.StripeInvoiceId.Get()
}

// GetStripeInvoiceIdOk returns a tuple with the StripeInvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceDto) GetStripeInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StripeInvoiceId.Get(), o.StripeInvoiceId.IsSet()
}

// HasStripeInvoiceId returns a boolean if a field has been set.
func (o *InvoiceDto) HasStripeInvoiceId() bool {
	if o != nil && o.StripeInvoiceId.IsSet() {
		return true
	}

	return false
}

// SetStripeInvoiceId gets a reference to the given NullableString and assigns it to the StripeInvoiceId field.
func (o *InvoiceDto) SetStripeInvoiceId(v string) {
	o.StripeInvoiceId.Set(&v)
}
// SetStripeInvoiceIdNil sets the value for StripeInvoiceId to be an explicit nil
func (o *InvoiceDto) SetStripeInvoiceIdNil() {
	o.StripeInvoiceId.Set(nil)
}

// UnsetStripeInvoiceId ensures that no value is present for StripeInvoiceId, not even an explicit nil
func (o *InvoiceDto) UnsetStripeInvoiceId() {
	o.StripeInvoiceId.Unset()
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *InvoiceDto) GetPrice() float64 {
	if o == nil || IsNil(o.Price) {
		var ret float64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDto) GetPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *InvoiceDto) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float64 and assigns it to the Price field.
func (o *InvoiceDto) SetPrice(v float64) {
	o.Price = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *InvoiceDto) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDto) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *InvoiceDto) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *InvoiceDto) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *InvoiceDto) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDto) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *InvoiceDto) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *InvoiceDto) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *InvoiceDto) GetDueDate() time.Time {
	if o == nil || IsNil(o.DueDate) {
		var ret time.Time
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDto) GetDueDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *InvoiceDto) HasDueDate() bool {
	if o != nil && !IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given time.Time and assigns it to the DueDate field.
func (o *InvoiceDto) SetDueDate(v time.Time) {
	o.DueDate = &v
}

func (o InvoiceDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.DocumentNumber.IsSet() {
		toSerialize["documentNumber"] = o.DocumentNumber.Get()
	}
	if !IsNil(o.OrganizationSubscriptionId) {
		toSerialize["organizationSubscriptionId"] = o.OrganizationSubscriptionId
	}
	if !IsNil(o.IsPaid) {
		toSerialize["isPaid"] = o.IsPaid
	}
	if !IsNil(o.RequiredPaymentAction) {
		toSerialize["requiredPaymentAction"] = o.RequiredPaymentAction
	}
	if o.StripeInvoiceId.IsSet() {
		toSerialize["stripeInvoiceId"] = o.StripeInvoiceId.Get()
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.DueDate) {
		toSerialize["dueDate"] = o.DueDate
	}
	return toSerialize, nil
}

type NullableInvoiceDto struct {
	value *InvoiceDto
	isSet bool
}

func (v NullableInvoiceDto) Get() *InvoiceDto {
	return v.value
}

func (v *NullableInvoiceDto) Set(val *InvoiceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceDto(val *InvoiceDto) *NullableInvoiceDto {
	return &NullableInvoiceDto{value: val, isSet: true}
}

func (v NullableInvoiceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


