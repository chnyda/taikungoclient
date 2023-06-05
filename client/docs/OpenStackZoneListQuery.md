# OpenStackZoneListQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpenStackUser** | Pointer to **string** |  | [optional] 
**OpenStackPassword** | Pointer to **string** |  | [optional] 
**OpenStackUrl** | Pointer to **string** |  | [optional] 
**OpenStackDomain** | Pointer to **string** |  | [optional] 
**OpenStackRegion** | Pointer to **string** |  | [optional] 
**ApplicationCredEnabled** | Pointer to **bool** |  | [optional] 
**IsAdmin** | Pointer to **bool** |  | [optional] 
**OpenstackProject** | Pointer to **string** |  | [optional] 

## Methods

### NewOpenStackZoneListQuery

`func NewOpenStackZoneListQuery() *OpenStackZoneListQuery`

NewOpenStackZoneListQuery instantiates a new OpenStackZoneListQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenStackZoneListQueryWithDefaults

`func NewOpenStackZoneListQueryWithDefaults() *OpenStackZoneListQuery`

NewOpenStackZoneListQueryWithDefaults instantiates a new OpenStackZoneListQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenStackUser

`func (o *OpenStackZoneListQuery) GetOpenStackUser() string`

GetOpenStackUser returns the OpenStackUser field if non-nil, zero value otherwise.

### GetOpenStackUserOk

`func (o *OpenStackZoneListQuery) GetOpenStackUserOk() (*string, bool)`

GetOpenStackUserOk returns a tuple with the OpenStackUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackUser

`func (o *OpenStackZoneListQuery) SetOpenStackUser(v string)`

SetOpenStackUser sets OpenStackUser field to given value.

### HasOpenStackUser

`func (o *OpenStackZoneListQuery) HasOpenStackUser() bool`

HasOpenStackUser returns a boolean if a field has been set.

### GetOpenStackPassword

`func (o *OpenStackZoneListQuery) GetOpenStackPassword() string`

GetOpenStackPassword returns the OpenStackPassword field if non-nil, zero value otherwise.

### GetOpenStackPasswordOk

`func (o *OpenStackZoneListQuery) GetOpenStackPasswordOk() (*string, bool)`

GetOpenStackPasswordOk returns a tuple with the OpenStackPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackPassword

`func (o *OpenStackZoneListQuery) SetOpenStackPassword(v string)`

SetOpenStackPassword sets OpenStackPassword field to given value.

### HasOpenStackPassword

`func (o *OpenStackZoneListQuery) HasOpenStackPassword() bool`

HasOpenStackPassword returns a boolean if a field has been set.

### GetOpenStackUrl

`func (o *OpenStackZoneListQuery) GetOpenStackUrl() string`

GetOpenStackUrl returns the OpenStackUrl field if non-nil, zero value otherwise.

### GetOpenStackUrlOk

`func (o *OpenStackZoneListQuery) GetOpenStackUrlOk() (*string, bool)`

GetOpenStackUrlOk returns a tuple with the OpenStackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackUrl

`func (o *OpenStackZoneListQuery) SetOpenStackUrl(v string)`

SetOpenStackUrl sets OpenStackUrl field to given value.

### HasOpenStackUrl

`func (o *OpenStackZoneListQuery) HasOpenStackUrl() bool`

HasOpenStackUrl returns a boolean if a field has been set.

### GetOpenStackDomain

`func (o *OpenStackZoneListQuery) GetOpenStackDomain() string`

GetOpenStackDomain returns the OpenStackDomain field if non-nil, zero value otherwise.

### GetOpenStackDomainOk

`func (o *OpenStackZoneListQuery) GetOpenStackDomainOk() (*string, bool)`

GetOpenStackDomainOk returns a tuple with the OpenStackDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackDomain

`func (o *OpenStackZoneListQuery) SetOpenStackDomain(v string)`

SetOpenStackDomain sets OpenStackDomain field to given value.

### HasOpenStackDomain

`func (o *OpenStackZoneListQuery) HasOpenStackDomain() bool`

HasOpenStackDomain returns a boolean if a field has been set.

### GetOpenStackRegion

`func (o *OpenStackZoneListQuery) GetOpenStackRegion() string`

GetOpenStackRegion returns the OpenStackRegion field if non-nil, zero value otherwise.

### GetOpenStackRegionOk

`func (o *OpenStackZoneListQuery) GetOpenStackRegionOk() (*string, bool)`

GetOpenStackRegionOk returns a tuple with the OpenStackRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackRegion

`func (o *OpenStackZoneListQuery) SetOpenStackRegion(v string)`

SetOpenStackRegion sets OpenStackRegion field to given value.

### HasOpenStackRegion

`func (o *OpenStackZoneListQuery) HasOpenStackRegion() bool`

HasOpenStackRegion returns a boolean if a field has been set.

### GetApplicationCredEnabled

`func (o *OpenStackZoneListQuery) GetApplicationCredEnabled() bool`

GetApplicationCredEnabled returns the ApplicationCredEnabled field if non-nil, zero value otherwise.

### GetApplicationCredEnabledOk

`func (o *OpenStackZoneListQuery) GetApplicationCredEnabledOk() (*bool, bool)`

GetApplicationCredEnabledOk returns a tuple with the ApplicationCredEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCredEnabled

`func (o *OpenStackZoneListQuery) SetApplicationCredEnabled(v bool)`

SetApplicationCredEnabled sets ApplicationCredEnabled field to given value.

### HasApplicationCredEnabled

`func (o *OpenStackZoneListQuery) HasApplicationCredEnabled() bool`

HasApplicationCredEnabled returns a boolean if a field has been set.

### GetIsAdmin

`func (o *OpenStackZoneListQuery) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *OpenStackZoneListQuery) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *OpenStackZoneListQuery) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *OpenStackZoneListQuery) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetOpenstackProject

`func (o *OpenStackZoneListQuery) GetOpenstackProject() string`

GetOpenstackProject returns the OpenstackProject field if non-nil, zero value otherwise.

### GetOpenstackProjectOk

`func (o *OpenStackZoneListQuery) GetOpenstackProjectOk() (*string, bool)`

GetOpenstackProjectOk returns a tuple with the OpenstackProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenstackProject

`func (o *OpenStackZoneListQuery) SetOpenstackProject(v string)`

SetOpenstackProject sets OpenstackProject field to given value.

### HasOpenstackProject

`func (o *OpenStackZoneListQuery) HasOpenstackProject() bool`

HasOpenstackProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


