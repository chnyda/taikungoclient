# ProjectFullListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Operation** | Pointer to **string** |  | [optional] 
**JobUrl** | Pointer to **string** |  | [optional] 
**ImageName** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**IsKubernetes** | Pointer to **bool** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**IsBackupEnabled** | Pointer to **bool** |  | [optional] 
**IsMonitoringEnabled** | Pointer to **bool** |  | [optional] 
**IsAutoUpgrade** | Pointer to **bool** |  | [optional] 
**AccessProfileRevision** | Pointer to **int32** |  | [optional] 
**CloudCredentialName** | Pointer to **string** |  | [optional] 
**OrganizationName** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Health** | Pointer to **string** |  | [optional] 
**AccessIp** | Pointer to **string** |  | [optional] 
**CloudType** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**AccessProfiles** | Pointer to [**AccessProfilesForProjectListDto**](AccessProfilesForProjectListDto.md) |  | [optional] 
**IsDeleteCluster** | Pointer to **bool** |  | [optional] 
**TaikunPrivateSSHKey** | Pointer to **string** |  | [optional] 
**TaikunPublicSSHKey** | Pointer to **string** |  | [optional] 
**StandaloneVms** | Pointer to [**[]StandAloneVmFullDto**](StandAloneVmFullDto.md) |  | [optional] 
**Cidr** | Pointer to **string** |  | [optional] 
**NetMask** | Pointer to **int32** |  | [optional] 
**PrivateIp** | Pointer to **string** |  | [optional] 
**PublicIp** | Pointer to **string** |  | [optional] 
**IsKubevapEnabled** | Pointer to **bool** |  | [optional] 
**TanzuReleaseVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewProjectFullListDto

`func NewProjectFullListDto() *ProjectFullListDto`

NewProjectFullListDto instantiates a new ProjectFullListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectFullListDtoWithDefaults

`func NewProjectFullListDtoWithDefaults() *ProjectFullListDto`

NewProjectFullListDtoWithDefaults instantiates a new ProjectFullListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectFullListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectFullListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectFullListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectFullListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProjectFullListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectFullListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectFullListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectFullListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperation

`func (o *ProjectFullListDto) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ProjectFullListDto) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ProjectFullListDto) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *ProjectFullListDto) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetJobUrl

`func (o *ProjectFullListDto) GetJobUrl() string`

GetJobUrl returns the JobUrl field if non-nil, zero value otherwise.

### GetJobUrlOk

`func (o *ProjectFullListDto) GetJobUrlOk() (*string, bool)`

GetJobUrlOk returns a tuple with the JobUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobUrl

`func (o *ProjectFullListDto) SetJobUrl(v string)`

SetJobUrl sets JobUrl field to given value.

### HasJobUrl

`func (o *ProjectFullListDto) HasJobUrl() bool`

HasJobUrl returns a boolean if a field has been set.

### GetImageName

`func (o *ProjectFullListDto) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *ProjectFullListDto) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *ProjectFullListDto) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *ProjectFullListDto) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetToken

`func (o *ProjectFullListDto) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ProjectFullListDto) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ProjectFullListDto) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ProjectFullListDto) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetIsKubernetes

`func (o *ProjectFullListDto) GetIsKubernetes() bool`

GetIsKubernetes returns the IsKubernetes field if non-nil, zero value otherwise.

### GetIsKubernetesOk

`func (o *ProjectFullListDto) GetIsKubernetesOk() (*bool, bool)`

GetIsKubernetesOk returns a tuple with the IsKubernetes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKubernetes

`func (o *ProjectFullListDto) SetIsKubernetes(v bool)`

SetIsKubernetes sets IsKubernetes field to given value.

### HasIsKubernetes

`func (o *ProjectFullListDto) HasIsKubernetes() bool`

HasIsKubernetes returns a boolean if a field has been set.

### GetIsLocked

`func (o *ProjectFullListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *ProjectFullListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *ProjectFullListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *ProjectFullListDto) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetIsBackupEnabled

`func (o *ProjectFullListDto) GetIsBackupEnabled() bool`

