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

// checks if the InvoiceListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceListDto{}

// InvoiceListDto struct for InvoiceListDto
type InvoiceListDto struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	StartDate NullableString `json:"startDate,omitempty"`
	EndDate NullableString `json:"endDate,omitempty"`
	RequiredPaymentAction *bool `json:"requiredPaymentAction,omitempty"`
	IsPaid *bool `json:"isPaid,omitempty"`
	InvoiceId NullableString `json:"invoiceId,omitempty"`
	SubscriptionType NullableString `json:"subscriptionType,omitempty"`
	SubscriptionName NullableString `json:"subscriptionName,omitempty"`
	Price *float64 `json:"price,omitempty"`
	OrganizationId *int32 `json:"organizationId,omitempty"`
	OrganizationName NullableString `json:"organizationName,omitempty"`
	InvoiceNumber NullableString `json:"invoiceNumber,omitempty"`
	OrganizationSubscriptionId *int32 `json:"organizationSubscriptionId,omitempty"`
}

// NewInvoiceListDto instantiates a new InvoiceListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceListDto() *InvoiceListDto {
	this := InvoiceListDto{}
	return &this
}

// NewInvoiceListDtoWithDefaults instantiates a new InvoiceListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceListDtoWithDefaults() *InvoiceListDto {
	this := InvoiceListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InvoiceListDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceListDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InvoiceListDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *InvoiceListDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceListDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *InvoiceListDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *InvoiceListDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *InvoiceListDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *InvoiceListDto) UnsetName() {
	o.Name.Unset()
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceListDto) GetStartDate() string {
	if o == nil || IsNil(o.StartDate.Get()) {
		var ret string
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceListDto) GetStartDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *InvoiceListDto) HasStartDate() bool {
	if o != nil && o.StartDate.IsSet() {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given NullableString and assigns it to the StartDate field.
func (o *InvoiceListDto) SetStartDate(v string) {
	o.StartDate.Set(&v)
}
// SetStartDateNil sets the value for StartDate to be an explicit nil
func (o *InvoiceListDto) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
func (o *InvoiceListDto) UnsetStartDate() {
	o.StartDate.Unset()
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceListDto) GetEndDate() string {
	if o == nil || IsNil(o.EndDate.Get()) {
		var ret string
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceListDto) GetEndDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *InvoiceListDto) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableString and assigns it to the EndDate field.
func (o *InvoiceListDto) SetEndDate(v string) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *InvoiceListDto) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *InvoiceListDto) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetRequiredPaymentAction returns the RequiredPaymentAction field value if set, zero value otherwise.
func (o *InvoiceListDto) GetRequiredPaymentAction() bool {
	if o == nil || IsNil(o.RequiredPaymentAction) {
		var ret bool
		return ret
	}
	return *o.RequiredPaymentAction
}

// GetRequiredPaymentActionOk returns a tuple with the RequiredPaymentAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceListDto) GetRequiredPaymentActionOk() (*bool, bool) {
	if o == nil || IsNil(o.RequiredPaymentAction) {
		return nil, false
	}
	return o.RequiredPaymentAction, true
}

// HasRequiredPaymentAction returns a boolean if a field has been set.
func (o *InvoiceListDto) HasRequiredPaymentAction() bool {
	if o != nil && !IsNil(o.RequiredPaymentAction) {
		return true
	}

	return false
}

// SetRequiredPaymentAction gets a reference to the given bool and assigns it to the RequiredPaymentAction field.
func (o *InvoiceListDto) SetRequiredPaymentAction(v bool) {
	o.RequiredPaymentAction = &v
}

// GetIsPaid returns the IsPaid field value if set, zero value otherwise.
func (o *InvoiceListDto) GetIsPaid() bool {
	if o == nil || IsNil(o.IsPaid) {
		var ret bool
		return ret
	}
	return *o.IsPaid
}

// GetIsPaidOk returns a tuple with the IsPaid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceListDto) GetIsPaidOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPaid) {
		return nil, false
	}
	return o.IsPaid, true
}

// HasIsPaid returns a boolean if a field has been set.
func (o *InvoiceListDto) HasIsPaid() bool {
	if o != nil && !IsNil(o.IsPaid) {
		return true
	}

	return false
}

