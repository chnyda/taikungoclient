# CreateProjectAppCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Namespace** | **string** |  | 
**ProjectId** | **int32** |  | 
**CatalogAppId** | **int32** |  | 
**ExtraValues** | Pointer to **string** |  | [optional] 
**AutoSync** | Pointer to **bool** |  | [optional] 
**Parameters** | Pointer to [**[]ProjectAppParamsDto**](ProjectAppParamsDto.md) |  | [optional] 

## Methods

### NewCreateProjectAppCommand

`func NewCreateProjectAppCommand(name string, namespace string, projectId int32, catalogAppId int32, ) *CreateProjectAppCommand`

NewCreateProjectAppCommand instantiates a new CreateProjectAppCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProjectAppCommandWithDefaults

`func NewCreateProjectAppCommandWithDefaults() *CreateProjectAppCommand`

NewCreateProjectAppCommandWithDefaults instantiates a new CreateProjectAppCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateProjectAppCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateProjectAppCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateProjectAppCommand) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *CreateProjectAppCommand) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CreateProjectAppCommand) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CreateProjectAppCommand) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetProjectId

`func (o *CreateProjectAppCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateProjectAppCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateProjectAppCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetCatalogAppId

`func (o *CreateProjectAppCommand) GetCatalogAppId() int32`

GetCatalogAppId returns the CatalogAppId field if non-nil, zero value otherwise.

### GetCatalogAppIdOk

`func (o *CreateProjectAppCommand) GetCatalogAppIdOk() (*int32, bool)`

GetCatalogAppIdOk returns a tuple with the CatalogAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogAppId

`func (o *CreateProjectAppCommand) SetCatalogAppId(v int32)`

SetCatalogAppId sets CatalogAppId field to given value.


### GetExtraValues

`func (o *CreateProjectAppCommand) GetExtraValues() string`

GetExtraValues returns the ExtraValues field if non-nil, zero value otherwise.

### GetExtraValuesOk

`func (o *CreateProjectAppCommand) GetExtraValuesOk() (*string, bool)`

GetExtraValuesOk returns a tuple with the ExtraValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraValues

`func (o *CreateProjectAppCommand) SetExtraValues(v string)`

SetExtraValues sets ExtraValues field to given value.

### HasExtraValues

`func (o *CreateProjectAppCommand) HasExtraValues() bool`

HasExtraValues returns a boolean if a field has been set.

### GetAutoSync

`func (o *CreateProjectAppCommand) GetAutoSync() bool`

GetAutoSync returns the AutoSync field if non-nil, zero value otherwise.

### GetAutoSyncOk

`func (o *CreateProjectAppCommand) GetAutoSyncOk() (*bool, bool)`

GetAutoSyncOk returns a tuple with the AutoSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSync

`func (o *CreateProjectAppCommand) SetAutoSync(v bool)`

SetAutoSync sets AutoSync field to given value.

### HasAutoSync

`func (o *CreateProjectAppCommand) HasAutoSync() bool`

HasAutoSync returns a boolean if a field has been set.

### GetParameters

`func (o *CreateProjectAppCommand) GetParameters() []ProjectAppParamsDto`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *CreateProjectAppCommand) GetParametersOk() (*[]ProjectAppParamsDto, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *CreateProjectAppCommand) SetParameters(v []ProjectAppParamsDto)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *CreateProjectAppCommand) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


