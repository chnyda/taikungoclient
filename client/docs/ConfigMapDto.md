# ConfigMapDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataName** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Age** | Pointer to **string** |  | [optional] 

## Methods

### NewConfigMapDto

`func NewConfigMapDto() *ConfigMapDto`

NewConfigMapDto instantiates a new ConfigMapDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigMapDtoWithDefaults

`func NewConfigMapDtoWithDefaults() *ConfigMapDto`

NewConfigMapDtoWithDefaults instantiates a new ConfigMapDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataName

`func (o *ConfigMapDto) GetMetadataName() string`

GetMetadataName returns the MetadataName field if non-nil, zero value otherwise.

### GetMetadataNameOk

`func (o *ConfigMapDto) GetMetadataNameOk() (*string, bool)`

GetMetadataNameOk returns a tuple with the MetadataName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataName

`func (o *ConfigMapDto) SetMetadataName(v string)`

SetMetadataName sets MetadataName field to given value.

### HasMetadataName

`func (o *ConfigMapDto) HasMetadataName() bool`

HasMetadataName returns a boolean if a field has been set.

### GetNamespace

`func (o *ConfigMapDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ConfigMapDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ConfigMapDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ConfigMapDto) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetAge

`func (o *ConfigMapDto) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *ConfigMapDto) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *ConfigMapDto) SetAge(v string)`

SetAge sets Age field to given value.

### HasAge

`func (o *ConfigMapDto) HasAge() bool`

HasAge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


