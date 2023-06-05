# UpdateHypervisorsCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Hypervisors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateHypervisorsCommand

`func NewUpdateHypervisorsCommand(id int32, ) *UpdateHypervisorsCommand`

NewUpdateHypervisorsCommand instantiates a new UpdateHypervisorsCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateHypervisorsCommandWithDefaults

`func NewUpdateHypervisorsCommandWithDefaults() *UpdateHypervisorsCommand`

NewUpdateHypervisorsCommandWithDefaults instantiates a new UpdateHypervisorsCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateHypervisorsCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateHypervisorsCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateHypervisorsCommand) SetId(v int32)`

SetId sets Id field to given value.


### GetHypervisors

`func (o *UpdateHypervisorsCommand) GetHypervisors() []string`

GetHypervisors returns the Hypervisors field if non-nil, zero value otherwise.

### GetHypervisorsOk

`func (o *UpdateHypervisorsCommand) GetHypervisorsOk() (*[]string, bool)`

GetHypervisorsOk returns a tuple with the Hypervisors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisors

`func (o *UpdateHypervisorsCommand) SetHypervisors(v []string)`

SetHypervisors sets Hypervisors field to given value.

### HasHypervisors

`func (o *UpdateHypervisorsCommand) HasHypervisors() bool`

HasHypervisors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


