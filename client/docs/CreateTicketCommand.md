# CreateTicketCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** |  | 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**Priority** | Pointer to [**TicketPriority**](TicketPriority.md) |  | [optional] 

## Methods

### NewCreateTicketCommand

`func NewCreateTicketCommand(name string, description string, ) *CreateTicketCommand`

NewCreateTicketCommand instantiates a new CreateTicketCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTicketCommandWithDefaults

`func NewCreateTicketCommandWithDefaults() *CreateTicketCommand`

NewCreateTicketCommandWithDefaults instantiates a new CreateTicketCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTicketCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTicketCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTicketCommand) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateTicketCommand) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTicketCommand) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTicketCommand) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetOrganizationId

`func (o *CreateTicketCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateTicketCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateTicketCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CreateTicketCommand) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetPriority

`func (o *CreateTicketCommand) GetPriority() TicketPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreateTicketCommand) GetPriorityOk() (*TicketPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreateTicketCommand) SetPriority(v TicketPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CreateTicketCommand) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


