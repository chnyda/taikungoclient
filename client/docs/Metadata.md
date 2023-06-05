# Metadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InfracostCommand** | Pointer to **string** |  | [optional] 
**VcsBranch** | Pointer to **string** |  | [optional] 
**VcsCommitSha** | Pointer to **string** |  | [optional] 
**VcsCommitAuthorName** | Pointer to **string** |  | [optional] 
**VcsCommitAuthorEmail** | Pointer to **string** |  | [optional] 
**VcsCommitTimestamp** | Pointer to **time.Time** |  | [optional] 
**VcsCommitMessage** | Pointer to **string** |  | [optional] 
**VcsRepositoryUrl** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**VcsSubPath** | Pointer to **string** |  | [optional] 

## Methods

### NewMetadata

`func NewMetadata() *Metadata`

NewMetadata instantiates a new Metadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithDefaults

`func NewMetadataWithDefaults() *Metadata`

NewMetadataWithDefaults instantiates a new Metadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfracostCommand

`func (o *Metadata) GetInfracostCommand() string`

GetInfracostCommand returns the InfracostCommand field if non-nil, zero value otherwise.

### GetInfracostCommandOk

`func (o *Metadata) GetInfracostCommandOk() (*string, bool)`

GetInfracostCommandOk returns a tuple with the InfracostCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfracostCommand

`func (o *Metadata) SetInfracostCommand(v string)`

SetInfracostCommand sets InfracostCommand field to given value.

### HasInfracostCommand

`func (o *Metadata) HasInfracostCommand() bool`

HasInfracostCommand returns a boolean if a field has been set.

### GetVcsBranch

`func (o *Metadata) GetVcsBranch() string`

GetVcsBranch returns the VcsBranch field if non-nil, zero value otherwise.

### GetVcsBranchOk

`func (o *Metadata) GetVcsBranchOk() (*string, bool)`

GetVcsBranchOk returns a tuple with the VcsBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsBranch

`func (o *Metadata) SetVcsBranch(v string)`

SetVcsBranch sets VcsBranch field to given value.

### HasVcsBranch

`func (o *Metadata) HasVcsBranch() bool`

HasVcsBranch returns a boolean if a field has been set.

### GetVcsCommitSha

`func (o *Metadata) GetVcsCommitSha() string`

GetVcsCommitSha returns the VcsCommitSha field if non-nil, zero value otherwise.

### GetVcsCommitShaOk

`func (o *Metadata) GetVcsCommitShaOk() (*string, bool)`

GetVcsCommitShaOk returns a tuple with the VcsCommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsCommitSha

`func (o *Metadata) SetVcsCommitSha(v string)`

SetVcsCommitSha sets VcsCommitSha field to given value.

### HasVcsCommitSha

`func (o *Metadata) HasVcsCommitSha() bool`

HasVcsCommitSha returns a boolean if a field has been set.

### GetVcsCommitAuthorName

`func (o *Metadata) GetVcsCommitAuthorName() string`

GetVcsCommitAuthorName returns the VcsCommitAuthorName field if non-nil, zero value otherwise.

### GetVcsCommitAuthorNameOk

`func (o *Metadata) GetVcsCommitAuthorNameOk() (*string, bool)`

GetVcsCommitAuthorNameOk returns a tuple with the VcsCommitAuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsCommitAuthorName

`func (o *Metadata) SetVcsCommitAuthorName(v string)`

SetVcsCommitAuthorName sets VcsCommitAuthorName field to given value.

### HasVcsCommitAuthorName

`func (o *Metadata) HasVcsCommitAuthorName() bool`

HasVcsCommitAuthorName returns a boolean if a field has been set.

### GetVcsCommitAuthorEmail

`func (o *Metadata) GetVcsCommitAuthorEmail() string`

GetVcsCommitAuthorEmail returns the VcsCommitAuthorEmail field if non-nil, zero value otherwise.

### GetVcsCommitAuthorEmailOk

`func (o *Metadata) GetVcsCommitAuthorEmailOk() (*string, bool)`

GetVcsCommitAuthorEmailOk returns a tuple with the VcsCommitAuthorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsCommitAuthorEmail

`func (o *Metadata) SetVcsCommitAuthorEmail(v string)`

SetVcsCommitAuthorEmail sets VcsCommitAuthorEmail field to given value.

### HasVcsCommitAuthorEmail

`func (o *Metadata) HasVcsCommitAuthorEmail() bool`

HasVcsCommitAuthorEmail returns a boolean if a field has been set.

### GetVcsCommitTimestamp

`func (o *Metadata) GetVcsCommitTimestamp() time.Time`

GetVcsCommitTimestamp returns the VcsCommitTimestamp field if non-nil, zero value otherwise.

### GetVcsCommitTimestampOk

`func (o *Metadata) GetVcsCommitTimestampOk() (*time.Time, bool)`

GetVcsCommitTimestampOk returns a tuple with the VcsCommitTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsCommitTimestamp

`func (o *Metadata) SetVcsCommitTimestamp(v time.Time)`

SetVcsCommitTimestamp sets VcsCommitTimestamp field to given value.

### HasVcsCommitTimestamp

`func (o *Metadata) HasVcsCommitTimestamp() bool`

HasVcsCommitTimestamp returns a boolean if a field has been set.

### GetVcsCommitMessage

`func (o *Metadata) GetVcsCommitMessage() string`

GetVcsCommitMessage returns the VcsCommitMessage field if non-nil, zero value otherwise.

### GetVcsCommitMessageOk

`func (o *Metadata) GetVcsCommitMessageOk() (*string, bool)`

GetVcsCommitMessageOk returns a tuple with the VcsCommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsCommitMessage

`func (o *Metadata) SetVcsCommitMessage(v string)`

SetVcsCommitMessage sets VcsCommitMessage field to given value.

### HasVcsCommitMessage

`func (o *Metadata) HasVcsCommitMessage() bool`

HasVcsCommitMessage returns a boolean if a field has been set.

### GetVcsRepositoryUrl

`func (o *Metadata) GetVcsRepositoryUrl() string`

GetVcsRepositoryUrl returns the VcsRepositoryUrl field if non-nil, zero value otherwise.

### GetVcsRepositoryUrlOk

`func (o *Metadata) GetVcsRepositoryUrlOk() (*string, bool)`

GetVcsRepositoryUrlOk returns a tuple with the VcsRepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsRepositoryUrl

`func (o *Metadata) SetVcsRepositoryUrl(v string)`

SetVcsRepositoryUrl sets VcsRepositoryUrl field to given value.

### HasVcsRepositoryUrl

`func (o *Metadata) HasVcsRepositoryUrl() bool`

HasVcsRepositoryUrl returns a boolean if a field has been set.

### GetPath

`func (o *Metadata) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Metadata) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Metadata) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Metadata) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetType

`func (o *Metadata) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Metadata) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Metadata) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Metadata) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVcsSubPath

`func (o *Metadata) GetVcsSubPath() string`

GetVcsSubPath returns the VcsSubPath field if non-nil, zero value otherwise.

### GetVcsSubPathOk

`func (o *Metadata) GetVcsSubPathOk() (*string, bool)`

GetVcsSubPathOk returns a tuple with the VcsSubPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsSubPath

`func (o *Metadata) SetVcsSubPath(v string)`

SetVcsSubPath sets VcsSubPath field to given value.

### HasVcsSubPath

`func (o *Metadata) HasVcsSubPath() bool`

HasVcsSubPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


