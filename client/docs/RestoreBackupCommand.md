# RestoreBackupCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **int32** |  | [optional] 
**BackupName** | Pointer to **string** |  | [optional] 
**RestoreName** | Pointer to **string** |  | [optional] 
**IncludeNamespaces** | Pointer to **[]string** |  | [optional] 
**ExcludeNamespaces** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRestoreBackupCommand

`func NewRestoreBackupCommand() *RestoreBackupCommand`

NewRestoreBackupCommand instantiates a new RestoreBackupCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreBackupCommandWithDefaults

`func NewRestoreBackupCommandWithDefaults() *RestoreBackupCommand`

NewRestoreBackupCommandWithDefaults instantiates a new RestoreBackupCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *RestoreBackupCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *RestoreBackupCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *RestoreBackupCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *RestoreBackupCommand) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetBackupName

`func (o *RestoreBackupCommand) GetBackupName() string`

GetBackupName returns the BackupName field if non-nil, zero value otherwise.

### GetBackupNameOk

`func (o *RestoreBackupCommand) GetBackupNameOk() (*string, bool)`

GetBackupNameOk returns a tuple with the BackupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupName

`func (o *RestoreBackupCommand) SetBackupName(v string)`

SetBackupName sets BackupName field to given value.

### HasBackupName

`func (o *RestoreBackupCommand) HasBackupName() bool`

HasBackupName returns a boolean if a field has been set.

### GetRestoreName

`func (o *RestoreBackupCommand) GetRestoreName() string`

GetRestoreName returns the RestoreName field if non-nil, zero value otherwise.

### GetRestoreNameOk

`func (o *RestoreBackupCommand) GetRestoreNameOk() (*string, bool)`

GetRestoreNameOk returns a tuple with the RestoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreName

`func (o *RestoreBackupCommand) SetRestoreName(v string)`

SetRestoreName sets RestoreName field to given value.

### HasRestoreName

`func (o *RestoreBackupCommand) HasRestoreName() bool`

HasRestoreName returns a boolean if a field has been set.

### GetIncludeNamespaces

`func (o *RestoreBackupCommand) GetIncludeNamespaces() []string`

GetIncludeNamespaces returns the IncludeNamespaces field if non-nil, zero value otherwise.

### GetIncludeNamespacesOk

`func (o *RestoreBackupCommand) GetIncludeNamespacesOk() (*[]string, bool)`

GetIncludeNamespacesOk returns a tuple with the IncludeNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeNamespaces

`func (o *RestoreBackupCommand) SetIncludeNamespaces(v []string)`

SetIncludeNamespaces sets IncludeNamespaces field to given value.

### HasIncludeNamespaces

`func (o *RestoreBackupCommand) HasIncludeNamespaces() bool`

HasIncludeNamespaces returns a boolean if a field has been set.

### GetExcludeNamespaces

`func (o *RestoreBackupCommand) GetExcludeNamespaces() []string`

GetExcludeNamespaces returns the ExcludeNamespaces field if non-nil, zero value otherwise.

### GetExcludeNamespacesOk

`func (o *RestoreBackupCommand) GetExcludeNamespacesOk() (*[]string, bool)`

GetExcludeNamespacesOk returns a tuple with the ExcludeNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeNamespaces

`func (o *RestoreBackupCommand) SetExcludeNamespaces(v []string)`

SetExcludeNamespaces sets ExcludeNamespaces field to given value.

### HasExcludeNamespaces

`func (o *RestoreBackupCommand) HasExcludeNamespaces() bool`

HasExcludeNamespaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


