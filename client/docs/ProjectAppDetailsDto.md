# ProjectAppDetailsDto

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
**AppRepoName** | Pointer to **string** |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**Values** | Pointer to **string** |  | [optional] 
**AutoSync** | Pointer to **bool** |  | [optional] 
**ReleaseNotes** | Pointer to **string** |  | [optional] 
**ProjectName** | Pointer to **string** |  | [optional] 
**HelmResult** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**HasJsonSchema** | Pointer to **bool** |  | [optional] 
**ProjectAppParams** | Pointer to [**[]ProjectAppParamDto**](ProjectAppParamDto.md) |  | [optional] 

## Methods

### NewProjectAppDetailsDto

`func NewProjectAppDetailsDto() *ProjectAppDetailsDto`

NewProjectAppDetailsDto instantiates a new ProjectAppDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectAppDetailsDtoWithDefaults

`func NewProjectAppDetailsDtoWithDefaults() *ProjectAppDetailsDto`

NewProjectAppDetailsDtoWithDefaults instantiates a new ProjectAppDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectAppDetailsDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectAppDetailsDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectAppDetailsDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectAppDetailsDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProjectAppDetailsDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectAppDetailsDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectAppDetailsDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectAppDetailsDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *ProjectAppDetailsDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ProjectAppDetailsDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ProjectAppDetailsDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ProjectAppDetailsDto) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetStatus

`func (o *ProjectAppDetailsDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectAppDetailsDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectAppDetailsDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProjectAppDetailsDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *ProjectAppDetailsDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ProjectAppDetailsDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ProjectAppDetailsDto) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ProjectAppDetailsDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCatalogId

`func (o *ProjectAppDetailsDto) GetCatalogId() int32`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *ProjectAppDetailsDto) GetCatalogIdOk() (*int32, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *ProjectAppDetailsDto) SetCatalogId(v int32)`

SetCatalogId sets CatalogId field to given value.

### HasCatalogId

`func (o *ProjectAppDetailsDto) HasCatalogId() bool`

HasCatalogId returns a boolean if a field has been set.

### GetCatalogName

`func (o *ProjectAppDetailsDto) GetCatalogName() string`

GetCatalogName returns the CatalogName field if non-nil, zero value otherwise.

### GetCatalogNameOk

`func (o *ProjectAppDetailsDto) GetCatalogNameOk() (*string, bool)`

GetCatalogNameOk returns a tuple with the CatalogName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogName

`func (o *ProjectAppDetailsDto) SetCatalogName(v string)`

SetCatalogName sets CatalogName field to given value.

### HasCatalogName

`func (o *ProjectAppDetailsDto) HasCatalogName() bool`

HasCatalogName returns a boolean if a field has been set.

### GetCatalogAppName

`func (o *ProjectAppDetailsDto) GetCatalogAppName() string`

GetCatalogAppName returns the CatalogAppName field if non-nil, zero value otherwise.

### GetCatalogAppNameOk

`func (o *ProjectAppDetailsDto) GetCatalogAppNameOk() (*string, bool)`

GetCatalogAppNameOk returns a tuple with the CatalogAppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogAppName

`func (o *ProjectAppDetailsDto) SetCatalogAppName(v string)`

SetCatalogAppName sets CatalogAppName field to given value.

### HasCatalogAppName

`func (o *ProjectAppDetailsDto) HasCatalogAppName() bool`

HasCatalogAppName returns a boolean if a field has been set.

### GetAppRepoName

`func (o *ProjectAppDetailsDto) GetAppRepoName() string`

GetAppRepoName returns the AppRepoName field if non-nil, zero value otherwise.

### GetAppRepoNameOk

`func (o *ProjectAppDetailsDto) GetAppRepoNameOk() (*string, bool)`

GetAppRepoNameOk returns a tuple with the AppRepoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRepoName

`func (o *ProjectAppDetailsDto) SetAppRepoName(v string)`

SetAppRepoName sets AppRepoName field to given value.

### HasAppRepoName

`func (o *ProjectAppDetailsDto) HasAppRepoName() bool`

HasAppRepoName returns a boolean if a field has been set.

### GetLogo

