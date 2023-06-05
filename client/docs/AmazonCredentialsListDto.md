# AmazonCredentialsListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ProjectCount** | Pointer to **int32** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**AvailabilityZones** | Pointer to **[]string** |  | [optional] 
**AvailabilityZonesCount** | Pointer to **int32** |  | [optional] 
**Projects** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**LastModified** | Pointer to **string** |  | [optional] 
**LastModifiedBy** | Pointer to **string** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**OrganizationName** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**ContinentName** | Pointer to **string** |  | [optional] 

## Methods

### NewAmazonCredentialsListDto

`func NewAmazonCredentialsListDto() *AmazonCredentialsListDto`

NewAmazonCredentialsListDto instantiates a new AmazonCredentialsListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonCredentialsListDtoWithDefaults

`func NewAmazonCredentialsListDtoWithDefaults() *AmazonCredentialsListDto`

NewAmazonCredentialsListDtoWithDefaults instantiates a new AmazonCredentialsListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AmazonCredentialsListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AmazonCredentialsListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AmazonCredentialsListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AmazonCredentialsListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProjectCount

`func (o *AmazonCredentialsListDto) GetProjectCount() int32`

GetProjectCount returns the ProjectCount field if non-nil, zero value otherwise.

### GetProjectCountOk

`func (o *AmazonCredentialsListDto) GetProjectCountOk() (*int32, bool)`

GetProjectCountOk returns a tuple with the ProjectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCount

`func (o *AmazonCredentialsListDto) SetProjectCount(v int32)`

SetProjectCount sets ProjectCount field to given value.

### HasProjectCount

`func (o *AmazonCredentialsListDto) HasProjectCount() bool`

HasProjectCount returns a boolean if a field has been set.

### GetIsLocked

`func (o *AmazonCredentialsListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *AmazonCredentialsListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *AmazonCredentialsListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *AmazonCredentialsListDto) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetName

`func (o *AmazonCredentialsListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AmazonCredentialsListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AmazonCredentialsListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AmazonCredentialsListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *AmazonCredentialsListDto) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AmazonCredentialsListDto) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AmazonCredentialsListDto) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AmazonCredentialsListDto) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetAvailabilityZones

`func (o *AmazonCredentialsListDto) GetAvailabilityZones() []string`

GetAvailabilityZones returns the AvailabilityZones field if non-nil, zero value otherwise.

### GetAvailabilityZonesOk

`func (o *AmazonCredentialsListDto) GetAvailabilityZonesOk() (*[]string, bool)`

GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZones

`func (o *AmazonCredentialsListDto) SetAvailabilityZones(v []string)`

SetAvailabilityZones sets AvailabilityZones field to given value.

### HasAvailabilityZones

`func (o *AmazonCredentialsListDto) HasAvailabilityZones() bool`

HasAvailabilityZones returns a boolean if a field has been set.

### GetAvailabilityZonesCount

`func (o *AmazonCredentialsListDto) GetAvailabilityZonesCount() int32`

GetAvailabilityZonesCount returns the AvailabilityZonesCount field if non-nil, zero value otherwise.

### GetAvailabilityZonesCountOk

`func (o *AmazonCredentialsListDto) GetAvailabilityZonesCountOk() (*int32, bool)`

GetAvailabilityZonesCountOk returns a tuple with the AvailabilityZonesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZonesCount

`func (o *AmazonCredentialsListDto) SetAvailabilityZonesCount(v int32)`

SetAvailabilityZonesCount sets AvailabilityZonesCount field to given value.

### HasAvailabilityZonesCount

`func (o *AmazonCredentialsListDto) HasAvailabilityZonesCount() bool`

HasAvailabilityZonesCount returns a boolean if a field has been set.

### GetProjects

`func (o *AmazonCredentialsListDto) GetProjects() []CommonDropdownDto`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *AmazonCredentialsListDto) GetProjectsOk() (*[]CommonDropdownDto, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *AmazonCredentialsListDto) SetProjects(v []CommonDropdownDto)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *AmazonCredentialsListDto) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AmazonCredentialsListDto) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AmazonCredentialsListDto) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AmazonCredentialsListDto) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AmazonCredentialsListDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetLastModified

`func (o *AmazonCredentialsListDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *AmazonCredentialsListDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *AmazonCredentialsListDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *AmazonCredentialsListDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *AmazonCredentialsListDto) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *AmazonCredentialsListDto) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *AmazonCredentialsListDto) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *AmazonCredentialsListDto) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetIsDefault

`func (o *AmazonCredentialsListDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *AmazonCredentialsListDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *AmazonCredentialsListDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *AmazonCredentialsListDto) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetOrganizationId

`func (o *AmazonCredentialsListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AmazonCredentialsListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AmazonCredentialsListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *AmazonCredentialsListDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetOrganizationName

`func (o *AmazonCredentialsListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *AmazonCredentialsListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *AmazonCredentialsListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *AmazonCredentialsListDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AmazonCredentialsListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AmazonCredentialsListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AmazonCredentialsListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AmazonCredentialsListDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetContinentName

`func (o *AmazonCredentialsListDto) GetContinentName() string`

GetContinentName returns the ContinentName field if non-nil, zero value otherwise.

### GetContinentNameOk

`func (o *AmazonCredentialsListDto) GetContinentNameOk() (*string, bool)`

GetContinentNameOk returns a tuple with the ContinentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinentName

`func (o *AmazonCredentialsListDto) SetContinentName(v string)`

SetContinentName sets ContinentName field to given value.

### HasContinentName

`func (o *AmazonCredentialsListDto) HasContinentName() bool`

HasContinentName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


