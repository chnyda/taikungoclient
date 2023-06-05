# GoogleCredentialsListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Projects** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**OrganizationName** | Pointer to **string** |  | [optional] 
**PartnerLogo** | Pointer to **string** |  | [optional] 
**PartnerName** | Pointer to **string** |  | [optional] 
**FolderId** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**BillingAccountId** | Pointer to **string** |  | [optional] 
**Zones** | Pointer to **[]string** |  | [optional] 
**AvailabilityZonesCount** | Pointer to **int32** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**BillingAccountName** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**ContinentName** | Pointer to **string** |  | [optional] 

## Methods

### NewGoogleCredentialsListDto

`func NewGoogleCredentialsListDto() *GoogleCredentialsListDto`

NewGoogleCredentialsListDto instantiates a new GoogleCredentialsListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCredentialsListDtoWithDefaults

`func NewGoogleCredentialsListDtoWithDefaults() *GoogleCredentialsListDto`

NewGoogleCredentialsListDtoWithDefaults instantiates a new GoogleCredentialsListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GoogleCredentialsListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GoogleCredentialsListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GoogleCredentialsListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GoogleCredentialsListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GoogleCredentialsListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GoogleCredentialsListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GoogleCredentialsListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GoogleCredentialsListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjects

`func (o *GoogleCredentialsListDto) GetProjects() []CommonDropdownDto`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *GoogleCredentialsListDto) GetProjectsOk() (*[]CommonDropdownDto, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *GoogleCredentialsListDto) SetProjects(v []CommonDropdownDto)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *GoogleCredentialsListDto) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetOrganizationId

`func (o *GoogleCredentialsListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GoogleCredentialsListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GoogleCredentialsListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *GoogleCredentialsListDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetOrganizationName

`func (o *GoogleCredentialsListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *GoogleCredentialsListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *GoogleCredentialsListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *GoogleCredentialsListDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetPartnerLogo

`func (o *GoogleCredentialsListDto) GetPartnerLogo() string`

GetPartnerLogo returns the PartnerLogo field if non-nil, zero value otherwise.

### GetPartnerLogoOk

`func (o *GoogleCredentialsListDto) GetPartnerLogoOk() (*string, bool)`

GetPartnerLogoOk returns a tuple with the PartnerLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerLogo

`func (o *GoogleCredentialsListDto) SetPartnerLogo(v string)`

SetPartnerLogo sets PartnerLogo field to given value.

### HasPartnerLogo

`func (o *GoogleCredentialsListDto) HasPartnerLogo() bool`

HasPartnerLogo returns a boolean if a field has been set.

### GetPartnerName

`func (o *GoogleCredentialsListDto) GetPartnerName() string`

GetPartnerName returns the PartnerName field if non-nil, zero value otherwise.

### GetPartnerNameOk

`func (o *GoogleCredentialsListDto) GetPartnerNameOk() (*string, bool)`

GetPartnerNameOk returns a tuple with the PartnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerName

`func (o *GoogleCredentialsListDto) SetPartnerName(v string)`

SetPartnerName sets PartnerName field to given value.

### HasPartnerName

`func (o *GoogleCredentialsListDto) HasPartnerName() bool`

HasPartnerName returns a boolean if a field has been set.

### GetFolderId

`func (o *GoogleCredentialsListDto) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *GoogleCredentialsListDto) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *GoogleCredentialsListDto) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *GoogleCredentialsListDto) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetProjectId

`func (o *GoogleCredentialsListDto) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GoogleCredentialsListDto) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GoogleCredentialsListDto) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *GoogleCredentialsListDto) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetBillingAccountId

`func (o *GoogleCredentialsListDto) GetBillingAccountId() string`

GetBillingAccountId returns the BillingAccountId field if non-nil, zero value otherwise.

### GetBillingAccountIdOk

`func (o *GoogleCredentialsListDto) GetBillingAccountIdOk() (*string, bool)`

GetBillingAccountIdOk returns a tuple with the BillingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAccountId

`func (o *GoogleCredentialsListDto) SetBillingAccountId(v string)`

SetBillingAccountId sets BillingAccountId field to given value.

### HasBillingAccountId

`func (o *GoogleCredentialsListDto) HasBillingAccountId() bool`

HasBillingAccountId returns a boolean if a field has been set.

### GetZones

`func (o *GoogleCredentialsListDto) GetZones() []string`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *GoogleCredentialsListDto) GetZonesOk() (*[]string, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *GoogleCredentialsListDto) SetZones(v []string)`

SetZones sets Zones field to given value.

### HasZones

`func (o *GoogleCredentialsListDto) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetAvailabilityZonesCount

`func (o *GoogleCredentialsListDto) GetAvailabilityZonesCount() int32`

GetAvailabilityZonesCount returns the AvailabilityZonesCount field if non-nil, zero value otherwise.

### GetAvailabilityZonesCountOk

`func (o *GoogleCredentialsListDto) GetAvailabilityZonesCountOk() (*int32, bool)`

GetAvailabilityZonesCountOk returns a tuple with the AvailabilityZonesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZonesCount

`func (o *GoogleCredentialsListDto) SetAvailabilityZonesCount(v int32)`

SetAvailabilityZonesCount sets AvailabilityZonesCount field to given value.

### HasAvailabilityZonesCount

`func (o *GoogleCredentialsListDto) HasAvailabilityZonesCount() bool`

HasAvailabilityZonesCount returns a boolean if a field has been set.

### GetRegion

`func (o *GoogleCredentialsListDto) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GoogleCredentialsListDto) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GoogleCredentialsListDto) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GoogleCredentialsListDto) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetIsLocked

`func (o *GoogleCredentialsListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *GoogleCredentialsListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *GoogleCredentialsListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *GoogleCredentialsListDto) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetIsDefault

`func (o *GoogleCredentialsListDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *GoogleCredentialsListDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *GoogleCredentialsListDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *GoogleCredentialsListDto) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetBillingAccountName

`func (o *GoogleCredentialsListDto) GetBillingAccountName() string`

GetBillingAccountName returns the BillingAccountName field if non-nil, zero value otherwise.

### GetBillingAccountNameOk

`func (o *GoogleCredentialsListDto) GetBillingAccountNameOk() (*string, bool)`

GetBillingAccountNameOk returns a tuple with the BillingAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAccountName

`func (o *GoogleCredentialsListDto) SetBillingAccountName(v string)`

SetBillingAccountName sets BillingAccountName field to given value.

### HasBillingAccountName

`func (o *GoogleCredentialsListDto) HasBillingAccountName() bool`

HasBillingAccountName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GoogleCredentialsListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GoogleCredentialsListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GoogleCredentialsListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GoogleCredentialsListDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetContinentName

`func (o *GoogleCredentialsListDto) GetContinentName() string`

GetContinentName returns the ContinentName field if non-nil, zero value otherwise.

### GetContinentNameOk

`func (o *GoogleCredentialsListDto) GetContinentNameOk() (*string, bool)`

GetContinentNameOk returns a tuple with the ContinentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinentName

`func (o *GoogleCredentialsListDto) SetContinentName(v string)`

SetContinentName sets ContinentName field to given value.

### HasContinentName

`func (o *GoogleCredentialsListDto) HasContinentName() bool`

HasContinentName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