// SetIsPaid gets a reference to the given bool and assigns it to the IsPaid field.
func (o *InvoiceListDto) SetIsPaid(v bool) {
	o.IsPaid = &v
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceListDto) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId.Get()) {
		var ret string
		return ret
	}
	return *o.InvoiceId.Get()
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceListDto) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceId.Get(), o.InvoiceId.IsSet()
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *InvoiceListDto) HasInvoiceId() bool {
	if o != nil && o.InvoiceId.IsSet() {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given NullableString and assigns it to the InvoiceId field.
func (o *InvoiceListDto) SetInvoiceId(v string) {
	o.InvoiceId.Set(&v)
}
// SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil
func (o *InvoiceListDto) SetInvoiceIdNil() {
	o.InvoiceId.Set(nil)
}

// UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
func (o *InvoiceListDto) UnsetInvoiceId() {
	o.InvoiceId.Unset()
}

// GetSubscriptionType returns the SubscriptionType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceListDto) GetSubscriptionType() string {
	if o == nil || IsNil(o.SubscriptionType.Get()) {
		var ret string
		return ret
	}
	return *o.SubscriptionType.Get()
}

// GetSubscriptionTypeOk returns a tuple with the SubscriptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceListDto) GetSubscriptionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionType.Get(), o.SubscriptionType.IsSet()
}

// HasSubscriptionType returns a boolean if a field has been set.
func (o *InvoiceListDto) HasSubscriptionType() bool {
	if o != nil && o.SubscriptionType.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionType gets a reference to the given NullableString and assigns it to the SubscriptionType field.
func (o *InvoiceListDto) SetSubscriptionType(v string) {
	o.SubscriptionType.Set(&v)
}
// SetSubscriptionTypeNil sets the value for SubscriptionType to be an explicit nil
func (o *InvoiceListDto) SetSubscriptionTypeNil() {
	o.SubscriptionType.Set(nil)
}

// UnsetSubscriptionType ensures that no value is present for SubscriptionType, not even an explicit nil
func (o *InvoiceListDto) UnsetSubscriptionType() {
	o.SubscriptionType.Unset()
}

// GetSubscriptionName returns the SubscriptionName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceListDto) GetSubscriptionName() string {
	if o == nil || IsNil(o.SubscriptionName.Get()) {
		var ret string
		return ret
	}
	return *o.SubscriptionName.Get()
}

// GetSubscriptionNameOk returns a tuple with the SubscriptionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceListDto) GetSubscriptionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionName.Get(), o.SubscriptionName.IsSet()
}

// HasSubscriptionName returns a boolean if a field has been set.
func (o *InvoiceListDto) HasSubscriptionName() bool {
	if o != nil && o.SubscriptionName.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionName gets a reference to the given NullableString and assigns it to the SubscriptionName field.
func (o *InvoiceListDto) SetSubscriptionName(v string) {
	o.SubscriptionName.Set(&v)
}
// SetSubscriptionNameNil sets the value for SubscriptionName to be an explicit nil
func (o *InvoiceListDto) SetSubscriptionNameNil() {
	o.SubscriptionName.Set(nil)
}

// UnsetSubscriptionName ensures that no value is present for SubscriptionName, not even an explicit nil
func (o *InvoiceListDto) UnsetSubscriptionName() {
	o.SubscriptionName.Unset()
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *InvoiceListDto) GetPrice() float64 {
	if o == nil || IsNil(o.Price) {
		var ret float64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceListDto) GetPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *InvoiceListDto) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float64 and assigns it to the Price field.
func (o *InvoiceListDto) SetPrice(v float64) {
	o.Price = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *InvoiceListDto) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId) {
		var ret int32
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceListDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *InvoiceListDto) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int32 and assigns it to the OrganizationId field.
func (o *InvoiceListDto) SetOrganizationId(v int32) {
	o.OrganizationId = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceListDto) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationName.Get()
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceListDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationName.Get(), o.OrganizationName.IsSet()
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *InvoiceListDto) HasOrganizationName() bool {
	if o != nil && o.OrganizationName.IsSet() {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given NullableString and assigns it to the OrganizationName field.
func (o *InvoiceListDto) SetOrganizationName(v string) {
	o.OrganizationName.Set(&v)
}
// SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil
func (o *InvoiceListDto) SetOrganizationNameNil() {
	o.OrganizationName.Set(nil)
}

// UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
func (o *InvoiceListDto) UnsetOrganizationName() {
	o.OrganizationName.Unset()
}

// GetInvoiceNumber returns the InvoiceNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceListDto) GetInvoiceNumber() string {
	if o == nil || IsNil(o.InvoiceNumber.Get()) {
		var ret string
		return ret
	}
	return *o.InvoiceNumber.Get()
}

// GetInvoiceNumberOk returns a tuple with the InvoiceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceListDto) GetInvoiceNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceNumber.Get(), o.InvoiceNumber.IsSet()
}

