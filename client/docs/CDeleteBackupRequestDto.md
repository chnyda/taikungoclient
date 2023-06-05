# CDeleteBackupRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataName** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**BackupName** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Phase** | Pointer to **string** |  | [optional] 

## Methods

### NewCDeleteBackupRequestDto

`func NewCDeleteBackupRequestDto() *CDeleteBackupRequestDto`

NewCDeleteBackupRequestDto instantiates a new CDeleteBackupRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCDeleteBackupRequestDtoWithDefaults

`func NewCDeleteBackupRequestDtoWithDefaults() *CDeleteBackupRequestDto`

NewCDeleteBackupRequestDtoWithDefaults instantiates a new CDeleteBackupRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataName

`func (o *CDeleteBackupRequestDto) GetMetadataName() string`

GetMetadataName returns the MetadataName field if non-nil, zero value otherwise.

### GetMetadataNameOk

`func (o *CDeleteBackupRequestDto) GetMetadataNameOk() (*string, bool)`

GetMetadataNameOk returns a tuple with the MetadataName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataName

`func (o *CDeleteBackupRequestDto) SetMetadataName(v string)`

SetMetadataName sets MetadataName field to given value.

### HasMetadataName

`func (o *CDeleteBackupRequestDto) HasMetadataName() bool`

HasMetadataName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CDeleteBackupRequestDto) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CDeleteBackupRequestDto) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CDeleteBackupRequestDto) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CDeleteBackupRequestDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetBackupName

`func (o *CDeleteBackupRequestDto) GetBackupName() string`

GetBackupName returns the BackupName field if non-nil, zero value otherwise.

### GetBackupNameOk

`func (o *CDeleteBackupRequestDto) GetBackupNameOk() (*string, bool)`

GetBackupNameOk returns a tuple with the BackupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupName

`func (o *CDeleteBackupRequestDto) SetBackupName(v string)`

SetBackupName sets BackupName field to given value.

### HasBackupName

`func (o *CDeleteBackupRequestDto) HasBackupName() bool`

HasBackupName returns a boolean if a field has been set.

### GetNamespace

`func (o *CDeleteBackupRequestDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CDeleteBackupRequestDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CDeleteBackupRequestDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CDeleteBackupRequestDto) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPhase

`func (o *CDeleteBackupRequestDto) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *CDeleteBackupRequestDto) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *CDeleteBackupRequestDto) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *CDeleteBackupRequestDto) HasPhase() bool`

HasPhase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


