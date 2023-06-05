# CreateStandAloneDiskCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StandaloneVmId** | **int32** |  | 
**Name** | **string** |  | 
**Size** | Pointer to **int64** |  | [optional] 
**VolumeType** | Pointer to **string** |  | [optional] 
**DeviceName** | Pointer to **string** |  | [optional] 
**LunId** | Pointer to **int32** |  | [optional] 

## Methods

### NewCreateStandAloneDiskCommand

`func NewCreateStandAloneDiskCommand(standaloneVmId int32, name string, ) *CreateStandAloneDiskCommand`

NewCreateStandAloneDiskCommand instantiates a new CreateStandAloneDiskCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStandAloneDiskCommandWithDefaults

`func NewCreateStandAloneDiskCommandWithDefaults() *CreateStandAloneDiskCommand`

NewCreateStandAloneDiskCommandWithDefaults instantiates a new CreateStandAloneDiskCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStandaloneVmId

`func (o *CreateStandAloneDiskCommand) GetStandaloneVmId() int32`

GetStandaloneVmId returns the StandaloneVmId field if non-nil, zero value otherwise.

### GetStandaloneVmIdOk

`func (o *CreateStandAloneDiskCommand) GetStandaloneVmIdOk() (*int32, bool)`

GetStandaloneVmIdOk returns a tuple with the StandaloneVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandaloneVmId

`func (o *CreateStandAloneDiskCommand) SetStandaloneVmId(v int32)`

SetStandaloneVmId sets StandaloneVmId field to given value.


### GetName

`func (o *CreateStandAloneDiskCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateStandAloneDiskCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateStandAloneDiskCommand) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *CreateStandAloneDiskCommand) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateStandAloneDiskCommand) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateStandAloneDiskCommand) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreateStandAloneDiskCommand) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetVolumeType

`func (o *CreateStandAloneDiskCommand) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *CreateStandAloneDiskCommand) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *CreateStandAloneDiskCommand) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *CreateStandAloneDiskCommand) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetDeviceName

`func (o *CreateStandAloneDiskCommand) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *CreateStandAloneDiskCommand) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *CreateStandAloneDiskCommand) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *CreateStandAloneDiskCommand) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetLunId

`func (o *CreateStandAloneDiskCommand) GetLunId() int32`

GetLunId returns the LunId field if non-nil, zero value otherwise.

### GetLunIdOk

`func (o *CreateStandAloneDiskCommand) GetLunIdOk() (*int32, bool)`

GetLunIdOk returns a tuple with the LunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunId

`func (o *CreateStandAloneDiskCommand) SetLunId(v int32)`

SetLunId sets LunId field to given value.

### HasLunId

`func (o *CreateStandAloneDiskCommand) HasLunId() bool`

HasLunId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


