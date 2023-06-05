# UpdateServerCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Ip** | **string** |  | 
**InstanceId** | Pointer to **string** |  | [optional] 
**ProviderID** | Pointer to **string** |  | [optional] 
**AwsHostName** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateServerCommand

`func NewUpdateServerCommand(id int32, ip string, ) *UpdateServerCommand`

NewUpdateServerCommand instantiates a new UpdateServerCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServerCommandWithDefaults

`func NewUpdateServerCommandWithDefaults() *UpdateServerCommand`

NewUpdateServerCommandWithDefaults instantiates a new UpdateServerCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateServerCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateServerCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateServerCommand) SetId(v int32)`

SetId sets Id field to given value.


### GetIp

`func (o *UpdateServerCommand) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *UpdateServerCommand) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *UpdateServerCommand) SetIp(v string)`

SetIp sets Ip field to given value.


### GetInstanceId

`func (o *UpdateServerCommand) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *UpdateServerCommand) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *UpdateServerCommand) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *UpdateServerCommand) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetProviderID

`func (o *UpdateServerCommand) GetProviderID() string`

GetProviderID returns the ProviderID field if non-nil, zero value otherwise.

### GetProviderIDOk

`func (o *UpdateServerCommand) GetProviderIDOk() (*string, bool)`

GetProviderIDOk returns a tuple with the ProviderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderID

`func (o *UpdateServerCommand) SetProviderID(v string)`

SetProviderID sets ProviderID field to given value.

### HasProviderID

`func (o *UpdateServerCommand) HasProviderID() bool`

HasProviderID returns a boolean if a field has been set.

### GetAwsHostName

`func (o *UpdateServerCommand) GetAwsHostName() string`

GetAwsHostName returns the AwsHostName field if non-nil, zero value otherwise.

### GetAwsHostNameOk

`func (o *UpdateServerCommand) GetAwsHostNameOk() (*string, bool)`

GetAwsHostNameOk returns a tuple with the AwsHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsHostName

`func (o *UpdateServerCommand) SetAwsHostName(v string)`

SetAwsHostName sets AwsHostName field to given value.

### HasAwsHostName

`func (o *UpdateServerCommand) HasAwsHostName() bool`

HasAwsHostName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


