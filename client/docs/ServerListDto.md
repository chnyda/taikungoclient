# ServerListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ProjectName** | Pointer to **string** |  | [optional] 
**OrganizationName** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**DiskSize** | Pointer to **int64** |  | [optional] 
**KubernetesHealth** | Pointer to **string** |  | [optional] 
**GoogleMachineType** | Pointer to **string** |  | [optional] 
**TanzuFlavor** | Pointer to **string** |  | [optional] 
**ProxmoxFlavor** | Pointer to **string** |  | [optional] 
**Cpu** | Pointer to **int32** |  | [optional] 
**Ram** | Pointer to **int64** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**OpenstackFlavor** | Pointer to **string** |  | [optional] 
**AwsInstanceType** | Pointer to **string** |  | [optional] 
**AzureVmSize** | Pointer to **string** |  | [optional] 
**CloudType** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**LastModified** | Pointer to **string** |  | [optional] 
**LastModifiedBy** | Pointer to **string** |  | [optional] 
**SpotPrice** | Pointer to **string** |  | [optional] 
**SpotInstance** | Pointer to **bool** |  | [optional] 
**ShutOff** | Pointer to **bool** |  | [optional] 
**AutoscalingGroup** | Pointer to **string** |  | [optional] 
**ProviderID** | Pointer to **string** |  | [optional] 
**AwsHostName** | Pointer to **string** |  | [optional] 
**AvailabilityZone** | Pointer to **string** |  | [optional] 
**Hypervisor** | Pointer to **string** |  | [optional] 
**ProxmoxRole** | Pointer to **string** |  | [optional] 
**ProxmoxNFSDiskSize** | Pointer to **int32** |  | [optional] 
**ActionButtons** | Pointer to [**ServerActionButtonVisibilityDto**](ServerActionButtonVisibilityDto.md) |  | [optional] 
**KubernetesNodeLabels** | Pointer to [**[]KubernetesNodeLabelsDto**](KubernetesNodeLabelsDto.md) |  | [optional] 

## Methods

### NewServerListDto

`func NewServerListDto() *ServerListDto`

NewServerListDto instantiates a new ServerListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerListDtoWithDefaults

`func NewServerListDtoWithDefaults() *ServerListDto`

NewServerListDtoWithDefaults instantiates a new ServerListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ServerListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ServerListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServerListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjectName

`func (o *ServerListDto) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ServerListDto) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ServerListDto) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *ServerListDto) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetOrganizationName

`func (o *ServerListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *ServerListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *ServerListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *ServerListDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetOrganizationId

`func (o *ServerListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ServerListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ServerListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ServerListDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetProjectId

`func (o *ServerListDto) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ServerListDto) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ServerListDto) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ServerListDto) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetIpAddress

