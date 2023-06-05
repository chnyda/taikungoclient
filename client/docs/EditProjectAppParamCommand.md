# EditProjectAppParamCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectAppId** | **int32** |  | 
**Parameters** | Pointer to [**[]EditProjectAppParamsDto**](EditProjectAppParamsDto.md) |  | [optional] 

## Methods

### NewEditProjectAppParamCommand

`func NewEditProjectAppParamCommand(projectAppId int32, ) *EditProjectAppParamCommand`

NewEditProjectAppParamCommand instantiates a new EditProjectAppParamCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditProjectAppParamCommandWithDefaults

`func NewEditProjectAppParamCommandWithDefaults() *EditProjectAppParamCommand`

NewEditProjectAppParamCommandWithDefaults instantiates a new EditProjectAppParamCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectAppId

`func (o *EditProjectAppParamCommand) GetProjectAppId() int32`

GetProjectAppId returns the ProjectAppId field if non-nil, zero value otherwise.

### GetProjectAppIdOk

`func (o *EditProjectAppParamCommand) GetProjectAppIdOk() (*int32, bool)`

GetProjectAppIdOk returns a tuple with the ProjectAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectAppId

`func (o *EditProjectAppParamCommand) SetProjectAppId(v int32)`

SetProjectAppId sets ProjectAppId field to given value.


### GetParameters

`func (o *EditProjectAppParamCommand) GetParameters() []EditProjectAppParamsDto`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *EditProjectAppParamCommand) GetParametersOk() (*[]EditProjectAppParamsDto, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *EditProjectAppParamCommand) SetParameters(v []EditProjectAppParamsDto)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *EditProjectAppParamCommand) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


