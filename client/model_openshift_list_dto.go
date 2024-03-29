/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikuncore

import (
	"encoding/json"
)

// checks if the OpenshiftListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenshiftListDto{}

// OpenshiftListDto struct for OpenshiftListDto
type OpenshiftListDto struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	BaseDomain NullableString `json:"baseDomain,omitempty"`
	StorageClass NullableString `json:"storageClass,omitempty"`
	OrganizationId *int32 `json:"organizationId,omitempty"`
	OrganizationName NullableString `json:"organizationName,omitempty"`
}

// NewOpenshiftListDto instantiates a new OpenshiftListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenshiftListDto() *OpenshiftListDto {
	this := OpenshiftListDto{}
	return &this
}

// NewOpenshiftListDtoWithDefaults instantiates a new OpenshiftListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenshiftListDtoWithDefaults() *OpenshiftListDto {
	this := OpenshiftListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OpenshiftListDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenshiftListDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OpenshiftListDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *OpenshiftListDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenshiftListDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenshiftListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *OpenshiftListDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *OpenshiftListDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *OpenshiftListDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *OpenshiftListDto) UnsetName() {
	o.Name.Unset()
}

// GetBaseDomain returns the BaseDomain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenshiftListDto) GetBaseDomain() string {
	if o == nil || IsNil(o.BaseDomain.Get()) {
		var ret string
		return ret
	}
	return *o.BaseDomain.Get()
}

// GetBaseDomainOk returns a tuple with the BaseDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenshiftListDto) GetBaseDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BaseDomain.Get(), o.BaseDomain.IsSet()
}

// HasBaseDomain returns a boolean if a field has been set.
func (o *OpenshiftListDto) HasBaseDomain() bool {
	if o != nil && o.BaseDomain.IsSet() {
		return true
	}

	return false
}

// SetBaseDomain gets a reference to the given NullableString and assigns it to the BaseDomain field.
func (o *OpenshiftListDto) SetBaseDomain(v string) {
	o.BaseDomain.Set(&v)
}
// SetBaseDomainNil sets the value for BaseDomain to be an explicit nil
func (o *OpenshiftListDto) SetBaseDomainNil() {
	o.BaseDomain.Set(nil)
}

// UnsetBaseDomain ensures that no value is present for BaseDomain, not even an explicit nil
func (o *OpenshiftListDto) UnsetBaseDomain() {
	o.BaseDomain.Unset()
}

// GetStorageClass returns the StorageClass field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenshiftListDto) GetStorageClass() string {
	if o == nil || IsNil(o.StorageClass.Get()) {
		var ret string
		return ret
	}
	return *o.StorageClass.Get()
}

// GetStorageClassOk returns a tuple with the StorageClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenshiftListDto) GetStorageClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageClass.Get(), o.StorageClass.IsSet()
}

// HasStorageClass returns a boolean if a field has been set.
func (o *OpenshiftListDto) HasStorageClass() bool {
	if o != nil && o.StorageClass.IsSet() {
		return true
	}

	return false
}

// SetStorageClass gets a reference to the given NullableString and assigns it to the StorageClass field.
func (o *OpenshiftListDto) SetStorageClass(v string) {
	o.StorageClass.Set(&v)
}
// SetStorageClassNil sets the value for StorageClass to be an explicit nil
func (o *OpenshiftListDto) SetStorageClassNil() {
	o.StorageClass.Set(nil)
}

// UnsetStorageClass ensures that no value is present for StorageClass, not even an explicit nil
func (o *OpenshiftListDto) UnsetStorageClass() {
	o.StorageClass.Unset()
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *OpenshiftListDto) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId) {
		var ret int32
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenshiftListDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *OpenshiftListDto) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int32 and assigns it to the OrganizationId field.
func (o *OpenshiftListDto) SetOrganizationId(v int32) {
	o.OrganizationId = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenshiftListDto) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationName.Get()
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenshiftListDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationName.Get(), o.OrganizationName.IsSet()
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *OpenshiftListDto) HasOrganizationName() bool {
	if o != nil && o.OrganizationName.IsSet() {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given NullableString and assigns it to the OrganizationName field.
func (o *OpenshiftListDto) SetOrganizationName(v string) {
	o.OrganizationName.Set(&v)
}
// SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil
func (o *OpenshiftListDto) SetOrganizationNameNil() {
	o.OrganizationName.Set(nil)
}

// UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
func (o *OpenshiftListDto) UnsetOrganizationName() {
	o.OrganizationName.Unset()
}

func (o OpenshiftListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenshiftListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.BaseDomain.IsSet() {
		toSerialize["baseDomain"] = o.BaseDomain.Get()
	}
	if o.StorageClass.IsSet() {
		toSerialize["storageClass"] = o.StorageClass.Get()
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if o.OrganizationName.IsSet() {
		toSerialize["organizationName"] = o.OrganizationName.Get()
	}
	return toSerialize, nil
}

type NullableOpenshiftListDto struct {
	value *OpenshiftListDto
	isSet bool
}

func (v NullableOpenshiftListDto) Get() *OpenshiftListDto {
	return v.value
}

func (v *NullableOpenshiftListDto) Set(val *OpenshiftListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenshiftListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenshiftListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenshiftListDto(val *OpenshiftListDto) *NullableOpenshiftListDto {
	return &NullableOpenshiftListDto{value: val, isSet: true}
}

func (v NullableOpenshiftListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenshiftListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


