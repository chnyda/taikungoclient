# PodDisruptionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**MinAvailable** | Pointer to **map[string]interface{}** |  | [optional] 
**MaxAvailable** | Pointer to **map[string]interface{}** |  | [optional] 
**AllowedDisruptions** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewPodDisruptionDto

`func NewPodDisruptionDto() *PodDisruptionDto`

NewPodDisruptionDto instantiates a new PodDisruptionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPodDisruptionDtoWithDefaults

`func NewPodDisruptionDtoWithDefaults() *PodDisruptionDto`

NewPodDisruptionDtoWithDefaults instantiates a new PodDisruptionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PodDisruptionDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PodDisruptionDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PodDisruptionDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PodDisruptionDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *PodDisruptionDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PodDisruptionDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PodDisruptionDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *PodDisruptionDto) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetMinAvailable

`func (o *PodDisruptionDto) GetMinAvailable() map[string]interface{}`

GetMinAvailable returns the MinAvailable field if non-nil, zero value otherwise.

### GetMinAvailableOk

`func (o *PodDisruptionDto) GetMinAvailableOk() (*map[string]interface{}, bool)`

GetMinAvailableOk returns a tuple with the MinAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAvailable

`func (o *PodDisruptionDto) SetMinAvailable(v map[string]interface{})`

SetMinAvailable sets MinAvailable field to given value.

### HasMinAvailable

`func (o *PodDisruptionDto) HasMinAvailable() bool`

HasMinAvailable returns a boolean if a field has been set.

### GetMaxAvailable

`func (o *PodDisruptionDto) GetMaxAvailable() map[string]interface{}`

GetMaxAvailable returns the MaxAvailable field if non-nil, zero value otherwise.

### GetMaxAvailableOk

`func (o *PodDisruptionDto) GetMaxAvailableOk() (*map[string]interface{}, bool)`

GetMaxAvailableOk returns a tuple with the MaxAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAvailable

`func (o *PodDisruptionDto) SetMaxAvailable(v map[string]interface{})`

SetMaxAvailable sets MaxAvailable field to given value.

### HasMaxAvailable

`func (o *PodDisruptionDto) HasMaxAvailable() bool`

HasMaxAvailable returns a boolean if a field has been set.

### GetAllowedDisruptions

`func (o *PodDisruptionDto) GetAllowedDisruptions() map[string]interface{}`

GetAllowedDisruptions returns the AllowedDisruptions field if non-nil, zero value otherwise.

### GetAllowedDisruptionsOk

`func (o *PodDisruptionDto) GetAllowedDisruptionsOk() (*map[string]interface{}, bool)`

GetAllowedDisruptionsOk returns a tuple with the AllowedDisruptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDisruptions

`func (o *PodDisruptionDto) SetAllowedDisruptions(v map[string]interface{})`

SetAllowedDisruptions sets AllowedDisruptions field to given value.

### HasAllowedDisruptions

`func (o *PodDisruptionDto) HasAllowedDisruptions() bool`

HasAllowedDisruptions returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PodDisruptionDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PodDisruptionDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PodDisruptionDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PodDisruptionDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


