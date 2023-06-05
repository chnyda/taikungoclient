# StsSearchCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**SearchTerm** | **string** |  | 

## Methods

### NewStsSearchCommand

`func NewStsSearchCommand(searchTerm string, ) *StsSearchCommand`

NewStsSearchCommand instantiates a new StsSearchCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStsSearchCommandWithDefaults

`func NewStsSearchCommandWithDefaults() *StsSearchCommand`

NewStsSearchCommandWithDefaults instantiates a new StsSearchCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *StsSearchCommand) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *StsSearchCommand) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *StsSearchCommand) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *StsSearchCommand) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *StsSearchCommand) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *StsSearchCommand) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *StsSearchCommand) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *StsSearchCommand) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetSearchTerm

`func (o *StsSearchCommand) GetSearchTerm() string`

GetSearchTerm returns the SearchTerm field if non-nil, zero value otherwise.

### GetSearchTermOk

`func (o *StsSearchCommand) GetSearchTermOk() (*string, bool)`

GetSearchTermOk returns a tuple with the SearchTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTerm

`func (o *StsSearchCommand) SetSearchTerm(v string)`

SetSearchTerm sets SearchTerm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


