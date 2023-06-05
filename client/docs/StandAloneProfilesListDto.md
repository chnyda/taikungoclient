# StandAloneProfilesListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**StandaloneVms** | Pointer to [**[]StandAloneVmSmallDetailDto**](StandAloneVmSmallDetailDto.md) |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**OrganizationName** | Pointer to **string** |  | [optional] 
**PartnerLogo** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewStandAloneProfilesListDto

`func NewStandAloneProfilesListDto() *StandAloneProfilesListDto`

NewStandAloneProfilesListDto instantiates a new StandAloneProfilesListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandAloneProfilesListDtoWithDefaults

`func NewStandAloneProfilesListDtoWithDefaults() *StandAloneProfilesListDto`

NewStandAloneProfilesListDtoWithDefaults instantiates a new StandAloneProfilesListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StandAloneProfilesListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StandAloneProfilesListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StandAloneProfilesListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *StandAloneProfilesListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *StandAloneProfilesListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StandAloneProfilesListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StandAloneProfilesListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StandAloneProfilesListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublicKey

`func (o *StandAloneProfilesListDto) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *StandAloneProfilesListDto) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *StandAloneProfilesListDto) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *StandAloneProfilesListDto) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetIsLocked

`func (o *StandAloneProfilesListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *StandAloneProfilesListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *StandAloneProfilesListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *StandAloneProfilesListDto) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetStandaloneVms

`func (o *StandAloneProfilesListDto) GetStandaloneVms() []StandAloneVmSmallDetailDto`

GetStandaloneVms returns the StandaloneVms field if non-nil, zero value otherwise.

### GetStandaloneVmsOk

`func (o *StandAloneProfilesListDto) GetStandaloneVmsOk() (*[]StandAloneVmSmallDetailDto, bool)`

GetStandaloneVmsOk returns a tuple with the StandaloneVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandaloneVms

`func (o *StandAloneProfilesListDto) SetStandaloneVms(v []StandAloneVmSmallDetailDto)`

SetStandaloneVms sets StandaloneVms field to given value.

### HasStandaloneVms

`func (o *StandAloneProfilesListDto) HasStandaloneVms() bool`

HasStandaloneVms returns a boolean if a field has been set.

### GetOrganizationId

`func (o *StandAloneProfilesListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *StandAloneProfilesListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *StandAloneProfilesListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *StandAloneProfilesListDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetOrganizationName

`func (o *StandAloneProfilesListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *StandAloneProfilesListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *StandAloneProfilesListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *StandAloneProfilesListDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetPartnerLogo

`func (o *StandAloneProfilesListDto) GetPartnerLogo() string`

GetPartnerLogo returns the PartnerLogo field if non-nil, zero value otherwise.

### GetPartnerLogoOk

`func (o *StandAloneProfilesListDto) GetPartnerLogoOk() (*string, bool)`

GetPartnerLogoOk returns a tuple with the PartnerLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerLogo

`func (o *StandAloneProfilesListDto) SetPartnerLogo(v string)`

SetPartnerLogo sets PartnerLogo field to given value.

### HasPartnerLogo

`func (o *StandAloneProfilesListDto) HasPartnerLogo() bool`

HasPartnerLogo returns a boolean if a field has been set.

### GetCreatedAt

`func (o *StandAloneProfilesListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StandAloneProfilesListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StandAloneProfilesListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *StandAloneProfilesListDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


