# UserTokenCreateCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpireDate** | Pointer to **time.Time** |  | [optional] 
**IsReadonly** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Endpoints** | Pointer to [**[]AvailableEndpointData**](AvailableEndpointData.md) |  | [optional] 
**BindALL** | Pointer to **bool** |  | [optional] 

## Methods

### NewUserTokenCreateCommand

`func NewUserTokenCreateCommand() *UserTokenCreateCommand`

NewUserTokenCreateCommand instantiates a new UserTokenCreateCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserTokenCreateCommandWithDefaults

`func NewUserTokenCreateCommandWithDefaults() *UserTokenCreateCommand`

NewUserTokenCreateCommandWithDefaults instantiates a new UserTokenCreateCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpireDate

`func (o *UserTokenCreateCommand) GetExpireDate() time.Time`

GetExpireDate returns the ExpireDate field if non-nil, zero value otherwise.

### GetExpireDateOk

`func (o *UserTokenCreateCommand) GetExpireDateOk() (*time.Time, bool)`

GetExpireDateOk returns a tuple with the ExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDate

`func (o *UserTokenCreateCommand) SetExpireDate(v time.Time)`

SetExpireDate sets ExpireDate field to given value.

### HasExpireDate

`func (o *UserTokenCreateCommand) HasExpireDate() bool`

HasExpireDate returns a boolean if a field has been set.

### GetIsReadonly

`func (o *UserTokenCreateCommand) GetIsReadonly() bool`

GetIsReadonly returns the IsReadonly field if non-nil, zero value otherwise.

### GetIsReadonlyOk

`func (o *UserTokenCreateCommand) GetIsReadonlyOk() (*bool, bool)`

GetIsReadonlyOk returns a tuple with the IsReadonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadonly

`func (o *UserTokenCreateCommand) SetIsReadonly(v bool)`

SetIsReadonly sets IsReadonly field to given value.

### HasIsReadonly

`func (o *UserTokenCreateCommand) HasIsReadonly() bool`

HasIsReadonly returns a boolean if a field has been set.

### GetName

`func (o *UserTokenCreateCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserTokenCreateCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserTokenCreateCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserTokenCreateCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEndpoints

`func (o *UserTokenCreateCommand) GetEndpoints() []AvailableEndpointData`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *UserTokenCreateCommand) GetEndpointsOk() (*[]AvailableEndpointData, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *UserTokenCreateCommand) SetEndpoints(v []AvailableEndpointData)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *UserTokenCreateCommand) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetBindALL

`func (o *UserTokenCreateCommand) GetBindALL() bool`

GetBindALL returns the BindALL field if non-nil, zero value otherwise.

### GetBindALLOk

`func (o *UserTokenCreateCommand) GetBindALLOk() (*bool, bool)`

GetBindALLOk returns a tuple with the BindALL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindALL

`func (o *UserTokenCreateCommand) SetBindALL(v bool)`

SetBindALL sets BindALL field to given value.

### HasBindALL

`func (o *UserTokenCreateCommand) HasBindALL() bool`

HasBindALL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