GetIsBackupEnabled returns the IsBackupEnabled field if non-nil, zero value otherwise.

### GetIsBackupEnabledOk

`func (o *ProjectFullListDto) GetIsBackupEnabledOk() (*bool, bool)`

GetIsBackupEnabledOk returns a tuple with the IsBackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBackupEnabled

`func (o *ProjectFullListDto) SetIsBackupEnabled(v bool)`

SetIsBackupEnabled sets IsBackupEnabled field to given value.

### HasIsBackupEnabled

`func (o *ProjectFullListDto) HasIsBackupEnabled() bool`

HasIsBackupEnabled returns a boolean if a field has been set.

### GetIsMonitoringEnabled

`func (o *ProjectFullListDto) GetIsMonitoringEnabled() bool`

GetIsMonitoringEnabled returns the IsMonitoringEnabled field if non-nil, zero value otherwise.

### GetIsMonitoringEnabledOk

`func (o *ProjectFullListDto) GetIsMonitoringEnabledOk() (*bool, bool)`

GetIsMonitoringEnabledOk returns a tuple with the IsMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMonitoringEnabled

`func (o *ProjectFullListDto) SetIsMonitoringEnabled(v bool)`

SetIsMonitoringEnabled sets IsMonitoringEnabled field to given value.

### HasIsMonitoringEnabled

`func (o *ProjectFullListDto) HasIsMonitoringEnabled() bool`

HasIsMonitoringEnabled returns a boolean if a field has been set.

### GetIsAutoUpgrade

`func (o *ProjectFullListDto) GetIsAutoUpgrade() bool`

GetIsAutoUpgrade returns the IsAutoUpgrade field if non-nil, zero value otherwise.

### GetIsAutoUpgradeOk

`func (o *ProjectFullListDto) GetIsAutoUpgradeOk() (*bool, bool)`

GetIsAutoUpgradeOk returns a tuple with the IsAutoUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoUpgrade

`func (o *ProjectFullListDto) SetIsAutoUpgrade(v bool)`

SetIsAutoUpgrade sets IsAutoUpgrade field to given value.

### HasIsAutoUpgrade

`func (o *ProjectFullListDto) HasIsAutoUpgrade() bool`

HasIsAutoUpgrade returns a boolean if a field has been set.

### GetAccessProfileRevision

`func (o *ProjectFullListDto) GetAccessProfileRevision() int32`

GetAccessProfileRevision returns the AccessProfileRevision field if non-nil, zero value otherwise.

### GetAccessProfileRevisionOk

`func (o *ProjectFullListDto) GetAccessProfileRevisionOk() (*int32, bool)`

GetAccessProfileRevisionOk returns a tuple with the AccessProfileRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfileRevision

`func (o *ProjectFullListDto) SetAccessProfileRevision(v int32)`

SetAccessProfileRevision sets AccessProfileRevision field to given value.

### HasAccessProfileRevision

`func (o *ProjectFullListDto) HasAccessProfileRevision() bool`

HasAccessProfileRevision returns a boolean if a field has been set.

### GetCloudCredentialName

`func (o *ProjectFullListDto) GetCloudCredentialName() string`

GetCloudCredentialName returns the CloudCredentialName field if non-nil, zero value otherwise.

### GetCloudCredentialNameOk

`func (o *ProjectFullListDto) GetCloudCredentialNameOk() (*string, bool)`

GetCloudCredentialNameOk returns a tuple with the CloudCredentialName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCredentialName

`func (o *ProjectFullListDto) SetCloudCredentialName(v string)`

SetCloudCredentialName sets CloudCredentialName field to given value.

### HasCloudCredentialName

`func (o *ProjectFullListDto) HasCloudCredentialName() bool`

HasCloudCredentialName returns a boolean if a field has been set.

### GetOrganizationName

