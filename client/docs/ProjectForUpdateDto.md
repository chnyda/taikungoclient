# ProjectForUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KubeConfig** | Pointer to **string** |  | [optional] 
**KubesprayCurrentVersion** | Pointer to **string** |  | [optional] 
**AccessIp** | Pointer to **string** |  | [optional] 
**ImageName** | Pointer to **string** |  | [optional] 
**MonitoringIp** | Pointer to **string** |  | [optional] 
**TanzuReleaseCurrentVersion** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**ProjectStatus**](ProjectStatus.md) |  | [optional] 
**Health** | Pointer to [**ProjectHealth**](ProjectHealth.md) |  | [optional] 
**IsBackupEnabled** | Pointer to **bool** |  | [optional] 
**IsMonitoringEnabled** | Pointer to **bool** |  | [optional] 
**IsOpaEnabled** | Pointer to **bool** |  | [optional] 
**IsAutoUpgrade** | Pointer to **bool** |  | [optional] 
**AppEnabled** | Pointer to **bool** |  | [optional] 
**IsKubevapEnabled** | Pointer to **bool** |  | [optional] 
**KubernetesCurrentVersion** | Pointer to **string** |  | [optional] 
**FailureReason** | Pointer to **string** |  | [optional] 

## Methods

### NewProjectForUpdateDto

`func NewProjectForUpdateDto() *ProjectForUpdateDto`

NewProjectForUpdateDto instantiates a new ProjectForUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectForUpdateDtoWithDefaults

`func NewProjectForUpdateDtoWithDefaults() *ProjectForUpdateDto`

NewProjectForUpdateDtoWithDefaults instantiates a new ProjectForUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKubeConfig

`func (o *ProjectForUpdateDto) GetKubeConfig() string`

GetKubeConfig returns the KubeConfig field if non-nil, zero value otherwise.

### GetKubeConfigOk

`func (o *ProjectForUpdateDto) GetKubeConfigOk() (*string, bool)`

GetKubeConfigOk returns a tuple with the KubeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeConfig

`func (o *ProjectForUpdateDto) SetKubeConfig(v string)`

SetKubeConfig sets KubeConfig field to given value.

### HasKubeConfig

`func (o *ProjectForUpdateDto) HasKubeConfig() bool`

HasKubeConfig returns a boolean if a field has been set.

### GetKubesprayCurrentVersion

`func (o *ProjectForUpdateDto) GetKubesprayCurrentVersion() string`

GetKubesprayCurrentVersion returns the KubesprayCurrentVersion field if non-nil, zero value otherwise.

### GetKubesprayCurrentVersionOk

`func (o *ProjectForUpdateDto) GetKubesprayCurrentVersionOk() (*string, bool)`

GetKubesprayCurrentVersionOk returns a tuple with the KubesprayCurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubesprayCurrentVersion

`func (o *ProjectForUpdateDto) SetKubesprayCurrentVersion(v string)`

SetKubesprayCurrentVersion sets KubesprayCurrentVersion field to given value.

### HasKubesprayCurrentVersion

`func (o *ProjectForUpdateDto) HasKubesprayCurrentVersion() bool`

HasKubesprayCurrentVersion returns a boolean if a field has been set.

### GetAccessIp

`func (o *ProjectForUpdateDto) GetAccessIp() string`

GetAccessIp returns the AccessIp field if non-nil, zero value otherwise.

### GetAccessIpOk

`func (o *ProjectForUpdateDto) GetAccessIpOk() (*string, bool)`

GetAccessIpOk returns a tuple with the AccessIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessIp

`func (o *ProjectForUpdateDto) SetAccessIp(v string)`

SetAccessIp sets AccessIp field to given value.

### HasAccessIp

`func (o *ProjectForUpdateDto) HasAccessIp() bool`

HasAccessIp returns a boolean if a field has been set.

### GetImageName

