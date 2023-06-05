# NetworkPolicyDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataName** | Pointer to **string** |  | [optional] 
**PodSelector** | Pointer to **map[string]string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Age** | Pointer to **string** |  | [optional] 

## Methods

### NewNetworkPolicyDto

`func NewNetworkPolicyDto() *NetworkPolicyDto`

NewNetworkPolicyDto instantiates a new NetworkPolicyDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPolicyDtoWithDefaults

`func NewNetworkPolicyDtoWithDefaults() *NetworkPolicyDto`

NewNetworkPolicyDtoWithDefaults instantiates a new NetworkPolicyDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataName

`func (o *NetworkPolicyDto) GetMetadataName() string`

GetMetadataName returns the MetadataName field if non-nil, zero value otherwise.

### GetMetadataNameOk

`func (o *NetworkPolicyDto) GetMetadataNameOk() (*string, bool)`

GetMetadataNameOk returns a tuple with the MetadataName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataName

`func (o *NetworkPolicyDto) SetMetadataName(v string)`

SetMetadataName sets MetadataName field to given value.

### HasMetadataName

`func (o *NetworkPolicyDto) HasMetadataName() bool`

HasMetadataName returns a boolean if a field has been set.

### GetPodSelector

`func (o *NetworkPolicyDto) GetPodSelector() map[string]string`

GetPodSelector returns the PodSelector field if non-nil, zero value otherwise.

### GetPodSelectorOk

`func (o *NetworkPolicyDto) GetPodSelectorOk() (*map[string]string, bool)`

GetPodSelectorOk returns a tuple with the PodSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodSelector

`func (o *NetworkPolicyDto) SetPodSelector(v map[string]string)`

SetPodSelector sets PodSelector field to given value.

### HasPodSelector

`func (o *NetworkPolicyDto) HasPodSelector() bool`

HasPodSelector returns a boolean if a field has been set.

### GetNamespace

`func (o *NetworkPolicyDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *NetworkPolicyDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *NetworkPolicyDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *NetworkPolicyDto) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetAge

`func (o *NetworkPolicyDto) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *NetworkPolicyDto) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *NetworkPolicyDto) SetAge(v string)`

SetAge sets Age field to given value.

### HasAge

`func (o *NetworkPolicyDto) HasAge() bool`

HasAge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


