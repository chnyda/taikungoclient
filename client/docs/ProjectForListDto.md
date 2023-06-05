# ProjectForListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Operation** | Pointer to **string** |  | [optional] 
**JobUrl** | Pointer to **string** |  | [optional] 
**TopicName** | Pointer to **string** |  | [optional] 
**ImageName** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**IsKubernetes** | Pointer to **bool** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**IsBackupEnabled** | Pointer to **bool** |  | [optional] 
**IsMonitoringEnabled** | Pointer to **bool** |  | [optional] 
**IsOpaEnabled** | Pointer to **bool** |  | [optional] 
**IsAutoUpgrade** | Pointer to **bool** |  | [optional] 
**HasKubeConfigFile** | Pointer to **bool** |  | [optional] 
**HasSelectedFlavors** | Pointer to **bool** |  | [optional] 
**Master** | Pointer to **int32** |  | [optional] 
**MasterReady** | Pointer to **int32** |  | [optional] 
**QuotaId** | Pointer to **int32** |  | [optional] 
**Bastion** | Pointer to **int32** |  | [optional] 
**AccessProfileRevision** | Pointer to **int32** |  | [optional] 
**OpaProfileRevision** | Pointer to **int32** |  | [optional] 
**CloudCredentialName** | Pointer to **string** |  | [optional] 
**CloudCredentialId** | Pointer to **int32** |  | [optional] 
**OrganizationName** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**PartnerId** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Health** | Pointer to **string** |  | [optional] 
**AccessIp** | Pointer to **string** |  | [optional] 
**TotalServersCount** | Pointer to **int32** |  | [optional] 
**CloudType** | Pointer to **string** |  | [optional] 
**KubesprayCurrentVersion** | Pointer to **string** |  | [optional] 
**KubesprayTargetVersion** | Pointer to **string** |  | [optional] 
**KubernetesCurrentVersion** | Pointer to **string** |  | [optional] 
**KubernetesTargetVersion** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**BoundUsers** | Pointer to [**[]UserDto**](UserDto.md) |  | [optional] 
**MonitoringCredential** | Pointer to [**MonitoringCredentialsListDto**](MonitoringCredentialsListDto.md) |  | [optional] 
**IsAutoscalingEnabled** | Pointer to **bool** |  | [optional] 
**Flavors** | Pointer to **[]string** |  | [optional] 
**AccessProfile** | Pointer to [**AccessProfilesForProjectListDto**](AccessProfilesForProjectListDto.md) |  | [optional] 
**KubernetesProfiles** | Pointer to [**KubernetesProfilesLisForPollerDto**](KubernetesProfilesLisForPollerDto.md) |  | [optional] 
**OpaProfile** | Pointer to [**OpaProfileListDto**](OpaProfileListDto.md) |  | [optional] 
**KubernetesAlerts** | Pointer to [**[]KubernetesAlertDto**](KubernetesAlertDto.md) |  | [optional] 
**S3BucketName** | Pointer to **string** |  | [optional] 
**S3AccessKeyId** | Pointer to **string** |  | [optional] 
**S3SecretKey** | Pointer to **string** |  | [optional] 
**S3Endpoint** | Pointer to **string** |  | [optional] 
**S3Region** | Pointer to **string** |  | [optional] 
**IsDeleteCluster** | Pointer to **bool** |  | [optional] 
**TaikunLBFlavor** | Pointer to **string** |  | [optional] 
**TaikunLBPrivateKey** | Pointer to **string** |  | [optional] 
**TaikunLBPublicKey** | Pointer to **string** |  | [optional] 
**RouterIdStartRange** | Pointer to **int32** |  | [optional] 
**RouterIdEndRange** | Pointer to **int32** |  | [optional] 
**TaikunPrivateSSHKey** | Pointer to **string** |  | [optional] 
**TaikunPublicSSHKey** | Pointer to **string** |  | [optional] 
**GoogleProjectId** | Pointer to **string** |  | [optional] 
**Cidr** | Pointer to **string** |  | [optional] 
**NetMask** | Pointer to **int32** |  | [optional] 
**PrivateIp** | Pointer to **string** |  | [optional] 
**PublicIp** | Pointer to **string** |  | [optional] 
**IsKubevapEnabled** | Pointer to **bool** |  | [optional] 
**TanzuReleaseVersion** | Pointer to **string** |  | [optional] 
**NfsDiskSize** | Pointer to **int32** |  | [optional] 
**KubevapEnabeledKubernetesVersions** | Pointer to **[]string** |  | [optional] 
**AwsProjectAZSubnets** | Pointer to [**[]AwsProjectAZSubnetDto**](AwsProjectAZSubnetDto.md) |  | [optional] 
**AvailabilityZones** | Pointer to **[]string** |  | [optional] 

## Methods

### NewProjectForListDto

`func NewProjectForListDto() *ProjectForListDto`

NewProjectForListDto instantiates a new ProjectForListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectForListDtoWithDefaults

`func NewProjectForListDtoWithDefaults() *ProjectForListDto`

NewProjectForListDtoWithDefaults instantiates a new ProjectForListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectForListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectForListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectForListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectForListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProjectForListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectForListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectForListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectForListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperation

`func (o *ProjectForListDto) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ProjectForListDto) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ProjectForListDto) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *ProjectForListDto) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetJobUrl

`func (o *ProjectForListDto) GetJobUrl() string`

GetJobUrl returns the JobUrl field if non-nil, zero value otherwise.

### GetJobUrlOk

`func (o *ProjectForListDto) GetJobUrlOk() (*string, bool)`

GetJobUrlOk returns a tuple with the JobUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobUrl

`func (o *ProjectForListDto) SetJobUrl(v string)`

SetJobUrl sets JobUrl field to given value.

### HasJobUrl

`func (o *ProjectForListDto) HasJobUrl() bool`

HasJobUrl returns a boolean if a field has been set.

### GetTopicName

`func (o *ProjectForListDto) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *ProjectForListDto) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *ProjectForListDto) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.

