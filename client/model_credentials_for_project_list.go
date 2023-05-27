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

// checks if the CredentialsForProjectList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialsForProjectList{}

// CredentialsForProjectList struct for CredentialsForProjectList
type CredentialsForProjectList struct {
	CloudType *CloudType `json:"cloudType,omitempty"`
	CloudCredentialRevision *int32 `json:"cloudCredentialRevision,omitempty"`
	BillingEnabled *bool `json:"billingEnabled,omitempty"`
	ContinentName NullableString `json:"continentName,omitempty"`
	RequiresVPN *bool `json:"requiresVPN,omitempty"`
	Openstack *OpenstackCredentialsForProjectDto `json:"openstack,omitempty"`
	Azure *AzureCredentialsForProjectDto `json:"azure,omitempty"`
	Aws *AwsCredentialsForProjectDto `json:"aws,omitempty"`
	Google *GoogleCredentialForProjectDto `json:"google,omitempty"`
	Tanzu *TanzuCredentialsForProjectDto `json:"tanzu,omitempty"`
	Proxmox *ProxmoxCredentialsForProjectDto `json:"proxmox,omitempty"`
}

// NewCredentialsForProjectList instantiates a new CredentialsForProjectList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialsForProjectList() *CredentialsForProjectList {
	this := CredentialsForProjectList{}
	return &this
}

// NewCredentialsForProjectListWithDefaults instantiates a new CredentialsForProjectList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialsForProjectListWithDefaults() *CredentialsForProjectList {
	this := CredentialsForProjectList{}
	return &this
}

// GetCloudType returns the CloudType field value if set, zero value otherwise.
func (o *CredentialsForProjectList) GetCloudType() CloudType {
	if o == nil || IsNil(o.CloudType) {
		var ret CloudType
		return ret
	}
	return *o.CloudType
}

// GetCloudTypeOk returns a tuple with the CloudType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsForProjectList) GetCloudTypeOk() (*CloudType, bool) {
	if o == nil || IsNil(o.CloudType) {
		return nil, false
	}
	return o.CloudType, true
}

// HasCloudType returns a boolean if a field has been set.
func (o *CredentialsForProjectList) HasCloudType() bool {
	if o != nil && !IsNil(o.CloudType) {
		return true
	}

	return false
}

// SetCloudType gets a reference to the given CloudType and assigns it to the CloudType field.
func (o *CredentialsForProjectList) SetCloudType(v CloudType) {
	o.CloudType = &v
}

// GetCloudCredentialRevision returns the CloudCredentialRevision field value if set, zero value otherwise.
func (o *CredentialsForProjectList) GetCloudCredentialRevision() int32 {
	if o == nil || IsNil(o.CloudCredentialRevision) {
		var ret int32
		return ret
	}
	return *o.CloudCredentialRevision
}

// GetCloudCredentialRevisionOk returns a tuple with the CloudCredentialRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsForProjectList) GetCloudCredentialRevisionOk() (*int32, bool) {
	if o == nil || IsNil(o.CloudCredentialRevision) {
		return nil, false
	}
	return o.CloudCredentialRevision, true
}

// HasCloudCredentialRevision returns a boolean if a field has been set.
func (o *CredentialsForProjectList) HasCloudCredentialRevision() bool {
	if o != nil && !IsNil(o.CloudCredentialRevision) {
		return true
	}

	return false
}

// SetCloudCredentialRevision gets a reference to the given int32 and assigns it to the CloudCredentialRevision field.
func (o *CredentialsForProjectList) SetCloudCredentialRevision(v int32) {
	o.CloudCredentialRevision = &v
}

// GetBillingEnabled returns the BillingEnabled field value if set, zero value otherwise.
func (o *CredentialsForProjectList) GetBillingEnabled() bool {
	if o == nil || IsNil(o.BillingEnabled) {
		var ret bool
		return ret
	}
	return *o.BillingEnabled
}

// GetBillingEnabledOk returns a tuple with the BillingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsForProjectList) GetBillingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BillingEnabled) {
		return nil, false
	}
	return o.BillingEnabled, true
}

// HasBillingEnabled returns a boolean if a field has been set.
func (o *CredentialsForProjectList) HasBillingEnabled() bool {
	if o != nil && !IsNil(o.BillingEnabled) {
		return true
	}

	return false
}

// SetBillingEnabled gets a reference to the given bool and assigns it to the BillingEnabled field.
func (o *CredentialsForProjectList) SetBillingEnabled(v bool) {
	o.BillingEnabled = &v
}

// GetContinentName returns the ContinentName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredentialsForProjectList) GetContinentName() string {
	if o == nil || IsNil(o.ContinentName.Get()) {
		var ret string
		return ret
	}
	return *o.ContinentName.Get()
}

// GetContinentNameOk returns a tuple with the ContinentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialsForProjectList) GetContinentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContinentName.Get(), o.ContinentName.IsSet()
}

