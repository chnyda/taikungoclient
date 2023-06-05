# Resource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**HourlyCost** | Pointer to **string** |  | [optional] 
**MonthlyCost** | Pointer to **string** |  | [optional] 
**CostComponents** | Pointer to [**[]CostComponent**](CostComponent.md) |  | [optional] 
**Subresources** | Pointer to [**[]Subresource**](Subresource.md) |  | [optional] 

## Methods

### NewResource

`func NewResource() *Resource`

NewResource instantiates a new Resource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceWithDefaults

`func NewResourceWithDefaults() *Resource`

NewResourceWithDefaults instantiates a new Resource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Resource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Resource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Resource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Resource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMetadata

`func (o *Resource) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Resource) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Resource) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Resource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetHourlyCost

`func (o *Resource) GetHourlyCost() string`

GetHourlyCost returns the HourlyCost field if non-nil, zero value otherwise.

### GetHourlyCostOk

`func (o *Resource) GetHourlyCostOk() (*string, bool)`

GetHourlyCostOk returns a tuple with the HourlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyCost

`func (o *Resource) SetHourlyCost(v string)`

SetHourlyCost sets HourlyCost field to given value.

### HasHourlyCost

`func (o *Resource) HasHourlyCost() bool`

HasHourlyCost returns a boolean if a field has been set.

### GetMonthlyCost

`func (o *Resource) GetMonthlyCost() string`

GetMonthlyCost returns the MonthlyCost field if non-nil, zero value otherwise.

### GetMonthlyCostOk

`func (o *Resource) GetMonthlyCostOk() (*string, bool)`

GetMonthlyCostOk returns a tuple with the MonthlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyCost

`func (o *Resource) SetMonthlyCost(v string)`

SetMonthlyCost sets MonthlyCost field to given value.

### HasMonthlyCost

`func (o *Resource) HasMonthlyCost() bool`

HasMonthlyCost returns a boolean if a field has been set.

### GetCostComponents

`func (o *Resource) GetCostComponents() []CostComponent`

GetCostComponents returns the CostComponents field if non-nil, zero value otherwise.

### GetCostComponentsOk

`func (o *Resource) GetCostComponentsOk() (*[]CostComponent, bool)`

GetCostComponentsOk returns a tuple with the CostComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostComponents

`func (o *Resource) SetCostComponents(v []CostComponent)`

SetCostComponents sets CostComponents field to given value.

### HasCostComponents

`func (o *Resource) HasCostComponents() bool`

HasCostComponents returns a boolean if a field has been set.

### GetSubresources

`func (o *Resource) GetSubresources() []Subresource`

GetSubresources returns the Subresources field if non-nil, zero value otherwise.

### GetSubresourcesOk

`func (o *Resource) GetSubresourcesOk() (*[]Subresource, bool)`

GetSubresourcesOk returns a tuple with the Subresources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubresources

`func (o *Resource) SetSubresources(v []Subresource)`

SetSubresources sets Subresources field to given value.

### HasSubresources

`func (o *Resource) HasSubresources() bool`

HasSubresources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


