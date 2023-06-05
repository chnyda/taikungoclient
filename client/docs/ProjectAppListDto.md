# ProjectAppListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**CatalogId** | Pointer to **int32** |  | [optional] 
**CatalogName** | Pointer to **string** |  | [optional] 
**CatalogAppName** | Pointer to **string** |  | [optional] 
**CatalogAppId** | Pointer to **int32** |  | [optional] 
**AppRepoName** | Pointer to **string** |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**AutoSync** | Pointer to **bool** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**LastModified** | Pointer to **string** |  | [optional] 
**LastModifiedBy** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**ProjectName** | Pointer to **string** |  | [optional] 

## Methods

### NewProjectAppListDto

`func NewProjectAppListDto() *ProjectAppListDto`

NewProjectAppListDto instantiates a new ProjectAppListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectAppListDtoWithDefaults

`func NewProjectAppListDtoWithDefaults() *ProjectAppListDto`

NewProjectAppListDtoWithDefaults instantiates a new ProjectAppListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectAppListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectAppListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectAppListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectAppListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProjectAppListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectAppListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectAppListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectAppListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *ProjectAppListDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ProjectAppListDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ProjectAppListDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ProjectAppListDto) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetStatus

`func (o *ProjectAppListDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectAppListDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectAppListDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProjectAppListDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *ProjectAppListDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ProjectAppListDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ProjectAppListDto) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ProjectAppListDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCatalogId

`func (o *ProjectAppListDto) GetCatalogId() int32`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *ProjectAppListDto) GetCatalogIdOk() (*int32, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *ProjectAppListDto) SetCatalogId(v int32)`

SetCatalogId sets CatalogId field to given value.

### HasCatalogId

`func (o *ProjectAppListDto) HasCatalogId() bool`

HasCatalogId returns a boolean if a field has been set.

### GetCatalogName

`func (o *ProjectAppListDto) GetCatalogName() string`

GetCatalogName returns the CatalogName field if non-nil, zero value otherwise.

### GetCatalogNameOk

`func (o *ProjectAppListDto) GetCatalogNameOk() (*string, bool)`

GetCatalogNameOk returns a tuple with the CatalogName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogName

`func (o *ProjectAppListDto) SetCatalogName(v string)`

SetCatalogName sets CatalogName field to given value.

### HasCatalogName

`func (o *ProjectAppListDto) HasCatalogName() bool`

HasCatalogName returns a boolean if a field has been set.

### GetCatalogAppName

`func (o *ProjectAppListDto) GetCatalogAppName() string`

GetCatalogAppName returns the CatalogAppName field if non-nil, zero value otherwise.

### GetCatalogAppNameOk

`func (o *ProjectAppListDto) GetCatalogAppNameOk() (*string, bool)`

GetCatalogAppNameOk returns a tuple with the CatalogAppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogAppName

`func (o *ProjectAppListDto) SetCatalogAppName(v string)`

SetCatalogAppName sets CatalogAppName field to given value.

### HasCatalogAppName

`func (o *ProjectAppListDto) HasCatalogAppName() bool`

HasCatalogAppName returns a boolean if a field has been set.

### GetCatalogAppId

`func (o *ProjectAppListDto) GetCatalogAppId() int32`

GetCatalogAppId returns the CatalogAppId field if non-nil, zero value otherwise.

### GetCatalogAppIdOk

`func (o *ProjectAppListDto) GetCatalogAppIdOk() (*int32, bool)`

GetCatalogAppIdOk returns a tuple with the CatalogAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogAppId

`func (o *ProjectAppListDto) SetCatalogAppId(v int32)`

SetCatalogAppId sets CatalogAppId field to given value.

### HasCatalogAppId

`func (o *ProjectAppListDto) HasCatalogAppId() bool`

HasCatalogAppId returns a boolean if a field has been set.

### GetAppRepoName

`func (o *ProjectAppListDto) GetAppRepoName() string`

GetAppRepoName returns the AppRepoName field if non-nil, zero value otherwise.

### GetAppRepoNameOk

`func (o *ProjectAppListDto) GetAppRepoNameOk() (*string, bool)`

GetAppRepoNameOk returns a tuple with the AppRepoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRepoName

`func (o *ProjectAppListDto) SetAppRepoName(v string)`

SetAppRepoName sets AppRepoName field to given value.

### HasAppRepoName

`func (o *ProjectAppListDto) HasAppRepoName() bool`

HasAppRepoName returns a boolean if a field has been set.

### GetLogo

`func (o *ProjectAppListDto) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *ProjectAppListDto) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *ProjectAppListDto) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *ProjectAppListDto) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetAutoSync

`func (o *ProjectAppListDto) GetAutoSync() bool`

GetAutoSync returns the AutoSync field if non-nil, zero value otherwise.

### GetAutoSyncOk

`func (o *ProjectAppListDto) GetAutoSyncOk() (*bool, bool)`

GetAutoSyncOk returns a tuple with the AutoSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSync

`func (o *ProjectAppListDto) SetAutoSync(v bool)`

SetAutoSync sets AutoSync field to given value.

### HasAutoSync

`func (o *ProjectAppListDto) HasAutoSync() bool`

HasAutoSync returns a boolean if a field has been set.

### GetCreated

`func (o *ProjectAppListDto) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ProjectAppListDto) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ProjectAppListDto) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ProjectAppListDto) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ProjectAppListDto) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ProjectAppListDto) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ProjectAppListDto) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ProjectAppListDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetLastModified

`func (o *ProjectAppListDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ProjectAppListDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ProjectAppListDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *ProjectAppListDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *ProjectAppListDto) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *ProjectAppListDto) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *ProjectAppListDto) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *ProjectAppListDto) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetProjectId

`func (o *ProjectAppListDto) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectAppListDto) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectAppListDto) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ProjectAppListDto) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProjectName

`func (o *ProjectAppListDto) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ProjectAppListDto) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ProjectAppListDto) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *ProjectAppListDto) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


