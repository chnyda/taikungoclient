# KubeConfigForUserDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**ProjectName** | Pointer to **string** |  | [optional] 
**IsAccessibleForAll** | Pointer to **bool** |  | [optional] 
**IsAccessibleForManager** | Pointer to **bool** |  | [optional] 
**KubeConfigRoleName** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**ExpirationDate** | Pointer to **string** |  | [optional] 

## Methods

### NewKubeConfigForUserDto

`func NewKubeConfigForUserDto() *KubeConfigForUserDto`

NewKubeConfigForUserDto instantiates a new KubeConfigForUserDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubeConfigForUserDtoWithDefaults

`func NewKubeConfigForUserDtoWithDefaults() *KubeConfigForUserDto`

NewKubeConfigForUserDtoWithDefaults instantiates a new KubeConfigForUserDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubeConfigForUserDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubeConfigForUserDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubeConfigForUserDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KubeConfigForUserDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *KubeConfigForUserDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *KubeConfigForUserDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *KubeConfigForUserDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *KubeConfigForUserDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetDisplayName

`func (o *KubeConfigForUserDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KubeConfigForUserDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KubeConfigForUserDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *KubeConfigForUserDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetProjectId

`func (o *KubeConfigForUserDto) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *KubeConfigForUserDto) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *KubeConfigForUserDto) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *KubeConfigForUserDto) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProjectName

`func (o *KubeConfigForUserDto) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *KubeConfigForUserDto) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *KubeConfigForUserDto) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *KubeConfigForUserDto) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetIsAccessibleForAll

`func (o *KubeConfigForUserDto) GetIsAccessibleForAll() bool`

GetIsAccessibleForAll returns the IsAccessibleForAll field if non-nil, zero value otherwise.

### GetIsAccessibleForAllOk

`func (o *KubeConfigForUserDto) GetIsAccessibleForAllOk() (*bool, bool)`

GetIsAccessibleForAllOk returns a tuple with the IsAccessibleForAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccessibleForAll

`func (o *KubeConfigForUserDto) SetIsAccessibleForAll(v bool)`

SetIsAccessibleForAll sets IsAccessibleForAll field to given value.

### HasIsAccessibleForAll

`func (o *KubeConfigForUserDto) HasIsAccessibleForAll() bool`

HasIsAccessibleForAll returns a boolean if a field has been set.

### GetIsAccessibleForManager

`func (o *KubeConfigForUserDto) GetIsAccessibleForManager() bool`

GetIsAccessibleForManager returns the IsAccessibleForManager field if non-nil, zero value otherwise.

### GetIsAccessibleForManagerOk

`func (o *KubeConfigForUserDto) GetIsAccessibleForManagerOk() (*bool, bool)`

GetIsAccessibleForManagerOk returns a tuple with the IsAccessibleForManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccessibleForManager

`func (o *KubeConfigForUserDto) SetIsAccessibleForManager(v bool)`

SetIsAccessibleForManager sets IsAccessibleForManager field to given value.

### HasIsAccessibleForManager

`func (o *KubeConfigForUserDto) HasIsAccessibleForManager() bool`

HasIsAccessibleForManager returns a boolean if a field has been set.

### GetKubeConfigRoleName

`func (o *KubeConfigForUserDto) GetKubeConfigRoleName() string`

GetKubeConfigRoleName returns the KubeConfigRoleName field if non-nil, zero value otherwise.

### GetKubeConfigRoleNameOk

`func (o *KubeConfigForUserDto) GetKubeConfigRoleNameOk() (*string, bool)`

GetKubeConfigRoleNameOk returns a tuple with the KubeConfigRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeConfigRoleName

`func (o *KubeConfigForUserDto) SetKubeConfigRoleName(v string)`

SetKubeConfigRoleName sets KubeConfigRoleName field to given value.

### HasKubeConfigRoleName

`func (o *KubeConfigForUserDto) HasKubeConfigRoleName() bool`

HasKubeConfigRoleName returns a boolean if a field has been set.

### GetCreatedBy

`func (o *KubeConfigForUserDto) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *KubeConfigForUserDto) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *KubeConfigForUserDto) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *KubeConfigForUserDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KubeConfigForUserDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KubeConfigForUserDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KubeConfigForUserDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KubeConfigForUserDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetNamespace

`func (o *KubeConfigForUserDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *KubeConfigForUserDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *KubeConfigForUserDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *KubeConfigForUserDto) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetExpirationDate

`func (o *KubeConfigForUserDto) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *KubeConfigForUserDto) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *KubeConfigForUserDto) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *KubeConfigForUserDto) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


