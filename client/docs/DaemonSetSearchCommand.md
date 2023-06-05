# DaemonSetSearchCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**SearchTerm** | **string** |  | 

## Methods

### NewDaemonSetSearchCommand

`func NewDaemonSetSearchCommand(searchTerm string, ) *DaemonSetSearchCommand`

NewDaemonSetSearchCommand instantiates a new DaemonSetSearchCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaemonSetSearchCommandWithDefaults

`func NewDaemonSetSearchCommandWithDefaults() *DaemonSetSearchCommand`

NewDaemonSetSearchCommandWithDefaults instantiates a new DaemonSetSearchCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *DaemonSetSearchCommand) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *DaemonSetSearchCommand) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *DaemonSetSearchCommand) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *DaemonSetSearchCommand) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *DaemonSetSearchCommand) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *DaemonSetSearchCommand) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *DaemonSetSearchCommand) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *DaemonSetSearchCommand) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetSearchTerm

`func (o *DaemonSetSearchCommand) GetSearchTerm() string`

GetSearchTerm returns the SearchTerm field if non-nil, zero value otherwise.

### GetSearchTermOk

`func (o *DaemonSetSearchCommand) GetSearchTermOk() (*string, bool)`

GetSearchTermOk returns a tuple with the SearchTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTerm

`func (o *DaemonSetSearchCommand) SetSearchTerm(v string)`

SetSearchTerm sets SearchTerm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


