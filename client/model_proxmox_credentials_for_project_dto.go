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

// checks if the ProxmoxCredentialsForProjectDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProxmoxCredentialsForProjectDto{}

// ProxmoxCredentialsForProjectDto struct for ProxmoxCredentialsForProjectDto
type ProxmoxCredentialsForProjectDto struct {
	TokenId *string `json:"tokenId,omitempty"`
	TokenSecret *string `json:"tokenSecret,omitempty"`
	Url *string `json:"url,omitempty"`
	Password *string `json:"password,omitempty"`
	Storage *string `json:"storage,omitempty"`
	VmTemplateName *string `json:"vmTemplateName,omitempty"`
	ProxmoxNetworks []ProxmoxNetworkListDto `json:"proxmoxNetworks,omitempty"`
}

// NewProxmoxCredentialsForProjectDto instantiates a new ProxmoxCredentialsForProjectDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxmoxCredentialsForProjectDto() *ProxmoxCredentialsForProjectDto {
	this := ProxmoxCredentialsForProjectDto{}
	return &this
}

// NewProxmoxCredentialsForProjectDtoWithDefaults instantiates a new ProxmoxCredentialsForProjectDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxmoxCredentialsForProjectDtoWithDefaults() *ProxmoxCredentialsForProjectDto {
	this := ProxmoxCredentialsForProjectDto{}
	return &this
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *ProxmoxCredentialsForProjectDto) GetTokenId() string {
	if o == nil || IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxmoxCredentialsForProjectDto) GetTokenIdOk() (*string, bool) {
	if o == nil || IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *ProxmoxCredentialsForProjectDto) HasTokenId() bool {
	if o != nil && !IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *ProxmoxCredentialsForProjectDto) SetTokenId(v string) {
	o.TokenId = &v
}

// GetTokenSecret returns the TokenSecret field value if set, zero value otherwise.
func (o *ProxmoxCredentialsForProjectDto) GetTokenSecret() string {
	if o == nil || IsNil(o.TokenSecret) {
		var ret string
		return ret
	}
	return *o.TokenSecret
}

// GetTokenSecretOk returns a tuple with the TokenSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxmoxCredentialsForProjectDto) GetTokenSecretOk() (*string, bool) {
	if o == nil || IsNil(o.TokenSecret) {
		return nil, false
	}
	return o.TokenSecret, true
}

// HasTokenSecret returns a boolean if a field has been set.
func (o *ProxmoxCredentialsForProjectDto) HasTokenSecret() bool {
	if o != nil && !IsNil(o.TokenSecret) {
		return true
	}

	return false
}

// SetTokenSecret gets a reference to the given string and assigns it to the TokenSecret field.
func (o *ProxmoxCredentialsForProjectDto) SetTokenSecret(v string) {
	o.TokenSecret = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ProxmoxCredentialsForProjectDto) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxmoxCredentialsForProjectDto) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ProxmoxCredentialsForProjectDto) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ProxmoxCredentialsForProjectDto) SetUrl(v string) {
	o.Url = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ProxmoxCredentialsForProjectDto) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxmoxCredentialsForProjectDto) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ProxmoxCredentialsForProjectDto) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ProxmoxCredentialsForProjectDto) SetPassword(v string) {
	o.Password = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *ProxmoxCredentialsForProjectDto) GetStorage() string {
	if o == nil || IsNil(o.Storage) {
		var ret string
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxmoxCredentialsForProjectDto) GetStorageOk() (*string, bool) {
	if o == nil || IsNil(o.Storage) {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *ProxmoxCredentialsForProjectDto) HasStorage() bool {
	if o != nil && !IsNil(o.Storage) {
		return true
	}

	return false
}

// SetStorage gets a reference to the given string and assigns it to the Storage field.
func (o *ProxmoxCredentialsForProjectDto) SetStorage(v string) {
	o.Storage = &v
}

// GetVmTemplateName returns the VmTemplateName field value if set, zero value otherwise.
func (o *ProxmoxCredentialsForProjectDto) GetVmTemplateName() string {
	if o == nil || IsNil(o.VmTemplateName) {
		var ret string
		return ret
	}
	return *o.VmTemplateName
}

// GetVmTemplateNameOk returns a tuple with the VmTemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxmoxCredentialsForProjectDto) GetVmTemplateNameOk() (*string, bool) {
	if o == nil || IsNil(o.VmTemplateName) {
		return nil, false
	}
	return o.VmTemplateName, true
}

// HasVmTemplateName returns a boolean if a field has been set.
func (o *ProxmoxCredentialsForProjectDto) HasVmTemplateName() bool {
	if o != nil && !IsNil(o.VmTemplateName) {
		return true
	}

	return false
}

// SetVmTemplateName gets a reference to the given string and assigns it to the VmTemplateName field.
func (o *ProxmoxCredentialsForProjectDto) SetVmTemplateName(v string) {
	o.VmTemplateName = &v
}

// GetProxmoxNetworks returns the ProxmoxNetworks field value if set, zero value otherwise.
func (o *ProxmoxCredentialsForProjectDto) GetProxmoxNetworks() []ProxmoxNetworkListDto {
	if o == nil || IsNil(o.ProxmoxNetworks) {
		var ret []ProxmoxNetworkListDto
		return ret
	}
	return o.ProxmoxNetworks
}

// GetProxmoxNetworksOk returns a tuple with the ProxmoxNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxmoxCredentialsForProjectDto) GetProxmoxNetworksOk() ([]ProxmoxNetworkListDto, bool) {
	if o == nil || IsNil(o.ProxmoxNetworks) {
		return nil, false
	}
	return o.ProxmoxNetworks, true
}

// HasProxmoxNetworks returns a boolean if a field has been set.
func (o *ProxmoxCredentialsForProjectDto) HasProxmoxNetworks() bool {
	if o != nil && !IsNil(o.ProxmoxNetworks) {
		return true
	}

	return false
}

// SetProxmoxNetworks gets a reference to the given []ProxmoxNetworkListDto and assigns it to the ProxmoxNetworks field.
func (o *ProxmoxCredentialsForProjectDto) SetProxmoxNetworks(v []ProxmoxNetworkListDto) {
	o.ProxmoxNetworks = v
}

func (o ProxmoxCredentialsForProjectDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProxmoxCredentialsForProjectDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TokenId) {
		toSerialize["tokenId"] = o.TokenId
	}
	if !IsNil(o.TokenSecret) {
		toSerialize["tokenSecret"] = o.TokenSecret
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Storage) {
		toSerialize["storage"] = o.Storage
	}
	if !IsNil(o.VmTemplateName) {
		toSerialize["vmTemplateName"] = o.VmTemplateName
	}
	if !IsNil(o.ProxmoxNetworks) {
		toSerialize["proxmoxNetworks"] = o.ProxmoxNetworks
	}
	return toSerialize, nil
}

type NullableProxmoxCredentialsForProjectDto struct {
	value *ProxmoxCredentialsForProjectDto
	isSet bool
}

func (v NullableProxmoxCredentialsForProjectDto) Get() *ProxmoxCredentialsForProjectDto {
	return v.value
}

func (v *NullableProxmoxCredentialsForProjectDto) Set(val *ProxmoxCredentialsForProjectDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProxmoxCredentialsForProjectDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProxmoxCredentialsForProjectDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxmoxCredentialsForProjectDto(val *ProxmoxCredentialsForProjectDto) *NullableProxmoxCredentialsForProjectDto {
	return &NullableProxmoxCredentialsForProjectDto{value: val, isSet: true}
}

func (v NullableProxmoxCredentialsForProjectDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxmoxCredentialsForProjectDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