`func (o *ProjectForUpdateDto) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *ProjectForUpdateDto) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *ProjectForUpdateDto) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *ProjectForUpdateDto) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetMonitoringIp

`func (o *ProjectForUpdateDto) GetMonitoringIp() string`

GetMonitoringIp returns the MonitoringIp field if non-nil, zero value otherwise.

### GetMonitoringIpOk

`func (o *ProjectForUpdateDto) GetMonitoringIpOk() (*string, bool)`

GetMonitoringIpOk returns a tuple with the MonitoringIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringIp

`func (o *ProjectForUpdateDto) SetMonitoringIp(v string)`

SetMonitoringIp sets MonitoringIp field to given value.

### HasMonitoringIp

`func (o *ProjectForUpdateDto) HasMonitoringIp() bool`

HasMonitoringIp returns a boolean if a field has been set.

### GetTanzuReleaseCurrentVersion

`func (o *ProjectForUpdateDto) GetTanzuReleaseCurrentVersion() string`

GetTanzuReleaseCurrentVersion returns the TanzuReleaseCurrentVersion field if non-nil, zero value otherwise.

### GetTanzuReleaseCurrentVersionOk

`func (o *ProjectForUpdateDto) GetTanzuReleaseCurrentVersionOk() (*string, bool)`

GetTanzuReleaseCurrentVersionOk returns a tuple with the TanzuReleaseCurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTanzuReleaseCurrentVersion

`func (o *ProjectForUpdateDto) SetTanzuReleaseCurrentVersion(v string)`

SetTanzuReleaseCurrentVersion sets TanzuReleaseCurrentVersion field to given value.

### HasTanzuReleaseCurrentVersion

`func (o *ProjectForUpdateDto) HasTanzuReleaseCurrentVersion() bool`

HasTanzuReleaseCurrentVersion returns a boolean if a field has been set.

### GetStatus

`func (o *ProjectForUpdateDto) GetStatus() ProjectStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectForUpdateDto) GetStatusOk() (*ProjectStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectForUpdateDto) SetStatus(v ProjectStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProjectForUpdateDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetHealth

`func (o *ProjectForUpdateDto) GetHealth() ProjectHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ProjectForUpdateDto) GetHealthOk() (*ProjectHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ProjectForUpdateDto) SetHealth(v ProjectHealth)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *ProjectForUpdateDto) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetIsBackupEnabled

`func (o *ProjectForUpdateDto) GetIsBackupEnabled() bool`

GetIsBackupEnabled returns the IsBackupEnabled field if non-nil, zero value otherwise.

### GetIsBackupEnabledOk

`func (o *ProjectForUpdateDto) GetIsBackupEnabledOk() (*bool, bool)`

GetIsBackupEnabledOk returns a tuple with the IsBackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBackupEnabled

`func (o *ProjectForUpdateDto) SetIsBackupEnabled(v bool)`

SetIsBackupEnabled sets IsBackupEnabled field to given value.

### HasIsBackupEnabled

`func (o *ProjectForUpdateDto) HasIsBackupEnabled() bool`

HasIsBackupEnabled returns a boolean if a field has been set.

### GetIsMonitoringEnabled

`func (o *ProjectForUpdateDto) GetIsMonitoringEnabled() bool`

GetIsMonitoringEnabled returns the IsMonitoringEnabled field if non-nil, zero value otherwise.

### GetIsMonitoringEnabledOk

`func (o *ProjectForUpdateDto) GetIsMonitoringEnabledOk() (*bool, bool)`

GetIsMonitoringEnabledOk returns a tuple with the IsMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMonitoringEnabled

`func (o *ProjectForUpdateDto) SetIsMonitoringEnabled(v bool)`

SetIsMonitoringEnabled sets IsMonitoringEnabled field to given value.

### HasIsMonitoringEnabled

`func (o *ProjectForUpdateDto) HasIsMonitoringEnabled() bool`

HasIsMonitoringEnabled returns a boolean if a field has been set.

### GetIsOpaEnabled

`func (o *ProjectForUpdateDto) GetIsOpaEnabled() bool`

GetIsOpaEnabled returns the IsOpaEnabled field if non-nil, zero value otherwise.

### GetIsOpaEnabledOk

`func (o *ProjectForUpdateDto) GetIsOpaEnabledOk() (*bool, bool)`

GetIsOpaEnabledOk returns a tuple with the IsOpaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpaEnabled

`func (o *ProjectForUpdateDto) SetIsOpaEnabled(v bool)`

SetIsOpaEnabled sets IsOpaEnabled field to given value.

### HasIsOpaEnabled

`func (o *ProjectForUpdateDto) HasIsOpaEnabled() bool`

HasIsOpaEnabled returns a boolean if a field has been set.

### GetIsAutoUpgrade

`func (o *ProjectForUpdateDto) GetIsAutoUpgrade() bool`

GetIsAutoUpgrade returns the IsAutoUpgrade field if non-nil, zero value otherwise.

### GetIsAutoUpgradeOk

`func (o *ProjectForUpdateDto) GetIsAutoUpgradeOk() (*bool, bool)`

GetIsAutoUpgradeOk returns a tuple with the IsAutoUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoUpgrade

`func (o *ProjectForUpdateDto) SetIsAutoUpgrade(v bool)`

SetIsAutoUpgrade sets IsAutoUpgrade field to given value.

### HasIsAutoUpgrade

`func (o *ProjectForUpdateDto) HasIsAutoUpgrade() bool`

HasIsAutoUpgrade returns a boolean if a field has been set.

### GetAppEnabled

`func (o *ProjectForUpdateDto) GetAppEnabled() bool`

GetAppEnabled returns the AppEnabled field if non-nil, zero value otherwise.

### GetAppEnabledOk

`func (o *ProjectForUpdateDto) GetAppEnabledOk() (*bool, bool)`

GetAppEnabledOk returns a tuple with the AppEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEnabled

`func (o *ProjectForUpdateDto) SetAppEnabled(v bool)`

SetAppEnabled sets AppEnabled field to given value.

### HasAppEnabled

`func (o *ProjectForUpdateDto) HasAppEnabled() bool`

HasAppEnabled returns a boolean if a field has been set.

### GetIsKubevapEnabled

`func (o *ProjectForUpdateDto) GetIsKubevapEnabled() bool`

GetIsKubevapEnabled returns the IsKubevapEnabled field if non-nil, zero value otherwise.

### GetIsKubevapEnabledOk

`func (o *ProjectForUpdateDto) GetIsKubevapEnabledOk() (*bool, bool)`

GetIsKubevapEnabledOk returns a tuple with the IsKubevapEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKubevapEnabled

`func (o *ProjectForUpdateDto) SetIsKubevapEnabled(v bool)`

SetIsKubevapEnabled sets IsKubevapEnabled field to given value.

### HasIsKubevapEnabled

`func (o *ProjectForUpdateDto) HasIsKubevapEnabled() bool`

HasIsKubevapEnabled returns a boolean if a field has been set.

### GetKubernetesCurrentVersion

`func (o *ProjectForUpdateDto) GetKubernetesCurrentVersion() string`

GetKubernetesCurrentVersion returns the KubernetesCurrentVersion field if non-nil, zero value otherwise.

### GetKubernetesCurrentVersionOk

`func (o *ProjectForUpdateDto) GetKubernetesCurrentVersionOk() (*string, bool)`

GetKubernetesCurrentVersionOk returns a tuple with the KubernetesCurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesCurrentVersion

`func (o *ProjectForUpdateDto) SetKubernetesCurrentVersion(v string)`

SetKubernetesCurrentVersion sets KubernetesCurrentVersion field to given value.

### HasKubernetesCurrentVersion

`func (o *ProjectForUpdateDto) HasKubernetesCurrentVersion() bool`

HasKubernetesCurrentVersion returns a boolean if a field has been set.

### GetFailureReason

`func (o *ProjectForUpdateDto) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *ProjectForUpdateDto) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *ProjectForUpdateDto) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *ProjectForUpdateDto) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


