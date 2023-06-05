# ProjectTemplateListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**MonitoringEnabled** | Pointer to **bool** |  | [optional] 
**BackupEnabled** | Pointer to **bool** |  | [optional] 
**AllowFullSpotKubernetes** | Pointer to **bool** |  | [optional] 
**AllowSpotVms** | Pointer to **bool** |  | [optional] 
**AllowSpotWorkers** | Pointer to **bool** |  | [optional] 
**KubernetesVersion** | Pointer to **string** |  | [optional] 
**OrganizationName** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 

## Methods

### NewProjectTemplateListDto

`func NewProjectTemplateListDto() *ProjectTemplateListDto`

NewProjectTemplateListDto instantiates a new ProjectTemplateListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectTemplateListDtoWithDefaults

`func NewProjectTemplateListDtoWithDefaults() *ProjectTemplateListDto`

NewProjectTemplateListDtoWithDefaults instantiates a new ProjectTemplateListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectTemplateListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectTemplateListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectTemplateListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectTemplateListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProjectTemplateListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectTemplateListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectTemplateListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectTemplateListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMonitoringEnabled

`func (o *ProjectTemplateListDto) GetMonitoringEnabled() bool`

GetMonitoringEnabled returns the MonitoringEnabled field if non-nil, zero value otherwise.

### GetMonitoringEnabledOk

`func (o *ProjectTemplateListDto) GetMonitoringEnabledOk() (*bool, bool)`

GetMonitoringEnabledOk returns a tuple with the MonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringEnabled

`func (o *ProjectTemplateListDto) SetMonitoringEnabled(v bool)`

SetMonitoringEnabled sets MonitoringEnabled field to given value.

### HasMonitoringEnabled

`func (o *ProjectTemplateListDto) HasMonitoringEnabled() bool`

HasMonitoringEnabled returns a boolean if a field has been set.

### GetBackupEnabled

`func (o *ProjectTemplateListDto) GetBackupEnabled() bool`

GetBackupEnabled returns the BackupEnabled field if non-nil, zero value otherwise.

### GetBackupEnabledOk

`func (o *ProjectTemplateListDto) GetBackupEnabledOk() (*bool, bool)`

GetBackupEnabledOk returns a tuple with the BackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupEnabled

`func (o *ProjectTemplateListDto) SetBackupEnabled(v bool)`

SetBackupEnabled sets BackupEnabled field to given value.

### HasBackupEnabled

`func (o *ProjectTemplateListDto) HasBackupEnabled() bool`

HasBackupEnabled returns a boolean if a field has been set.

### GetAllowFullSpotKubernetes

`func (o *ProjectTemplateListDto) GetAllowFullSpotKubernetes() bool`

GetAllowFullSpotKubernetes returns the AllowFullSpotKubernetes field if non-nil, zero value otherwise.

### GetAllowFullSpotKubernetesOk

`func (o *ProjectTemplateListDto) GetAllowFullSpotKubernetesOk() (*bool, bool)`

GetAllowFullSpotKubernetesOk returns a tuple with the AllowFullSpotKubernetes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFullSpotKubernetes

`func (o *ProjectTemplateListDto) SetAllowFullSpotKubernetes(v bool)`

SetAllowFullSpotKubernetes sets AllowFullSpotKubernetes field to given value.

### HasAllowFullSpotKubernetes

`func (o *ProjectTemplateListDto) HasAllowFullSpotKubernetes() bool`

HasAllowFullSpotKubernetes returns a boolean if a field has been set.

### GetAllowSpotVms

`func (o *ProjectTemplateListDto) GetAllowSpotVms() bool`

GetAllowSpotVms returns the AllowSpotVms field if non-nil, zero value otherwise.

### GetAllowSpotVmsOk

`func (o *ProjectTemplateListDto) GetAllowSpotVmsOk() (*bool, bool)`

GetAllowSpotVmsOk returns a tuple with the AllowSpotVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSpotVms

`func (o *ProjectTemplateListDto) SetAllowSpotVms(v bool)`

SetAllowSpotVms sets AllowSpotVms field to given value.

### HasAllowSpotVms

`func (o *ProjectTemplateListDto) HasAllowSpotVms() bool`

HasAllowSpotVms returns a boolean if a field has been set.

### GetAllowSpotWorkers

`func (o *ProjectTemplateListDto) GetAllowSpotWorkers() bool`

GetAllowSpotWorkers returns the AllowSpotWorkers field if non-nil, zero value otherwise.

### GetAllowSpotWorkersOk

`func (o *ProjectTemplateListDto) GetAllowSpotWorkersOk() (*bool, bool)`

GetAllowSpotWorkersOk returns a tuple with the AllowSpotWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSpotWorkers

`func (o *ProjectTemplateListDto) SetAllowSpotWorkers(v bool)`

SetAllowSpotWorkers sets AllowSpotWorkers field to given value.

### HasAllowSpotWorkers

`func (o *ProjectTemplateListDto) HasAllowSpotWorkers() bool`

HasAllowSpotWorkers returns a boolean if a field has been set.

### GetKubernetesVersion

`func (o *ProjectTemplateListDto) GetKubernetesVersion() string`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *ProjectTemplateListDto) GetKubernetesVersionOk() (*string, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *ProjectTemplateListDto) SetKubernetesVersion(v string)`

SetKubernetesVersion sets KubernetesVersion field to given value.

### HasKubernetesVersion

`func (o *ProjectTemplateListDto) HasKubernetesVersion() bool`

HasKubernetesVersion returns a boolean if a field has been set.

### GetOrganizationName

`func (o *ProjectTemplateListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *ProjectTemplateListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *ProjectTemplateListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *ProjectTemplateListDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetOrganizationId

`func (o *ProjectTemplateListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ProjectTemplateListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ProjectTemplateListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ProjectTemplateListDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


