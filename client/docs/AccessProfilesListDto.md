# AccessProfilesListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**HttpProxy** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**OrganizationName** | Pointer to **string** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**DnsServers** | Pointer to [**[]DnsServerListDto**](DnsServerListDto.md) |  | [optional] 
**NtpServers** | Pointer to [**[]NtpServerListDto**](NtpServerListDto.md) |  | [optional] 
**AllowedHosts** | Pointer to [**[]AllowedHostListDto**](AllowedHostListDto.md) |  | [optional] 
**Projects** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**LastModified** | Pointer to **string** |  | [optional] 
**LastModifiedBy** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewAccessProfilesListDto

`func NewAccessProfilesListDto() *AccessProfilesListDto`

NewAccessProfilesListDto instantiates a new AccessProfilesListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessProfilesListDtoWithDefaults

`func NewAccessProfilesListDtoWithDefaults() *AccessProfilesListDto`

NewAccessProfilesListDtoWithDefaults instantiates a new AccessProfilesListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccessProfilesListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessProfilesListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessProfilesListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AccessProfilesListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AccessProfilesListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessProfilesListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessProfilesListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccessProfilesListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHttpProxy

`func (o *AccessProfilesListDto) GetHttpProxy() string`

GetHttpProxy returns the HttpProxy field if non-nil, zero value otherwise.

### GetHttpProxyOk

`func (o *AccessProfilesListDto) GetHttpProxyOk() (*string, bool)`

GetHttpProxyOk returns a tuple with the HttpProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxy

`func (o *AccessProfilesListDto) SetHttpProxy(v string)`

SetHttpProxy sets HttpProxy field to given value.

### HasHttpProxy

`func (o *AccessProfilesListDto) HasHttpProxy() bool`

HasHttpProxy returns a boolean if a field has been set.

### GetOrganizationId

`func (o *AccessProfilesListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AccessProfilesListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AccessProfilesListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *AccessProfilesListDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetOrganizationName

`func (o *AccessProfilesListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *AccessProfilesListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *AccessProfilesListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *AccessProfilesListDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetIsLocked

`func (o *AccessProfilesListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *AccessProfilesListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *AccessProfilesListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *AccessProfilesListDto) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetDnsServers

`func (o *AccessProfilesListDto) GetDnsServers() []DnsServerListDto`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *AccessProfilesListDto) GetDnsServersOk() (*[]DnsServerListDto, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *AccessProfilesListDto) SetDnsServers(v []DnsServerListDto)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *AccessProfilesListDto) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### GetNtpServers

`func (o *AccessProfilesListDto) GetNtpServers() []NtpServerListDto`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *AccessProfilesListDto) GetNtpServersOk() (*[]NtpServerListDto, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *AccessProfilesListDto) SetNtpServers(v []NtpServerListDto)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *AccessProfilesListDto) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### GetAllowedHosts

`func (o *AccessProfilesListDto) GetAllowedHosts() []AllowedHostListDto`

GetAllowedHosts returns the AllowedHosts field if non-nil, zero value otherwise.

### GetAllowedHostsOk

`func (o *AccessProfilesListDto) GetAllowedHostsOk() (*[]AllowedHostListDto, bool)`

GetAllowedHostsOk returns a tuple with the AllowedHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedHosts

`func (o *AccessProfilesListDto) SetAllowedHosts(v []AllowedHostListDto)`

SetAllowedHosts sets AllowedHosts field to given value.

### HasAllowedHosts

`func (o *AccessProfilesListDto) HasAllowedHosts() bool`

HasAllowedHosts returns a boolean if a field has been set.

### GetProjects

`func (o *AccessProfilesListDto) GetProjects() []CommonDropdownDto`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *AccessProfilesListDto) GetProjectsOk() (*[]CommonDropdownDto, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *AccessProfilesListDto) SetProjects(v []CommonDropdownDto)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *AccessProfilesListDto) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AccessProfilesListDto) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AccessProfilesListDto) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AccessProfilesListDto) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AccessProfilesListDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetLastModified

`func (o *AccessProfilesListDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *AccessProfilesListDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *AccessProfilesListDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *AccessProfilesListDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *AccessProfilesListDto) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *AccessProfilesListDto) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *AccessProfilesListDto) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *AccessProfilesListDto) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AccessProfilesListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccessProfilesListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccessProfilesListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AccessProfilesListDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


