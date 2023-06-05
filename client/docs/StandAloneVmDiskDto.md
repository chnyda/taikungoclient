# StandAloneVmDiskDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Size** | Pointer to **int64** |  | [optional] 
**VolumeType** | Pointer to **string** |  | [optional] 
**DeviceName** | Pointer to **string** |  | [optional] 
**LunId** | Pointer to **int32** |  | [optional] 

## Methods

### NewStandAloneVmDiskDto

`func NewStandAloneVmDiskDto(name string, ) *StandAloneVmDiskDto`

NewStandAloneVmDiskDto instantiates a new StandAloneVmDiskDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandAloneVmDiskDtoWithDefaults

`func NewStandAloneVmDiskDtoWithDefaults() *StandAloneVmDiskDto`

NewStandAloneVmDiskDtoWithDefaults instantiates a new StandAloneVmDiskDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StandAloneVmDiskDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StandAloneVmDiskDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StandAloneVmDiskDto) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *StandAloneVmDiskDto) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StandAloneVmDiskDto) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StandAloneVmDiskDto) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *StandAloneVmDiskDto) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetVolumeType

`func (o *StandAloneVmDiskDto) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *StandAloneVmDiskDto) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *StandAloneVmDiskDto) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *StandAloneVmDiskDto) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetDeviceName

`func (o *StandAloneVmDiskDto) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *StandAloneVmDiskDto) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *StandAloneVmDiskDto) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *StandAloneVmDiskDto) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetLunId

`func (o *StandAloneVmDiskDto) GetLunId() int32`

GetLunId returns the LunId field if non-nil, zero value otherwise.

### GetLunIdOk

`func (o *StandAloneVmDiskDto) GetLunIdOk() (*int32, bool)`

GetLunIdOk returns a tuple with the LunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunId

`func (o *StandAloneVmDiskDto) SetLunId(v int32)`

SetLunId sets LunId field to given value.

### HasLunId

`func (o *StandAloneVmDiskDto) HasLunId() bool`

HasLunId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


