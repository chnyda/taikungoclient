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

// checks if the AwsExtendedImagesListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsExtendedImagesListDto{}

// AwsExtendedImagesListDto struct for AwsExtendedImagesListDto
type AwsExtendedImagesListDto struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	PlatformDetails *string `json:"platformDetails,omitempty"`
	OwnerId *string `json:"ownerId,omitempty"`
	OwnerAlias *string `json:"ownerAlias,omitempty"`
	Logo *string `json:"logo,omitempty"`
}

// NewAwsExtendedImagesListDto instantiates a new AwsExtendedImagesListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsExtendedImagesListDto() *AwsExtendedImagesListDto {
	this := AwsExtendedImagesListDto{}
	return &this
}

// NewAwsExtendedImagesListDtoWithDefaults instantiates a new AwsExtendedImagesListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsExtendedImagesListDtoWithDefaults() *AwsExtendedImagesListDto {
	this := AwsExtendedImagesListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AwsExtendedImagesListDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsExtendedImagesListDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AwsExtendedImagesListDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AwsExtendedImagesListDto) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AwsExtendedImagesListDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsExtendedImagesListDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AwsExtendedImagesListDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AwsExtendedImagesListDto) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AwsExtendedImagesListDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsExtendedImagesListDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AwsExtendedImagesListDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AwsExtendedImagesListDto) SetDescription(v string) {
	o.Description = &v
}

// GetPlatformDetails returns the PlatformDetails field value if set, zero value otherwise.
func (o *AwsExtendedImagesListDto) GetPlatformDetails() string {
	if o == nil || IsNil(o.PlatformDetails) {
		var ret string
		return ret
	}
	return *o.PlatformDetails
}

// GetPlatformDetailsOk returns a tuple with the PlatformDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsExtendedImagesListDto) GetPlatformDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.PlatformDetails) {
		return nil, false
	}
	return o.PlatformDetails, true
}

// HasPlatformDetails returns a boolean if a field has been set.
func (o *AwsExtendedImagesListDto) HasPlatformDetails() bool {
	if o != nil && !IsNil(o.PlatformDetails) {
		return true
	}

	return false
}

// SetPlatformDetails gets a reference to the given string and assigns it to the PlatformDetails field.
func (o *AwsExtendedImagesListDto) SetPlatformDetails(v string) {
	o.PlatformDetails = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *AwsExtendedImagesListDto) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsExtendedImagesListDto) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *AwsExtendedImagesListDto) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *AwsExtendedImagesListDto) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetOwnerAlias returns the OwnerAlias field value if set, zero value otherwise.
func (o *AwsExtendedImagesListDto) GetOwnerAlias() string {
	if o == nil || IsNil(o.OwnerAlias) {
		var ret string
		return ret
	}
	return *o.OwnerAlias
}

// GetOwnerAliasOk returns a tuple with the OwnerAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsExtendedImagesListDto) GetOwnerAliasOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerAlias) {
		return nil, false
	}
	return o.OwnerAlias, true
}

// HasOwnerAlias returns a boolean if a field has been set.
func (o *AwsExtendedImagesListDto) HasOwnerAlias() bool {
	if o != nil && !IsNil(o.OwnerAlias) {
		return true
	}

	return false
}

// SetOwnerAlias gets a reference to the given string and assigns it to the OwnerAlias field.
func (o *AwsExtendedImagesListDto) SetOwnerAlias(v string) {
	o.OwnerAlias = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *AwsExtendedImagesListDto) GetLogo() string {
	if o == nil || IsNil(o.Logo) {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsExtendedImagesListDto) GetLogoOk() (*string, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *AwsExtendedImagesListDto) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *AwsExtendedImagesListDto) SetLogo(v string) {
	o.Logo = &v
}

func (o AwsExtendedImagesListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsExtendedImagesListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.PlatformDetails) {
		toSerialize["platformDetails"] = o.PlatformDetails
	}
	if !IsNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}
	if !IsNil(o.OwnerAlias) {
		toSerialize["ownerAlias"] = o.OwnerAlias
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	return toSerialize, nil
}

type NullableAwsExtendedImagesListDto struct {
	value *AwsExtendedImagesListDto
	isSet bool
}

func (v NullableAwsExtendedImagesListDto) Get() *AwsExtendedImagesListDto {
	return v.value
}

func (v *NullableAwsExtendedImagesListDto) Set(val *AwsExtendedImagesListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsExtendedImagesListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsExtendedImagesListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsExtendedImagesListDto(val *AwsExtendedImagesListDto) *NullableAwsExtendedImagesListDto {
	return &NullableAwsExtendedImagesListDto{value: val, isSet: true}
}

func (v NullableAwsExtendedImagesListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsExtendedImagesListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


