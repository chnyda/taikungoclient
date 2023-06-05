# StandAloneProfileCreateCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**PublicKey** | **string** |  | 
**SecurityGroups** | Pointer to [**[]StandAloneProfileSecurityGroupDto**](StandAloneProfileSecurityGroupDto.md) |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 

## Methods

### NewStandAloneProfileCreateCommand

`func NewStandAloneProfileCreateCommand(name string, publicKey string, ) *StandAloneProfileCreateCommand`

NewStandAloneProfileCreateCommand instantiates a new StandAloneProfileCreateCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandAloneProfileCreateCommandWithDefaults

`func NewStandAloneProfileCreateCommandWithDefaults() *StandAloneProfileCreateCommand`

NewStandAloneProfileCreateCommandWithDefaults instantiates a new StandAloneProfileCreateCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StandAloneProfileCreateCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StandAloneProfileCreateCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StandAloneProfileCreateCommand) SetName(v string)`

SetName sets Name field to given value.


### GetPublicKey

`func (o *StandAloneProfileCreateCommand) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *StandAloneProfileCreateCommand) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *StandAloneProfileCreateCommand) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetSecurityGroups

`func (o *StandAloneProfileCreateCommand) GetSecurityGroups() []StandAloneProfileSecurityGroupDto`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *StandAloneProfileCreateCommand) GetSecurityGroupsOk() (*[]StandAloneProfileSecurityGroupDto, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *StandAloneProfileCreateCommand) SetSecurityGroups(v []StandAloneProfileSecurityGroupDto)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *StandAloneProfileCreateCommand) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetOrganizationId

`func (o *StandAloneProfileCreateCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *StandAloneProfileCreateCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *StandAloneProfileCreateCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *StandAloneProfileCreateCommand) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


