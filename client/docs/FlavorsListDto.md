# FlavorsListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ram** | Pointer to **int64** |  | [optional] 
**Cpu** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **map[string]interface{}** |  | [optional] 
**MaxDataDiskCount** | Pointer to **float64** |  | [optional] 

## Methods

### NewFlavorsListDto

`func NewFlavorsListDto() *FlavorsListDto`

NewFlavorsListDto instantiates a new FlavorsListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlavorsListDtoWithDefaults

`func NewFlavorsListDtoWithDefaults() *FlavorsListDto`

NewFlavorsListDtoWithDefaults instantiates a new FlavorsListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRam

`func (o *FlavorsListDto) GetRam() int64`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *FlavorsListDto) GetRamOk() (*int64, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *FlavorsListDto) SetRam(v int64)`

SetRam sets Ram field to given value.

### HasRam

`func (o *FlavorsListDto) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetCpu

`func (o *FlavorsListDto) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *FlavorsListDto) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *FlavorsListDto) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *FlavorsListDto) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetName

`func (o *FlavorsListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlavorsListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlavorsListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FlavorsListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *FlavorsListDto) GetDescription() map[string]interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FlavorsListDto) GetDescriptionOk() (*map[string]interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FlavorsListDto) SetDescription(v map[string]interface{})`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FlavorsListDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaxDataDiskCount

`func (o *FlavorsListDto) GetMaxDataDiskCount() float64`

GetMaxDataDiskCount returns the MaxDataDiskCount field if non-nil, zero value otherwise.

### GetMaxDataDiskCountOk

`func (o *FlavorsListDto) GetMaxDataDiskCountOk() (*float64, bool)`

GetMaxDataDiskCountOk returns a tuple with the MaxDataDiskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDataDiskCount

`func (o *FlavorsListDto) SetMaxDataDiskCount(v float64)`

SetMaxDataDiskCount sets MaxDataDiskCount field to given value.

### HasMaxDataDiskCount

`func (o *FlavorsListDto) HasMaxDataDiskCount() bool`

HasMaxDataDiskCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