// HasInvoiceNumber returns a boolean if a field has been set.
func (o *InvoiceListDto) HasInvoiceNumber() bool {
	if o != nil && o.InvoiceNumber.IsSet() {
		return true
	}

	return false
}

// SetInvoiceNumber gets a reference to the given NullableString and assigns it to the InvoiceNumber field.
func (o *InvoiceListDto) SetInvoiceNumber(v string) {
	o.InvoiceNumber.Set(&v)
}
// SetInvoiceNumberNil sets the value for InvoiceNumber to be an explicit nil
func (o *InvoiceListDto) SetInvoiceNumberNil() {
	o.InvoiceNumber.Set(nil)
}

// UnsetInvoiceNumber ensures that no value is present for InvoiceNumber, not even an explicit nil
func (o *InvoiceListDto) UnsetInvoiceNumber() {
	o.InvoiceNumber.Unset()
}

// GetOrganizationSubscriptionId returns the OrganizationSubscriptionId field value if set, zero value otherwise.
func (o *InvoiceListDto) GetOrganizationSubscriptionId() int32 {
	if o == nil || IsNil(o.OrganizationSubscriptionId) {
		var ret int32
		return ret
	}
	return *o.OrganizationSubscriptionId
}

// GetOrganizationSubscriptionIdOk returns a tuple with the OrganizationSubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceListDto) GetOrganizationSubscriptionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrganizationSubscriptionId) {
		return nil, false
	}
	return o.OrganizationSubscriptionId, true
}

// HasOrganizationSubscriptionId returns a boolean if a field has been set.
func (o *InvoiceListDto) HasOrganizationSubscriptionId() bool {
	if o != nil && !IsNil(o.OrganizationSubscriptionId) {
		return true
	}

	return false
}

// SetOrganizationSubscriptionId gets a reference to the given int32 and assigns it to the OrganizationSubscriptionId field.
func (o *InvoiceListDto) SetOrganizationSubscriptionId(v int32) {
	o.OrganizationSubscriptionId = &v
}

func (o InvoiceListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.StartDate.IsSet() {
		toSerialize["startDate"] = o.StartDate.Get()
	}
	if o.EndDate.IsSet() {
		toSerialize["endDate"] = o.EndDate.Get()
	}
	if !IsNil(o.RequiredPaymentAction) {
		toSerialize["requiredPaymentAction"] = o.RequiredPaymentAction
	}
	if !IsNil(o.IsPaid) {
		toSerialize["isPaid"] = o.IsPaid
	}
	if o.InvoiceId.IsSet() {
		toSerialize["invoiceId"] = o.InvoiceId.Get()
	}
	if o.SubscriptionType.IsSet() {
		toSerialize["subscriptionType"] = o.SubscriptionType.Get()
	}
	if o.SubscriptionName.IsSet() {
		toSerialize["subscriptionName"] = o.SubscriptionName.Get()
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if o.OrganizationName.IsSet() {
		toSerialize["organizationName"] = o.OrganizationName.Get()
	}
	if o.InvoiceNumber.IsSet() {
		toSerialize["invoiceNumber"] = o.InvoiceNumber.Get()
	}
	if !IsNil(o.OrganizationSubscriptionId) {
		toSerialize["organizationSubscriptionId"] = o.OrganizationSubscriptionId
	}
	return toSerialize, nil
}

type NullableInvoiceListDto struct {
	value *InvoiceListDto
	isSet bool
}

func (v NullableInvoiceListDto) Get() *InvoiceListDto {
	return v.value
}

func (v *NullableInvoiceListDto) Set(val *InvoiceListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceListDto(val *InvoiceListDto) *NullableInvoiceListDto {
	return &NullableInvoiceListDto{value: val, isSet: true}
}

func (v NullableInvoiceListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