`func (o *ProjectFullListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *ProjectFullListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *ProjectFullListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *ProjectFullListDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetOrganizationId

`func (o *ProjectFullListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ProjectFullListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ProjectFullListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ProjectFullListDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetStatus

`func (o *ProjectFullListDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectFullListDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectFullListDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProjectFullListDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetHealth

`func (o *ProjectFullListDto) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ProjectFullListDto) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ProjectFullListDto) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *ProjectFullListDto) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetAccessIp

`func (o *ProjectFullListDto) GetAccessIp() string`

GetAccessIp returns the AccessIp field if non-nil, zero value otherwise.

### GetAccessIpOk

`func (o *ProjectFullListDto) GetAccessIpOk() (*string, bool)`

GetAccessIpOk returns a tuple with the AccessIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessIp

`func (o *ProjectFullListDto) SetAccessIp(v string)`

SetAccessIp sets AccessIp field to given value.

### HasAccessIp

`func (o *ProjectFullListDto) HasAccessIp() bool`

HasAccessIp returns a boolean if a field has been set.

### GetCloudType

`func (o *ProjectFullListDto) GetCloudType() string`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *ProjectFullListDto) GetCloudTypeOk() (*string, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *ProjectFullListDto) SetCloudType(v string)`

SetCloudType sets CloudType field to given value.

### HasCloudType

`func (o *ProjectFullListDto) HasCloudType() bool`

HasCloudType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProjectFullListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectFullListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectFullListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProjectFullListDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ProjectFullListDto) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProjectFullListDto) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProjectFullListDto) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProjectFullListDto) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetAccessProfiles

`func (o *ProjectFullListDto) GetAccessProfiles() AccessProfilesForProjectListDto`

GetAccessProfiles returns the AccessProfiles field if non-nil, zero value otherwise.

### GetAccessProfilesOk

`func (o *ProjectFullListDto) GetAccessProfilesOk() (*AccessProfilesForProjectListDto, bool)`

GetAccessProfilesOk returns a tuple with the AccessProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfiles

`func (o *ProjectFullListDto) SetAccessProfiles(v AccessProfilesForProjectListDto)`

SetAccessProfiles sets AccessProfiles field to given value.

### HasAccessProfiles

`func (o *ProjectFullListDto) HasAccessProfiles() bool`

HasAccessProfiles returns a boolean if a field has been set.

### GetIsDeleteCluster

`func (o *ProjectFullListDto) GetIsDeleteCluster() bool`

GetIsDeleteCluster returns the IsDeleteCluster field if non-nil, zero value otherwise.

### GetIsDeleteClusterOk

`func (o *ProjectFullListDto) GetIsDeleteClusterOk() (*bool, bool)`

GetIsDeleteClusterOk returns a tuple with the IsDeleteCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleteCluster

`func (o *ProjectFullListDto) SetIsDeleteCluster(v bool)`

SetIsDeleteCluster sets IsDeleteCluster field to given value.

### HasIsDeleteCluster

`func (o *ProjectFullListDto) HasIsDeleteCluster() bool`

HasIsDeleteCluster returns a boolean if a field has been set.

### GetTaikunPrivateSSHKey

`func (o *ProjectFullListDto) GetTaikunPrivateSSHKey() string`

GetTaikunPrivateSSHKey returns the TaikunPrivateSSHKey field if non-nil, zero value otherwise.

### GetTaikunPrivateSSHKeyOk

`func (o *ProjectFullListDto) GetTaikunPrivateSSHKeyOk() (*string, bool)`

GetTaikunPrivateSSHKeyOk returns a tuple with the TaikunPrivateSSHKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaikunPrivateSSHKey

`func (o *ProjectFullListDto) SetTaikunPrivateSSHKey(v string)`

SetTaikunPrivateSSHKey sets TaikunPrivateSSHKey field to given value.

### HasTaikunPrivateSSHKey

`func (o *ProjectFullListDto) HasTaikunPrivateSSHKey() bool`

HasTaikunPrivateSSHKey returns a boolean if a field has been set.

### GetTaikunPublicSSHKey

`func (o *ProjectFullListDto) GetTaikunPublicSSHKey() string`

GetTaikunPublicSSHKey returns the TaikunPublicSSHKey field if non-nil, zero value otherwise.

### GetTaikunPublicSSHKeyOk

`func (o *ProjectFullListDto) GetTaikunPublicSSHKeyOk() (*string, bool)`

GetTaikunPublicSSHKeyOk returns a tuple with the TaikunPublicSSHKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaikunPublicSSHKey

`func (o *ProjectFullListDto) SetTaikunPublicSSHKey(v string)`

SetTaikunPublicSSHKey sets TaikunPublicSSHKey field to given value.

### HasTaikunPublicSSHKey

`func (o *ProjectFullListDto) HasTaikunPublicSSHKey() bool`

HasTaikunPublicSSHKey returns a boolean if a field has been set.

### GetStandaloneVms

`func (o *ProjectFullListDto) GetStandaloneVms() []StandAloneVmFullDto`

GetStandaloneVms returns the StandaloneVms field if non-nil, zero value otherwise.

### GetStandaloneVmsOk

`func (o *ProjectFullListDto) GetStandaloneVmsOk() (*[]StandAloneVmFullDto, bool)`

GetStandaloneVmsOk returns a tuple with the StandaloneVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandaloneVms

`func (o *ProjectFullListDto) SetStandaloneVms(v []StandAloneVmFullDto)`

SetStandaloneVms sets StandaloneVms field to given value.

### HasStandaloneVms

`func (o *ProjectFullListDto) HasStandaloneVms() bool`

HasStandaloneVms returns a boolean if a field has been set.

### GetCidr

`func (o *ProjectFullListDto) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *ProjectFullListDto) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *ProjectFullListDto) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *ProjectFullListDto) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetNetMask

