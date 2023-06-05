# PodsSearchCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**SearchTerm** | **string** |  | 

## Methods

### NewPodsSearchCommand

`func NewPodsSearchCommand(searchTerm string, ) *PodsSearchCommand`

NewPodsSearchCommand instantiates a new PodsSearchCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPodsSearchCommandWithDefaults

`func NewPodsSearchCommandWithDefaults() *PodsSearchCommand`

NewPodsSearchCommandWithDefaults instantiates a new PodsSearchCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *PodsSearchCommand) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PodsSearchCommand) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PodsSearchCommand) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PodsSearchCommand) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *PodsSearchCommand) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PodsSearchCommand) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PodsSearchCommand) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *PodsSearchCommand) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetSearchTerm

`func (o *PodsSearchCommand) GetSearchTerm() string`

GetSearchTerm returns the SearchTerm field if non-nil, zero value otherwise.

### GetSearchTermOk

`func (o *PodsSearchCommand) GetSearchTermOk() (*string, bool)`

GetSearchTermOk returns a tuple with the SearchTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTerm

`func (o *PodsSearchCommand) SetSearchTerm(v string)`

SetSearchTerm sets SearchTerm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


