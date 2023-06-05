# CreateOpenstackCloudCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**OpenStackUser** | **string** |  | 
**OpenStackPassword** | **string** |  | 
**OpenStackUrl** | **string** |  | 
**OpenStackProject** | Pointer to **string** |  | [optional] 
**OpenStackPublicNetwork** | **string** |  | 
**OpenStackAvailabilityZone** | Pointer to **string** |  | [optional] 
**OpenStackDomain** | Pointer to **string** |  | [optional] 
**OpenStackRegion** | **string** |  | 
**OpenStackContinent** | Pointer to **string** |  | [optional] 
**OpenStackVolumeType** | Pointer to **string** |  | [optional] 
**OpenStackImportNetwork** | Pointer to **bool** |  | [optional] 
**OpenStackInternalSubnetId** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**ApplicationCredEnabled** | Pointer to **bool** |  | [optional] 
**IsAdmin** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateOpenstackCloudCommand

`func NewCreateOpenstackCloudCommand(name string, openStackUser string, openStackPassword string, openStackUrl string, openStackPublicNetwork string, openStackRegion string, ) *CreateOpenstackCloudCommand`

NewCreateOpenstackCloudCommand instantiates a new CreateOpenstackCloudCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOpenstackCloudCommandWithDefaults

`func NewCreateOpenstackCloudCommandWithDefaults() *CreateOpenstackCloudCommand`