### HasTopicName

`func (o *ProjectForListDto) HasTopicName() bool`

HasTopicName returns a boolean if a field has been set.

### GetImageName

`func (o *ProjectForListDto) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *ProjectForListDto) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *ProjectForListDto) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *ProjectForListDto) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetToken

`func (o *ProjectForListDto) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ProjectForListDto) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ProjectForListDto) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ProjectForListDto) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetIsKubernetes

`func (o *ProjectForListDto) GetIsKubernetes() bool`

GetIsKubernetes returns the IsKubernetes field if non-nil, zero value otherwise.

### GetIsKubernetesOk

`func (o *ProjectForListDto) GetIsKubernetesOk() (*bool, bool)`

GetIsKubernetesOk returns a tuple with the IsKubernetes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKubernetes

`func (o *ProjectForListDto) SetIsKubernetes(v bool)`

SetIsKubernetes sets IsKubernetes field to given value.

### HasIsKubernetes

`func (o *ProjectForListDto) HasIsKubernetes() bool`

HasIsKubernetes returns a boolean if a field has been set.

### GetIsLocked

`func (o *ProjectForListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *ProjectForListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *ProjectForListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *ProjectForListDto) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetIsBackupEnabled

`func (o *ProjectForListDto) GetIsBackupEnabled() bool`

GetIsBackupEnabled returns the IsBackupEnabled field if non-nil, zero value otherwise.

### GetIsBackupEnabledOk

`func (o *ProjectForListDto) GetIsBackupEnabledOk() (*bool, bool)`

GetIsBackupEnabledOk returns a tuple with the IsBackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBackupEnabled

`func (o *ProjectForListDto) SetIsBackupEnabled(v bool)`

SetIsBackupEnabled sets IsBackupEnabled field to given value.

### HasIsBackupEnabled

`func (o *ProjectForListDto) HasIsBackupEnabled() bool`

HasIsBackupEnabled returns a boolean if a field has been set.

### GetIsMonitoringEnabled

`func (o *ProjectForListDto) GetIsMonitoringEnabled() bool`

GetIsMonitoringEnabled returns the IsMonitoringEnabled field if non-nil, zero value otherwise.

### GetIsMonitoringEnabledOk

`func (o *ProjectForListDto) GetIsMonitoringEnabledOk() (*bool, bool)`

GetIsMonitoringEnabledOk returns a tuple with the IsMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMonitoringEnabled

`func (o *ProjectForListDto) SetIsMonitoringEnabled(v bool)`

SetIsMonitoringEnabled sets IsMonitoringEnabled field to given value.

### HasIsMonitoringEnabled

`func (o *ProjectForListDto) HasIsMonitoringEnabled() bool`

HasIsMonitoringEnabled returns a boolean if a field has been set.

### GetIsOpaEnabled

`func (o *ProjectForListDto) GetIsOpaEnabled() bool`

GetIsOpaEnabled returns the IsOpaEnabled field if non-nil, zero value otherwise.

### GetIsOpaEnabledOk

`func (o *ProjectForListDto) GetIsOpaEnabledOk() (*bool, bool)`

GetIsOpaEnabledOk returns a tuple with the IsOpaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpaEnabled

`func (o *ProjectForListDto) SetIsOpaEnabled(v bool)`

SetIsOpaEnabled sets IsOpaEnabled field to given value.

### HasIsOpaEnabled

`func (o *ProjectForListDto) HasIsOpaEnabled() bool`

HasIsOpaEnabled returns a boolean if a field has been set.

### GetIsAutoUpgrade

`func (o *ProjectForListDto) GetIsAutoUpgrade() bool`

GetIsAutoUpgrade returns the IsAutoUpgrade field if non-nil, zero value otherwise.

### GetIsAutoUpgradeOk

`func (o *ProjectForListDto) GetIsAutoUpgradeOk() (*bool, bool)`

GetIsAutoUpgradeOk returns a tuple with the IsAutoUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoUpgrade

`func (o *ProjectForListDto) SetIsAutoUpgrade(v bool)`

SetIsAutoUpgrade sets IsAutoUpgrade field to given value.

### HasIsAutoUpgrade

`func (o *ProjectForListDto) HasIsAutoUpgrade() bool`

HasIsAutoUpgrade returns a boolean if a field has been set.

### GetHasKubeConfigFile

`func (o *ProjectForListDto) GetHasKubeConfigFile() bool`

GetHasKubeConfigFile returns the HasKubeConfigFile field if non-nil, zero value otherwise.

### GetHasKubeConfigFileOk

`func (o *ProjectForListDto) GetHasKubeConfigFileOk() (*bool, bool)`

GetHasKubeConfigFileOk returns a tuple with the HasKubeConfigFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasKubeConfigFile

`func (o *ProjectForListDto) SetHasKubeConfigFile(v bool)`

SetHasKubeConfigFile sets HasKubeConfigFile field to given value.

### HasHasKubeConfigFile

`func (o *ProjectForListDto) HasHasKubeConfigFile() bool`

HasHasKubeConfigFile returns a boolean if a field has been set.

### GetHasSelectedFlavors

`func (o *ProjectForListDto) GetHasSelectedFlavors() bool`

GetHasSelectedFlavors returns the HasSelectedFlavors field if non-nil, zero value otherwise.

### GetHasSelectedFlavorsOk

`func (o *ProjectForListDto) GetHasSelectedFlavorsOk() (*bool, bool)`

GetHasSelectedFlavorsOk returns a tuple with the HasSelectedFlavors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSelectedFlavors

`func (o *ProjectForListDto) SetHasSelectedFlavors(v bool)`

SetHasSelectedFlavors sets HasSelectedFlavors field to given value.

### HasHasSelectedFlavors

`func (o *ProjectForListDto) HasHasSelectedFlavors() bool`

HasHasSelectedFlavors returns a boolean if a field has been set.

### GetMaster

`func (o *ProjectForListDto) GetMaster() int32`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *ProjectForListDto) GetMasterOk() (*int32, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *ProjectForListDto) SetMaster(v int32)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *ProjectForListDto) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### GetMasterReady

`func (o *ProjectForListDto) GetMasterReady() int32`

GetMasterReady returns the MasterReady field if non-nil, zero value otherwise.

### GetMasterReadyOk

`func (o *ProjectForListDto) GetMasterReadyOk() (*int32, bool)`

GetMasterReadyOk returns a tuple with the MasterReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterReady

`func (o *ProjectForListDto) SetMasterReady(v int32)`

SetMasterReady sets MasterReady field to given value.

### HasMasterReady

`func (o *ProjectForListDto) HasMasterReady() bool`

HasMasterReady returns a boolean if a field has been set.

### GetQuotaId

`func (o *ProjectForListDto) GetQuotaId() int32`

GetQuotaId returns the QuotaId field if non-nil, zero value otherwise.

### GetQuotaIdOk

`func (o *ProjectForListDto) GetQuotaIdOk() (*int32, bool)`

GetQuotaIdOk returns a tuple with the QuotaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaId

`func (o *ProjectForListDto) SetQuotaId(v int32)`

SetQuotaId sets QuotaId field to given value.

### HasQuotaId

`func (o *ProjectForListDto) HasQuotaId() bool`

HasQuotaId returns a boolean if a field has been set.

### GetBastion

`func (o *ProjectForListDto) GetBastion() int32`

GetBastion returns the Bastion field if non-nil, zero value otherwise.

### GetBastionOk

`func (o *ProjectForListDto) GetBastionOk() (*int32, bool)`

GetBastionOk returns a tuple with the Bastion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBastion

`func (o *ProjectForListDto) SetBastion(v int32)`

SetBastion sets Bastion field to given value.

### HasBastion

`func (o *ProjectForListDto) HasBastion() bool`

HasBastion returns a boolean if a field has been set.

### GetAccessProfileRevision

`func (o *ProjectForListDto) GetAccessProfileRevision() int32`

GetAccessProfileRevision returns the AccessProfileRevision field if non-nil, zero value otherwise.

### GetAccessProfileRevisionOk

`func (o *ProjectForListDto) GetAccessProfileRevisionOk() (*int32, bool)`

GetAccessProfileRevisionOk returns a tuple with the AccessProfileRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfileRevision

`func (o *ProjectForListDto) SetAccessProfileRevision(v int32)`

SetAccessProfileRevision sets AccessProfileRevision field to given value.

### HasAccessProfileRevision

`func (o *ProjectForListDto) HasAccessProfileRevision() bool`

HasAccessProfileRevision returns a boolean if a field has been set.

### GetOpaProfileRevision

`func (o *ProjectForListDto) GetOpaProfileRevision() int32`

GetOpaProfileRevision returns the OpaProfileRevision field if non-nil, zero value otherwise.

### GetOpaProfileRevisionOk

`func (o *ProjectForListDto) GetOpaProfileRevisionOk() (*int32, bool)`

GetOpaProfileRevisionOk returns a tuple with the OpaProfileRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpaProfileRevision

`func (o *ProjectForListDto) SetOpaProfileRevision(v int32)`

SetOpaProfileRevision sets OpaProfileRevision field to given value.

### HasOpaProfileRevision

`func (o *ProjectForListDto) HasOpaProfileRevision() bool`

HasOpaProfileRevision returns a boolean if a field has been set.

### GetCloudCredentialName

`func (o *ProjectForListDto) GetCloudCredentialName() string`

GetCloudCredentialName returns the CloudCredentialName field if non-nil, zero value otherwise.

### GetCloudCredentialNameOk

`func (o *ProjectForListDto) GetCloudCredentialNameOk() (*string, bool)`

GetCloudCredentialNameOk returns a tuple with the CloudCredentialName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCredentialName

`func (o *ProjectForListDto) SetCloudCredentialName(v string)`

SetCloudCredentialName sets CloudCredentialName field to given value.

### HasCloudCredentialName

`func (o *ProjectForListDto) HasCloudCredentialName() bool`

HasCloudCredentialName returns a boolean if a field has been set.

### GetCloudCredentialId

`func (o *ProjectForListDto) GetCloudCredentialId() int32`

GetCloudCredentialId returns the CloudCredentialId field if non-nil, zero value otherwise.

### GetCloudCredentialIdOk

`func (o *ProjectForListDto) GetCloudCredentialIdOk() (*int32, bool)`

GetCloudCredentialIdOk returns a tuple with the CloudCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCredentialId

`func (o *ProjectForListDto) SetCloudCredentialId(v int32)`

SetCloudCredentialId sets CloudCredentialId field to given value.

### HasCloudCredentialId

`func (o *ProjectForListDto) HasCloudCredentialId() bool`

HasCloudCredentialId returns a boolean if a field has been set.

### GetOrganizationName

`func (o *ProjectForListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *ProjectForListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *ProjectForListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *ProjectForListDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetOrganizationId

`func (o *ProjectForListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ProjectForListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ProjectForListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ProjectForListDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetPartnerId

`func (o *ProjectForListDto) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *ProjectForListDto) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *ProjectForListDto) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *ProjectForListDto) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetStatus

`func (o *ProjectForListDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectForListDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectForListDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProjectForListDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetHealth

`func (o *ProjectForListDto) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ProjectForListDto) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ProjectForListDto) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *ProjectForListDto) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetAccessIp

`func (o *ProjectForListDto) GetAccessIp() string`

GetAccessIp returns the AccessIp field if non-nil, zero value otherwise.

### GetAccessIpOk

`func (o *ProjectForListDto) GetAccessIpOk() (*string, bool)`

GetAccessIpOk returns a tuple with the AccessIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessIp

`func (o *ProjectForListDto) SetAccessIp(v string)`

SetAccessIp sets AccessIp field to given value.

### HasAccessIp

`func (o *ProjectForListDto) HasAccessIp() bool`

HasAccessIp returns a boolean if a field has been set.

### GetTotalServersCount

`func (o *ProjectForListDto) GetTotalServersCount() int32`

GetTotalServersCount returns the TotalServersCount field if non-nil, zero value otherwise.

### GetTotalServersCountOk

`func (o *ProjectForListDto) GetTotalServersCountOk() (*int32, bool)`

GetTotalServersCountOk returns a tuple with the TotalServersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalServersCount

`func (o *ProjectForListDto) SetTotalServersCount(v int32)`

SetTotalServersCount sets TotalServersCount field to given value.