`func (o *ServerListDto) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ServerListDto) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ServerListDto) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ServerListDto) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetDiskSize

`func (o *ServerListDto) GetDiskSize() int64`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *ServerListDto) GetDiskSizeOk() (*int64, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *ServerListDto) SetDiskSize(v int64)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *ServerListDto) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### GetKubernetesHealth

`func (o *ServerListDto) GetKubernetesHealth() string`

GetKubernetesHealth returns the KubernetesHealth field if non-nil, zero value otherwise.

### GetKubernetesHealthOk

`func (o *ServerListDto) GetKubernetesHealthOk() (*string, bool)`

GetKubernetesHealthOk returns a tuple with the KubernetesHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesHealth

`func (o *ServerListDto) SetKubernetesHealth(v string)`

SetKubernetesHealth sets KubernetesHealth field to given value.

### HasKubernetesHealth

`func (o *ServerListDto) HasKubernetesHealth() bool`

HasKubernetesHealth returns a boolean if a field has been set.

### GetGoogleMachineType

`func (o *ServerListDto) GetGoogleMachineType() string`

GetGoogleMachineType returns the GoogleMachineType field if non-nil, zero value otherwise.

### GetGoogleMachineTypeOk

`func (o *ServerListDto) GetGoogleMachineTypeOk() (*string, bool)`

GetGoogleMachineTypeOk returns a tuple with the GoogleMachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleMachineType

`func (o *ServerListDto) SetGoogleMachineType(v string)`

SetGoogleMachineType sets GoogleMachineType field to given value.

### HasGoogleMachineType

`func (o *ServerListDto) HasGoogleMachineType() bool`

HasGoogleMachineType returns a boolean if a field has been set.

### GetTanzuFlavor

`func (o *ServerListDto) GetTanzuFlavor() string`

GetTanzuFlavor returns the TanzuFlavor field if non-nil, zero value otherwise.

### GetTanzuFlavorOk

`func (o *ServerListDto) GetTanzuFlavorOk() (*string, bool)`

GetTanzuFlavorOk returns a tuple with the TanzuFlavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTanzuFlavor

`func (o *ServerListDto) SetTanzuFlavor(v string)`

SetTanzuFlavor sets TanzuFlavor field to given value.

### HasTanzuFlavor

`func (o *ServerListDto) HasTanzuFlavor() bool`

HasTanzuFlavor returns a boolean if a field has been set.

### GetProxmoxFlavor

`func (o *ServerListDto) GetProxmoxFlavor() string`

GetProxmoxFlavor returns the ProxmoxFlavor field if non-nil, zero value otherwise.

### GetProxmoxFlavorOk

`func (o *ServerListDto) GetProxmoxFlavorOk() (*string, bool)`

GetProxmoxFlavorOk returns a tuple with the ProxmoxFlavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxmoxFlavor

`func (o *ServerListDto) SetProxmoxFlavor(v string)`

SetProxmoxFlavor sets ProxmoxFlavor field to given value.

### HasProxmoxFlavor

`func (o *ServerListDto) HasProxmoxFlavor() bool`

HasProxmoxFlavor returns a boolean if a field has been set.

### GetCpu

`func (o *ServerListDto) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ServerListDto) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ServerListDto) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ServerListDto) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetRam

`func (o *ServerListDto) GetRam() int64`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *ServerListDto) GetRamOk() (*int64, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *ServerListDto) SetRam(v int64)`

SetRam sets Ram field to given value.

### HasRam

`func (o *ServerListDto) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetRole

`func (o *ServerListDto) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ServerListDto) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ServerListDto) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ServerListDto) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStatus

`func (o *ServerListDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServerListDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServerListDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServerListDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ServerListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServerListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServerListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServerListDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetOpenstackFlavor

`func (o *ServerListDto) GetOpenstackFlavor() string`

GetOpenstackFlavor returns the OpenstackFlavor field if non-nil, zero value otherwise.

### GetOpenstackFlavorOk

`func (o *ServerListDto) GetOpenstackFlavorOk() (*string, bool)`

GetOpenstackFlavorOk returns a tuple with the OpenstackFlavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenstackFlavor

`func (o *ServerListDto) SetOpenstackFlavor(v string)`

SetOpenstackFlavor sets OpenstackFlavor field to given value.

### HasOpenstackFlavor

`func (o *ServerListDto) HasOpenstackFlavor() bool`

HasOpenstackFlavor returns a boolean if a field has been set.

### GetAwsInstanceType

`func (o *ServerListDto) GetAwsInstanceType() string`

GetAwsInstanceType returns the AwsInstanceType field if non-nil, zero value otherwise.

### GetAwsInstanceTypeOk

`func (o *ServerListDto) GetAwsInstanceTypeOk() (*string, bool)`

GetAwsInstanceTypeOk returns a tuple with the AwsInstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsInstanceType

`func (o *ServerListDto) SetAwsInstanceType(v string)`

SetAwsInstanceType sets AwsInstanceType field to given value.

### HasAwsInstanceType

`func (o *ServerListDto) HasAwsInstanceType() bool`

HasAwsInstanceType returns a boolean if a field has been set.

### GetAzureVmSize

`func (o *ServerListDto) GetAzureVmSize() string`

GetAzureVmSize returns the AzureVmSize field if non-nil, zero value otherwise.

### GetAzureVmSizeOk

`func (o *ServerListDto) GetAzureVmSizeOk() (*string, bool)`

GetAzureVmSizeOk returns a tuple with the AzureVmSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureVmSize

`func (o *ServerListDto) SetAzureVmSize(v string)`

SetAzureVmSize sets AzureVmSize field to given value.

### HasAzureVmSize

`func (o *ServerListDto) HasAzureVmSize() bool`

HasAzureVmSize returns a boolean if a field has been set.

### GetCloudType

`func (o *ServerListDto) GetCloudType() string`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *ServerListDto) GetCloudTypeOk() (*string, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *ServerListDto) SetCloudType(v string)`

SetCloudType sets CloudType field to given value.

### HasCloudType

`func (o *ServerListDto) HasCloudType() bool`

HasCloudType returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ServerListDto) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ServerListDto) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ServerListDto) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ServerListDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetLastModified

`func (o *ServerListDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ServerListDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ServerListDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *ServerListDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *ServerListDto) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *ServerListDto) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *ServerListDto) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *ServerListDto) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetSpotPrice

