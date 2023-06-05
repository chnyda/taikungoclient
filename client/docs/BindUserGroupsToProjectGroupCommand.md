# BindUserGroupsToProjectGroupCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserGroups** | Pointer to [**[]UpdateProjectUserGroupDto**](UpdateProjectUserGroupDto.md) |  | [optional] 
**ProjectGroupId** | **int32** |  | 
**ProjectGroupName** | Pointer to **string** |  | [optional] 

## Methods

### NewBindUserGroupsToProjectGroupCommand

`func NewBindUserGroupsToProjectGroupCommand(projectGroupId int32, ) *BindUserGroupsToProjectGroupCommand`

NewBindUserGroupsToProjectGroupCommand instantiates a new BindUserGroupsToProjectGroupCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBindUserGroupsToProjectGroupCommandWithDefaults

`func NewBindUserGroupsToProjectGroupCommandWithDefaults() *BindUserGroupsToProjectGroupCommand`

NewBindUserGroupsToProjectGroupCommandWithDefaults instantiates a new BindUserGroupsToProjectGroupCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserGroups

`func (o *BindUserGroupsToProjectGroupCommand) GetUserGroups() []UpdateProjectUserGroupDto`

GetUserGroups returns the UserGroups field if non-nil, zero value otherwise.

### GetUserGroupsOk

`func (o *BindUserGroupsToProjectGroupCommand) GetUserGroupsOk() (*[]UpdateProjectUserGroupDto, bool)`

GetUserGroupsOk returns a tuple with the UserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroups

`func (o *BindUserGroupsToProjectGroupCommand) SetUserGroups(v []UpdateProjectUserGroupDto)`

SetUserGroups sets UserGroups field to given value.

### HasUserGroups

`func (o *BindUserGroupsToProjectGroupCommand) HasUserGroups() bool`

HasUserGroups returns a boolean if a field has been set.

### GetProjectGroupId

`func (o *BindUserGroupsToProjectGroupCommand) GetProjectGroupId() int32`

GetProjectGroupId returns the ProjectGroupId field if non-nil, zero value otherwise.

### GetProjectGroupIdOk

`func (o *BindUserGroupsToProjectGroupCommand) GetProjectGroupIdOk() (*int32, bool)`

GetProjectGroupIdOk returns a tuple with the ProjectGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectGroupId

`func (o *BindUserGroupsToProjectGroupCommand) SetProjectGroupId(v int32)`

SetProjectGroupId sets ProjectGroupId field to given value.


### GetProjectGroupName

`func (o *BindUserGroupsToProjectGroupCommand) GetProjectGroupName() string`

GetProjectGroupName returns the ProjectGroupName field if non-nil, zero value otherwise.

### GetProjectGroupNameOk

`func (o *BindUserGroupsToProjectGroupCommand) GetProjectGroupNameOk() (*string, bool)`

GetProjectGroupNameOk returns a tuple with the ProjectGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectGroupName

`func (o *BindUserGroupsToProjectGroupCommand) SetProjectGroupName(v string)`

SetProjectGroupName sets ProjectGroupName field to given value.

### HasProjectGroupName

`func (o *BindUserGroupsToProjectGroupCommand) HasProjectGroupName() bool`

HasProjectGroupName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