### HasTotalServersCount

`func (o *ProjectForListDto) HasTotalServersCount() bool`

HasTotalServersCount returns a boolean if a field has been set.

### GetCloudType

`func (o *ProjectForListDto) GetCloudType() string`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *ProjectForListDto) GetCloudTypeOk() (*string, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *ProjectForListDto) SetCloudType(v string)`

SetCloudType sets CloudType field to given value.

### HasCloudType

`func (o *ProjectForListDto) HasCloudType() bool`

HasCloudType returns a boolean if a field has been set.

### GetKubesprayCurrentVersion

`func (o *ProjectForListDto) GetKubesprayCurrentVersion() string`

GetKubesprayCurrentVersion returns the KubesprayCurrentVersion field if non-nil, zero value otherwise.

### GetKubesprayCurrentVersionOk

`func (o *ProjectForListDto) GetKubesprayCurrentVersionOk() (*string, bool)`

GetKubesprayCurrentVersionOk returns a tuple with the KubesprayCurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubesprayCurrentVersion

`func (o *ProjectForListDto) SetKubesprayCurrentVersion(v string)`

SetKubesprayCurrentVersion sets KubesprayCurrentVersion field to given value.

### HasKubesprayCurrentVersion

`func (o *ProjectForListDto) HasKubesprayCurrentVersion() bool`

HasKubesprayCurrentVersion returns a boolean if a field has been set.

### GetKubesprayTargetVersion

`func (o *ProjectForListDto) GetKubesprayTargetVersion() string`

GetKubesprayTargetVersion returns the KubesprayTargetVersion field if non-nil, zero value otherwise.

### GetKubesprayTargetVersionOk

`func (o *ProjectForListDto) GetKubesprayTargetVersionOk() (*string, bool)`

GetKubesprayTargetVersionOk returns a tuple with the KubesprayTargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubesprayTargetVersion

`func (o *ProjectForListDto) SetKubesprayTargetVersion(v string)`

SetKubesprayTargetVersion sets KubesprayTargetVersion field to given value.

### HasKubesprayTargetVersion

`func (o *ProjectForListDto) HasKubesprayTargetVersion() bool`

HasKubesprayTargetVersion returns a boolean if a field has been set.

### GetKubernetesCurrentVersion

`func (o *ProjectForListDto) GetKubernetesCurrentVersion() string`

GetKubernetesCurrentVersion returns the KubernetesCurrentVersion field if non-nil, zero value otherwise.

### GetKubernetesCurrentVersionOk

`func (o *ProjectForListDto) GetKubernetesCurrentVersionOk() (*string, bool)`

GetKubernetesCurrentVersionOk returns a tuple with the KubernetesCurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesCurrentVersion

`func (o *ProjectForListDto) SetKubernetesCurrentVersion(v string)`

SetKubernetesCurrentVersion sets KubernetesCurrentVersion field to given value.

### HasKubernetesCurrentVersion

`func (o *ProjectForListDto) HasKubernetesCurrentVersion() bool`

HasKubernetesCurrentVersion returns a boolean if a field has been set.

### GetKubernetesTargetVersion

`func (o *ProjectForListDto) GetKubernetesTargetVersion() string`

GetKubernetesTargetVersion returns the KubernetesTargetVersion field if non-nil, zero value otherwise.

### GetKubernetesTargetVersionOk

`func (o *ProjectForListDto) GetKubernetesTargetVersionOk() (*string, bool)`

GetKubernetesTargetVersionOk returns a tuple with the KubernetesTargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesTargetVersion

`func (o *ProjectForListDto) SetKubernetesTargetVersion(v string)`

SetKubernetesTargetVersion sets KubernetesTargetVersion field to given value.

### HasKubernetesTargetVersion

`func (o *ProjectForListDto) HasKubernetesTargetVersion() bool`

HasKubernetesTargetVersion returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProjectForListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectForListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectForListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProjectForListDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ProjectForListDto) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProjectForListDto) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProjectForListDto) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProjectForListDto) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetBoundUsers

`func (o *ProjectForListDto) GetBoundUsers() []UserDto`

GetBoundUsers returns the BoundUsers field if non-nil, zero value otherwise.

### GetBoundUsersOk

`func (o *ProjectForListDto) GetBoundUsersOk() (*[]UserDto, bool)`

GetBoundUsersOk returns a tuple with the BoundUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundUsers

`func (o *ProjectForListDto) SetBoundUsers(v []UserDto)`

SetBoundUsers sets BoundUsers field to given value.

### HasBoundUsers

`func (o *ProjectForListDto) HasBoundUsers() bool`

HasBoundUsers returns a boolean if a field has been set.

### GetMonitoringCredential

`func (o *ProjectForListDto) GetMonitoringCredential() MonitoringCredentialsListDto`

GetMonitoringCredential returns the MonitoringCredential field if non-nil, zero value otherwise.

### GetMonitoringCredentialOk

`func (o *ProjectForListDto) GetMonitoringCredentialOk() (*MonitoringCredentialsListDto, bool)`

GetMonitoringCredentialOk returns a tuple with the MonitoringCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringCredential

`func (o *ProjectForListDto) SetMonitoringCredential(v MonitoringCredentialsListDto)`

SetMonitoringCredential sets MonitoringCredential field to given value.

### HasMonitoringCredential

`func (o *ProjectForListDto) HasMonitoringCredential() bool`

HasMonitoringCredential returns a boolean if a field has been set.

### GetIsAutoscalingEnabled

`func (o *ProjectForListDto) GetIsAutoscalingEnabled() bool`

GetIsAutoscalingEnabled returns the IsAutoscalingEnabled field if non-nil, zero value otherwise.

### GetIsAutoscalingEnabledOk

`func (o *ProjectForListDto) GetIsAutoscalingEnabledOk() (*bool, bool)`

GetIsAutoscalingEnabledOk returns a tuple with the IsAutoscalingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoscalingEnabled

`func (o *ProjectForListDto) SetIsAutoscalingEnabled(v bool)`

SetIsAutoscalingEnabled sets IsAutoscalingEnabled field to given value.

### HasIsAutoscalingEnabled

`func (o *ProjectForListDto) HasIsAutoscalingEnabled() bool`

HasIsAutoscalingEnabled returns a boolean if a field has been set.

### GetFlavors

`func (o *ProjectForListDto) GetFlavors() []string`

GetFlavors returns the Flavors field if non-nil, zero value otherwise.

### GetFlavorsOk

`func (o *ProjectForListDto) GetFlavorsOk() (*[]string, bool)`

GetFlavorsOk returns a tuple with the Flavors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavors

`func (o *ProjectForListDto) SetFlavors(v []string)`

SetFlavors sets Flavors field to given value.

### HasFlavors

`func (o *ProjectForListDto) HasFlavors() bool`

HasFlavors returns a boolean if a field has been set.

### GetAccessProfile

`func (o *ProjectForListDto) GetAccessProfile() AccessProfilesForProjectListDto`

GetAccessProfile returns the AccessProfile field if non-nil, zero value otherwise.

### GetAccessProfileOk

`func (o *ProjectForListDto) GetAccessProfileOk() (*AccessProfilesForProjectListDto, bool)`

GetAccessProfileOk returns a tuple with the AccessProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfile

`func (o *ProjectForListDto) SetAccessProfile(v AccessProfilesForProjectListDto)`

SetAccessProfile sets AccessProfile field to given value.

### HasAccessProfile

`func (o *ProjectForListDto) HasAccessProfile() bool`

HasAccessProfile returns a boolean if a field has been set.

### GetKubernetesProfiles

`func (o *ProjectForListDto) GetKubernetesProfiles() KubernetesProfilesLisForPollerDto`

GetKubernetesProfiles returns the KubernetesProfiles field if non-nil, zero value otherwise.

### GetKubernetesProfilesOk

`func (o *ProjectForListDto) GetKubernetesProfilesOk() (*KubernetesProfilesLisForPollerDto, bool)`

GetKubernetesProfilesOk returns a tuple with the KubernetesProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesProfiles

`func (o *ProjectForListDto) SetKubernetesProfiles(v KubernetesProfilesLisForPollerDto)`

SetKubernetesProfiles sets KubernetesProfiles field to given value.

### HasKubernetesProfiles

`func (o *ProjectForListDto) HasKubernetesProfiles() bool`

HasKubernetesProfiles returns a boolean if a field has been set.

### GetOpaProfile

`func (o *ProjectForListDto) GetOpaProfile() OpaProfileListDto`

GetOpaProfile returns the OpaProfile field if non-nil, zero value otherwise.

### GetOpaProfileOk

`func (o *ProjectForListDto) GetOpaProfileOk() (*OpaProfileListDto, bool)`

GetOpaProfileOk returns a tuple with the OpaProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpaProfile

`func (o *ProjectForListDto) SetOpaProfile(v OpaProfileListDto)`

SetOpaProfile sets OpaProfile field to given value.

### HasOpaProfile

`func (o *ProjectForListDto) HasOpaProfile() bool`

HasOpaProfile returns a boolean if a field has been set.

### GetKubernetesAlerts

`func (o *ProjectForListDto) GetKubernetesAlerts() []KubernetesAlertDto`

GetKubernetesAlerts returns the KubernetesAlerts field if non-nil, zero value otherwise.

### GetKubernetesAlertsOk

`func (o *ProjectForListDto) GetKubernetesAlertsOk() (*[]KubernetesAlertDto, bool)`

GetKubernetesAlertsOk returns a tuple with the KubernetesAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesAlerts

`func (o *ProjectForListDto) SetKubernetesAlerts(v []KubernetesAlertDto)`

SetKubernetesAlerts sets KubernetesAlerts field to given value.

### HasKubernetesAlerts

`func (o *ProjectForListDto) HasKubernetesAlerts() bool`

HasKubernetesAlerts returns a boolean if a field has been set.

### GetS3BucketName

`func (o *ProjectForListDto) GetS3BucketName() string`

GetS3BucketName returns the S3BucketName field if non-nil, zero value otherwise.

### GetS3BucketNameOk

`func (o *ProjectForListDto) GetS3BucketNameOk() (*string, bool)`

GetS3BucketNameOk returns a tuple with the S3BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3BucketName

`func (o *ProjectForListDto) SetS3BucketName(v string)`

SetS3BucketName sets S3BucketName field to given value.

### HasS3BucketName

`func (o *ProjectForListDto) HasS3BucketName() bool`

HasS3BucketName returns a boolean if a field has been set.

### GetS3AccessKeyId

`func (o *ProjectForListDto) GetS3AccessKeyId() string`

GetS3AccessKeyId returns the S3AccessKeyId field if non-nil, zero value otherwise.

### GetS3AccessKeyIdOk

`func (o *ProjectForListDto) GetS3AccessKeyIdOk() (*string, bool)`

GetS3AccessKeyIdOk returns a tuple with the S3AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3AccessKeyId

`func (o *ProjectForListDto) SetS3AccessKeyId(v string)`

SetS3AccessKeyId sets S3AccessKeyId field to given value.

### HasS3AccessKeyId

`func (o *ProjectForListDto) HasS3AccessKeyId() bool`

HasS3AccessKeyId returns a boolean if a field has been set.

### GetS3SecretKey

`func (o *ProjectForListDto) GetS3SecretKey() string`

GetS3SecretKey returns the S3SecretKey field if non-nil, zero value otherwise.

### GetS3SecretKeyOk

`func (o *ProjectForListDto) GetS3SecretKeyOk() (*string, bool)`

GetS3SecretKeyOk returns a tuple with the S3SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3SecretKey

`func (o *ProjectForListDto) SetS3SecretKey(v string)`

SetS3SecretKey sets S3SecretKey field to given value.

### HasS3SecretKey

`func (o *ProjectForListDto) HasS3SecretKey() bool`

HasS3SecretKey returns a boolean if a field has been set.

### GetS3Endpoint

`func (o *ProjectForListDto) GetS3Endpoint() string`

GetS3Endpoint returns the S3Endpoint field if non-nil, zero value otherwise.

### GetS3EndpointOk

`func (o *ProjectForListDto) GetS3EndpointOk() (*string, bool)`

GetS3EndpointOk returns a tuple with the S3Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Endpoint

`func (o *ProjectForListDto) SetS3Endpoint(v string)`

SetS3Endpoint sets S3Endpoint field to given value.

### HasS3Endpoint

`func (o *ProjectForListDto) HasS3Endpoint() bool`

HasS3Endpoint returns a boolean if a field has been set.

### GetS3Region

`func (o *ProjectForListDto) GetS3Region() string`

GetS3Region returns the S3Region field if non-nil, zero value otherwise.

### GetS3RegionOk

`func (o *ProjectForListDto) GetS3RegionOk() (*string, bool)`

GetS3RegionOk returns a tuple with the S3Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Region

`func (o *ProjectForListDto) SetS3Region(v string)`

SetS3Region sets S3Region field to given value.

### HasS3Region

`func (o *ProjectForListDto) HasS3Region() bool`

HasS3Region returns a boolean if a field has been set.

### GetIsDeleteCluster

`func (o *ProjectForListDto) GetIsDeleteCluster() bool`

GetIsDeleteCluster returns the IsDeleteCluster field if non-nil, zero value otherwise.

### GetIsDeleteClusterOk

`func (o *ProjectForListDto) GetIsDeleteClusterOk() (*bool, bool)`

GetIsDeleteClusterOk returns a tuple with the IsDeleteCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleteCluster

`func (o *ProjectForListDto) SetIsDeleteCluster(v bool)`

SetIsDeleteCluster sets IsDeleteCluster field to given value.

### HasIsDeleteCluster

`func (o *ProjectForListDto) HasIsDeleteCluster() bool`

HasIsDeleteCluster returns a boolean if a field has been set.

### GetTaikunLBFlavor

`func (o *ProjectForListDto) GetTaikunLBFlavor() string`

GetTaikunLBFlavor returns the TaikunLBFlavor field if non-nil, zero value otherwise.

### GetTaikunLBFlavorOk

`func (o *ProjectForListDto) GetTaikunLBFlavorOk() (*string, bool)`

GetTaikunLBFlavorOk returns a tuple with the TaikunLBFlavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaikunLBFlavor

`func (o *ProjectForListDto) SetTaikunLBFlavor(v string)`

SetTaikunLBFlavor sets TaikunLBFlavor field to given value.

### HasTaikunLBFlavor

`func (o *ProjectForListDto) HasTaikunLBFlavor() bool`

HasTaikunLBFlavor returns a boolean if a field has been set.

### GetTaikunLBPrivateKey

`func (o *ProjectForListDto) GetTaikunLBPrivateKey() string`

GetTaikunLBPrivateKey returns the TaikunLBPrivateKey field if non-nil, zero value otherwise.

### GetTaikunLBPrivateKeyOk

`func (o *ProjectForListDto) GetTaikunLBPrivateKeyOk() (*string, bool)`

GetTaikunLBPrivateKeyOk returns a tuple with the TaikunLBPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaikunLBPrivateKey

`func (o *ProjectForListDto) SetTaikunLBPrivateKey(v string)`

SetTaikunLBPrivateKey sets TaikunLBPrivateKey field to given value.

### HasTaikunLBPrivateKey

`func (o *ProjectForListDto) HasTaikunLBPrivateKey() bool`

HasTaikunLBPrivateKey returns a boolean if a field has been set.

### GetTaikunLBPublicKey

`func (o *ProjectForListDto) GetTaikunLBPublicKey() string`

GetTaikunLBPublicKey returns the TaikunLBPublicKey field if non-nil, zero value otherwise.

### GetTaikunLBPublicKeyOk

`func (o *ProjectForListDto) GetTaikunLBPublicKeyOk() (*string, bool)`

GetTaikunLBPublicKeyOk returns a tuple with the TaikunLBPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaikunLBPublicKey

`func (o *ProjectForListDto) SetTaikunLBPublicKey(v string)`

SetTaikunLBPublicKey sets TaikunLBPublicKey field to given value.

### HasTaikunLBPublicKey

`func (o *ProjectForListDto) HasTaikunLBPublicKey() bool`

HasTaikunLBPublicKey returns a boolean if a field has been set.

### GetRouterIdStartRange

`func (o *ProjectForListDto) GetRouterIdStartRange() int32`

GetRouterIdStartRange returns the RouterIdStartRange field if non-nil, zero value otherwise.

### GetRouterIdStartRangeOk

`func (o *ProjectForListDto) GetRouterIdStartRangeOk() (*int32, bool)`

GetRouterIdStartRangeOk returns a tuple with the RouterIdStartRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterIdStartRange

`func (o *ProjectForListDto) SetRouterIdStartRange(v int32)`

SetRouterIdStartRange sets RouterIdStartRange field to given value.

### HasRouterIdStartRange

`func (o *ProjectForListDto) HasRouterIdStartRange() bool`

HasRouterIdStartRange returns a boolean if a field has been set.

### GetRouterIdEndRange

`func (o *ProjectForListDto) GetRouterIdEndRange() int32`

GetRouterIdEndRange returns the RouterIdEndRange field if non-nil, zero value otherwise.

### GetRouterIdEndRangeOk

`func (o *ProjectForListDto) GetRouterIdEndRangeOk() (*int32, bool)`

GetRouterIdEndRangeOk returns a tuple with the RouterIdEndRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterIdEndRange

`func (o *ProjectForListDto) SetRouterIdEndRange(v int32)`

SetRouterIdEndRange sets RouterIdEndRange field to given value.

### HasRouterIdEndRange

`func (o *ProjectForListDto) HasRouterIdEndRange() bool`

HasRouterIdEndRange returns a boolean if a field has been set.

### GetTaikunPrivateSSHKey

`func (o *ProjectForListDto) GetTaikunPrivateSSHKey() string`

GetTaikunPrivateSSHKey returns the TaikunPrivateSSHKey field if non-nil, zero value otherwise.

### GetTaikunPrivateSSHKeyOk

`func (o *ProjectForListDto) GetTaikunPrivateSSHKeyOk() (*string, bool)`

GetTaikunPrivateSSHKeyOk returns a tuple with the TaikunPrivateSSHKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaikunPrivateSSHKey

`func (o *ProjectForListDto) SetTaikunPrivateSSHKey(v string)`

SetTaikunPrivateSSHKey sets TaikunPrivateSSHKey field to given value.

### HasTaikunPrivateSSHKey

`func (o *ProjectForListDto) HasTaikunPrivateSSHKey() bool`

HasTaikunPrivateSSHKey returns a boolean if a field has been set.

### GetTaikunPublicSSHKey

`func (o *ProjectForListDto) GetTaikunPublicSSHKey() string`

GetTaikunPublicSSHKey returns the TaikunPublicSSHKey field if non-nil, zero value otherwise.

### GetTaikunPublicSSHKeyOk

`func (o *ProjectForListDto) GetTaikunPublicSSHKeyOk() (*string, bool)`

GetTaikunPublicSSHKeyOk returns a tuple with the TaikunPublicSSHKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaikunPublicSSHKey

`func (o *ProjectForListDto) SetTaikunPublicSSHKey(v string)`

SetTaikunPublicSSHKey sets TaikunPublicSSHKey field to given value.

### HasTaikunPublicSSHKey

`func (o *ProjectForListDto) HasTaikunPublicSSHKey() bool`

HasTaikunPublicSSHKey returns a boolean if a field has been set.

### GetGoogleProjectId

`func (o *ProjectForListDto) GetGoogleProjectId() string`

GetGoogleProjectId returns the GoogleProjectId field if non-nil, zero value otherwise.

### GetGoogleProjectIdOk

`func (o *ProjectForListDto) GetGoogleProjectIdOk() (*string, bool)`

GetGoogleProjectIdOk returns a tuple with the GoogleProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleProjectId

`func (o *ProjectForListDto) SetGoogleProjectId(v string)`

SetGoogleProjectId sets GoogleProjectId field to given value.

### HasGoogleProjectId

`func (o *ProjectForListDto) HasGoogleProjectId() bool`

HasGoogleProjectId returns a boolean if a field has been set.

### GetCidr

`func (o *ProjectForListDto) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *ProjectForListDto) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *ProjectForListDto) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *ProjectForListDto) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetNetMask

`func (o *ProjectForListDto) GetNetMask() int32`

GetNetMask returns the NetMask field if non-nil, zero value otherwise.

### GetNetMaskOk

`func (o *ProjectForListDto) GetNetMaskOk() (*int32, bool)`

GetNetMaskOk returns a tuple with the NetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetMask

`func (o *ProjectForListDto) SetNetMask(v int32)`

SetNetMask sets NetMask field to given value.

### HasNetMask

`func (o *ProjectForListDto) HasNetMask() bool`

HasNetMask returns a boolean if a field has been set.

### GetPrivateIp

`func (o *ProjectForListDto) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *ProjectForListDto) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *ProjectForListDto) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *ProjectForListDto) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetPublicIp

`func (o *ProjectForListDto) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *ProjectForListDto) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *ProjectForListDto) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *ProjectForListDto) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetIsKubevapEnabled

`func (o *ProjectForListDto) GetIsKubevapEnabled() bool`

GetIsKubevapEnabled returns the IsKubevapEnabled field if non-nil, zero value otherwise.

### GetIsKubevapEnabledOk

`func (o *ProjectForListDto) GetIsKubevapEnabledOk() (*bool, bool)`

GetIsKubevapEnabledOk returns a tuple with the IsKubevapEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKubevapEnabled

