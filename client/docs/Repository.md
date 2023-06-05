# Repository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Official** | Pointer to **bool** |  | [optional] 
**RepositoryId** | Pointer to **string** |  | [optional] 
**ScannerDisabled** | Pointer to **bool** |  | [optional] 
**OrganizationName** | Pointer to **string** |  | [optional] 
**VerifiedPublisher** | Pointer to **bool** |  | [optional] 
**OrganizationDisplayName** | Pointer to **string** |  | [optional] 

## Methods

### NewRepository

`func NewRepository() *Repository`

NewRepository instantiates a new Repository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryWithDefaults

`func NewRepositoryWithDefaults() *Repository`

NewRepositoryWithDefaults instantiates a new Repository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *Repository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Repository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Repository) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Repository) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetKind

`func (o *Repository) GetKind() int64`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Repository) GetKindOk() (*int64, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Repository) SetKind(v int64)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Repository) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *Repository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Repository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Repository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Repository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOfficial

`func (o *Repository) GetOfficial() bool`

GetOfficial returns the Official field if non-nil, zero value otherwise.

### GetOfficialOk

`func (o *Repository) GetOfficialOk() (*bool, bool)`

GetOfficialOk returns a tuple with the Official field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficial

`func (o *Repository) SetOfficial(v bool)`

SetOfficial sets Official field to given value.

### HasOfficial

`func (o *Repository) HasOfficial() bool`

HasOfficial returns a boolean if a field has been set.

### GetRepositoryId

`func (o *Repository) GetRepositoryId() string`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *Repository) GetRepositoryIdOk() (*string, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *Repository) SetRepositoryId(v string)`

SetRepositoryId sets RepositoryId field to given value.

### HasRepositoryId

`func (o *Repository) HasRepositoryId() bool`

HasRepositoryId returns a boolean if a field has been set.

### GetScannerDisabled

`func (o *Repository) GetScannerDisabled() bool`

GetScannerDisabled returns the ScannerDisabled field if non-nil, zero value otherwise.

### GetScannerDisabledOk

`func (o *Repository) GetScannerDisabledOk() (*bool, bool)`

GetScannerDisabledOk returns a tuple with the ScannerDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScannerDisabled

`func (o *Repository) SetScannerDisabled(v bool)`

SetScannerDisabled sets ScannerDisabled field to given value.

### HasScannerDisabled

`func (o *Repository) HasScannerDisabled() bool`

HasScannerDisabled returns a boolean if a field has been set.

### GetOrganizationName

`func (o *Repository) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *Repository) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *Repository) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *Repository) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetVerifiedPublisher

`func (o *Repository) GetVerifiedPublisher() bool`

GetVerifiedPublisher returns the VerifiedPublisher field if non-nil, zero value otherwise.

### GetVerifiedPublisherOk

`func (o *Repository) GetVerifiedPublisherOk() (*bool, bool)`

GetVerifiedPublisherOk returns a tuple with the VerifiedPublisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedPublisher

`func (o *Repository) SetVerifiedPublisher(v bool)`

SetVerifiedPublisher sets VerifiedPublisher field to given value.

### HasVerifiedPublisher

`func (o *Repository) HasVerifiedPublisher() bool`

HasVerifiedPublisher returns a boolean if a field has been set.

### GetOrganizationDisplayName

`func (o *Repository) GetOrganizationDisplayName() string`

GetOrganizationDisplayName returns the OrganizationDisplayName field if non-nil, zero value otherwise.

### GetOrganizationDisplayNameOk

`func (o *Repository) GetOrganizationDisplayNameOk() (*string, bool)`

GetOrganizationDisplayNameOk returns a tuple with the OrganizationDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationDisplayName

`func (o *Repository) SetOrganizationDisplayName(v string)`

SetOrganizationDisplayName sets OrganizationDisplayName field to given value.

### HasOrganizationDisplayName

`func (o *Repository) HasOrganizationDisplayName() bool`

HasOrganizationDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