`func (o *ProjectAppDetailsDto) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *ProjectAppDetailsDto) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *ProjectAppDetailsDto) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *ProjectAppDetailsDto) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetValues

`func (o *ProjectAppDetailsDto) GetValues() string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ProjectAppDetailsDto) GetValuesOk() (*string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ProjectAppDetailsDto) SetValues(v string)`

SetValues sets Values field to given value.

### HasValues

`func (o *ProjectAppDetailsDto) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetAutoSync

`func (o *ProjectAppDetailsDto) GetAutoSync() bool`

GetAutoSync returns the AutoSync field if non-nil, zero value otherwise.

### GetAutoSyncOk

`func (o *ProjectAppDetailsDto) GetAutoSyncOk() (*bool, bool)`

GetAutoSyncOk returns a tuple with the AutoSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSync

`func (o *ProjectAppDetailsDto) SetAutoSync(v bool)`

SetAutoSync sets AutoSync field to given value.

### HasAutoSync

`func (o *ProjectAppDetailsDto) HasAutoSync() bool`

HasAutoSync returns a boolean if a field has been set.

### GetReleaseNotes

`func (o *ProjectAppDetailsDto) GetReleaseNotes() string`

GetReleaseNotes returns the ReleaseNotes field if non-nil, zero value otherwise.

### GetReleaseNotesOk

`func (o *ProjectAppDetailsDto) GetReleaseNotesOk() (*string, bool)`

GetReleaseNotesOk returns a tuple with the ReleaseNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseNotes

`func (o *ProjectAppDetailsDto) SetReleaseNotes(v string)`

SetReleaseNotes sets ReleaseNotes field to given value.

### HasReleaseNotes

`func (o *ProjectAppDetailsDto) HasReleaseNotes() bool`

HasReleaseNotes returns a boolean if a field has been set.

### GetProjectName

`func (o *ProjectAppDetailsDto) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ProjectAppDetailsDto) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ProjectAppDetailsDto) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *ProjectAppDetailsDto) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetHelmResult

`func (o *ProjectAppDetailsDto) GetHelmResult() string`

GetHelmResult returns the HelmResult field if non-nil, zero value otherwise.

### GetHelmResultOk

`func (o *ProjectAppDetailsDto) GetHelmResultOk() (*string, bool)`

GetHelmResultOk returns a tuple with the HelmResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmResult

`func (o *ProjectAppDetailsDto) SetHelmResult(v string)`

SetHelmResult sets HelmResult field to given value.

### HasHelmResult

`func (o *ProjectAppDetailsDto) HasHelmResult() bool`

HasHelmResult returns a boolean if a field has been set.

### GetProjectId

`func (o *ProjectAppDetailsDto) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectAppDetailsDto) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectAppDetailsDto) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ProjectAppDetailsDto) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetHasJsonSchema

`func (o *ProjectAppDetailsDto) GetHasJsonSchema() bool`

GetHasJsonSchema returns the HasJsonSchema field if non-nil, zero value otherwise.

### GetHasJsonSchemaOk

`func (o *ProjectAppDetailsDto) GetHasJsonSchemaOk() (*bool, bool)`

GetHasJsonSchemaOk returns a tuple with the HasJsonSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasJsonSchema

`func (o *ProjectAppDetailsDto) SetHasJsonSchema(v bool)`

SetHasJsonSchema sets HasJsonSchema field to given value.

### HasHasJsonSchema

`func (o *ProjectAppDetailsDto) HasHasJsonSchema() bool`

HasHasJsonSchema returns a boolean if a field has been set.

### GetProjectAppParams

`func (o *ProjectAppDetailsDto) GetProjectAppParams() []ProjectAppParamDto`

GetProjectAppParams returns the ProjectAppParams field if non-nil, zero value otherwise.

### GetProjectAppParamsOk

`func (o *ProjectAppDetailsDto) GetProjectAppParamsOk() (*[]ProjectAppParamDto, bool)`

GetProjectAppParamsOk returns a tuple with the ProjectAppParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectAppParams

`func (o *ProjectAppDetailsDto) SetProjectAppParams(v []ProjectAppParamDto)`

SetProjectAppParams sets ProjectAppParams field to given value.

### HasProjectAppParams

`func (o *ProjectAppDetailsDto) HasProjectAppParams() bool`

HasProjectAppParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


