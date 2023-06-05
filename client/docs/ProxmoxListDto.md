# ProxmoxListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ProjectCount** | Pointer to **int32** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Projects** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**LastModified** | Pointer to **string** |  | [optional] 
**LastModifiedBy** | Pointer to **string** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**OrganizationName** | Pointer to **string** |  | [optional] 
**ContinentName** | Pointer to **string** |  | [optional] 
**Hypervisors** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**TokenId** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Storage** | Pointer to **string** |  | [optional] 
**VmTemplateName** | Pointer to **string** |  | [optional] 
**ProxmoxNetworks** | Pointer to [**[]ProxmoxNetworkListDto**](ProxmoxNetworkListDto.md) |  | [optional] 

## Methods

### NewProxmoxListDto

`func NewProxmoxListDto() *ProxmoxListDto`

NewProxmoxListDto instantiates a new ProxmoxListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxmoxListDtoWithDefaults

`func NewProxmoxListDtoWithDefaults() *ProxmoxListDto`

NewProxmoxListDtoWithDefaults instantiates a new ProxmoxListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProxmoxListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProxmoxListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProxmoxListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProxmoxListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProjectCount

`func (o *ProxmoxListDto) GetProjectCount() int32`

GetProjectCount returns the ProjectCount field if non-nil, zero value otherwise.

### GetProjectCountOk

`func (o *ProxmoxListDto) GetProjectCountOk() (*int32, bool)`

GetProjectCountOk returns a tuple with the ProjectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCount

`func (o *ProxmoxListDto) SetProjectCount(v int32)`

SetProjectCount sets ProjectCount field to given value.

### HasProjectCount

`func (o *ProxmoxListDto) HasProjectCount() bool`

HasProjectCount returns a boolean if a field has been set.

### GetIsLocked

`func (o *ProxmoxListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *ProxmoxListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *ProxmoxListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *ProxmoxListDto) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetName

`func (o *ProxmoxListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProxmoxListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProxmoxListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProxmoxListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjects

`func (o *ProxmoxListDto) GetProjects() []CommonDropdownDto`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *ProxmoxListDto) GetProjectsOk() (*[]CommonDropdownDto, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *ProxmoxListDto) SetProjects(v []CommonDropdownDto)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *ProxmoxListDto) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ProxmoxListDto) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ProxmoxListDto) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ProxmoxListDto) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ProxmoxListDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProxmoxListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProxmoxListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProxmoxListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProxmoxListDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastModified

`func (o *ProxmoxListDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ProxmoxListDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ProxmoxListDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *ProxmoxListDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *ProxmoxListDto) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *ProxmoxListDto) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *ProxmoxListDto) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *ProxmoxListDto) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetIsDefault

`func (o *ProxmoxListDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ProxmoxListDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ProxmoxListDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *ProxmoxListDto) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetOrganizationId

`func (o *ProxmoxListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ProxmoxListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ProxmoxListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ProxmoxListDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetOrganizationName

`func (o *ProxmoxListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *ProxmoxListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *ProxmoxListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *ProxmoxListDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetContinentName

`func (o *ProxmoxListDto) GetContinentName() string`

GetContinentName returns the ContinentName field if non-nil, zero value otherwise.

### GetContinentNameOk

`func (o *ProxmoxListDto) GetContinentNameOk() (*string, bool)`

GetContinentNameOk returns a tuple with the ContinentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinentName

`func (o *ProxmoxListDto) SetContinentName(v string)`

SetContinentName sets ContinentName field to given value.

### HasContinentName

`func (o *ProxmoxListDto) HasContinentName() bool`

HasContinentName returns a boolean if a field has been set.

### GetHypervisors

`func (o *ProxmoxListDto) GetHypervisors() []CommonDropdownDto`

GetHypervisors returns the Hypervisors field if non-nil, zero value otherwise.

### GetHypervisorsOk

`func (o *ProxmoxListDto) GetHypervisorsOk() (*[]CommonDropdownDto, bool)`

GetHypervisorsOk returns a tuple with the Hypervisors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisors

`func (o *ProxmoxListDto) SetHypervisors(v []CommonDropdownDto)`

SetHypervisors sets Hypervisors field to given value.

### HasHypervisors

`func (o *ProxmoxListDto) HasHypervisors() bool`

HasHypervisors returns a boolean if a field has been set.

### GetTokenId

`func (o *ProxmoxListDto) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ProxmoxListDto) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ProxmoxListDto) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *ProxmoxListDto) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetUrl

`func (o *ProxmoxListDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProxmoxListDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProxmoxListDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProxmoxListDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStorage

`func (o *ProxmoxListDto) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ProxmoxListDto) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ProxmoxListDto) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *ProxmoxListDto) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetVmTemplateName

`func (o *ProxmoxListDto) GetVmTemplateName() string`

GetVmTemplateName returns the VmTemplateName field if non-nil, zero value otherwise.

### GetVmTemplateNameOk

`func (o *ProxmoxListDto) GetVmTemplateNameOk() (*string, bool)`

GetVmTemplateNameOk returns a tuple with the VmTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTemplateName

`func (o *ProxmoxListDto) SetVmTemplateName(v string)`

SetVmTemplateName sets VmTemplateName field to given value.

### HasVmTemplateName

`func (o *ProxmoxListDto) HasVmTemplateName() bool`

HasVmTemplateName returns a boolean if a field has been set.

### GetProxmoxNetworks

`func (o *ProxmoxListDto) GetProxmoxNetworks() []ProxmoxNetworkListDto`

GetProxmoxNetworks returns the ProxmoxNetworks field if non-nil, zero value otherwise.

### GetProxmoxNetworksOk

`func (o *ProxmoxListDto) GetProxmoxNetworksOk() (*[]ProxmoxNetworkListDto, bool)`

GetProxmoxNetworksOk returns a tuple with the ProxmoxNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxmoxNetworks

`func (o *ProxmoxListDto) SetProxmoxNetworks(v []ProxmoxNetworkListDto)`

SetProxmoxNetworks sets ProxmoxNetworks field to given value.

### HasProxmoxNetworks

`func (o *ProxmoxListDto) HasProxmoxNetworks() bool`

HasProxmoxNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