NewCreateOpenstackCloudCommandWithDefaults instantiates a new CreateOpenstackCloudCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateOpenstackCloudCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOpenstackCloudCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOpenstackCloudCommand) SetName(v string)`

SetName sets Name field to given value.


### GetOpenStackUser

`func (o *CreateOpenstackCloudCommand) GetOpenStackUser() string`

GetOpenStackUser returns the OpenStackUser field if non-nil, zero value otherwise.

### GetOpenStackUserOk

`func (o *CreateOpenstackCloudCommand) GetOpenStackUserOk() (*string, bool)`

GetOpenStackUserOk returns a tuple with the OpenStackUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackUser

`func (o *CreateOpenstackCloudCommand) SetOpenStackUser(v string)`

SetOpenStackUser sets OpenStackUser field to given value.


### GetOpenStackPassword

`func (o *CreateOpenstackCloudCommand) GetOpenStackPassword() string`

GetOpenStackPassword returns the OpenStackPassword field if non-nil, zero value otherwise.

### GetOpenStackPasswordOk

`func (o *CreateOpenstackCloudCommand) GetOpenStackPasswordOk() (*string, bool)`

GetOpenStackPasswordOk returns a tuple with the OpenStackPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackPassword

`func (o *CreateOpenstackCloudCommand) SetOpenStackPassword(v string)`

SetOpenStackPassword sets OpenStackPassword field to given value.


### GetOpenStackUrl

`func (o *CreateOpenstackCloudCommand) GetOpenStackUrl() string`

GetOpenStackUrl returns the OpenStackUrl field if non-nil, zero value otherwise.

### GetOpenStackUrlOk

`func (o *CreateOpenstackCloudCommand) GetOpenStackUrlOk() (*string, bool)`

GetOpenStackUrlOk returns a tuple with the OpenStackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackUrl

`func (o *CreateOpenstackCloudCommand) SetOpenStackUrl(v string)`

SetOpenStackUrl sets OpenStackUrl field to given value.


### GetOpenStackProject

`func (o *CreateOpenstackCloudCommand) GetOpenStackProject() string`

GetOpenStackProject returns the OpenStackProject field if non-nil, zero value otherwise.

### GetOpenStackProjectOk

`func (o *CreateOpenstackCloudCommand) GetOpenStackProjectOk() (*string, bool)`

GetOpenStackProjectOk returns a tuple with the OpenStackProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackProject

`func (o *CreateOpenstackCloudCommand) SetOpenStackProject(v string)`

SetOpenStackProject sets OpenStackProject field to given value.

### HasOpenStackProject

`func (o *CreateOpenstackCloudCommand) HasOpenStackProject() bool`

HasOpenStackProject returns a boolean if a field has been set.

### GetOpenStackPublicNetwork

`func (o *CreateOpenstackCloudCommand) GetOpenStackPublicNetwork() string`

GetOpenStackPublicNetwork returns the OpenStackPublicNetwork field if non-nil, zero value otherwise.

### GetOpenStackPublicNetworkOk

`func (o *CreateOpenstackCloudCommand) GetOpenStackPublicNetworkOk() (*string, bool)`

GetOpenStackPublicNetworkOk returns a tuple with the OpenStackPublicNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackPublicNetwork

`func (o *CreateOpenstackCloudCommand) SetOpenStackPublicNetwork(v string)`

SetOpenStackPublicNetwork sets OpenStackPublicNetwork field to given value.


### GetOpenStackAvailabilityZone

`func (o *CreateOpenstackCloudCommand) GetOpenStackAvailabilityZone() string`

GetOpenStackAvailabilityZone returns the OpenStackAvailabilityZone field if non-nil, zero value otherwise.

### GetOpenStackAvailabilityZoneOk

`func (o *CreateOpenstackCloudCommand) GetOpenStackAvailabilityZoneOk() (*string, bool)`

GetOpenStackAvailabilityZoneOk returns a tuple with the OpenStackAvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackAvailabilityZone

`func (o *CreateOpenstackCloudCommand) SetOpenStackAvailabilityZone(v string)`

SetOpenStackAvailabilityZone sets OpenStackAvailabilityZone field to given value.

### HasOpenStackAvailabilityZone

`func (o *CreateOpenstackCloudCommand) HasOpenStackAvailabilityZone() bool`

HasOpenStackAvailabilityZone returns a boolean if a field has been set.

### GetOpenStackDomain

`func (o *CreateOpenstackCloudCommand) GetOpenStackDomain() string`

GetOpenStackDomain returns the OpenStackDomain field if non-nil, zero value otherwise.

### GetOpenStackDomainOk

`func (o *CreateOpenstackCloudCommand) GetOpenStackDomainOk() (*string, bool)`

GetOpenStackDomainOk returns a tuple with the OpenStackDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackDomain

`func (o *CreateOpenstackCloudCommand) SetOpenStackDomain(v string)`

SetOpenStackDomain sets OpenStackDomain field to given value.

### HasOpenStackDomain

`func (o *CreateOpenstackCloudCommand) HasOpenStackDomain() bool`

HasOpenStackDomain returns a boolean if a field has been set.

### GetOpenStackRegion

`func (o *CreateOpenstackCloudCommand) GetOpenStackRegion() string`

GetOpenStackRegion returns the OpenStackRegion field if non-nil, zero value otherwise.

### GetOpenStackRegionOk

`func (o *CreateOpenstackCloudCommand) GetOpenStackRegionOk() (*string, bool)`

GetOpenStackRegionOk returns a tuple with the OpenStackRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackRegion

`func (o *CreateOpenstackCloudCommand) SetOpenStackRegion(v string)`

SetOpenStackRegion sets OpenStackRegion field to given value.


### GetOpenStackContinent

`func (o *CreateOpenstackCloudCommand) GetOpenStackContinent() string`

GetOpenStackContinent returns the OpenStackContinent field if non-nil, zero value otherwise.

### GetOpenStackContinentOk

`func (o *CreateOpenstackCloudCommand) GetOpenStackContinentOk() (*string, bool)`

GetOpenStackContinentOk returns a tuple with the OpenStackContinent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackContinent

`func (o *CreateOpenstackCloudCommand) SetOpenStackContinent(v string)`

SetOpenStackContinent sets OpenStackContinent field to given value.

### HasOpenStackContinent

`func (o *CreateOpenstackCloudCommand) HasOpenStackContinent() bool`

HasOpenStackContinent returns a boolean if a field has been set.

### GetOpenStackVolumeType

`func (o *CreateOpenstackCloudCommand) GetOpenStackVolumeType() string`

GetOpenStackVolumeType returns the OpenStackVolumeType field if non-nil, zero value otherwise.

### GetOpenStackVolumeTypeOk

`func (o *CreateOpenstackCloudCommand) GetOpenStackVolumeTypeOk() (*string, bool)`

GetOpenStackVolumeTypeOk returns a tuple with the OpenStackVolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackVolumeType

`func (o *CreateOpenstackCloudCommand) SetOpenStackVolumeType(v string)`

SetOpenStackVolumeType sets OpenStackVolumeType field to given value.

### HasOpenStackVolumeType

`func (o *CreateOpenstackCloudCommand) HasOpenStackVolumeType() bool`

HasOpenStackVolumeType returns a boolean if a field has been set.

### GetOpenStackImportNetwork

`func (o *CreateOpenstackCloudCommand) GetOpenStackImportNetwork() bool`

GetOpenStackImportNetwork returns the OpenStackImportNetwork field if non-nil, zero value otherwise.

### GetOpenStackImportNetworkOk

`func (o *CreateOpenstackCloudCommand) GetOpenStackImportNetworkOk() (*bool, bool)`

GetOpenStackImportNetworkOk returns a tuple with the OpenStackImportNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackImportNetwork

`func (o *CreateOpenstackCloudCommand) SetOpenStackImportNetwork(v bool)`

SetOpenStackImportNetwork sets OpenStackImportNetwork field to given value.

### HasOpenStackImportNetwork

`func (o *CreateOpenstackCloudCommand) HasOpenStackImportNetwork() bool`

HasOpenStackImportNetwork returns a boolean if a field has been set.

### GetOpenStackInternalSubnetId

`func (o *CreateOpenstackCloudCommand) GetOpenStackInternalSubnetId() string`

GetOpenStackInternalSubnetId returns the OpenStackInternalSubnetId field if non-nil, zero value otherwise.

### GetOpenStackInternalSubnetIdOk

`func (o *CreateOpenstackCloudCommand) GetOpenStackInternalSubnetIdOk() (*string, bool)`

GetOpenStackInternalSubnetIdOk returns a tuple with the OpenStackInternalSubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackInternalSubnetId

`func (o *CreateOpenstackCloudCommand) SetOpenStackInternalSubnetId(v string)`

SetOpenStackInternalSubnetId sets OpenStackInternalSubnetId field to given value.

### HasOpenStackInternalSubnetId

`func (o *CreateOpenstackCloudCommand) HasOpenStackInternalSubnetId() bool`

HasOpenStackInternalSubnetId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *CreateOpenstackCloudCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateOpenstackCloudCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateOpenstackCloudCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CreateOpenstackCloudCommand) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetApplicationCredEnabled

`func (o *CreateOpenstackCloudCommand) GetApplicationCredEnabled() bool`

GetApplicationCredEnabled returns the ApplicationCredEnabled field if non-nil, zero value otherwise.

### GetApplicationCredEnabledOk

`func (o *CreateOpenstackCloudCommand) GetApplicationCredEnabledOk() (*bool, bool)`

GetApplicationCredEnabledOk returns a tuple with the ApplicationCredEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCredEnabled

`func (o *CreateOpenstackCloudCommand) SetApplicationCredEnabled(v bool)`

SetApplicationCredEnabled sets ApplicationCredEnabled field to given value.

### HasApplicationCredEnabled

`func (o *CreateOpenstackCloudCommand) HasApplicationCredEnabled() bool`

HasApplicationCredEnabled returns a boolean if a field has been set.

### GetIsAdmin

`func (o *CreateOpenstackCloudCommand) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *CreateOpenstackCloudCommand) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *CreateOpenstackCloudCommand) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *CreateOpenstackCloudCommand) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