`func (o *ProjectForListDto) SetIsKubevapEnabled(v bool)`

SetIsKubevapEnabled sets IsKubevapEnabled field to given value.

### HasIsKubevapEnabled

`func (o *ProjectForListDto) HasIsKubevapEnabled() bool`

HasIsKubevapEnabled returns a boolean if a field has been set.

### GetTanzuReleaseVersion

`func (o *ProjectForListDto) GetTanzuReleaseVersion() string`

GetTanzuReleaseVersion returns the TanzuReleaseVersion field if non-nil, zero value otherwise.

### GetTanzuReleaseVersionOk

`func (o *ProjectForListDto) GetTanzuReleaseVersionOk() (*string, bool)`

GetTanzuReleaseVersionOk returns a tuple with the TanzuReleaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTanzuReleaseVersion

`func (o *ProjectForListDto) SetTanzuReleaseVersion(v string)`

SetTanzuReleaseVersion sets TanzuReleaseVersion field to given value.

### HasTanzuReleaseVersion

`func (o *ProjectForListDto) HasTanzuReleaseVersion() bool`

HasTanzuReleaseVersion returns a boolean if a field has been set.

### GetNfsDiskSize

`func (o *ProjectForListDto) GetNfsDiskSize() int32`

GetNfsDiskSize returns the NfsDiskSize field if non-nil, zero value otherwise.

