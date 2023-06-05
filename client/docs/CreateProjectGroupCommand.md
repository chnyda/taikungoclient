# CreateProjectGroupCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**ProjectIds** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewCreateProjectGroupCommand

`func NewCreateProjectGroupCommand(name string, ) *CreateProjectGroupCommand`

NewCreateProjectGroupCommand instantiates a new CreateProjectGroupCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProjectGroupCommandWithDefaults

`func NewCreateProjectGroupCommandWithDefaults() *CreateProjectGroupCommand`

NewCreateProjectGroupCommandWithDefaults instantiates a new CreateProjectGroupCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateProjectGroupCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateProjectGroupCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateProjectGroupCommand) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *CreateProjectGroupCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateProjectGroupCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateProjectGroupCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CreateProjectGroupCommand) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetProjectIds

`func (o *CreateProjectGroupCommand) GetProjectIds() []int32`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *CreateProjectGroupCommand) GetProjectIdsOk() (*[]int32, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *CreateProjectGroupCommand) SetProjectIds(v []int32)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *CreateProjectGroupCommand) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