// HasContinentName returns a boolean if a field has been set.
func (o *CredentialsForProjectList) HasContinentName() bool {
	if o != nil && o.ContinentName.IsSet() {
		return true
	}

	return false
}

// SetContinentName gets a reference to the given NullableString and assigns it to the ContinentName field.
func (o *CredentialsForProjectList) SetContinentName(v string) {
	o.ContinentName.Set(&v)
}
// SetContinentNameNil sets the value for ContinentName to be an explicit nil
func (o *CredentialsForProjectList) SetContinentNameNil() {
	o.ContinentName.Set(nil)
}

// UnsetContinentName ensures that no value is present for ContinentName, not even an explicit nil
func (o *CredentialsForProjectList) UnsetContinentName() {
	o.ContinentName.Unset()
}

// GetRequiresVPN returns the RequiresVPN field value if set, zero value otherwise.
func (o *CredentialsForProjectList) GetRequiresVPN() bool {
	if o == nil || IsNil(o.RequiresVPN) {
		var ret bool
		return ret
	}
	return *o.RequiresVPN
}

// GetRequiresVPNOk returns a tuple with the RequiresVPN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsForProjectList) GetRequiresVPNOk() (*bool, bool) {
	if o == nil || IsNil(o.RequiresVPN) {
		return nil, false
	}
	return o.RequiresVPN, true
}

// HasRequiresVPN returns a boolean if a field has been set.
func (o *CredentialsForProjectList) HasRequiresVPN() bool {
	if o != nil && !IsNil(o.RequiresVPN) {
		return true
	}

	return false
}

// SetRequiresVPN gets a reference to the given bool and assigns it to the RequiresVPN field.
func (o *CredentialsForProjectList) SetRequiresVPN(v bool) {
	o.RequiresVPN = &v
}

// GetOpenstack returns the Openstack field value if set, zero value otherwise.
func (o *CredentialsForProjectList) GetOpenstack() OpenstackCredentialsForProjectDto {
	if o == nil || IsNil(o.Openstack) {
		var ret OpenstackCredentialsForProjectDto
		return ret
	}
	return *o.Openstack
}

// GetOpenstackOk returns a tuple with the Openstack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsForProjectList) GetOpenstackOk() (*OpenstackCredentialsForProjectDto, bool) {
	if o == nil || IsNil(o.Openstack) {
		return nil, false
	}
	return o.Openstack, true
}

// HasOpenstack returns a boolean if a field has been set.
func (o *CredentialsForProjectList) HasOpenstack() bool {
	if o != nil && !IsNil(o.Openstack) {
		return true
	}

	return false
}

// SetOpenstack gets a reference to the given OpenstackCredentialsForProjectDto and assigns it to the Openstack field.
func (o *CredentialsForProjectList) SetOpenstack(v OpenstackCredentialsForProjectDto) {
	o.Openstack = &v
}

// GetAzure returns the Azure field value if set, zero value otherwise.
func (o *CredentialsForProjectList) GetAzure() AzureCredentialsForProjectDto {
	if o == nil || IsNil(o.Azure) {
		var ret AzureCredentialsForProjectDto
		return ret
	}
	return *o.Azure
}

// GetAzureOk returns a tuple with the Azure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsForProjectList) GetAzureOk() (*AzureCredentialsForProjectDto, bool) {
	if o == nil || IsNil(o.Azure) {
		return nil, false
	}
	return o.Azure, true
}

// HasAzure returns a boolean if a field has been set.
func (o *CredentialsForProjectList) HasAzure() bool {
	if o != nil && !IsNil(o.Azure) {
		return true
	}

	return false
}

// SetAzure gets a reference to the given AzureCredentialsForProjectDto and assigns it to the Azure field.
func (o *CredentialsForProjectList) SetAzure(v AzureCredentialsForProjectDto) {
	o.Azure = &v
}

// GetAws returns the Aws field value if set, zero value otherwise.
func (o *CredentialsForProjectList) GetAws() AwsCredentialsForProjectDto {
	if o == nil || IsNil(o.Aws) {
		var ret AwsCredentialsForProjectDto
		return ret
	}
	return *o.Aws
}

// GetAwsOk returns a tuple with the Aws field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsForProjectList) GetAwsOk() (*AwsCredentialsForProjectDto, bool) {
	if o == nil || IsNil(o.Aws) {
		return nil, false
	}
	return o.Aws, true
}

// HasAws returns a boolean if a field has been set.
func (o *CredentialsForProjectList) HasAws() bool {
	if o != nil && !IsNil(o.Aws) {
		return true
	}

	return false
}

// SetAws gets a reference to the given AwsCredentialsForProjectDto and assigns it to the Aws field.
func (o *CredentialsForProjectList) SetAws(v AwsCredentialsForProjectDto) {
	o.Aws = &v
}

