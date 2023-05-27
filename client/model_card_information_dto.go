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

// checks if the CardInformationDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardInformationDto{}

// CardInformationDto struct for CardInformationDto
type CardInformationDto struct {
	ExpirationMonth NullableString `json:"expirationMonth,omitempty"`
	ExpirationYear NullableString `json:"expirationYear,omitempty"`
	Last4 NullableString `json:"last4,omitempty"`
	Brand NullableString `json:"brand,omitempty"`
	HolderName NullableString `json:"holderName,omitempty"`
	Balance *int64 `json:"balance,omitempty"`
}

// NewCardInformationDto instantiates a new CardInformationDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardInformationDto() *CardInformationDto {
	this := CardInformationDto{}
	return &this
}

// NewCardInformationDtoWithDefaults instantiates a new CardInformationDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardInformationDtoWithDefaults() *CardInformationDto {
	this := CardInformationDto{}
	return &this
}

// GetExpirationMonth returns the ExpirationMonth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CardInformationDto) GetExpirationMonth() string {
	if o == nil || IsNil(o.ExpirationMonth.Get()) {
		var ret string
		return ret
	}
	return *o.ExpirationMonth.Get()
}

// GetExpirationMonthOk returns a tuple with the ExpirationMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardInformationDto) GetExpirationMonthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpirationMonth.Get(), o.ExpirationMonth.IsSet()
}

// HasExpirationMonth returns a boolean if a field has been set.
func (o *CardInformationDto) HasExpirationMonth() bool {
	if o != nil && o.ExpirationMonth.IsSet() {
		return true
	}

	return false
}

// SetExpirationMonth gets a reference to the given NullableString and assigns it to the ExpirationMonth field.
func (o *CardInformationDto) SetExpirationMonth(v string) {
	o.ExpirationMonth.Set(&v)
}
// SetExpirationMonthNil sets the value for ExpirationMonth to be an explicit nil
func (o *CardInformationDto) SetExpirationMonthNil() {
	o.ExpirationMonth.Set(nil)
}

// UnsetExpirationMonth ensures that no value is present for ExpirationMonth, not even an explicit nil
func (o *CardInformationDto) UnsetExpirationMonth() {
	o.ExpirationMonth.Unset()
}

// GetExpirationYear returns the ExpirationYear field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CardInformationDto) GetExpirationYear() string {
	if o == nil || IsNil(o.ExpirationYear.Get()) {
		var ret string
		return ret
	}
	return *o.ExpirationYear.Get()
}

// GetExpirationYearOk returns a tuple with the ExpirationYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardInformationDto) GetExpirationYearOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpirationYear.Get(), o.ExpirationYear.IsSet()
}

// HasExpirationYear returns a boolean if a field has been set.
func (o *CardInformationDto) HasExpirationYear() bool {
	if o != nil && o.ExpirationYear.IsSet() {
		return true
	}

	return false
}

// SetExpirationYear gets a reference to the given NullableString and assigns it to the ExpirationYear field.
func (o *CardInformationDto) SetExpirationYear(v string) {
	o.ExpirationYear.Set(&v)
}
// SetExpirationYearNil sets the value for ExpirationYear to be an explicit nil
func (o *CardInformationDto) SetExpirationYearNil() {
	o.ExpirationYear.Set(nil)
}

// UnsetExpirationYear ensures that no value is present for ExpirationYear, not even an explicit nil
func (o *CardInformationDto) UnsetExpirationYear() {
	o.ExpirationYear.Unset()
}

// GetLast4 returns the Last4 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CardInformationDto) GetLast4() string {
	if o == nil || IsNil(o.Last4.Get()) {
		var ret string
		return ret
	}
	return *o.Last4.Get()
}

// GetLast4Ok returns a tuple with the Last4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardInformationDto) GetLast4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Last4.Get(), o.Last4.IsSet()
}

// HasLast4 returns a boolean if a field has been set.
func (o *CardInformationDto) HasLast4() bool {
	if o != nil && o.Last4.IsSet() {
		return true
	}

	return false
}