`func (o *ServerListDto) GetSpotPrice() string`

GetSpotPrice returns the SpotPrice field if non-nil, zero value otherwise.

### GetSpotPriceOk

`func (o *ServerListDto) GetSpotPriceOk() (*string, bool)`

GetSpotPriceOk returns a tuple with the SpotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotPrice

`func (o *ServerListDto) SetSpotPrice(v string)`

SetSpotPrice sets SpotPrice field to given value.

### HasSpotPrice

`func (o *ServerListDto) HasSpotPrice() bool`

HasSpotPrice returns a boolean if a field has been set.

### GetSpotInstance

`func (o *ServerListDto) GetSpotInstance() bool`

GetSpotInstance returns the SpotInstance field if non-nil, zero value otherwise.

### GetSpotInstanceOk

`func (o *ServerListDto) GetSpotInstanceOk() (*bool, bool)`

GetSpotInstanceOk returns a tuple with the SpotInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotInstance

`func (o *ServerListDto) SetSpotInstance(v bool)`

SetSpotInstance sets SpotInstance field to given value.

### HasSpotInstance

`func (o *ServerListDto) HasSpotInstance() bool`

HasSpotInstance returns a boolean if a field has been set.

### GetShutOff

`func (o *ServerListDto) GetShutOff() bool`

GetShutOff returns the ShutOff field if non-nil, zero value otherwise.

### GetShutOffOk

`func (o *ServerListDto) GetShutOffOk() (*bool, bool)`

GetShutOffOk returns a tuple with the ShutOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutOff

`func (o *ServerListDto) SetShutOff(v bool)`

SetShutOff sets ShutOff field to given value.

### HasShutOff

`func (o *ServerListDto) HasShutOff() bool`

HasShutOff returns a boolean if a field has been set.

### GetAutoscalingGroup

`func (o *ServerListDto) GetAutoscalingGroup() string`

GetAutoscalingGroup returns the AutoscalingGroup field if non-nil, zero value otherwise.

### GetAutoscalingGroupOk

`func (o *ServerListDto) GetAutoscalingGroupOk() (*string, bool)`

GetAutoscalingGroupOk returns a tuple with the AutoscalingGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalingGroup

`func (o *ServerListDto) SetAutoscalingGroup(v string)`

SetAutoscalingGroup sets AutoscalingGroup field to given value.

### HasAutoscalingGroup

`func (o *ServerListDto) HasAutoscalingGroup() bool`

HasAutoscalingGroup returns a boolean if a field has been set.

### GetProviderID

`func (o *ServerListDto) GetProviderID() string`

GetProviderID returns the ProviderID field if non-nil, zero value otherwise.

### GetProviderIDOk

`func (o *ServerListDto) GetProviderIDOk() (*string, bool)`

GetProviderIDOk returns a tuple with the ProviderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderID

`func (o *ServerListDto) SetProviderID(v string)`

SetProviderID sets ProviderID field to given value.

### HasProviderID

`func (o *ServerListDto) HasProviderID() bool`

HasProviderID returns a boolean if a field has been set.

### GetAwsHostName

`func (o *ServerListDto) GetAwsHostName() string`

GetAwsHostName returns the AwsHostName field if non-nil, zero value otherwise.

### GetAwsHostNameOk

`func (o *ServerListDto) GetAwsHostNameOk() (*string, bool)`

