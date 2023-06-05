# BackupStorageLocationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataName** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**LastValidated** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**AccessMode** | Pointer to **string** |  | [optional] 
**Phase** | Pointer to **string** |  | [optional] 
**BackupCredentialId** | Pointer to **int32** |  | [optional] 

## Methods

### NewBackupStorageLocationDto

`func NewBackupStorageLocationDto() *BackupStorageLocationDto`

NewBackupStorageLocationDto instantiates a new BackupStorageLocationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupStorageLocationDtoWithDefaults

`func NewBackupStorageLocationDtoWithDefaults() *BackupStorageLocationDto`

NewBackupStorageLocationDtoWithDefaults instantiates a new BackupStorageLocationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataName

`func (o *BackupStorageLocationDto) GetMetadataName() string`

GetMetadataName returns the MetadataName field if non-nil, zero value otherwise.

### GetMetadataNameOk

`func (o *BackupStorageLocationDto) GetMetadataNameOk() (*string, bool)`

GetMetadataNameOk returns a tuple with the MetadataName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataName

`func (o *BackupStorageLocationDto) SetMetadataName(v string)`

SetMetadataName sets MetadataName field to given value.

### HasMetadataName

`func (o *BackupStorageLocationDto) HasMetadataName() bool`

HasMetadataName returns a boolean if a field has been set.

### GetProvider

`func (o *BackupStorageLocationDto) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *BackupStorageLocationDto) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *BackupStorageLocationDto) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *BackupStorageLocationDto) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetNamespace

`func (o *BackupStorageLocationDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *BackupStorageLocationDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *BackupStorageLocationDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *BackupStorageLocationDto) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetLastValidated

`func (o *BackupStorageLocationDto) GetLastValidated() time.Time`

GetLastValidated returns the LastValidated field if non-nil, zero value otherwise.

### GetLastValidatedOk

`func (o *BackupStorageLocationDto) GetLastValidatedOk() (*time.Time, bool)`

GetLastValidatedOk returns a tuple with the LastValidated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastValidated

`func (o *BackupStorageLocationDto) SetLastValidated(v time.Time)`

SetLastValidated sets LastValidated field to given value.

### HasLastValidated

`func (o *BackupStorageLocationDto) HasLastValidated() bool`

HasLastValidated returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BackupStorageLocationDto) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BackupStorageLocationDto) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BackupStorageLocationDto) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BackupStorageLocationDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetAccessMode

`func (o *BackupStorageLocationDto) GetAccessMode() string`

GetAccessMode returns the AccessMode field if non-nil, zero value otherwise.

### GetAccessModeOk

`func (o *BackupStorageLocationDto) GetAccessModeOk() (*string, bool)`

GetAccessModeOk returns a tuple with the AccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMode

`func (o *BackupStorageLocationDto) SetAccessMode(v string)`

SetAccessMode sets AccessMode field to given value.

### HasAccessMode

`func (o *BackupStorageLocationDto) HasAccessMode() bool`

HasAccessMode returns a boolean if a field has been set.

### GetPhase

`func (o *BackupStorageLocationDto) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *BackupStorageLocationDto) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *BackupStorageLocationDto) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *BackupStorageLocationDto) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetBackupCredentialId

`func (o *BackupStorageLocationDto) GetBackupCredentialId() int32`

GetBackupCredentialId returns the BackupCredentialId field if non-nil, zero value otherwise.

### GetBackupCredentialIdOk

`func (o *BackupStorageLocationDto) GetBackupCredentialIdOk() (*int32, bool)`

GetBackupCredentialIdOk returns a tuple with the BackupCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupCredentialId

`func (o *BackupStorageLocationDto) SetBackupCredentialId(v int32)`

SetBackupCredentialId sets BackupCredentialId field to given value.

### HasBackupCredentialId

`func (o *BackupStorageLocationDto) HasBackupCredentialId() bool`

HasBackupCredentialId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