// SetLast4 gets a reference to the given NullableString and assigns it to the Last4 field.
func (o *CardInformationDto) SetLast4(v string) {
	o.Last4.Set(&v)
}
// SetLast4Nil sets the value for Last4 to be an explicit nil
func (o *CardInformationDto) SetLast4Nil() {
	o.Last4.Set(nil)
}

// UnsetLast4 ensures that no value is present for Last4, not even an explicit nil
func (o *CardInformationDto) UnsetLast4() {
	o.Last4.Unset()
}

// GetBrand returns the Brand field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CardInformationDto) GetBrand() string {
	if o == nil || IsNil(o.Brand.Get()) {
		var ret string
		return ret
	}
	return *o.Brand.Get()
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardInformationDto) GetBrandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Brand.Get(), o.Brand.IsSet()
}

// HasBrand returns a boolean if a field has been set.
func (o *CardInformationDto) HasBrand() bool {
	if o != nil && o.Brand.IsSet() {
		return true
	}

	return false
}

// SetBrand gets a reference to the given NullableString and assigns it to the Brand field.
func (o *CardInformationDto) SetBrand(v string) {
	o.Brand.Set(&v)
}
// SetBrandNil sets the value for Brand to be an explicit nil
func (o *CardInformationDto) SetBrandNil() {
	o.Brand.Set(nil)
}

// UnsetBrand ensures that no value is present for Brand, not even an explicit nil
func (o *CardInformationDto) UnsetBrand() {
	o.Brand.Unset()
}

// GetHolderName returns the HolderName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CardInformationDto) GetHolderName() string {
	if o == nil || IsNil(o.HolderName.Get()) {
		var ret string
		return ret
	}
	return *o.HolderName.Get()
}

// GetHolderNameOk returns a tuple with the HolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardInformationDto) GetHolderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HolderName.Get(), o.HolderName.IsSet()
}

// HasHolderName returns a boolean if a field has been set.
func (o *CardInformationDto) HasHolderName() bool {
	if o != nil && o.HolderName.IsSet() {
		return true
	}

	return false
}

// SetHolderName gets a reference to the given NullableString and assigns it to the HolderName field.
func (o *CardInformationDto) SetHolderName(v string) {
	o.HolderName.Set(&v)
}
// SetHolderNameNil sets the value for HolderName to be an explicit nil
func (o *CardInformationDto) SetHolderNameNil() {
	o.HolderName.Set(nil)
}

// UnsetHolderName ensures that no value is present for HolderName, not even an explicit nil
func (o *CardInformationDto) UnsetHolderName() {
	o.HolderName.Unset()
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *CardInformationDto) GetBalance() int64 {
	if o == nil || IsNil(o.Balance) {
		var ret int64
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInformationDto) GetBalanceOk() (*int64, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *CardInformationDto) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given int64 and assigns it to the Balance field.
func (o *CardInformationDto) SetBalance(v int64) {
	o.Balance = &v
}

func (o CardInformationDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardInformationDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpirationMonth.IsSet() {
		toSerialize["expirationMonth"] = o.ExpirationMonth.Get()
	}
	if o.ExpirationYear.IsSet() {
		toSerialize["expirationYear"] = o.ExpirationYear.Get()
	}
	if o.Last4.IsSet() {
		toSerialize["last4"] = o.Last4.Get()
	}
	if o.Brand.IsSet() {
		toSerialize["brand"] = o.Brand.Get()
	}
	if o.HolderName.IsSet() {
		toSerialize["holderName"] = o.HolderName.Get()
	}
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	return toSerialize, nil
}

type NullableCardInformationDto struct {
	value *CardInformationDto
	isSet bool
}

func (v NullableCardInformationDto) Get() *CardInformationDto {
	return v.value
}

func (v *NullableCardInformationDto) Set(val *CardInformationDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCardInformationDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCardInformationDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardInformationDto(val *CardInformationDto) *NullableCardInformationDto {
	return &NullableCardInformationDto{value: val, isSet: true}
}

func (v NullableCardInformationDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardInformationDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