### GetNfsDiskSizeOk

`func (o *ProjectForListDto) GetNfsDiskSizeOk() (*int32, bool)`

GetNfsDiskSizeOk returns a tuple with the NfsDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsDiskSize

`func (o *ProjectForListDto) SetNfsDiskSize(v int32)`

SetNfsDiskSize sets NfsDiskSize field to given value.

### HasNfsDiskSize

`func (o *ProjectForListDto) HasNfsDiskSize() bool`

HasNfsDiskSize returns a boolean if a field has been set.

### GetKubevapEnabeledKubernetesVersions

`func (o *ProjectForListDto) GetKubevapEnabeledKubernetesVersions() []string`

GetKubevapEnabeledKubernetesVersions returns the KubevapEnabeledKubernetesVersions field if non-nil, zero value otherwise.

### GetKubevapEnabeledKubernetesVersionsOk

`func (o *ProjectForListDto) GetKubevapEnabeledKubernetesVersionsOk() (*[]string, bool)`

GetKubevapEnabeledKubernetesVersionsOk returns a tuple with the KubevapEnabeledKubernetesVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubevapEnabeledKubernetesVersions

`func (o *ProjectForListDto) SetKubevapEnabeledKubernetesVersions(v []string)`

SetKubevapEnabeledKubernetesVersions sets KubevapEnabeledKubernetesVersions field to given value.

