# PodDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataName** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**RestartCount** | Pointer to **int32** |  | [optional] 
**Namespace** | Pointer to **NullableString** |  | [optional] 
**Age** | Pointer to **NullableTime** |  | [optional] 
**Node** | Pointer to **NullableString** |  | [optional] 
**Phase** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPodDto

`func NewPodDto() *PodDto`

NewPodDto instantiates a new PodDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPodDtoWithDefaults

`func NewPodDtoWithDefaults() *PodDto`

NewPodDtoWithDefaults instantiates a new PodDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataName

`func (o *PodDto) GetMetadataName() string`

GetMetadataName returns the MetadataName field if non-nil, zero value otherwise.

### GetMetadataNameOk

`func (o *PodDto) GetMetadataNameOk() (*string, bool)`

GetMetadataNameOk returns a tuple with the MetadataName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataName

`func (o *PodDto) SetMetadataName(v string)`

SetMetadataName sets MetadataName field to given value.

### HasMetadataName

`func (o *PodDto) HasMetadataName() bool`

HasMetadataName returns a boolean if a field has been set.

### SetMetadataNameNil

`func (o *PodDto) SetMetadataNameNil(b bool)`

 SetMetadataNameNil sets the value for MetadataName to be an explicit nil

### UnsetMetadataName
`func (o *PodDto) UnsetMetadataName()`

UnsetMetadataName ensures that no value is present for MetadataName, not even an explicit nil
### GetStatus

`func (o *PodDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PodDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PodDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PodDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *PodDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *PodDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetRestartCount

`func (o *PodDto) GetRestartCount() int32`

GetRestartCount returns the RestartCount field if non-nil, zero value otherwise.

### GetRestartCountOk

`func (o *PodDto) GetRestartCountOk() (*int32, bool)`

GetRestartCountOk returns a tuple with the RestartCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCount

`func (o *PodDto) SetRestartCount(v int32)`

SetRestartCount sets RestartCount field to given value.

### HasRestartCount

`func (o *PodDto) HasRestartCount() bool`

HasRestartCount returns a boolean if a field has been set.

### GetNamespace

`func (o *PodDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PodDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PodDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *PodDto) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *PodDto) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *PodDto) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetAge

`func (o *PodDto) GetAge() time.Time`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *PodDto) GetAgeOk() (*time.Time, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *PodDto) SetAge(v time.Time)`

SetAge sets Age field to given value.

### HasAge

`func (o *PodDto) HasAge() bool`

HasAge returns a boolean if a field has been set.

### SetAgeNil

`func (o *PodDto) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *PodDto) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
### GetNode

`func (o *PodDto) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *PodDto) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *PodDto) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *PodDto) HasNode() bool`

HasNode returns a boolean if a field has been set.

### SetNodeNil

`func (o *PodDto) SetNodeNil(b bool)`

 SetNodeNil sets the value for Node to be an explicit nil

### UnsetNode
`func (o *PodDto) UnsetNode()`

UnsetNode ensures that no value is present for Node, not even an explicit nil
### GetPhase

`func (o *PodDto) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *PodDto) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *PodDto) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *PodDto) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### SetPhaseNil

`func (o *PodDto) SetPhaseNil(b bool)`

 SetPhaseNil sets the value for Phase to be an explicit nil

### UnsetPhase
`func (o *PodDto) UnsetPhase()`

UnsetPhase ensures that no value is present for Phase, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