`func (o *ProjectFullListDto) GetNetMask() int32`

GetNetMask returns the NetMask field if non-nil, zero value otherwise.

### GetNetMaskOk

`func (o *ProjectFullListDto) GetNetMaskOk() (*int32, bool)`

GetNetMaskOk returns a tuple with the NetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetMask

`func (o *ProjectFullListDto) SetNetMask(v int32)`

SetNetMask sets NetMask field to given value.

### HasNetMask

`func (o *ProjectFullListDto) HasNetMask() bool`

HasNetMask returns a boolean if a field has been set.

### GetPrivateIp

`func (o *ProjectFullListDto) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *ProjectFullListDto) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *ProjectFullListDto) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *ProjectFullListDto) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetPublicIp

`func (o *ProjectFullListDto) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *ProjectFullListDto) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *ProjectFullListDto) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *ProjectFullListDto) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetIsKubevapEnabled

`func (o *ProjectFullListDto) GetIsKubevapEnabled() bool`

GetIsKubevapEnabled returns the IsKubevapEnabled field if non-nil, zero value otherwise.

### GetIsKubevapEnabledOk

`func (o *ProjectFullListDto) GetIsKubevapEnabledOk() (*bool, bool)`

GetIsKubevapEnabledOk returns a tuple with the IsKubevapEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKubevapEnabled

`func (o *ProjectFullListDto) SetIsKubevapEnabled(v bool)`

SetIsKubevapEnabled sets IsKubevapEnabled field to given value.

### HasIsKubevapEnabled

`func (o *ProjectFullListDto) HasIsKubevapEnabled() bool`

HasIsKubevapEnabled returns a boolean if a field has been set.

### GetTanzuReleaseVersion

`func (o *ProjectFullListDto) GetTanzuReleaseVersion() string`

GetTanzuReleaseVersion returns the TanzuReleaseVersion field if non-nil, zero value otherwise.

### GetTanzuReleaseVersionOk

`func (o *ProjectFullListDto) GetTanzuReleaseVersionOk() (*string, bool)`

GetTanzuReleaseVersionOk returns a tuple with the TanzuReleaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTanzuReleaseVersion

`func (o *ProjectFullListDto) SetTanzuReleaseVersion(v string)`

SetTanzuReleaseVersion sets TanzuReleaseVersion field to given value.

### HasTanzuReleaseVersion

`func (o *ProjectFullListDto) HasTanzuReleaseVersion() bool`

HasTanzuReleaseVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


