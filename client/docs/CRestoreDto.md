# CRestoreDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataName** | Pointer to **string** |  | [optional] 
**BackupName** | Pointer to **string** |  | [optional] 
**ScheduleName** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**ExcludeNamespaces** | Pointer to **[]string** |  | [optional] 
**IncludeNamespaces** | Pointer to **[]string** |  | [optional] 
**CompletionDateTime** | Pointer to **time.Time** |  | [optional] 
**StartTimeStamp** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Warnings** | Pointer to **int64** |  | [optional] 
**Phase** | Pointer to **string** |  | [optional] 

## Methods

### NewCRestoreDto

`func NewCRestoreDto() *CRestoreDto`

NewCRestoreDto instantiates a new CRestoreDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCRestoreDtoWithDefaults

`func NewCRestoreDtoWithDefaults() *CRestoreDto`

NewCRestoreDtoWithDefaults instantiates a new CRestoreDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataName

`func (o *CRestoreDto) GetMetadataName() string`

GetMetadataName returns the MetadataName field if non-nil, zero value otherwise.

### GetMetadataNameOk

`func (o *CRestoreDto) GetMetadataNameOk() (*string, bool)`

GetMetadataNameOk returns a tuple with the MetadataName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataName

`func (o *CRestoreDto) SetMetadataName(v string)`

SetMetadataName sets MetadataName field to given value.

### HasMetadataName

`func (o *CRestoreDto) HasMetadataName() bool`

HasMetadataName returns a boolean if a field has been set.

### GetBackupName

`func (o *CRestoreDto) GetBackupName() string`

GetBackupName returns the BackupName field if non-nil, zero value otherwise.

### GetBackupNameOk

`func (o *CRestoreDto) GetBackupNameOk() (*string, bool)`

GetBackupNameOk returns a tuple with the BackupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupName

`func (o *CRestoreDto) SetBackupName(v string)`

SetBackupName sets BackupName field to given value.

### HasBackupName

`func (o *CRestoreDto) HasBackupName() bool`

HasBackupName returns a boolean if a field has been set.

### GetScheduleName

`func (o *CRestoreDto) GetScheduleName() string`

GetScheduleName returns the ScheduleName field if non-nil, zero value otherwise.

### GetScheduleNameOk

`func (o *CRestoreDto) GetScheduleNameOk() (*string, bool)`

GetScheduleNameOk returns a tuple with the ScheduleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleName

`func (o *CRestoreDto) SetScheduleName(v string)`

SetScheduleName sets ScheduleName field to given value.

### HasScheduleName

`func (o *CRestoreDto) HasScheduleName() bool`

HasScheduleName returns a boolean if a field has been set.

### GetNamespace

`func (o *CRestoreDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CRestoreDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CRestoreDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CRestoreDto) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetExcludeNamespaces

`func (o *CRestoreDto) GetExcludeNamespaces() []string`

GetExcludeNamespaces returns the ExcludeNamespaces field if non-nil, zero value otherwise.

### GetExcludeNamespacesOk

`func (o *CRestoreDto) GetExcludeNamespacesOk() (*[]string, bool)`

GetExcludeNamespacesOk returns a tuple with the ExcludeNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeNamespaces

`func (o *CRestoreDto) SetExcludeNamespaces(v []string)`

SetExcludeNamespaces sets ExcludeNamespaces field to given value.

### HasExcludeNamespaces

`func (o *CRestoreDto) HasExcludeNamespaces() bool`

HasExcludeNamespaces returns a boolean if a field has been set.

### GetIncludeNamespaces

`func (o *CRestoreDto) GetIncludeNamespaces() []string`

GetIncludeNamespaces returns the IncludeNamespaces field if non-nil, zero value otherwise.

### GetIncludeNamespacesOk

`func (o *CRestoreDto) GetIncludeNamespacesOk() (*[]string, bool)`

GetIncludeNamespacesOk returns a tuple with the IncludeNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeNamespaces

`func (o *CRestoreDto) SetIncludeNamespaces(v []string)`

SetIncludeNamespaces sets IncludeNamespaces field to given value.

### HasIncludeNamespaces

`func (o *CRestoreDto) HasIncludeNamespaces() bool`

HasIncludeNamespaces returns a boolean if a field has been set.

### GetCompletionDateTime

`func (o *CRestoreDto) GetCompletionDateTime() time.Time`

GetCompletionDateTime returns the CompletionDateTime field if non-nil, zero value otherwise.

### GetCompletionDateTimeOk

`func (o *CRestoreDto) GetCompletionDateTimeOk() (*time.Time, bool)`

GetCompletionDateTimeOk returns a tuple with the CompletionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDateTime

`func (o *CRestoreDto) SetCompletionDateTime(v time.Time)`

SetCompletionDateTime sets CompletionDateTime field to given value.

### HasCompletionDateTime

`func (o *CRestoreDto) HasCompletionDateTime() bool`

HasCompletionDateTime returns a boolean if a field has been set.

### GetStartTimeStamp

`func (o *CRestoreDto) GetStartTimeStamp() time.Time`

GetStartTimeStamp returns the StartTimeStamp field if non-nil, zero value otherwise.

### GetStartTimeStampOk

`func (o *CRestoreDto) GetStartTimeStampOk() (*time.Time, bool)`

GetStartTimeStampOk returns a tuple with the StartTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeStamp

`func (o *CRestoreDto) SetStartTimeStamp(v time.Time)`

SetStartTimeStamp sets StartTimeStamp field to given value.

### HasStartTimeStamp

`func (o *CRestoreDto) HasStartTimeStamp() bool`

HasStartTimeStamp returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CRestoreDto) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CRestoreDto) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CRestoreDto) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CRestoreDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetWarnings

`func (o *CRestoreDto) GetWarnings() int64`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *CRestoreDto) GetWarningsOk() (*int64, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *CRestoreDto) SetWarnings(v int64)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *CRestoreDto) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetPhase

`func (o *CRestoreDto) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *CRestoreDto) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *CRestoreDto) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *CRestoreDto) HasPhase() bool`

HasPhase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