GetAwsHostNameOk returns a tuple with the AwsHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsHostName

`func (o *ServerListDto) SetAwsHostName(v string)`

SetAwsHostName sets AwsHostName field to given value.

### HasAwsHostName

`func (o *ServerListDto) HasAwsHostName() bool`

HasAwsHostName returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *ServerListDto) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *ServerListDto) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *ServerListDto) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *ServerListDto) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetHypervisor

`func (o *ServerListDto) GetHypervisor() string`

GetHypervisor returns the Hypervisor field if non-nil, zero value otherwise.

### GetHypervisorOk

`func (o *ServerListDto) GetHypervisorOk() (*string, bool)`

GetHypervisorOk returns a tuple with the Hypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisor

`func (o *ServerListDto) SetHypervisor(v string)`

SetHypervisor sets Hypervisor field to given value.

### HasHypervisor

`func (o *ServerListDto) HasHypervisor() bool`

HasHypervisor returns a boolean if a field has been set.

### GetProxmoxRole

`func (o *ServerListDto) GetProxmoxRole() string`

GetProxmoxRole returns the ProxmoxRole field if non-nil, zero value otherwise.

### GetProxmoxRoleOk

`func (o *ServerListDto) GetProxmoxRoleOk() (*string, bool)`

GetProxmoxRoleOk returns a tuple with the ProxmoxRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxmoxRole

`func (o *ServerListDto) SetProxmoxRole(v string)`

SetProxmoxRole sets ProxmoxRole field to given value.

### HasProxmoxRole

`func (o *ServerListDto) HasProxmoxRole() bool`

HasProxmoxRole returns a boolean if a field has been set.

### GetProxmoxNFSDiskSize

`func (o *ServerListDto) GetProxmoxNFSDiskSize() int32`

GetProxmoxNFSDiskSize returns the ProxmoxNFSDiskSize field if non-nil, zero value otherwise.

### GetProxmoxNFSDiskSizeOk

`func (o *ServerListDto) GetProxmoxNFSDiskSizeOk() (*int32, bool)`

GetProxmoxNFSDiskSizeOk returns a tuple with the ProxmoxNFSDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxmoxNFSDiskSize

`func (o *ServerListDto) SetProxmoxNFSDiskSize(v int32)`

SetProxmoxNFSDiskSize sets ProxmoxNFSDiskSize field to given value.

### HasProxmoxNFSDiskSize

`func (o *ServerListDto) HasProxmoxNFSDiskSize() bool`

HasProxmoxNFSDiskSize returns a boolean if a field has been set.

### GetActionButtons

`func (o *ServerListDto) GetActionButtons() ServerActionButtonVisibilityDto`

GetActionButtons returns the ActionButtons field if non-nil, zero value otherwise.

### GetActionButtonsOk

`func (o *ServerListDto) GetActionButtonsOk() (*ServerActionButtonVisibilityDto, bool)`

GetActionButtonsOk returns a tuple with the ActionButtons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionButtons

`func (o *ServerListDto) SetActionButtons(v ServerActionButtonVisibilityDto)`

SetActionButtons sets ActionButtons field to given value.

### HasActionButtons

`func (o *ServerListDto) HasActionButtons() bool`

HasActionButtons returns a boolean if a field has been set.

### GetKubernetesNodeLabels

`func (o *ServerListDto) GetKubernetesNodeLabels() []KubernetesNodeLabelsDto`

GetKubernetesNodeLabels returns the KubernetesNodeLabels field if non-nil, zero value otherwise.

### GetKubernetesNodeLabelsOk

`func (o *ServerListDto) GetKubernetesNodeLabelsOk() (*[]KubernetesNodeLabelsDto, bool)`

GetKubernetesNodeLabelsOk returns a tuple with the KubernetesNodeLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesNodeLabels

`func (o *ServerListDto) SetKubernetesNodeLabels(v []KubernetesNodeLabelsDto)`

SetKubernetesNodeLabels sets KubernetesNodeLabels field to given value.

### HasKubernetesNodeLabels

`func (o *ServerListDto) HasKubernetesNodeLabels() bool`

HasKubernetesNodeLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


