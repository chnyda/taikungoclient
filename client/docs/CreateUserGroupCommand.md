# CreateUserGroupCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**UserIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateUserGroupCommand

`func NewCreateUserGroupCommand(name string, ) *CreateUserGroupCommand`

NewCreateUserGroupCommand instantiates a new CreateUserGroupCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserGroupCommandWithDefaults

`func NewCreateUserGroupCommandWithDefaults() *CreateUserGroupCommand`

NewCreateUserGroupCommandWithDefaults instantiates a new CreateUserGroupCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateUserGroupCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUserGroupCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUserGroupCommand) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *CreateUserGroupCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateUserGroupCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateUserGroupCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CreateUserGroupCommand) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetUserIds

`func (o *CreateUserGroupCommand) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *CreateUserGroupCommand) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *CreateUserGroupCommand) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *CreateUserGroupCommand) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


