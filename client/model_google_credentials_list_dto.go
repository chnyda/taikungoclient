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

// checks if the GoogleCredentialsListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GoogleCredentialsListDto{}

// GoogleCredentialsListDto struct for GoogleCredentialsListDto
type GoogleCredentialsListDto struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Projects []CommonDropdownDto `json:"projects,omitempty"`
	OrganizationId *int32 `json:"organizationId,omitempty"`
	OrganizationName *string `json:"organizationName,omitempty"`
	PartnerLogo *string `json:"partnerLogo,omitempty"`
	PartnerName *string `json:"partnerName,omitempty"`
	FolderId *string `json:"folderId,omitempty"`
	ProjectId *string `json:"projectId,omitempty"`
	BillingAccountId *string `json:"billingAccountId,omitempty"`
	Zones []string `json:"zones,omitempty"`
	AvailabilityZonesCount *int32 `json:"availabilityZonesCount,omitempty"`
	Region *string `json:"region,omitempty"`
	IsLocked *bool `json:"isLocked,omitempty"`
	IsDefault *bool `json:"isDefault,omitempty"`
	BillingAccountName *string `json:"billingAccountName,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	ContinentName *string `json:"continentName,omitempty"`
}

// NewGoogleCredentialsListDto instantiates a new GoogleCredentialsListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleCredentialsListDto() *GoogleCredentialsListDto {
	this := GoogleCredentialsListDto{}
	return &this
}

// NewGoogleCredentialsListDtoWithDefaults instantiates a new GoogleCredentialsListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCredentialsListDtoWithDefaults() *GoogleCredentialsListDto {
	this := GoogleCredentialsListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GoogleCredentialsListDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCredentialsListDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GoogleCredentialsListDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GoogleCredentialsListDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCredentialsListDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GoogleCredentialsListDto) SetName(v string) {
	o.Name = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *GoogleCredentialsListDto) GetProjects() []CommonDropdownDto {
	if o == nil || IsNil(o.Projects) {
		var ret []CommonDropdownDto
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCredentialsListDto) GetProjectsOk() ([]CommonDropdownDto, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasProjects() bool {
	if o != nil && !IsNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []CommonDropdownDto and assigns it to the Projects field.
func (o *GoogleCredentialsListDto) SetProjects(v []CommonDropdownDto) {
	o.Projects = v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *GoogleCredentialsListDto) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId) {
		var ret int32
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCredentialsListDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int32 and assigns it to the OrganizationId field.
func (o *GoogleCredentialsListDto) SetOrganizationId(v int32) {
	o.OrganizationId = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise.
func (o *GoogleCredentialsListDto) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName) {
		var ret string
		return ret
	}
	return *o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCredentialsListDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationName) {
		return nil, false
	}
	return o.OrganizationName, true
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasOrganizationName() bool {
	if o != nil && !IsNil(o.OrganizationName) {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given string and assigns it to the OrganizationName field.
func (o *GoogleCredentialsListDto) SetOrganizationName(v string) {
	o.OrganizationName = &v
}

// GetPartnerLogo returns the PartnerLogo field value if set, zero value otherwise.
func (o *GoogleCredentialsListDto) GetPartnerLogo() string {
	if o == nil || IsNil(o.PartnerLogo) {
		var ret string
		return ret
	}
	return *o.PartnerLogo
}

// GetPartnerLogoOk returns a tuple with the PartnerLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCredentialsListDto) GetPartnerLogoOk() (*string, bool) {
	if o == nil || IsNil(o.PartnerLogo) {
		return nil, false
	}
	return o.PartnerLogo, true
}

// HasPartnerLogo returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasPartnerLogo() bool {
	if o != nil && !IsNil(o.PartnerLogo) {
		return true
	}

	return false
}

// SetPartnerLogo gets a reference to the given string and assigns it to the PartnerLogo field.
func (o *GoogleCredentialsListDto) SetPartnerLogo(v string) {
	o.PartnerLogo = &v
}

// GetPartnerName returns the PartnerName field value if set, zero value otherwise.
func (o *GoogleCredentialsListDto) GetPartnerName() string {
	if o == nil || IsNil(o.PartnerName) {
		var ret string
		return ret
	}
	return *o.PartnerName
}

// GetPartnerNameOk returns a tuple with the PartnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCredentialsListDto) GetPartnerNameOk() (*string, bool) {
	if o == nil || IsNil(o.PartnerName) {
		return nil, false
	}
	return o.PartnerName, true
}

// HasPartnerName returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasPartnerName() bool {
	if o != nil && !IsNil(o.PartnerName) {
		return true
	}

	return false
}

// SetPartnerName gets a reference to the given string and assigns it to the PartnerName field.
func (o *GoogleCredentialsListDto) SetPartnerName(v string) {
	o.PartnerName = &v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *GoogleCredentialsListDto) GetFolderId() string {
	if o == nil || IsNil(o.FolderId) {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCredentialsListDto) GetFolderIdOk() (*string, bool) {
	if o == nil || IsNil(o.FolderId) {
		return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasFolderId() bool {
	if o != nil && !IsNil(o.FolderId) {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
func (o *GoogleCredentialsListDto) SetFolderId(v string) {
	o.FolderId = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *GoogleCredentialsListDto) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCredentialsListDto) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *GoogleCredentialsListDto) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetBillingAccountId returns the BillingAccountId field value if set, zero value otherwise.
func (o *GoogleCredentialsListDto) GetBillingAccountId() string {
	if o == nil || IsNil(o.BillingAccountId) {
		var ret string
		return ret
	}
	return *o.BillingAccountId
}

// GetBillingAccountIdOk returns a tuple with the BillingAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCredentialsListDto) GetBillingAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.BillingAccountId) {
		return nil, false
	}
	return o.BillingAccountId, true
}

// HasBillingAccountId returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasBillingAccountId() bool {
	if o != nil && !IsNil(o.BillingAccountId) {
		return true
	}

	return false
}

// SetBillingAccountId gets a reference to the given string and assigns it to the BillingAccountId field.
func (o *GoogleCredentialsListDto) SetBillingAccountId(v string) {
	o.BillingAccountId = &v
}

// GetZones returns the Zones field value if set, zero value otherwise.
func (o *GoogleCredentialsListDto) GetZones() []string {
	if o == nil || IsNil(o.Zones) {
		var ret []string
		return ret
	}
	return o.Zones
}

// GetZonesOk returns a tuple with the Zones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCredentialsListDto) GetZonesOk() ([]string, bool) {
	if o == nil || IsNil(o.Zones) {
		return nil, false
	}
	return o.Zones, true
}

// HasZones returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasZones() bool {
	if o != nil && !IsNil(o.Zones) {
		return true
	}

	return false
}

// SetZones gets a reference to the given []string and assigns it to the Zones field.
func (o *GoogleCredentialsListDto) SetZones(v []string) {
	o.Zones = v
}

// GetAvailabilityZonesCount returns the AvailabilityZonesCount field value if set, zero value otherwise.
func (o *GoogleCredentialsListDto) GetAvailabilityZonesCount() int32 {
	if o == nil || IsNil(o.AvailabilityZonesCount) {
		var ret int32
		return ret
	}
	return *o.AvailabilityZonesCount
}

// GetAvailabilityZonesCountOk returns a tuple with the AvailabilityZonesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCredentialsListDto) GetAvailabilityZonesCountOk() (*int32, bool) {
	if o == nil || IsNil(o.AvailabilityZonesCount) {
		return nil, false
	}
	return o.AvailabilityZonesCount, true
}

// HasAvailabilityZonesCount returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasAvailabilityZonesCount() bool {
	if o != nil && !IsNil(o.AvailabilityZonesCount) {
		return true
	}

	return false
}

// SetAvailabilityZonesCount gets a reference to the given int32 and assigns it to the AvailabilityZonesCount field.
func (o *GoogleCredentialsListDto) SetAvailabilityZonesCount(v int32) {
	o.AvailabilityZonesCount = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *GoogleCredentialsListDto) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCredentialsListDto) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *GoogleCredentialsListDto) SetRegion(v string) {
	o.Region = &v
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *GoogleCredentialsListDto) GetIsLocked() bool {
	if o == nil || IsNil(o.IsLocked) {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCredentialsListDto) GetIsLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocked) {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasIsLocked() bool {
	if o != nil && !IsNil(o.IsLocked) {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *GoogleCredentialsListDto) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *GoogleCredentialsListDto) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCredentialsListDto) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *GoogleCredentialsListDto) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetBillingAccountName returns the BillingAccountName field value if set, zero value otherwise.
func (o *GoogleCredentialsListDto) GetBillingAccountName() string {
	if o == nil || IsNil(o.BillingAccountName) {
		var ret string
		return ret
	}
	return *o.BillingAccountName
}

// GetBillingAccountNameOk returns a tuple with the BillingAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCredentialsListDto) GetBillingAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.BillingAccountName) {
		return nil, false
	}
	return o.BillingAccountName, true
}

// HasBillingAccountName returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasBillingAccountName() bool {
	if o != nil && !IsNil(o.BillingAccountName) {
		return true
	}

	return false
}

// SetBillingAccountName gets a reference to the given string and assigns it to the BillingAccountName field.
func (o *GoogleCredentialsListDto) SetBillingAccountName(v string) {
	o.BillingAccountName = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GoogleCredentialsListDto) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCredentialsListDto) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GoogleCredentialsListDto) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetContinentName returns the ContinentName field value if set, zero value otherwise.
func (o *GoogleCredentialsListDto) GetContinentName() string {
	if o == nil || IsNil(o.ContinentName) {
		var ret string
		return ret
	}
	return *o.ContinentName
}

// GetContinentNameOk returns a tuple with the ContinentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCredentialsListDto) GetContinentNameOk() (*string, bool) {
	if o == nil || IsNil(o.ContinentName) {
		return nil, false
	}
	return o.ContinentName, true
}

// HasContinentName returns a boolean if a field has been set.
func (o *GoogleCredentialsListDto) HasContinentName() bool {
	if o != nil && !IsNil(o.ContinentName) {
		return true
	}

	return false
}

// SetContinentName gets a reference to the given string and assigns it to the ContinentName field.
func (o *GoogleCredentialsListDto) SetContinentName(v string) {
	o.ContinentName = &v
}

func (o GoogleCredentialsListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GoogleCredentialsListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Projects) {
		toSerialize["projects"] = o.Projects
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !IsNil(o.OrganizationName) {
		toSerialize["organizationName"] = o.OrganizationName
	}
	if !IsNil(o.PartnerLogo) {
		toSerialize["partnerLogo"] = o.PartnerLogo
	}
	if !IsNil(o.PartnerName) {
		toSerialize["partnerName"] = o.PartnerName
	}
	if !IsNil(o.FolderId) {
		toSerialize["folderId"] = o.FolderId
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.BillingAccountId) {
		toSerialize["billingAccountId"] = o.BillingAccountId
	}
	if !IsNil(o.Zones) {
		toSerialize["zones"] = o.Zones
	}
	if !IsNil(o.AvailabilityZonesCount) {
		toSerialize["availabilityZonesCount"] = o.AvailabilityZonesCount
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.IsLocked) {
		toSerialize["isLocked"] = o.IsLocked
	}
	if !IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	if !IsNil(o.BillingAccountName) {
		toSerialize["billingAccountName"] = o.BillingAccountName
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.ContinentName) {
		toSerialize["continentName"] = o.ContinentName
	}
	return toSerialize, nil
}

type NullableGoogleCredentialsListDto struct {
	value *GoogleCredentialsListDto
	isSet bool
}

func (v NullableGoogleCredentialsListDto) Get() *GoogleCredentialsListDto {
	return v.value
}

func (v *NullableGoogleCredentialsListDto) Set(val *GoogleCredentialsListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleCredentialsListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleCredentialsListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleCredentialsListDto(val *GoogleCredentialsListDto) *NullableGoogleCredentialsListDto {
	return &NullableGoogleCredentialsListDto{value: val, isSet: true}
}

func (v NullableGoogleCredentialsListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleCredentialsListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