### HasKubevapEnabeledKubernetesVersions

`func (o *ProjectForListDto) HasKubevapEnabeledKubernetesVersions() bool`

HasKubevapEnabeledKubernetesVersions returns a boolean if a field has been set.

### GetAwsProjectAZSubnets

`func (o *ProjectForListDto) GetAwsProjectAZSubnets() []AwsProjectAZSubnetDto`

GetAwsProjectAZSubnets returns the AwsProjectAZSubnets field if non-nil, zero value otherwise.

### GetAwsProjectAZSubnetsOk

`func (o *ProjectForListDto) GetAwsProjectAZSubnetsOk() (*[]AwsProjectAZSubnetDto, bool)`

GetAwsProjectAZSubnetsOk returns a tuple with the AwsProjectAZSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsProjectAZSubnets

`func (o *ProjectForListDto) SetAwsProjectAZSubnets(v []AwsProjectAZSubnetDto)`

SetAwsProjectAZSubnets sets AwsProjectAZSubnets field to given value.

### HasAwsProjectAZSubnets

`func (o *ProjectForListDto) HasAwsProjectAZSubnets() bool`

HasAwsProjectAZSubnets returns a boolean if a field has been set.

### GetAvailabilityZones

`func (o *ProjectForListDto) GetAvailabilityZones() []string`

GetAvailabilityZones returns the AvailabilityZones field if non-nil, zero value otherwise.

### GetAvailabilityZonesOk

`func (o *ProjectForListDto) GetAvailabilityZonesOk() (*[]string, bool)`

GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZones

`func (o *ProjectForListDto) SetAvailabilityZones(v []string)`

SetAvailabilityZones sets AvailabilityZones field to given value.

### HasAvailabilityZones

`func (o *ProjectForListDto) HasAvailabilityZones() bool`

HasAvailabilityZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


