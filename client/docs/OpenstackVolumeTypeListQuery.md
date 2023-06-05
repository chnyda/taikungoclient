# OpenstackVolumeTypeListQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **int32** |  | [optional] 
**OpenStackUser** | Pointer to **string** |  | [optional] 
**OpenStackPassword** | Pointer to **string** |  | [optional] 
**OpenStackUrl** | Pointer to **string** |  | [optional] 
**OpenStackDomain** | Pointer to **string** |  | [optional] 
**OpenStackRegion** | Pointer to **string** |  | [optional] 
**ApplicationCredEnabled** | Pointer to **bool** |  | [optional] 
**IsAdmin** | Pointer to **bool** |  | [optional] 
**OpenstackProject** | Pointer to **string** |  | [optional] 

## Methods

### NewOpenstackVolumeTypeListQuery

`func NewOpenstackVolumeTypeListQuery() *OpenstackVolumeTypeListQuery`

NewOpenstackVolumeTypeListQuery instantiates a new OpenstackVolumeTypeListQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenstackVolumeTypeListQueryWithDefaults

`func NewOpenstackVolumeTypeListQueryWithDefaults() *OpenstackVolumeTypeListQuery`

NewOpenstackVolumeTypeListQueryWithDefaults instantiates a new OpenstackVolumeTypeListQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *OpenstackVolumeTypeListQuery) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *OpenstackVolumeTypeListQuery) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *OpenstackVolumeTypeListQuery) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *OpenstackVolumeTypeListQuery) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetOpenStackUser

`func (o *OpenstackVolumeTypeListQuery) GetOpenStackUser() string`

GetOpenStackUser returns the OpenStackUser field if non-nil, zero value otherwise.

### GetOpenStackUserOk

`func (o *OpenstackVolumeTypeListQuery) GetOpenStackUserOk() (*string, bool)`

GetOpenStackUserOk returns a tuple with the OpenStackUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackUser

`func (o *OpenstackVolumeTypeListQuery) SetOpenStackUser(v string)`

SetOpenStackUser sets OpenStackUser field to given value.

### HasOpenStackUser

`func (o *OpenstackVolumeTypeListQuery) HasOpenStackUser() bool`

HasOpenStackUser returns a boolean if a field has been set.

### GetOpenStackPassword

`func (o *OpenstackVolumeTypeListQuery) GetOpenStackPassword() string`

GetOpenStackPassword returns the OpenStackPassword field if non-nil, zero value otherwise.

### GetOpenStackPasswordOk

`func (o *OpenstackVolumeTypeListQuery) GetOpenStackPasswordOk() (*string, bool)`

GetOpenStackPasswordOk returns a tuple with the OpenStackPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackPassword

`func (o *OpenstackVolumeTypeListQuery) SetOpenStackPassword(v string)`

SetOpenStackPassword sets OpenStackPassword field to given value.

### HasOpenStackPassword

`func (o *OpenstackVolumeTypeListQuery) HasOpenStackPassword() bool`

HasOpenStackPassword returns a boolean if a field has been set.

### GetOpenStackUrl

`func (o *OpenstackVolumeTypeListQuery) GetOpenStackUrl() string`

GetOpenStackUrl returns the OpenStackUrl field if non-nil, zero value otherwise.

### GetOpenStackUrlOk

`func (o *OpenstackVolumeTypeListQuery) GetOpenStackUrlOk() (*string, bool)`

GetOpenStackUrlOk returns a tuple with the OpenStackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackUrl

`func (o *OpenstackVolumeTypeListQuery) SetOpenStackUrl(v string)`

SetOpenStackUrl sets OpenStackUrl field to given value.

### HasOpenStackUrl

`func (o *OpenstackVolumeTypeListQuery) HasOpenStackUrl() bool`

HasOpenStackUrl returns a boolean if a field has been set.

### GetOpenStackDomain

`func (o *OpenstackVolumeTypeListQuery) GetOpenStackDomain() string`

GetOpenStackDomain returns the OpenStackDomain field if non-nil, zero value otherwise.

### GetOpenStackDomainOk

`func (o *OpenstackVolumeTypeListQuery) GetOpenStackDomainOk() (*string, bool)`

GetOpenStackDomainOk returns a tuple with the OpenStackDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackDomain

`func (o *OpenstackVolumeTypeListQuery) SetOpenStackDomain(v string)`

SetOpenStackDomain sets OpenStackDomain field to given value.

### HasOpenStackDomain

`func (o *OpenstackVolumeTypeListQuery) HasOpenStackDomain() bool`

HasOpenStackDomain returns a boolean if a field has been set.

### GetOpenStackRegion

`func (o *OpenstackVolumeTypeListQuery) GetOpenStackRegion() string`

GetOpenStackRegion returns the OpenStackRegion field if non-nil, zero value otherwise.

### GetOpenStackRegionOk

`func (o *OpenstackVolumeTypeListQuery) GetOpenStackRegionOk() (*string, bool)`

GetOpenStackRegionOk returns a tuple with the OpenStackRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackRegion

`func (o *OpenstackVolumeTypeListQuery) SetOpenStackRegion(v string)`

SetOpenStackRegion sets OpenStackRegion field to given value.

### HasOpenStackRegion

`func (o *OpenstackVolumeTypeListQuery) HasOpenStackRegion() bool`

HasOpenStackRegion returns a boolean if a field has been set.

### GetApplicationCredEnabled

`func (o *OpenstackVolumeTypeListQuery) GetApplicationCredEnabled() bool`

GetApplicationCredEnabled returns the ApplicationCredEnabled field if non-nil, zero value otherwise.

### GetApplicationCredEnabledOk

`func (o *OpenstackVolumeTypeListQuery) GetApplicationCredEnabledOk() (*bool, bool)`

GetApplicationCredEnabledOk returns a tuple with the ApplicationCredEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCredEnabled

`func (o *OpenstackVolumeTypeListQuery) SetApplicationCredEnabled(v bool)`

SetApplicationCredEnabled sets ApplicationCredEnabled field to given value.

### HasApplicationCredEnabled

`func (o *OpenstackVolumeTypeListQuery) HasApplicationCredEnabled() bool`

HasApplicationCredEnabled returns a boolean if a field has been set.

### GetIsAdmin

`func (o *OpenstackVolumeTypeListQuery) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *OpenstackVolumeTypeListQuery) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *OpenstackVolumeTypeListQuery) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *OpenstackVolumeTypeListQuery) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetOpenstackProject

`func (o *OpenstackVolumeTypeListQuery) GetOpenstackProject() string`

GetOpenstackProject returns the OpenstackProject field if non-nil, zero value otherwise.

### GetOpenstackProjectOk

`func (o *OpenstackVolumeTypeListQuery) GetOpenstackProjectOk() (*string, bool)`

GetOpenstackProjectOk returns a tuple with the OpenstackProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenstackProject

`func (o *OpenstackVolumeTypeListQuery) SetOpenstackProject(v string)`

SetOpenstackProject sets OpenstackProject field to given value.

### HasOpenstackProject

`func (o *OpenstackVolumeTypeListQuery) HasOpenstackProject() bool`

HasOpenstackProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