// GetGoogle returns the Google field value if set, zero value otherwise.
func (o *CredentialsForProjectList) GetGoogle() GoogleCredentialForProjectDto {
	if o == nil || IsNil(o.Google) {
		var ret GoogleCredentialForProjectDto
		return ret
	}
	return *o.Google
}

// GetGoogleOk returns a tuple with the Google field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsForProjectList) GetGoogleOk() (*GoogleCredentialForProjectDto, bool) {
	if o == nil || IsNil(o.Google) {
		return nil, false
	}
	return o.Google, true
}

// HasGoogle returns a boolean if a field has been set.
func (o *CredentialsForProjectList) HasGoogle() bool {
	if o != nil && !IsNil(o.Google) {
		return true
	}

	return false
}

// SetGoogle gets a reference to the given GoogleCredentialForProjectDto and assigns it to the Google field.
func (o *CredentialsForProjectList) SetGoogle(v GoogleCredentialForProjectDto) {
	o.Google = &v
}

// GetTanzu returns the Tanzu field value if set, zero value otherwise.
func (o *CredentialsForProjectList) GetTanzu() TanzuCredentialsForProjectDto {
	if o == nil || IsNil(o.Tanzu) {
		var ret TanzuCredentialsForProjectDto
		return ret
	}
	return *o.Tanzu
}

// GetTanzuOk returns a tuple with the Tanzu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsForProjectList) GetTanzuOk() (*TanzuCredentialsForProjectDto, bool) {
	if o == nil || IsNil(o.Tanzu) {
		return nil, false
	}
	return o.Tanzu, true
}

// HasTanzu returns a boolean if a field has been set.
func (o *CredentialsForProjectList) HasTanzu() bool {
	if o != nil && !IsNil(o.Tanzu) {
		return true
	}

	return false
}

// SetTanzu gets a reference to the given TanzuCredentialsForProjectDto and assigns it to the Tanzu field.
func (o *CredentialsForProjectList) SetTanzu(v TanzuCredentialsForProjectDto) {
	o.Tanzu = &v
}

// GetProxmox returns the Proxmox field value if set, zero value otherwise.
func (o *CredentialsForProjectList) GetProxmox() ProxmoxCredentialsForProjectDto {
	if o == nil || IsNil(o.Proxmox) {
		var ret ProxmoxCredentialsForProjectDto
		return ret
	}
	return *o.Proxmox
}

// GetProxmoxOk returns a tuple with the Proxmox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsForProjectList) GetProxmoxOk() (*ProxmoxCredentialsForProjectDto, bool) {
	if o == nil || IsNil(o.Proxmox) {
		return nil, false
	}
	return o.Proxmox, true
}

// HasProxmox returns a boolean if a field has been set.
func (o *CredentialsForProjectList) HasProxmox() bool {
	if o != nil && !IsNil(o.Proxmox) {
		return true
	}

	return false
}

// SetProxmox gets a reference to the given ProxmoxCredentialsForProjectDto and assigns it to the Proxmox field.
func (o *CredentialsForProjectList) SetProxmox(v ProxmoxCredentialsForProjectDto) {
	o.Proxmox = &v
}

func (o CredentialsForProjectList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialsForProjectList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudType) {
		toSerialize["cloudType"] = o.CloudType
	}
	if !IsNil(o.CloudCredentialRevision) {
		toSerialize["cloudCredentialRevision"] = o.CloudCredentialRevision
	}
	if !IsNil(o.BillingEnabled) {
		toSerialize["billingEnabled"] = o.BillingEnabled
	}
	if o.ContinentName.IsSet() {
		toSerialize["continentName"] = o.ContinentName.Get()
	}
	if !IsNil(o.RequiresVPN) {
		toSerialize["requiresVPN"] = o.RequiresVPN
	}
	if !IsNil(o.Openstack) {
		toSerialize["openstack"] = o.Openstack
	}
	if !IsNil(o.Azure) {
		toSerialize["azure"] = o.Azure
	}
	if !IsNil(o.Aws) {
		toSerialize["aws"] = o.Aws
	}
	if !IsNil(o.Google) {
		toSerialize["google"] = o.Google
	}
	if !IsNil(o.Tanzu) {
		toSerialize["tanzu"] = o.Tanzu
	}
	if !IsNil(o.Proxmox) {
		toSerialize["proxmox"] = o.Proxmox
	}
	return toSerialize, nil
}

type NullableCredentialsForProjectList struct {
	value *CredentialsForProjectList
	isSet bool
}

func (v NullableCredentialsForProjectList) Get() *CredentialsForProjectList {
	return v.value
}

func (v *NullableCredentialsForProjectList) Set(val *CredentialsForProjectList) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialsForProjectList) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialsForProjectList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialsForProjectList(val *CredentialsForProjectList) *NullableCredentialsForProjectList {
	return &NullableCredentialsForProjectList{value: val, isSet: true}
}

func (v NullableCredentialsForProjectList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialsForProjectList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


