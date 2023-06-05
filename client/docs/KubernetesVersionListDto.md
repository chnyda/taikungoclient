# KubernetesVersionListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** |  | [optional] 
**IsKubevapEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewKubernetesVersionListDto

`func NewKubernetesVersionListDto() *KubernetesVersionListDto`

NewKubernetesVersionListDto instantiates a new KubernetesVersionListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesVersionListDtoWithDefaults

`func NewKubernetesVersionListDtoWithDefaults() *KubernetesVersionListDto`

NewKubernetesVersionListDtoWithDefaults instantiates a new KubernetesVersionListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *KubernetesVersionListDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KubernetesVersionListDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KubernetesVersionListDto) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *KubernetesVersionListDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetIsKubevapEnabled

`func (o *KubernetesVersionListDto) GetIsKubevapEnabled() bool`

GetIsKubevapEnabled returns the IsKubevapEnabled field if non-nil, zero value otherwise.

### GetIsKubevapEnabledOk

`func (o *KubernetesVersionListDto) GetIsKubevapEnabledOk() (*bool, bool)`

GetIsKubevapEnabledOk returns a tuple with the IsKubevapEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKubevapEnabled

`func (o *KubernetesVersionListDto) SetIsKubevapEnabled(v bool)`

SetIsKubevapEnabled sets IsKubevapEnabled field to given value.

### HasIsKubevapEnabled

`func (o *KubernetesVersionListDto) HasIsKubevapEnabled() bool`

HasIsKubevapEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


