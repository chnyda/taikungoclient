# CScheduleDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**MetadataName** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Schedule** | Pointer to **string** |  | [optional] 
**Ttl** | Pointer to **string** |  | [optional] 
**LastBackup** | Pointer to **time.Time** |  | [optional] 
**Phase** | Pointer to **string** |  | [optional] 
**ExcludedNamespaces** | Pointer to **[]string** |  | [optional] 
**IncludedNamespaces** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCScheduleDto

`func NewCScheduleDto() *CScheduleDto`

NewCScheduleDto instantiates a new CScheduleDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCScheduleDtoWithDefaults

`func NewCScheduleDtoWithDefaults() *CScheduleDto`

NewCScheduleDtoWithDefaults instantiates a new CScheduleDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CScheduleDto) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CScheduleDto) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CScheduleDto) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CScheduleDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMetadataName

`func (o *CScheduleDto) GetMetadataName() string`

GetMetadataName returns the MetadataName field if non-nil, zero value otherwise.

### GetMetadataNameOk

`func (o *CScheduleDto) GetMetadataNameOk() (*string, bool)`

GetMetadataNameOk returns a tuple with the MetadataName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataName

`func (o *CScheduleDto) SetMetadataName(v string)`

SetMetadataName sets MetadataName field to given value.

### HasMetadataName

`func (o *CScheduleDto) HasMetadataName() bool`

HasMetadataName returns a boolean if a field has been set.

### GetNamespace

`func (o *CScheduleDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CScheduleDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CScheduleDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CScheduleDto) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CScheduleDto) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CScheduleDto) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CScheduleDto) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CScheduleDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetSchedule

`func (o *CScheduleDto) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CScheduleDto) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CScheduleDto) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CScheduleDto) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetTtl

`func (o *CScheduleDto) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *CScheduleDto) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *CScheduleDto) SetTtl(v string)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *CScheduleDto) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetLastBackup

`func (o *CScheduleDto) GetLastBackup() time.Time`

GetLastBackup returns the LastBackup field if non-nil, zero value otherwise.

### GetLastBackupOk

`func (o *CScheduleDto) GetLastBackupOk() (*time.Time, bool)`

GetLastBackupOk returns a tuple with the LastBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackup

`func (o *CScheduleDto) SetLastBackup(v time.Time)`

SetLastBackup sets LastBackup field to given value.

### HasLastBackup

`func (o *CScheduleDto) HasLastBackup() bool`

HasLastBackup returns a boolean if a field has been set.

### GetPhase

`func (o *CScheduleDto) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *CScheduleDto) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *CScheduleDto) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *CScheduleDto) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetExcludedNamespaces

`func (o *CScheduleDto) GetExcludedNamespaces() []string`

GetExcludedNamespaces returns the ExcludedNamespaces field if non-nil, zero value otherwise.

### GetExcludedNamespacesOk

`func (o *CScheduleDto) GetExcludedNamespacesOk() (*[]string, bool)`

GetExcludedNamespacesOk returns a tuple with the ExcludedNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedNamespaces

`func (o *CScheduleDto) SetExcludedNamespaces(v []string)`

SetExcludedNamespaces sets ExcludedNamespaces field to given value.

### HasExcludedNamespaces

`func (o *CScheduleDto) HasExcludedNamespaces() bool`

HasExcludedNamespaces returns a boolean if a field has been set.

### GetIncludedNamespaces

`func (o *CScheduleDto) GetIncludedNamespaces() []string`

GetIncludedNamespaces returns the IncludedNamespaces field if non-nil, zero value otherwise.

### GetIncludedNamespacesOk

`func (o *CScheduleDto) GetIncludedNamespacesOk() (*[]string, bool)`

GetIncludedNamespacesOk returns a tuple with the IncludedNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedNamespaces

`func (o *CScheduleDto) SetIncludedNamespaces(v []string)`

SetIncludedNamespaces sets IncludedNamespaces field to given value.

### HasIncludedNamespaces

`func (o *CScheduleDto) HasIncludedNamespaces() bool`

HasIncludedNamespaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


